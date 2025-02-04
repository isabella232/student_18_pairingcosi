package byzcoin_ng_test

/*
The test-file should at the very least run the protocol for a varying number
of nodes. It is even better practice to test the different methods of the
protocol, as in Test Driven Development.
*/

import (
	"testing"
	"time"

	"bls-ftcosi/cothority/log"
	"bls-ftcosi/cothority/network"
	"bls-ftcosi/cothority/protocols/template"
	"bls-ftcosi/cothority/sda"
)

func TestMain(m *testing.M) {
	log.MainTest(m)
}

// Tests a 2, 5 and 13-node system. It is good practice to test different
// sizes of trees to make sure your protocol is stable.
func TestNode(t *testing.T) {
	local := sda.NewLocalTest()
	nodes := []int{2, 5, 13}
	for _, nbrNodes := range nodes {
		_, _, tree := local.GenTree(nbrNodes, false, true, true)
		log.Lvl3(tree.Dump())

		pi, err := local.StartProtocol("Template", tree)
		if err != nil {
			t.Fatal("Couldn't start protocol:", err)
		}
		protocol := pi.(*template.ProtocolTemplate)
		timeout := network.WaitRetry * time.Duration(network.MaxRetryConnect*nbrNodes*2) * time.Millisecond
		select {
		case children := <-protocol.ChildCount:
			log.Lvl2("Instance 1 is done")
			if children != nbrNodes {
				t.Fatal("Didn't get a child-cound of", nbrNodes)
			}
		case <-time.After(timeout):
			t.Fatal("Didn't finish in time")
		}
		local.CloseAll()
	}
}
