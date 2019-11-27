package com.pharbers.process

import com.pharbers.factory.Listener

case class KafkaStartProcess() {
	def start(klst: List[Listener]): Unit = klst.foreach(_.start())
}
