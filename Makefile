.PHONY: protos

protos:
	# $ is used as an escape character in makefile.
	# export PATH=$$PATH:/usr/local/go/bin
	# export PATH=$$PATH:$$HOME/go/bin
	protoc -I=proto --go_out=plugins=grpc:proto proto/blockchain.proto
