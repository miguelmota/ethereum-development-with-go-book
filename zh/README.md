---
概述: 
这本迷你书可以帮助你学习如何用Go语言部署，编译，与智能合约交互，发交易，使用Swarm和Whisper协议。就这么简单：）

---

# 用Go来做以太坊开发

这本迷你书的本意是给任何想用Go进行以太坊开发的同学一个概括的介绍。本意是如果你已经对以太坊和Go有一些熟悉，但是对于怎么把两者结合起来还有些无从下手，那这本书就是一个好的起点。你会学习如何用Go与智能合约交互，还有如何完成一些日常的查询和任务。

这本书里有很多我希望我当初学习用Go以太坊开发的时候能有的代码范例。你上手Go语言以太坊开发的大部分所需知识，这本书里面都会手把手介绍到。

当然了，以太坊还是一直在飞速的发展的进化的。所以难免会有些过期的内容，或者你认为有可以值得提升的地方，那就是你提 [issue](https://github.com/miguelmota/ethereum-development-with-go-book/issues) 或者 [pull request](https://github.com/miguelmota/ethereum-development-with-go-book/pulls) 的好机会了：）这本书是完全开源并且免费的，你可以在 [github](https://github.com/miguelmota/ethereum-development-with-go-book) 上看到源码.

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

#### About the Author

This book was written by [Miguel Mota](https://github.com/miguelmota), a software developer working in the blockchain space from the always sunny Southern California. You can find him on Twitter [@miguelmota](https://twitter.com/miguelmotah)

---

Enough with the introduction, let's get [started](../client)!
