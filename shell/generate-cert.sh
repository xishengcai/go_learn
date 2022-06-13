#!/bin/bash

usage() {
cat <<EOF

usage:
    --CN common name: localhost,www.abc.cn，为域名生成证书
    --IP 1.1.1.1, 为IP地址生成证书
    --rsa 2048
    --ns k8s 命名空间,在该命名空间中创建secret: tls-secret
    --dir  证书存放的目录
    --pkcs12  是否生成浏览器证书，example：true, 可以是任意值
    --type   generate certificate is  pem or crt
    -h print help
    example: sh generate-cert.sh --CN localhost
EOF
}

set -e

if [ "$1" = "-h" ]; then
  usage
  exit 0
fi

while [ $# -gt 0 ]
do
    key="$1"
    case $key in
        --IP)
            export IP=$2
            shift
        ;;
        --CN)
            export CN=$2
            shift
        ;;
        --rsa)
            export rsa=$2
            shift
        ;;
        --ns)
            export ns=$2
            shift
        ;;
        --dir)
            export dir=$2
            shift
        ;;
        --pkcs12)
            export pkcs12=$2
            shift
        ;;
        --type)
            export type=$2
            shift
        ;;
        *)
            echo "unknown option [$key]"
            usage
            exit 1
        ;;
    esac
    shift
done

if [ -z "$CN" ]; then
  CN='localhost'
fi

if [ -z "$rsa" ]; then
  rsa=2048
fi

if [ -z "$dir" ]; then
  echo $dir
  dir=cert
fi

genServerCrt(){
  # 生成根秘钥及证书
  openssl req -x509 -newkey rsa:4096 -sha256  -keyout privkey.pem -out cacert.pem -days 3650 -nodes -subj '/CN=Launcher LStack Authority'
  openssl req -x509 -sha256 -newkey rsa:${rsa} -keyout ${dir}/ca.key -out ${dir}/ca.crt -days 3650 -nodes -subj '/CN=Launcher LStack Authority'

  IFS=","
  for i in $CN; do

    # 生成服务器密钥，证书并使用CA证书签名
    openssl genrsa -out ${dir}/${i}.key ${rsa}
    openssl req -new -key ${dir}/${i}.key -subj "/CN=${i}" -out ${dir}/${i}.csr

  # ip domain
    if [ -n "$IP" ]; then
      echo "ip: $IP"
      echo subjectAltName = IP:${IP} > ${dir}/extfile.cnf
      openssl x509 -req -in ${dir}/${i}.csr -CA ${dir}/ca.crt -CAkey ${dir}/ca.key -CAcreateserial -extfile ${dir}/extfile.cnf -out ${dir}/${i}.crt -days 3650
    else
      openssl x509 -req -days 3650 -in ${dir}/${i}.csr\
      -CA ${dir}/ca.crt \
      -CAkey ${dir}/ca.key \
      -CAcreateserial -out ${dir}/${i}.crt
    fi

    openssl req -new -newkey rsa:${rsa}\
     -keyout ${dir}/${i}.key\
     -subj "/CN=${i}" \
#    -config <(cat /System/Library/OpenSSL/openssl.cnf <(printf "[SAN]\nsubjectAltName=DNS:${i}")) \
     -out ${dir}/${i}.csr -nodes
    openssl x509 -req -sha256 -days 3650 -in ${dir}/${i}.csr -CA ${dir}/ca.crt -CAkey ${dir}/ca.key -set_serial 02 -out ${dir}/${i}.crt
  done

}

# 生成p12
creteP12(){
  openssl pkcs12 -export -clcerts -inkey ${dir}/client.key -in ${dir}/client.crt -out ${dir}/client.p12
}

createCert() {
  kubectl delete secret tls-secret -n test

  kubectl -n test \
  create secret generic tls-secret \
--from-file=tls.crt=server.crt \
--from-file=tls.key=server.key \
--from-file=ca.crt=ca.crt
}

createTls() {

  set +e
  # 忽略删除错误
  kubectl delete secret tls-secret -n ${ns}
  set -e

  kubectl -n ${ns} \
  create secret tls tls-secret \
--cert=server.crt \
--key=server.key
}

testCurl() {
  curl -v https://xxx:6443/apis/launchercontroller.k8s.io/v1/rsoverviews/lau-crd-resource-overview?timeout=5s \
--cert /etc/kubernetes/pki/apiserver-kubelet-client.crt  \
--key /etc/kubernetes/pki/apiserver-kubelet-client.key \
--cacert /etc/kubernetes/pki/ca.crt

curl https://xx \
--cert ./cert/client.crt \
--key ./cert/client.key \
--cacert ./cert/ca.crt
}

clean(){
  rm -rf ./extfile.cnf
  rm -rf ./client.*
  rm -rf ./ca.*
  rm -rf ./server.*
}

main() {
  mkdir -p $dir

  if [ "$type" = "pem" ]; then
    generatePem
  else
    genServerCrt
  fi

  if [ "$pkcs12" ]; then
    creteP12
  fi

  if [ "$ns" ]; then
    createTls
  fi
}

generatePem(){
  openssl genrsa -out ${dir}/ca.key 2048
  openssl req -x509 -new -nodes -key ${dir}/ca.key -days 3650 -out ${dir}/ca.pem -subj "/CN=Launcher LStack Authority"

  IFS=","
  for i in $CN; do
    openssl genrsa -out ${dir}/${i}.key 2048
    openssl req -new -key ${dir}/${i}.key -out ${dir}/${i}.csr -subj "/CN=${i}"
    openssl x509 -req -in ${dir}/${i}.csr -CA ${dir}/ca.pem -CAkey ${dir}/ca.key -CAcreateserial -out ${dir}/${i}.pem -days 365 -extensions v3_req
  done

}
main
