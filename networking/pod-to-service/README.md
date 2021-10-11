# Pod-to-service Networking
Use Kubernetes services to expose pods internally to the cluster with a reliable IP address. The service will loadbalance traffic across pods that the service targets via labels.

# Usage
1. Spin up the resources
    > kubectl apply -f pod-to-service.yaml
1. Check if the resources are running
    > kubectl get all -l lab=pod-to-service
1. Start a shell on the container in `shell` pod
    > kubectl exec shell -c shell --stdin --tty -- /bin/sh
1. Run the following command in the shell multiple times. The output will show the same count a few times because the service is loadbalancing requests across the replicas of the counters deployment.
    ```
    [shell]$ echo $(wget -qO - http://10.96.7.14:8080/)
    Count: 1
    [shell]$ echo $(wget -qO - http://10.96.7.14:8080/)
    Count: 1
    [shell]$ echo $(wget -qO - http://10.96.7.14:8080/)
    Count: 2
    ```
1. Delete the lab resources
    > k delete all -l lab=pod-to-service --grace-period=0

# Benefits
* Service resources decouple clients from backends
    * For instance, a service can expose port 8000 for incoming requests, then redirect those requests to port 80 on the target pods
* Built-in load balancing
* Consistent addressing to ephemeral pods

# Drawbacks
* (something to do with load balancing. not being able to configure it)