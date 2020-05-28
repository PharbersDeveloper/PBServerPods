"use strict"

import DataSet from "../models/DataSet"
import Job from "../models/Job"
import mongoose = require("mongoose")
import Asset from "../models/Asset"

/**
 * 血缘数据处理，针对dfs、job与asset的关联
 */
// TODO 整体需要重构 ！！Alex留
export default class JobBloodHandler {
    // TODO 有时间改写法，现在先这样，这个暂时给老邓那边用  //body.description || result.description
    async createDataSetsAndJob(body: any) {
        const jobRes = await new Job().getModel().findOne({jobContainerId: body.jobId})
        if (jobRes) {
            const dsRes = await new DataSet().getModel().findOne({job: jobRes.id || "", url: ""})
            if (dsRes) {
                dsRes.colNames = body.columnNames || dsRes.colNames
                dsRes.length = body.length || dsRes.length
                dsRes.tabName = body.tabName || dsRes.tabName
                dsRes.url = body.url || dsRes.url
                dsRes.status = body.status || dsRes.status
                await dsRes.save()
            }
            return {status: "ok"}
        } else {
            const datasetModel = new DataSet()
            const jobModel = new Job()
            jobModel.jobContainerId = body.jobId
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

        // const result = await new DataSet().getModel().findById(new mongoose.mongo.ObjectId(body.mongoId))
        // if (!result) {
        //     const model = new DataSet()
        //     const jobModel = new Job()
        //     jobModel.jobContainerId = body.jobId
        //     jobModel.create = new Date().getTime()
        //     const job = await new Job().getModel().create(jobModel)
        //
        //     model._id = new mongoose.mongo.ObjectId(body.mongoId)
        //     model.parent = body.parentIds
        //     model.colNames = body.colName
        //     model.length = body.length
        //     model.tabName = body.tabName
        //     model.url = body.url
        //     model.description = body.description
        //     model.status = body.status
        //     model.job = job
        //     const ds = await new DataSet().getModel().create(model)
        //     const asset = await new Asset()
        //         .getModel().findOne({_id: new mongoose.mongo.ObjectId(body.assetId), isNewVersion: true})
        //     asset.dfs = asset.dfs.concat(ds)
        //     await asset.save()
        // } else {
        //     result.parent = body.parentIds || result.parent
        //     result.colNames = body.colName || result.colNames
        //     result.length = body.length || result.length
        //     result.tabName = body.tabName || result.tabName
        //     result.url = body.url || result.url
        //     result.description = result.description
        //     result.status = body.status || result.status
        //     await result.save()
        //     const asset = await new Asset()
        //         .getModel().findOne({_id: new mongoose.mongo.ObjectId(body.assetId), isNewVersion: true})
        //     if (!asset.dfs.map((item) => item.toString()).includes(result.id)) {
        //         asset.dfs = asset.dfs.concat(result)
        //     }
        //     await asset.save()
        // }
    }

    async initJobs(body: any) {

        async function push(j: Job) {
            const dsModel = new DataSet()
            dsModel._id = new mongoose.mongo.ObjectId(body.mongoId)
            dsModel.parent = body.parentIds
            dsModel.colNames = body.colName
            dsModel.length = body.length
            dsModel.tabName = body.tabName
            dsModel.url = body.url
            dsModel.description = body.description
            dsModel.status = body.status
            dsModel.job = j
            const ds = await new DataSet().getModel().create(dsModel)
            const asset = await new Asset()
                .getModel().findOne({_id: new mongoose.mongo.ObjectId(body.assetId), isNewVersion: true})
            asset.dfs = asset.dfs.concat(ds)
            await asset.save()
        }

        const jobRes = await new Job().getModel().findOne({jobContainerId: body.jobId})

        if (jobRes) {
            await push(jobRes)
        } else {
            const jobModel = new Job()
            jobModel.jobContainerId = body.jobId
            jobModel.create = new Date().getTime()
            const job = await new Job().getModel().create(jobModel)
            await push(job)
        }

    }

    async pushDs(body: any) {
        const jobRes = await new Job().getModel().findOne({jobContainerId: body.jobId})

        const dsRes = await new DataSet().getModel().findOne({job: jobRes.id}).or([{status: "pending"},{status: "start"}])
        if (dsRes) {
            dsRes.colNames = body.columnNames || dsRes.colNames
            dsRes.length = body.length || dsRes.length
            dsRes.tabName = body.tabName || dsRes.tabName
            dsRes.url = body.url || dsRes.url
            dsRes.description = body.description || dsRes.description
            dsRes.status = body.status || dsRes.status
            await dsRes.save()
            return {
                "dsId": dsRes.id,
                "status": dsRes.status,
                "path": dsRes.url,
                "description": dsRes.description
            }
        } else {
            return {
                "dsId": "-1",
                "status": "error"
            }
        }


        // if (jobRes) {
        //     const dsRes = await new DataSet().getModel().findOne({job: jobRes.id})
        //     dsRes.colNames = body.colName || dsRes.colNames
        //     dsRes.length = body.length || dsRes.length
        //     dsRes.tabName = body.tabName || dsRes.tabName
        //     dsRes.url = body.url || dsRes.url
        //     dsRes.description = body.description || dsRes.description
        //     dsRes.status = body.status || dsRes.status
        //     await dsRes.save()
        //     const asset = await new Asset().getModel().findOne({isNewVersion: true}).in("dfs", [dsRes.id])
        //     if (!asset.dfs.map((item) => item.toString()).includes(dsRes.id)) {
        //         asset.dfs = asset.dfs.concat(dsRes)
        //     }
        //     await asset.save()
        // } else {
        //     const dsModel = new DataSet()
        //     const jobModel = new Job()
        //     jobModel.jobContainerId = body.jobId
        //     jobModel.create = new Date().getTime()
        //     const job = await new Job().getModel().create(jobModel)
        //
        //     dsModel._id = new mongoose.mongo.ObjectId(body.mongoId)
        //     dsModel.parent = body.parentIds
        //     dsModel.colNames = body.colName
        //     dsModel.length = body.length
        //     dsModel.tabName = body.tabName
        //     dsModel.url = body.url
        //     dsModel.description = body.description
        //     dsModel.status = body.status
        //     dsModel.job = job
        //     const ds = await new DataSet().getModel().create(dsModel)
        //     const asset = await new Asset()
        //         .getModel().findOne({_id: new mongoose.mongo.ObjectId(body.assetId), isNewVersion: true})
        //     asset.dfs = asset.dfs.concat(ds)
        //     await asset.save()
        // }
    }

    // 根据当前的DsId 查找下一个JobId（在没有DAG前）
    async findNextJob(body: any) {
        const dsRes = await new DataSet().getModel().findOne({"description": body.jobName}).in("parent", [body.dsId])
        if (dsRes) {
           const job = await new Job().getModel().findById(dsRes.job)
           return {
               "jobId": job.jobContainerId
           }
        } else {
            return {
                "jobId": "-1"
            }
        }
    }
}