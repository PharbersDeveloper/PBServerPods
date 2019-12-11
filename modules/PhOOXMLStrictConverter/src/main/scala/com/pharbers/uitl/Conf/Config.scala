package com.pharbers.uitl.Conf

import com.alibaba.fastjson.JSON

import scala.io.Source.fromFile
import scala.collection.JavaConversions._

object Config extends App{
	def loadOssConfig(): Map[String, Map[String, String]] = {
		val source = fromFile(sys.env("PHPRODSHOME") + "/conf/oss_config.json")
		val jsonStr = source.getLines().mkString("").replace(" ", "")
		val jsonObject = JSON.parseObject(jsonStr, classOf[java.util.Map[String, java.util.Map[String, String]]])
		jsonObject.flatMap (m => Map(m._1 -> m._2.toMap)).toMap
	}
}
