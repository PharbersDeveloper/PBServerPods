"use strict"
import {JsonObject, JsonProperty} from "json2typescript"
import { KfkConf } from "./kfkConf"
import { ModelConf } from "./modelConf"
import { MongoConf } from "./mongoConf"
import { OAuthConf } from "./oauthConf"
import { OssConf } from "./ossConf"

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
}
