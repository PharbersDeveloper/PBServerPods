package com.pharbers.convert

import com.pharbers.factory.Convert
import javax.xml.namespace.QName
import javax.xml.stream.{XMLEventFactory, XMLInputFactory, XMLOutputFactory}

case class ConvertStrictExcel() extends Convert {
		override def exec(parameter: Map[String, Any]): (Boolean, String) = {
			val mappings = PhReadMapping.exec()
			val ooXml = new OoXmlStrictConverter(XMLEventFactory.newInstance,
				XMLInputFactory.newInstance,
				XMLOutputFactory.newInstance,
				new QName("conformance"))

			val inputPath = s"${parameter("inputPath")}"
			val outputPath = s"${mappings.getProperty("output")}/${System.currentTimeMillis()}"

			try {
				ooXml.transform(inputPath, outputPath, mappings)
				(true, outputPath)
			} catch {
				case _: Exception => (false, "")
			}
		}
}
