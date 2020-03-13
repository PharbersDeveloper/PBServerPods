FROM maven:3.6-ibmjava-8-alpine as build

MAINTAINER Pharbers "pqian@pharbers.com"

# 错误文件的输入、输出目录
WORKDIR /app/convert_excel/input
WORKDIR /app/convert_excel/output
ENV CONVERTINPUT=/app/convert_excel/input
ENV CONVERTOUTPUT=/app/convert_excel/output

WORKDIR /app
ENV MICROSERVICELIBS=/app/micro-service-libs
ENV KAFKALIB=${MICROSERVICELIBS}/kafka

WORKDIR /app/PBServerPods
ENV PHPRODSHOME=/app/PBServerPods
ENV PHOOXMLSTRICTCONVERTER=${PHPRODSHOME}/modules/PhOOXMLStrictConverter
ENV PHOOXMLCONF=${PHPRODSHOME}/conf/ooxml_config.json
ENV PHOSSCONF=${PHPRODSHOME}/conf/oss_config.json

#Scala micro-service-libs库需要，就离谱哦，lib包里面还有直接获取运行地址的
ENV PHA_CONF_HOME=""

COPY ./PBServerPods ${PHPRODSHOME}
COPY ./micro-service-libs ${MICROSERVICELIBS}

WORKDIR ${MICROSERVICELIBS}
RUN mvn clean install

WORKDIR ${KAFKALIB}
RUN mvn clean install

WORKDIR ${PHOOXMLSTRICTCONVERTER}
RUN mvn clean package -DskipTests

WORKDIR ${PHOOXMLSTRICTCONVERTER}/target


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

# 提取资源文件
COPY --from=0 /app/PBServerPods/conf/ooxml_config.json ./
COPY --from=0 /app/PBServerPods/conf/oss_config.json ./

WORKDIR /app

COPY --from=0  /app/PBServerPods/pharbers_config/kafka_config.xml ./
COPY --from=0  /app/PBServerPods/conf/kafka.broker1.truststore.jks ./
COPY --from=0  /app/PBServerPods/conf/kafka.broker1.keystore.jks ./
COPY --from=0  /app/PBServerPods/modules/PhOOXMLStrictConverter/target/PhOOXMLStrictConverter.jar ./

ENTRYPOINT ["java", "-jar", "PhOOXMLStrictConverter.jar"]

#ENTRYPOINT cp /app/PBServerPods/pharbers_config/kafka_config.xml /app/PBServerPods/modules/PhOOXMLStrictConverter/target \
#	&& cp /app/PBServerPods/conf/kafka.broker1.truststore.jks /app/PBServerPods/modules/PhOOXMLStrictConverter/target \
#	&& cp /app/PBServerPods/conf/kafka.broker1.keystore.jks /app/PBServerPods/modules/PhOOXMLStrictConverter/target \
#	&& java -jar /app/PBServerPods/modules/PhOOXMLStrictConverter/target/PhOOXMLStrictConverter.jar
