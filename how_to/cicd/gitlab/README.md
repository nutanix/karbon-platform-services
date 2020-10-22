# Gitlab with Karbon Platform Services

This guide will show you how to implement a CI/CD pipeline in Gitlab with KPS.

## Setup

You will start by setting up your environment.

1. Install [Docker](https://www.docker.com/) on your local machine.
2. Create an API Token in your KPS account
* You will find a section labeled **Manage API Keys** under your username in the top right corner of the KPS UI
2. Set up a Python and Docker Runner in your Gitlab project.
* Follow this [guide](https://angristan.xyz/2018/09/build-push-docker-images-gitlab-ci/) on how to set up your Runners
3. Create environment variables in your Gitlab project at **Settings** -> **CI/CD** -> **Variables**.
* Create a variable called **REGISTRY_USER** which will refer to your container registry username
* Create a variable called **REGISTRY_PASS** which will refer to your container registry password
* Create a variable called **TOKEN** which will refer to your KPS API Token