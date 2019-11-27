"use strict"
import Asset from "../models/Asset"
import DataSet from "../models/DataSet"
import Job from "../models/Job"

export class UpdateJobId2MongoHandler {
    async updateJobId2DataSets(body: any) {
        const job = await new Job().getModel().findOne({jobId: body.jobId})
        const ds = await new DataSet().getModel().findOne({job: job.id})

        const asset = await new Asset().getModel().findOne({traceId: body.traceId})
        asset.dfs = asset.dfs.concat(ds)
        await asset.save()
        return {status: "ok"}
    }
}