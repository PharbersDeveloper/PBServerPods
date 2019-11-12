package com.pharbers.ConvertSchema2Mongo.PhCommon.PhMongo

import scala.xml.{Elem, XML}
import com.mongodb.casbah.Imports._

sealed trait MongoInstance {
//	lazy val mongoXml: Elem = XML.load("src/main/resources/sandbox_mongo_connect.xml")
//	lazy val url: String = (mongoXml \ "server_host" \ "@value").toString()
//	lazy val port: Int = (mongoXml \ "server_port" \ "@value").toString().toInt
//	lazy val databaseName: String = (mongoXml \ "conn_name" \ "@value").toString()
//	lazy val dbIns = MongoClient(url, port)
	
	lazy val url: String = "127.0.0.1"
	lazy val port: Int = 27017
	lazy val databaseName: String = "pharbers-sandbox-4"
}

trait MongoTrait extends MongoInstance {
	import collection.JavaConverters._
	def queryObject(condition: DBObject, coll: String): Option[DBObject] = {
		val dbIns = MongoClient(url, port)
		val collect = dbIns.getDB(databaseName).getCollection(coll)
		val result = collect.find(condition).toArray.asScala.toList match {
			case Nil => None
			case res => Some(res.head)
		}
		dbIns.close()
		result
	}

	def updateObject(condition: DBObject, coll: String, obj: DBObject): Int = {
		val dbIns = MongoClient(url, port)
		val collect = dbIns.getDB(databaseName).getCollection(coll)
		val result = collect.update(condition, obj).getN
		dbIns.close()
		result
	}

	def queryAll(coll: String, condition: DBObject): List[DBObject] = {
		import collection.JavaConverters._
		val dbIns = MongoClient(url, port)
		val collect = dbIns.getDB(databaseName).getCollection(coll)
		val result = collect.find(condition).toArray().asScala.toList
		dbIns.close()
		result
	}
}

