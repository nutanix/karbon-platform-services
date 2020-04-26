# XI IOT - FACEFEED APPLICATION DEPLOYMENT GUIDE¶

## Xi IoT Overview
The Nutanix Xi IoT platform delivers local compute and AI for IoT edge devices, converging the edge and cloud into one seamless data processing platform. The Xi IoT platform eliminates complexity, accelerates deployments, and elevates developers to focus on the business logic powering IoT applications and services. Now developers can use a low-code development platform to create application software via APIs instead of arduous programming methods.

In this tutorial you’ll deploy an application called “Facefeed” using the Xi IoT SaaS control plane. This application ingests a video stream using the real time streaming protocol (RTSP), and uses machine learning to detect known faces. The application and its data pipelines are deployed to a Xi Edge device for local execution. This tutorial assumes your edge has already been deployed using steps from the [Xi IoT Infrastructure Admin Guide](https://portal.nutanix.com/#/page/docs/details?targetId=Xi-IoT-Infra-Admin-Guide:Xi-IoT-Infra-Admin-Guide). Access to the Admin Guide is provided via the My Nutanix Portal and requires an account to login.

## Accessing Xi IoT
1. Open [https://my.nutanix.com/](https://my.nutanix.com/) in your browser. If you don’t already have a My Nutanix account, follow steps to create one.
1. Scroll to the Xi Cloud Services section and click Launch to access the Xi IoT SaaS control plane.
    * At this point you should have a dashboard with a default User (you), Project, Category.

### Download App Files
The Facefeed application utilizes five Functions within two Data Pipelines to transform the incoming data (RTSP video stream)and draw inference (Face IDs) to identify faces. For convenience, these pre-made functions are available on GitHub:
* aggregatefeed.py - Maintains a buffer of known and unknown faces for output.
* face_register.py - Combines UI data inputs for storage in the database.
* facematch.py - Matches inference results to database of registered faces.
* facerecognition.py - Uses a TensorFlow machine learning model to draw inference.
* raw_to_jpeg.py - Converts the raw binary stream into a readable image format.
* es_datamover.go - Streams output into a local Elasticsearch instance for later recall.

1. Open [https://github.com/nutanix/xi-iot](https://github.com/nutanix/xi-iot) in a new browser tab and click **Clone or download > Download ZIP**.
Extract the .zip file to a directory. These are required when referencing .yaml, .py, and .go files later in this tutorial.
1. Extract the .zip file to a directory. These are required when referencing .yaml, .py, and .go files later in this tutorial.

## Defining Categories
In Xi IoT, categories help you assign various attributes to edges and data sources which can be further used to query and select them when creating Data Pipelines or deploying Applications.

An example of a category could be “City” with values in [San Francisco, San Jose, San Diego] or “State” with values in [California, Washington, Oregon] and so on. It can be anything meaningful to your environment. For this tutorial, we’ll categorize types of cameras by their function.
1. From the **Xi IoT** management portal, select **More > Infrastructure > Categories**.
1. Click on the default **Data Type** to explore the default values.
    * Data Type is one way to categorize and sort forms of data processed by an Edge.
1. Return to **Categories** and click **+ Create** to add your own, custom category with two values.
1. Fill out the following fields and click **Create**:
    * **Name** - Camera Type
    * **Purpose** - Identifies the intended use of the camera
    * Select + **Add Value**
    * **Value** - Facial Recognition
    * Select + **Add Value**
    * **Value** - Face Registration

## Creating a Project
In Xi IoT, Projects are used to segment resources such as applications and edges so that only assigned users can view and modify them. This allows different departments or teams to utilize shared data sources, edges, or cloud resources without interfering with each other.

As part of this tutorial, you’ll create a new Project to deploy your sample Data Pipelines and Applications.
1. From the **Xi IoT** management portal, select **More > Projects > + Create**.
1. Fill out the following fields and click Next:
    * **Name** - Facefeed
    * **Description** - Optional
    * Select + **Add Users**
    * Select your user name and click **Done**
1. Click + **Add Edges** and select your Edge.
    * Xi IoT has the ability to natively output Data Pipelines from the edge to several public cloud services such as AWS S3, or GCP Cloud Datastore. For this tutorial, Cloud Profile Selection can be left blank because no cloud resources will be used.
    * Xi IoT can also natively run Applications (Docker containers) at the edge using Kubernetes formated yaml as the only required input. Each yaml definition refers to a container image stored in a public or private registry. Private registries can be accessed by creating a Xi IoT Container Registry Profile to store required access information. Because this tutorial utilizes containers hosted in a public registry, Container Registry Selection can be left blank.
1. Click **Create**

## Staging Source Data
The tutorial depends on the availability of a video stream from which to identify faces.

Xi IoT supports direct ingest of RTSP (commonly used in retail/security) and GigE Vision (commonly used in manufacturing/industrial) video streaming protocols, as well as MQTT messaging protocol (commonly used by IoT sensor devices). For other industry specific protocols, numerous hardware & software “gateways” exist to translate those data formats & protocols into MQTT.

Outside of a tutorial environment, these video streams would likely originate on a camera or network video recorder external to the Edge device. However, for the purposes of the tutorial, we can leverage Xi IoT’s **Application** construct to deploy a pre-configured containerized application hosting an RTSP video stream running directly on your Edge VM.

As mentioned above, Xi IoT Applications are simply Docker containers that can be deployed to the edge using Kubernetes formated yaml as the only required input. This is considered Containers-as-a-Service (CaaS) functionality and is sold as a specific Xi IoT service SKU.

### Deploying RTSP Sample Feed Application
1. From the **Xi IoT** management portal, select **More > Projects > Facefeed > Apps & Data > Applications > + Create Application**.
1. Fill out the following fields and click **Next**:
* **Name** - facefeed-rtsp-samples
* **Description** - Optional
* Select + **Add Edges**
* Select your **Edge**
1. Click **Choose File** and select xi-iot-master\projects\facefeed\applications\facefeed-rtsp-sample.yaml.
    * Note the environment variables and values defined in the YAML file, namely RTSP_USERNAME and RTSP_PASSWORD.
1. Click **Create**
1. Click **facefeed-rtsp-sample** to see a Summary of the application performance, alerts, deployments, etc.
    * Edge Deployments should list “1 of 1 Running” on your Edge device once the application has successfully launched.

**Note**
    * Deployment of the application may take a few minutes as the ~200MB container needs to be downloaded from the Internet to the Edge VM.

### Adding RTSP Sample Feed as a Data Source
1. From the Xi IoT management portal, select **More > Infrastructure > Edges**.
1. Record your Edge IP address. You’ll need this in the next step.
1. From the Xi IoT management portal, select **More > Infrastructure > Data Sources > + Add Data Source**.
1. Fill out the following fields and click **Next**:
    * Type - Sensor
    * Name - rtsp-sample-feed
    * Associated Edge - your Edge
    * Protocol - RTSP
    * Authentication Type - Username and Password
    * IP Address - your Edge IP address recorded earlier
    * Username - Found in facefeed-rtsp-sample.yaml
    * Password - Found in facefeed-rtsp-sample.yaml
Next you will define what data is extracted from the source, in this case, we require the specific address used to host the stream.
1. Click **Add New Field** and fill out the following fields:
    * **Name** - VideoFeed
    * **RTSP URL** - live.sdp
1. Click the check to add the data extraction field.
1. Click **Next**.
    * Finally you will assign the category attributes which will be used to identify the sample feed as the data source for the facial recognition Data Pipeline you will build in later exercises.
1. From the Attribute drop down menu, select **Camera Type : Facial Recognition**.
1. Click **Add**

## Deploy Functions
Xi IoT Functions allow developers to directly build and execute business logic to correlate, filter, or transform data in standard languages such as Python or Go without the burden of maintaining underlying operating systems or runtimes.
1. From the **Xi IoT** management portal, select **More > Projects > Facefeed > Apps & Data > Functions > + Add Function**.
1. Fill out the following fields to create the first function:
    * **Name** - aggregatefeed
    * **Description** - Optional
    * **Project** - Facefeed
    * **Language** - Python
    * **Runtime Environment** - Tensorflow Python
Xi IoT Functions may be written in well known software languages most commonly used for edge computing and machine learning. These currently include Python, Go, and Node.js. This allows developers to re-use existing code, or quickly write new logic utilizing standard libraries, and without the burden of learning a new platform or language.
3. Click **Next**.
4. Click **Choose File** and select xi-iot-master\projects\facefeed\functions\aggregatefeed.py.
5. Click **Create**.
6. Repeat Steps 1-5 to add the remaining four .py python functions. The **Name** should follow the script name (without .py).
7. Repeat Steps 1-4 once more to add the es_datamover.go function. Note that the Language and Runtime Environment should be golang, and Golang Env.
    * Before clicking Create, click + **Add parameter** in the left pane. Enter esIndex in the **Name** field. Select string form the Type dropdown. Click the check to save the parameter.
8. Click Create.

## Deploying Data Pipelines
Data Pipelines in Xi IoT allow you to transform data by injecting your own code. In this exercise, we will use Data Pipelines to transform frames (from the video feed) into Face IDs (by using machine learning).

### Data Pipeline 1 - faceregister¶
This Data Pipeline will source the frames from a local webcam or uploaded image (using a containerized UI application you’ll deploy), apply a TensorFlow machine learning model to detect faces, calculate a unique Face ID, and persist the data in the local Elasticsearch instance running on your edge.

1. From the Xi IoT management portal, select **More > Projects > Facefeed > Apps & Data > Data Pipelines > + Create Data Pipeline**.
2. Select the Facefeed project and click **Next**.
3. Fill out the following fields to build the pipeline:

**Note**
Use the exact **Pipeline** and **Endpoint** Names used in this guide, as the Endpoint name is used as the name of the Elasticsearch index on the edge. The application that you will deploy to leverage these pipelines is hardcoded to look for these specific index names within the local Elasticsearch instance.
    * **Data Pipeline Name** - faceregister
    * Select + **Add Data Source > Data Source**
    * **Category** - Camera Type
    * **Value** - Face Registration
    * Select + **Add Function > facerecognition**
    * Select + to add an additional function
    * Select **face_register**
    * Select + to add an additional function
    * Select **es_datamover**
    * Type datastream-face_register in the esIndex (string) field.
    * Select + **Add Destination > Infrastructure**
    * **Endpoint Type** - Realtime Data Stream
    * **Endpoint Name** - datastream-faceregister
4. Click **Create**

### Data Pipeline 2 - facerecognitionlivefeed
This Data Pipeline will source from the RTSP sample feed you configured as a Data Source, apply a TensorFlow Machine Learning model to detect faces, calculate Face IDs, and search your Elasticsearch instance to find a match. The containerized UI application you’ll deploy will show the known vs unknown faces based on inference results.
1. Click + **Create** to define your next Data Pipeline.
2. Select the **Facefeed** project and click **Next**.
3. Fill out the following fields to build the pipeline:
    * **Data Pipeline Name** - facerecognitionlivefeed
    * Select + **Add Data Source > Data Source**
    * **Category** - Camera Type
    * **Value** - Face Recognition
    * Select + **Add Function > raw_to_jpeg**
    * Select **Enable Sampling Interval** and keep the default 1s interval
    * Select + to add an additional function
    * Select **facerecognition**
    * Select + to add an additional function
    * Select **facematch**
    * Select + to add an additional function
    * Select **aggregatefeed**
    * Select + to add an additional function
    * Select **es_datamover**
    * Type datastream-face_register in the esIndex (string) field.
    * Select + **Add Destination > Infrastructure**
    * **Endpoint Type** - Realtime Data Stream
    * **Endpoint Name** - datastream-facerecognitionlivefeed
4. Click **Create**
At this point, your Data Sources, Functions, and Data Pipelines are all configured and automatically deployed by Xi IoT onto your edge based on your earlier Edge assignment within the Facefeed Project.

In this tutorial you’re outputting Data Pipeline results to an Elasticsearch instance hosted on your edge, but Xi IoT has native capability to output in many ways. From the Destination dropdown you’ll notice the ability to output to your edge, or to a cloud.

Here’s a breakdown of options and typical use cases:
    * **Infrastructure**
        * **Kafka** - real-time streaming between edge local applications
        * **MQTT** - real-time streaming devices (actuators or other edge devices)
        * **Realtime Data Stream** - real-time streaming between Xi IoT Data Pipelines
    * **Cloud**
        * **AWS**
            * **Kinesis** - real-time streaming for large volumes of data
            * **SQS** - sending messages via web service applications
            * **S3** - simple file storage
        * **GCP**
            * **PubSub** - real-time streaming
            * **Cloud Datastore** - simple file storage

## Deploying Facefeed
So far you have deployed a data source, functions for processing that data, and pipelines to tie the functions together and direct output back to our Edge VM. The final step is to deploy the Facefeed application.

Like the sample RTSP stream, Facefeed is a containerized application described by a YAML file provided in the Git repository. It provides the GUI used to upload images to be analyzed by the **faceregister** pipeline, as well as a log of all recognized and unrecognized faces outputted by the **facerecognitionlivefeed** pipeline.
1. From the Xi IoT management portal, select **More > Projects > Facefeed > Apps & Data > Applications > + Create**.
2. Fill out the following fields and click **Next**:
    * **Name** - facefeedui
    * **Description** - Optional
    * Select + **Add Edges**
    * Select your Initials-edge Edge
3. Click **Choose File** and select xi-iot-master\projects\facefeed\applications\facefeed.yaml.
    * Note the host port that will be used to access the application.
4. Click **Create**.
5. Click facefeed and monitor the deployment status until it reaches **1 of 1 Running**.
6. Open https://EDGE-VM-IP:8888/ in a new browser tab and log into Facefeed using the default credentials:
    * **Username** - demo
    * **Password** - facefeed
7. Download the following linked images and add the users to the Registered Faces database:
    * Maurice Moss:
        * Designation - Administrator
        * Department - IT
        * Employee ID - 1738WUH
    * Jen Barber:
        * Designation - Supervisor
        * Department - IT
        * Employee ID - 8675309
    **Note:** If the **Add to Database** button spins and stops without adding an entry to the **List of Registered Faces**, validate that the **Endpoint Name** of the **faceregister** data pipeline is accurate.
8. Once the desired faces have been registered, click **Go to application >** to access the log of known and unknown faces.
9. Return to the Dashboard for the summary view of both projects and infrastructure.

Congratulations! You’ve successfully deployed a facial recognition application to your edge from Xi IoT. This base application could be modified for use in retail, banking, municipalities and more. Xi IoT would then make it simple to manage the deployment and monitoring of both the edge servers as well as the applications and data residing on them.

This tutorial is but one edge application example. Xi IoT has already been deployed by customers to:
* Identify objects on a manufacturing assembly line and control a robot to remove unsanctioned objects automatically.
* Collect multiple parameters from various sensors on a manufacturing assembly line, correlate them, and send aggregated data to the cloud.
* Implement ‘Amazon Go’ for cafeterias. Ingesting camera data at the edge for real-time checkout processing and supply-chain updates.

## Takeaways

What are the key things you should know about **Nutanix Xi IoT**?
* A single platform that can run AI-based apps, containers, and functions as a service.
* Easy to deploy containerized applications at scale with a SaaS control plane.
* Reduced time to setup and configure edge intelligence (i.e. kubernetes and analytics platform).
* Operate edge locations offline with limited internet connectivity.
* Can choose cloud connectivity without heavy lifting via learning APIs.
* Supports serverless and development languages like Python, Node.js and Go and integrates into existing CI/CD pipelines.
* Developer APIs and pluggable architecture enables bring your own framework and functions for simplified integrations without having to rewrite your code.