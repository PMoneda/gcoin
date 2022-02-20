package ledger

import (
	"fmt"
	"os"
	"testing"

	"github.com/PMoneda/gcoin/config"
)

func Test_ShouldCreateNewLedger(t *testing.T) {
	l, err := NewLedgerFromPath("ledger.db")
	if err != nil {
		t.Fail()
	}
	l.Close()

	ledger, err := OpenLedger("ledger.db")
	if ledger.powDifficult != config.PoW_Difficult {
		t.Fail()
	}
	cleanUp()

}

func Test_ShouldCreateBlocksWhenPreviousMatchWithHead(t *testing.T) {

	ledger, _ := OpenLedger("ledger.db")
	if ledger.powDifficult != config.PoW_Difficult {
		t.Fail()
	}
	for i := 4; i < 25; i++ {
		ledger.AppendBlock([]byte(fmt.Sprintf("%d", i)))
	}
	if err := ledger.CheckConsistency(); err != nil {
		fmt.Printf(err.Error())
		t.Fail()
	}

	ledger.Close()
	cleanUp()

}

func cleanUp() {
	os.Remove("ledger.db")
}
