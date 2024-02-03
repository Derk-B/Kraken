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
gcloud container clusters create <name-of-your-cluster> --location=us-west1 --num-nodes 1
# Or: if you want to enable network policy enforcement on your cluster
gcloud container clusters create <name-of-your-cluster> --enable-network-policy --location=us-west1 --num-nodes 1
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


_Now you can check the deployments on Google Cloud Console._