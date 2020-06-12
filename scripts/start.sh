source conf/.env

bash scripts/compile.sh

bash scripts/build_images.sh

bash scripts/create_namespace.sh

bash scripts/create_configmap.sh

bash scripts/create_proxy.sh





