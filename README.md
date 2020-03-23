# K8S Cluster API Certification Generator

cluster api 에서는 3종류 파일이 필요하며 이 파일들은 같은 타입이면 모두 동일한 내용으로 제공되고 있다.

* Private Key
  * 아래 CA cert나 public key를 만들때 사용되고 이들과 pair로 존재해야한다.
* x509 Certification(for Certificate Authority)
  * private key를 통해 생성되며 대부분 인증의 용도로 사용된다.
* Public Key 
  * private key를 통해 생성되며 Service account에 인증 용도로 사용된다.

| usage           | real file                              | type               | matched file |
| --------------- | -------------------------------------- | ------------------ | ------------ |
| CA              | /etc/kubernetes/pki/ca.crt             | x509 Certification | tls.crt      |
| CA              | /etc/kubernetes/pki/ca.key             | Private Key        | tls.key      |
| ETCD CA         | /etc/kubernetes/pki/etcd/ca.crt        | x509 Certification | tls.crt      |
| ETCD CA         | /etc/kubernetes/pki/etcd/ca.key        | Private Key        | tls.key      |
| Front Proxy CA  | /etc/kubernetes/pki/front-proxy-ca.crt | x509 Certification | tls.crt      |
| Front Proxy CA  | /etc/kubernetes/pki/front-proxy-ca.key | Private Key        | tls.key      |
| Service Account | /etc/kubernetes/pki/sa.pub             | Public Key         | tls.pub      |
| Service Account | /etc/kubernetes/pki/sa.key             | Private Key        | tls.key      |
