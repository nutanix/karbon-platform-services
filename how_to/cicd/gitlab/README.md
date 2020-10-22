# Gitlab with Karbon Platform Services

This guide will show you how to implement a CI/CD pipeline in Gitlab with KPS.
![Pipeline!](img/pipeline.png "CI/CD Pipeline")

## Setup

You will start by setting up your environment.

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

This section will take a deep dive into the different steps being used in the CI/CD pipeline.

### Build

The following is a job configured to build the cat-dog application.
```yaml
docker build:
  image: $DOCKER_RUNNER
  stage: Build image
  tags: 
  - docker
  only:
  - master
  script:
  - docker build -t registry.hub.docker.com/$REGISTRY_USER/$app .
  - docker tag registry.hub.docker.com/$REGISTRY_USER/$app:latest $REGISTRY_USER/$app:$VERSION
```

There are a few labels to note in this job. The first is __tag__ which lets the project runners know who is meant to pick up the job. The __only__ label
is used to signify which branches this job should be triggered on. Finally, __script__ signifies the commands that are to be executed inside the docker container.