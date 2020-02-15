ARG rootBuildFolder=/app

## BUILD
FROM golang:1.13-alpine3.11 AS builder

ARG rootBuildFolder
ARG appPackage
ARG appName

ENV _rootBuildFolder=${rootBuildFolder}
ENV _appPackage=${appPackage}
ENV _appName=${appName}

ENV _appFolder=${rootBuildFolder}/${appPackage}

RUN test -n "$_appPackage"
RUN test -n "$_appName"

RUN echo "rootBuildFolder $_rootBuildFolder"
RUN echo "_appPackage $_appPackage"
RUN echo "_appName $_appName"
RUN echo "appFolder $_appFolder"

# Set the Current Working Directory inside the container
WORKDIR ${_rootBuildFolder}

# Copy go mod and sum files
COPY . .

WORKDIR ${_appFolder}
# Build the Go app
RUN go mod download
RUN go build

## RUN
FROM alpine:3.11

ARG rootBuildFolder
ARG appPackage
ARG appName

ENV _rootBuildFolder=${rootBuildFolder}
ENV _appPackage=${appPackage}
ENV _appName=${appName}

ENV _appFolder=${rootBuildFolder}/${appPackage}

RUN test -n "$_appPackage"
RUN test -n "$_appName"

RUN echo "rootBuildFolder $_rootBuildFolder"
RUN echo "_appPackage $_appPackage"
RUN echo "_appName $_appName"
RUN echo "appFolder $_appFolder"

RUN apk --no-cache add ca-certificates
WORKDIR ${_rootBuildFolder}
COPY --from=builder /app/${_appPackage}/${_appName} .
RUN mv ${_appName} app
CMD ["./app"]  