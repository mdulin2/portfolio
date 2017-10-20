#!/bin/bash

## Be friendly and say hello
echo "🤖 - Hello human-friend. Im gonna install some go dependencies for you.";

## Dependencies
go get github.com/stretchr/testify/assert;
go get github.com/siddontang/go-mysql/client;
go get github.com/gorilla/mux;
go get golang.org/x/crypto/bcrypt;
go get github.com/gorilla/handlers;
go get github.com/goware/emailx;

## Be friendly and say goodbye
echo "🤖 - Okay human-friend, all done. ✨";