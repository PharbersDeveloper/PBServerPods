package com.pharbers.convert

import java.io.File

import com.pharbers.factory.Convert
import com.pharbers.oss.Oss

case class ConvertXls2Xlsx() extends Convert {
	override def exec(parameter: Map[String, Any]): (Boolean, String) = {
		val mappings = PhReadMapping.exec()
		val fileConversionXLSToXLXS = new FileConversionXLSToXLXS()
		val outputPath = s"${mappings.getProperty("output")}/${System.currentTimeMillis()}"
		try {
			isExistFile(s"${parameter("inputPath")}")
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
//
object main extends App {
//	val ossPath = "290999ab-ecc4-45c4-84ad-0215730bd8ad/1574303412548"
//	val downloadPath = PhReadMapping.exec().getProperty("input") + "/" + ossPath.substring(ossPath.lastIndexOf("/") + 1)
	val downloadPath = "/Users/qianpeng/Desktop/convert_excel/input/1574303592028.xls"
//	Oss().down(downloadPath, ossPath)
	val result = ConvertXls2Xlsx().exec(Map("inputPath" -> downloadPath))
	println(result)
}
