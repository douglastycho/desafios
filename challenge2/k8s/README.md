# Run Back
```
kubectl apply -f api/deployment.yaml
kubectl expose -f api/deployment.yaml
kubectl port-forward service/challenge3-api 3000
```

# Run Front
```
kubectl apply -f front/deployment.yaml
kubectl expose -f front/deployment.yaml
kubectl port-forward service/challenge3-front 3001:80
```


Access: http://localhost:3001 to see result