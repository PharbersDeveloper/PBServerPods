import com.pharbers.kafka.consumer.PharbersKafkaConsumer
import com.pharbers.uitl.ThreadExecutor.ThreadExecutor
import org.apache.avro.specific.SpecificRecord
import org.apache.kafka.clients.consumer.ConsumerRecord
import org.scalatest.FunSuite

class ConsumerTest extends FunSuite {
	test("consumer test") {
		val process: ConsumerRecord[String, SpecificRecord] => Unit = (record: ConsumerRecord[String, SpecificRecord]) => {
			println(record.value())
		}
		
		val pkc = new PharbersKafkaConsumer[String, SpecificRecord](
			"oss_task_submit" :: Nil,
			1000,
			Int.MaxValue, process
		)
		ThreadExecutor().execute(pkc)
		
		
		
	}
}
