#!/bin/sh


## SM CA

gmssl genpkey -algorithm EC -pkeyopt ec_paramgen_curve:sm2p256v1 -out ca-gm-key.pem
gmssl req -x509 -new -nodes -key ca-gm-key.pem -subj "/CN=myca.com" -days 5000 -out ca-gm-cert.crt


### server: sign key and cert

gmssl genpkey -algorithm EC -pkeyopt ec_paramgen_curve:sm2p256v1 -out server-gm-sign-key.pem
gmssl req -new -key server-gm-sign-key.pem -subj "/CN=umf.com" -out server-gm-sign.csr
gmssl x509 -req -in server-gm-sign.csr -CA ca-gm-cert.crt -CAkey ca-gm-key.pem -CAcreateserial -out server-gm-sign-cert.crt -days 5000 -extfile ./server.cnf -extensions ext

gmssl verify -CAfile ca-gm-cert.crt server-gm-sign-cert.crt

### server: enc key and cer

gmssl genpkey -algorithm EC -pkeyopt ec_paramgen_curve:sm2p256v1 -out server-gm-enc-key.pem
gmssl req -new -key server-gm-enc-key.pem -subj "/CN=umf.com" -out server-gm-enc.csr
gmssl x509 -req -in server-gm-enc.csr -CA ca-gm-cert.crt -CAkey ca-gm-key.pem -CAcreateserial -out server-gm-enc-cert.crt -days 5000 -extfile ./server.cnf -extensions ext

gmssl verify -CAfile ca-gm-cert.crt server-gm-enc-cert.crt

### client: auth key and cert

gmssl genpkey -algorithm EC -pkeyopt ec_paramgen_curve:sm2p256v1 -out client-gm-auth-key.pem
gmssl req -new -key client-gm-auth-key.pem -subj "/CN=client1.com" -out client-gm-auth.csr
gmssl x509 -req -in client-gm-auth.csr -CA ca-gm-cert.crt -CAkey ca-gm-key.pem -CAcreateserial -out client-gm-auth-cert.crt -days 5000 -extfile ./client.cnf -extensions ext

gmssl verify -CAfile ca-gm-cert.crt client-gm-auth-cert.crt