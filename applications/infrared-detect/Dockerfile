FROM ubuntu:18.04

RUN mkdir -p /app
WORKDIR /app

ENV LANG C.UTF-8
RUN apt-get update && apt-get install -y python3 python3-pip
RUN pip3 --no-cache-dir install --upgrade pip setuptools

RUN groupadd -r dev && useradd -r -g dev dev && mkdir /url-feed && chown dev:dev /url-feed && mkdir /url-feed/uwsgi

# see e.g. https://github.com/janza/docker-python3-opencv/blob/master/Dockerfile
RUN apt-get update && apt-get install -y \
        build-essential \
        cmake \
        git \
        wget \
        unzip \
        yasm \
        pkg-config \
        libswscale-dev \
        libtbb-dev \
        libjpeg-dev \
        libpng-dev \
        libtiff-dev \
        libavformat-dev \
        libpq-dev \
        vim \
        ffmpeg

# OpenCV compilation requires numpy to be installed
RUN pip install numpy

ENV OPENCV_VERSION="3.4.6"
RUN wget -nv https://github.com/opencv/opencv/archive/${OPENCV_VERSION}.zip \
&& unzip -q ${OPENCV_VERSION}.zip \
&& mkdir /app/opencv-${OPENCV_VERSION}/cmake_binary \
&& cd /app/opencv-${OPENCV_VERSION}/cmake_binary \
&& cmake -DBUILD_TIFF=ON \
  -DBUILD_opencv_java=OFF \
  -DWITH_CUDA=OFF \
  -DWITH_OPENGL=ON \
  -DWITH_OPENCL=ON \
  -DWITH_IPP=ON \
  -DWITH_TBB=ON \
  -DWITH_EIGEN=ON \
  -DWITH_V4L=ON \
  -DBUILD_TESTS=OFF \
  -DBUILD_PERF_TESTS=OFF \
  -DCMAKE_BUILD_TYPE=RELEASE \
  -DCMAKE_INSTALL_PREFIX=$(python3 -c "import sys; print(sys.prefix)") \
  -DPYTHON_EXECUTABLE=$(which python3) \
  -DPYTHON_INCLUDE_DIR=$(python3 -c "from distutils.sysconfig import get_python_inc; print(get_python_inc())") \
  -DPYTHON_PACKAGES_PATH=$(python3 -c "from distutils.sysconfig import get_python_lib; print(get_python_lib())") .. \
&& make install \
&& rm /app/${OPENCV_VERSION}.zip \
&& rm -r /app/opencv-${OPENCV_VERSION}

# Python dependencies
RUN pip3 --no-cache-dir install tensorflow==1.12.0 \
  asyncio-nats-client \
  protobuf \
  boto3 \
  configparser \
  psycopg2 \
  Pillow

RUN wget https://yolo-model-configurations.s3.us-east-2.amazonaws.com/yolov3.weights
RUN wget https://yolo-model-configurations.s3.us-east-2.amazonaws.com/yolov3.cfg
RUN wget https://yolo-model-configurations.s3.us-east-2.amazonaws.com/database.ini
RUN wget https://github.com/shadiakiki1986/nats-cli/releases/download/0.0.4.2/nats-amd64 -O /sbin/nats && chmod +x /sbin/nats

COPY *.py /app/
CMD ["python3", "infrared-app.py"]
