#!/bin/bash
cd lq
go test -coverprofile=coverage.out
go tool cover -func=coverage.out
