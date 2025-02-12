package resolution

import (
	"net/http"
	"strconv"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/unstoppabledomains/resolution-go/v2/uns/contracts/proxyreader"
)

// UnsBuilder is a builder to setup and build instance of Uns
type UnsBuilder interface {
	// SetContractBackend set Ethereum backend for communication with UNS registry
	SetContractBackend(backend bind.ContractBackend) UnsBuilder

	// SetL2ContractBackend set Ethereum backend for communication with UNS L2registry
	SetL2ContractBackend(backend bind.ContractBackend) UnsBuilder

	// SetMetadataClient set http backend for communication with ERC721 metadata server
	SetMetadataClient(backend MetadataClient) UnsBuilder

	// SetEthereumNetwork set Ethereum network for communication with UNS registry
	SetEthereumNetwork(network string) UnsBuilder

	// SetL2EthereumNetwork set Ethereum network for communication with UNS L2 registry
	SetL2EthereumNetwork(network string) UnsBuilder

	// Build Uns instance
	Build() (*Uns, error)
}

type unsBuilder struct {
	l1ContractBackend bind.ContractBackend
	l2ContractBackend bind.ContractBackend
	metadataClient    MetadataClient
	l1Network         string
	l2Network         string
}

// NewUnsBuilder Creates builder to setup new instance of Uns
func NewUnsBuilder() UnsBuilder {
	return &unsBuilder{
		l1Network: "mainnet",
		l2Network: "polygon",
	}
}

// SetContractBackend set Ethereum backend for communication with UNS registry
func (cb *unsBuilder) SetContractBackend(backend bind.ContractBackend) UnsBuilder {
	cb.l1ContractBackend = backend
	return cb
}

// SetContractBackend set Ethereum backend for communication with UNS registry
func (cb *unsBuilder) SetL2ContractBackend(backend bind.ContractBackend) UnsBuilder {
	cb.l2ContractBackend = backend
	return cb
}

func (cb *unsBuilder) SetMetadataClient(client MetadataClient) UnsBuilder {
	cb.metadataClient = client
	return cb
}

func (cb *unsBuilder) SetEthereumNetwork(network string) UnsBuilder {
	cb.l1Network = network
	return cb
}

func (cb *unsBuilder) SetL2EthereumNetwork(network string) UnsBuilder {
	cb.l2Network = network
	return cb
}

func (cb *unsBuilder) BuildService(netContracts contracts, contractBackend bind.ContractBackend, provider string) (*UnsService, error) {
	unsProxyReader := common.HexToAddress(netContracts["ProxyReader"].Address)
	cnsDefaultResolver := common.HexToAddress(netContracts["Resolver"].Address)
	unsRegistry := common.HexToAddress(netContracts["UNSRegistry"].Address)
	unsStartingEventsBlock, _ := strconv.ParseUint(netContracts["UNSRegistry"].DeploymentBlock[2:], 16, 32)
	cnsStartingEventsBlock, _ := strconv.ParseUint(netContracts["Resolver"].DeploymentBlock[2:], 16, 32)

	if contractBackend == nil {
		backend, err := ethclient.Dial(provider)
		if err != nil {
			return nil, err
		}
		contractBackend = backend
	}

	proxyReaderContract, err := proxyreader.NewContract(unsProxyReader, contractBackend)
	if err != nil {
		return nil, err
	}
	if err != nil {
		return nil, err
	}

	supportedKeys, err := newSupportedKeys()
	if err != nil {
		return nil, err
	}

	return &UnsService{proxyReader: proxyReaderContract, supportedKeys: supportedKeys, contractBackend: contractBackend, metadataClient: cb.metadataClient, cnsDefaultResolver: cnsDefaultResolver, unsRegistry: unsRegistry, unsStartingEventsBlock: unsStartingEventsBlock, cnsStartingEventsBlock: cnsStartingEventsBlock}, nil
}

// Build Uns instance
func (cb *unsBuilder) Build() (*Uns, error) {
	contracts, err := newContracts()
	if err != nil {
		return nil, err
	}

	if cb.metadataClient == nil {
		cb.metadataClient = &http.Client{}
	}

	if cb.l1Network == "" {
		return nil, &UnsConfigurationError{Layer: Layer1, InvalidField: "network"}
	}
	if cb.l2Network == "" {
		return nil, &UnsConfigurationError{Layer: Layer2, InvalidField: "network"}
	}
	l1Custom, l2Custom := cb.l1ContractBackend != nil, cb.l2ContractBackend != nil
	if l1Custom != l2Custom {
		if l1Custom {
			return nil, &UnsConfigurationError{Layer: Layer2, InvalidField: "contractBackend"}
		} else {
			return nil, &UnsConfigurationError{Layer: Layer1, InvalidField: "contractBackend"}
		}
	}

	l1Service, err := cb.BuildService(contracts[cb.l1Network], cb.l1ContractBackend, NetworkProviders[cb.l1Network])
	if err != nil {
		return nil, err
	}
	l1Service.networkId = NetworkNameToId[cb.l1Network]
	l1Service.blockchainProviderUrl = NetworkProviders[cb.l1Network]
	l1Service.Layer = Layer1

	l2Service, err := cb.BuildService(contracts[cb.l2Network], cb.l2ContractBackend, NetworkProviders[cb.l2Network])
	if err != nil {
		return nil, err
	}
	l2Service.Layer = Layer2
	l2Service.networkId = NetworkNameToId[cb.l2Network]
	l2Service.blockchainProviderUrl = NetworkProviders[cb.l2Network]

	zService, err := NewZnsBuilder().Build()
	if err != nil {
		return nil, err
	}

	return &Uns{
		*l1Service,
		*l2Service,
		*zService,
	}, nil
}
