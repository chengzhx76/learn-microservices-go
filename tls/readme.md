## SAN证书生成

https://segmentfault.com/a/1190000041068876

openssl genrsa -out ca.key 2048
openssl req -new -x509 -days 3650 -key ca.key -out ca.crt


openssl genpkey -algorithm RSA -out server.key
openssl req -new -nodes -key server.key -out server.csr -days 3650 -subj "/C=cn/OU=custer/O=custer/CN=localhost" -config /etc/pki/tls/openssl.cnf -extensions v3_req
openssl x509 -req -days 3650 -in server.csr -out server.crt -CA ca.crt -CAkey ca.key -CAcreateserial -extfile /etc/pki/tls/openssl.cnf -extensions v3_req

openssl genpkey -algorithm RSA -out client.key
openssl req -new -nodes -key client.key -out client.csr -days 3650 -subj "/C=cn/OU=custer/O=custer/CN=localhost" -config /etc/pki/tls/openssl.cnf -extensions v3_req
openssl x509 -req -days 3650 -in client.csr -out client.crt -CA ca.crt -CAkey ca.key -CAcreateserial -extfile /etc/pki/tls/openssl.cnf -extensions v3_req