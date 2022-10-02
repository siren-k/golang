#!/usr/bin/env bash
mockgen -destination mocks_test.go -package dbinterface golang/chap6/dbinterface DB,Transaction