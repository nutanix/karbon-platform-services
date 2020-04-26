# XI IOT - QUICK START FOR AI INFERENCE

## Xi IoT Overview
The Nutanix Xi IoT platform delivers local compute and AI for IoT edge devices, converging the edge and cloud into one seamless data processing platform. The Xi IoT platform eliminates complexity, accelerates deployments, and elevates developers to focus on the business logic powering IoT applications and services. Now developers can use a low-code development platform to create application software via APIs instead of arduous programming methods.

## Xi IoT Trial
This Xi IoT Quick Start leverages the Xi IoT Trial. The trial is a limited-time, ready-to-deploy implementation of the Xi IoT edge computing platform. The Xi IoT Trial provides pre-built applications and data connectors hosted on its own infrastructure. This instant architecture demonstrates how to quickly develop and test IoT applications in the cloud for seamless deployment to the edge.

Nutanix has already created the basic infrastructure you need to get started.

**What’s In the Xi IoT Trial?**
1. Xi IoT management console, which provides the base for your Xi IoT trial.
1. A Starter project that includes:
    * You (the project user).
    * Xi Edge stack, connected and ready to go: no cluster or bare-metal resources required on your part.
    * YouTube-8M application, just waiting for your YouTube-8M video URL.
    * Xi IoT Sensor smartphone app, if you want to use your own video instead of YouTube-8M.

**What Can I Do with the Xi IoT Trial?**
* Stream video from YouTube-8M video or your smartphone to the Xi Cloud edge.
* Automatically run containerized apps at the edge to perform object recognition on your video.
* Stream the results back to the Xi IoT console or your smartphone, with recognized objects highlighted in your video.

### Signing Up For the Xi IoT Trial¶
Do any of these steps to sign up for the Xi IoT Trial.
1. Click Start Trial at [https://www.nutanix.com/products/iot/](https://www.nutanix.com/products/iot/) or [https://iot.nutanix.com](https://iot.nutanix.com).
1. Sign up now for a My Nutanix account at [https://my.nutanix.com](https://my.nutanix.com).
1. If you already have an account, log on to [https://my.nutanix.com](https://my.nutanix.com) with your existing account credentials and click Learn More in the Xi IoT panel.

SUPPORT FOR AND LEARNING MORE ABOUT XI IOT

The most support for the Xi IoT trial is available through the Nutanix Next Xi IoT trial forum. Nutanix asks that you share your experiences and lessons learned with your fellow users.

You can also visit the following pages for more information about Xi IoT.

* Connect with other users at [Xi IoT User Forum](https://next.nutanix.com/xi-iot-72).
* Connect on [Twitter](https://twitter.com/NutanixIoT) with the Nutanix Xi IoT team.
* Check out articles about Xi IoT at the [Nutanix Developer site](https://developer.nutanix.com/iot).
* View videos about Xi IoT at [Nutanix University YouTube channel](https://www.youtube.com/watch?v#wmUkz-XZLJo).
* Get more details about Xi IoT features in the [Nutanix documentation](https://portal.nutanix.com/?filterKey#type&filterVal#Xi#/page/docs/list).

**Getting Started With the Xi IoT Trial**
1. Log on to the Xi IoT management console.
1. Have a YouTube-8M URL handy or create and upload a video from your smartphone.
1. On your smartphone, download the Xi IoT Sensor app (available from the Google Play Store).

**Logging On to the Xi IoT Console**
Before you begin:

Supported web browsers include the current and two previous versions of Google Chrome. You’ll need your My Nutanix credentials for this step.
* Open [https://iot.nutanix.com/](https://iot.nutanix.com/) in a web browser, click **Log in with My Nutanix** and log on with your My Nutanix credentials.
* If you are logging on for the first time, click to read the Terms and Conditions, then click to Accept and Continue.
* Take a few moments to read about Xi IoT, then click Get Started.

**Xi IoT Quick Start Menu**
The Xi IoT management console includes a Quick Start menu next to your user name. You can click Quick Start, then click the links to:
1. See object detection in action by using a YouTube-8M video.
1. Try object detection on your phone.
1. Invite your colleagues to try out Xi IoT
1. Edit a data pipeline
1. Create an application

**This tutorial utilizes the Quick Start menu, but does not follow the same steps. Please continue to steps below.**

## Using the Xi IoT Sensor App to Detect Objects in Your Smartphone Video
About this task

Connect your Android based phone (iPhone coming soon) as a data source to stream video and perform object detection in near realtime using Xi IoT. Output can be viewed on your phone and from an HTTP Live Stream (HLS) in your browser.
1. If you are not logged on, open [https://iot.nutanix.com/](https://iot.nutanix.com/) in a web browser and log in.
1. Connect your phone through the Quick Start menu.
    * Click Quick Start, then click **Scan QR code** to connect a phone.
    * Open the Google Play Store on your Android based phone, search for Xi IoT Sensor, and install the app on your phone.
    * After downloading and installing the Xi IoT Sensor app, scan the QR code to authenticate.
    * Enter a name for your phone, then click Next.
1. From the Xi IoT management portal, Click **More > Apps and Data > Data Pipelines**.
    * The phone-object-detection data pipeline tile should show **Status: Healthy**.
1. If it shows Status: Stopped, click **Actions**, then **Start**.
1. Open the Xi IoT Sensor app on your phone, click **Capture Video**, and wait for up to 30 seconds for the inference engine to start.
1. Switch to the **phone-object-detection** tab to view the results. Point your phone’s camera around the room to identify objects in near realtime!
1. From the Xi IoT management portal, click **View Http Live Stream** on the phone-object-detection data pipeline tile. This opens HLS output for viewing the results in your browser.
1. Click **x** to close the HLS page.
1. Stop capturing video
1. From the Xi IoT management portal, click **Actions**, then **Stop**, then Stop again on the phone-object-detection data pipeline tile.


## Using the Xi IoT App Library to Detect Objects in a YouTube-8M Video
About this task

Use a YouTube-8M video to demonstrate object recognition in Xi IoT. We recommend a short video showing city scenes, drone footage, or a sporting event.
1. If you are not logged on, open [https://iot.nutanix.com/](https://iot.nutanix.com/) in a web browser and log in.
1. From the Xi IoT management portal, Click **More > Apps and Data > Kubernetes Apps**.
1. On the **youtube-8m-object-detection-app** application tile, click **Actions**, then **Start**, then **Start** again.
1. Click View App UI.
1. Copy and paste the following YouTube-8M URL in the field, then press play.
    * [https://www.youtube.com/watch?v=HqqsJkonXs](https://www.youtube.com/watch?v=HqqsJkonXs)  
1. Close the App UI tab.
1. Back on the **youtube-8m-object-detection-app** application tile, click **Actions**, then click **Stop**, then click **Stop** again.

## Using Xi IoT Data Pipelines to Detect Objects in a YouTube-8M Video
About this task

Use data pipelines and two YouTube-8M videos to demonstrate object detection using only python code in Xi IoT.
1. If you are not logged on, open [https://iot.nutanix.com/](https://iot.nutanix.com/) in a web browser and log in.
1. From the Xi IoT management portal, Click **More > Apps and Data > Data Pipelines**.
1. On the youtube-8m-object-detection data pipeline tile, click **Actions**, then **Start**, then **Start** again.
1. Now click **View Http Live Stream** to view object detection via HLS output.
1. After viewing the output, click **x** to close the HLS page.
1. On the same youtube-8m-object-detection data pipeline tile, click **Actions**, then **Edit**.
1. Click **x** to close the data pipeline without making any changes.

In Xi IoT, categories help you assign various attributes to edges and data sources which can be further used to query and select them when creating Data Pipelines or deploying Applications.

An example of a category could be “City” with values in [San Francisco, San Jose, San Diego] or “State” with values in [California, Washington, Oregon] and so on. It can be anything meaningful to your environment.

In the next steps, you’ll add a new category for assignment to a new YouTube-8M channel, add the channel to the YouTube-8M data source, and modify the data pipeline to use this new channel.
1. Click **More > Administration > Categories**.
1. Click the check box beside the youtube-8m Category, then click **Edit**.
1. Click **Add Value**, enter **channel2**, click the round, blue check, then click **Update**.
1. Click **More > Infrastructure > Data Sources**.
1. Click the check box beside the youtube-8m data source, then click **Edit**.
1. Click **Add New URL**, enter **youtube-8m-2** in the Name field, copy and paste **https://www.youtube.com/watch?v=PYbrTRE1bZg** into the URL field, click the round, blue **check**, then click **Next**.
1. On the Category Assignment page, click inside the **Select Fields dropdown** and choose **Select Fields…**.
1. From the Select Fields dialog, select click the check box beside the youtube-8m-1 field to select it and click **OK**.
1. Click **Add** to add a category assignment for the **youtube-8m-2** field just created.
1. Click inside the **Select Fields dropdown** of the newly added category assignment and choose **Select Fields**.
1. From the Select Fields dialog, click the check box beside the youtube-8m-2 field to select it and click **OK**.
1. Click inside the first (left) **Attribute dropdown** of the newly added category assignment and choose **youtube-8m**.
1. Click inside the second (right) **Attribute dropdown** of the newly added category assignment and choose **channel2**.
1. Click **Update** to update the data source with the new channel.
1. Click **More > Apps and Data > Data Pipelines**.
1. On the youtube-8m-object-detection data pipeline tile, click **Actions**, then **Edit**.
1. In the Input section, click inside the second (right) **Select by Categories dropdown** and change **channel1** to **channel2**.
1. Click Update to update the data pipeline with the new input data source category.
    * The data pipeline will automatically update to use the YouTube-8M stream created and assigned the channel2 category as the new data source.
1. To verify the new stream is being used, click **View Http Live Stream** on the youtube-8m-object-detection data pipeline tile to view object detection via HLS output.
1. Click **x** to close the HLS page.
1. On the youtube-8m-object-detection data pipeline tile, click **Actions**, then **Stop**, then **Stop** again to stop the data pipeline.

## Using Xi IoT Input and Output Connectors for Applications
About this task

Learn more about using your phone or a YouTube-8M video as a data source, and a HTTP Live Stream as output when writing your own applications for Xi IoT by exploring the echoapp sample application provided in the Application Library.

1. If you are not logged on, open [https://iot.nutanix.com/](https://iot.nutanix.com/) in a web browser and log in.
1. Click **More > Apps and Data > Applications**.
1. On the echoapp application tile, click **Actions**, then **Edit**.
    * The General Information page displays information about the application such as its Name, Description, the Project its assigned to, and the edges on which its assigned to run.
1. Click **Next**
    * The Yaml Configuration page lists the application pod’s specification YAML in Kubernetes format.
```
kind: Deployment
apiVersion: apps/v1
metadata:
  name: sample-python-nats-app
  labels:
    app: sample-python-nats-app
spec:
  replicas: 1
  selector:
    matchLabels:
      app: sample-python-nats-app
  template:
    metadata:
      labels:
        app: sample-python-nats-app
    spec:
      containers:
      - name: app-container
        image: 770301640873.dkr.ecr.us-west-2.amazonaws.com/edgecomputing/dataifc/python-nats-app:pankit.thapar
        imagePullPolicy: Always
        command: ["/usr/local/bin/python", "demo-app.py"]
```
1. Click **Next**
    * The Input and Output page provides the option to use a YouTube-8M video or Xi IoT Sensor phone app as input and a HTTP Live Stream (HLS) as an output for applications. Simply check the appropriate boxes, and install a NATS client within your application. The selected input will be available on the NATS topic name stored in the NATS_SRC_TOPIC environment variable. Subscribe to it using the NATS server name stored in the NATS_ENDPOINT environment variable. Application output in jpeg format sent to the topic name stored in NATS_DST_TOPIC will be available via the application’s HTTP Live Stream.
1. Use one of the YouTube-8M sample (or your own) videos, or the Xi IoT Sensor phone app to demonstrate. Choose **phone** or **youtube-8m** in the **Type of Input** dropdown, and the channel as appropriate in the **Field** dropdown.
1. Click **Update**.
1. Click **Actions**, then **Start**, then **Start** again on the echoapp application tile to start the application.
1. Click **View Http Live Stream** on the echoapp application tile to view the application’s HLS output.
1. Click **x** to close the HLS page.
1. Click **Actions**, then **Stop**, then **Stop** again on the echoapp application tile to stop the application.

## Using the Xi IoT AI Inferencing Service to Detect Objects
About this task

Use data pipelines and a YouTube-8M video to demonstrate object detection using the Xi IoT AI Inferencing Service.

1. If you are not logged on, open [https://iot.nutanix.com/](https://iot.nutanix.com/) in a web browser and log in.
1. Click **More > Apps and Data > Data Pipelines**.
1. On the ai-inference-service-demo data pipeline tile, click **Actions**, then **Edit**.
    * This data pipeline will look very similar to the youtube-8m-object-detection pipeline used in the earlier exercise. However, there’s one major difference. Take notice of the transformation function used in this pipeline. It’s named **ml_objectdetect_func-python**.
    * To better understand how the AI Inferencing Service works, first take a look at the **objdetect_func-python** used in the youtube-8m-object-detection data pipeline in the earlier exercise, then compare it to the **ml_objectdetect_func-python** function.
1. Click **x** to close the data pipeline without making any changes.
1. Click **More > Apps and Data > Functions**.
1. Click the check box beside the **objdetect_func-python** function, then click **Edit**.
    * Notice that the function is written in python and uses the Xi IoT Tensorflow Python runtime.
1. Click **Next**
    * The function’s python code is now displayed. Take notice of lines 16-19 excerpted below:
```python
BASE_PATH = "/mllib/objectdetection"

# ssd_inception_v2_coco   latency - 42ms
PATH_TO_CKPT = BASE_PATH + '/ssd_inception_v2_coco_2017_11_17/frozen_inference_graph.pb'
```
As mentioned in the earlier exercise, the ssd inception v2 model is embedded in the Xi IoT Tensorflow Python runtime. This is fine for an example, but not suitable for production deployments. For example, the model cannot be updated.

Now compare this python code to that used in the ml_objectdetect_func-python function used in the ai-inference-service-demo data pipeline.

8. Click **x** to close the function without making any changes.
9. Uncheck the check box beside the **objdetect_func-python** function, click the check box beside the **ml_objdetect_func-python** function, then click **Edit**.
    * Notice that this function is also written in python and uses the Xi IoT Tensorflow Python runtime.
10. Click **Next**.
    * The function’s python code is now displayed. This time, first take notice of line 23:
```python
ai_inference_endpoint = os.environ['AI_INFERENCE_ENDPOINT']
```
Object detection, or inference, will now be performed by the Inferencing Service, so the function must know the service endpoint for submission at the edge. As you can see, its automatically passed to the runtime as an environment variable.

Now take notice of lines excerpted below:
```python
def detect(image):
   image_np = np.asarray(image, dtype="int32")
   image_np_expanded = np.expand_dims(image_np, axis=0)
   data = json.dumps({"signature_name": "serving_default",
                     "instances": image_np_expanded.tolist()})
   headers = {"content-type": "application/json"}
   model_name = "objectdetect"
   model_version = 1
   url = "http://%s/v1/models/%s/versions/%d:infer" % (
      ai_inference_endpoint, model_name, model_version)
   response = requests.post(url, data=data, headers=headers)
   if response.status_code != 200:
      logging.error(response.json())
      return None
   text = response.text
   inference_payload = json.loads(text)
   predictions = inference_payload['predictions']
   return predictions[0]
```
This excerpt is of the detect function utilizing the Inferencing Service. Of particular note are lines 159-161 where the model_name is set to objectdetect, model_version is set to 1, and the connection url is built with the ai_inference_endpoint (remember this was passed automatically).For this example, the Inferencing Service has already been pre-loaded with the same ssd inception v2 model, but this time trained using the [Open Images dataset](https://github.com/openimages/dataset).

11. Click **x** to close the function without making any changes.
12. View how this model and others can be managed by clicking **More > Apps and Data > ML Models**.
13. Click the check box beside the objectdetect model, then click **Edit**.
    * Example model version 1 is listed along with the Tensorflow Framework Type. A new version of the model could be uploaded by clicking **Add new**. This version could then be referenced in functions similarly to line 160 in the example code above.
14. Click **x** to close the ML model without making any changes.

If you’d like to view the data pipeline and Inferencing Service in action, simply navigate back to the **ai-inference-service-demo** data pipeline tile, start the pipeline, then click to view the Http Live Stream.

## Takeaways

What are the key takeaways and other things you should know about **Nutanix Xi IoT**?
* Get started with AI Inference in minutes with Xi Cloud based edges
* Use Xi IoT Sensor app as instant video data source
* Use HTTP Live Stream output to instantly view application and data pipeline results
* A single platform that can run AI-based apps, containers, and functions.
* Easy to deploy applications at scale with a SaaS control plane.
* Reduced time to setup and configure edge intelligence (i.e. kubernetes and analytics platform).
* Operate edge locations offline with limited internet connectivity.
* Can choose cloud connectivity without heavy lifting via native public cloud APIs.
* Supports development languages like Python, Node.js and Go and integrates into existing CI/CD pipelines.
* Developer APIs and pluggable architecture enables “bring your own framework and functions” for simplified integrations without having to rewrite your code.