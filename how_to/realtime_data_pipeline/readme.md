
# How to Consume Data from Realtime Data Pipeline

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

#### What is MQTT