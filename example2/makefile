pbdir=`pwd`/proto
.PHONY: clean srv cli

all: clean srv cli

clean:
	# go clean
	rm -rf s c
	cd ${pbdir} && rm -rf *.go

g:
	cd ${pbdir} &&  protoc  --micro_out=. --go_out=. *.proto 

srv: g
	cd srv && go build -o ../s

cli:
	cd cli && go build -o ../c
