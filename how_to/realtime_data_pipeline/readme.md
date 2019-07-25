
**How to Consume Data from Realtime Data Pipeline**

This Realtime Data Pipeline will output the data to [NATS message broker](https://nats.io/).
To consume that data from NATS, we need to connect to the broker and subscribe to a topic.

The topic name is the same as the endpoint name given when you created the data pipeline.
The NATS service is exposed on a DNS name of nats. 

**Payload Format**

Data received from nats is encoded in protobuf format. We must decode the data before using it. The generated protobuf file is available in the custom runtime package. 

 
**Build the Sample Application** 

The sample script uses the topic name "datapipeline-demo", so change the topic name to the your realtime data pipeline endpoint name. 

This example uses docker ([install docker](https://docs.docker.com/install/)) to build an image and push it to public docker repository ([docker hub](https://hub.docker.com/))

To push to docker hub, log on to your account (https://docs.docker.com/engine/reference/commandline/login/)

 ```
 git clone https://github.com/nutanix/xi-iot.git
 cd how_to/realtime_data_pipeline
 cd nodejs (for other languages choose appropriate folder)
 docker build -t realtimedatapipeline:latest .
 docker tag realtimedatapipeline:latest xiiot/sample:latest
 docker push xiiot/sample:latest
 
 ```

**Run the Sample Application**

Create an Application on Xi IoT platform using the kubernetes yaml file (sample.yaml)
