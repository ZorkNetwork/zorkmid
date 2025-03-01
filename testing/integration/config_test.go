package integration

import (
	"io/ioutil"
	"testing"
	"time"

	"github.com/kaspanet/kaspad/domain/dagconfig"
	"github.com/kaspanet/kaspad/infrastructure/config"
)

const (
	p2pAddress1 = "127.0.0.1:54321"
	p2pAddress2 = "127.0.0.1:54322"
	p2pAddress3 = "127.0.0.1:54323"
	p2pAddress4 = "127.0.0.1:54324"
	p2pAddress5 = "127.0.0.1:54325"

	rpcAddress1 = "127.0.0.1:12345"
	rpcAddress2 = "127.0.0.1:12346"
	rpcAddress3 = "127.0.0.1:12347"
	rpcAddress4 = "127.0.0.1:12348"
	rpcAddress5 = "127.0.0.1:12349"

	miningAddress1           = "zorksim:qqrut86v094tl3gmzd5dna694p2gzplfwfgl0g7nt5vcnka9t3d5clczq93y2"
	miningAddress1PrivateKey = "f74c54eaea70d6fddf3ebfa6a74c571b9efdb152aa25493acabd9d768ecd529d"

	miningAddress2           = "zorksim:qqrut86v094tl3gmzd5dna694p2gzplfwfgl0g7nt5vcnka9t3d5clczq93y2"
	miningAddress2PrivateKey = "f74c54eaea70d6fddf3ebfa6a74c571b9efdb152aa25493acabd9d768ecd529d"

	miningAddress3           = "zorksim:qznryhlsz2w3f8rfrja3qm5gsw99x4ry4hfalpurf28dxez9vkr3jc6tffmh2"
	miningAddress3PrivateKey = "f42805a8e868d8ca1db1a6bea4ddb00d86b6355c1a4b4c8ea147542da924b4f9"

	defaultTimeout = 30 * time.Second
)

func setConfig(t *testing.T, harness *appHarness, protocolVersion uint32) {
	harness.config = commonConfig()
	harness.config.AppDir = randomDirectory(t)
	harness.config.Listeners = []string{harness.p2pAddress}
	harness.config.RPCListeners = []string{harness.rpcAddress}
	harness.config.UTXOIndex = harness.utxoIndex
	harness.config.AllowSubmitBlockWhenNotSynced = true
	if protocolVersion != 0 {
		harness.config.ProtocolVersion = protocolVersion
	}

	if harness.overrideDAGParams != nil {
		harness.config.ActiveNetParams = harness.overrideDAGParams
	}
}

func commonConfig() *config.Config {
	commonConfig := config.DefaultConfig()

	*commonConfig.ActiveNetParams = dagconfig.SimnetParams // Copy so that we can make changes safely
	commonConfig.ActiveNetParams.BlockCoinbaseMaturity = 10
	commonConfig.TargetOutboundPeers = 0
	commonConfig.DisableDNSSeed = true
	commonConfig.Simnet = true

	return commonConfig
}

func randomDirectory(t *testing.T) string {
	dir, err := ioutil.TempDir("", "integration-test")
	if err != nil {
		t.Fatalf("Error creating temporary directory for test: %+v", err)
	}

	return dir
}
