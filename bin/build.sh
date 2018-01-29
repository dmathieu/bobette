#!/bin/sh
set -e

echo $GCR_TOKEN | docker login \
  --username=oauth2accesstoken \
  --password-stdin \
  gcr.io

imageName=gcr.io/dmathieu-191516/bobette-$ARCH
docker build -t $imageName .
docker push $imageName
