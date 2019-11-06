"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
const log4js_1 = require("log4js");
class PhLogger {
    constructor() {
        log4js_1.configure(process.env.PH_TS_SANDBOX_HOME + "/log4js.json");
    }
    startConnectLog(app) {
        // tslint:disable-next-line: max-line-length
        app.use(log4js_1.connectLogger(log4js_1.getLogger("http"), { level: "auto", format: (req, res, format) => format(`:remote-addr - :method :url HTTP/:http-version :status :referrer`) }));
    }
    trace(msg, ...params) {
        log4js_1.getLogger().trace(msg, params);
    }
    debug(msg, ...params) {
        log4js_1.getLogger().debug(msg, params);
    }
    info(msg, ...params) {
        log4js_1.getLogger().info(msg, params);
    }
    warn(msg, ...params) {
        log4js_1.getLogger().warn(msg, params);
    }
    error(msg, ...params) {
        log4js_1.getLogger().error(msg, params);
    }
    fatal(msg, ...params) {
        log4js_1.getLogger().fatal(msg, params);
    }
}
exports.default = new PhLogger();
//# sourceMappingURL=phLogger.js.map