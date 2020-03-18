"use strict"
import Asset from "../models/Asset"
import PhLogger from "../logger/phLogger"
import mongoose = require("mongoose")
import Mart from "../models/Mart"
import DataSet from "../models/DataSet"

export class AssetDataMartHandler {
    // TODO Alex自己留，记得重构
    async assetDataMart(body: any) {
        PhLogger.info("进入DataMart")

        async function createDataMart() {
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
        }

        const assetResult = await new Asset().getModel().
            find({"name": body.assetName, "isNewVersion": true}).
            sort({"createTime": -1}).skip(0).limit(1)

        if (body.saveMode === "append") {
            if (assetResult.length > 0) {
                const latestAsset = assetResult[0]
                const mart = await new Mart().getModel().findById(latestAsset.mart)
                const dsIds = body.dfs.map((id: string) => new mongoose.mongo.ObjectId(id))
                const ds = await new DataSet().getModel().find({"_id": {$in: dsIds}})
                mart.dfs = mart.dfs.concat(ds)
                latestAsset.mart = mart
                await mart.save()
                await latestAsset.save()
            } else {
                await createDataMart()
            }
        } else if (body.saveMode === "overwrite") {
            if (assetResult.length > 0) {
                const latestAsset = assetResult[0]
                latestAsset.isNewVersion = false
                await latestAsset.save()
            }
            await createDataMart()
        } else {
            PhLogger.info("其他情况")
        }
        return {status: "ok"}
    }
}