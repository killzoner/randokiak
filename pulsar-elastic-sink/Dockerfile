FROM apachepulsar/pulsar:2.5.0

# install elastic search connector
COPY install-pulsar-es-connector.sh /pulsar
COPY create-connector-pulsar.sh /pulsar
RUN chmod a+x /pulsar/*.sh
RUN /pulsar/install-pulsar-es-connector.sh

