#!/usr/bin/env bash
protoc  --go_out=../go-proto *.proto
protoc-c  --c_out=../fw/logic *.proto
