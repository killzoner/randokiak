# Pulsar deploy validation

Use <http://pulsar-ip:pulsar-port> in manager <http://pulsar-manager-ip:pulsar-manager-port>
with pulsar/pulsar credentials

Test startup with commands :

`pulsar-client --url pulsar://<pulsar-ip>:<pulsar-port> produce my-topic --messages "hello-pulsar"`
For example: `pulsar-client --url pulsar://172.17.0.2:32757 produce my-topic --messages "hello-pulsar"`

`pulsar-client --url pulsar://<pulsar-ip>:<pulsar-port> consume my-topic -s "first-subscription"`
For example: `pulsar-client --url pulsar://172.17.0.2:32757 consume my-topic -s "first-subscription"`

to be sure broker has finished starting and everything is ok
