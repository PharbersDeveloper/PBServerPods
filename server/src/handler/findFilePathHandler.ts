"use strict"
import PhLogger from "../logger/phLogger"
import Asset from "../models/Asset"
import File from "../models/File"

export class FindFilePathHandler {
    async findFilePathWithTraceId(body: any) {
        const asset = await new Asset().getModel().findOne({traceId: body.traceId})
        const file = await new File().getModel().findById(asset.file)
        return {"ossPath": file.url}
    }

    async findFilePathWithJobId(body: any) {
        return {}
    }
}