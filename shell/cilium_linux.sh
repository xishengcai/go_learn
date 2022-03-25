#!/bin/bash

#Requirements:
#
#Kubernetes must be configured to use CNI (see Network Plugin Requirements)
#Linux kernel >= 4.9.17
#内核升级指南：https://www.cnblogs.com/ding2016/p/10429640.html

set -e

kernel=`uname -r | cut -d "-" -f 1`

kernel_v1=`echo $kernel | cut -d "." -f 1`
kernel_v2=`echo $kernel | cut -d "." -f 2`
kernel_v3=`echo $kernel | cut -d "." -f 3`

echo `uname -r`
if [[ $kernel_v1 -lt 4 ]];then
  echo "$kernel 版本低于 4.9.17"
  exit 1
fi

if [[ $kernel_v2 -lt 9 ]];then
  echo "$kernel 版本低于 4.9.17"
  exit 1
fi

if [[ $kernel_v2 -eq 9 ]] &&  [[ $kernel_v3 -lt 17 ]];then
  echo "$kernel 版本低于 4.9.17"
  exit 1
fi

#CILIUM_VERSION=$(curl -s https://raw.githubusercontent.com/cilium/hubble/master/stable.txt)
#HUBBLE_VERSION=$(curl -s https://raw.githubusercontent.com/cilium/hubble/master/stable.txt)
#curl -LO https://github.com/cilium/cilium-cli/archive/refs/tags/v0.8.6.tar.gz

CILIUM_VERSION=1.10.3

downloadCilium(){
  curl -LO https://soft-package-xisheng.oss-cn-hangzhou.aliyuncs.com/k8s/cilium/v0.8.4/cilium-linux-amd64.tar.gz
  sudo tar xzvfC cilium-linux-amd64.tar.gz /usr/local/bin
  rm -rf cilium-linux-amd64.tar.gz
}

installCilium(){
#  cilium install
#  yum install tc -y
  helm repo add cilium https://helm.cilium.io/
  helm install cilium cilium/cilium --version $CILIUM_VERSION \
  --namespace kube-system

}

enableHubble() {
  # Enabling Hubble requires the TCP port 4245 to be open on all nodes running Cilium. T
  # his is required for Relay to operate correctly.
#  cilium hubble enable
  helm upgrade cilium cilium/cilium --version $CILIUM_VERSION \
   --namespace kube-system \
   --reuse-values \
   --set hubble.relay.enabled=true \
   --set hubble.ui.enabled=true

}

downloadHubbleClient(){
  curl -LO "https://github.com/cilium/hubble/releases/download/$HUBBLE_VERSION/hubble-linux-amd64.tar.gz"
  curl -LO "https://github.com/cilium/hubble/releases/download/$HUBBLE_VERSION/hubble-linux-amd64.tar.gz.sha256sum"
  sha256sum --check hubble-linux-amd64.tar.gz.sha256sum
  tar zxf hubble-linux-amd64.tar.gz
}

main(){
  installCilium
  enableHubble
}

main