import { JsonConvert, OperationMode, ValueCheckingMode } from "json2typescript"
import { slow, suite, test, timeout } from "mocha-typescript"
import mongoose = require("mongoose")
import PhLogger from "../../src/logger/phLogger"

@suite(timeout(1000 * 60), slow(2000))
class UploadFileTransmit {

    public static before() {
        PhLogger.info(`before starting the test`)
        mongoose.connect("mongodb://localhost:27017/pharbers-ntm-client")
    }

    public static after() {
        PhLogger.info(`after starting the test`)
        mongoose.disconnect()
    }

    @test public async excelModelData() {
        PhLogger.info(`start input data with excel`)
    }
}
