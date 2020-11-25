FROM jjanzic/docker-python3-opencv:opencv-4.0.1

RUN mkdir -p /code /model

# (!!!) Download manually from https://drive.google.com/file/d/1EXPBSXwTaqrSC0OhUdXNmKSh9qJUQ55-/view
ADD ./20180402-114759/20180402-114759.pb /model/20180402-114759.pb
ADD ./src/requirements.txt /code/requirements.txt

RUN pip install -r /code/requirements.txt
ADD ./src /code

CMD ["/usr/local/bin/python3", "/code/nats_loop.py", "/model/20180402-114759.pb"]

