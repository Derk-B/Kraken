# Presentation

The following sections are the command used in the presentation.  
For the deployment of the project, please refer to the `docs` folder.

### Part 1


### Part 2

1. Run the project using helm:
```bash
kubectl create namespace ns
helm install helm1 ./helm -n ns
```

2. Check permission of different roles
```bash
kubectl auth can-i list pods --as monitor -n ns
```

### Part 3
