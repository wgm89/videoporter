#!/bin/bash

rm templates/bindata.go
go-bindata templates/...
sed -i "s/package main/package templates/g" ./bindata.go
mv bindata.go templates


rm public/bindata_assetfs.go
go-bindata-assetfs public/...
sed -i "s/package main/package public/g" ./bindata_assetfs.go
mv bindata_assetfs.go public
