.PHONY: run update apply

run:
	sh build.sh && sh output/bootstrap.sh

update:
	cd ../.. && kitex -module tiktok idl/user.thrift

apply:
	kitex -module tiktok -service user -use tiktok/kitex_gen  ../../idl/user.thrift