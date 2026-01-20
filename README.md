# Kubernetes Revisit

## Build & Deploy
```bash
docker build -t kubernetes-revisit:latest .
kubectl apply -f deployment.yaml
kubectl apply -f service.yaml
```

## PostgreSQL Setup
```bash
kubectl apply -f postgres-storage.yaml
kubectl apply -f postgres-deployment.yaml
kubectl apply -f postgres-service.yaml
```

## Access App
```bash
minikube service kubernetes-revisit-service
minikube service kubernetes-revisit-service --url
```

## API Usage
```bash
curl <URL>/health

curl <URL>/users

curl -X POST <URL>/users/create \
  -H "Content-Type: application/json" \
  -d '{"name":"Ayman","email":"ayman@local"}'
```


## To redploy app updates..
```bash
docker build -t kubernetes-revisit:latest .
kubectl rollout restart deployment kubernetes-revisit
```

## Debugging
```bash
kubectl get pods
kubectl logs -l app=kubernetes-revisit
kubectl describe pod <pod-name>
kubectl scale deployment kubernetes-revisit --replicas=3
```
