# Kraken
Monorepo for Kraken (yet another To-do list web app)

## Building & Running (the Docker way)
The docker-compose file creates 3 containers: api, web and db.

1. Dependencies
   Make sure you have installed:
   - Docker (Compose)
   - Go

2. Copy required environment file
    ```bash
    cp .env.example .env
    # make changes to .env file
    ```

3. Build & run the managed containers
    ```bash
    # build containers with latest code changes
    docker compose build
    # start managed containers
    docker compose up -d
    ```

4. Access the app
   If you want to access the web interface or API, view the website you have to visit the hostname of the docker container.
   Run:
    ```bash
    docker inspect web
    docker inspect api
    ```
   This will print the properties of the docker container, in there you will find the hostname. Visit this address in your browser and you should see the website or API.

## Building & Running (the K8s way)

1. Dependencies
   Make sure you have installed:
   - Docker (Compose)
   - Go
   - Microk8s

2. Build the images for each container:
   ```bash
   docker-compose build
   ```
   The containers must be named: `kraken_web` and `kraken_api`
   If you still need to change the tag of the images, run:
   ```bash
   docker tag oldWebImageName kraken_web;
   docker tag oldApiImageName kraken_api
   ```

3. Make sure that the registry add-on is running for microk8s.
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

### Load balancer and ingress
Run the following command

`microk8s enable ingress; microk8s enable metallb`

Metallb will ask you for a range of ip addresses, type: `10.50.100.5-10.50.100.25`
```bash
Enter each ip addresses range: 10.50.100.5-10.50.100.25
```

Now apply the yaml files for ingress and metallb
`kubectl apply -f loadbalancer.yaml`
`kubectl apply -f ingress.yaml`

Test this using curl:
```
curl localhost/ping
{"message":"pong"}
```

### Add certificate
add cert-manager to your microk8s
```
kubectl apply -f https://github.com/cert-manager/cert-manager/releases/download/v1.11.0/cert-manager.yaml
```

Then:
```
kubectl create ns sandbox
kubectl apply -f issuer.yaml
kubectl get issuer -n sandbox # This should print an issuer.
```

#### Add self-signed certificate
`kubectl apply -f root-ca.yaml`

`kubectl get certificate -n sandbox` should print a certificate

Create a certificate authority
`kubectl apply -f ca-issuer.yaml`

Update ingress to use the CA
`kubectl apply -f ingress.yaml`

Verify the certificate:
`kctl get secret myingress-cert -n sandbox -o yaml`

Verify the tls connectgion
`openssl s_client -showcerts --connect kraken.com:443`

This should show somewhere that it is unable to verify the certificate
(that is because we signed it ourselves)

## Helm Chart
1. Make sure you install helm
   ```bash
   sudo snap install helm --classic
   ```
   or check out the official website for installation https://helm.sh/

2. Initialize an empty helm chart to deploy your app in a single action
    ```helm create todolist```

3. Reuse all the Kubernetes deployment files we created previously 
    ```cp ../k8s-services/*.yaml todolist/templates/```

4. Decide which hardcoded values from these YAML files should become variable
   create a values.yaml file that stored these properties

5. Package the helm chart
    ```bash
   helm package todolist
   ```




