package com.pharbers.ConvertSchema2Mongo.PhHandler

import com.pharbers.ConvertSchema2Mongo.PhCommon.PhMongo.MongoTrait
import com.mongodb.casbah.Imports._
import com.pharbers.ConvertSchema2Mongo.PhCommon.PhHDFS.HDFSUtil

class PhConvert2MetaDataHandler extends MongoTrait {
    val jobIdsWithUrl: List[(ObjectId, String)] = queryAll("datasets", "url" $ne "")
	    .map(x => (x.getAsOrElse[ObjectId]("_id", null), x.getAsOrElse[String]("url", "")))
	
	val jobIds: List[ObjectId] = jobIdsWithUrl.map(_._1)
	
	val fileNameWithJobId: List[(ObjectId, String)] = queryAll("assets", "dfs" $in jobIds) match {
		case Nil => println("IS Nil"); Nil
		case res =>
			res.flatMap (x =>
				x.get("dfs").asInstanceOf[BasicDBList].toList.map(y =>
					(y.asInstanceOf[ObjectId], x.getAsOrElse[String]("name", "")))
			)
	}
	
	
	// 路径, 文件名
	val realData: List[(String, String)] = jobIdsWithUrl.map{x =>
		val d = fileNameWithJobId.find(y => y._1 == x._1).get
		(x._2, d._2)
	}

// Astellas_1401_1412_CPA.xlsx, c02ec-1547-4f94-b1b7-448310, /test/alex/ff89f6cf-7f52-4ae1-a5ec-2609169b3995/metadata/c02ec-1547-4f94-b1b7-448310
	var num = 0
	realData.foreach {x =>
		val realContent = new String(HDFSUtil.readHDFSFile(x._1))
		if (!realContent.contains("fileName") && !realContent.isEmpty) {
			println(x._1)
			println(x._2)
			num += 1
			println(s"==============>>>>>$num")
//			val flag: Boolean = HDFSUtil.append(x._1, s"""{"fileName":"${x._2}"}""")
//			println(flag)
//			Thread.sleep(1000)
		}
	}
	
//	val flag = HDFSUtil.append("/test/alex/ff89f6cf-7f52-4ae1-a5ec-2609169b3995/metadata/c02ec-1547-4f94-b1b7-448310", """{"fileName":"Fuck"}""")
//	println(flag)
	

	
}
