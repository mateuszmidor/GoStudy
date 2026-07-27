# backup
Backup your files as you develop

## install
go get github.com/cheekybits/is

## run
Configure what directories to monitor:
cd backup/cmds/backup
./backup -db=../backupdata add test1 test2
./backup -db=../backupdata list

Prepare some directories to monitor:
cd backup/cmds/backupd
mkdir test1
mkdir test2

Start monitor:
./backupd -db=../backupdata -archive=archive -interval=5s
touch test1/hello

Observe backupd archives test1