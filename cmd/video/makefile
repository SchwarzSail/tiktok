.PHONY: run update apply

run:
	sh build.sh && sh output/bootstrap.sh

update:
	cd ../.. && kitex -module tiktok idl/video.thrift

apply:
	kitex -module tiktok -service video -use tiktok/kitex_gen  ../../idl/video.thrift