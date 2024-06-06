.PHONY: run update apply

run:
	sh build.sh && sh output/bootstrap.sh

update:
	cd ../.. && kitex -module tiktok idl/social.thrift

apply:
	kitex -module tiktok -service social -use tiktok/kitex_gen  ../../idl/social.thrift