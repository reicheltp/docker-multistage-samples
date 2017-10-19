## BUILD STAGE
FROM maven:alpine AS build

WORKDIR /app
COPY . .

RUN mvn install

## FINAL STAGE
FROM openjdk:jre-alpine

WORKDIR /app
# Copy jar from build
COPY --from=build /app/target/*-fat.jar .

ENV PORT=80

# Launch the verticle
ENTRYPOINT ["sh", "-c"]
CMD ["exec java -jar *-fat.jar"]
