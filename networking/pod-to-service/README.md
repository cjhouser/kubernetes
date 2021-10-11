# Pod-to-service Networking
Use Kubernetes services to expose pods internally to the cluster with a reliable IP address. The service will loadbalance traffic across pods that the service targets via labels.

# Usage
1. Spin up the resources
    > kubectl apply -f pod-to-service.yaml
1. Check if the resources are running. Note the service's CLUSTER-IP.
    > kubectl get all -l lab=pod-to-service
1. Start a shell on the container in `shell` pod
    > kubectl run shell -i --tty --image=bash --rm
1. Run the following command in the shell. The output will show the same count a few times because the service is loadbalancing requests across the replicas of the counters deployment.
    > for _ in {1..10}; do echo $(wget -qO - http://\<service-ip>:8080/); done
1. Delete the lab resources
    > kubectl delete all -l lab=pod-to-service

# Benefits
* Service resources decouple clients from backends
    * For instance, a service can expose port 8000 for incoming requests, then redirect those requests to port 80 on the target pods
* Built-in load balancing
* Consistent addressing to ephemeral pods

# Drawbacks
* (something to do with load balancing. not being able to configure it)