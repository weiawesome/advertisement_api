#!/bin/bash

kubectl apply -f secret.yaml

helm install helm-mysql -f value.yaml oci://registry-1.docker.io/bitnamicharts/mysql
