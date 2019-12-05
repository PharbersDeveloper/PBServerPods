"use strict"
import PhLogger from "../logger/phLogger"
import Asset from "../models/Asset"
import File from "../models/File"
import mongoose = require("mongoose")

export class UpdateFilePathHandler {
    constructor() {
        PhLogger.info("凸(艹皿艹 )")
    }
    // TODO: 未做异常处理
    async updateAssetVersion(body: any) {
        // 上一个版本的历史
        const preAssetVersion = await new Asset().getModel().findById(new mongoose.mongo.ObjectId(body.assetId))
        preAssetVersion.isNewVersion = false
        await preAssetVersion.save()

        const preFileVersion = await new File().getModel().findById(preAssetVersion.file)
        const file = new File()
        file.fileName = preFileVersion.fileName
        file.extension = preFileVersion.extension
        file.uploaded = new Date().getTime()
        file.size = preFileVersion.size
        file.url = body.url
        const fileModel = await new File().getModel().create(file)

        const asset = new Asset()
        asset._id = new mongoose.mongo.ObjectId()
        asset.name = preAssetVersion.name
        asset.owner = preAssetVersion.owner
        asset.accessibility = preAssetVersion.accessibility
        asset.version = preAssetVersion.version + 1
        asset.isNewVersion = true
        asset.dataType = preAssetVersion.dataType
        asset.providers = preAssetVersion.providers
        asset.markets = preAssetVersion.markets
        asset.molecules = preAssetVersion.molecules
        asset.dataCover = preAssetVersion.dataCover
        asset.geoCover = preAssetVersion.geoCover
        asset.labels = preAssetVersion.labels
        asset.file = fileModel
        asset.dfs = preAssetVersion.dfs
        asset.description = preAssetVersion.description

        await new Asset().getModel().create(asset)
        return {"status": "ok", "assetId": asset._id.toString()}
    }
}