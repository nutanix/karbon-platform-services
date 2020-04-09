
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

* **Golang**
* **NodeJS**
	* Node.js functions can be run in context of a data pipeline. A transformation function must accept context as well as message payload as parameters. Context can be used to query function parameters passed in when function has been instantiated. Moreover context is used to send messages to next stage in data pipeline.
	* Following is a basic Node.js function template:
	```javascript
	function main(ctx, msg) {
	   return new Promise(function(resolve, reject) {
	      // log list of transformation parameters
	      console.log("Config", ctx.config)
	      // log length of message payload
	      console.log(msg.length)
	      // forward message to next stage in pipeline
	      ctx.send(msg)
	      // complete promise
	      resolve()
	   })
	}
	exports.main = main
	```

	All functions must export main which returns a promise.

	Expected output:
	```console
	Config { IntParam: '42', StringParam: 'hello' }
	2764855
	```

	**Note**
	Packages available in NodeJS Runtime
		* alpine-baselayout
		* alpine-keys
		* apk-tools
		* busybox
		* libc-utils
		* libgcc
		* libressl2.5-libcrypto
		* libressl2.5-libssl
		* libressl2.5-libtls
		* libstdc++
		* musl
		* musl-utils
		* scanelf
		* ssl_client
		* zlib
* **Python**
	* Functions can be executed in data pipelines to transform and filter data. Transformations are functions used to process single messages and optionally forward them to next stage in data pipeline. The next stage could be another transformation or destination of the data pipeline on edge or in the cloud. Transformation can accept parameters. In Python parameters are passed as dictionary to transformation. The following script demonstrates some basic concepts:
	```python
	import logging
	# Python function are invoked with context and message payload.
	# The context can be used to retrieve metadata about the message and allows
	# function to send mesagges to next stage in stream. In this sample we just
	# log message payload and forward it as is to next stage.
	def main(ctx, msg):
	      logging.info("Parameters: %s", ctx.get_config())
	      logging.info("Process %d bytes from %s at %s", msg, ctx.get_topic(), ctx.get_timestamp())
	      # Forward to next stage in pipeline.
	      ctx.send(msg)
	```
	Pass two parameters to the function:

		* MyStringParam like the name suggests is a parameter of type string.
		* MyIntParam is a number.
		The function would produce the following console output when processing images from a camera:
		```console
		[2019-03-12 04:57:26,820 root INFO] Parameters: {u'MyIntParam':u'42', u'MyStringParam': u'hello'}
		[2019-03-12 04:57:26,820 root INFO] Process 2764855 bytes from rtsp://184.72.239.149:554/vod/mp4:BigBuckBunny_175k.mov at 1552366646754939017
		```
	**Methods provided by ctx**

	* **ctx.get_config()** - returns a dict of parameters passed to the function.
	* **ctx.get_topic()** - returns the topic (string) on which the current message was received. In this case, it is the topic is set to RTSP topic from which image has been received.
	* **ctx.get_timestamp()** - returns the time in nanoseconds since epoch (Jan 1st, 1970 UTC).
	* **ctx.send()** - Takes bytes as input and forwards it to the next stage in the pipeline. If the input is not of type bytes, an error is thrown and a corresponding alert is raised in Xi IoT.

	**In memory caching**
	```python
	import logging
	counter=0
	def main(ctx, msg):
	      global counter
	      logging.info("This is message number %d", counter)
	      counter+=1
	      # Forward to next stage in pipeline.
	      ctx.send(msg)
	```






