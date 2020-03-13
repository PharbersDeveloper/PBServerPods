import com.pharbers.convert.ConvertXls2Xlsx
import org.scalatest.FunSuite

class ConvertXls extends FunSuite {
	test("xls") {
		println(sys.env("CONVERTOUTPUT"))
		val downloadPath = s"${sys.env("CONVERTINPUT")}/1574303607650"
		val result = ConvertXls2Xlsx().exec(Map("inputPath" -> downloadPath))
		println(result)
	}
}
