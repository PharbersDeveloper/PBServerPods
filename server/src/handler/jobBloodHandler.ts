"use strict"

import DataSet from "../models/DataSet"
import Job from "../models/Job"
import mongoose = require("mongoose")

export default class JobBloodHandler {

    async createDataSetsAndJob(body: any) {
        const datasetModel = new DataSet()
        const jobModel = new Job()
        jobModel.jobContainerId = body.jobContainerId
        jobModel.create = new Date().getTime()
        const job = await new Job().getModel().create(jobModel)

        datasetModel._id = new mongoose.mongo.ObjectId(body.mongoId)
        datasetModel.parent = body.parentIds
        datasetModel.colNames = body.colName
        datasetModel.length = body.length
        datasetModel.tabName = body.tabName
        datasetModel.url = body.url
        datasetModel.description = body.description
        datasetModel.job = job

        const ds = await new DataSet().getModel().create(datasetModel)
        if (ds != null) {
            return {status: "ok"}
        } else {
            return {status: "no"}
        }
    }
}