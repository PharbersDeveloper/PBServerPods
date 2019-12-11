package com.pharbers.oss

import java.io.{File, FileInputStream}

import com.aliyun.oss.{OSS, OSSClientBuilder}
import com.aliyun.oss.model.GetObjectRequest
import com.pharbers.uitl.Conf

case class Oss() {
	lazy val conf: Map[String, String] = Conf.Config.loadOssConfig()("sandbox")
	
	lazy val endpoint: String = conf("endpoint")
	lazy val accessKeyId: String = conf("accessKeyId")
	lazy val accessKeySecret: String = conf("accessKeySecret")
	lazy val bucketName: String = conf("bucketName")
	def down(outPath: String, objectName: String): Unit = {
		val downOssClient: OSS = new OSSClientBuilder().build(endpoint, accessKeyId, accessKeySecret)
		downOssClient.getObject(new GetObjectRequest(bucketName, objectName), new File(outPath))
		downOssClient.shutdown()
	}
	
	def upload(inputPath: String, objectName: String): Unit = {
		val uploadOssClient = new OSSClientBuilder().build(endpoint, accessKeyId, accessKeySecret)
		val inputStream = new FileInputStream(inputPath)
		uploadOssClient.putObject(bucketName, objectName, inputStream)
		uploadOssClient.shutdown()
	}
}
