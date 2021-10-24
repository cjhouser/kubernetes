# Volumes
Ephemeral volumes are bound to the lifetime of a pod.

## Lab
The shell pod, created via a deployment, manages a container that may be accessed and modified via a remote shell. The interesting bit of the container is the "/volume" directory which is the mount point for an ephemeral volume.

## Usage
1. Spin up the lab resources
    > kubectl apply -f volumes.yaml
1. Get the name of the pod created by the deployment
    > kubectl get pods -l lab=volumes
1. Open a remote shell to the container
    > kubectl exec -it \<pod-name> -- sh
1. Create a file in mounted directory, then exit
    ```
    (container)# touch /volume/emptyFile
    (container)# exit
    ```
1. Delete the pod
    > kubectl delete pod \<pod-name>
1. Remote into the new pod to verify that the file created is still there
    > (container)# ls /volume
1. Clean up the lab resources
    > kubectl delete all -l lab=volumes

## Benefits
* Shared filesystem for containers in the same pod.

## Drawbacks
* Volume contents are lost when the pod is deleted.

## Bonus
On a cluster with more than a single node, what happens when a node fails that is running the pod with the ephemeral volume? A new pod will eventually be scheduled on an available node. Ephemeral volumes on the original pod are not preserved.