# Externally Exposed Service
A service can expose an application. For minikube (and cloud environments), this is done by adding the LoadBalancer type to the service spec.

## Lab
The counter pod runs a small server that listens for requests, increments a local counter, and sends an HTTP response back to the requester with the current count. This application is exposed by a service to allow access from outside the cluster.

## Usage
1. Spin up the resources
    > kubectl apply -f external-to-service.yaml
1. Ensure the resources are ready
    > kubectl get all -l lab=external-to-service
1. Get the external IP of the service. This step may differ depending on your Kubernetes environment. For minikube:
    > minikube service service
1. Send requests to the service from the host machine
    > echo $(wget -qO - http://\<external-ip>:<external-port>/)
1. Delete the lab resources
    > kubectl delete all -l lab=external-to-service
## Benefits

## Drawbacks