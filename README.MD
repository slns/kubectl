# k8s

## Docker

### Build Dockerfile

```docker
docker build -t slns/hello-go .
```

### Run Container

```docker
docker run --rm -p 8087:8087 slns/hello-go
```

## KUBECTL

### Get nodes

```kubectl
kubectl get nodes
```

### Apply File

```kubectl
kubectl apply -f k8s/pod.yaml
```

### Get pods

```kubectl
kubectl get pods
```

### Generate ReplicaSet

```kubectl
kubectl apply -f k8s/replicaset.yaml
```

### Generate Deployments

```kubectl
kubectl apply -f k8s/deployment.yaml
```

### Get History

```kubectl
kubectl get replicasets
kubectl rollout history deployment goserver
```

### Generate Deployments

```kubectl
kubectl apply -f k8s/deployment.yaml
```

### Undo Deployment

```kubectl
kubctl rollout undo deployment goserver
```

### Undo Deployment to spefic revision

```kubectl
kubectl rollout history deployment goserver
kubctl rollout undo deployment goserver --to-revision=<numer_revision>
```

### Get Description 

```kubectl
kubectl describe pod <name_pod>
kubectl describe replicasets <name_replicasets>
kubectl describe deployment <name_deployment>
```

### Generate Service

```kubectl
kubectl apply -f k8s/service.yaml
```

### Get Services
```kubectl
kubectl get services
kubectl get svc
```

### Get All Endpoints Kubernets
```kubectl
kubectl proxy --port=<numberPort>
```

### Forward to port
```kubectl
kubectl port-forward service/goserver-service 8087:8087
```

### Generate ConfigMap

```kubectl
kubectl apply -f k8s/configmap-env.yaml
```