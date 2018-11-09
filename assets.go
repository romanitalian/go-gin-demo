package main

import (
	"time"

	"github.com/jessevdk/go-assets"
)

var _Assets3737a75b5254ed1f6d588b40a3449721f9ea86c2 = "<!DOCTYPE html>\r\n<html lang=\"en\">\r\n<head>\r\n    <meta charset=\"UTF-8\">\r\n    <title>{{ .Foo }}</title>\r\n</head>\r\n<body>\r\n<h1>{{ .Foo }}</h1>\r\n</body>\r\n</html>"

// Assets returns go-assets FileSystem
var Assets = assets.NewFileSystem(map[string][]string{"/": []string{"html"}, "/html": []string{"index.tmpl"}}, map[string]*assets.File{
	"/": &assets.File{
		Path:     "/",
		FileMode: 0x800001ff,
		Mtime:    time.Unix(1541807402, 1541807402114305700),
		Data:     nil,
	}, "/html": &assets.File{
		Path:     "/html",
		FileMode: 0x800001ff,
		Mtime:    time.Unix(1541807477, 1541807477927502200),
		Data:     nil,
	}, "/html/index.tmpl": &assets.File{
		Path:     "/html/index.tmpl",
		FileMode: 0x1b6,
		Mtime:    time.Unix(1541807477, 1541807477913542200),
		Data:     []byte(_Assets3737a75b5254ed1f6d588b40a3449721f9ea86c2),
	}}, "")
