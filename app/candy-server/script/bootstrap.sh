#!/bin/bash
CURDIR=$(cd $(dirname $0); pwd)
BinaryName=candy-server-cluster
echo "$CURDIR/bin/${BinaryName}"
exec $CURDIR/bin/${BinaryName}