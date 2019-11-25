package com.pharbers.convert

import java.io.{FileInputStream, FileOutputStream, FilterInputStream, FilterOutputStream, InputStream, OutputStream}
import java.util.Properties
import java.util.zip.{ZipEntry, ZipInputStream, ZipOutputStream}

import javax.xml.namespace._
import javax.xml.stream._
import javax.xml.stream.events._

import scala.collection.JavaConversions._

class PhTransform(inFile: String, outFile: String, mapping: Properties) {
	
	final val XEF = XMLEventFactory.newInstance()
	final val XIF = XMLInputFactory.newInstance()
	final val XOF = XMLOutputFactory.newInstance()
	final val CONFORMANCE = new QName("conformance")
	
	def exec(): Unit = {
		XOF.setProperty(XMLOutputFactory.IS_REPAIRING_NAMESPACES, true)
		transform()
	}
	
	def transform(): Unit = {
		
		try {
			val zis = new ZipInputStream(new FileInputStream(inFile))
			val zos = new ZipOutputStream(new FileOutputStream(outFile))
			var ze: ZipEntry = null
			def getNextEntry: ZipEntry = {
				ze = zis.getNextEntry
				ze
			}
			while (getNextEntry != null) {
				val newZipEntry = new ZipEntry(ze.getName)
				zos.putNextEntry(newZipEntry)
				
				val filterIs = new FilterInputStream(zis) {
					override def close(): Unit = {}
				}
				val filterOs = new FilterOutputStream(zos) {
					override def close(): Unit = {}
				}
				
				if(isXml(ze.getName)) {
					try {
						val xer = XIF.createXMLEventReader(filterIs)
						val xew = XOF.createXMLEventWriter(filterOs)
						var depth = 0
						while (xer.hasNext) {
							var xe = xer.nextEvent()
							if (xe.isStartElement) {
								val se = xe.asStartElement()
								xe = XEF.createStartElement(updateQName(se.getName, mapping),
								processAttributes(se.getAttributes.toIterator.asInstanceOf[Iterator[Attribute]], mapping, se.getName.getNamespaceURI, depth == 0),
								processNamespaces(se.getNamespaces.toIterator.asInstanceOf[Iterator[Namespace]], mapping))
								depth = depth + 1
							} else if(xe.isEndElement) {
								val ee = xe.asEndElement()
								xe = XEF.createEndElement(updateQName(ee.getName, mapping),
								processNamespaces(ee.getNamespaces.toIterator.asInstanceOf[Iterator[Namespace]], mapping))
								depth = depth - 1
							}
							xew.add(xe)
						}
						xer.close()
						xew.close()
					} catch {
						case e: Exception => println(e.getMessage); throw new Exception(e.getMessage)
					}
				} else {
					copy(filterIs, filterOs)
				}
				zis.closeEntry()
				zos.closeEntry()
			}
		} catch {
			case e: Exception => println(e.getMessage)
		}
	}
	
	def processAttributes(iter: Iterator[Attribute],
	                      mappings: Properties,
	                      elementNamespaceUri: String,
	                      rootElement: Boolean): Iterator[Attribute] = {
		var list: List[Attribute] = Nil
		while (iter.hasNext) {
			val att = iter.next()
			val qn = updateQName(att.getName, mappings)
			if(rootElement && mappings.containsKey(elementNamespaceUri) && att.getName == CONFORMANCE) {
				//drop attribute
			} else {
				var newValue = att.getValue
				import scala.util.control.Breaks._
				breakable {
					mappings.stringPropertyNames().foreach{ key =>
						if (att.getValue.startsWith(key)) {
							newValue = att.getValue.replace(key, mappings.getProperty(key))
//							att.getValue.replace(key, mappings.getProperty(key))
							break()
						}
					}
				}
				list = list :+ XEF.createAttribute(qn, newValue)
			}
		}
		list.toIterator
	}
	
	def processNamespaces(iter: Iterator[Namespace], mappings: Properties): Iterator[Namespace] = {
		var list: List[Namespace] = Nil
		while(iter.hasNext) {
			val ns = iter.next()
			if(!ns.isDefaultNamespaceDeclaration && !mappings.containsKey(ns.getNamespaceURI)) {
				list = list :+ ns
			}
		}
		list.toIterator
		
//		 iter.
//			 toList.
//			 filter(ite => !ite.isDefaultNamespaceDeclaration && !mappings.containsKey(ite.getNamespaceURI)).
//			 toIterator
	}
	
	def updateQName(qn: QName, mappings: Properties): QName = {
		var nqn = qn
		val namespaceUri = qn.getNamespaceURI
		if (isNotBlank(namespaceUri)) {
			val mappedUri = mappings.getProperty(namespaceUri)
			if(mappedUri != null) {
				if(isBlank(qn.getPrefix)) {
					nqn = new QName(mappedUri, qn.getLocalPart)
				} else {
					nqn = new QName(mappedUri, qn.getLocalPart, qn.getPrefix)
				}
			}
		}
		nqn
	}
	
	def isXml(fileName: String): Boolean = {
		var flag = false
		if(isNotBlank(fileName)) {
			val pos = fileName.lastIndexOf(".")
			if (pos != -1) {
				val ext = fileName.substring(pos + 1).toLowerCase()
				flag = ext == "xml" || ext == "vml" || ext == "rels"
			}
		}
		flag
	}
	
	def isBlank(str: String): Boolean = str == null || str.trim.length == 0
	
	def isNotBlank(str: String): Boolean = !isBlank(str)
	
	def copy(inp: InputStream, out: OutputStream): Unit = {
		val buff = new Array[Byte](4 * 1024)
		var count = 0
		while (count != -1) {
			count = inp.read(buff)
			if (count > 0) {
				out.write(buff, 0, count)
			}
		}
	}
}
