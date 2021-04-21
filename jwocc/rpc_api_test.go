package jwocc

import (
	"github.com/blocktree/whitecoin-adapter/types"
	"testing"

	"github.com/blocktree/openwallet/v2/log"
)

func TestWalletClient_GetBlockchainInfo(t *testing.T) {
	b, err := tw.Api.GetBlockchainInfo()
	if err != nil {
		t.Errorf("GetBlockchainInfo failed unexpected error: %v\n", err)
	} else {
		log.Infof("GetBlockchainInfo info: %+v\n", b)
	}
}

func TestWalletClient_GetBlockByHeight(t *testing.T) {
	block, err := tw.Api.GetBlockByHeight(48460)
	if err != nil {
		t.Errorf("GetBlockByHeight failed unexpected error: %v\n", err)
		return
	}

	log.Infof("GetBlockByHeight info: %+v", block)
	for _, tx := range block.Transactions {
		for _, operation := range tx.Operations {
			if transferOperation, ok := operation.(*types.TransferOperation); ok {
				log.Infof("transferOperation: %+v", transferOperation)
			}
		}

	}
}

func TestWalletClient_GetTransaction(t *testing.T) {
	//8f2a2ed7620ff564a1588f7a2c090593c7195ad0
	tx, err := tw.Api.GetTransaction("550313d7c7f9fc7ed83fbad21a232a833ccb6c48")
	if err != nil {
		t.Errorf("GetTransaction failed unexpected error: %v\n", err)
		return
	} else {
		log.Infof("GetTransaction info: %+v", tx)
	}
}

func TestWalletClient_GetAccountID(t *testing.T) {
	id, err := tw.Api.GetAccountID("jjuu369")
	if err != nil {
		t.Errorf("AccountID failed unexpected error: %v\n", err)
	} else {
		log.Infof("AccountID info: %+v", id)
	}
}

func TestWalletClient_GetAccounts(t *testing.T) {
	id, err := tw.Api.GetAccounts("zbalice111")
	if err != nil {
		t.Errorf("get Accounts failed unexpected error: %v\n", err)
	} else {
		log.Infof("get Accounts info: %+v", id)
	}
}

func TestWalletClient_GetAddrBalance(t *testing.T) {
	balances, err := tw.Api.GetAddrBalance("JWOCCNPDAwLMzmrLMFhm3Z8wo5fpBzk5uHhopuD", types.MustParseObjectID("1.3.0"))
	if err != nil {
		t.Errorf("Balances failed unexpected error: %v\n", err)
	} else {
		log.Infof("Balances info: %+v", balances)
	}
}

func TestWalletClient_GetAccountByAddr(t *testing.T) {
	account, err := tw.Api.GetAccountByAddr("XWCNXAnMmSrv8eAkEnXa9ARFzpQJUzz8UDyWF")
	if err != nil {
		t.Errorf("GetAccountByAddr failed unexpected error: %v\n", err)
	} else {
		log.Infof("account info: %+v", account)
	}
}
