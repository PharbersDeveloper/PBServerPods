"use strict"
import {JsonObject, JsonProperty} from "json2typescript"

@JsonObject("OAuthConf")
export class OAuthConf {

    @JsonProperty("debugging", Boolean)
    public debugging: boolean = false

    @JsonProperty("oauthHost", String)
    public oauthHost: string = undefined

    @JsonProperty("oauthPort", String)
    public oauthPort: string = undefined

    @JsonProperty("oauthApiNamespace", String)
    public oauthApiNamespace: string = undefined
}
