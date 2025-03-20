IMG_NAME=api-gateway
#VTAG=$(shell date +"%y%m%d-%H%M")
#VTAG=$(shell date +"%y%m%d")-$(shell git branch --show-current)
VTAG=$(shell date +"%y%m%d")
CT_NAME=${IMG_NAME}_ct
MAINTAINER=$(shell whoami)
CT_FILE=Containerfile
CT_PORT=8080
HOST_PORT=8080
PLATFORM=linux/amd64,linux/arm64


build:
	podman build -f ${CT_FILE} -t ${MAINTAINER}/${IMG_NAME}:${VTAG}
	podman image prune -f
	podman builder prune -f
	podman system prune -f

build-multiarch:
#	podman build --jobs=4 --platform=${PLATFORM} --manifest ${MAINTAINER}/${IMG_NAME}:${VTAG} -f Containerfile
	podman manifest create ${IMG_NAME}:${VTAG}
	podman build --platform=linux/amd64 -t ${IMG_NAME}:${VTAG}-amd64 -f Containerfile
	podman build --platform=linux/arm64 -t ${IMG_NAME}:${VTAG}-arm64 -f Containerfile
	podman manifest add ${IMG_NAME}:${VTAG} localhost/${IMG_NAME}:${VTAG}-amd64
	podman manifest add ${IMG_NAME}:${VTAG} localhost/${IMG_NAME}:${VTAG}-arm64
	podman manifest push --all ${IMG_NAME}:${VTAG} docker://docker.io/dnsobc/${IMG_NAME}:${VTAG}
	podman image prune -f
	podman builder prune -f
	podman system prune -f

run-multiarch:
#	podman run -it --rm -p ${HOST_PORT}:${CT_PORT} --name ${CT_NAME} localhost/${MAINTAINER}/${IMG_NAME}:${VTAG}
	podman run -it --rm -p ${HOST_PORT}:${CT_PORT} --name ${CT_NAME} localhost/${IMG_NAME}:${VTAG}

run:
#	podman run -d --net internal-net -p ${HOST_PORT}:${CT_PORT} --name ${CT_NAME} ${MAINTAINER}/${IMG_NAME}:${VTAG}
	podman run -it --rm -p ${HOST_PORT}:${CT_PORT} --name ${CT_NAME} ${MAINTAINER}/${IMG_NAME}:${VTAG}

view-logs:
	podman logs ${CT_NAME}

run-app-locally:
	mvn spring-boot:run

init-maven-wrapper:
	mvn -N wrapper:wrapper -Dmaven=3.5.2
