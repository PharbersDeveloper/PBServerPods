"use strict"
import PhLogger from "../logger/phLogger"
import {CONFIG} from "../shared/config"
import * as mongoose from "mongoose"
import {Request, Response} from "express"
import phLogger from "../logger/phLogger"
import {NextFunction} from "express-serve-static-core"
import Asset from "../models/Asset"
import DataSet from "../models/DataSet"

export class UpdateJobId2MongoHandler {
    constructor() {
        PhLogger.info("凸(艹皿艹 )")
    }
    // TODO: 未做异常处理
    async createJobId2Datasets(body: any) {
        const assets = await new Asset().getModel().findOne({traceId: body.traceId})
        const ds = new DataSet()
        ds.jobId = body.jobId

        const dfs = await new DataSet().getModel().create(ds)

        assets.dfs = assets.dfs.concat(dfs)
        await assets.save()

        return {"status": "ok"}
    }

    async updateJobId2Datasets(body: any) {
        const jobId = body.jobId
        const ds = await new DataSet().getModel().findOne({jobId})
        ds.colNames = body.columnName
        ds.length = body.length
        ds.url = body.path
        ds.description = body.sheetName
        await ds.save()
    }
}