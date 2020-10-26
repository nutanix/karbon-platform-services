# Jenkins with Karbon Platform Services

This guide will show you how to implement a CI/CD pipeline in Gitlab with KPS.
![Pipeline!](img/pipeline.png "CI/CD Pipeline")

## Setup

You will start by setting up your environment. This guides assumes you already have a Jenkins environment set up.

1. Install [Docker](https://www.docker.com/) on your local machine.
2. Create an API Token in your KPS account
    * You will find a section labeled **Manage API Keys** under your username in the top right corner of the KPS UI
3. Set up a Python and Docker Runner in your Gitlab project.
    * Follow this [guide](https://angristan.xyz/2018/09/build-push-docker-images-gitlab-ci/) on how to set up your Runners
4. Create environment variables in your Gitlab project at **Settings** -> **CI/CD** -> **Variables**.
    * Create a variable called **REGISTRY_USER** which will refer to your container registry username
    * Create a variable called **REGISTRY_PASS** which will refer to your container registry password
    * Create a variable called **TOKEN** which will refer to your KPS API Token

**Note**: This example will be using Dockerhub as a container registry, but this framework can be used with any container registry.

## CI/CD Pipeline

This section will take a deep dive into the different steps being used in the CI/CD pipeline. These steps have been configured in [Jenkinsfile](Jenkinsfile).

### Build

The following is a job configured to build the cat-dog image.



## Using the Pipeline

When you merge a change into the branches you have tagged in your CI yaml, the pipeline will trigger the jobs 
that have been specified. You can view the results and logs of these jobs at **CI/CD** -> **Pipelines**. You can view the logs of the job to debug
failures in the pipeline.
![Jobs!](img/jobs.png "Pipeline Jobs")

## Takeaways

* How to set up Runners in a Gitlab project
* How to configure a Jenkinsfile
* Configuring environment variables in a Jenkins Pipeline
* How to configure a Jenkins CI job
* How to trigger a Jenkins CI/CD pipeline
* Using the KPS API to run a Jenkins deployment job