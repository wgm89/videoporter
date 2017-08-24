package handles

import (
	"net/http"
	"strings"

	"videoporter/env"
	. "videoporter/public"

	"github.com/elazarl/go-bindata-assetfs"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

type binaryFileSystem struct {
	fs http.FileSystem
}

func (b *binaryFileSystem) Open(name string) (http.File, error) {
	return b.fs.Open(name)
}

func (b *binaryFileSystem) Exists(prefix string, filepath string) bool {
	if p := strings.TrimPrefix(filepath, prefix); len(p) < len(filepath) {
		if _, err := b.fs.Open(p); err != nil {
			return false
		}
		return true
	}
	return false
}

func BinaryFileSystem(root string) *binaryFileSystem {
	fs := &assetfs.AssetFS{Asset, AssetDir, AssetInfo, root}
	return &binaryFileSystem{
		fs,
	}
}

func InitHandle(r *gin.Engine) {
	InitHomeHandle(r)
	InitVideoHandle(r)
	if env.IsProduction {
		r.Use(static.Serve("/", BinaryFileSystem("")))
	} else {
		r.Static("/public", "./public")
	}
}
