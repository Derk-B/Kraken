# Kraken
Monorepo for Software Containerization project

## Dependencies
Make sure you have installed:
- Docker (Compose)
- Go

## Building & Running
The docker-compose file creates 3 containers: api, web and db.

1. Build the web application
    ```bash
    cd web
    go build
    ```

2. Copy required environment file
    ```bash
    cp .env.example .env; cp .env.example ./api/.env
    # make changes to .env file
    ```

3. Build & run the managed containers
    ```bash
    # build containers with latest code changes
    docker compose build
    # start managed containers
    docker compose up -d
    ```

4. Access each deployed component  
   If you want to access the web interface or API, view the website you have to visit the hostname of the docker container.
   Run:
    ```bash
    docker inspect web
    docker inspect api
    ```
This will print the properties of the docker container, in there you will find the hostname. Visit this address in your
browser and you should see the website or API.

## Setup k8s
Make sure you build the images for each container:
```bash
docker-compose build
```
The containers must be named: `kraken_web` and `kraken_api`
If you still need to change the tag of the images, run:
```bash
docker tag oldWebImageName kraken_web;
docker tag oldApiImageName kraken_api
```

Make sure that the registry add-on is running for microk8s.
```bash
> sudo microk8s status
...
addons:
    enabled:
        registry ...
    disabled:
        registry # If registry is mentioned here instead, run: sudo microk8s enable registry
    ...
```

You might need to add the file: `/etc/docker/daemon.json` and put the following inside the file:
```json
{
    "insecure-registries": ["localhost:32000"]
}
```
Test the registry using:
```bash
> curl http://localhost:32000/v2/

[] # Works correctly

# or

curl: (7) Failed to connect to localhost port 32000 after 0 ms: Connection refused # Does not work, check if registry is enabled!
```

Push the image to the registry:
```bash
docker tag kraken_api localhost:32000/kraken_api
docker push localhost:32000/kraken_api
```

The deployment file for the api is already created: `k8s-services/kraken-api-deployment.yaml`

Go to the root of the project and create the pods for the api:
```bash
kubectl apply -f k8s-services/kraken-api-deployment.yaml
```

Also create the ClusterIP using:
```bash
kubectl apply -f k8s-services/kraken-api-service.yaml
```

To confirm everything is working, type:
```
kubectl get svc
```
This should print the ip of the clusterIp (kraken-api-service)
Navigate to this api in your browser: `[clusterip]:8080/ping`, this should return `{"message":"pong"}`