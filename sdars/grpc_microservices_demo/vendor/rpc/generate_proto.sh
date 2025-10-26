#!/bin/bash

protoc *.proto --go_out=plugins=grpc:.
