#!/bin/bash

USERNAME=donpebru
docker build -t $USERNAME/trial-container:latest .
docker push $USERNAME/trial-container:latest

docker run -d --name trialContainer -p 3000:8080 $USERNAME/
trial-docker:latest