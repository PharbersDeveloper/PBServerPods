package com.pharbers.convert

import com.pharbers.factory.Convert
import com.pharbers.uitl.Conf.Config
import javax.xml.namespace.QName
import javax.xml.stream.{XMLEventFactory, XMLInputFactory, XMLOutputFactory}

case class ConvertStrictExcel() extends Convert {
		override def exec(parameter: Map[String, Any]): (Boolean, String) = {
			val mappings = Config.loadExcelMappingConfig()
			val ooXml = new OoXmlStrictConverter(XMLEventFactory.newInstance,
				XMLInputFactory.newInstance,
				XMLOutputFactory.newInstance,
				new QName("conformance"))
			val outputPath = s"${sys.env("CONVERTOUTPUT")}/${System.currentTimeMillis()}"
			try {
				ooXml.transform(parameter("inputPath").toString, outputPath, mappings)
				(true, outputPath)
			} catch {
				case _: Exception => (false, "")
			}
		}
}
