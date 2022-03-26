package transaction

//GetInputsFromWallet with sufficient funds to transfer
func GetInputsFromWallet(wallet string, amount float64) []*TxDirection {
	input := make([]*TxDirection, 0)
	//TODO query on block chain
	input = append(input, &TxDirection{Address: wallet, Amount: 1, LockedTo: ""})
	input = append(input, &TxDirection{Address: wallet, Amount: 3, LockedTo: "a"})
	input = append(input, &TxDirection{Address: wallet, Amount: 1, LockedTo: ""})
	input = append(input, &TxDirection{Address: wallet, Amount: 2, LockedTo: "c"})
	input = append(input, &TxDirection{Address: wallet, Amount: 5, LockedTo: ""})
	input = append(input, &TxDirection{Address: wallet, Amount: 2, LockedTo: "d"})
	input = append(input, &TxDirection{Address: wallet, Amount: 10, LockedTo: ""})
	input = append(input, &TxDirection{Address: wallet, Amount: 0.5, LockedTo: ""})
	input = append(input, &TxDirection{Address: wallet, Amount: 20, LockedTo: ""})
	input = append(input, &TxDirection{Address: wallet, Amount: 1, LockedTo: ""})
	input = append(input, &TxDirection{Address: wallet, Amount: 3, LockedTo: "a"})
	input = append(input, &TxDirection{Address: wallet, Amount: 1, LockedTo: ""})
	input = append(input, &TxDirection{Address: wallet, Amount: 2, LockedTo: "c"})
	input = append(input, &TxDirection{Address: wallet, Amount: 5, LockedTo: ""})
	input = append(input, &TxDirection{Address: wallet, Amount: 2, LockedTo: "d"})
	input = append(input, &TxDirection{Address: wallet, Amount: 10, LockedTo: ""})
	input = append(input, &TxDirection{Address: wallet, Amount: 0.5, LockedTo: ""})
	input = append(input, &TxDirection{Address: wallet, Amount: 20, LockedTo: ""})
	input = append(input, &TxDirection{Address: wallet, Amount: 1, LockedTo: ""})
	input = append(input, &TxDirection{Address: wallet, Amount: 3, LockedTo: "a"})
	input = append(input, &TxDirection{Address: wallet, Amount: 1, LockedTo: ""})
	input = append(input, &TxDirection{Address: wallet, Amount: 2, LockedTo: "c"})
	input = append(input, &TxDirection{Address: wallet, Amount: 5, LockedTo: ""})
	input = append(input, &TxDirection{Address: wallet, Amount: 2, LockedTo: "d"})
	input = append(input, &TxDirection{Address: wallet, Amount: 10, LockedTo: ""})
	input = append(input, &TxDirection{Address: wallet, Amount: 0.5, LockedTo: ""})
	input = append(input, &TxDirection{Address: wallet, Amount: 20, LockedTo: ""})

	result := make([]*TxDirection, 0)
	currentAmount := 0.0
	for _, tx := range input {
		if currentAmount < amount && tx.LockedTo == "" {
			result = append(result, tx)
			currentAmount += tx.Amount
		}
	}
	return result
}
