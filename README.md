# drone-google-chat

[![Build Status](https://drone.seattleslow.com/api/badges/josmo/drone-google-chat/status.svg)](https://drone.seattleslow.com/josmo/drone-google-chat)

Drone plugin for sending google chat notifications. For the usage information and a listing of the available options please reference TBD.

## Build

Build the binary with the following commands:

```
go build
go test
```

## Docker

Build the docker image with the following commands:

```
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -tags netgo
docker build -t peloton/drone-google-chat .
```


## Usage

Execute from the working directory:

You need to create an incomming hook from the google chat room. The webhook is everything minus the &token=... 
and the token is just the token itself.

```
docker run --rm \
  -e GOOGLE_CHAT_WEBHOOK=https://dynamite.sandbox.googleapis.com/v1/rooms/... \
  -e GOOGLE_CHAT_TOKEN=sometoken
  -e DRONE_REPO_OWNER=octocat \
  -e DRONE_REPO_NAME=hello-world \
  -e DRONE_COMMIT_SHA=7fd1a60b01f91b314f59955a4e4d4e80d8edf11d \
  -e DRONE_COMMIT_BRANCH=master \
  -e DRONE_COMMIT_AUTHOR=octocat \
  -e DRONE_BUILD_NUMBER=1 \
  -e DRONE_BUILD_STATUS=success \
  -e DRONE_BUILD_LINK=http://github.com/octocat/hello-world \
  -e DRONE_TAG=1.0.0 \
  peloton/drone-google-chat
```
