#!/bin/sh
ID="~Everybody+(Backstreet's+Back)+(Radio+Edit)"
AUDIO=`base64 -i "$ID".wav`
RESOURCE=localhost:3002/cooltown
echo "{ \"Audio\":\"$AUDIO\" }" > input
curl -v -X POST -d @input $RESOURCE
