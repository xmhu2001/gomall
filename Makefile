.PHONY: gen-demo-proto
gen-demo-proto:
	@cd demo/demo_proto && cwgo server -I ..\..\idl\ --type RPC --module github.com/xmhu2001/gomall/demo/demo_proto --service demo_proto --idl ..\..\idl\echo.proto

.PHONY: gen-demo_thrift
gen-demo-thrift:
	@cd demo/demo_proto && cwgo server -I ..\..\idl\ --type RPC --module github.com/xmhu2001/gomall/demo/demo_thrift --service demo_thrift --idl ..\..\idl\echo.thrift

.PHONY: gen-frontend
gen-frontend:
	@cd app/frontend && cwgo server --type HTTP --idl ../../idl/frontend/home.proto --service frontend --module github.coom/xmhu2001/gomall/app/frontend -I ..\..\idl
