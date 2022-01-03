```部署httpserver
kubectl create ns httpserver
kubectl label ns httpserver istio-injection=enabled
heml install httpserver -n httpserver ../httpserver
```
```创建HTTPS
kubectl create -n istio-system secret tls httpserver-credential --key=httpserver.io.key --cert=httpserver.io.crt
kubectl apply -f istio-specs.yaml -n httpserver
```


