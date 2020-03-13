"use strict"
import PhLogger from "../logger/phLogger"
import Asset from "../models/Asset"
import File from "../models/File"
import mongoose = require("mongoose")

export class FindFilePathHandler {
    async findFilePathWithId(body: any) {
        const asset = await new Asset().getModel().findById(new mongoose.mongo.ObjectId(body.assetId))
        const file = await new File().getModel().findById(asset.file)
        return {"ossPath": file.url}
    }

    async findFilePathWithJobId(body: any) {
        return {}
    }
}