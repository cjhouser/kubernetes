# Pod-to-pod Networking
Every pod gets a cluster-internal IP address which can be used to communicate with other pods.

## Lab
This lab sets up two pods with one container in each to show that the two containers can communicate using cluster-internal IP addresses and ports. `interpod-a` contains a small HTTP server that listens for requests on port 8080, increments a local counter when receiving a request, and sends an HTTP response back to the requester with the current count.

## Usage
1. Spin up the pods
    > kubectl apply -f interpod.yaml
1. Note the IP address of `interpod-a`
    > kubectl get pods -o wide
1. Start a shell in the container on pod `interpod-b`
    > kubectl exec intrapod -c shell --stdin --tty -- /bin/sh

1. Send a request to the counter server on `interpod-a` and print the response to stdout
    > wget -qO - http://<i>\<interpod-a-ip></i>:8080/
1. Terminate the pods
    > kubectl delete pods -l lab=intrapod --grace-period=0

## Benefits
* Pods can be treated like standard machines from a network perspective
## Drawbacks
* IP addresses are just as ephemeral as the pods they are associated with