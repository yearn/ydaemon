package addresses

import (
	"reflect"

	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/logs"
)

func ToMixedcase(address any) common.MixedcaseAddress {
	valueType := reflect.TypeOf(address)
	if valueType.Kind() == reflect.String {
		return common.NewMixedcaseAddress(common.HexToAddress(address.(string)))
	}
	if valueType.String() == `common.Address` {
		return common.NewMixedcaseAddress(address.(common.Address))
	}
	if valueType.String() == `common.MixedcaseAddress` {
		return address.(common.MixedcaseAddress)
	}
	logs.Warning(`unknown`, valueType.String())
	return common.MixedcaseAddress{}
}

func ToAddress(address any) common.Address {
	valueType := reflect.TypeOf(address)
	if valueType.Kind() == reflect.String {
		return common.HexToAddress(address.(string))
	}
	if valueType.String() == `common.MixedcaseAddress` {
		addr := address.(common.MixedcaseAddress)
		return addr.Address()
	}
	if valueType.String() == `common.Address` {
		return address.(common.Address)
	}
	logs.Warning(`unknown`, valueType.String())
	return common.Address{}
}

func Equals(a, b any) bool {
	return ToAddress(a).Hex() == ToAddress(b).Hex()
}
