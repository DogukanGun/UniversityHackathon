FROM ubuntu:latest

RUN apt-get update
RUN apt-get install ca-certificates -y
RUN update-ca-certificates

WORKDIR /
ADD main ./main
ADD server.env ./server.env

ENTRYPOINT ./main