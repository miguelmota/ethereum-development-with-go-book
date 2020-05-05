---
description: Tutorial on setting up swarm node.
---

# Setting up Swarm

To run swarm you first need to install `geth` and `bzzd` which is the swarm daemon.

```go
go install github.com/ethereum/go-ethereum/cmd/geth
go install github.com/ethersphere/swarm/cmd/swarm
```

Now we'll generate a new geth account.

```bash
$ geth account new

Your new account is locked with a password. Please give a password. Do not forget this password.
Passphrase:
Repeat passphrase:
Address: {970ef9790b54425bea2c02e25cab01e48cf92573}
```

Export the environment variable `BZZKEY` mapping to the geth account address we just generated.

```bash
export BZZKEY=970ef9790b54425bea2c02e25cab01e48cf92573
```

And now run swarm with the specified account to be our swarm account. Swarm by default will run on port `8500`.

```bash
$ swarm --bzzaccount $BZZKEY
Unlocking swarm account 0x970EF9790B54425BEA2C02e25cAb01E48CF92573 [1/3]
Passphrase:
WARN [06-12|13:11:41] Starting Swarm service
```

Now that we have the swarm daemon set up and running, let's learn how to upload files to swarm in the [next section](../swarm-upload).

---

### Full code

Commands

```bash
go install github.com/ethereum/go-ethereum/cmd/geth
go install github.com/ethersphere/swarm/cmd/swarm
geth account new
export BZZKEY=970ef9790b54425bea2c02e25cab01e48cf92573
swarm --bzzaccount $BZZKEY
```
