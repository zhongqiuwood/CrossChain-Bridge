package tools

import (
	"bytes"

	"github.com/fsn-dev/fsn-go-sdk/efsn/common"
)

var mbtcExtendedCodeParts = map[string][]byte{
	// Mbtc extended interfaces
	"SwapinFuncHash":  common.FromHex("0xec126c77"),
	"LogSwapinTopic":  common.FromHex("0x05d0634fe981be85c22e2942a880821b70095d84e152c3ea3c17a4e4250d9d61"),
	"SwapoutFuncHash": common.FromHex("0xad54056d"),
	"LogSwapoutTopic": common.FromHex("0x9c92ad817e5474d30a4378deface765150479363a897b0590fbb12ae9d89396b"),
}

var erc20CodeParts = map[string][]byte{
	// Erc20 interfaces
	"name":         common.FromHex("0x06fdde03"),
	"symbol":       common.FromHex("0x95d89b41"),
	"decimals":     common.FromHex("0x313ce567"),
	"totalSupply":  common.FromHex("0x18160ddd"),
	"balanceOf":    common.FromHex("0x70a08231"),
	"transfer":     common.FromHex("0xa9059cbb"),
	"transferFrom": common.FromHex("0x23b872dd"),
	"approve":      common.FromHex("0x095ea7b3"),
	"allowance":    common.FromHex("0xdd62ed3e"),
	"LogTransfer":  common.FromHex("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
	"LogApproval":  common.FromHex("0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925"),
}

func CheckContractCode(code []byte, codePartsSlice ...map[string][]byte) bool {
	for _, codeParts := range codePartsSlice {
		for _, part := range codeParts {
			if bytes.Index(code, part) == -1 {
				return false
			}
		}
	}
	return true
}

func IsErc20Contract(code []byte) bool {
	return CheckContractCode(code, erc20CodeParts)
}

func IsMbtcContract(code []byte) bool {
	return CheckContractCode(code, mbtcExtendedCodeParts, erc20CodeParts)
}
