"use strict"
import {JsonObject, JsonProperty} from "json2typescript"
import { KfkConf } from "./kfkConf"
import { ModelConf } from "./modelConf"
import { MongoConf } from "./mongoConf"
import { OAuthConf } from "./oauthConf"
import { OssConf } from "./ossConf"
import { ModuleConf } from "./moduleConf"

@JsonObject("ServerConf")
export class ServerConf {

    @JsonProperty("models", [ModelConf])
    public models: ModelConf[] = undefined

    @JsonProperty("mongo", MongoConf)
    public mongo: MongoConf = undefined

    @JsonProperty("oauth", OAuthConf)
    public oauth: OAuthConf = undefined

    @JsonProperty("oss", OssConf)
    public oss: OssConf = undefined

    @JsonProperty("kfk", KfkConf)
    public kfk: KfkConf = undefined

    @JsonProperty("modules", [ModuleConf])
    public modules: ModuleConf[] = undefined

    // @JsonProperty("handlers", [String])
    // public handlers: string[] = undefined
}
