<p align="center">
  <a href="https://goethereumbook.org"><img src="https://github.com/miguelmota/ethereum-development-with-go-book/raw/master/assets/cover.jpg" width="320" alt="Book cover" /></a>
</p>
<br>

# Ethereum Development with Go

> A little guide book on [Ethereum](https://www.ethereum.org/) Development with [Go](https://golang.org/) (golang)

[![License](http://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/miguelmota/merkletreejs/master/LICENSE)
[![Mentioned in Awesome Go](https://awesome.re/mentioned-badge.svg)](https://github.com/avelino/awesome-go)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg)](#contributing)

## Online

[https://goethereumbook.org](https://goethereumbook.org/)

## E-book

The e-book is avaiable in different formats.

- [PDF](https://goethereumbook.org/ethereum-development-with-go.pdf)
- [EPUB](https://goethereumbook.org/ethereum-development-with-go.epub)
- [MOBI](https://goethereumbook.org/ethereum-development-with-go.mobi)

## Languages

* [English](en/)
* [Chinese中文](zh/)

## Contents

* [Introduction](en/README.md)
* [Client](en/client/README.md)
  * [Setting up the Client](en/client-setup/README.md)
* [Accounts](en/accounts/README.md)
  * [Account Balances](en/account-balance/README.md)
  * [Account Token Balances](en/account-balance-token/README.md)
  * [Generating New Wallets](en/wallet-generate/README.md)
  * [Keystores](en/keystore/README.md)
  * [HD Wallets](en/hd-wallet/README.md)
  * [Address Check](address-check/README.md)
* [Transactions](en/transactions/README.md)
  * [Querying Blocks](en/block-query/README.md)
  * [Querying Transactions](en/transaction-query/README.md)
  * [Transferring ETH](en/transfer-eth/README.md)
  * [Transferring Tokens](en/transfer-tokens/README.md)
  * [Subscribing to New Blocks](en/block-subscribe/README.md)
  * [Create Raw Transaction](en/transaction-raw-create/README.md)
  * [Send Raw Transaction](en/transaction-raw-send/README.md)
* [Smart Contracts](en/smart-contracts/README.md)
  * [Smart Contract Compilation & ABI](en/smart-contract-compile/README.md)
  * [Deploying a Smart Contract](en/smart-contract-deploy/README.md)
  * [Loading a Smart Contract](en/smart-contract-load/README.md)
  * [Querying a Smart Contract](en/smart-contract-read/README.md)
  * [Writing to a Smart Contract](en/smart-contract-write/README.md)
  * [Reading Smart Contract Bytecode](en/smart-contract-bytecode/README.md)
  * [Querying an ERC20 Token Smart Contract](en/smart-contract-read-erc20/README.md)
* [Event Logs](en/events/README.md)
  * [Subscribing to Event Logs](en/event-subscribe/README.md)
  * [Reading Event Logs](en/event-read/README.md)
  * [Reading ERC-20 Token Event Logs](en/event-read-erc20/README.md)
  * [Reading 0x Protocol Event Logs](en/event-read-0xprotocol/README.md)
* [Signatures](en/signatures/README.md)
  * [Generating Signatures](en/signature-generate/README.md)
  * [Verifying Signatures](en/signature-verify/README.md)
* [Testing](en/test/README.md)
  * [Faucets](en/faucets/README.md)
  * [Using a Simulated Client](en/client-simulated/README.md)
* [Swarm](en/swarm/README.md)
  * [Setting Up Swarm](en/swarm-setup/README.md)
  * [Uploading Files to Swarm](en/swarm-upload/README.md)
  * [Download Files From Swarm](en/swarm-download/README.md)
* [Whisper](en/whisper/README.md)
  * [Connecting Whisper Client](en/whisper-client/README.md)
  * [Generating Whisper Key Pair](en/whisper-keys/README.md)
  * [Sending Messages on Whisper](en/whisper-send/README.md)
  * [Subscribing to Whisper Messages](en/whisper-subscribe/README.md)
* [Utilities](en/util/README.md)
  * [Collection of Utility Functions](en/util-go/README.md)
* [Glossary](en/GLOSSARY.md)
* [Resources](en/resources/README.md)

## Help & Support

- Join the [#ethereum](https://gophers.slack.com/messages/C9HP1S9V2/) channel on the [gophers slack](https://invite.slack.golangbridge.org/) for Go (golang) help

- The [Ethereum StackExchange](https://ethereum.stackexchange.com/) is a great place to ask general Ethereum question and Go specific questions

## Development

Install dependencies:

```bash
make install
```

Run gitbook server:

```bash
make serve
```

Generating e-book in pdf, mobi, and epub format:

```bash
make ebooks
```

Visit [http://localhost:4000](http://localhost:4000)

## Contributing

Pull requests are welcome!

If making general content fixes:

- please double check for typos and cite any relevant sources in the comments.

If updating code examples:

- make sure to update both the code in the markdown files as well as the code in the [code](code/) folder.

If wanting to add a new translation, follow these instructions:

1. Set up [development environment](#development)

2. Add language to `LANGS.md`

3. Copy the the `en` directory and rename it with the 2 letter language code of the language you're translating to (e.g. `zh`)

4. Translate content

5. Set `"root"` to `"./"` in `book.json` if not already set

## Thanks

Thanks to [@qbig](https://github.com/qbig) and [@gzuhlwang](https://github.com/gzuhlwang) for the Chinese translation.

And thanks to all the [contributors](https://github.com/miguelmota/ethereum-development-with-go-book/graphs/contributors) who have contributed to this guide book.

## License

[CC0-1.0](./LICENSE)
