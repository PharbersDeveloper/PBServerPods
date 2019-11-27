package com.pharbers.listener

import java.io.{File, FileInputStream}
import java.util.UUID

import com.alibaba.fastjson.JSON
import com.aliyun.oss.OSSClientBuilder
import com.aliyun.oss.model.GetObjectRequest
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
//		record.value().getTraceId
//		record.value().getJobId
//		record.value().getType

		val uuid = UUID.randomUUID().toString
		val objectName = s"$uuid/${System.currentTimeMillis()}"

		val p = JSON.toJSONString(Map("traceId" -> record.value().getTraceId).asJava, true)
		val response = Http.Post("http://localhost:8080/findFilePathWithTraceId", p, "application/json").exec()
		val ossPath = JSON.parseObject(response).getString("ossPath")
		val downloadPath = PhReadMapping.exec().getProperty("input") + "/" + ossPath.substring(ossPath.lastIndexOf("/") + 1)
		Oss.down(downloadPath, ossPath)
		val result = new ConvertStrictExcel().exec(Map("inputPath" -> downloadPath))
		if (result._1) {
			Oss.upload(result._2, objectName)
			
			Http.Post("http://localhost:8080/updateAssetVersion",
				JSON.toJSONString(Map("traceId" -> record.value().getTraceId, "url" -> objectName).asJava, true),
				"application/json").exec()
			
			// TODO：这块儿应该发送kafka到调度中心（老铁那边）
			Http.Post("http://localhost:8080/reCommitJobWithTraceId",
				JSON.toJSONString(Map("traceId" -> record.value().getTraceId).asJava, true),
				"application/json").exec()
		} else {
			// TODO:再次进入错误队列
		}
	}
}

//  TODO: 测试用
//object OSS {
//	val endpoint = "oss-cn-beijing.aliyuncs.com"
//	val accessKeyId = "LTAIEoXgk4DOHDGi"
//	val accessKeySecret = "x75sK6191dPGiu9wBMtKE6YcBBh8EI"
//	val bucketName = "pharbers-sandbox"
//	def down(outPath: String, objectName: String): Unit = {
//		val ossClient = new OSSClientBuilder().build(endpoint, accessKeyId, accessKeySecret)
//		ossClient.getObject(new GetObjectRequest(bucketName, objectName), new File(outPath))
//		ossClient.shutdown()
//	}
//
//	def upload(inputPath: String, objectName: String): Unit = {
//		val ossClient = new OSSClientBuilder().build(endpoint, accessKeyId, accessKeySecret)
//		val inputStream = new FileInputStream(inputPath)
//		println(objectName)
//		ossClient.putObject(bucketName, objectName, inputStream)
//		ossClient.shutdown()
//	}
//}

//object test extends App {
//	val traceId = "08300-8033-494a-9cb5-3acee"
//	val response = Http.Post("http://localhost:8080/findFilePathWithTraceId",
//		JSON.toJSONString(Map("traceId" -> traceId).asJava, true),
//		"application/json").exec()
//
//	val ossPath = JSON.parseObject(response).getString("ossPath")
//	val downloadPath = PhReadMapping.exec().getProperty("input") + "/" + ossPath.substring(ossPath.lastIndexOf("/") + 1)
//	OSS.down(downloadPath, ossPath)
//	val result = new ConvertStrictExcel().exec(Map("inputPath" -> downloadPath))
//	val uuid = UUID.randomUUID().toString
//	val objectName = s"$uuid/${System.currentTimeMillis()}"
//	if (result._1) {
////		OSS.upload(result._2, objectName)
//		println("上传啦")
//	}
//
//	Http.Post("http://localhost:8080/updateAssetVersion",
//		JSON.toJSONString(Map("traceId" -> traceId, "url" -> objectName).asJava, true),
//		"application/json").exec()
//}
