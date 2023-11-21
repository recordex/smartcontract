.PHONY: build
build:
	npx hardhat compile
	npx hardhat typechain
	sh scripts/make_go_package_from_abi.sh
