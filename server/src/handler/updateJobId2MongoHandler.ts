"use strict"
import Asset from "../models/Asset"
import PhLogger from "../logger/phLogger"
import DataSet from "../models/DataSet"
import mongoose = require("mongoose")

export class UpdateJobId2MongoHandler {
    async uploadFileEnd(body: any) {
        PhLogger.info("进入修改DS")
        // const ds = await new DataSet().getModel().findById(new mongoose.mongo.ObjectId(body.dataSetId))
        // const asset = await new Asset().getModel().findOne({_id: new mongoose.mongo.ObjectId(body.assetId), isNewVersion: true})
        // asset.dfs = asset.dfs.concat(ds)
        // await asset.save()

        function sleep(ms: number){
            return new Promise((resolve)=>setTimeout(resolve,ms))
        }

        function getDs() {
            return new DataSet()
                .getModel().findById(new mongoose.mongo.ObjectId(body.dataSetId)).then(async ds => {
                if (ds !== null) {
                    const asset = await new Asset()
                        .getModel().findOne({_id: new mongoose.mongo.ObjectId(body.assetId), isNewVersion: true})
                    asset.dfs = asset.dfs.concat(ds)
                    await asset.save()
                    return {status: "ok"}
                } else {
                    PhLogger.info("DS Is Null，进入等待")
                    await sleep(1000)
                    await getDs()
                }
            })
        }
        return getDs()
    }
}