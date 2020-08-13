## Accessing Xi IoT 

Open [https://my.nutanix.com/](https://my.nutanix.com/) in your browser. If you don’t already have a My Nutanix account, follow steps to create one.
Scroll to the Xi Cloud Services section and click Launch to access the Xi IoT SaaS control plane.
At this point you should have a dashboard with a default User (you), Project, Category.
You will be able to find Data Sources under the Infrastructure tab.

## Xi IoT CLI

The data interfaces created as part of the following steps require the Xi IoT CLI.  Open [cli](https://github.com/nutanix/xi-iot/tree/master/cli) in your browser and follow the steps to install the Xi IoT CLI. 

## Registering a Data Sources

In Xi IoT, data sources connect cloud entities such as Applications and Data Pipelines to edge devices. Support for configuring MQTT, RTSP, and GigE Vision data sources is built into the UI. 

Other data sources can be deployed using the Xi IoT CLI. The CLI commands accept YAML input to specify criteria such as service domains and topics. The following examples assume you have installed and configured the CLI.
