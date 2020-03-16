"use strict"
import axios from "axios"
import bodyParser from "body-parser"
import express from "express"
import API, { ResourceTypeRegistry } from "json-api"
import { APIControllerOpts } from "json-api/build/src/controllers/API"
import mongoose = require("mongoose")
import KafkaDelegate from "../kafka/KafkaDelegate"
import PhLogger from "../logger/phLogger"
import { urlEncodeFilterParser } from "./urlEncodeFilterParser"
import phLogger from "../logger/phLogger"
import { CONFIG } from "../shared/config"
import {FindFilePathHandler} from "../handler/findFilePathHandler"
import {UpdateFilePathHandler} from "../handler/updateFilePathHandler"
import {ReCommitJobHandler} from "../handler/reCommitJobHandler"
import {UpdateJobId2MongoHandler} from "../handler/updateJobId2MongoHandler"
import {JobBloodHandler} from "../handler/jobBloodHandler"
import { AssetDataMartHandler } from "../handler/AssetDataMartHandler"

/**
 * The summary section should be brief. On a documentation web site,
 * it will be shown on a page that lists summaries for many different
 * API items.  On a detail page for a single item, the summary will be
 * shown followed by the remarks section (if any).
 *
 */
export default class AppDelegate {

    /**
     * @returns the configuration of the server
     */

    private app = express()
    private router = express.Router()
    private kafka: KafkaDelegate = null // new KafkaDelegate()

    public exec() {
        this.loadConfiguration()
        this.configMiddleware()
        this.connect2MongoDB()
        this.generateRoutes(this.getModelRegistry())
        this.generateModules()
        this.listen2Port(8080)
    }

    protected configMiddleware() {

        this.app.use(bodyParser.json())
        this.app.use( bodyParser.urlencoded( {
            extended: true
        } ) )

        if (!CONFIG.oauth.debugging) {
            // a middleware function with no mount path. This code is executed for every request to the router
            this.router.use((req, res, next) => {
                const auth = req.get("Authorization")
                if (auth === undefined) {
                    PhLogger.error("no auth")
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
                        PhLogger.error("auth error")
                        res.status(500).send(response.data)
                        return
                    } else {
                        next()
                    }
                }).catch((error) => {
                    PhLogger.error("auth error")
                    res.status(500).send(error)
                    return
                })
            })
        }

        this.app.use("/", this.router)
    }

    protected loadConfiguration() {
        try {
            this.kafka = new KafkaDelegate(CONFIG.kfk)
        } catch (e) {
            PhLogger.fatal( e as Error )
        }
    }

    protected generateModels(): any {
        const prefix = "/server/dist/src/models/"
        const path = process.env.PHPRODSHOME + prefix
        const suffix = ".js"
        const result: {[index: string]: any} = {}
        CONFIG.models.forEach((ele) => {
                const filename = path + ele.file + suffix
                const one = require(filename).default
                result[ele.file] = new one().getModel()
            })
        return result
    }

    protected connect2MongoDB() {
        const prefix = CONFIG.mongo.algorithm
        const host = CONFIG.mongo.host
        const port = `${CONFIG.mongo.port}`
        const username = CONFIG.mongo.username
        const pwd = CONFIG.mongo.pwd
        const coll = CONFIG.mongo.coll
        const auth = CONFIG.mongo.auth
        if (auth) {
            PhLogger.info(`connect mongodb with ${ username } and ${ pwd }`)
            mongoose.connect(prefix + "://" + username + ":" + pwd + "@" + host + ":" + port + "/" + coll,
                (err: any) => {
                    if (err != null) {
                        PhLogger.error(err)
                    }
                })
        } else {
            PhLogger.info(`connect mongodb without auth`)
            mongoose.connect(prefix + "://" + host + ":" + port + "/" + coll,
                (err: any) => {
                if (err != null) {
                    PhLogger.error(err)
                }
            })
        }
    }

    protected getModelRegistry(): ResourceTypeRegistry {
        const result: {[index: string]: any} = {}
        CONFIG.models.forEach((ele) => {
            result[ele.reg] = {}
        })
        return new API.ResourceTypeRegistry(result, {
            dbAdapter: new API.dbAdapters.Mongoose(this.generateModels()),
            info: {
                description: "Blackmirror inc. Alfred Yang 2019"
            },
            urlTemplates: {
                self: "/{type}/{id}"
            },
        })
    }

    protected generateRoutes(registry: ResourceTypeRegistry) {

        const opts: APIControllerOpts = {
            filterParser: urlEncodeFilterParser
        }

        const Front = new API.httpStrategies.Express(
            new API.controllers.API(registry, opts),
            new API.controllers.Documentation(registry, {name: "Pharbers API"})
        )

        PhLogger.startConnectLog(this.app)
        this.app.get("/", Front.docsRequest)
        const perfix = "/:type"
        const ms = CONFIG.models.map((x) => x.reg).join("|")
        const suffix = "/:id"

        const all = perfix + "(" + ms + ")"
        const one = all + suffix
        const relation = one + "/relationships/:relationship"

        // Add routes for basic list, read, create, update, delete operations
        this.app.get(all, Front.apiRequest)
        this.app.get(one, Front.apiRequest)
        this.app.post(all, Front.apiRequest)
        this.app.patch(one, Front.apiRequest)
        this.app.delete(one, Front.apiRequest)

        // Add routes for adding to, removing from, or updating resource relationships
        this.app.post(relation, Front.apiRequest)
        this.app.patch(relation, Front.apiRequest)
        this.app.delete(relation, Front.apiRequest)
    }

    protected generateModules() {
        phLogger.info(CONFIG.modules)

        // TODO：先实现功能，在结构
        this.router.post("/uploadFileEnd" , async (req, res) => {
            res.json(await new UpdateJobId2MongoHandler().uploadFileEnd(req.body))
        } )

        this.router.post("/findFilePathWithId" , async (req, res) => {
            res.json(await new FindFilePathHandler().findFilePathWithId(req.body))
        } )

        this.router.post("/updateAssetVersion" , async (req, res) => {
            res.json(await new UpdateFilePathHandler().updateAssetVersion(req.body))
        } )

        this.router.post("/reCommitJobWithAssetId" , async (req, res) => {
            res.json(await new ReCommitJobHandler().reCommitJobWithAssetId(req.body))
        } )

        this.router.post("/createDataSetsAndJob" , async (req, res) => {
            res.json(await new JobBloodHandler().createDataSetsAndJob(req.body))
        } )

        this.router.post("/assetDataMart" , async (req, res) => {
            res.json(await new AssetDataMartHandler().assetDataMart(req.body))
        } )

        CONFIG.modules.forEach( (module) => {
            const host = module.host
            const port = module.port
            // TODO: 改成策略模式会好看很多
            if (module.protocol === "http") {
                module.routers.forEach( (router) => {
                    this.router.post("/" + router, async (req, res) => {
                        const result = await axios.post(`http://${host}:${port}/${router}`, req.body)
                        res.json(result.data)
                    } )
                } )
            } else {
                phLogger.fatal("not implemented!!")
            }
        } )
    }

    protected listen2Port(port: number) {
        // start the Express server
        this.app.listen( port, () => {
            PhLogger.info( `server started at http://localhost:${ port }` )
        } )
    }
}
