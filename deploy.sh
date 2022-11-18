#!/usr/bin/env bash

set -e
YD_HOME=$HOME/ydaemon
TAG=${1:-latest}
ENV_FILE=${2:-"$HOME/.yd.dev.env"}
BRANCH=${3:-develop}

echo "[*] Starting deployment"
echo "[*] Sourcing environment"
if ! [ -f "$ENV_FILE" ]; then
        echo "[!] $ENV_FILE file not found. Exiting..."
        exit 1
fi

rm -fr $YD_HOME || true
git clone https://github.com/yearn/ydaemon $YD_HOME
cd $YD_HOME
git checkout $BRANCH

source $ENV_FILE
docker pull "ghcr.io/yearn/ydaemon:$TAG"
docker-compose -f docker-compose.dev.yml stop
docker-compose -f docker-compose.dev.yml up -d

echo "[*] Finished!"
