package com.pharbers.convert

import java.io.File

import com.pharbers.factory.Convert

case class ConvertXls2Xlsx() extends Convert {
	override def exec(parameter: Map[String, Any]): (Boolean, String) = {
		val fileConversionXLSToXLXS = new FileConversionXLSToXLXS()
		val outputPath = s"${sys.env("CONVERTOUTPUT")}/${System.currentTimeMillis()}"
		try {
			isExistFile(parameter("inputPath").toString)
			fileConversionXLSToXLXS.convertXLS2XLSX(s"${parameter("inputPath")}", outputPath)
			(true, outputPath)
		} catch {
			case x: Exception =>
				println(x.getMessage)
				x.printStackTrace()
				(false, "")
		}
	}
	
	def isExistFile(path: String): Unit = {
		val file = new File(path)
		if (!file.exists()) {
			println("File is None")
			Thread.sleep(1000)
			isExistFile(path)
		}
	}
}
