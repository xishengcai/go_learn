#!/bin/bash

kubectl run sockjs --image=xishengcai/sockjs
kubectl expose deployment sockjs --port=8080 --target-port=8080
