"use strict"

import DataSet from "../models/DataSet"
import Job from "../models/Job"
import mongoose = require("mongoose")

/**
 * 血缘数据处理，针对dfs、job与asset的关联
 */

export default class JobBloodHandler {

    async createDataSetsAndJob(body: any) {
        const result = await new DataSet().getModel().findById(new mongoose.mongo.ObjectId(body.mongoId))
        if (!result) {
            const model = new DataSet()
            const jobModel = new Job()
            jobModel.jobContainerId = body.jobContainerId
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
            await new DataSet().getModel().create(model)
        } else {
            result.parent = body.parentIds
            result.colNames = body.colName
            result.length = body.length
            result.tabName = body.tabName
            result.url = body.url
            result.description = body.description
            result.status = body.status
            result.save()
        }
        return {status: "ok"}
    }
}