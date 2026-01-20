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
