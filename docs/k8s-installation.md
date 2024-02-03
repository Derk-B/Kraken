# Building & Running (the K8s way)

### Dependencies
Make sure you have installed:
- Docker (w/ Compose)
- Microk8s (or any other K8s implementations, w/ registry add-on enabled)

### Build the images for each container:
```bash
GIT_COMMIT=$(git rev-parse HEAD) docker-compose build
```
The containers must be named: `kraken_web` and `kraken_api`

### Push the image to the local K8s registry:
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