package com.pharbers.uitl.ThreadExecutor

import java.util.concurrent.{CountDownLatch, ExecutorService, Executors}

object ThreadExecutor {
    var executorService: Option[ExecutorService] = None
    val count = new CountDownLatch(1)
    def apply(): ExecutorService = {
        executorService match {
            case None => executorService = Some(Executors.newFixedThreadPool(10))
            case _ =>
        }
        executorService.get
    }

    def waitForShutdown(): Unit ={
        count.await()
    }

    def shutdown(): Unit ={
        count.countDown()
        executorService.get.shutdown()
    }
}
