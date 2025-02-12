package resolution

import (
	"net/http"
	"testing"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/assert"
)

func TestUnsBuilder(t *testing.T) {
	t.Parallel()
	builder := NewUnsBuilder().SetEthereumNetwork("goerli").SetL2EthereumNetwork("mumbai")
	_, err := builder.Build()
	assert.Nil(t, err)
	assert.NotNil(t, uns.l1Service.contractBackend)
	assert.NotNil(t, uns.l1Service.metadataClient)
	assert.NotNil(t, uns.l1Service.supportedKeys)
	assert.NotNil(t, uns.l1Service.proxyReader)

	assert.NotNil(t, uns.l2Service.contractBackend)
	assert.NotNil(t, uns.l2Service.metadataClient)
	assert.NotNil(t, uns.l2Service.supportedKeys)
	assert.NotNil(t, uns.l2Service.proxyReader)
}

func TestUnsBuilderSetBackend(t *testing.T) {
	t.Parallel()
	backendl1, _ := ethclient.Dial("https://goerli.infura.io/v3/c5da69dfac9c4d9d96dd232580d4124e")
	backendl2, _ := ethclient.Dial("https://polygon-mumbai.infura.io/v3/c5da69dfac9c4d9d96dd232580d4124e")
	builder := NewUnsBuilder().SetEthereumNetwork("goerli").SetL2EthereumNetwork("mumbai")
	builder.SetContractBackend(backendl1)
	builder.SetL2ContractBackend(backendl2)
	uns, err := builder.Build()
	assert.Nil(t, err)
	assert.Equal(t, backendl1, uns.l1Service.contractBackend)
	assert.Equal(t, backendl2, uns.l2Service.contractBackend)
}

func TestUnsBuilderSetMetadataClient(t *testing.T) {
	t.Parallel()
	client := &http.Client{}
	builder := NewUnsBuilder().SetEthereumNetwork("goerli").SetL2EthereumNetwork("mumbai")
	builder.SetMetadataClient(client)
	uns, err := builder.Build()
	assert.Nil(t, err)
	assert.Equal(t, client, uns.l1Service.metadataClient)
	assert.Equal(t, client, uns.l2Service.metadataClient)
}

func TestUnsBuilderChecksL2ContractBackend(t *testing.T) {
	t.Parallel()
	var expectedError *UnsConfigurationError
	backendl1, _ := ethclient.Dial("https://goerli.infura.io/v3/c5da69dfac9c4d9d96dd232580d4124e")
	builder := NewUnsBuilder().SetEthereumNetwork("goerli").SetL2EthereumNetwork("mumbai")
	builder.SetContractBackend(backendl1)
	_, err := builder.Build()
	assert.ErrorAs(t, err, &expectedError)
	assert.Equal(t, "Invalid UNS configuration value of contractBackend for Layer 2", err.Error())
}

func TestUnsBuilderChecksL1ContractBackend(t *testing.T) {
	t.Parallel()
	var expectedError *UnsConfigurationError
	backendl2, _ := ethclient.Dial("https://goerli.infura.io/v3/c5da69dfac9c4d9d96dd232580d4124e")
	builder := NewUnsBuilder().SetEthereumNetwork("goerli").SetL2EthereumNetwork("mumbai")
	builder.SetL2ContractBackend(backendl2)
	_, err := builder.Build()
	assert.ErrorAs(t, err, &expectedError)
	assert.Equal(t, "Invalid UNS configuration value of contractBackend for Layer 1", err.Error())
}
