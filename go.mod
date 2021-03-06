module github.com/assetsadapterstore/jwocc-adapter

go 1.13

require (
	github.com/astaxie/beego v1.12.0
	github.com/blocktree/go-owaddress v1.1.13
	github.com/blocktree/go-owcdrivers v1.2.19
	github.com/blocktree/go-owcrypt v1.1.4
	github.com/blocktree/openwallet/v2 v2.0.10
	github.com/blocktree/whitecoin-adapter v1.2.1
	github.com/imroc/req v0.2.4
	github.com/pkg/errors v0.8.1
	github.com/shopspring/decimal v0.0.0-20200105231215-408a2507e114
	github.com/tidwall/gjson v1.3.5
)

replace github.com/imroc/req => github.com/blocktree/req v0.2.5
