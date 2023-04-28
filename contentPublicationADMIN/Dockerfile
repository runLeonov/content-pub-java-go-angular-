FROM maven AS build
COPY src /admin/home/app/src
COPY pom.xml /admin/home/app
RUN mvn -f /admin/home/app/pom.xml clean package

#
# Package stage
#
FROM openjdk
COPY --from=build /admin/home/app/target/spring-boot-mvc-0.0.1-SNAPSHOT.jar /usr/local/lib/app.jar
EXPOSE 8080
ENTRYPOINT ["java","-jar","/usr/local/lib/app.jar"]