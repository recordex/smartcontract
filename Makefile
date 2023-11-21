.PHONY: build
build:
	npx hardhat compile
	./scripts/make_go_package_from_abi.sh
