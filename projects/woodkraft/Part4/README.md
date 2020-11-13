In this section will cover the Istio and make use of some of the features provided by Istio like Blue/Green testing.

# Enable Istio
Select the Project **E-Commerce Application** on the [Projects](https://karbon.nutanix.com/projects/list) page and go to **Manage Services**. Enable Istio and Confirm the selection

**NOTE**: This will restart all apps in the project to enable Istio sidecar. If you prefer not to add Istio to a particular application, please add the following annotation to your Deployment metadata. Here are addition [details](https://istio.io/latest/docs/setup/additional-setup/sidecar-injection/)

```
sidecar.istio.io/inject: "false"
```

Once Istio is enabled on the Service Domain for the project we can deploy V2 of the recommendation-service and experiment with the routing rules.

# Installing Recommendation-Service V2
* Go to **Ecommerce Application** project
* Create an application via the Admin Console
  **Kubernetes Apps** -> **Create**
* Name it **recommendation-service-v2**
* Add the required Service Domains to the app
* Use the [recommendation-service-v2.yaml](istio/recommendation-service-v2.yaml) as the application manifest

This will install the v2 version of the app, which essentially generate the same graph for top five products but in color blue. It will also add a single [VirtualService](https://istio.io/latest/docs/reference/config/networking/virtual-service/) which will route users to the version v1 of the service by default and add a special route for Firefox users to route to version v2. So with this VirtualService if you access the e-commmerce portal from Firefox you should seed the graph on Recommendations tab in green, otherwise it should be in blue.

You can verify the VirtualService in the Project console here:
  **Istio** -> **Virtual Services**

Here is the VirtualService used. You can tweak it to excercise different behaviour of Istio Routing

```
apiVersion: networking.istio.io/v1alpha3
kind: VirtualService
metadata:
  name: recomm-svc
spec:
  hosts:
  - recommendation-service
  http:
  - match:
    - headers:
        user-agent:
          regex: .*Firefox.*
    route:
    - destination:
        host: recommendation-service
        subset: v2
  - route:
    - destination:
        host: recommendation-service
        subset: v1
```