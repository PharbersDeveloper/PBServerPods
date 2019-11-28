"use strict"
import Asset from "../models/Asset"
import DataSet from "../models/DataSet"
import mongoose = require("mongoose")

export class UpdateJobId2MongoHandler {
    async uploadFileEnd(body: any) {
        const ds = await new DataSet().getModel().findById(new mongoose.mongo.ObjectId(body.dataSetId))

        const asset = await new Asset().getModel().findOne({traceId: body.traceId, isNewVersion: true})
        asset.dfs = asset.dfs.concat(ds)
        await asset.save()
        return {status: "ok"}
    }
}