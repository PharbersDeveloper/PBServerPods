package com.pharbers.convert

import java.util.Properties

import scala.io.Source._


object PhReadMapping {
	def exec(): Properties = {
		val props = new Properties()
		val source = fromFile(sys.env("PH_TS_SANDBOX_HOME") + "/conf/ooxml-strict-mappings.properties", "ISO-8859-1")
		val lines = source.getLines().toList
		lines.foreach { line =>
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
