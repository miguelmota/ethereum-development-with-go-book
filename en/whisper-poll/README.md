FilterMessages retrieves all messages that are received between the last call to this function and match the criteria that where given when the filter was created.

NewMessageFilter creates a filter within the node. This filter can be used to poll for new messages (see FilterMessages) that satisfy the given criteria. A filter can timeout when it was polled for in whisper.filterTimeout.
