IMG_NAME=golang-grpc
#VTAG=$(shell date +"%y%m%d-%H%M")
VTAG=$(shell date +"%y%m%d")
CT_NAME=${IMG_NAME}_ct
MAINTAINER=$(shell whoami)
CT_FILE=Containerfile
CT_PORT=8080
HOST_PORT=8080


build:
	podman build -f ${CT_FILE} -t ${MAINTAINER}/${IMG_NAME}:${VTAG}
	podman image prune
	podman builder prune
	podman system prune

run:
	podman run -it --rm -p ${HOST_PORT}:${CT_PORT} --name ${CT_NAME} ${MAINTAINER}/${IMG_NAME}:${VTAG}

view-logs:
	podman logs ${CT_NAME}
