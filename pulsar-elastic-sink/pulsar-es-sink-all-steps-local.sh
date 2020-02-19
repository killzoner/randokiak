MY_DIR=$(dirname $(readlink -f $0))
$MY_DIR/install-pulsar-es-connector.sh
$MY_DIR/create-connector-pulsar.sh