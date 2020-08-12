.. title:: Nutanix Karbon Services for IoT Deepomatic Demo Application Guide

.. toctree::
  :maxdepth: 2
  :caption:     Contents
  :hidden:

  index

------------------------------------------
Karbon Services for IoT - Deepomatic Demo Application Guide
------------------------------------------

Explore a gesture recognition demo application, as seen at .NEXT 2019. Some other gestures have been included, can you find them all?

Karbon Services for IoT Overview
###############


The Nutanix Karbon Services for IoT platform delivers local compute and AI for IoT edge devices, converging the edge and cloud into one seamless data processing platform.
The Karbon Services for IoT platform eliminates complexity, accelerates deployments, and elevates developers to focus on the business logic powering IoT applications and services.
Now developers can use a low-code development platform to create application software via APIs instead of arduous programming methods.

Karbon Services for IoT Trial
############

This Karbon Services for IoT demo leverages the Karbon Services for IoT Trial. The trial is a limited-time, ready-to-deploy
implementation of the Karbon Services for IoT edge computing platform. The Karbon Services for IoT Trial provides pre-built applications and data
connectors hosted on its own infrastructure. This instant architecture demonstrates how 
to quickly develop and test IoT applications in the cloud for seamless deployment to the edge.

Nutanix has already created the basic infrastructure you need to get
started.

**What's In the Karbon Services for IoT Trial?**

1. Karbon Services for IoT management console, which provides the base for your Karbon Services for IoT
   trial.

2. A Starter project that includes:

   -  You (the project user).

   -  Xi Edge stack, connected and ready to go: no cluster or bare-metal
      resources required on your part.

   -  Deepomatic demo application, just waiting for your gestures.

   -  Karbon Services for IoT Sensor smartphone app for video input.

..

**What else can I do with the Karbon Services for IoT Trial?**

-  Stream video from YouTube-8M video or your smartphone to the Xi Cloud
   edge.

-  Automatically run containerized apps at the edge to perform object
   recognition on your video.

-  Stream the results back to the Karbon Services for IoT console or your smartphone,
   with recognized objects highlighted in your video.

..

Signing Up For the Karbon Services for IoT Trial
+++++++++++++++++++++++++++++++

Do any of these steps to sign up for the Karbon Services for IoT Trial.

a. Click Start Trial at https://www.nutanix.com/products/iot/ or
   `https://iot.nutanix.com <https://iot.nutanix.com/>`__.

b. Sign up now for a My Nutanix account at
   `https://my.nutanix.com <https://my.nutanix.com/>`__.

c. If you already have an account, log on to
   `https://my.nutanix.com <https://my.nutanix.com/>`__ with your
   existing account credentials and click Learn More in the Karbon Services for IoT
   panel.

..

SUPPORT FOR AND LEARNING MORE ABOUT Karbon Services for IoT

The most support for the Karbon Services for IoT trial is available through the
Nutanix Next Karbon Services for IoT trial forum. Nutanix asks that you share your
experiences and lessons learned with your fellow users.

You can also visit the following pages for more information about Xi
IoT.

-  Connect with other users at `Karbon Services for IoT User
   Forum. <https://next.nutanix.com/xi-iot-72>`__

-  Connect on `Twitter <https://twitter.com/NutanixIoT>`__ with the
   Nutanix Karbon Services for IoT team.

-  Check out articles about Karbon Services for IoT at the
   `Nutanix Developer site <https://developer.nutanix.com/iot>`__.

-  View videos about Karbon Services for IoT at `Nutanix University YouTube
   channel <https://www.youtube.com/watch?v#wmUkz-XZLJo>`__.

-  Get more details about Karbon Services for IoT features in the `Nutanix
   documentation <https://portal.nutanix.com/?filterKey#type&filterVal#Xi#/page/docs/list>`__.


Getting Started With the Karbon Services for IoT Trial
+++++++++++++++++++++++++++++++++++++

1. Log on to the Karbon Services for IoT management console.

2. On your smartphone, download the Karbon Services for IoT Sensor app (available from the
   Google Play Store).

Logging On to the Karbon Services for IoT Console
--------------------------------

Before you begin:

Supported web browsers include the current and two previous versions of
Google Chrome. You'll need your My Nutanix credentials for this step.
  
1. Open https://iot.nutanix.com/ in a web browser, click **Log in with
   My Nutanix** and log on with your My Nutanix credentials.

2. If you are logging on for the first time, click to read the Terms and Conditions, then click to Accept and Continue.

3. Take a few moments to read about Karbon Services for IoT, then click Get Started.

Your web browser displays the Karbon Services for IoT dashboard and the Karbon Services for IoT Quick Start Menu.

Karbon Services for IoT Quick Start Menu
-----------------------

The Karbon Services for IoT management console includes a Quick Start menu next to
your user name. You can click Quick Start, then click the links to:

1. See object detection in action by using a YouTube-8M video.

2. Try object detection on your phone.

3. Invite your colleagues to try out Karbon Services for IoT.

4. Edit a data pipeline.

5. Create an application.

**This tutorial utilizes the Quick Start menu, but does not follow the same steps. Please continue to steps below.**

Using the Deepomatic Demo and Karbon Services for IoT Sensor Apps to Detect Gestures in Your Smartphone Video
############################################################################################

About this task

Connect your Android based phone (iPhone coming soon) as a data source to stream video and perform
gesture detection in near realtime using Karbon Services for IoT. Output can be viewed on your phone
and from an HTTP Live Stream (HLS) in your browser.
   
#. If you are not logged on, open https://iot.nutanix.com/ in a web
   browser and log in.

#. Connect your phone through the Quick Start menu.

   a. Click Quick Start, then click **Scan QR code** to connect a phone.

   .. figure:: contents/images/image5.png

      Figure 1: Quick Start: Phone as Data Source
      
   b. Open the Google Play Store on your Android based phone, search for Karbon Services for IoT Sensor, 
      and install the app on your phone.

   c. After downloading and installing the Karbon Services for IoT Sensor app, scan the QR code
      to authenticate.

   d. Enter a name for your phone, then click Next.
   
#. From the Karbon Services for IoT management portal, Click :fa:`bars` **> Apps and Data > Data Pipelines**.

   The phone-object-detection data pipeline is connected to the phone data source and started by default. Since this data pipeline is not used as part of this tutorial, it can be stopped.

#. On the **phone-object-detection** data pipeline tile, click **Actions**, then **Stop**, then **Stop** again.

#. Click :fa:`bars` **> Apps and Data > Applications**.

#. On the **deepomatic-demo** application tile, click **Actions**, then **Start**, then **Start** again.

#. Open the Karbon Services for IoT Sensor app on your phone, and tap the button to switch to the front facing camera.

#. Tap **Capture Video**, and wait up to 30 seconds for the deepomatic-demo application to initialize.

#. Switch to the **deepomatic-demo** tab to view the results. Try making the gestures seen on stage at .NEXT 2019, or try your own to see if they're recognized!

#. From the Karbon Services for IoT management portal, click **View Http Live Stream** on the **deepomatic-demo** application tile. This opens HLS output for viewing the results in your browser.

#. Click :fa:`remove` to close the HLS page.

#. From the Karbon Services for IoT Sensor app, press :fa:`stop` to stop capturing video.

The containerized deepomatic-demo application utilizes built-in Karbon Services for IoT input and output connectors. Learn more about using your phone or a YouTube-8M video as a 
data source, and a HTTP Live Stream as output when writing your own applications for Karbon Services for IoT by further exploring the deepomatic-demo application.

#. On the **deepomatic-demo** application tile, click **Actions**, then **Edit**.

   The General Information page displays information about the application such as its Name, Description, the Project its assigned to, and 
   the edges on which its assigned to run.

#. Click **Next**.

   The Yaml Configuration page lists the application pod's specification YAML in Kubernetes format.
   
#. Click **Next**.

   The Input and Output page provides the option to use a YouTube-8M video or Karbon Services for IoT Sensor phone app as input and a 
   HTTP Live Stream (HLS) as an output for applications. Simply check the appropriate boxes, and install a `NATS <https://nats.io>`__
   client within your application. The selected input will be available on the NATS topic name stored in the NATS_SRC_TOPIC environment 
   variable. Subscribe to it using the NATS server name stored in the NATS_ENDPOINT environment variable. Application output 
   in jpeg format sent to the topic name stored in NATS_DST_TOPIC will be available via the application's HTTP Live Stream.

#. Click :fa:`remove` to close the application without making any changes.

#. From the Karbon Services for IoT management portal, click **Actions**, then **Stop**, then **Stop** again on the **deepomatic-demo** application tile.

Ready to try more sample apps?
##############################

For more information and guided tutorials on using the other sample applications and data pipelines provided as part of the Karbon Services for IoT trial, check
out the `Karbon Services for IoT - Quick Start for AI Inference <https://nutanix.handsonworkshops.com/workshops/2470db21-2e9a-47a3-bc25-be8eb7521a68/p>`__.



