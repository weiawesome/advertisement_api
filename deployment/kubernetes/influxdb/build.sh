#!/bin/bash

kubectl apply -f secret.yaml

helm install helm-influxdb -f value.yaml oci://registry-1.docker.io/bitnamicharts/influxdb
