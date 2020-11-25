# Facial Recognition

This is a Kubernetes application designed to detect the faces of characters from the TV show Friends.

## Building the Application

Perform the following steps containerize the application and push it to a Container Registry.

1. Run the following command to build the docker container:
```
docker build -t <registry>/<repo>:<image tag> .
```
2. Run the following command to push the image to a Container Registry:
```
docker push <registry>/<repo>:<image tag>
```

## Deploying the Application

Fill in the `<image path>` in [app.yaml](app.yaml) and upload it to Karbon Platform Services as a Kubernetes App. Finally add a data source such as youtube-8m which is streaming a video of Friends to Karbon Platform Services.

Ex:
* https://www.youtube.com/watch?v=jKvcO6IZy6I