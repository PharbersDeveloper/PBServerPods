"use strict"
import Asset from "../models/Asset"
import PhLogger from "../logger/phLogger"
import mongoose = require("mongoose")
import Mart from "../models/Mart"

export class AssetDataMartHandler {
    async assetDataMart(body: any) {
        PhLogger.info("进入新增DataMart")

        const martModel = new Mart()
        const assetModel = new Asset()


        martModel.name = body.martName
        martModel.dfs = body.dfs.map((id: string) => new mongoose.mongo.ObjectId(id))
        martModel.url = body.martUrl
        martModel.dataType = body.martDataType
        const mart =  await new Mart().getModel().create(martModel)

        assetModel._id = new mongoose.mongo.ObjectId()
        assetModel.name = body.assetName
        assetModel.description = body.assetDescription
        assetModel.version = body.assetVersion
        assetModel.dataType = body.assetDataType
        assetModel.providers = body.providers
        assetModel.markets = body.markets
        assetModel.molecules = body.molecules
        assetModel.dataCover = body.dataCover
        assetModel.geoCover = body.geoCover
        assetModel.labels = body.labels
        assetModel.mart = mart
        assetModel.accessibility = "w"
        assetModel.isNewVersion = true
        assetModel.createTime = new Date().getTime()
        await new Asset().getModel().create(assetModel)

        return {status: "ok"}
    }
}