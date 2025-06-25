package ethclient

type TxPoolResponse struct {
	Pending map[string]TxPoolTxn `json:"pending,omitempty"`
	Queued  map[string]TxPoolTxn `json:"queued,omitempty"`
}

type TxPoolTxn struct {
	BlockHash        string `json:"blockHash,omitempty"`
	BlockNumber      string `json:"blockNumber,omitempty"`
	From             string `json:"from,omitempty"`
	Gas              string `json:"gas,omitempty"`
	GasPrice         string `json:"gasPrice,omitempty"`
	Hash             string `json:"hash,omitempty"`
	Input            string `json:"input,omitempty"`
	Nonce            string `json:"nonce,omitempty"`
	To               string `json:"to,omitempty"`
	TransactionIndex string `json:"transactionIndex,omitempty"`
	Value            string `json:"value,omitempty"`
	Type             string `json:"type,omitempty"`
	ChainId          string `json:"chainId,omitempty"`

	V string `json:"v,omitempty"`
	R string `json:"r,omitempty"`
	S string `json:"s,omitempty"`
}
