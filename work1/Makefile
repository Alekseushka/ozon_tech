CURDIR=$(shell pwd)
BINDIR=${CURDIR}/bin 
GOVER=$(shell go version | perl -nle '/(go\d\S+)/; print $$1;')
MOCKGEN=${BINDIR}/mockgen_${GOVER}
SMARTIMPORTS=${BINDIR}/smartimports_${GOVER}
LINTVER=v1.49.0
LINTBIN=${BINDIR}/lint_${GOVER}_${LINTVER}
PACKAGE=github.com/Alekseushka/ozon_tech/work1/cmd/bot

all: build test lint

build: bindir
	go build -o ${BINDIR}/bot ${PACKAGE}

test:
	go test ./...

run:
	go run ${PACKAGE}

lint: install-lint
	${LINTBIN} run

bindir: 
	mkdir -p ${BINDIR}

install-lint: bindir
	test -f ${LINTBIN} || \
	(GOBIN=${BINDIR} go install github.com/golangci/golangci-lint/cmd/golangci-lint@${LINTVER} && \
	mv ${BINDIR}/golangci-lint ${LINTBIN})