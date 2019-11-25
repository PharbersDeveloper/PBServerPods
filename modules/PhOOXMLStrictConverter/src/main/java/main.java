import javax.xml.stream.XMLOutputFactory;
import java.util.Properties;

class Main {
	public static void main(String[] args) {
		OoXmlStrictConverter.XOF.setProperty(XMLOutputFactory.IS_REPAIRING_NAMESPACES, true);
		try {
			Properties mappings = OoXmlStrictConverter.readMappings();
			System.out.println("loaded mappings entries=" + mappings.size());
			OoXmlStrictConverter.transform(
					"/Users/qianpeng/Desktop/sample.strict.xlsx",
					"/Users/qianpeng/Desktop/aa.xlsx", mappings);
		} catch(Throwable t) {
			t.printStackTrace();
		}
	}
}

