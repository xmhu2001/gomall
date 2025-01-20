.PHONY: gen-demo-proto
gen-demo-proto:
	@cd demo/demo_proto && cwgo server -I ..\..\idl\ --type RPC --module github.com/xmhu2001/gomall/demo/demo_proto --service demo_proto --idl ..\..\idl\echo.proto

.PHONY: gen-demo_thrift
gen-demo-thrift:
	@cd demo/demo_proto && cwgo server -I ..\..\idl\ --type RPC --module github.com/xmhu2001/gomall/demo/demo_thrift --service demo_thrift --idl ..\..\idl\echo.thrift

.PHONY: gen-frontend
gen-frontend:
	@cd app/frontend && cwgo server --type HTTP --idl ../../idl/frontend/auth_page.proto --service frontend --module github.com/xmhu2001/gomall/app/frontend -I ..\..\idl

.PHONY: gen-user
gen-user:
	@cd rpc_gen && cwgo client --type RPC --service user --module github.com/xmhu2001/gomall/rpc_gen -I ..\idl --idl ../idl/user.proto
	@cd app/user && cwgo server --type RPC --service user --module github.com/xmhu2001/gomall/app/user --pass "-use github.com/xmhu2001/gomall/rpc_gen/kitex_gen" -I ../../idl --idl ../../idl/user.proto

.PHONY: gen-product
gen-product:
	@cd rpc_gen && cwgo client --type RPC --service product --module github.com/xmhu2001/gomall/rpc_gen -I ..\idl --idl ../idl/product.proto
	@cd app/product && cwgo server --type RPC --service product --module github.com/xmhu2001/gomall/app/product --pass "-use github.com/xmhu2001/gomall/rpc_gen/kitex_gen" -I ../../idl --idl ../../idl/product.proto