// import { JsonConvert, OperationMode, ValueCheckingMode } from "json2typescript"
import { slow, suite, test, timeout } from "mocha-typescript"
import phLogger from "../../src/logger/phLogger"
import mongoose = require("mongoose")
import PhLogger from "../../src/logger/phLogger"
import { ObjectId } from "bson"

import FileDetail from "../../src/models/FileDetail"
import FileVersion from "../../src/models/FileVersion"
import SandboxIndex from "../../src/models/SandboxIndex"

import Assent from "../../src/models/Assent"
import File from "../../src/models/File"
import DataSet from "../../src/models/DataSet"

@suite(timeout(1000 * 60), slow(2000))
class UploadFileTransmit {

    public static before() {
        PhLogger.info(`before starting the test`)
        mongoose.connect("mongodb://pharbers.com:5555/pharbers-sandbox-3")
    }

    public static after() {
        PhLogger.info(`after starting the test`)
        mongoose.disconnect()
    }

    @test public async fileTransmit() {
        PhLogger.info(`start trans data from old version`)
        await this.fileTransmitImpl()
    }

    public async fileTransmitImpl() {
        const sim = new SandboxIndex().getModel()
        const fdm = new FileDetail().getModel()
        const fvm = new FileVersion().getModel()
        const fm = new File().getModel()

        const contents = await sim.find({})
        await Promise.all(contents.map( async (content) => {
            const owner = content.account
            const filesIds = content.files
            await Promise.all(filesIds.map( async (id) => {
                const fd = await fdm.findOne({
                    _id: id
                } )
                const fvid = fd.versions[0]
                const fv = await fvm.findOne({
                    _id: fvid
                } )

                /**
                 * 1. 将FileDetail转成File
                 */
                const f = new File()
                f.url = fv.where
                f.fileName = fd.name
                f.extension = fd.extension
                f.uploaded = fd.created
                await fm.create(f)
            } ))
        } ))
        // phLogger.info(await tmp[0])
    }
}
