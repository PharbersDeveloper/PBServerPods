package com.pharbers.ConvertSchema2Mongo.PhHandler

import com.alibaba.fastjson.{JSON, JSONArray, JSONObject}
import com.pharbers.ConvertSchema2Mongo.PhCommon.PhMongo.MongoTrait
import com.mongodb.casbah.Imports._
import com.mongodb.casbah.query.Imports
import com.mongodb.casbah.query.dsl.QueryExpressionObject
import com.pharbers.ConvertSchema2Mongo.PhCommon.PhHDFS.HDFSUtil

import scala.collection.JavaConversions._

class PhConvert2MongoHandler() extends MongoTrait {
	// filedetails先Query出来符合要求的jobid
	val fileDetailsJobIDRes: List[String] = queryAll("filedetails", DBObject()) match {
		case Nil => println("is Nil"); Nil
		case res =>
			res.filter(dbo =>
				dbo.getAsOrElse[String]("tag", "") == "" &&
				!dbo.containsField("tag"))
			.flatMap(r => r.get("jobIds").asInstanceOf[BasicDBList].toList.map(_.toString))
	}
	println(fileDetailsJobIDRes.size)
	
//	val fileDetailsJobIDRes: List[String] = "47ef7-3ec3-402f-a4fa-c9fd90" :: Nil
	
	def findByKey(listStr: List[String], key: String): String = {
		listStr.find(x => x.contains(key)) match {
			case None => ""
			case Some(x) => x
		}
	}
	
	fileDetailsJobIDRes.foreach { jobId =>
		// 读取HDFS中meta data文件
		// 黄鑫
//		val path = s"/test/alex/0829b025-48ac-450c-843c-6d4ee91765ca/metadata/$jobId"
//		val path = s"/test/alex/a50dfe1a-572e-4dce-9b45-ef877e92b380/metadata/$jobId"
//		val path = s"/test/alex/07b8411a-5064-4271-bfd3-73079f2b42b2/metadata/$jobId"
		// MAX 5年
//		val path = s"/test/alex/ff89f6cf-7f52-4ae1-a5ec-2609169b3995/metadata/$jobId"
		// 更新CPA_MKT.csv
		val path = s"/test/alex/ff303a9d-50aa-491c-8dac-88571f6cf9f4/metadata/$jobId"
		val content: List[String] = new String(HDFSUtil.readHDFSFile(path)).split("\n")
				.map(x => x.replaceAll("\r", "")).toList
		if (content.head != "") {
			val schemaMap: Map[String, String] = Map(
				"schema" -> {
					val c = findByKey(content, "[{")
					c.substring(c.indexOf("[{"), c.lastIndexOf("}]") + 2)
				},
				"length" -> findByKey(content, """{"length":""")
					.split(":").last
					.replaceAll("_", "")
					.replaceAll("}", "")
			)
			val parseSchema2Json: List[JSONObject] = JSON.parseArray(schemaMap("schema"))
				.iterator().toList.map(_.asInstanceOf[JSONObject])

			// 读取datasets对应的jobid进行colNames的修改
			val condition: DBObject = DBObject("jobId" -> jobId)
			queryObject(condition, "datasets") match {
				case None => println("is None")
				case Some(dbo) =>
					dbo += "colNames" -> parseSchema2Json.map(_.get("key").toString)
					dbo += "url" -> path
					dbo += "length" -> schemaMap("length").toLong.asInstanceOf[Number]
					updateObject(condition, "datasets", dbo)
			}
		} else {
			println(s"jobID =====> $jobId")
		}

	}
}