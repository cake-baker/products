# Products - Sample Backend

Docker Image: [cakebakery/products:v1]()

## 1. Test Sample Backend

Resource Usage
```yaml
resources:
  limits:
    cpu: "5m"
    memory: "10Mi"
  requests:
    cpu: "2m"
    memory: "5Mi"
```

### 1.1. Docker

```sh
$ docker run -d -p 8080:8080 --name products-backend cakebakery/products:v1
```

#### Sample cURL
```sh
curl http://localhost:8080/products
```

```json
[{"id":1,"name":"Apples","category":"Food","price":149},{"id":2,"name":"Macaroni and Cheese","category":"Food","price":769},{"id":3,"name":"ABC Smart TV","category":"Electronics","price":39999},{"id":4,"name":"Motor Oil","category":"Automobile","price":2288},{"id":5,"name":"Floral Sleeveless Blouse","category":"Clothing","price":2150}]
```

#### Swagger Documentation

- [https://app.swaggerhub.com/apis-docs/renuka-fernando/Products/v1](https://app.swaggerhub.com/apis-docs/renuka-fernando/Products/v1)
- [Swagger.yaml](swagger.yaml)

####Remove running service
```
docker rm -f products-backend
```

## 2. Build Docker Image

Rename `IMAGE_NAME` and `VERSION` in the script `build-docker.sh` and execute the script.
```sh
$ ./build-docker.sh
```
