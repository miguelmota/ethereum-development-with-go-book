# Smart Contract Compilation & ABI

In order to interact with a smart contract, we first must generate the ABI (application binary interface) of the contract and compile the ABI to a format that we can import into our Go application.

The first step is to install the [Solidity compiler](https://solidity.readthedocs.io/en/latest/installing-solidity.html) (`solc`).

Solc is available as a snapcraft package for Ubuntu.

```bash
sudo snap install solc --edge
```

Solc is available as a Homebrew package for macOS.

```bash
brew update
brew tap ethereum/ethereum
brew install solidity
```

For other platforms or for installing from source, check out the official solidity [install guide](https://solidity.readthedocs.io/en/latest/installing-solidity.html#building-from-source).

We also need to install a tool called `abigen` for generating the ABI from a solidity smart contract.

Assuming you have Go all set up on your computer, simply run the following to install the `abigen` tool.

```bash
go get -u github.com/ethereum/go-ethereum
cd $GOPATH/src/github.com/ethereum/go-ethereum/
make
make devtools
```

We'll create a simple smart contract to test with. This simple contract will be a key/value store with only 1 external method to set a key/value pair by anyone. We also added an event to emit after the value is set.

```solidity
pragma solidity ^0.4.24;

contract Store {
  event ItemSet(bytes32 key, bytes32 value);

  string public version;
  mapping (bytes32 => bytes32) public items;

  constructor(string _version) public {
    version = _version;
  }

  function setItem(bytes32 key, bytes32 value) external {
    items[key] = value;
    emit ItemSet(key, value);
  }
}
```

Now we can generate the ABI from a solidity source file.

```bash
solc --abi Store.sol
```

We'll store it in a file.

```bash
solc --abi Store.sol | awk '/JSON ABI/{x=1;next}x' > Store.abi
```

Now let's convert the ABI to a Go file that we can import. This new file will contain all the available methods the we can use to interact with the smart contract from our Go application.

```bash
abigen --abi=Store.abi --pkg=store --out=Store.go
```

In order to deploy a smart contract from Go, we also need to compile the solidity smart contract to EVM bytecode. The EVM bytecode is what will be sent in the data field of the transaction. The bin file is required for generating the deploy methods on the Go contract file.

```bash
solc --bin Store.sol | awk '/Binary:/{x=1;next}x' > Store.bin
```

Now we compile the Go contract file which will include the deploy methods because we includes the bin file.

```bash
abigen --bin=Store.bin --abi=Store.abi --pkg=store --out=Store.go
```

That's it for this lesson. In the next lessons we'll learn how to deploy the smart contract, and then interact with it.


**Full code**

Commands

```bash
solc --abi Store.sol | awk '/JSON ABI/{x=1;next}x' > Store.abi
solc --bin Store.sol | awk '/Binary:/{x=1;next}x' > Store.bin
abigen --bin=Store.bin --abi=Store.abi --pkg=store --out=Store.go
```

[Store.sol](https://github.com/miguelmota/ethereum-development-with-go-book/blob/master/code/contracts/Store.sol)

```solidity
pragma solidity ^0.4.24;

contract Store {
  event ItemSet(bytes32 key, bytes32 value);

  string public version;
  mapping (bytes32 => bytes32) public items;

  constructor(string _version) public {
    version = _version;
  }

  function setItem(bytes32 key, bytes32 value) external {
    items[key] = value;
    emit ItemSet(key, value);
  }
}
```
