IMG_NAME=datastructures
#VTAG=$(shell date +"%y%m%d-%H%M")
VTAG=$(shell date +"%y%m%d")
CT_NAME=${IMG_NAME}_ct
MAINTAINER=$(shell whoami)
CT_FILE=Containerfile

build:
	podman build -f ${CT_FILE} -t ${MAINTAINER}/${IMG_NAME}:${VTAG}
	podman image prune
	podman builder prune
	podman system prune

run:
	podman run -it --rm --name ${CT_NAME} ${MAINTAINER}/${IMG_NAME}:${VTAG}

