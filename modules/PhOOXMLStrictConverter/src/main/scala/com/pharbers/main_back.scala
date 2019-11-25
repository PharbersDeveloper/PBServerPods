package com.pharbers

import com.pharbers.convert.{PhReadMapping, PhTransform}

object main extends App {
	val mappings = PhReadMapping.exec()
	val t = new PhTransform("/Users/qianpeng/Desktop/sample.strict.xlsx",
		"/Users/qianpeng/Desktop/bb.xlsx",
		mappings)
	
	t.transform()
}
