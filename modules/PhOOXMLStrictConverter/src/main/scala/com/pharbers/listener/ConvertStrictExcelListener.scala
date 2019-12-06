package com.pharbers.listener
import java.util.UUID

import com.alibaba.fastjson.JSON
import com.pharbers.convert.{OoXmlStrictConverter, PhReadMapping}
import com.pharbers.factory.{Convert, Listener}
import com.pharbers.kafka.consumer.PharbersKafkaConsumer
import com.pharbers.kafka.schema.ConvertExcel
import com.pharbers.oss.Oss
import com.pharbers.uitl.ThreadExecutor.ThreadExecutor
import javax.xml.namespace.QName
import javax.xml.stream.{XMLEventFactory, XMLInputFactory, XMLOutputFactory}
import org.apache.kafka.clients.consumer.ConsumerRecord
import com.pharbers.uitl.Http

import collection.JavaConverters._

class ConvertStrictExcel extends Convert {
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

case class ConvertStrictExcelListener() extends Listener {
	override def start(): Unit = {
		val pkc = new PharbersKafkaConsumer[String, ConvertExcel](
			"convert_excel_job" :: Nil,
			1000,
			Int.MaxValue, process
		)
		ThreadExecutor().execute(pkc)
	}
	
	val process: ConsumerRecord[String, ConvertExcel] => Unit = (record: ConsumerRecord[String, ConvertExcel]) => {
//		println("进入错误Handler")
		val uuid = UUID.randomUUID().toString
		val objectName = s"$uuid/${System.currentTimeMillis()}"

		val p = JSON.toJSONString(Map("assetId" -> record.value().getAssetId.toString).asJava, true)
		val response = Http.Post("http://localhost:8080/findFilePathWithId", p, "application/json").exec()
		val ossPath = JSON.parseObject(response).getString("ossPath")
		val downloadPath = PhReadMapping.exec().getProperty("input") + "/" + ossPath.substring(ossPath.lastIndexOf("/") + 1)
		Oss.down(downloadPath, ossPath)
		val result = new ConvertStrictExcel().exec(Map("inputPath" -> downloadPath))
		if (result._1) {
			Oss.upload(result._2, objectName)
			
			// TODO：这块儿有问题，不能一直等待返回值出现，如果崩了，就都崩溃了
			val updateResponse = Http.Post("http://localhost:8080/updateAssetVersion",
				JSON.toJSONString(Map("assetId" -> record.value().getAssetId.toString, "url" -> objectName).asJava, true),
				"application/json").exec()
			
			
			// TODO：暂时在我这里做，其实不在我这边进行重新提交，只是发送一个消息
//			Http.Post("http://localhost:8080/reCommitJobWithAssetId",
//				JSON.toJSONString(Map("assetId" -> JSON.parseObject(updateResponse).getString("assetId")).asJava, true),
//				"application/json").exec()
		} else {
			// TODO:再次进入错误队列
		}
	}
}