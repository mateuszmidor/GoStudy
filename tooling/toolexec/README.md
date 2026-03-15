# toolexec
This allows to run an interceptor application before each tool `go build` is about to run.
```sh
go build -toolexec='./toolexec-wrapper'
```
https://www.youtube.com/watch?v=8Rw-fVEjihw&list=PLDWZ5uzn69ewrYyHTNrXlrWVDjLiOX0Yb&index=18

## Run

```sh
make run
```
output:
```
# command-line-arguments
TOOLEXEC: /usr/lib/go/pkg/tool/linux_amd64/compile
  -o
  $WORK/b001/_pkg_.a
  -trimpath
  $WORK/b001=>
  -p
  main
  -lang=go1.25
  -complete
  -buildid
  ezuZVBPygmf846BUPQj5/ezuZVBPygmf846BUPQj5
  -goversion
  go1.25.7 X:nodwarf5
  -c=4
  -nolocalimports
  -importcfg
  $WORK/b001/importcfg
  -pack
  ./main.go
# command-line-arguments
TOOLEXEC: /usr/lib/go/pkg/tool/linux_amd64/link
  -o
  $WORK/b001/exe/a.out
  -importcfg
  $WORK/b001/importcfg.link
  -X=runtime.godebugDefault=asynctimerchan=1,containermaxprocs=0,decoratemappings=0,gotestjsonbuildtext=1,gotypesalias=0,httpcookiemaxnum=0,httplaxcontentlength=1,httpmuxgo121=1,httpservecontentkeepheaders=1,multipathtcp=0,randseednop=0,rsa1024min=0,tls10server=1,tls3des=1,tlsmlkem=0,tlsrsakex=1,tlssha1=1,tlsunsafeekm=1,updatemaxprocs=0,urlmaxqueryparams=0,winreadlinkvolume=0,winsymlink=0,x509keypairleaf=0,x509negativeserial=1,x509rsacrt=0,x509sha256skid=0,x509usepolicies=0
  -buildmode=exe
  -buildid=18gekAM0HNDeZCpwVvhJ/ezuZVBPygmf846BUPQj5/yPlSVoFfBK7O2VgZwzPE/18gekAM0HNDeZCpwVvhJ
  -extld=gcc
  $WORK/b001/_pkg_.a
```