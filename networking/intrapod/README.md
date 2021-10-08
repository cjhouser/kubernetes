# Container-to-container Networking
Containers in a pod share the same networking namespace and thus can communicate with each other using ports on localhost.

## Lab
This lab sets up two containers in a pod without any network configuration to show that the two can communicate using localhost and a port number. The counter pod runs a small server that listens for requests, increments a local counter, and sends an HTTP response back to the requester with the current count.

## Usage
1. Spin up the pod
    > kubectl apply -f intrapod.yaml
2. Start a shell in the shell container
    > kubectl exec intrapod -c shell --stdin --tty -- /bin/sh
3. The file /count in the shell container shows the number of requests that have been sent to the counter container.
    ```
    $ cat /count
    Count: 2/
    ```
4. Terminate the pod
    > kubectl delete pod intrapod
## Benefits
* Containers are co-located on node
* Intuitive organization for applications which use tightly coupled containers
## Drawbacks
* Containers are tightly coupled