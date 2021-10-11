# Pod-to-service Networking
Kubernetes services can expose pods internally to the cluster with a reliable IP address. A service will balance traffic across target pods.

# Lab
This lab sets up a deployment with one replica (two pods).
The deployment contains two small HTTP servers.
Each listen for requests on port 8080, increment their local counter when receiving a request, and send an HTTP response back to the requester with the current local count.
Instead of using the IP addresses of the pods directly, a service is put in place that knows how to find those pods in the cluster, whatever their IP addresses may be.

# Usage
1. Spin up the resources
    > kubectl apply -f pod-to-service.yaml
1. Check if the resources are running. Note the service's CLUSTER-IP.
    > kubectl get all -l lab=pod-to-service
1. Start a shell inside the cluster
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
* Service load balancing technique is dependent on the kube-proxy's mode:
    * iptables: random selection
    * IPVS: multiple techniques
    * userspace: round-robin + permute list of pods