.PHONY: build
build:
	npx hardhat compile
	npx hardhat typechain
	./scripts/make_go_package_from_abi.sh
