package com.pharbers

import com.pharbers.convert.{OoXmlStrictConverter, PhReadMapping}
import com.pharbers.listener.ConvertStrictExcelListener
import com.pharbers.process.KafkaStartProcess
import javax.xml.namespace.QName
import javax.xml.stream.{XMLEventFactory, XMLInputFactory, XMLOutputFactory}

object main extends App {
//	val mappings = PhReadMapping.exec()
//	val ooXml = new OoXmlStrictConverter(XMLEventFactory.newInstance,
//		XMLInputFactory.newInstance,
//		XMLOutputFactory.newInstance,
//		new QName("conformance"))
//	val fileName = "sample.strict.xlsx"
//	val inputPath = s"/Users/qianpeng/Desktop/$fileName"
//	ooXml.transform(inputPath, s"${mappings.getProperty("output")}/$fileName", mappings)
	
	val list = ConvertStrictExcelListener() :: Nil
	KafkaStartProcess().start(list)
	println("Consumer Start")
}