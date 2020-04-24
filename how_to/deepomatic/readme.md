# XI IOT - DEEPOMATIC DEMO APPLICATION GUIDE

Explore a gesture recognition demo application, as seen at .NEXT 2019. Some other gestures have been included, can you find them all?

## Xi IoT Overview

The Nutanix Xi IoT platform delivers local compute and AI for IoT edge devices, converging the edge and cloud into one seamless data processing platform. The Xi IoT platform eliminates complexity, accelerates deployments, and elevates developers to focus on the business logic powering IoT applications and services. Now developers can use a low-code development platform to create application software via APIs instead of arduous programming methods.

## Xi IoT Trial
This Xi IoT demo leverages the Xi IoT Trial. The trial is a limited-time, ready-to-deploy implementation of the Xi IoT edge computing platform. The Xi IoT Trial provides pre-built applications and data connectors hosted on its own infrastructure. This instant architecture demonstrates how to quickly develop and test IoT applications in the cloud for seamless deployment to the edge.

Nutanix has already created the basic infrastructure you need to get started.

**What’s In the Xi IoT Trial?**

1. Xi IoT management console, which provides the base for your Xi IoT trial.
1. A Starter project that includes:
    * You (the project user).
    * Xi Edge stack, connected and ready to go: no cluster or bare-metal resources required on your part.
    * Deepomatic demo application, just waiting for your gestures.
    * Xi IoT Sensor smartphone app for video input.

**What else can I do with the Xi IoT Trial?**
* Stream video from YouTube-8M video or your smartphone to the Xi Cloud edge.
* Automatically run containerized apps at the edge to perform object recognition on your video.
* Stream the results back to the Xi IoT console or your smartphone, with recognized objects highlighted in your video.

### Signing Up For the Xi IoT Trial

Do any of these steps to sign up for the Xi IoT Trial.

1. Click Start Trial at [https://www.nutanix.com/products/iot/](https://www.nutanix.com/products/iot/) or [https://iot.nutanix.com](https://iot.nutanix.com).
1. Sign up now for a [My Nutanix](https://my.nutanix.com) account.
1. If you already have an account, log on to [My Nutanix](https://my.nutanix.com) with your existing account credentials and click Learn More in the Xi IoT panel.

SUPPORT FOR AND LEARNING MORE ABOUT XI IOT

The most support for the Xi IoT trial is available through the Nutanix Next Xi IoT trial forum. Nutanix asks that you share your experiences and lessons learned with your fellow users.

You can also visit the following pages for more information about Xi IoT.
* Connect with other users at [Xi IoT User Forum](https://next.nutanix.com/xi-iot-72).
* Connect on [Twitter](https://twitter.com/NutanixIoT) with the Nutanix Xi IoT team.
* Check out articles about Xi IoT at the [Nutanix Developer site](https://developer.nutanix.com/iot).
* View videos about Xi IoT at [Nutanix University YouTube channel](https://www.youtube.com/watch?v#wmUkz-XZLJo).
* Get more details about Xi IoT features in the [Nutanix documentation](https://portal.nutanix.com/?filterKey#type&filterVal#Xi#/page/docs/list).

### Getting Started With the Xi IoT Trial
1. Log on to the Xi IoT management console.
1. On your smartphone, download the Xi IoT Sensor app (available from the Google Play Store).

Logging On to the Xi IoT Console¶
Before you begin:

Supported web browsers include the current and two previous versions of Google Chrome. You’ll need your My Nutanix credentials for this step.
1. Open [https://iot.nutanix.com/](https://iot.nutanix.com/) in a web browser, click **Log in with My Nutanix** and log on with your My Nutanix credentials.
1. If you are logging on for the first time, click to read the Terms and Conditions, then click to Accept and Continue.
1. Take a few moments to read about Xi IoT, then click Get Started.

Your web browser displays the Xi IoT dashboard and the Xi IoT Quick Start Menu.

#### Xi IoT Quick Start Menu
The Xi IoT management console includes a Quick Start menu next to your user name. You can click Quick Start, then click the links to:
1. See object detection in action by using a YouTube-8M video.
1. Try object detection on your phone.
1. Invite your colleagues to try out Xi IoT.
1. Edit a data pipeline.
1. Create an application.

## Using the Deepomatic Demo and Xi IoT Sensor Apps to Detect Gestures in Your Smartphone Video

Connect your Android based phone (iPhone coming soon) as a data source to stream video and perform gesture detection in near realtime using Xi IoT. Output can be viewed on your phone and from an HTTP Live Stream (HLS) in your browser.
1. If you are not logged on, open [https://iot.nutanix.com/](https://iot.nutanix.com/) in a web browser and log in.
1. Connect your phone through the Quick Start menu.
2. Click Quick Start, then click **Scan QR code** to connect a phone.
2. Open the Google Play Store on your Android based phone, search for Xi IoT Sensor, and install the app on your phone.
2. After downloading and installing the Xi IoT Sensor app, scan the QR code to authenticate.
2. Enter a name for your phone, then click Next.
1. From the Xi IoT management portal, Click **More > Apps and Data > Data Pipelines**.
1. On the **phone-object-detection** data pipeline tile, click **Actions**, then **Stop**, then **Stop** again.
1. Click **More > Apps and Data > Applications**.
1. On the **deepomatic-demo** application tile, click **Actions**, then **Start**, then **Start** again.
1. Open the Xi IoT Sensor app on your phone, and tap the button to switch to the front facing camera.
1. Tap **Capture Video**, and wait up to 30 seconds for the deepomatic-demo application to initialize.
1. Switch to the **deepomatic-demo** tab to view the results. Try making the gestures seen on stage at .NEXT 2019, or try your own to see if they’re recognized!
1. From the Xi IoT management portal, click View **Http Live Stream** on the **deepomatic-demo** application tile. This opens HLS output for viewing the results in your browser.
1. Click **x** to close the HLS page.
1. From the Xi IoT Sensor app, stop capturing video.

The containerized deepomatic-demo application utilizes built-in Xi IoT input and output connectors. Learn more about using your phone or a YouTube-8M video as a data source, and a HTTP Live Stream as output when writing your own applications for Xi IoT by further exploring the deepomatic-demo application.
1. On the **deepomatic-demo** application tile, click **Actions**, then **Edit**.
    * The General Information page displays information about the application such as its Name, Description, the Project its assigned to, and the edges on which its assigned to run.
1. Click **Next**.
    * The Yaml Configuration page lists the application pod’s specification YAML in Kubernetes format.
1. Click **Next**.
    * The Input and Output page provides the option to use a YouTube-8M video or Xi IoT Sensor phone app as input and a HTTP Live Stream (HLS) as an output for applications. Simply check the appropriate boxes, and install a [NATS](https://nats.io/) client within your application. The selected input will be available on the NATS topic name stored in the NATS_SRC_TOPIC environment variable. Subscribe to it using the NATS server name stored in the NATS_ENDPOINT environment variable. Application output in jpeg format sent to the topic name stored in NATS_DST_TOPIC will be available via the application’s HTTP Live Stream.
1. Click **x** to close the application without making any changes.
1. From the Xi IoT management portal, click **Actions**, then **Stop**, then **Stop** again on the **deepomatic-demo** application tile.
