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

About this task:
Connect your Android based phone (iPhone coming soon) as a data source to stream video and perform gesture detection in near realtime using Xi IoT. Output can be viewed on your phone and from an HTTP Live Stream (HLS) in your browser.