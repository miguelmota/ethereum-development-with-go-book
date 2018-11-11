---
description: Tutorial on swarm with Go.
---

# Swarm

Swarm in Ethereum's decentralized and distributed storage solution, comparable to IPFS. Swarm is a peer to peer data sharing network in which files are addressed by the hash of their content. Similar to Bittorrent, it is possible to fetch the data from many nodes at once and as long as a single node hosts a piece of data, it will remain accessible everywhere. This approach makes it possible to distribute data without having to host any kind of server - data accessibility is location independent. Other nodes in the network can be incentivised to replicate and store the data themselves, obviating the need for hosting services when the original nodes are not connected to the network.

Swarm's incentive mechanism, Swap (Swarm Accounting Protocol), is a protocol by which peers in the Swarm network keep track of chunks delivered and received and the resulting (micro-) payments owed. On its own, SWAP can function in a wider context however it's usually presented as a generic micropayment scheme suited for pairwise accounting between peers. while generic by design, the first use of it is for accounting of bandwidth as part of the incentivisation of data transfer in the Swarm decentralised peer to peer storage network.

