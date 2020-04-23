import { slow, suite, test, timeout } from "mocha-typescript"
import mongoose = require("mongoose")
import PhLogger from "../../src/logger/phLogger"

import Asset from "../../src/models/Asset"
import File from "../../src/models/File"

@suite(timeout(1000 * 60), slow(2000))
class ReplaceTime {

    public static before() {
        PhLogger.info(`before starting the test`)
        mongoose.connect("mongodb://127.0.0.1:27017/pharbers-sandbox-merge")
    }

    public static after() {
        PhLogger.info(`after starting the test`)
        mongoose.disconnect()
    }

    @test public async assetTransmit() {
        PhLogger.info(`start trans data from old version`)
        await this.assetTransmitImpl()
    }

    public async assetTransmitImpl() {
        const assetModel = new Asset().getModel()

        const assetResult = await assetModel.find()


        await Promise.all(assetResult.map(async data => {
            const fileData = await new File().getModel().findById(data.file)
            // data.createTime = fileData.uploaded
            await data.save()
        }))
    }
}
