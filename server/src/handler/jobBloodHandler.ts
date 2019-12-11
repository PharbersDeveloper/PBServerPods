"use strict"

import DataSet from "../models/DataSet"
import Job from "../models/Job"
import mongoose = require("mongoose")

export class JobBloodHandler {

    async createDataSetsAndJob(body: any) {
        const datasetModel = new DataSet()
        const jobModel = new Job()
        jobModel.jobContainerId = body.jobContainerId
        jobModel.create = new Date().getTime()
        const job = await new Job().getModel().create(jobModel)

        // tslint:disable-next-line:no-console
        console.info(body)
        datasetModel._id = new mongoose.mongo.ObjectId(body.mongoId)
        datasetModel.parent = body.parent
        datasetModel.colNames = body.colNames
        datasetModel.length = body.length
        datasetModel.tabName = body.tabName
        datasetModel.url = body.url
        datasetModel.description = body.description
        datasetModel.job = job

        await new DataSet().getModel().create(datasetModel)
        return {status: "ok"}
        // const datasetModel = new DataSet()
        // const jobModel = new Job()
        // jobModel.jobContainerId = body.jobContainerId
        // jobModel.create = new Date().getTime()
        // return await new Job().getModel().create(jobModel).then(async job => {
        //     datasetModel._id = new mongoose.mongo.ObjectId(body.mongoId)
        //     datasetModel.parent = body.parent
        //     datasetModel.colNames = body.colNames
        //     datasetModel.length = body.length
        //     datasetModel.tabName = body.tabName
        //     datasetModel.url = body.url
        //     datasetModel.description = body.description
        //     datasetModel.job = job
        //
        //     await new DataSet().getModel().create(datasetModel)
        //     return {status: "ok"}
        // })
    }
}