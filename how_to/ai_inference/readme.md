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

## Using Xi IoT Data Pipelines to Detect Objects in a YouTube-8M Video¶
About this task

Use data pipelines and two YouTube-8M videos to demonstrate object detection using only python code in Xi IoT.
1. If you are not logged on, open [https://iot.nutanix.com/](https://iot.nutanix.com/) in a web browser and log in.
1. From the Xi IoT management portal, Click **More > Apps and Data > Data Pipelines**.