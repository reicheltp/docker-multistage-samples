# BUILD STAGE
FROM golang:alpine AS build

# git is required for go get
RUN apk --no-cache upgrade && \
    apk --no-cache add git

WORKDIR /go/src/app
COPY ./src .

RUN go-wrapper download
RUN go-wrapper install

# FINAL STAGE
FROM alpine

COPY --from=build /go/bin/app /bin/app

ENV GIN_MODE=release
ENV PORT=80

CMD ["app"]
