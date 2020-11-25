# FaceNet

This is a Kubernetes application designed to detect faces in a video feed which can be deployed on Karbon Platform Services.

## Building the Application

Perform the following steps containerize the application and push it to a Container Registry.

1. Download the required `.pb` file as shown in the Dockerfile
2. Run the following command to build the docker container:
```
docker build -t <registry>/<repo>:<image tag> .
```
3. Run the following command to push the image to a Container Registry:
```
docker push <registry>/<repo>:<image tag>
```

## Deploying the Application

Fill in the `<image path>` in [facenet-deployment.yaml](facenet-deployment.yaml) and upload it to Karbon Platform Services as a Kubernetes App. Finally add the data source with the video of faces as an input.