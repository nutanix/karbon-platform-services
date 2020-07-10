## Kafka 

The Kafka data interface allows the user to produce and consume data from the Kafka service in Xi IoT.

### Kafka-In

The *kafka-in* data interface can be used to consume data from the Kafka service running in Xi IoT. To start setting up this datasource, you will 
first need to expose Kafka on your edge. You can find documentation on how to expose kafka [here](https://github.com/nutanix/xi-iot/tree/master/services/kafka)

After you have exposed Kafka, you can create your data interface using the following YAML.

**kafka_in.yaml**
```yaml
kind: dataSource
name: kafka-in
svcDomain: <Service Domain>
protocol: DATAINTERFACE
type: Sensor
authType: CERTIFICATE
ifcInfo:
  class: DATAINTERFACE
  img: dataifc/kafka-in:v1.1
  kind: IN
  protocol: nats 
  ports:
    - name: rtmp
      port: 1935
edge: myprovideredge
fields:
- name: host
  topic: "kafkahost-<IP Address>"
- name: port
  topic: "kafkaport-<Port>"
- name: group
  topic: "kafkagroup-<Group>"
- name: topic
  topic: "<Topic>"
```

**Note:** Your Topic is the Kafka topic you wish to consume. 

Run the following command to create an instance of the datasource:
```console
xi-iot create -f kafka_in.yaml
```
* If the datasource was configured correctly, you should now see a datasource in your UI called **kafka-in**
    * You can now use this data source to pass Kafka messages to your Kubernetes Apps and Data Pipelines.  

You can now use this data source to pass Kafka messages to your Kubernetes Apps and Data Pipelines.