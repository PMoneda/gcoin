package transaction

type Transaction struct {
	TxID          string
	Address       string
	ScriptPubKey  string
	Amount        float64
	Confirmations int
	Spendable     bool
	Solvable      bool
	Timestamp     int
}
