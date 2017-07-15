#!/bin/sh
export REALM_CORE_PREFIX="$HOME/realm/realm-core"
export CGO_LDFLAGS="-framework CoreFoundation $REALM_CORE_PREFIX/src/realm/librealm.a"
#export CGO_LDFLAGS="-L$REALM_CORE_PREFIX/src/realm -lrealm"
export CGO_CXXFLAGS="-std=c++14 -I$REALM_CORE_PREFIX/src"
go build -a -x -ldflags -s
go install
