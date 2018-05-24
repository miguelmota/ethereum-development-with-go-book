# Introduction

## What is Ethereum

> Ethereum is an open-source, public, blockchain-based distributed computing platform and operating system featuring smart contract (scripting) functionality.[3] It supports a modified version of Nakamoto consensus via transaction based state transitions. [Wikipedia](https://en.wikipedia.org/wiki/Ethereum)

Ethereum is a blockchain that allows developers to create applications that can be ran completely decentralized, meaning that no single entity can take it down or modify it. Each application deployed to Ethereum is executed by every single full client on the Ethereum network.

Solidity is a turing complete programming language for writing smart contracts. Solidity gets compiled to bytecode which is what the Ethereum virtual machine exectues.

Etherscan is a website for exploring and drilling down on data that lives on the blockchain. These type of website are known as "Block Explorers" because they allow you to explore the contents of a block, which is a fundamental component of the blockchain. The block contains the data of all the transactions that have been mined within the allocated block time. The block explorer also allows you to view events that were emitted during the execution of the smart contract as well as things such as how much was paid for the gas and amount of ether was transacted, etc.

In this book we'll be using the [go-ethereum](https://github.com/ethereum/go-ethereum), the official Ethereum implementation in Go, to interact with the blockchain. Go-ethereum, also known as `geth` for short, is the most popular Ethereum client and because it's in Go, it provides everything we'll ever need for reading and writing to the blockchain.
