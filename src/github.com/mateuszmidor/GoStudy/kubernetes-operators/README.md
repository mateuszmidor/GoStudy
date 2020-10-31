# Kubernetes operator using Operator-SDK

Operator-SDK v1.0 workshop  
<https://www.youtube.com/watch?v=1iJKDbQzL-k>

The below RedHat OpenShift Operator-SDK workshop uses some earlier version of operator-sdk so the commands differ, but the resulting operator is the same:  
<https://www.youtube.com/watch?v=pTbuHoMp68s>

Migrating to new operator-sdk:  
<https://sdk.operatorframework.io/docs/building-operators/golang/migration/>

## Install minikube

<https://github.com/mateuszmidor/DevOpsStudy/tree/master/Kubernetes>
```bash
sudo pacman -S minikube
```

## Install operator-sdk

<https://sdk.operatorframework.io/docs/installation/install-operator-sdk/#compile-and-install-from-master>
```bash
mkdir $GOPATH/src/github.com/operator-framework
cd $GOPATH/src/github.com/operator-framework
git clone https://github.com/operator-framework/operator-sdk
cd operator-sdk
git checkout master # in this example used v1.1
make tidy
make install
```
## Steps (from nothing to working operator)

<https://sdk.operatorframework.io/docs/building-operators/golang/quickstart/>
- generate just general boilerplate code for operator  
  `operator-sdk init --project-name podset-operator --domain mateuszmidor.com --owner "Mateusz Midor" --license apache2`
- generate operator scaffolding, including the resource and controller, resulting apiVersion: podset.mateuszmidor.com/v1alpha1  
  `operator-sdk create api --kind PodSet --group podset --version v1alpha1 --resource=true --controller=true`  
- modify our PodSet definition under api/v1alpha1/podset_types.go, then run deep-copy code generators and build manager  
  `make` (used to be: operator-sdk generate k8s in earlied operator-sdk versions)
- modify our PodSet controller under controllers/podset_controller.go to handle POD replicas    
  `wget https://raw.githubusercontent.com/openshift-labs/learn-katacoda/master/operatorframework/go-operator-podset/assets/podset_controller.go`
- generate CRD yaml (config/crd/bases/podset.mateuszmidor.com_podsets.yaml), RBAC and install them into cluster  
  `make manifests`
- install new CRD into cluster  
  `make install # this actually also runs: make manifests`
- run operator OUTSIDE the cluster, as a separate go app (it connects to kubernetes API server that is accessible from outside the cluster)  
  `make run`  
- see new CRD "podsets" is now available:  
  `kubectl api-resources`  
```
NAME                              SHORTNAMES   APIGROUP                       NAMESPACED   KIND 
...   
podsets                                        podset.mateuszmidor.com        true         PodSet
...  
```
- create example podset  
  `kubectl apply -f example-podset.yaml`
- see podset created  
  `kubectl get podset`  
```
NAME             AGE
example-podset   26s
```
- see pods managed by podset  
  `kubectl get pods`
```
NAME                      READY   STATUS    RESTARTS   AGE
example-podset-podgkgdb   1/1     Running   0          77s
example-podset-podmqsvv   1/1     Running   0          77s
example-podset-podxjpp2   1/1     Running   0          76s
```
- update podset replicas from 3 to 4 and see it gets reconciled  
`kubectl patch podset example-podset --type='json' -p='[{"op": "replace", "path": "/spec/replicas", "value":4}]'`
```json
2020-10-28T11:10:05.392+0100    INFO    controllers.PodSet      Scaling up pods {"Currently available": 3, "Required replicas": 4}
2020-10-28T11:10:05.406+0100    DEBUG   controller      Successfully Reconciled {"reconcilerGroup": "podset.mateuszmidor.com", "reconcilerKind": "PodSet", "controller": "podset", "name": "example-podset", "namespace": "default"}
```

## Exercises 

outdated; uses operator-sdk v0.6 while there is v1.1 out there and the commands have changed
<http://workshop.coreostrain.me/exercises/>