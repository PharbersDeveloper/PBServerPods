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
Object.defineProperty(exports, "__esModule", { value: true });
const json2typescript_1 = require("json2typescript");
const kfkConf_1 = require("./kfkConf");
const modelConf_1 = require("./modelConf");
const mongoConf_1 = require("./mongoConf");
const oauthConf_1 = require("./oauthConf");
const ossConf_1 = require("./ossConf");
const moduleConf_1 = require("./moduleConf");
let ServerConf = class ServerConf {
    constructor() {
        this.models = undefined;
        this.mongo = undefined;
        this.oauth = undefined;
        this.oss = undefined;
        this.kfk = undefined;
        this.modules = undefined;
    }
};
__decorate([
    json2typescript_1.JsonProperty("models", [modelConf_1.ModelConf]),
    __metadata("design:type", Array)
], ServerConf.prototype, "models", void 0);
__decorate([
    json2typescript_1.JsonProperty("mongo", mongoConf_1.MongoConf),
    __metadata("design:type", mongoConf_1.MongoConf)
], ServerConf.prototype, "mongo", void 0);
__decorate([
    json2typescript_1.JsonProperty("oauth", oauthConf_1.OAuthConf),
    __metadata("design:type", oauthConf_1.OAuthConf)
], ServerConf.prototype, "oauth", void 0);
__decorate([
    json2typescript_1.JsonProperty("oss", ossConf_1.OssConf),
    __metadata("design:type", ossConf_1.OssConf)
], ServerConf.prototype, "oss", void 0);
__decorate([
    json2typescript_1.JsonProperty("kfk", kfkConf_1.KfkConf),
    __metadata("design:type", kfkConf_1.KfkConf)
], ServerConf.prototype, "kfk", void 0);
__decorate([
    json2typescript_1.JsonProperty("modules", [moduleConf_1.ModuleConf]),
    __metadata("design:type", Array)
], ServerConf.prototype, "modules", void 0);
ServerConf = __decorate([
    json2typescript_1.JsonObject("ServerConf")
], ServerConf);
exports.ServerConf = ServerConf;
//# sourceMappingURL=serverConf.js.map