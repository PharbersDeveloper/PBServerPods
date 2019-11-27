"use strict"
import Asset from "../models/Asset"
import File from "../models/File"
import axios from "axios"

export class ReCommitJobHandler {
    async reCommitJobWithTraceId(body: any) {
        const asset = await new Asset().getModel().findOne({traceId: body.traceId, isNewVersion: true})
        const file = await new File().getModel().findById(asset.file)
        const reqbody = {
            "traceId": body.traceId,
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
        await axios.post(`http://localhost:8080/putJob2Stream`, reqbody)
        return {status: "ok"}
    }
}