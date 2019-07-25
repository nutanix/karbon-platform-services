FROM node:8
ADD main.js /
ADD package*.json /
RUN npm install
RUN wget https://s3-us-west-2.amazonaws.com/ntnxsherlock-runtimes/node-env.tgz
RUN tar -xvzf node-env.tgz
CMD ["node", "main.js"]