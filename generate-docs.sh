#!/bin/bash

~/go/bin/swag init -g cmd/api/main.go --parseDependency --parseInternal --parseVendor -o ./api/docs
