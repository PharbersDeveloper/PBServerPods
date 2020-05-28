
import phLogger from "../logger/phLogger"
import {CONFIG} from "../shared/config"
import axios from "axios"

export default class TokenCheckMiddleWare {
    public exec(router: any) {
        router.use(this.check)
    }

    protected check(req: any, res: any, next: any) {
        // phLogger.info("Token 验证")

        if (!CONFIG.oauth.debugging) {
            // a middleware function with no mount path. This code is executed for every request to the router
            const auth = req.get("Authorization")
            if (auth === undefined) {
                phLogger.error("no auth")
                res.status(500).send({error: "no auth!"})
                return
            }

            const host = CONFIG.oauth.oauthHost
            const port = CONFIG.oauth.oauthPort
            const namespace = CONFIG.oauth.oauthApiNamespace

            axios.post(`http://${host}${port}/${namespace}/TokenValidation`, null, {
                headers: {
                    Authorization: auth,
                },
            }).then((response) => {
                if (response.data.error !== undefined) {
                    phLogger.error("auth error")
                    res.status(500).send(response.data)
                    return
                } else {
                    next()
                }
            }).catch((error) => {
                phLogger.error("auth error")
                res.status(500).send(error)
                return
            })
            next()
        } else {
            next()
        }
    }
}