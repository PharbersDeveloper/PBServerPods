package com.pharbers.uitl.Http

import scalaj.http.{Http, HttpOptions}

case class Post(url: String, body: String, contentType: String) {
	def exec(): String = {
		val rsp = Http(url).postData(body)
			.header("Content-Type", contentType)
			.header("Charset", "UTF-8")
			.option(HttpOptions.readTimeout(10000)).asString
		rsp.code match {
			case 200 => rsp.body
			case _ => ""
		}
	}
}
