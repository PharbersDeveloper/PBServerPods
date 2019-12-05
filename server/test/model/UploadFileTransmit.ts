// import { JsonConvert, OperationMode, ValueCheckingMode } from "json2typescript"
import { slow, suite, test, timeout } from "mocha-typescript"
import phLogger from "../../src/logger/phLogger"
import mongoose = require("mongoose")
import PhLogger from "../../src/logger/phLogger"
import { ObjectId } from "bson"

import FileDetail from "../../src/models/FileDetail"
import FileVersion from "../../src/models/FileVersion"
import SandboxIndex from "../../src/models/SandboxIndex"

import Asset from "../../src/models/Asset"
import File from "../../src/models/File"
import DataSet from "../../src/models/DataSet"

@suite(timeout(1000 * 60), slow(2000))
class UploadFileTransmit {

    public static before() {
        PhLogger.info(`before starting the test`)
        mongoose.connect("mongodb://127.0.0.1:27017/pharbers-sandbox-4")
    }

    public static after() {
        PhLogger.info(`after starting the test`)
        mongoose.disconnect()
    }

    @test public async fileTransmit() {
        PhLogger.info(`start trans data from old version`)
        await this.fileTransmitImpl()
    }

    public async fileTransmitImpl() {
        const sim = new SandboxIndex().getModel()
        const fdm = new FileDetail().getModel()
        const fvm = new FileVersion().getModel()

        const fm = new File().getModel()
        const dsm = new DataSet().getModel()
        const am = new Asset().getModel()

        const contents = await sim.find({})
        // @ts-ignore
        await Promise.all(contents.map( async (content) => {
            const owner = content.account
            const filesIds = content.files
            // @ts-ignore
            await Promise.all(filesIds.map( async (id) => {
                const fd = await fdm.findOne({
                    _id: id
                } )
                const fvId = fd.versions[0]
                const fv = await fvm.findOne({
                    _id: fvId
                } )

                /**
                 * 1. 将FileDetail, FileVersion转成File
                 */
                const f = new File()
                f.url = fv.where
                f.size = fv.size
                f.fileName = fd.name
                f.extension = fd.extension
                f.uploaded = fd.created
                const fc = await fm.create(f)

                /**
                 * 2. 将JobID 创建出来的DataSet MetaData化
                 */
                const jIds = fd.jobIds
                // @ts-ignore
                const dfs = await Promise.all(jIds.map( async (jid) => {
                    const ds = new DataSet()
                    // ds.jobId = jid
                    return await dsm.create(ds)
                } ) )

                /**
                 * 3. 将用户上传的内容，抽象成平台所需要的Assents
                 */
                const asset = new Asset()
                asset.name= fd.name
                asset.description= fd.name
                asset.dataType = "file"
                asset.file = fc
                // @ts-ignore
                asset.dfs = dfs
                asset.owner = fd.ownerID
                asset.accessibility = "w"
                asset.version = 0

                /**
                 * 4. 为数据添加tags
                 */
                if (fd.name.indexOf("_") > 0) {
                    phLogger.info("cpa gyc data")
                    this.cpa_gyc_name_2_tags(fd.name, asset)
                } else {
                    phLogger.info("chc data")
                    this.chc_name_2_tags(fd.name, asset)
                }

                await am.create(asset)
            } ))
        } ))
        // phLogger.info(await tmp[0])
    }

    private cpa_gyc_name_2_tags(name: string, asset: Asset) {
        const tags = name.split("_")
        if (tags.length < 4) {
            if (name.indexOf("Lilly") !== -1) {
                asset.providers = [tags[0]]
                asset.dataCover = [tags[1], tags[2]]
            } else if (name.indexOf("cpa") !== -1 || name.indexOf("gyc") !== -1) {
                asset.providers = [tags[0]]
            } else {
                phLogger.info(name)
                this.chc_name_2_tags(name, asset)
            }
        } else {
            asset.providers = [tags[0], tags[3]]
            asset.dataCover = [tags[1], tags[2]]
        }
    }

    private chc_name_2_tags(name: string, asset: Asset) {
        phLogger.info(name)
        const geo = []
        if (name.indexOf("北京") !== -1) {
            geo.push("北京")
        }

        if (name.indexOf("上海") !== -1) {
            geo.push("上海")
        }

        if (name.indexOf("安徽") !== -1) {
            geo.push("安徽")
        }

        if (name.indexOf("山东") !== -1) {
            geo.push("山东")
        }

        if (name.indexOf("广州") !== -1) {
            geo.push("广州")
        }

        if (name.indexOf("福建") !== -1) {
            geo.push("福建")
        }

        if (name.indexOf("江苏") !== -1) {
            geo.push("江苏")
        }
        asset.geoCover = geo

        if (name.indexOf("【") !== -1) {
            const start = name.indexOf("【")
            const end = name.indexOf("】")
            const length = end - start - 1
            const provider = name.substr(start + 1, length)
            asset.providers = [provider]
        }
    }
}
