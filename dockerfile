FROM golang:latest
RUN apt-get update
RUN apt-get -y install npm
RUN npm install -g nodemon
RUN apt -y install chromium

COPY ./ /app/
WORKDIR /app