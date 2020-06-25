## Accessing Xi IoT 

Open [https://my.nutanix.com/](https://my.nutanix.com/) in your browser. If you don’t already have a My Nutanix account, follow steps to create one.
Scroll to the Xi Cloud Services section and click Launch to access the Xi IoT SaaS control plane.
At this point you should have a dashboard with a default User (you), Project, Category.
You will be able to find Data Sources under the Infrastructure tab.

## Xi IoT CLI

Open [https://github.com/nutanix/xi-iot/tree/master/cli](https://github.com/nutanix/xi-iot/tree/master/cli) in your browser and follow the steps to install the Xi IoT CLI. 

## Data Sources

In Xi IoT, data sources can be set up in order to connect cloud entities such as Applications and Data Pipelines to edge devices. Xi IoT ia equipped with a few ready to use protocols
such as MQTT, RTSP, and GigE Vision.  

We also support many other data sources as well which can be deployed using the Xi IoT CLI amd a yaml file to specifying criteria such as service domain and topics. You can find examples of these 
yaml files inside each of the data source subfolders as well as how to implement and use them in the platform. 