package com.pharbers.convert

import java.util.Properties

import com.alibaba.fastjson.JSON

import scala.io.Source._

import scala.collection.JavaConversions._


object PhReadMapping {
	def exec(): Properties = {
		val props = new Properties()
		val source = fromFile(sys.env("PHPRODSHOME") + "/conf/ooxml_config.json", "ISO-8859-1")
		val jsonStr = source.getLines().mkString("").replace(" ", "")
		val jsonObject = JSON.parseObject(jsonStr)
		props.setProperty("output", jsonObject.getString("output"))
		props.setProperty("input", jsonObject.getString("input"))
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
