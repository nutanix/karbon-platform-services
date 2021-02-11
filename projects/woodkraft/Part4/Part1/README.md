In this part we will get basic Wordpress working with MySQL DB. We will use a separate instance of MySQL DB so it can be managed independently from the Wordpress service components
The Wordpress and MySQL are already setup with ecommerce application plugins and product catalog to make the deployment easier.

Please refer [here](../README.md) for pre-requisites before proceeding.

# Installing MySQL 

We will use the Admin Console to create the MySQL App.
* Select **E-Commerce Application** Project from the drop down
* Navigate to **Kubernetes Apps** -> **Create**
* Name the app **mysql** and add it to the project **E-Commerce Application**
* Add the Service Domains created earlier to the app
* Use the [mysql-app.yaml](mysql/mysql-app.yaml) as the application manifest 

**NOTE**: The docker image used for MySQL in the manifest above is pre-loaded with data and configuration to make the demo app easier to deploy.

# Installing Wordpress
Similar to MySQL we will create another application called **wordpress** in the console and use the app manifest [wordpress-app.yaml](wordpress/wordpress-app.yaml)

In the manifest we have enabled ingress for this application, with a hostname **woodkraft.ntnxdomain.com**. In the next section we will enable Ingress Controller

NOTE: As with MySQL, this Wordpress container is pre-loaded with product catalog.

# Enabling Ingress Controller
Now that we have MySQL and Wordpress installed, we want to expose the application outside the cluster using K8s Ingress. 

Since we have already enabled Ingress resource when deploying Wordpress, all we have to do is enable the KPS managed Ingress Controller](https://karbon.nutanix.com/services) to make use of the [Ingress Resource](https://kubernetes.io/docs/concepts/services-networking/ingress/#the-ingress-resource)

KPS support 2 types of ingress controllers, Nginx and Traefik. We will be using Nginx for this project. For details on ingress controller support in KPS please refer to [ingress documentation](https://portal.nutanix.com/page/documents/details?targetId=Karbon-Platform-Services-Admin-Guide:ks-using-an-ingress-controller-c.html)

## Adding custom certificates
By default TLS ingress is initialized with a fake certificate. If you would like to change the following fields in the [wordpress-app-tls.yaml](wordpress-app-tls.yaml) file
``` 
tls.cert: 
tls.key:
```
Both values are base64 encoded of the certification and key file in PEM format

## Accessing the application
You can setup your own DNS and maps the hostname used in the ingress rules **woodkraft.ntnxdomain.com** to any node IP of the Service Domain. You can find the node IP by navigating to the Project and **Nginx-Ingress** -> **Rules**

For simplicity we can add the hostname to the **/etc/hosts** file in your system to map to any node IP of the service domain, e.g.
```
echo "10.45.27.51 woodkraft.ntnxdomain.com" | sudo tee -a /etc/hosts
```

Now you should be able to go to https://woodkraft.ntnxdomain.com and access the site.

You can also login using the credentials **admin/wordpress** to further customize the application. Please note that the settings will be removed if you restart the app, since the data is stored in the docker container for simmplicity.
