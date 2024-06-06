.PHONY: run update apply

run:
	sh build.sh && sh output/bootstrap.sh

update:
	cd ../.. && kitex -module tiktok idl/interaction.thrift

apply:
	kitex -module tiktok -service interaction -use tiktok/kitex_gen  ../../idl/interaction.thrift