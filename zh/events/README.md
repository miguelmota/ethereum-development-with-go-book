---
description: Tutorial on Ethereum event logs with Go.
---

# Events

Smart contracts have the ability to "emit" events during execution. Events are also known as "logs" in Ethereum. The output of the events are stored in transaction receipts under a logs section. Events have become pretty widely used in Ethereum smart contracts to log when a significant action has occured, particularly in token contracts (i.e. ERC-20) to indicate that a token transfer has occured. These sections will walk you through the process of reading events from the blockchain as well as subscribing to events so that you get notified in real time as the transaction gets mined.
