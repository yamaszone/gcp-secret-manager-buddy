version:
	test -n "$(VERSION)"  # $$VERSION must be set.

github-token:
	test -n "$(GITHUB_TOKEN)" # $$GITHUB_TOKEN must be set.

fetch:
	go get \
	github.com/mitchellh/gox \
	github.com/aktau/github-release

clean:
	rm -rf ./release

build: version
	gox -verbose \
	-osarch="windows/amd64 linux/amd64 darwin/amd64" \
	-output="release/{{.Dir}}-${VERSION}-{{.OS}}-{{.Arch}}" .

test:
	go clean -testcache
	go test ./...

publish: version github-token clean test build
	github-release release --user yamaszone --repo gcp-secret-manager-buddy --tag ${VERSION} \
	--name "${VERSION}" --description "${VERSION}" && \
	github-release upload --user yamaszone --repo gcp-secret-manager-buddy --tag ${VERSION} \
	--name "gcp-secret-manager-buddy-${VERSION}-windows-amd64.exe" --file release/gcp-secret-manager-buddy-${VERSION}-windows-amd64.exe; \
	github-release upload --user yamaszone --repo gcp-secret-manager-buddy --tag ${VERSION} \
	--name "gcp-secret-manager-buddy-${VERSION}-windows-amd64.exe.asc" --file release/gcp-secret-manager-buddy-${VERSION}-windows-amd64.exe.asc; \
	for qualifier in darwin-amd64 linux-amd64 ; do \
		github-release upload --user yamaszone --repo gcp-secret-manager-buddy --tag ${VERSION} \
		--name "gcp-secret-manager-buddy-${VERSION}-$$qualifier" --file release/gcp-secret-manager-buddy-${VERSION}-$$qualifier; \
		github-release upload --user yamaszone --repo gcp-secret-manager-buddy --tag ${VERSION} \
		--name "gcp-secret-manager-buddy-${VERSION}-$$qualifier.asc" --file release/gcp-secret-manager-buddy-${VERSION}-$$qualifier.asc; \
	done
