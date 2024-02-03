For Derk's part of the presentation

For the part where i show that the roles are working
run the project using helm:
  ```
kubectl create namespace ns
helm install helm1 ./helm -n ns
```
run kubectl auth can-i list pods --as monitor -n ns
