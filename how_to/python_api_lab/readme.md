## WELCOME
Welcome to Nutanix API Getting Started Lab (Python) - v1.1.

## WHAT WE ARE DOING
The Nutanix Python API Lab will cover a couple of key points.

Creation of a simple Python Flask web application.
* Creation of a single basic view to display cluster details for the user.
* A backend model to talk to the Nutanix APIs.
* JavaScript to create the interface between the front- and back-end parts of the application.

## WHAT WE AREN’T DOING
This lab is not intended as a guide that can be used to learn Python development. While the copy & paste steps will allow you to create a working application, previous experience with Python will aid you in understanding what each section does.

However, the lab will include links to valuable explanation and learning resources that can be used at any time for more information on each section. For example, the general structure of this application is almost identical to the one provided by the official Python Flask tutorial and will often link to resources there.

## REQUIREMENTS
To successfully complete this lab, you will need an environment that meets the following specifications.
* Previous experience with Python is recommended but not strictly mandatory
* An installation of Python 3.6 or later. For OS-specific information, please see the next section.
* Python pip for Python 3.6.
* Python Flask. On most systems this should be case of running pip3 install flask. On Windows, pip3.exe is in the Scripts folder within the Python install location.
* The text editor of your choice. A good suggestion is [Visual Studio Code](https://code.visualstudio.com/) as it is free and supports Python development via plugin.
* cURL
* cURL (for Windows - see below).

## Python 3.6 on OS X
Install Python 3.6 on OS X
```
brew install python
```

## Python 3.6 on Ubuntu 18.04
From the terminal, the following commands can be used to install Python 3.6:
```
sudo apt-get -y update
sudo apt-get -y install curl
sudo apt-get -y install python3-dev python3-pip
sudo apt-get -y install python3-venv
sudo apt-get -y install python3-setuptools
```

## Python 3.6 on CentOS 7
From the terminal, the following commands can be used to install Python 3.6:
```
sudo yum -y update
sudo yum -y install curl
sudo yum -y install epel-release
sudo yum -y install python36
python3.6 -m ensurepip
sudo yum -y install python36-setuptools
```

## Python 3.6 and cURL on Windows
* Install Python 3.6 by downloading the Python 3.6 [installer](https://www.python.org/downloads/release/python-360/).
* Install [cURL](https://curl.haxx.se/windows/).


### Notes

* If you are running through this lab using Nutanix Frame, Python 3.6 has been installed in the c:python36 directory. cURL has also been installed in the c:tools directory.
* Note that cURL is not required to create the demo app. cURL command samples are provided throughout the lab and may be used for reference at any time. This is due to its cross-platform nature vs supporting platform-specific commands (e.g. PowerShell).
* As at January 2019, a default installation of Python 3.6 will be installed in the following folder:
```
C:\Users\<username>\AppData\Local\Programs\Python\Python36
```

# Optional Components
In addition to the requirement components above, the following things are “nice to have”. They are not mandatory for these labs.
* A Github account. This can be created by signing up directly through [GitHub](https://github.com/).
* The GitHub [Desktop](https://desktop.github.com/) application (available for Windows and Mac only)
* [Postman](https://www.getpostman.com/), one of the most popular API testing tools available.

# Cluster Details
In a presenter-led environment you will be using a shared Nutanix cluster. Please use this cluster when carrying out your cURL and application testing.

In a self-paced environment you will need access to a Nutanix cluster along with the credentials required to access it.