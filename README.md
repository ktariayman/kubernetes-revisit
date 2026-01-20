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

## PostgreSQL Database

### Deploy PostgreSQL
```bash
# Create storage
kubectl apply -f postgres-storage.yaml

# Deploy postgres
kubectl apply -f postgres-deployment.yaml

# Create service
kubectl apply -f postgres-service.yaml

# Check postgres is running
kubectl get pods -l app=postgres
```

### API Endpoints
To get the API URL:
```bash
minikube service kubernetes-revisit-service --url
# Example output: http://192.168.49.2:30001
```

Then use the URL to access endpoints:
```bash
# Health check
curl <URL>/health

# Get all users
curl <URL>/users

# Create user
curl -X POST <URL>/users/create \
  -H "Content-Type: application/json" \
  -d '{"name":"John","email":"john@example.com"}'
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
