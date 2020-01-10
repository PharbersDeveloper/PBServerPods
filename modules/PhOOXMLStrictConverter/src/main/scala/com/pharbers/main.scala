package com.pharbers

import com.pharbers.listener.ConvertStrictExcelListener
import com.pharbers.process.KafkaStartProcess

object main extends App {
	val list = ConvertStrictExcelListener() :: Nil
	KafkaStartProcess().start(list)
	println("Consumer Start Complete")
}