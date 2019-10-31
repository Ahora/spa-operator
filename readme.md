# Ahora SPA Operator
## Introduction 
Manage [SPA](https://github.com/ahora/spa) deployment with crds in kubernetes.

**You are more than welcome to contribute to this project.**

## Usage
``` sh
# Setup Service Account
$ kubectl create -f deploy/service_account.yaml
# Setup RBAC
$ kubectl create -f deploy/role.yaml
$ kubectl create -f deploy/role_binding.yaml
# Setup the CRD
$ kubectl create -f deploy/crds/ahora.dev_spas_crd.yaml
# Deploy the spa-operator
$ kubectl create -f deploy/operator.yaml

# Create a SPA CR
$ kubectl create -f examples/example.yaml
```

## CR Example
``` yaml
apiVersion: ahora.dev/v1alpha1
kind: SPA
metadata:
  name: example-spa
  annotations:
    nginx.ingress.kubernetes.io/ssl-redirect: "true"
spec:
  SPAArchiveURL: https://storage.googleapis.com/ahora-spa-archives/spa-demo.tar.gz
  replicas: 2
  hosts: 
  - www.example.com
  - example.com
  tls:
  - secretName: example-ssl
    hosts:
    - www.example.com
    - example.com
  livenessProbe:
    httpGet:
      path: /status
      port: 80
    initialDelaySeconds: 3
    periodSeconds: 3
  readinessProbe:
    httpGet:
      path: /status
      port: 80
    initialDelaySeconds: 3
    periodSeconds: 3
  resources:
    requests:
      memory: "32Mi"
      cpu: "50m"
    limits:
      memory: "64Mi"
      cpu: "100m"
  paths:
  - path: /api
    backend:
      serviceName: backendservice
      servicePort: 80
```

## TODOs
* Writing tests