# Building & Running (on GKE)

### Dependencies
Make sure you have installed:
- Kubectl
- Helm
- Google Cloud SDK

### Initiate the Google Cloud CLI
```bash
gcloud init
```

### Setup clusters on GKE and local `kubectl` connection credentials
```bash
gcloud container clusters create <name-of-your-cluster> --location=us-east1 --num-nodes 1
# Or: if you want to enable network policy enforcement on your cluster
gcloud container clusters create <name-of-your-cluster> --enable-network-policy --location=us-east1 --num-nodes 1
# To check if clusters are properly created
gcloud container clusters list
gcloud container clusters get-credentials kraken-cluster --zone us-east1
```
> [!NOTE]  
> If you decide not to enable network policy enforcement, the network policies in the project won't work. 

### (Optional) Enable network policy enforcement
If you have enabled the network policy enforcement on step 3, you can skip this step. 
```bash
gcloud container clusters update <name-of-your-cluster> --update-addons=NetworkPolicy=ENABLED
gcloud container clusters update <name-of-your-cluster> --enable-network-policy
```

### Install with Helm
```bash
helm install <name-of-your-helm-install> helm -n <name-of-your-k8s-namespace>
```
For example:
```bash
helm install kraken-instance helm -n dev
```
will start an Helm install called `kraken-instance` in namespace `dev`.  

_Now you can check the deployments on Google Cloud Console, or with other tools of your choice._  

### (Optional) Setup local domain name and hosts
Depending on your operating system, you can edit the `hosts` file to achieve optimal testing experience.  

> [!WARNING]  
> Although the IP address is public, it's not a fixed address, which mean it may not survive over time or
> across multiple deployments. Therefore, its purpose is mostly for proof-of-concept. 

1. Get public IP address from GKE
```bash
kubectl get ingress
```
Under the `ADDRESS` column there's your IP of ingress.  

2. Edit your `hosts` file
For most Linux distributions and macOS, you can edit `/etc/hosts` and add the following content:
```
<your-public-ip> kraken.com
<your-public-ip> api.kraken.com
```

3. Start the broswer with less strict security policies
Assuming your browser is `chromium`-based, you need to close all the current browser windows, can launch with:
```bash
<path-to-your-browser-binary> --ignore-certificate-errors
# For example, Google Chrome on Linux
`which google-chrome` --ignore-certificate-errors
# Or, Google Chrome on macOS
/Applications/Google\ Chrome.app/Contents/MacOS/Google\ Chrome --ignore-certificate-errors
```
After the browser launches, you can navigate to [https://kraken.com](https://kraken.com). 