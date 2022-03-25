#!/bin/bash

usage() {
cat <<EOF

usage:
    --image-name sockjs
    --version v0.1
    --image-repo xishengcai
    -h print help

    example: sh build.sh --image-name lsh-mcp-lcs-out-gateway --version dev --image-repo xlauncher

EOF
}
BIN_FILE="sockjs"

#echo "set build info"
#GIT_COMMIT=$(git rev-parse --short HEAD || echo "GitNotFound")
#BUILD_TIME=$(date +%FT%T%z)

#echo $GIT_COMMIT $BUILD_TIME
#LDFLAGS="-X main.GitCommit=${GIT_COMMIT} -X main.BuildTime=${BUILD_TIME}"

buildProgram(){
  CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./bin/${BIN_FILE} ./main.go

  if [ $? -ne 0 ]; then
    echo "compile failed"
    exit 1
  fi
  echo "***** compile success *****"
}

set -e

buildImage() {
  docker build -t ${IMAGE_REPO}/${IMAGE_NAME}:${VERSION} ./
  echo "***** docker build success *****"
}

pushImage() {

  docker push ${IMAGE_REPO}/${IMAGE_NAME}:${VERSION}
  if [ $? -ne 0 ]; then
    echo "*****  docker push failed *****"
    exit 1
  fi
  echo "***** docker push success *****"
  docker rmi ${IMAGE_REPO}/${IMAGE_NAME}:${VERSION}
}

main() {
  buildProgram
  buildImage
  pushImage
}

if [ "$1" = "-h" ]; then
  usage
  exit 0
fi

while [ $# -gt 0 ]
do
    key="$1"
    case $key in
        --image-name)
            export IMAGE_NAME=$2
            shift
        ;;
        --version)
            export VERSION=$2
            shift
        ;;
        --image-repo)
            export IMAGE_REPO=$2
            shift
        ;;
        *)
            echo "unknown option [$key]"
            usage
            exit 1
        ;;
    esac
    shift
done

if [ -z "$IMAGE_NAME" ]; then
  IMAGE_NAME='sockjs'
fi

if [ -z "$VERSION" ]; then
  VERSION='v12'
fi

if [ -z "$IMAGE_REPO" ]; then
  IMAGE_REPO='xishengcai'
fi

main
