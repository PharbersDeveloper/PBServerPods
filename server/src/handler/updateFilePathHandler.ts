"use strict"
import PhLogger from "../logger/phLogger"
import Asset from "../models/Asset"
import File from "../models/File"
import mongoose = require("mongoose")

export default class UpdateFilePathHandler {
    async updateAssetVersion(body: any) {

        function convertVersion(version: string) {
            const nums = version.split(".")
            const lastNum = Number(nums.pop()) + 1
            nums.concat(lastNum.toString())
            return nums.join(".")
        }

        PhLogger.info("更新Asset版本")
        // 上一个版本的历史
        const preAssetVersion = await new Asset().getModel().findById(new mongoose.mongo.ObjectId(body.assetId))
        preAssetVersion.isNewVersion = false
        await preAssetVersion.save()

        const preFileVersion = await new File().getModel().findById(preAssetVersion.file)
        const file = new File()
        file.fileName = preFileVersion.fileName
        file.extension = preFileVersion.extension === "xls" ? "xlsx" : preFileVersion.extension
        file.size = preFileVersion.size
        file.url = body.url
        const fileModel = await new File().getModel().create(file)

        const asset = new Asset()
        asset.name = preAssetVersion.name
        asset.owner = preAssetVersion.owner
        asset.accessibility = preAssetVersion.accessibility
        asset.version = convertVersion(preAssetVersion.version)
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
        asset.createTime = new Date().getTime()

        const assetResult = await new Asset().getModel().create(asset)
        return {"status": "ok", "assetId": assetResult.id.toString()}
    }
}