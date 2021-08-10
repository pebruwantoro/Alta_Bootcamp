#!/bin/bash

echo "Run and create container (image should be exist)"
docker run -d --name trialContainer -p 3000:8080 trial-docker:latest

# docker run -d --name <nama-container> \
# -p <host-port>:<container-port> \
# -p <host-port>:<container-port> \
# -e <ENVNAME=VALUE>
# <image-name>