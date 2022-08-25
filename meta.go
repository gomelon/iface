package iface

import (
	_ "embed"
	"github.com/gomelon/meta"
)

//go:embed iface.tmpl
var TmplIface string

func DefaultPkgGenFactory() meta.PkgGenFactory {
	pkgParser := meta.NewPkgParser()
	return meta.NewTmplPkgGenFactory(TmplIface,
		meta.WithOutputFilename("iface"),
		meta.WithPkgParser(pkgParser),
		meta.WithMetaParser(meta.NewParser(pkgParser)),
	)
}

const (
	MetaIface   = "+iface.Iface"
	MetaComment = "+iface.Comment"
)

func Name() string {
	return ""
}

//Iface to generate interface for struct
type Iface struct {
	//Name specifies the generated interface name
	Name string
}

//Comment copy comment to generated interface
type Comment struct {
	Value []string
}

//TODO 缺少Factory
