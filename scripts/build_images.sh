source conf/.env

docker build -t continuous-upload-proxy:$VERSION -f ./dockerfiles/worker.dockerfile \
    --build-arg DOCKER_DATA_DIR=$DOCKER_DATA_DIR $PWD || exit 1