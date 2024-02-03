# Building & Running (on GKE)

1. Dependencies
    Make sure you have installed:
    - Kubectl
    - Helm
    - Google Cloud SDK

2. Initiate your Google Cloud CLI
    ```bash
    gcloud init
    ```
3. Setup clusters on GKE and local `kubectl` connection credentials
    ```bash
    gcloud container clusters create kraken-cluster --location=us-west1 --num-nodes 1
    # To check if clusters are properly created
    gcloud container clusters list
    gcloud container clusters get-credentials kraken-cluster --zone us-east1
    ```

4. Install with Helm
    ```bash
    helm install <name-of-your-helm-install> helm -n <name-of-your-k8s-namespace>
    ```
    For example:
    ```bash
    helm install kraken-instance helm -n dev
    ```
    will start an Helm install called `kraken-instance` in namespace `dev`.

5. Check the deployment on Google Cloud Console