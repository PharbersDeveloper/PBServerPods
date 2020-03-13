package com.pharbers.uitl.Conf

import java.util.Properties

import com.alibaba.fastjson.JSON

import scala.io.Source.fromFile
import scala.collection.JavaConversions._

object Config {
	def loadOssConfig(): Map[String, Map[String, String]] = {
		val source = fromFile(sys.env("PHOSSCONF"))
		val jsonStr = source.getLines().mkString("").replace(" ", "")
		val jsonObject = JSON.parseObject(jsonStr, classOf[java.util.Map[String, java.util.Map[String, String]]])
		jsonObject.flatMap (m => Map(m._1 -> m._2.toMap)).toMap
	}
	
	def loadExcelMappingConfig(): Properties = {
		val props = new Properties()
		val source = fromFile(sys.env("PHOOXMLCONF"), "ISO-8859-1")
		val jsonStr = source.getLines().mkString("").replace(" ", "")
		val jsonObject = JSON.parseObject(jsonStr)
		jsonObject.getJSONArray("replace").toList.asInstanceOf[List[String]].foreach { line =>
			val spKV = line.split("=")
			if (spKV.length >= 2) {
				props.setProperty(spKV(0), spKV(1))
			} else if(spKV.length == 1) {
				props.setProperty(spKV(0), "")
			}
		}
		source.close()
		props
	}
}
