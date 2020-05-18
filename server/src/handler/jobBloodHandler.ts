"use strict"

import DataSet from "../models/DataSet"
import Job from "../models/Job"
import mongoose = require("mongoose")
import Asset from "../models/Asset"

/**
 * 血缘数据处理，针对dfs、job与asset的关联
 */

export default class JobBloodHandler {
    // TODO 有时间改写法，现在先这样
    async createDataSetsAndJob(body: any) {
        const result = await new DataSet().getModel().findById(new mongoose.mongo.ObjectId(body.mongoId))
        if (!result) {
            const model = new DataSet()
            const jobModel = new Job()
            jobModel.jobContainerId = body.jobId
            jobModel.create = new Date().getTime()
            const job = await new Job().getModel().create(jobModel)

            model._id = new mongoose.mongo.ObjectId(body.mongoId)
            model.parent = body.parentIds
            model.colNames = body.colName
            model.length = body.length
            model.tabName = body.tabName
            model.url = body.url
            model.description = body.description
            model.status = body.status
            model.job = job
            const ds = await new DataSet().getModel().create(model)
            const asset = await new Asset()
                .getModel().findOne({_id: new mongoose.mongo.ObjectId(body.assetId), isNewVersion: true})
            asset.dfs = asset.dfs.concat(ds)
            await asset.save()
        } else {
            result.parent = body.parentIds || result.parent
            result.colNames = body.colName || result.colNames
            result.length = body.length || result.length
            result.tabName = body.tabName || result.tabName
            result.url = body.url || result.url
            result.description = body.description || result.description
            result.status = body.status || result.status
            await result.save()
            const asset = await new Asset()
                .getModel().findOne({_id: new mongoose.mongo.ObjectId(body.assetId), isNewVersion: true})
            if (!asset.dfs.map((item) => item.toString()).includes(result.id)) {
                asset.dfs = asset.dfs.concat(result)
            }
            await asset.save()
        }

        return {status: "ok"}
    }
}