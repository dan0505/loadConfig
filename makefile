release:
	# to call > make release new_version=v0.1.2
	git tag $(new_version)
	git push origin $(new_version)
	GOPROXY=proxy.golang.org go list -m github.com/dan0505/loadConfig@$(new_version)
	echo released version $(new_version)