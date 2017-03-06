FROM node:6.10-alpine

RUN mkdir -p /opt/node
ADD . /opt/node

WORKDIR /opt/node
RUN yarn install

CMD node /opt/node/src/index.js

