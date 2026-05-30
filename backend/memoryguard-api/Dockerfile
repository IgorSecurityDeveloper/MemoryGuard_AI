FROM eclipse-temurin:21-jdk

WORKDIR /app

COPY target/quarkus-app/lib/ ./lib/
COPY target/quarkus-app/*.jar ./
COPY target/quarkus-app/app/ ./app/
COPY target/quarkus-app/quarkus/ ./quarkus/

CMD ["java", "-jar", "quarkus-run.jar"]