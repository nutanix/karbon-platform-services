FROM node:9

RUN wget https://s3-us-west-2.amazonaws.com/ntnxsherlock-runtimes/node-env.tgz
RUN tar xf /node-env.tgz

WORKDIR /node-env
RUN npm install
# Containers should NOT run as root as a good practice
USER 10001
CMD ["/node-env/run.sh"]
