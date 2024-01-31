# Kraken
Monorepo for Kraken (yet another To-do list web app)

## Building & Running
There are three ways to install and run Kraken.
1. [With Vanilla Docker](docs/docker-installation.md)
2. [With Kubernetes](docs/k8s-installation.md)

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




