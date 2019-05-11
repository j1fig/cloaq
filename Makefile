BUILD_DIR = "/tmp/__cloaq_build__"
RELEASE_DIR = "${BUILD_DIR}/${CLOAQ_VERSION}"
LINUX_RELEASE_DIR = ${RELEASE_DIR}/linux-amd64
MACOSX_RELEASE_DIR = ${RELEASE_DIR}/darwin-amd64
WINDOWS_RELEASE_DIR = ${RELEASE_DIR}/windows-amd64

.PHONY: install
install:
	go install

.PHONY: release
release:
	mkdir -p ${RELEASE_DIR}
	mkdir ${LINUX_RELEASE_DIR}
	mkdir ${MACOSX_RELEASE_DIR}
	mkdir ${WINDOWS_RELEASE_DIR}
	GOOS=linux GOARCH=amd64 go build -o ${LINUX_RELEASE_DIR} -v -race cmd
	GOOS=darwin GOARCH=amd64 go build -o ${MACOSX_RELEASE_DIR} -v -race cmd
	GOOS=windows GOARCH=amd64 go build -o ${WINDOWS_RELEASE_DIR} -v -race cmd
	git tag ${CLOAQ_VERSION}
	git push origin ${CLOAQ_VERSION}


.PHONY: clean
clean:
	rm -rf ${BUILD_DIR}
