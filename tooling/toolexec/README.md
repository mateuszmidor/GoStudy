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
# github.com/mateuszmidor/GoStudy/tooling/toolexec
TOOLEXEC: C:\Program Files (x86)\Go\pkg\tool\windows_386\compile.exe [-o $WORK\b001\_pkg_.a -trimpath $WORK\b001=> -p main -lang=go1.21 -complete -buildid 06uRGBL4UR-qo1H4XpYb/06uRGBL4UR-qo1H4XpYb -goversion go1.24.2 -c=4 -nolocalimports -importcfg $WORK\b001\importcfg -pack .\main.go]
# github.com/mateuszmidor/GoStudy/tooling/toolexec
TOOLEXEC: C:\Program Files (x86)\Go\pkg\tool\windows_386\link.exe [-o $WORK\b001\exe\a.out.exe -importcfg $WORK\b001\importcfg.link -X=runtime.godebugDefault=asynctimerchan=1,gotestjsonbuildtext=1,gotypesalias=0,httplaxcontentlength=1,httpmuxgo121=1,httpservecontentkeepheaders=1,multipathtcp=0,randseednop=0,rsa1024min=0,tls10server=1,tls3des=1,tlsmlkem=0,tlsrsakex=1,tlsunsafeekm=1,winreadlinkvolume=0,winsymlink=0,x509keypairleaf=0,x509negativeserial=1,x509rsacrt=0,x509usepolicies=0 -buildmode=pie -buildid=NkViBFhfkQzq8qeR1Ram/06uRGBL4UR-qo1H4XpYb/hhcPRqJJK-lDoxU0AeAM/NkViBFhfkQzq8qeR1Ram -extld=gcc $WORK\b001\_pkg_.a]
```