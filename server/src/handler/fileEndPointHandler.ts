"use strict"
import Asset from "../models/Asset"
import PhLogger from "../logger/phLogger"
import DataSet from "../models/DataSet"
import mongoose = require("mongoose")

export default class FileEndPointHandler {
    async uploadFileEnd(body: any) {
        // TODO: 尝试10次，没有就出大问题，应该通知错误处理，但现在没有
        let count = 10
        PhLogger.info("进入修改DS")

        function sleep(ms: number){
            return new Promise((resolve)=>setTimeout(resolve,ms))
        }

        async function getDs() {
            const ds = await new DataSet().getModel().findById(new mongoose.mongo.ObjectId(body.dataSetId))

            if (ds !== null) {
                PhLogger.info("ObjectId AssetId ======> " + body.assetId)
                const asset = await new Asset()
                    .getModel().findOne({_id: new mongoose.mongo.ObjectId(body.assetId), isNewVersion: true})
                if (!asset.dfs.map((item) => item.toString()).includes(body.dataSetId)) {
                    asset.dfs = asset.dfs.concat(ds)
                }
                await asset.save()
                return {status: "ok"}
            } else {
                await sleep(1000)
                PhLogger.info(count)
                count--
                if (count <= 0 ) {
                    PhLogger.info("DataSet 数据库中未存在，id为 =====> " + body.dataSetId)
                    return {status: "no"}
                }
                await getDs()
            }
        }
        return getDs()
    }
}