In this part we will get basic Wordpress working with MySQL DB. We will use a separate instance of MySQL DB so it can be managed independently from the Wordpress service components
The Wordpress and MySQL are already setup with ecommerce application plugins and product catalog to make the deployment easier.

Please refer [here](../README.md) for pre-requisites before proceeding.

# Installing MySQL 

We will use the Admin Console to create the MySQL App.
* Navigate to **Apps and Data** -> **Kubernetes App** -> **Create** 
* Name it **mysql** and add it to the project **Ecommerce Application**.
* Add the required Service Domains to the app. 
* Use the [mysql-app.yaml](mysql/mysql-app.yaml) as the application manifest 

NOTE: The docker image used for MySQL in the manifest above is pre-loaded with data and configuration to make the demo app easier to deploy.

<!---
<details>
<summary>Customizing yaml</summary>
If you need to add custom configuration you can do that using the helm chart. Create a values files and add the appropriate values for your installation

[mysql-custom-values.yaml](mysql-custom-values.yaml)

```
root:
  password: <root password>
  forcePassword: false
  injectSecretsAsVolume: true

db:
  user: wordpress
  password: <wordpress user password>
  name: wordpress_db
  forcePassword: false
  injectSecretsAsVolume: false

replication:
  enabled: false
```

Dowload the helm chart and convert it to a yaml. [mysq-app.yaml](mysql-app.yaml)
```
helm fetch --untar bitnami/mysql
helm template -n bloganalytics mysql --values mysql-custom-values.yaml > mysql-app.yaml
```
---

Create the applicaton spec file for XKS

[mysql-xks.yaml](mysql-xks.yaml)
```
kind: application
name: mysql
description: mysql db
project: Ecommerce Application
appYamlPath: mysql-app.yaml 
 ```
---
Use the xi-iot CLI to deploy the app
```
xi-iot create -f mysql-xks.yaml
```
</details>
--->

# Installing Wordpress
Similar to MySQL we will create another application called **wordpress** in the Admin console and use the app manifest [wordpress-app.yaml](wordpress/wordpress-app.yaml)

In the manifest we have enabled ingress for this application, with a hostname **woodkraft.ntnxsherlock.com**. We will come back to the usage of Ingress Controller [later](#Adding-Ingress). If you need to change the hostname, please see the section [Changing hostname for the site]()

NOTE: As with MySQL, this Wordpress container is pre-loaded with product catalog.

<!---
<details>
<summary>Customize yaml</summary>
Create custom values for the chart to used the external DB created above. Please make sure that the database and password match the values created for MySQL.
[wordpress-custom-values.yaml](wordpress-custom-values.yaml)
```
externalDatabase:
  host: bloganalytics-mysql
  user: wordpress
  password: <password>
  database: wordpress_db
  port: 3306
mariadb:
  enabled: false
service:
  type: ClusterIP
ingress:
  enabled: true
  hostname: bloganalytics.foo.org
```
---
Download the chart for Wordpress and convert it to yaml [wordpress-app.yaml](wordpress-app.yaml)
```
helm fetch --untar bitnami/wordpress
helm template -n bloganalytics wordpress --values wordpress-custom-values.yaml > wordpress-app.yaml
```
---
Create the applicaton spec file for XKS

[wordpress-xks.yaml](wordpress-xks.yaml)
```
kind: application
name: wordpress
description: wordpress app
project: Ecommerce Application
appYamlPath: wordpress-app.yaml
```
---
Use the xi-iot CLI to deploy the app
```
$ xi-iot create -f wordpress-xks.yaml
Successfully created application: wordpress
```
</details>
-->

# Adding Ingress
Now that we have MySQL and Wordpress install, we want to expose the application outside the cluster using Ingress. 

Since we have already enable Ingress resource when deploying Wordpress, all we have to do is enable the [XKS managed Ingress Controller](https://github.com/nutanix/xi-iot/tree/master/services/ingress) to make use of the [Ingress Resource](https://kubernetes.io/docs/concepts/services-networking/ingress/#the-ingress-resource)

XKS support 2 types of ingress controllers, Nginx and Traefik. We will be using Nginx for this project. For details on ingress controller support in XKS please refer to [ingress documentation](https://github.com/nutanix/xi-iot/tree/master/services/ingress)

## Enabling Nginx Ingress Controller
We can enable it using the CLI. Here **svcdomain** is the name of the Service Domain where the application is running
```
xi-iot update svcdomain <svcdomain> --ingress-type=NGINX
```
---
Confirm the setting
```
$ xi-iot get svcdomain svcdomain1 -o yaml
kind: edge
name: svcdomain1
connected: true
profile:
  privileged: false
  enableSSH: false
  ingressType: NGINX
effectiveProfile:
  privileged: false
  enableSSH: false
  ingressType: NGINX
 ```

## Accessing the application
You can setup your own DNS and maps the hostname used in the Ingress **woodkraft.ntnxsherlock.com** to any node IP of the Service Domain.

For simplicity we can add the hostname to the **/etc/hosts** file in your system to map to any node IP of the service domain, e.g.
```
echo "10.45.27.51 woodkraft.ntnxsherlock.com" | sudo tee -a /etc/hosts
```

Now you should be able to go to http://woodkraft.ntnxsherlock.com and access the site.

You can also login using the credentials **admin/wordpress** to further customize the application. Please note that the settings will be removed if you restart the app, since the data is stored in the docker container for simmplicity.
