"use strict";
var __decorate = (this && this.__decorate) || function (decorators, target, key, desc) {
    var c = arguments.length, r = c < 3 ? target : desc === null ? desc = Object.getOwnPropertyDescriptor(target, key) : desc, d;
    if (typeof Reflect === "object" && typeof Reflect.decorate === "function") r = Reflect.decorate(decorators, target, key, desc);
    else for (var i = decorators.length - 1; i >= 0; i--) if (d = decorators[i]) r = (c < 3 ? d(r) : c > 3 ? d(target, key, r) : d(target, key)) || r;
    return c > 3 && r && Object.defineProperty(target, key, r), r;
};
var __metadata = (this && this.__metadata) || function (k, v) {
    if (typeof Reflect === "object" && typeof Reflect.metadata === "function") return Reflect.metadata(k, v);
};
var __awaiter = (this && this.__awaiter) || function (thisArg, _arguments, P, generator) {
    return new (P || (P = Promise))(function (resolve, reject) {
        function fulfilled(value) { try { step(generator.next(value)); } catch (e) { reject(e); } }
        function rejected(value) { try { step(generator["throw"](value)); } catch (e) { reject(e); } }
        function step(result) { result.done ? resolve(result.value) : new P(function (resolve) { resolve(result.value); }).then(fulfilled, rejected); }
        step((generator = generator.apply(thisArg, _arguments || [])).next());
    });
};
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
// import { JsonConvert, OperationMode, ValueCheckingMode } from "json2typescript"
const mocha_typescript_1 = require("mocha-typescript");
const phLogger_1 = __importDefault(require("../../src/logger/phLogger"));
const mongoose = require("mongoose");
const phLogger_2 = __importDefault(require("../../src/logger/phLogger"));
const FileDetail_1 = __importDefault(require("../../src/models/FileDetail"));
const FileVersion_1 = __importDefault(require("../../src/models/FileVersion"));
const SandboxIndex_1 = __importDefault(require("../../src/models/SandboxIndex"));
const Asset_1 = __importDefault(require("../../src/models/Asset"));
const File_1 = __importDefault(require("../../src/models/File"));
const DataSet_1 = __importDefault(require("../../src/models/DataSet"));
let UploadFileTransmit = class UploadFileTransmit {
    static before() {
        phLogger_2.default.info(`before starting the test`);
        mongoose.connect("mongodb://pharbers.com:5555/pharbers-sandbox-3");
    }
    static after() {
        phLogger_2.default.info(`after starting the test`);
        mongoose.disconnect();
    }
    fileTransmit() {
        return __awaiter(this, void 0, void 0, function* () {
            phLogger_2.default.info(`start trans data from old version`);
            yield this.fileTransmitImpl();
        });
    }
    fileTransmitImpl() {
        return __awaiter(this, void 0, void 0, function* () {
            const sim = new SandboxIndex_1.default().getModel();
            const fdm = new FileDetail_1.default().getModel();
            const fvm = new FileVersion_1.default().getModel();
            const fm = new File_1.default().getModel();
            const dsm = new DataSet_1.default().getModel();
            const am = new Asset_1.default().getModel();
            const contents = yield sim.find({});
            yield Promise.all(contents.map((content) => __awaiter(this, void 0, void 0, function* () {
                const owner = content.account;
                const filesIds = content.files;
                yield Promise.all(filesIds.map((id) => __awaiter(this, void 0, void 0, function* () {
                    const fd = yield fdm.findOne({
                        _id: id
                    });
                    const fvId = fd.versions[0];
                    const fv = yield fvm.findOne({
                        _id: fvId
                    });
                    /**
                     * 1. 将FileDetail, FileVersion转成File
                     */
                    const f = new File_1.default();
                    f.url = fv.where;
                    f.size = fv.size;
                    f.fileName = fd.name;
                    f.extension = fd.extension;
                    f.uploaded = fd.created;
                    const fc = yield fm.create(f);
                    /**
                     * 2. 将JobID 创建出来的DataSet MetaData化
                     */
                    const jIds = fd.jobIds;
                    const dfs = yield Promise.all(jIds.map((jid) => __awaiter(this, void 0, void 0, function* () {
                        const ds = new DataSet_1.default();
                        ds.jobId = jid;
                        return yield dsm.create(ds);
                    })));
                    /**
                     * 3. 将用户上传的内容，抽象成平台所需要的Assents
                     */
                    const asset = new Asset_1.default();
                    asset.name = fd.name;
                    asset.description = fd.name;
                    asset.traceId = fd.traceID;
                    asset.dataType = "file";
                    asset.file = fc;
                    asset.dfs = dfs;
                    asset.owner = fd.ownerID;
                    asset.accessibility = "w";
                    asset.version = 0;
                    /**
                     * 4. 为数据添加tags
                     */
                    if (fd.name.indexOf("_") > 0) {
                        phLogger_1.default.info("cpa gyc data");
                        this.cpa_gyc_name_2_tags(fd.name, asset);
                    }
                    else {
                        phLogger_1.default.info("chc data");
                        this.chc_name_2_tags(fd.name, asset);
                    }
                    yield am.create(asset);
                })));
            })));
            // phLogger.info(await tmp[0])
        });
    }
    cpa_gyc_name_2_tags(name, asset) {
        const tags = name.split("_");
        if (tags.length < 4) {
            if (name.indexOf("Lilly") !== -1) {
                asset.providers = [tags[0]];
                asset.dataCover = [tags[1], tags[2]];
            }
            else if (name.indexOf("cpa") !== -1 || name.indexOf("gyc") !== -1) {
                asset.providers = [tags[0]];
            }
            else {
                phLogger_1.default.info(name);
                this.chc_name_2_tags(name, asset);
            }
        }
        else {
            asset.providers = [tags[0], tags[3]];
            asset.dataCover = [tags[1], tags[2]];
        }
    }
    chc_name_2_tags(name, asset) {
        phLogger_1.default.info(name);
        const geo = [];
        if (name.indexOf("北京") !== -1) {
            geo.push("北京");
        }
        if (name.indexOf("上海") !== -1) {
            geo.push("上海");
        }
        if (name.indexOf("安徽") !== -1) {
            geo.push("安徽");
        }
        if (name.indexOf("山东") !== -1) {
            geo.push("山东");
        }
        if (name.indexOf("广州") !== -1) {
            geo.push("广州");
        }
        if (name.indexOf("福建") !== -1) {
            geo.push("福建");
        }
        if (name.indexOf("江苏") !== -1) {
            geo.push("江苏");
        }
        asset.geoCover = geo;
        if (name.indexOf("【") !== -1) {
            const start = name.indexOf("【");
            const end = name.indexOf("】");
            const length = end - start - 1;
            const provider = name.substr(start + 1, length);
            asset.providers = [provider];
        }
    }
};
__decorate([
    mocha_typescript_1.test,
    __metadata("design:type", Function),
    __metadata("design:paramtypes", []),
    __metadata("design:returntype", Promise)
], UploadFileTransmit.prototype, "fileTransmit", null);
UploadFileTransmit = __decorate([
    mocha_typescript_1.suite(mocha_typescript_1.timeout(1000 * 60), mocha_typescript_1.slow(2000))
], UploadFileTransmit);
//# sourceMappingURL=UploadFileTransmit.js.map