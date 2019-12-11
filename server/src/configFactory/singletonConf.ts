import {ServerConf} from "./serverConf"
import {JsonConvert, ValueCheckingMode} from "json2typescript"
import * as yaml from "js-yaml"
import * as fs from "fs"
import PhLogger from "../logger/phLogger"

export class SingletonInitConf {
    private static conf: ServerConf

    getConf() {
        if (!SingletonInitConf.conf) SingletonInitConf.init()
        return SingletonInitConf.conf
    }

    private static init() {
        if (!SingletonInitConf.conf) {
            PhLogger.info("Init Config......")
            const path = process.env.PHPRODSHOME + "/conf"
            const jsonConvert: JsonConvert = new JsonConvert()
            const doc = yaml.safeLoad(fs.readFileSync(path + "/server.yml", "utf8"))
            jsonConvert.ignorePrimitiveChecks = false
            jsonConvert.valueCheckingMode = ValueCheckingMode.DISALLOW_NULL
            SingletonInitConf.conf = jsonConvert.deserializeObject(doc, ServerConf)
        }
    }
}