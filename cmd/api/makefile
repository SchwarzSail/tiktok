.PHONY: all user video interaction

update:
	hz update -idl ../../idl/api.thrift

all: user video interaction api

user:
	cd ../user && sh build.sh && sh output/bootstrap.sh &

video:
	cd ../video && sh build.sh && sh output/bootstrap.sh &

interaction:
	cd ../interaction && sh build.sh && sh output/bootstrap.sh &

run:
	sh build.sh && sh output/bootstrap.sh
