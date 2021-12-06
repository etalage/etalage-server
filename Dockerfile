FROM openjdk:8-alpine

COPY target/uberjar/etalage.jar /etalage/app.jar

EXPOSE 3000

CMD ["java", "-jar", "/etalage/app.jar"]
