# go-client-kubernetes

## Prerequisites
- Docker Engine
- Kubernetes cluster (Eg. Minikube) configured with kubectl and nginx
- Golang

## Instructions to run

```
git clone https://github.com/mtbstn24/go-client-kubernetes.git
```
```
cd go-client-kubernetes
```
```
go run .
```
```
go run . -i IMAGE_NAME -d DEPLOYMENT_NAME -p PORT_NUMBER -r REPLICA_COUNT -kubeconfig PATH_TO_KUBE_CONFIG
```
### In windows
Edit the C:/Windows/System32/drivers/etc/hosts file to include the exposed host with proper intendations (opent the file as administrator).
  - ```127.0.0.1 EXPOSED_HOST```


## Sample Demo
```
go run . -i mtbstn24/k8s-web-hello:latest -d project2 -p 3001
```

### Output
```
mtbstn24/k8s-web-hello:latest project2 1 3001
Default namespace Pods
Pod name : nginx-deployment-5fbdf85c67-wtwf9

Default namespace deployments
Deployment name : nginx-deployment
Creating Deployment....
Created Deployment "project2"
Creating Service....
Created Service "project2-svc"
Exposing the service using the Ingress....
Service Exposed. Access the service using "http://project2.info"

----------------End of Program----------------
```
