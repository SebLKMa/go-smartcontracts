// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bankapi

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// BankapiMetaData contains all meta data concerning the Bankapi contract.
var BankapiMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"name\":\"EventLog\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"API\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"losers\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"anteWei\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gameFeeWei\",\"type\":\"uint256\"}],\"name\":\"Reconcile\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Withdraw\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f80fd5b506040518060400160405280600581526020017f302e312e30000000000000000000000000000000000000000000000000000000815250600190816100549190610294565b50610363565b5f81519050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f60028204905060018216806100d557607f821691505b6020821081036100e8576100e7610091565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f6008830261014a7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8261010f565b610154868361010f565b95508019841693508086168417925050509392505050565b5f819050919050565b5f819050919050565b5f61019861019361018e8461016c565b610175565b61016c565b9050919050565b5f819050919050565b6101b18361017e565b6101c56101bd8261019f565b84845461011b565b825550505050565b5f90565b6101d96101cd565b6101e48184846101a8565b505050565b5b81811015610207576101fc5f826101d1565b6001810190506101ea565b5050565b601f82111561024c5761021d816100ee565b61022684610100565b81016020851015610235578190505b61024961024185610100565b8301826101e9565b50505b505050565b5f82821c905092915050565b5f61026c5f1984600802610251565b1980831691505092915050565b5f610284838361025d565b9150826002028217905092915050565b61029d8261005a565b67ffffffffffffffff8111156102b6576102b5610064565b5b6102c082546100be565b6102cb82828561020b565b5f60209050601f8311600181146102fc575f84156102ea578287015190505b6102f48582610279565b86555061035b565b601f19841661030a866100ee565b5f5b828110156103315784890151825560018201915060208501945060208101905061030c565b8683101561034e578489015161034a601f89168261025d565b8355505b6001600288020188555050505b505050505050565b611a41806103705f395ff3fe608060405260043610610054575f3560e01c806357ea89b6146100585780637d7b009914610062578063b4a99a4e1461008c578063bb62860d146100b6578063ed21248c146100e0578063fa84fd8e146100ea575b5f80fd5b610060610112565b005b34801561006d575f80fd5b506100766102c8565b6040516100839190610da4565b60405180910390f35b348015610097575f80fd5b506100a06102eb565b6040516100ad9190610da4565b60405180910390f35b3480156100c1575f80fd5b506100ca610310565b6040516100d79190610e2d565b60405180910390f35b6100e861039c565b005b3480156100f5575f80fd5b50610110600480360381019061010b9190610ffb565b610496565b005b5f3390505f60035f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205403610195576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161018c906110c5565b60405180910390fd5b5f60035f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205490508173ffffffffffffffffffffffffffffffffffffffff166108fc8290811502906040515f60405180830381858888f19350505050158015610219573d5f803e3d5ffd5b505f60035f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20819055507fd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a610286336109ec565b61028f83610ba2565b6040516020016102a092919061118f565b6040516020818303038152906040526040516102bc9190610e2d565b60405180910390a15050565b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6001805461031d9061120c565b80601f01602080910402602001604051908101604052809291908181526020018280546103499061120c565b80156103945780601f1061036b57610100808354040283529160200191610394565b820191905f5260205f20905b81548152906001019060200180831161037757829003601f168201915b505050505081565b3460035f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8282546103e89190611269565b925050819055507fd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a610419336109ec565b61045f60035f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2054610ba2565b6040516020016104709291906112e8565b60405160208183030381529060405260405161048c9190610e2d565b60405180910390a1565b5f8290505f5b8451811015610712578360035f8784815181106104bc576104bb611338565b5b602002602001015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20541015610689577fd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a61058560035f88858151811061053e5761053d611338565b5b602002602001015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2054610ba2565b61058e86610ba2565b60405160200161059f9291906113b1565b6040516020818303038152906040526040516105bb9190610e2d565b60405180910390a160035f8683815181106105d9576105d8611338565b5b602002602001015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2054826106269190611269565b91505f60035f87848151811061063f5761063e611338565b5b602002602001015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2081905550610705565b83826106959190611269565b91508360035f8784815181106106ae576106ad611338565b5b602002602001015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8282546106fd91906113f2565b925050819055505b808060010191505061049c565b507fd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a61073d84610ba2565b61074684610ba2565b61074f84610ba2565b60405160200161076193929190611497565b60405160208183030381529060405260405161077d9190610e2d565b60405180910390a15f81036107c7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107be90611573565b60405180910390fd5b818110156108a7578060035f60025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f82825461083c9190611269565b925050819055507fd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a61086d82610ba2565b60405160200161087d9190611601565b6040516020818303038152906040526040516108999190610e2d565b60405180910390a1506109e6565b81816108b391906113f2565b90508060035f8773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8282546109019190611269565b925050819055508160035f60025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8282546109759190611269565b925050819055507fd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a6109a682610ba2565b6109af84610ba2565b6040516020016109c092919061167d565b6040516020818303038152906040526040516109dc9190610e2d565b60405180910390a1505b50505050565b60605f602867ffffffffffffffff811115610a0a57610a09610e8c565b5b6040519080825280601f01601f191660200182016040528015610a3c5781602001600182028036833780820191505090505b5090505f5b6014811015610b98575f816013610a5891906113f2565b6008610a6491906116cd565b6002610a70919061183d565b8573ffffffffffffffffffffffffffffffffffffffff16610a9191906118b4565b60f81b90505f60108260f81c610aa791906118f0565b60f81b90505f8160f81c6010610abd9190611920565b8360f81c610acb919061195c565b60f81b9050610ad982610d20565b85856002610ae791906116cd565b81518110610af857610af7611338565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191690815f1a905350610b2f81610d20565b856001866002610b3f91906116cd565b610b499190611269565b81518110610b5a57610b59611338565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191690815f1a9053505050508080600101915050610a41565b5080915050919050565b60605f8203610be8576040518060400160405280600181526020017f30000000000000000000000000000000000000000000000000000000000000008152509050610d1b565b5f8290505f5b5f8214610c17578080610c0090611990565b915050600a82610c1091906118b4565b9150610bee565b5f8167ffffffffffffffff811115610c3257610c31610e8c565b5b6040519080825280601f01601f191660200182016040528015610c645781602001600182028036833780820191505090505b5090505f8290505b5f8614610d1357600181610c8091906113f2565b90505f600a8088610c9191906118b4565b610c9b91906116cd565b87610ca691906113f2565b6030610cb291906119d7565b90505f8160f81b905080848481518110610ccf57610cce611338565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191690815f1a905350600a88610d0a91906118b4565b97505050610c6c565b819450505050505b919050565b5f600a8260f81c60ff161015610d4a5760308260f81c610d4091906119d7565b60f81b9050610d60565b60578260f81c610d5a91906119d7565b60f81b90505b919050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f610d8e82610d65565b9050919050565b610d9e81610d84565b82525050565b5f602082019050610db75f830184610d95565b92915050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f601f19601f8301169050919050565b5f610dff82610dbd565b610e098185610dc7565b9350610e19818560208601610dd7565b610e2281610de5565b840191505092915050565b5f6020820190508181035f830152610e458184610df5565b905092915050565b5f604051905090565b5f80fd5b5f80fd5b610e6781610d84565b8114610e71575f80fd5b50565b5f81359050610e8281610e5e565b92915050565b5f80fd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b610ec282610de5565b810181811067ffffffffffffffff82111715610ee157610ee0610e8c565b5b80604052505050565b5f610ef3610e4d565b9050610eff8282610eb9565b919050565b5f67ffffffffffffffff821115610f1e57610f1d610e8c565b5b602082029050602081019050919050565b5f80fd5b5f610f45610f4084610f04565b610eea565b90508083825260208201905060208402830185811115610f6857610f67610f2f565b5b835b81811015610f915780610f7d8882610e74565b845260208401935050602081019050610f6a565b5050509392505050565b5f82601f830112610faf57610fae610e88565b5b8135610fbf848260208601610f33565b91505092915050565b5f819050919050565b610fda81610fc8565b8114610fe4575f80fd5b50565b5f81359050610ff581610fd1565b92915050565b5f805f806080858703121561101357611012610e56565b5b5f61102087828801610e74565b945050602085013567ffffffffffffffff81111561104157611040610e5a565b5b61104d87828801610f9b565b935050604061105e87828801610fe7565b925050606061106f87828801610fe7565b91505092959194509250565b7f6e6f7420656e6f7567682062616c616e636500000000000000000000000000005f82015250565b5f6110af601283610dc7565b91506110ba8261107b565b602082019050919050565b5f6020820190508181035f8301526110dc816110a3565b9050919050565b7f77697468647261775b0000000000000000000000000000000000000000000000815250565b5f81905092915050565b5f61111d82610dbd565b6111278185611109565b9350611137818560208601610dd7565b80840191505092915050565b7f5d20616d6f756e745b0000000000000000000000000000000000000000000000815250565b7f5d00000000000000000000000000000000000000000000000000000000000000815250565b5f611199826110e3565b6009820191506111a98285611113565b91506111b482611143565b6009820191506111c48284611113565b91506111cf82611169565b6001820191508190509392505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061122357607f821691505b602082108103611236576112356111df565b5b50919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f61127382610fc8565b915061127e83610fc8565b92508282019050808211156112965761129561123c565b5b92915050565b7f6465706f7369745b000000000000000000000000000000000000000000000000815250565b7f5d2062616c616e63655b00000000000000000000000000000000000000000000815250565b5f6112f28261129c565b6008820191506113028285611113565b915061130d826112c2565b600a8201915061131d8284611113565b915061132882611169565b6001820191508190509392505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b7f6163636f756e742062616c616e63652000000000000000000000000000000000815250565b7f206973206c657373207468616e20616e74652000000000000000000000000000815250565b5f6113bb82611365565b6010820191506113cb8285611113565b91506113d68261138b565b6013820191506113e68284611113565b91508190509392505050565b5f6113fc82610fc8565b915061140783610fc8565b925082820390508181111561141f5761141e61123c565b5b92915050565b7f616e74655b000000000000000000000000000000000000000000000000000000815250565b7f5d2067616d654665655b00000000000000000000000000000000000000000000815250565b7f5d20706f745b0000000000000000000000000000000000000000000000000000815250565b5f6114a182611425565b6005820191506114b18286611113565b91506114bc8261144b565b600a820191506114cc8285611113565b91506114d782611471565b6006820191506114e78284611113565b91506114f282611169565b600182019150819050949350505050565b7f6e6f20706f74207761732063726561746564206261736564206f6e206163636f5f8201527f756e742062616c616e6365730000000000000000000000000000000000000000602082015250565b5f61155d602c83610dc7565b915061156882611503565b604082019050919050565b5f6020820190508181035f83015261158a81611551565b9050919050565b7f706f74206c657373207468616e206665653a2077696e6e65725b305d206f776e5f8201527f65725b0000000000000000000000000000000000000000000000000000000000602082015250565b5f6115eb602383611109565b91506115f682611591565b602382019050919050565b5f61160b826115df565b91506116178284611113565b915061162282611169565b60018201915081905092915050565b7f77696e6e65725b00000000000000000000000000000000000000000000000000815250565b7f5d206f776e65725b000000000000000000000000000000000000000000000000815250565b5f61168782611631565b6007820191506116978285611113565b91506116a282611657565b6008820191506116b28284611113565b91506116bd82611169565b6001820191508190509392505050565b5f6116d782610fc8565b91506116e283610fc8565b92508282026116f081610fc8565b915082820484148315176117075761170661123c565b5b5092915050565b5f8160011c9050919050565b5f808291508390505b60018511156117635780860481111561173f5761173e61123c565b5b600185161561174e5780820291505b808102905061175c8561170e565b9450611723565b94509492505050565b5f8261177b5760019050611836565b81611788575f9050611836565b816001811461179e57600281146117a8576117d7565b6001915050611836565b60ff8411156117ba576117b961123c565b5b8360020a9150848211156117d1576117d061123c565b5b50611836565b5060208310610133831016604e8410600b841016171561180c5782820a9050838111156118075761180661123c565b5b611836565b611819848484600161171a565b925090508184048111156118305761182f61123c565b5b81810290505b9392505050565b5f61184782610fc8565b915061185283610fc8565b925061187f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff848461176c565b905092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b5f6118be82610fc8565b91506118c983610fc8565b9250826118d9576118d8611887565b5b828204905092915050565b5f60ff82169050919050565b5f6118fa826118e4565b9150611905836118e4565b92508261191557611914611887565b5b828204905092915050565b5f61192a826118e4565b9150611935836118e4565b9250828202611943816118e4565b91508082146119555761195461123c565b5b5092915050565b5f611966826118e4565b9150611971836118e4565b9250828203905060ff81111561198a5761198961123c565b5b92915050565b5f61199a82610fc8565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036119cc576119cb61123c565b5b600182019050919050565b5f6119e1826118e4565b91506119ec836118e4565b9250828201905060ff811115611a0557611a0461123c565b5b9291505056fea264697066735822122003605204e904b0c97c96ae76c3e6f4710042d0cfe5b180147725486ef0e4b7eb64736f6c634300081a0033",
}

// BankapiABI is the input ABI used to generate the binding from.
// Deprecated: Use BankapiMetaData.ABI instead.
var BankapiABI = BankapiMetaData.ABI

// BankapiBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BankapiMetaData.Bin instead.
var BankapiBin = BankapiMetaData.Bin

// DeployBankapi deploys a new Ethereum contract, binding an instance of Bankapi to it.
func DeployBankapi(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Bankapi, error) {
	parsed, err := BankapiMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BankapiBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Bankapi{BankapiCaller: BankapiCaller{contract: contract}, BankapiTransactor: BankapiTransactor{contract: contract}, BankapiFilterer: BankapiFilterer{contract: contract}}, nil
}

// Bankapi is an auto generated Go binding around an Ethereum contract.
type Bankapi struct {
	BankapiCaller     // Read-only binding to the contract
	BankapiTransactor // Write-only binding to the contract
	BankapiFilterer   // Log filterer for contract events
}

// BankapiCaller is an auto generated read-only Go binding around an Ethereum contract.
type BankapiCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BankapiTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BankapiTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BankapiFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BankapiFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BankapiSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BankapiSession struct {
	Contract     *Bankapi          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BankapiCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BankapiCallerSession struct {
	Contract *BankapiCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// BankapiTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BankapiTransactorSession struct {
	Contract     *BankapiTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// BankapiRaw is an auto generated low-level Go binding around an Ethereum contract.
type BankapiRaw struct {
	Contract *Bankapi // Generic contract binding to access the raw methods on
}

// BankapiCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BankapiCallerRaw struct {
	Contract *BankapiCaller // Generic read-only contract binding to access the raw methods on
}

// BankapiTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BankapiTransactorRaw struct {
	Contract *BankapiTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBankapi creates a new instance of Bankapi, bound to a specific deployed contract.
func NewBankapi(address common.Address, backend bind.ContractBackend) (*Bankapi, error) {
	contract, err := bindBankapi(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Bankapi{BankapiCaller: BankapiCaller{contract: contract}, BankapiTransactor: BankapiTransactor{contract: contract}, BankapiFilterer: BankapiFilterer{contract: contract}}, nil
}

// NewBankapiCaller creates a new read-only instance of Bankapi, bound to a specific deployed contract.
func NewBankapiCaller(address common.Address, caller bind.ContractCaller) (*BankapiCaller, error) {
	contract, err := bindBankapi(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BankapiCaller{contract: contract}, nil
}

// NewBankapiTransactor creates a new write-only instance of Bankapi, bound to a specific deployed contract.
func NewBankapiTransactor(address common.Address, transactor bind.ContractTransactor) (*BankapiTransactor, error) {
	contract, err := bindBankapi(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BankapiTransactor{contract: contract}, nil
}

// NewBankapiFilterer creates a new log filterer instance of Bankapi, bound to a specific deployed contract.
func NewBankapiFilterer(address common.Address, filterer bind.ContractFilterer) (*BankapiFilterer, error) {
	contract, err := bindBankapi(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BankapiFilterer{contract: contract}, nil
}

// bindBankapi binds a generic wrapper to an already deployed contract.
func bindBankapi(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BankapiMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bankapi *BankapiRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bankapi.Contract.BankapiCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bankapi *BankapiRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bankapi.Contract.BankapiTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bankapi *BankapiRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bankapi.Contract.BankapiTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bankapi *BankapiCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bankapi.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bankapi *BankapiTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bankapi.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bankapi *BankapiTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bankapi.Contract.contract.Transact(opts, method, params...)
}

// API is a free data retrieval call binding the contract method 0x7d7b0099.
//
// Solidity: function API() view returns(address)
func (_Bankapi *BankapiCaller) API(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bankapi.contract.Call(opts, &out, "API")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// API is a free data retrieval call binding the contract method 0x7d7b0099.
//
// Solidity: function API() view returns(address)
func (_Bankapi *BankapiSession) API() (common.Address, error) {
	return _Bankapi.Contract.API(&_Bankapi.CallOpts)
}

// API is a free data retrieval call binding the contract method 0x7d7b0099.
//
// Solidity: function API() view returns(address)
func (_Bankapi *BankapiCallerSession) API() (common.Address, error) {
	return _Bankapi.Contract.API(&_Bankapi.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0xb4a99a4e.
//
// Solidity: function Owner() view returns(address)
func (_Bankapi *BankapiCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bankapi.contract.Call(opts, &out, "Owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0xb4a99a4e.
//
// Solidity: function Owner() view returns(address)
func (_Bankapi *BankapiSession) Owner() (common.Address, error) {
	return _Bankapi.Contract.Owner(&_Bankapi.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0xb4a99a4e.
//
// Solidity: function Owner() view returns(address)
func (_Bankapi *BankapiCallerSession) Owner() (common.Address, error) {
	return _Bankapi.Contract.Owner(&_Bankapi.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0xbb62860d.
//
// Solidity: function Version() view returns(string)
func (_Bankapi *BankapiCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Bankapi.contract.Call(opts, &out, "Version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0xbb62860d.
//
// Solidity: function Version() view returns(string)
func (_Bankapi *BankapiSession) Version() (string, error) {
	return _Bankapi.Contract.Version(&_Bankapi.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0xbb62860d.
//
// Solidity: function Version() view returns(string)
func (_Bankapi *BankapiCallerSession) Version() (string, error) {
	return _Bankapi.Contract.Version(&_Bankapi.CallOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xed21248c.
//
// Solidity: function Deposit() payable returns()
func (_Bankapi *BankapiTransactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bankapi.contract.Transact(opts, "Deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xed21248c.
//
// Solidity: function Deposit() payable returns()
func (_Bankapi *BankapiSession) Deposit() (*types.Transaction, error) {
	return _Bankapi.Contract.Deposit(&_Bankapi.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xed21248c.
//
// Solidity: function Deposit() payable returns()
func (_Bankapi *BankapiTransactorSession) Deposit() (*types.Transaction, error) {
	return _Bankapi.Contract.Deposit(&_Bankapi.TransactOpts)
}

// Reconcile is a paid mutator transaction binding the contract method 0xfa84fd8e.
//
// Solidity: function Reconcile(address winner, address[] losers, uint256 anteWei, uint256 gameFeeWei) returns()
func (_Bankapi *BankapiTransactor) Reconcile(opts *bind.TransactOpts, winner common.Address, losers []common.Address, anteWei *big.Int, gameFeeWei *big.Int) (*types.Transaction, error) {
	return _Bankapi.contract.Transact(opts, "Reconcile", winner, losers, anteWei, gameFeeWei)
}

// Reconcile is a paid mutator transaction binding the contract method 0xfa84fd8e.
//
// Solidity: function Reconcile(address winner, address[] losers, uint256 anteWei, uint256 gameFeeWei) returns()
func (_Bankapi *BankapiSession) Reconcile(winner common.Address, losers []common.Address, anteWei *big.Int, gameFeeWei *big.Int) (*types.Transaction, error) {
	return _Bankapi.Contract.Reconcile(&_Bankapi.TransactOpts, winner, losers, anteWei, gameFeeWei)
}

// Reconcile is a paid mutator transaction binding the contract method 0xfa84fd8e.
//
// Solidity: function Reconcile(address winner, address[] losers, uint256 anteWei, uint256 gameFeeWei) returns()
func (_Bankapi *BankapiTransactorSession) Reconcile(winner common.Address, losers []common.Address, anteWei *big.Int, gameFeeWei *big.Int) (*types.Transaction, error) {
	return _Bankapi.Contract.Reconcile(&_Bankapi.TransactOpts, winner, losers, anteWei, gameFeeWei)
}

// Withdraw is a paid mutator transaction binding the contract method 0x57ea89b6.
//
// Solidity: function Withdraw() payable returns()
func (_Bankapi *BankapiTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bankapi.contract.Transact(opts, "Withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x57ea89b6.
//
// Solidity: function Withdraw() payable returns()
func (_Bankapi *BankapiSession) Withdraw() (*types.Transaction, error) {
	return _Bankapi.Contract.Withdraw(&_Bankapi.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x57ea89b6.
//
// Solidity: function Withdraw() payable returns()
func (_Bankapi *BankapiTransactorSession) Withdraw() (*types.Transaction, error) {
	return _Bankapi.Contract.Withdraw(&_Bankapi.TransactOpts)
}

// BankapiEventLogIterator is returned from FilterEventLog and is used to iterate over the raw logs and unpacked data for EventLog events raised by the Bankapi contract.
type BankapiEventLogIterator struct {
	Event *BankapiEventLog // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BankapiEventLogIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BankapiEventLog)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BankapiEventLog)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BankapiEventLogIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BankapiEventLogIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BankapiEventLog represents a EventLog event raised by the Bankapi contract.
type BankapiEventLog struct {
	Value string
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterEventLog is a free log retrieval operation binding the contract event 0xd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a.
//
// Solidity: event EventLog(string value)
func (_Bankapi *BankapiFilterer) FilterEventLog(opts *bind.FilterOpts) (*BankapiEventLogIterator, error) {

	logs, sub, err := _Bankapi.contract.FilterLogs(opts, "EventLog")
	if err != nil {
		return nil, err
	}
	return &BankapiEventLogIterator{contract: _Bankapi.contract, event: "EventLog", logs: logs, sub: sub}, nil
}

// WatchEventLog is a free log subscription operation binding the contract event 0xd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a.
//
// Solidity: event EventLog(string value)
func (_Bankapi *BankapiFilterer) WatchEventLog(opts *bind.WatchOpts, sink chan<- *BankapiEventLog) (event.Subscription, error) {

	logs, sub, err := _Bankapi.contract.WatchLogs(opts, "EventLog")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BankapiEventLog)
				if err := _Bankapi.contract.UnpackLog(event, "EventLog", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseEventLog is a log parse operation binding the contract event 0xd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a.
//
// Solidity: event EventLog(string value)
func (_Bankapi *BankapiFilterer) ParseEventLog(log types.Log) (*BankapiEventLog, error) {
	event := new(BankapiEventLog)
	if err := _Bankapi.contract.UnpackLog(event, "EventLog", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
