"use strict"
import Asset from "../models/Asset"
import File from "../models/File"
import axios from "axios"
import uuid from "uuid/v1"
import mongoose = require("mongoose")
import PhLogger from "../logger/phLogger"

export default class ReCommitJobHandler {
    async reCommitJobWithAssetId(body: any) {
        PhLogger.info("进入重新提交")
        const asset = await new Asset().getModel().findOne({_id: new mongoose.mongo.ObjectId(body.assetId), isNewVersion: true})
        const file = await new File().getModel().findById(asset.file)

        const reqBody = {
            "assetId": asset._id.toString(),
            "jobId": uuid(),
            "traceId": uuid(),
            "ossKey": file.url,
            "fileType": file.extension,
            "fileName": file.fileName,
            "labels": asset.labels,
            "dataCover": asset.dataCover,
            "geoCover": asset.geoCover,
            "markets": asset.markets,
            "molecules": asset.molecules,
            "providers": asset.providers
        }
        await axios.post(`http://localhost:8080/putJob2Stream`, reqBody)
        return {status: "ok"}
    }
}