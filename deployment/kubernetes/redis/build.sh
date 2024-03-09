#!/bin/bash

kubectl apply -f secret.yaml

helm install helm-redis -f value.yaml oci://registry-1.docker.io/bitnamicharts/redis-cluster
