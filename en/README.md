---
description: Learn how to deploy, compile, interact with smart contracts, send transactions, use the swarm and whisper protocols, and much more with this little guide book on Ethereum Development with Go.
---

# Ethereum Development with Go

This little guide book is to serve as a general help guide for anyone wanting to develop Ethereum applications using the Go programming language. It's meant to provide a starting point if you're already pretty familiar with Ethereum and Go but don't know where to to start on bringing it all together. You'll learn how to interact with smart contracts and perform general blockchain tasks and queries using Golang.

This book is composed of many examples that I wish I had encountered before when I first started doing Ethereum development with Go. This book will walk you through most things that you should be aware of in order for you to be a productive Ethereum developer using Go.

Ethereum is quickly evolving and things may go out of date sooner than anticipated. I strongly suggest opening an [issue](https://github.com/miguelmota/ethereum-development-with-go-book/issues) or making a [pull request](https://github.com/miguelmota/ethereum-development-with-go-book/pulls) if you observe things that can be improved. This book is completely open and free and available on [github](https://github.com/miguelmota/ethereum-development-with-go-book).

#### Online

[https://goethereumbook.org](https://goethereumbook.org/)

#### E-book

The e-book is avaiable in different formats.

- [PDF](https://goethereumbook.org/ethereum-development-with-go.pdf)
- [EPUB](https://goethereumbook.org/ethereum-development-with-go.epub)
- [MOBI](https://goethereumbook.org/ethereum-development-with-go.mobi)

## Introduction

> Ethereum is an open-source, public, blockchain-based distributed computing platform and operating system featuring smart contract (scripting) functionality. It supports a modified version of Nakamoto consensus via transaction based state transitions.

-[Wikipedia](https://en.wikipedia.org/wiki/Ethereum)

Ethereum is a blockchain that allows developers to create applications that can be ran completely decentralized, meaning that no single entity can take it down or modify it. Each application deployed to Ethereum is executed by every single full client on the Ethereum network.

#### Solidity

Solidity is a Turing complete programming language for writing smart contracts. Solidity gets compiled to bytecode which is what the Ethereum virtual machine executes.

#### go-ethereum

In this book we'll be using the [go-ethereum](https://github.com/ethereum/go-ethereum), the official Ethereum implementation in Go, to interact with the blockchain. Go-ethereum, also known as *geth* for short, is the most popular Ethereum client and because it's in Go, it provides everything we'll ever need for reading and writing to the blockchain when developing applications using Golang.

The examples in this book were tested with go-ethereum version `1.8.10-stable` and Go version `go1.10.2`.

#### Block Explorers

[Etherscan](https://etherscan.io/) is a website for exploring and drilling down on data that lives on the blockchain. These type of websites are known as *Block Explorers* because they allow you to explore the contents of blocks (which contain transactions). Blocks are fundamental components of the blockchain. The block contains the data of all the transactions that have been mined within the allocated block time. The block explorer also allows you to view events that were emitted during the execution of the smart contract as well as things such as how much was paid for the gas and amount of ether was transacted, etc.

### Swarm and Whisper

We'll also be diving a little bit into Swarm and Whisper, a file storage protocol, and a peer-to-peer messaging protocol respectively, which are the other two pillars required for achieving completely decentralized and distributed applications.

<img src="https://user-images.githubusercontent.com/168240/41317815-2e287afe-6e4b-11e8-89d8-4ec959988b64.png" width="600"/>

<sup><sub><a href="https://ethereum.stackexchange.com/a/388/5093">image credit</a></sub></sup>

## Support

Join the [#ethereum](https://gophers.slack.com/messages/C9HP1S9V2/) channel on the [gophers slack](https://invite.slack.golangbridge.org/) for Go (golang) help.

The [Ethereum StackExchange](https://ethereum.stackexchange.com/) is also a great place to ask general Ethereum question and Go specific questions.

---

Enough with the introduction, let's get [started](../en/client)!
