# ListAccounts query

ListAccounts keeps cached list of accounts, pupulated by subscribing to event bus - see: projector.go.
The cache is populated from event store on startup.