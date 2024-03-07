#!/bin/bash

kubectl apply -f secret.yaml

helm install helm-grafana -f value.yaml oci://registry-1.docker.io/bitnamicharts/influxdb
