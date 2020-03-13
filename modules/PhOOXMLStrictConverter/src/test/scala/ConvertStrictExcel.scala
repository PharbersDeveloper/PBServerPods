import com.pharbers.convert.ConvertStrictExcel
import org.scalatest.FunSuite

class ConvertStrictExcel extends FunSuite {
	
	test("Strict") {
		println(sys.env("CONVERTOUTPUT"))
		val downloadPath = s"${sys.env("CONVERTINPUT")}/1574303592028_strict"
		val result = ConvertStrictExcel().exec(Map("inputPath" -> downloadPath))
		println(result)
	}
}
