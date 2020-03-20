# prod
FROM maven:3.6-ibmjava-8-alpine

# 错误文件的输入、输出目录
WORKDIR /app/convert_excel/input
WORKDIR /app/convert_excel/output
ENV CONVERTINPUT=/app/convert_excel/input
ENV CONVERTOUTPUT=/app/convert_excel/output

ENV PHPRODSHOME=/app
ENV PHOOXMLCONF=${PHPRODSHOME}/conf/ooxml_config.json
ENV PHOSSCONF=${PHPRODSHOME}/conf/oss_config.json
#Scala micro-service-libs库需要，就离谱哦，lib包里面还有直接获取运行地址的
ENV PHA_CONF_HOME=""


WORKDIR /app/conf
#COPY ./ooxml_config.json ./
#COPY ./oss_config.json ./


WORKDIR /app


#COPY ./kafka_config.xml ./
#COPY ./kafka.broker1.truststore.jks ./
#COPY ./kafka.broker1.keystore.jks ./

COPY ./target/PhOOXMLStrictConverter.jar .

#ENTRYPOINT ["java", "-jar", "PhOOXMLStrictConverter.jar"]
ENTRYPOINT ["/bin/sh", "-c", "sleep 600s"]
