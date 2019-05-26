#!/bin/bash
protoc --go_out=plugins=grpc+special-annontation:. *.proto
