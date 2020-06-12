source conf/.env

cp conf/storage_template.yaml conf/storage.yaml
sed -ri "s!HOST_DATA_DIR!$HOST_DATA_DIR!g" conf/storage.yaml
sed -ri "s/HOST_DATA_SERVER/$HOST_DATA_SERVER/g" conf/storage.yaml
sed -ri "s/STORAGE_SIZE/$STORAGE_SIZE/g" conf/storage.yaml


cp conf/service_template.yaml conf/service.yaml
sed -ri "s/EX_PORT/$EX_PORT/g" conf/service.yaml
sed -ri "s/DOCKER_PORT/$DOCKER_PORT/g" conf/service.yaml


cp conf/deployment_template.yaml conf/deployment.yaml
sed -ri "s/VERSION/$VERSION/g" conf/deployment_template.yaml
sed -ri "s!DOCKER_DATA_DIR!$DOCKER_DATA_DIR!g" conf/deployment_template.yaml
sed -ri "s/DOCKER_PORT/$DOCKER_PORT/g" conf/deployment_template.yaml

kubectl apply -f conf/storage.yaml
kubectl apply -f conf/service.yaml
kubectl apply -f conf/deployment.yaml