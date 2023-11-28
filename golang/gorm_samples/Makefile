IMG_NAME=psgres
#VTAG=$(shell date +"%y%m%d-%H%M")
VTAG=$(shell date +"%y%m%d")
CT_NAME=${IMG_NAME}_ct
MAINTAINER=$(shell whoami)
CT_FILE=Containerfile
CT_PORT=8080
HOST_PORT=8080

test:
	echo ${IMG_NAME}
	echo ${IMG_VERSION}
	echo ${CT_NAME}
	echo ${MAINTAINER}
	echo ${CT_FILE}

build:
	podman build -f ${CT_FILE} -t ${MAINTAINER}/${IMG_NAME}:${VTAG}

run:
	podman run -it --rm -p ${HOST_PORT}:${CT_PORT} --name ${CT_NAME} ${MAINTAINER}/${IMG_NAME}:${VTAG}

view-logs:
	podman logs ${CT_NAME}

destroy:
	podman stop cass_cluster
	podman rm cass_cluster
