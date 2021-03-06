In this part we will add [Debezium](https://debezium.io/) to the project to captures the events from the MySql DB generated by creating orders on the e-commerce portal. Debezium will then feed the changes to the managed Kafka provided by the KPS platform.

# Enabling Kafka
Since Kafka is a managed service provided by KPS, there is no installing required, we have to simply enable it in the portal.

Select the Project **E-Commerce Application** on the [Projects](https://karbon.nutanix.com/projects/list) page and go to **Manage Services**. Enable Kafka and Confirm the selection

# Install Debezium
Since the Debezium connector for Kafka is available as a container we will install it as an applicatioon in KPS.

* Select **E-Commerce Application** Project from the drop down
* Navigate to **Kubernetes Apps** -> **Create**
* Name the app **debezium** and add it to the project **E-Commerce Application**
* Add the Service Domains created earlier to the app
* Use the [debezium-app.yaml](debezium-app.yaml) as the application manifest 

# Install Recommendation Service 
Recommendation Service is a Python application that consumes the events via Kafka topic created by debezium from the wordpress DB and creates a basic graph of the top five most bought products. 

Similar to the other apps create an Application called **recommendation-service** using the Admin Console

* Use the app yaml [recommendation-service.yaml](recommendation-service/recommendation-service.yaml)

# Verification
Once all the above apps are installed and running, you can verify the Kafka topics created by Debezium under **Kafka** -> **Topics**