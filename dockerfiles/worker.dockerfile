FROM ubuntu:18.04

ARG DOCKER_DATA_DIR
WORKDIR /root

RUN mkdir $DOCKER_DATA_DIR
COPY src/ContinuousUpload /root/ContinuousUpload

CMD ["ContinuousUpload"]