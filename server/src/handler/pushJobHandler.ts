"use strict"

import mongoose = require("mongoose")
import Asset from "../models/Asset"
import File from "../models/File"
import axios from "axios"
import PhLogger from "../logger/phLogger"

export default class PushJobHandler {
    async pushJob(body: any) {
        const asset = await new Asset().getModel().findById(new mongoose.mongo.ObjectId(body.assetId))
        const file = await new File().getModel().findById(asset.file)
        if (file === null) {
           return {status: "no"}
        }

        let labels = []
        let providers = []
        try {
            labels = asset.labels.map(x => JSON.parse(x))
        } catch (e) {
            labels = asset.labels.map(x => x)
        }
        const pv = labels.find(x => x.providers)
        if (pv !== undefined) {
            providers = pv.providers.map((item: any) => Object.keys(item).map(key => item[key])).flat()
        } else {
            providers = asset.providers
        }

        const condition: any = {
            assetId:    body.assetId,
            jobId:      body.jobId,
            traceId:    body.traceId,
            owner:      asset.owner,
            createTime: asset.createTime,
            ossKey:     file.url,
            fileType:   file.extension,
            fileName:   file.fileName,
            providers,
            molecules:  asset.molecules,
            markets:    asset.markets,
            geoCover:   asset.geoCover,
            dataCover:  asset.dataCover,
            labels:     asset.labels
        }

        axios.post(`http://127.0.0.1:30001/putJob2Stream`, condition)

        return {status: "ok"}
    }
}