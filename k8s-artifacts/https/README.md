# HTTPS Products Sample

## 1. Create Certs

```sh
CN=products-https
SAN1=localhost
SAN2=products

openssl req -x509 -newkey rsa:4096 -nodes -days 365 \
    -subj "/O=foo.com Inc./CN=${CN}" \
    -reqexts SAN \
        -config <(cat /etc/ssl/openssl.cnf \
            <(printf "\n[SAN]\nsubjectAltName=DNS:${SAN2},DNS:${SAN2}")) \
    -keyout products-certs.key \
    -out products-certs.crt
```

## 2. Create K8s Secret

```sh
kubectl create secret tls https-products-certs \
  --cert=products-certs.crt \
  --key=products-certs.key
```

## 3. Create Deployment

```sh
kubectl apply -f products-https.yaml
```

## 4. Clean Up

```sh
rm products-certs.crt products-certs.key
kubectl delete secret https-products-certs
kubectl delete -f products-https.yaml
```


