# buildinfo

Since go 1.18, build information including VCS like Git info is embedded into binary.  
The information is available only in binaries built with module support.  

## Run

```bash
go run .
```

```text
go      go1.20.4
path    github.com/mateuszmidor/GoStudy/buildinfo
mod     github.com/mateuszmidor/GoStudy/buildinfo       (devel)
build   -buildmode=exe
build   -compiler=gc
build   CGO_ENABLED=1
build   CGO_CFLAGS=
build   CGO_CPPFLAGS=
build   CGO_CXXFLAGS=
build   CGO_LDFLAGS=
build   GOARCH=amd64
build   GOOS=darwin
build   GOAMD64=v1
build   vcs=git
build   vcs.revision=67e98f957b83407980fa525cd97ddd0462117808
build   vcs.time=2023-07-27T05:13:37Z
build   vcs.modified=true
```

, then you can `git describe --contains 67e98f957b83407980fa525cd97ddd0462117808` to  get git tag (if there is any).

## Get version from compiled binary "buildinfo"

You can dig out the version information from the binary like this:

```bash
 go version -m ./buildinfo
 ```

```text
go      go1.20.4
path    github.com/mateuszmidor/GoStudy/buildinfo
mod     github.com/mateuszmidor/GoStudy/buildinfo       (devel)
build   -buildmode=exe
build   -compiler=gc
build   CGO_ENABLED=1
build   CGO_CFLAGS=
build   CGO_CPPFLAGS=
build   CGO_CXXFLAGS=
build   CGO_LDFLAGS=
build   GOARCH=amd64
build   GOOS=darwin
build   GOAMD64=v1
build   vcs=git
build   vcs.revision=67e98f957b83407980fa525cd97ddd0462117808
build   vcs.time=2023-07-27T05:13:37Z
build   vcs.modified=true
```