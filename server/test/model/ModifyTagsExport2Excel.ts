// import { JsonConvert, OperationMode, ValueCheckingMode } from "json2typescript"
import { slow, suite, test, timeout } from "mocha-typescript"
import phLogger from "../../src/logger/phLogger"
import mongoose = require("mongoose")
import PhLogger from "../../src/logger/phLogger"
import XLSX = require("xlsx")
import uuidv4 from "uuid/v4"
import { ObjectId } from "bson"

import Asset from "../../src/models/Asset"
import File from "../../src/models/File"
import DataSet from "../../src/models/DataSet"

@suite(timeout(1000 * 60), slow(2000))
class ModifyTagsExport2Excel {
    public static before() {
        PhLogger.info(`before starting the test`)
        mongoose.connect("mongodb://pharbers.com:5555/pharbers-sandbox-4")
    }

    public static after() {
        PhLogger.info(`after starting the test`)
        mongoose.disconnect()
    }

    @test public async ModifyTagsExport() {
        PhLogger.info(`start export data to excel`)

        /**
         * 1. 先将Assets数据
         */
        const fm = new File().getModel()
        const dsm = new DataSet().getModel()
        const am = new Asset().getModel()

        const allAm = await am.find({})
        phLogger.info(allAm.length)

        /**
         * 1.1 确定headers
         */
        const headers = [
            [
                "_id", "traceId", "文件名", "文件描述", "所有人ID", "资产权限",
                "资产版本", "资产类型", "资产客户来源", "资产数据来源", "资产覆盖市场", "资产覆盖分子",
                "资产覆盖时间", "资产覆盖区域", "其他标签", "文件本身属性", "文件名", "文件类型",
                "文件上传时间", "线上文件地址", "文件大小"
            ]
        ]

        const result = await Promise.all(allAm.map ( async (asset) => {
            let sourceType = this.concatArray2String(asset.providers).split(",")
            const source:string[] = []

            if (sourceType.indexOf("CPA") !== -1) {
                source.push("CPA")
            } else if (sourceType.indexOf("GYC") !== -1) {
                source.push("GYC")
            } else {
                source.push("CHC")
            }
            sourceType = sourceType.filter(x => x !== "CPA" && x !== "GYC")

            /**
             * 2. 对每个Asset读出其对应的文件
             */
            const file = await fm.findOne({
                "_id": asset.file
            })

            return [
                asset._id.toString(),
                asset.traceId,
                asset.name,
                asset.description,
                asset.owner.toString(),
                asset.accessibility,
                asset.version,
                asset.dataType,
                this.concatArray2String(sourceType),
                this.concatArray2String(source),
                this.concatArray2String(asset.markets),
                this.concatArray2String(asset.molecules),
                this.modifyDateFormat(this.concatArray2String(asset.dataCover)),
                this.concatArray2String(asset.geoCover),
                this.concatArray2String(asset.labels),
                "",     // 文件本身属性数据分割
                file.fileName,
                file.extension,
                this.timestamp2String(file.uploaded),
                file.url,
                this.bytes2MB(file.size)
            ]
        } ))
        const data = headers.concat(result)

        const workbook = XLSX.utils.book_new()
        const worksheet = XLSX.utils.aoa_to_sheet(data)

        workbook.Props = {
            Author: "Alfred Yang",
            CreatedDate: new Date(),
            Subject: "Asset Tag Export",
            Title: "Asset Tag Export"
        }

        XLSX.utils.book_append_sheet(workbook, worksheet, "Asset Tag Export")
        XLSX.writeFile(workbook, process.env.PH_TS_SANDBOX_HOME + "/export/" + uuidv4() + ".xlsx")
    }

    private concatArray2String(arr: any[], mid: string = ","): string {
        let result = ""
        arr.forEach(item => {
            if (result.length <= 0)
                result = this.cleanDataSuffix(item.toString())
            else
                result = result + mid + this.cleanDataSuffix(item.toString())
        } )
        return result
    }

    private cleanDataSuffix(str: string): string {
        if (str.indexOf(".") !== -1)
            return str.substr(0, str.indexOf("."))
        else
            return str
    }

    private modifyDateFormat(str: string): string {
        const arr = str.split(",")
        const tmp: number[] = arr.map((item) => {
            if (!isNaN(parseInt(item, 10)))
                return +("20" + item)
            else
                return NaN
        } )
        const result = tmp.filter(x => !isNaN(x))
        return this.concatArray2String(result)
    }

    private timestamp2String(stamp: number): string {
        const date = new Date()
        date.setTime(stamp)
        return date.toUTCString()
    }

    private bytes2MB(size: number): string {
        const oneMb = 1024 * 1024
        return (Math.floor((size / oneMb) * 100) / 100) + "M"
    }
}
