#!/bin/bash
docker run --rm -it --name golang -v $PWD/src:/go/src --workdir="/go/src" golang
