package com.pharbers.factory

trait Convert {
	def exec(parameter: Map[String, Any]): (Boolean, String)
}
