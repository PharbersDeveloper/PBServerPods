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
let OAuthConf = class OAuthConf {
    constructor() {
        this.debugging = false;
        this.oauthHost = undefined;
        this.oauthPort = undefined;
        this.oauthApiNamespace = undefined;
    }
};
__decorate([
    json2typescript_1.JsonProperty("debugging", Boolean),
    __metadata("design:type", Boolean)
], OAuthConf.prototype, "debugging", void 0);
__decorate([
    json2typescript_1.JsonProperty("oauthHost", String),
    __metadata("design:type", String)
], OAuthConf.prototype, "oauthHost", void 0);
__decorate([
    json2typescript_1.JsonProperty("oauthPort", String),
    __metadata("design:type", String)
], OAuthConf.prototype, "oauthPort", void 0);
__decorate([
    json2typescript_1.JsonProperty("oauthApiNamespace", String),
    __metadata("design:type", String)
], OAuthConf.prototype, "oauthApiNamespace", void 0);
OAuthConf = __decorate([
    json2typescript_1.JsonObject("OAuthConf")
], OAuthConf);
exports.OAuthConf = OAuthConf;
//# sourceMappingURL=oauthConf.js.map