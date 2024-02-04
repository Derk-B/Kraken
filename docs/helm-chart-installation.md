# Building & Running (the Helm Chart way)

### Dependencies
Make sure you have installed:
- Docker (w/ Compose, optional if other runtime is used)
- Microk8s (or any other K8s implementations, w/ registry add-on enabled)
- Helm
- [Helmify](https://github.com/arttor/helmify)

### Generating the Helm Chart
```bash
helmify -f k8s -r <name-of-your-chart>
```
> [!NOTE]  
> It is recommended to do modifications based on the generated Chart for your actual use-case. 
* In `kraken/templates/kraken-config.yaml`, change the value of `POSTGRES_HOSTNAME` to `{{ include "kraken.fullname" . }}-{{ .Values.krakenConfig.postgresHostname }}`
* Customize your database password in `value.yaml`
* etc ...

### Installing the Helm package
```bash
helm install <name-of-your-install> <name-of-your-chart> -n <name-of-your-k8s-namespace>
```
For example:
```bash
helm install kraken-instance helm -n dev
```
will start an install called `kraken-instance` in namespace `dev` with Chart folder `helm`.