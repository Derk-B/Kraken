# Building & Running (the Helm Chart way)

1. Dependencies
   Make sure you have installed:
   - Docker (w/ Compose)
   - Go
   - Microk8s (or any other K8s implementations, w/ registry add-on enabled)
   - Helm

2. Initialize an empty helm chart to deploy your app in a single action
   ```bash
   helm create todolist
   ```

3. Reuse all the Kubernetes deployment files we created previously 
   ```bash
   cp ../k8s-services/*.yaml todolist/templates/
   ```

4. Decide which hardcoded values from these YAML files should become variable
   create a values.yaml file that stored these properties

5. Package the helm chart
   ```bash
   helm package todolist
   ```