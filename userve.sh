#!/bin/bash

src=$GOPATH/src
vp=github.com/varunamachi/vaali
sp=${src}/${vp}/
dp=${src}/github.com/varunamachi/sprw/vendor/${vp}/

rsync -av --exclude='vendor/' --exclude='.git/' ${sp} ${dp} && \
cd cmd/sprw && \
go install && \
sprw serve

