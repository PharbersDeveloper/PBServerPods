package com.pharbers.ConvertSchema2Mongo.PhCommon.PhHDFS

import java.io._
import java.net.URI
import java.util._

import org.apache.commons.lang3.StringUtils
import org.apache.hadoop.conf.Configuration
import org.apache.hadoop.fs._
import org.apache.zookeeper.common.IOUtils

object HDFSUtil {
	val hdfsUrl = "hdfs://192.168.100.137:9000"
	var realUrl = ""

	def mkdir(dir : String) : Boolean = {
		var result = false
		if (StringUtils.isNoneBlank(dir)) {
			realUrl = hdfsUrl + dir
			val config = new Configuration()
			val fs = FileSystem.get(URI.create(realUrl), config)
			if (!fs.exists(new Path(realUrl))) {
				fs.mkdirs(new Path(realUrl))
			}
			fs.close()
			result = true
		}
		result
	}
	
	def deleteDir(dir : String) : Boolean = {
		var result = false
		if (StringUtils.isNoneBlank(dir)) {
			realUrl = hdfsUrl + dir
			val config = new Configuration()
			val fs = FileSystem.get(URI.create(realUrl), config)
			fs.delete(new Path(realUrl), true)
			fs.close()
			result = true
		}
		result
	}
	
	def listAll(dir : String) : List[String] = {
		val names : List[String] = new ArrayList[String]()
		if (StringUtils.isNoneBlank(dir)) {
			realUrl = hdfsUrl + dir
			val config = new Configuration()
			val fs = FileSystem.get(URI.create(realUrl), config)
			val stats = fs.listStatus(new Path(realUrl))
			for (i <- 0 to stats.length - 1) {
				if (stats(i).isFile) {
					names.add(stats(i).getPath.toString)
				} else if (stats(i).isDirectory) {
					names.add(stats(i).getPath.toString)
				} else if (stats(i).isSymlink) {
					names.add(stats(i).getPath.toString)
				}
			}
		}
		names
	}

	def uploadLocalFile2HDFS(localFile : String, hdfsFile : String) : Boolean = {
		var result = false
		if (StringUtils.isNoneBlank(localFile) && StringUtils.isNoneBlank(hdfsFile)) {
			realUrl = hdfsUrl + hdfsFile
			val config = new Configuration()
			val hdfs = FileSystem.get(URI.create(hdfsUrl), config)
			val src = new Path(localFile)
			val dst = new Path(realUrl)
			hdfs.copyFromLocalFile(src, dst)
			hdfs.close()
			result = true
		}
		result
	}
	
	def createNewHDFSFile(newFile : String, content : String) : Boolean = {
		var result = false
		if (StringUtils.isNoneBlank(newFile) && null != content) {
			realUrl = hdfsUrl + newFile
			val config = new Configuration()
			val hdfs = FileSystem.get(URI.create(realUrl), config)
			val os = hdfs.create(new Path(realUrl))
			os.write(content.getBytes("UTF-8"))
			os.close()
			hdfs.close()
			result = true
		}
		result
	}

	def deleteHDFSFile(hdfsFile : String) : Boolean = {
		var result = false
		if (StringUtils.isNoneBlank(hdfsFile)) {
			realUrl = hdfsUrl + hdfsFile
			val config = new Configuration()
			val hdfs = FileSystem.get(URI.create(realUrl), config)
			val path = new Path(realUrl)
			val isDeleted = hdfs.delete(path, true)
			hdfs.close()
			result = isDeleted
		}
		result
	}
	
	def readHDFSFile(hdfsFile : String) : Array[Byte] = {
		var result =  new Array[Byte](0)
		if (StringUtils.isNoneBlank(hdfsFile)) {
//			realUrl =hdfsUrl + hdfsFile
			val config = new Configuration()
			val hdfs = FileSystem.get(URI.create(hdfsUrl + hdfsFile), config)
			val path = new Path(hdfsUrl + hdfsFile)
			if (hdfs.exists(path)) {
				val inputStream = hdfs.open(path)
				val stat = hdfs.getFileStatus(path)
				val length = stat.getLen.toInt
				val buffer = new Array[Byte](length)
				inputStream.readFully(buffer)
				inputStream.close()
				hdfs.close()
				result = buffer
			}
		}
		result
	}

	def append(hdfsFile : String, content : String) : Boolean = {
		var result = false
		if (StringUtils.isNoneBlank(hdfsFile) && null != content) {
			realUrl = hdfsUrl + hdfsFile
			val config = new Configuration()
			config.set("dfs.client.block.write.replace-datanode-on-failure.policy", "NEVER")
			config.set("dfs.client.block.write.replace-datanode-on-failure.enable", "true")
			val hdfs = FileSystem.get(URI.create(realUrl), config)
			val path = new Path(realUrl)
			if (hdfs.exists(path)) {
				val inputStream = new ByteArrayInputStream(content.getBytes())
				val outputStream = hdfs.append(path)
				IOUtils.copyBytes(inputStream, outputStream, 4096, true)
				outputStream.close()
				inputStream.close()
				hdfs.close()
				result = true
			}
		} else {
			HDFSUtil.createNewHDFSFile(hdfsFile, content);
			result = true
		}
		result
	}
}
