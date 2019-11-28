package com.pharbers.convert;

import java.io.*;
import java.util.ArrayList;
import java.util.Collections;
import java.util.Iterator;
import java.util.Properties;
import java.util.zip.ZipEntry;
import java.util.zip.ZipInputStream;
import java.util.zip.ZipOutputStream;

import javax.xml.namespace.QName;
import javax.xml.stream.XMLEventFactory;
import javax.xml.stream.XMLEventReader;
import javax.xml.stream.XMLEventWriter;
import javax.xml.stream.XMLInputFactory;
import javax.xml.stream.XMLOutputFactory;
import javax.xml.stream.events.Attribute;
import javax.xml.stream.events.EndElement;
import javax.xml.stream.events.Namespace;
import javax.xml.stream.events.StartElement;
import javax.xml.stream.events.XMLEvent;

public class OoXmlStrictConverter {

	private static XMLInputFactory XIF;
	private static XMLEventFactory XEF;
	private static XMLOutputFactory XOF;
	private static QName CONFORMANCE;

	public OoXmlStrictConverter(XMLEventFactory XEF, XMLInputFactory XIF, XMLOutputFactory XOF, QName CONFORMANCE) {
		OoXmlStrictConverter.XEF = XEF;
		OoXmlStrictConverter.XIF = XIF;
		OoXmlStrictConverter.XOF = XOF;
		OoXmlStrictConverter.CONFORMANCE = CONFORMANCE;
		OoXmlStrictConverter.XOF.setProperty(XMLOutputFactory.IS_REPAIRING_NAMESPACES, true);
	}

	public void transform(final String inFile, final String outFile, final Properties mappings) throws Exception {
		System.out.println("transforming " + inFile + " to " + outFile);
		try(ZipInputStream zis = new ZipInputStream(new FileInputStream(inFile));
		    ZipOutputStream zos = new ZipOutputStream(new FileOutputStream(outFile))) {
			ZipEntry ze;
			while((ze = zis.getNextEntry()) != null) {
				ZipEntry newZipEntry = new ZipEntry(ze.getName());
				zos.putNextEntry(newZipEntry);
				FilterInputStream filterIs = new FilterInputStream(zis) {
					@Override
					public void close() throws IOException {
					}
				};
				FilterOutputStream filterOs = new FilterOutputStream(zos) {
					@Override
					public void close() throws IOException {
					}
				};
				if(isXml(ze.getName())) {
					try {
						XMLEventReader xer = XIF.createXMLEventReader(filterIs);
						XMLEventWriter xew = XOF.createXMLEventWriter(filterOs);
						int depth = 0;
						while(xer.hasNext()) {
							XMLEvent xe = xer.nextEvent();
							if(xe.isStartElement()) {
								StartElement se = xe.asStartElement();
								xe = XEF.createStartElement(updateQName(se.getName(), mappings),
										processAttributes(se.getAttributes(), mappings, se.getName().getNamespaceURI(), (depth == 0)),
										processNamespaces(se.getNamespaces(), mappings));
								depth++;
							} else if(xe.isEndElement()) {
								EndElement ee = xe.asEndElement();
								xe = XEF.createEndElement(updateQName(ee.getName(), mappings),
										processNamespaces(ee.getNamespaces(), mappings));
								depth--;
							}
							xew.add(xe);
						}
						xer.close();
						xew.close();
					} catch(Throwable t) {
						throw new IOException("Problem paraing " + ze.getName(), t);
					}
				} else {
					copy(filterIs, filterOs);
				}
				zis.closeEntry();
				zos.closeEntry();
			}
		}
	}

	private static boolean isXml(final String fileName) {
		if(isNotBlank(fileName)) {
			int pos = fileName.lastIndexOf(".");
			if(pos != -1) {
				String ext = fileName.substring(pos + 1).toLowerCase();
				return ext.equals("xml") || ext.equals("vml") || ext.equals("rels");
			}
		}
		return false;
	}


	private static Iterator<Attribute> processAttributes(final Iterator<Attribute> iter,
	                                                     final Properties mappings, final String elementNamespaceUri, final boolean rootElement) {
		ArrayList<Attribute> list = new ArrayList<>();
		while(iter.hasNext()) {
			Attribute att = iter.next();
			QName qn = updateQName(att.getName(), mappings);
			if(rootElement && mappings.containsKey(elementNamespaceUri) && att.getName().equals(CONFORMANCE)) {
				//drop attribute
			} else {
				String newValue = att.getValue();
				for(String key : mappings.stringPropertyNames()) {
					if(att.getValue().startsWith(key)) {
						newValue = att.getValue().replace(key, mappings.getProperty(key));
//						att.getValue().replace(key, mappings.getProperty(key));
						break;
					}
				}
				list.add(XEF.createAttribute(qn, newValue));
			}
		}
		return Collections.unmodifiableList(list).iterator();
	}

	private static Iterator<Namespace> processNamespaces(final Iterator<Namespace> iter,
	                                                     final Properties mappings) {
		ArrayList<Namespace> list = new ArrayList<>();
		while(iter.hasNext()) {
			Namespace ns = iter.next();
			if(!ns.isDefaultNamespaceDeclaration() && !mappings.containsKey(ns.getNamespaceURI())) {
				list.add(ns);
			}
		}
		return Collections.unmodifiableList(list).iterator();
	}

	private static QName updateQName(QName qn, Properties mappings) {
		String namespaceUri = qn.getNamespaceURI();
		if(isNotBlank(namespaceUri)) {
			String mappedUri = mappings.getProperty(namespaceUri);
			if(mappedUri != null) {
				qn = isBlank(qn.getPrefix()) ? new QName(mappedUri, qn.getLocalPart())
						: new QName(mappedUri, qn.getLocalPart(), qn.getPrefix());
			}
		}
		return qn;
	}

	private static boolean isBlank(final String str) {
		return str == null || str.trim().length() == 0;
	}

	private static boolean isNotBlank(final String str) {
		return !isBlank(str);
	}

	private static void copy(InputStream inp, OutputStream out) throws IOException {
		byte[] buff = new byte[4096];
		int count;
		while ((count = inp.read(buff)) != -1) {
			if (count > 0) {
				out.write(buff, 0, count);
			}
		}
	}
}