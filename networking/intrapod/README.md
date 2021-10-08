# Container-to-container Networking
Containers in a pod share the same networking namespace and thus can communicate with each other using ports on localhost.

## Usage
1. Spin up the pod
    > kubectl apply -f counter-pod.yaml

## Benefits
* ?
## Drawbacks
* Intrapod network communications couples the containers in a pod.