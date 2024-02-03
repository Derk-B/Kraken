# Building & Running (the Helm Chart way)

1. Dependencies
    Make sure you have installed:
    - Docker (w/ Compose, optional if other runtime is used)
    - Microk8s (or any other K8s implementations, w/ registry add-on enabled)
    - Helm

2. Install the Helm package
    ```bash
    helm install <name-of-your-install> helm -n <name-of-your-k8s-namespace>
    ```
    For example:
    ```bash
    helm install kraken-instance helm -n dev
    ```
    will start an install called `kraken-instance` in namespace `dev`.