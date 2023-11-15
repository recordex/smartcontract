.PHONY: compile
compile:
	npx hardhat compile
	sh scripts/make_go_package_from_abi.sh

.PHONY: deploy-dev
deploy-dev:
	git checkout dev
	git fetch origin
	git reset --hard origin/main
	git merge ${branch}
	git log -n 2
	git push -f origin dev
	git checkout ${branch}
