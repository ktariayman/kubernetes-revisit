# kubernetes-revisit

Simple Go app with Kubernetes deployment.

## Build Docker Image
```bash
docker build -t kubernetes-revisit:latest .
```

## Deploy to Kubernetes
```bash
# Apply deployment
kubectl apply -f deployment.yaml

# Apply service
kubectl apply -f service.yaml

# Check status
kubectl get pods
kubectl get services
```

## Access the App (Minikube)
```bash
minikube service kubernetes-revisit-service
```

## Pod Management

### View logs
```bash
kubectl logs <pod-name>
```

### Describe pod (detailed info)
```bash
kubectl describe pod <pod-name>
```

### Self-healing
When you delete a pod, Kubernetes automatically creates a new one:
```bash
kubectl delete pod <pod-name>
kubectl get pods  # New pod appears
```

### Scale replicas
```bash
# Scale to 5 replicas
kubectl scale deployment kubernetes-revisit --replicas=5

# Check all running pods
kubectl get pods
```
