#!/usr/bin/env bash

go build . # -gcflags=-B .
./finder_cli <<< $'krk gdn\ngdn dxb\ndxb sez\nsez hkg\nhkg bkk\nbkk syd\nsyd ppt\nppt hnl\nhnl sfo\nsfo jfk\njfk bcn\nbcn fra\nfra krk\nexit\n'