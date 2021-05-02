# Products - Sample Backend

Docker Image: [cakebakery/products:v1](https://hub.docker.com/r/cakebakery/products)

## 1. Test Sample Backend

Resource Usage
```yaml
resources:
  limits:
    cpu: "5m"
    memory: "20Mi"
  requests:
    cpu: "2m"
    memory: "10Mi"
```

### 1.1. Docker

```sh
$ docker run -d -p 8080:8080 --name products-backend cakebakery/products:v1
```

#### Swagger Documentation

- [https://app.swaggerhub.com/apis-docs/renuka-fernando/Products/v1](https://app.swaggerhub.com/apis-docs/renuka-fernando/Products/v1)
- [Swagger.yaml](swagger.yaml)

#### Sample cURL
```sh
curl -X 'GET' \
  'http://localhost:8080/products' \
  -H 'accept: application/json'
```

```json
[
    {
        "id": 1,
        "name": "Apples",
        "category": "Food",
        "price": 149
    },
    {
        "id": 2,
        "name": "Macaroni and Cheese",
        "category": "Food",
        "price": 769
    },
    {
        "id": 3,
        "name": "ABC Smart TV",
        "category": "Electronics",
        "price": 39999
    },
    {
        "id": 4,
        "name": "Motor Oil",
        "category": "Automobile",
        "price": 2288
    },
    {
        "id": 5,
        "name": "Floral Sleeveless Blouse",
        "category": "Clothing",
        "price": 2150
    }
]
```

Other sample cURL
```sh
curl -X 'POST' \
  'http://localhost:8080/products' \
  -H 'accept: application/json' \
  -H 'Content-Type: applicatoin/json' \
  -d '{
  "id": 3,
  "name": "ABC Smart TV",
  "category": "Electronics",
  "price": 39999
}'
```
```sh
curl -X 'GET' \
  'http://localhost:8080/products/3' \
  -H 'accept: application/json'
```
```sh
curl -X 'PUT' \
  'http://localhost:8080/products/3' \
  -H 'accept: application/json' \
  -H 'Content-Type: applicatoin/json' \
  -d '{
  "id": 3,
  "name": "ABC Smart TV",
  "category": "Electronics",
  "price": 39999
}'
```
```sh
curl -X 'DELETE' \
  'http://localhost:8080/products/3' \
  -H 'accept: application/json'
```

#### Remove running service
```
docker rm -f products-backend
```

## 2. Build Docker Image

Rename `IMAGE_NAME` and `VERSION` in the script `build-docker.sh` and execute the script.
```sh
$ ./build-docker.sh
```
