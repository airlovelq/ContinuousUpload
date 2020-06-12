source ./conf/.env

kubectl delete configmap continuous-upload-proxy-configmaps

kubectl create configmap continuous-upload-proxy-configmaps \
    --namespace=continuous-upload \
    --from-literal=VERSION=$VERSION \
    --from-literal=STORAGE_SIZE=$STORAGE_SIZE \
    --from-literal=DOCKER_DATA_DIR=$DOCKER_DATA_DIR \
    --from-literal=DOCKER_PORT=$DOCKER_PORT \
    --from-literal=HOST_DATA_SERVER=$HOST_DATA_SERVER \
    --from-literal=HOST_DATA_DIR=$HOST_DATA_DIR \
    --from-literal=EX_PORT=$EX_PORT