
# XI IOT - DATA PIPELINES - GETTING STARTED GUIDE

## Xi IoT Overview

The Nutanix Xi IoT platform delivers local compute and AI for IoT edge devices, converging the edge and cloud into one seamless data processing platform. The Xi IoT platform eliminates complexity, accelerates deployments, and elevates developers to focus on the business logic powering IoT applications and services. Now developers can use a low-code development platform to create application software via APIs instead of arduous programming methods.

## Introducing Data Pipelines

The main steps in this guide are excerpts from the [Xi IoT Infrastructure Admin Guide](https://portal.nutanix.com/page/documents/details/?targetId=Xi-IoT-Infra-Admin-Guide:Xi-IoT-Infra-Admin-Guide), available from the Nutanix Support Portal.

Data Pipelines are paths for data that include:
* **Input**. An existing data source or real-time data stream.
* **Transformation**. Code block such as a script defined in a Function to process or transform input data.
* **Output**. A destination for your data. Publish data to the cloud or cloud data service (such as AWS Simple Queue Service) at the edge.

They also enable you to process and transform captured data for further consumption or processing.

Data pipelines have the following components used in the examples in this guide:

* Data Sources (defined as MQTT in this guide)
* Runtime Environments
* Functions

### Using MQTT Data Sources in Data Pipelines

#### What is MQTT?

If you are looking to understand the internals of how MQTT works, please read the 10 part series on [MQTT Essentials](https://www.hivemq.com/tags/mqtt-essentials/) by HiveMQ.

### Adding a Data Source

You can add one or more data sources (a collection of sensors, gateways, or other input devices providing data) to associate with an edge.

Each defined data source consists of:

* Data source type (sensor, input device like a camera, or gateway) - the origin of the data
* Communication protocol typically associated with the data source
* Authentication type to secure access to the data source and data
* One or more fields specifying the data extraction method - the data pipeline specification
* Categories which are attributes that can be metadata you define to associate with the captured data

#### Add a Data Source - MQTT

* Data Extraction - MQTT
* Categories - MQTT

View these topics in the [Xi IoT Infrastructure Admin Guide](https://portal.nutanix.com/page/documents/details/?targetId=Xi-IoT-Infra-Admin-Guide:Xi-IoT-Infra-Admin-Guide), available from the Nutanix Support Portal.

### MQTT Client Samples for Testing

If you are looking to understand the internals of how MQTT works, please read the 10 part series on [MQTT Essentials](https://www.hivemq.com/tags/mqtt-essentials/) by HiveMQ.

#### Javascript

Please refer to [mqtt package](https://www.npmjs.com/package/mqtt) and examples [here](https://github.com/mqttjs/MQTT.js/blob/master/examples/client/secure-client.js) for creating secure mqtt clients in javascript.

#### Python 2

Prerequisites

* A Nutanix edge with an IP address onboarded to Xi IoT
* X509 certificates generated using Xi IoT
* Python 2.7.10
* pip 10.0.1 (python 2.7)
* paho-mqtt. Install it for python 2.7.10 using the following command:

```console
$ sudo pip2.7 install paho-mqtt
```

Sample

Below is a simple example that shows how to connect to an mqtt broker, publish a single message to a specific topic and receive the published message back.

```python
# TODO
```

Running the example

1. Download the certificates from Xi IoT and store them locally under certs. directory. Name the files as follows:
    * ca.crt - Root CA certificate
    * client.crt - client certificate
    * client.key - client private key
1. Modify broker_address to point to the Xi IoT edge IP address.
1. Run the example as follows:
```console
$ python2.7 mqtt-example.py
```
Expect output:
```console
Connecting...
Connected to broker
Published!
New message received!
Topic: test
Message: Hello, World!
```

## Runtime Environments

A runtime environment is a command execution environment to run applications written in a particular language or associated with a specific Docker registry or file. Each Function added to a Data Pipeline is executed via its own specified Runtime Environment.

Xi IoT includes standard runtime environments including but not limited to the following. These runtimes are read-only and cannot be edited, updated, or deleted by users. They are available to all projects, functions, and associated container registries.
