copied from https://github.com/urfave/cli/issues/1583#issuecomment-1328190259


What components are there?
 - Flags wrapper:
   - Help prompt
     - Version: encourage setting compile-time variables, such as git hash. https://blog.alexellis.io/inject-build-time-vars-golang/ If undocumented, could lead to byte-level unreproducible builds, but it's way better than incrementing versions manually. [^ver], [^notv3]
   - Shell completions [^notv3]
   - Default values
     - In `Flag`, replace `Required bool` with `Action`, renamed to `Condition`. It is trivial to write common implementations with generic functions. The library can provide commonly used ones, such as `NotDefault` (must be set by user, `--foo=""` is valid) and `NotEmpty`. In most cases, omitted or looking like `cli.StringFlag{Name: "foo", Condition: cli.NotEmpty}`. 
Ideas for changes I've had so far. I've asked around for feedback, and clarifying what `Required` did come up.
   - Validation (after type conversion)
   - Supporting yaml with nested flag definition structure and/or exposing setting flags in context [^notv3]
- Context
  - Signals [^notv3]
    - Make it clear on how interrupts are expected to be handled
  - Exit codes [^notv3]
    - Documenting and handling exit codes: Exit code caller should reference an already-defined code, similar to calling up a flag, not just 'exit 1'. [^notv3]
- Commands
  - Before/After
  - Subcommands
    - Flag-dependant subcommands: Quick example: `$ app` → `ls $(pwd)` `$ app <flag>` → `cd <flag>` [^notv3]

It should also include everything in the v3 milestone. Anything I missed? The only parts I expect to get stuck on are shell completions, and yaml. They are not critical for a v3 beta, though.

[^notv3]: Not needed for a functional beta release. No changes are necessary for adding it without breaking changes.
[^ver]: ```
	$ kubectl version
	Client Version: version.Info{Major:"1", Minor:"25", GitVersion:"v1.25.4", GitCommit:"872a965c6c6526caa949f0c6ac028ef7aff3fb78", GitTreeState:"clean", BuildDate:"2022-11-09T13:36:36Z", GoVersion:"go1.19.3", Compiler:"gc", Platform:"linux/amd64"}
	```
	Implementation:
	```
	var BuildInfoGitCommit string
	```
	```
	export GIT_COMMIT=$(git rev-list -1 HEAD) && \
 	go build -ldflags "-X cli.BuildInfoGitCommit=$GIT_COMMIT"
	```
