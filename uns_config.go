package resolution

import (
	"encoding/json"
)

// contracts struct of contracts
type contracts map[string]struct {
	Address         string
	Implementation  string
	LegacyAddresses []string
	DeploymentBlock string
}

const (
	Mainnet string = "mainnet"
	Rinkeby string = "rinkeby"
	Polygon string = "polygon"
	Mumbai  string = "mumbai"
)

const (
	Layer1 string = "Layer 1"
	Layer2 string = "Layer 2"
)

type NetworkContracts map[string]contracts

var NetworkProviders = map[string]string{
	Mainnet: "https://mainnet.infura.io/v3/c5da69dfac9c4d9d96dd232580d4124e",
	Rinkeby: "https://rinkeby.infura.io/v3/c5da69dfac9c4d9d96dd232580d4124e",
	Polygon: "https://polygon-mainnet.infura.io/v3/c5da69dfac9c4d9d96dd232580d4124e",
	Mumbai:  "https://polygon-mumbai.infura.io/v3/c5da69dfac9c4d9d96dd232580d4124e",
}

var NetworkNameToId = map[string]int{
	Mainnet: 1,
	Rinkeby: 4,
	Polygon: 137,
	Mumbai:  80001,
}

func parseContracts(data []byte) (contracts, error) {
	var contractsObject struct {
		Contracts contracts
	}
	err := json.Unmarshal(data, &contractsObject)
	if err != nil {
		return nil, err
	}
	return contractsObject.Contracts, nil
}

func newContracts() (NetworkContracts, error) {
	networks := make(NetworkContracts)
	var err error
	networks[Mainnet], err = parseContracts(unsMainnetConfigJSON)
	if err != nil {
		return nil, err
	}
	networks[Rinkeby], err = parseContracts(unsRinkebyConfigJSON)
	if err != nil {
		return nil, err
	}
	networks[Polygon], err = parseContracts(unsPolygonConfigJSON)
	if err != nil {
		return nil, err
	}
	networks[Mumbai], err = parseContracts(unsMumbaiConfigJSON)
	if err != nil {
		return nil, err
	}
	return networks, nil
}

var unsMainnetConfigJSON = []byte(`
{
	"contracts": {
		"UNSRegistry": {
			"address": "0x049aba7510f45BA5b64ea9E658E342F904DB358D",
			"implementation": "0x2351f8557Ec733F35a278C922F9DcAac32c687AF",
			"legacyAddresses": [],
			"deploymentBlock": "0xc2fede",
			"forwarder": "0x049aba7510f45BA5b64ea9E658E342F904DB358D"
		},
		"CNSRegistry": {
			"address": "0xD1E5b0FF1287aA9f9A268759062E4Ab08b9Dacbe",
			"legacyAddresses": [],
			"deploymentBlock": "0x8a958b",
			"forwarder": "0x97B0E89fC1B7eD4A8B237D9d8Fcce9b234f25A37"
		},
		"MintingManager": {
			"address": "0x2a7084870bB724175a3C96Da8FaA55128fa3E19D",
			"implementation": "0x1c776e8D286a35e8B4bc51388A77dD2044E5Fa7d",
			"legacyAddresses": [],
			"deploymentBlock": "0xc2fee0",
			"forwarder": "0xb970fbCF52cd8111c76c379D4f2FE12E7f8AE7fb"
		},
		"ProxyAdmin": {
			"address": "0xAA16DA78110D9A9742c760a1a064F28654Ab93de",
			"legacyAddresses": [],
			"deploymentBlock": "0xc2fedc"
		},
		"SignatureController": {
			"address": "0x82EF94294C95aD0930055f31e53A34509227c5f7",
			"legacyAddresses": [],
			"deploymentBlock": "0x8a95a6"
		},
		"MintingController": {
			"address": "0xb0EE56339C3253361730F50c08d3d7817ecD60Ca",
			"legacyAddresses": [],
			"deploymentBlock": "0x8a95aa",
			"deprecated": true
		},
		"WhitelistedMinter": {
			"address": "0xd3fF3377b0ceade1303dAF9Db04068ef8a650757",
			"legacyAddresses": [],
			"deploymentBlock": "0xa76ad3",
			"deprecated": true
		},
		"URIPrefixController": {
			"address": "0x09B091492759737C03da9dB7eDF1CD6BCC3A9d91",
			"legacyAddresses": [],
			"deploymentBlock": "0x8a95ae",
			"deprecated": true
		},
		"DomainZoneController": {
			"address": "0xeA70777e28E00E81f58b8921fC47F78B8a72eFE7",
			"legacyAddresses": [],
			"deploymentBlock": "0x98ca20",
			"deprecated": true
		},
		"Resolver": {
			"address": "0xb66DcE2DA6afAAa98F2013446dBCB0f4B0ab2842",
			"legacyAddresses": [
				"0xa1cac442be6673c49f8e74ffc7c4fd746f3cbd0d",
				"0x878bc2f3f717766ab69c0a5f9a6144931e61aed3"
			],
			"deploymentBlock": "0x960844",
			"forwarder": "0x92660E5F403679aB45e0C6DB7c35B5629d265fDd"
		},
		"ProxyReader": {
			"address": "0xc3C2BAB5e3e52DBF311b2aAcEf2e40344f19494E",
			"legacyAddresses": [
				"0xfEe4D4F0aDFF8D84c12170306507554bC7045878",
				"0xa6E7cEf2EDDEA66352Fd68E5915b60BDbb7309f5",
				"0x7ea9Ee21077F84339eDa9C80048ec6db678642B1"
			],
			"deploymentBlock": "0xca7ab2"
		},
		"TwitterValidationOperator": {
			"address": "0x2F659766E3D08561CA3408FbAba7C0749ab2c402",
			"legacyAddresses": [
				"0xbb486C6E9cF1faA86a6E3eAAFE2e5665C0507855"
			],
			"deploymentBlock": "0xc300b5"
		},
		"FreeMinter": {
			"address": "0x1fC985cAc641ED5846b631f96F35d9b48Bc3b834",
			"legacyAddresses": [],
			"deploymentBlock": "0xacc390",
			"deprecated": true
		}
	}
}`)

var unsRinkebyConfigJSON = []byte(`
{
	"contracts": {
		"UNSRegistry": {
			"address": "0x7fb83000B8eD59D3eAD22f0D584Df3a85fBC0086",
			"implementation": "0xc479D7A65243f7Eb1641F06a6C04E5F06cb5c4F7",
			"legacyAddresses": [],
			"deploymentBlock": "0x85e628",
			"forwarder": "0x7fb83000B8eD59D3eAD22f0D584Df3a85fBC0086"
		},
		"CNSRegistry": {
			"address": "0xAad76bea7CFEc82927239415BB18D2e93518ecBB",
			"legacyAddresses": [],
			"deploymentBlock": "0x7232bc",
			"forwarder": "0xdf5CC97216785398D5C77348e68fc9461108f85d"
		},
		"MintingManager": {
			"address": "0xdAAf99A920D31F4f5720e4667b12b24e54A03070",
			"implementation": "0x38Fa95a0AC0E59D6e2845eFADBc17aF0FF9c7089",
			"legacyAddresses": [],
			"deploymentBlock": "0x85e629",
			"forwarder": "0xfB13e29C4D31a48B4Cd61131Cf3b681416e11681"
		},
		"ProxyAdmin": {
			"address": "0xaf9815005A208d1460b6fC60B4f90B9f2185E88c",
			"legacyAddresses": [],
			"deploymentBlock": "0x85e627"
		},
		"SignatureController": {
			"address": "0x66a5e3e2C27B4ce4F46BBd975270BE154748D164",
			"legacyAddresses": [],
			"deploymentBlock": "0x7232be"
		},
		"MintingController": {
			"address": "0x51765307AeB3Df2E647014a2C501d5324212467c",
			"legacyAddresses": [],
			"deploymentBlock": "0x7232bf",
			"deprecated": true
		},
		"WhitelistedMinter": {
			"address": "0xbcB32f13f90978a9e059E8Cb40FaA9e6619d98e7",
			"legacyAddresses": [],
			"deploymentBlock": "0x7232c6",
			"deprecated": true
		},
		"URIPrefixController": {
			"address": "0xe1d2e4B9f0518CA5c803073C3dFa886470627237",
			"legacyAddresses": [],
			"deploymentBlock": "0x7232c0",
			"deprecated": true
		},
		"DomainZoneController": {
			"address": "0x6f8F96A566663C1d4fEe70edD37E9b62Fe39dE5D",
			"legacyAddresses": [],
			"deploymentBlock": "0x7232c2",
			"deprecated": true
		},
		"Resolver": {
			"address": "0x95AE1515367aa64C462c71e87157771165B1287A",
			"legacyAddresses": [],
			"deploymentBlock": "0x7232cf",
			"forwarder": "0xE172D8557d6F342b1b2976dE784F6Dff6ABC0a37"
		},
		"ProxyReader": {
			"address": "0xE6729D224D00b3dd4FC731C4Ee3274E35Da06578",
			"legacyAddresses": [
				"0x299974AeD8911bcbd2C61262605b89F591a53E83",
				"0x9F19473F6a98a715176291c930558E1954fd3D1e",
				"0x3A2e74CF832cbA3d77E72708d55370119E4323a6"
			],
			"deploymentBlock": "0x8dc79a"
		},
		"TwitterValidationOperator": {
			"address": "0x9ea4A63184ebE9CBA55CD1af473D98075Aa02b4C",
			"legacyAddresses": [
				"0x1CB337b3b208dc29a6AcE8d11Bb591b66c5Dd83d"
			],
			"deploymentBlock": "0x86935e"
		},
		"FreeMinter": {
			"address": "0x84214215904cDEbA9044ECf95F3eBF009185AAf4",
			"legacyAddresses": [],
			"deploymentBlock": "0x740d93",
			"deprecated": true
		}
	}
}`)

var unsPolygonConfigJSON = []byte(`
{
	"contracts": {
		"UNSRegistry": {
			"address": "0xa9a6A3626993D487d2Dbda3173cf58cA1a9D9e9f",
			"legacyAddresses": [],
			"deploymentBlock": "0x01272eb5",
			"implementation": "0x0a7b33E986f2c8BF2a16bdda6004d3FaFFC27695",
			"forwarder": "0xa9a6A3626993D487d2Dbda3173cf58cA1a9D9e9f"
		},
		"CNSRegistry": {
			"address": "0x0000000000000000000000000000000000000000",
			"legacyAddresses": [],
			"deploymentBlock": "0x0",
			"forwarder": "0x0000000000000000000000000000000000000000"
		},
		"MintingManager": {
			"address": "0x7be83293BeeDc9Eba1bd76c66A65F10F3efaeC26",
			"legacyAddresses": [],
			"deploymentBlock": "0x01272f41",
			"implementation": "0x10e91753eC98cd259A62085002C25E92C9dc8Aed",
			"forwarder": "0xC37d3c4326ab0E1D2b9D8b916bBdf5715f780fcF"
		},
		"ProxyAdmin": {
			"address": "0xe1D668052D52388F52b90f4d1798DB2b04bC3b88",
			"legacyAddresses": [],
			"deploymentBlock": "0x01272d15"
		},
		"SignatureController": {
			"address": "0x0000000000000000000000000000000000000000",
			"legacyAddresses": [],
			"deploymentBlock": "0x0"
		},
		"MintingController": {
			"address": "0x0000000000000000000000000000000000000000",
			"legacyAddresses": [],
			"deploymentBlock": "0x0",
			"deprecated": true
		},
		"WhitelistedMinter": {
			"address": "0x0000000000000000000000000000000000000000",
			"legacyAddresses": [],
			"deploymentBlock": "0x0",
			"deprecated": true
		},
		"URIPrefixController": {
			"address": "0x0000000000000000000000000000000000000000",
			"legacyAddresses": [],
			"deploymentBlock": "0x0",
			"deprecated": true
		},
		"DomainZoneController": {
			"address": "0x0000000000000000000000000000000000000000",
			"legacyAddresses": [],
			"deploymentBlock": "0x0",
			"deprecated": true
		},
		"Resolver": {
			"address": "0x0000000000000000000000000000000000000000",
			"legacyAddresses": [],
			"deploymentBlock": "0x0",
			"forwarder": "0x0000000000000000000000000000000000000000"
		},
		"ProxyReader": {
			"address": "0xA3f32c8cd786dc089Bd1fC175F2707223aeE5d00",
			"legacyAddresses": [],
			"deploymentBlock": "0x01273234"
		},
		"TwitterValidationOperator": {
			"address": "0x0000000000000000000000000000000000000000",
			"legacyAddresses": [],
			"deploymentBlock": "0x0"
		},
		"FreeMinter": {
			"address": "0x0000000000000000000000000000000000000000",
			"legacyAddresses": [],
			"deploymentBlock": "0x0",
			"deprecated": true
		}
	}
}`)

var unsMumbaiConfigJSON = []byte(`
{
	"contracts": {
		"UNSRegistry": {
			"address": "0x2a93C52E7B6E7054870758e15A1446E769EdfB93",
			"legacyAddresses": [],
			"deploymentBlock": "0x01213f43",
			"implementation": "0x6EBa8D5fD76a7C2A760BE0fB993F18FB54920010",
			"forwarder": "0x2a93C52E7B6E7054870758e15A1446E769EdfB93"
		},
		"CNSRegistry": {
			"address": "0x0000000000000000000000000000000000000000",
			"legacyAddresses": [],
			"deploymentBlock": "0x0",
			"forwarder": "0x0000000000000000000000000000000000000000"
		},
		"MintingManager": {
			"address": "0x428189346bb3CC52f031A1092fd47C919AC30A9f",
			"legacyAddresses": [],
			"deploymentBlock": "0x01213f4a",
			"implementation": "0x4f4c3a4B75346A546d309934726e7FfbdA13262D",
			"forwarder": "0xEf3a491A8750BEC2Dff5339CF6Df94436d432C4d"
		},
		"ProxyAdmin": {
			"address": "0x460d63117c7Ab1624b7474C45BF46eC6702f57ce",
			"legacyAddresses": [],
			"deploymentBlock": "0x01213b22"
		},
		"SignatureController": {
			"address": "0x0000000000000000000000000000000000000000",
			"legacyAddresses": [],
			"deploymentBlock": "0x0"
		},
		"MintingController": {
			"address": "0x0000000000000000000000000000000000000000",
			"legacyAddresses": [],
			"deploymentBlock": "0x0",
			"deprecated": true
		},
		"WhitelistedMinter": {
			"address": "0x0000000000000000000000000000000000000000",
			"legacyAddresses": [],
			"deploymentBlock": "0x0",
			"deprecated": true
		},
		"URIPrefixController": {
			"address": "0x0000000000000000000000000000000000000000",
			"legacyAddresses": [],
			"deploymentBlock": "0x0",
			"deprecated": true
		},
		"DomainZoneController": {
			"address": "0x0000000000000000000000000000000000000000",
			"legacyAddresses": [],
			"deploymentBlock": "0x0",
			"deprecated": true
		},
		"Resolver": {
			"address": "0x0000000000000000000000000000000000000000",
			"legacyAddresses": [],
			"deploymentBlock": "0x0",
			"forwarder": "0x0000000000000000000000000000000000000000"
		},
		"ProxyReader": {
			"address": "0x332A8191905fA8E6eeA7350B5799F225B8ed30a9",
			"legacyAddresses": [],
			"deploymentBlock": "0x01213f87"
		},
		"TwitterValidationOperator": {
			"address": "0x0000000000000000000000000000000000000000",
			"legacyAddresses": [],
			"deploymentBlock": "0x0"
		},
		"FreeMinter": {
			"address": "0x0000000000000000000000000000000000000000",
			"legacyAddresses": [],
			"deploymentBlock": "0x0",
			"deprecated": true
		}
	}
}`)