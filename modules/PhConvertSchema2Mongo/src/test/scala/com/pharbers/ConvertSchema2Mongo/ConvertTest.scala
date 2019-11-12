package com.pharbers.ConvertSchema2Mongo

import com.pharbers.ConvertSchema2Mongo.PhHandler.{PhConvert2MetaDataHandler, PhConvert2MongoHandler}
import org.scalatest.FunSuite

class ConvertTest extends FunSuite{
	test("convert 2 mongo") {
		new PhConvert2MongoHandler()
	}
	
	test("convert 2 meta data") {
		new PhConvert2MetaDataHandler()
	}
}
