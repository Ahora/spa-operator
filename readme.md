# Ahora SPA Operator
## Introduction
An operator which manage [SPA](https://github.com/ahora/spa) crd in kubrenetes.

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
$ kubectl create -f deploy/crds/ahora.dev_v1alpha1_spa_cr.yaml
```

## CR Example
``` yaml
apiVersion: ahora.dev/v1alpha1
kind: SPA
metadata:
  name: example-spa
spec:
  SPAArchiveURL: https://ahora.dev/archive.tar.gz
  replicas: 2
  host: www.yourdomain.com
```

## TODOs
* Support SSL.
* Support custom backend for API service (/api)
* Patch service, ingress & Deployment if needed
* Writing tests
* Add limits