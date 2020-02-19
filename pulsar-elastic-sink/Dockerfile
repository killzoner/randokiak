FROM apachepulsar/pulsar:2.5.0

COPY install-pulsar-es-connector.sh /pulsar
COPY create-connector-pulsar.sh /pulsar
COPY pulsar-es-sink.yml /pulsar
RUN chmod a+x /pulsar/*.sh
RUN /pulsar/install-pulsar-es-connector.sh
