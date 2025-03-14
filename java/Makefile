IMG_NAME=api-gateway
#VTAG=$(shell date +"%y%m%d-%H%M")
#VTAG=$(shell date +"%y%m%d")-$(shell git branch --show-current)
VTAG=$(shell date +"%y%m%d")
CT_NAME=${IMG_NAME}_ct
MAINTAINER=$(shell whoami)
CT_FILE=Containerfile
CT_PORT=8080
HOST_PORT=8080


build:
	podman build -f ${CT_FILE} -t ${MAINTAINER}/${IMG_NAME}:${VTAG}
	podman image prune -f
	podman builder prune -f
	podman system prune -f

run:
#	podman run -d --net internal-net -p ${HOST_PORT}:${CT_PORT} --name ${CT_NAME} ${MAINTAINER}/${IMG_NAME}:${VTAG}
	podman run -it --rm -p ${HOST_PORT}:${CT_PORT} --name ${CT_NAME} ${MAINTAINER}/${IMG_NAME}:${VTAG}

view-logs:
	podman logs ${CT_NAME}

run-app-locally:
	mvn spring-boot:run

init-maven-wrapper:
	mvn -N wrapper:wrapper -Dmaven=3.5.2
