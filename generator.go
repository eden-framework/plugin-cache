package main

import (
	"fmt"
	"github.com/eden-framework/plugins"
	"path"
)

var Plugin GenerationPlugin

type GenerationPlugin struct {
}

func (g *GenerationPlugin) GenerateEntryPoint(opt plugins.Option, cwd string) string {
	globalPkgPath := path.Join(opt.PackageName, "internal/global")
	globalFilePath := path.Join(cwd, "internal/global")
	tpl := fmt.Sprintf(`,
		{{ .UseWithoutAlias "github.com/eden-framework/eden-framework/pkg/application" "" }}.WithConfig(&{{ .UseWithoutAlias "%s" "%s" }}.CacheConfig)`, globalPkgPath, globalFilePath)
	return tpl
}

func (g *GenerationPlugin) GenerateFilePoint(opt plugins.Option, cwd string) []*plugins.FileTemplate {
	file := plugins.NewFileTemplate("global", path.Join(cwd, "internal/global/cache.go"))
	file.WithBlock(`
var CacheConfig = struct {
	Cache *{{ .UseWithoutAlias "github.com/eden-framework/plugin-cache/cache" "" }}.Cache
}{
	Cache: &{{ .UseWithoutAlias "github.com/eden-framework/plugin-cache/cache" "" }}.Cache{
		Driver: {{ .UseWithoutAlias "github.com/eden-framework/plugin-cache/cache" "" }}.DRIVER__REDIS,
		Host:   "localhost",
		Port:   6379,
	},
}
`)

	return []*plugins.FileTemplate{file}
}
