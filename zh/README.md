---
概述: 这本迷你书可以帮助你学习如何用Go语言部署，编译，与智能合约交互，发交易，使用Swarm和Whisper协议。就这么简单：）
---

# 用Go来做以太坊开发

这本迷你书的本意是给任何想用Go进行以太坊开发的同学一个概括的介绍。本意是如果你已经对以太坊和Go有一些熟悉，但是对于怎么把两者结合起来还有些无从下手，那这本书就是一个好的起点。你会学习如何用Go与智能合约交互，还有如何完成一些日常的查询和任务。

这本书里有很多我希望我当初学习用Go以太坊开发的时候能有的代码范例。你上手Go语言以太坊开发的大部分所需知识，这本书里面都会手把手介绍到。

当然了，以太坊还是一直在飞速的发展的进化的。所以难免会有些过期的内容，或者你认为有可以值得提升的地方，那就是你提 [issue](https://github.com/miguelmota/ethereum-development-with-go-book/issues) 或者 [pull request](https://github.com/miguelmota/ethereum-development-with-go-book/pulls) 的好机会了：）这本书是完全开源并且免费的，你可以在 [github](https://github.com/miguelmota/ethereum-development-with-go-book) 上看到源码.

#### 在线电子书

[https://goethereumbook.org](https://goethereumbook.org/)

#### 电子书

电子书有三种格式。

- [PDF](https://goethereumbook.org/ethereum-development-with-go-zh.pdf)
- [EPUB](https://goethereumbook.org/ethereum-development-with-go-zh.epub)
- [MOBI](https://goethereumbook.org/ethereum-development-with-go-zh.mobi)

## 介绍

> 以太坊是一个开源，公开，基于区块链的分布式计算平台和具备智能合约（脚本）功能的操作系统。它通过基于交易的状态转移支持中本聪共识的一个改进算法。

-[维基百科](https://en.wikipedia.org/wiki/Ethereum)

以太坊是一个区块链，允许开发者创建完全去中心化运行的应用程序，这意味着没有单个实体可以将其删除或修改它。部署到以太坊上的每个应用都由以太坊网络上每个完整客户端执行。

#### Solidity

Solidity是一种用于编写智能合约的图灵完备编程语言。Solidity被编译成以太坊虚拟机可执行的字节码。

#### go-ethereum

本书中，我们将使用Go的官方以太坊实现[go-ethereum](https://github.com/ethereum/go-ethereum)来和以太坊区块链进行交互。Go-ethereum，也被简称为Geth，是最流行的以太坊客户端。因为它是用Go开发的，当使用Golang开发应用程序时，Geth提供了读写区块链的一切功能。

本书的例子在go-ethereum版本`1.8.10-stable`和Go版本`go1.10.2`下完成测试。

#### Block Explorers

[Etherscan](https://etherscan.io)是一个用于查看和深入研究区块链上数据的网站。这些类型的网站被称为*区块浏览器*，因为它们允许您查看区块（包含交易）的内容。区块是区块链的基础构成要素。区块包含在已分配的出块时间内开采出的所有交易数据。区块浏览器也允许您查看智能合约执行期间释放的事件以及诸如支付的gas和交易的以太币数量等。

### Swarm and Whisper

我们还将深入研究蜂群(Swarm)和耳语(Whisper)，分别是一个文件存储协议和一个点对点的消息传递协议，它们是实现完全去中心化和分布式应用程序需要的另外两个核心。

<img src="https://user-images.githubusercontent.com/168240/41317815-2e287afe-6e4b-11e8-89d8-4ec959988b64.png" width="600"/>

<sup><sub><a href="https://ethereum.stackexchange.com/a/388/5093">图片来源</a></sub></sup>

## 寻求帮助

寻求Go(Golang)帮助可以加入[gophers slack]((https://invite.slack.golangbridge.org/))上的[#ethereum]((https://gophers.slack.com/messages/C9HP1S9V2/))频道。

---

介绍部分足够了，让我们[开始](../zh/client)吧。
