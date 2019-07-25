FROM python:2.7
RUN pip install nats-client==0.8.2
RUN pip install protobuf==3.6.1
RUN wget https://s3-us-west-2.amazonaws.com/ntnxsherlock-runtimes/python2-env.tgz
RUN tar -xvzf python2-env.tgz
RUN mv ./python-env/datastream_pb2.py ./
ADD main.py ./
CMD ["python", "main.py"]

