"use strict"
import Asset from "../models/Asset"
import File from "../models/File"
import mongoose = require("mongoose")

/**
 * 查询File的在oss中的地址
 * 使用者  Convert Excel的项目中使用
 */
export default class FindFilePathHandler {
    async findFilePathWithId(body: any) {
        const asset = await new Asset().getModel().findById(new mongoose.mongo.ObjectId(body.assetId))
        const file = await new File().getModel().findById(asset.file)
        return {"ossPath": file.url}
    }
    // async findFilePathWithJobId(body: any) {
    //     return {}
    // }
}