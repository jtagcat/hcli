package hcli

var (
	semVer    string
	gitCommit string

	VersionGiver string
	// TODO: ...
)

// TODO: don't export gitCommit, as it then only accessible at build time
// TODO: interface for printing version? allow calling it from elsewhere without exit?
func version() {
	// TODO:
	// to json and print to stdout
}

// https://blog.alexellis.io/inject-build-time-vars-golang/
// var GitCommit string
// export GIT_COMMIT=$(git rev-list -1 HEAD) && \
//   go build -ldflags "-X hcli.GitCommit=$GIT_COMMIT"

// note: does not work with go run
