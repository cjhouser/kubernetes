
# Networking

## Glossary
* Network namespace: A copy of the network stack, with its own routes, firewall rules, and network devices. By default, a process inherits its network namespace from its parent. Initially, all processes share the same default network namespace from the init process.<sup>2</sup>
## Resources
[1] [A Guide to the Kubernetes Networking Model](https://sookocheff.com/post/kubernetes/understanding-kubernetes-networking-model/) - Kevin Sookocheff

[2] [ip-netns(8)](https://man7.org/linux/man-pages/man8/ip-netns.8.html)