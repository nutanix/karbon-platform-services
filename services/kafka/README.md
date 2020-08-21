# Kafka Data Service

We provide support for Kafka which is an extensible streaming platform. Clients can manage, publish and subscribe to topics using the native Kafka protocol.

Data Pipelines can use Kafka as destination. Similarly applications can use Kafka client of their choice to access the Kafka Data Service.

## Instantiate a Kafka Data Service

Services are enabled per project to be consumed by Pipelines and Kubernetes Applications within that project. Moreover Kafka can optionally be exposed for external clients.

A sophisticated data service like Kafka has many config options which a user might want to change from their default values. We allow various settings to be tuned for throughput, latency, availability or durability reasons. We treat all Kafka brokers as equals. That is all brokers will have same configuration settings.

You can configure services by using the project service API. The CLI supports this API which we'll demonstrate shortly. 

### A simple Kafka instance with default settings

Let's setup Kafka service with default settings:

```
kind: service
name: Kafka
project: KafkaTest
```

The service name must be Kafka, as we only support a single Kafka instance within a project. In that sense name is equal to type of service requested.
We automatically size the number of Kafka brokers to be equal to number of nodes in Service Domain. The Kafka cluster size adjusts if a node is added or removed.

The service can be instantiated using CLI as follows:

```
$ kpsctl create -f kafka-with-defaults.yaml
```

We can query back service instance in project Kafka:

```
kpsctl get service -p KafkaTest Kafka -o yaml
kind: service
name: Kafka
project: KafkaTest
serviceYaml: |
  null
```

You can see serviceYaml property which is null in config. That means default configuration is used. The YAML configuring Kafka is embedded into service config. 

You can update the configuration after creating it. Depending on the updates, Kafka cluster availability might be affected.

```
$ kpsctl update -f changed-kafka-config.yaml
```

If you no longer require the Kafka data service, you can delete it from the project.

```
$ kpsctl delete service -p KafkaTest Kafka
```

Of course, deleting the service might affect running applications and data pipelines.

### A Kafka instance with non-default settings

Here is a more complex YAML snippet where we configure the Kafka data service:

```
kind: service
name: Kafka
project: KafkaTest
serviceYaml: |
  apiVersion: sherlock.nutanix.com/v1
  kind: Kafka
  metadata:
    name: kafka
  spec:
    # Expose Kafka on node port 32092
    nodePort: 32092
    # Retain 1 megabyte
    logRetentinBytes: 1000000
    # Give each Kafka broker 4G of memory.
    kafkaMemory: 4G
    # Allocate 2 CPU cores for each Kafka broker.
    kafkaCPU: 2000m
    # Give each Zookeeper node 2G of memory.
    zookeeperMemory: 2G
    # Allocate 0.6 CPU cores for each ZK node.
    zookeeperCPU: 600m
    # Allocate 500GB storage for each Kafka broker.
    kafkaVolumeSize: 500Gi
```

## Access Kafka Data Service from Pipelines

Data Pipelines can specify Kafka as destination. Pipelines will use Kafka cluster enabled for the project. If no Kafka cluster has been explicitly enabled for project, a new cluster will be instantiated with default settings.

## Access Kafka Data Service from Kubernetes Applications

Applications are not required to discover Kafka clusters. We rely on applications templating for injecting actual Kafka broker list into applications:

```
apiVersion: apps/v1beta1
kind: Deployment
metadata:
  name: my-kafka-app
spec:
  template:
    metadata:
      labels:
        app: my-kakfa-app
    spec:
      containers:
      - name: some-container
        image: my-kafka-app
        env:
        - name: KAFKA_SERVER
          value: "{{.Services.Kafka.Endpoint}}"
        ports:
        - containerPort: 9000
```

KAFKA_SERVER will be expanded to Kafka API endpoint as a headless k8s service. That is, a DNS A entry with IPs for each broker in the cluster.

Kafka deployments come with their own Zookeeper Ensemble. Similar to Kafka API endpoint, Zookeeper's endpoint is injected into app YAML by using template parameter.

```
apiVersion: apps/v1beta1
kind: Deployment
metadata:
  name: kafdrop
spec:
  replicas: 1
  template:
    metadata:
      labels:
        app: kafdrop
    spec:
      containers:
      - name: kafdrop
        image: thomsch98/kafdrop
        env:
        - name: LISTEN
          value: "9000"
        - name: ZK_HOSTS
          value: "{{.Services.KafkaZookeeper.Endpoint}}"
        ports:
        - containerPort: 9000
---
apiVersion: v1
kind: Service
metadata:
  name: kafdrop
  annotations:
    sherlock.nutanix.com/http-ingress-path: /kafdrop
spec:
  ports:
  - port: 9000
  selector:
    app: kafdrop

```

This app YAML starts Kafdrop, a Kafka UI on any Service Domain.

## Access Kafka Data Service from external clients

External clients must specify a list of Kafka brokers (IP:port) as usual for Kafka clients. We publish list of Kafka endpoints in UI.

To see externally exposed Kafka brokers for your project: From the Projects page in the Karbon Platform Services for IoT cloud management console, click your project name. Open Data Services > Kafka and click the Deployments tab.

## Kafka settings

Following is a list of all settings and their defaults values. Change config default settings with caution since those affect availability and durability of Kafka data.

```
autoCreateTopicsEnable: true
defaultReplicationFactor: 2
deleteTopicEnable: true
kafkaCPU: 500m
kafkaMemory: 512M
kafkaVolumeSize: 6Gi
logFlushIntervalMS: 30000
logFlushIntervalMessages: 9223372036854775807
logFlushOffsetCheckpointIntervalMS: 60000
logFlushSchedulerIntervalMS: 9223372036854775807
logRetentionBytes: -1
logRetentionHours: 168
logRollHours: 168
logSegmentBytes: 1073741824
logSegmentDeleteDelayMS: 60000
messageMaxBytes: 52428800
minInsyncReplicas: 2
numPartitions: 6
numReplicaFetchers: 1
offsetsLoadBufferSize: 5242880
offsetsMetadataMaxBytes: 4096
offsetsRetentionCheckIntervalMS: 600000
offsetsTopicNumPartitions: 6
offsetsTopicReplicationFactor: 2
offsetsTopicRetentionMinutes: 1440
offsetsTopicSegmentBytes: 104857600
replicaFetchMaxBytes: 52428800
replicaFetchMinBytes: 1
replicaFetchWaitMaxMS: 500
replicaHighWatermarkCheckpointIntervalMS: 5000
replicaLagTimeMaxMS: 10000
uncleanLeaderElectionEnable: false
zookeeperCPU: 500m
zookeeperMemory: 512M
zookeeperReplicationFactor: 3
zookeeperVolumeSize: 2Gi
```