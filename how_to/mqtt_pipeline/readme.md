# XI IOT - QUICK START FOR MQTT DATA PIPELINES

## Xi IoT Overview

The Nutanix Xi IoT platform delivers local compute and AI for IoT edge devices, converging the edge and cloud into one seamless data processing platform. The Xi IoT platform eliminates complexity, accelerates deployments, and elevates developers to focus on the business logic powering IoT applications and services. Now developers can use a low-code development platform to create application software via APIs instead of arduous programming methods.

SUPPORT FOR AND LEARNING MORE ABOUT XI IOT

The most support for the Xi IoT trial is available through the Nutanix Next Xi IoT trial forum. Nutanix asks that you share your experiences and lessons learned with your fellow users.

You can also visit the following pages for more information about Xi IoT.

* Connect with other users at [Xi IoT User Forum](https://next.nutanix.com/xi-iot-72).
* Connect on [Twitter](https://twitter.com/NutanixIoT) with the Nutanix Xi IoT team.
* Check out articles about Xi IoT at the [Nutanix Developer site](https://developer.nutanix.com/iot).
* View videos about Xi IoT at [Nutanix University YouTube channel](https://www.youtube.com/watch?v#wmUkz-XZLJo).
* Get more details about Xi IoT features in the [Nutanix documentation](https://portal.nutanix.com/?filterKey#type&filterVal#Xi#/page/docs/list).

### Logging On to the Xi IoT Console

Before you begin:

Supported web browsers include the current and two previous versions of Google Chrome. You’ll need your My Nutanix credentials for this step.

1. Open [https://iot.nutanix.com/](https://iot.nutanix.com/) in a web browser, click **Log in with My Nutanix** and log on with your My Nutanix credentials.
1. If you are logging on for the first time, click to read the Terms and Conditions, then click to Accept and Continue.
1. Take a few moments to read about Xi IoT, then click Get Started.

Your web browser displays the Xi IoT dashboard and the Xi IoT Quick Start Menu.

#### Creating a Project

In Xi IoT, Projects are used to segment resources such as applications and edges so that only assigned users can view and modify them. This allows different departments or teams to utilize shared data sources, edges, or cloud resources without interfering with each other.

As part of this tutorial, you’ll create a new Project to deploy your sample Data Pipelines and Applications.
1. From the **Xi IoT** management portal, select **More > Projects > + Create**.
1. Fill out the following fields and click **Next**:
    * **Name** - MQTT Pipeline
    * **Description** - Optional
    * Select + **Add Users**
    * Select your user name and click **Done**
1. Click + **Add Infrastructure**, select your Edge, and click **Done**.
Xi IoT has the ability to natively output Data Pipelines from the edge to several public cloud services such as AWS S3, or GCP Cloud Datastore. For this tutorial, Cloud Profile Selection can be left blank because no cloud resources will be used.

Xi IoT can also natively run Applications (Docker containers) at the edge using Kubernetes formated yaml as the only required input. Each yaml definition refers to a container image stored in a public or private registry. Private registries can be accessed by creating a Xi IoT Container Registry Profile to store required access information. Because this tutorial utilizes containers hosted in a public registry, Container Registry Selection can be left blank.
1. Click **Create**

#### Stagin Source Data

The tutorial depends on the availability of MQTT sample data.

Xi IoT supports direct ingest of [MQTT](https://www.hivemq.com/blog/how-to-get-started-with-mqtt/) messaging protocol (commonly used by IoT sensor devices). For other industry specific protocols, numerous hardware & software “gateways” exist to translate those data formats & protocols into MQTT.

Outside of a tutorial environment, this data would likely originate on a device external to the Edge device. However, for the purposes of the tutorial, we can leverage Xi IoT’s **Application** construct to deploy a [pre-configured containerized application](https://cloud.docker.com/u/xiiot/repository/docker/xiiot/mqtt-sensor) running directly on your Edge to generate MQTT data.

As mentioned above, Xi IoT Applications are simply Docker containers that can be deployed to the edge using Kubernetes formated yaml as the only required input. This is considered Containers-as-a-Service (CaaS) functionality and is sold as a specific Xi IoT service SKU.

### Using a Xi IoT App to Generate Sample MQTT Data

“MQTT Sensor” is a single container application that executes a python script to download a CSV file input and publish contents of each row as an MQTT message to the local Xi Edge where it is running. It can be used to test data pipeline transforms and outputs.

#### Creating the MQTT Sensor Data Source

1. If you are not logged on, open [https://iot.nutanix.com/](https://iot.nutanix.com/) in a web browser and log in.
1. Click **More > Infratructure > Data Sources.**
1. Click **+ Add Data Source.**
1. Enter **mqtt-sensor** in the Name field.
1. In the Associated Infrastructure dropdown, choose the appropriate Edge (this is the same edge where the MQTT Sensor app will run).
1. In the Protocol dropdown, choose **MQTT.**
1. Click **+ Generate Certificates**, then click **+ Download** to save the zip file to a location for reference in a future step.
1. Click **Next**.
1. On the Data Extraction page, click **Add New Field**, enter **temp** in the Name field, enter **temp** into the MQTT Topic field, click the round, blue **check**, then click **Next**.
1. Click inside the first (left) **Attribute dropdown** of the newly added category assignment and choose **Data Type**.
1. Click inside the second (right) **Attribute dropdown** of the newly added category assignment and choose **Temperature**.
1. Click **Add**.
1. On your local machine, open a shell and base64 encode the certificate bundle (zip) downloaded above. Keep the raw output available for a future step.

```console
base64 -i 1561481707433_certificates.zip
```