#!/bin/sh

./build.sh
podman build -t mygoapi .

podman run -p 5000:5000 mygoapi  --name mygoapi -d 