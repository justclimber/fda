package main

import (
	"fmt"
	"go/ast"
	"go/constant"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"golang.org/x/tools/go/ast/inspector"

	"github.com/justclimber/fda/common/generators/entityrepo/gentemplates"
)

const keyPrefix = "Key"

type repositoryGenerator struct {
	entitiesValueSpec *ast.ValueSpec
	systemsValueSpec  *ast.ValueSpec
	info              *types.Info
	packageName       string
}

type key struct {
	Value            int64
	ShortStr         string
	FullStr          string
	PackageName      string
	StrWithoutPrefix string
}

type ecGroupData struct {
	Mask         int64
	PackageName  string
	MaskName     string
	Keys         []key
	KeysPackages map[string]bool
}

type allECGroupData struct {
	ECGroups    map[int64]ecGroupData
	PackageName string
}

type repoData struct {
	Mask         int64
	PackageName  string
	MaskName     string
	ECGroups     []string
	Keys         []key
	KeysPackages map[string]bool
}

func joinKeys(sep string, prefix, postfix string, k []key) string {
	var s []string
	for _, kk := range k {
		s = append(s, prefix+kk.StrWithoutPrefix+postfix)
	}
	return strings.Join(s, sep)
}

func (r repositoryGenerator) Generate() {
	templates := template.Must(template.New("").
		Delims("[[", "]]").
		Funcs(template.FuncMap{"joinKeys": joinKeys}).
		ParseFS(gentemplates.EmbeddedFS, "*.go.tpl"),
	)

	newPath := filepath.Join(".", "generated", r.packageName)
	err := os.MkdirAll(newPath, os.ModePerm)
	if err != nil {
		panic(err)
	}

	var entityMasks []int64
	allECGroups := allECGroupData{
		ECGroups: map[int64]ecGroupData{},
	}

	for _, ecGroupAst := range r.entitiesValueSpec.Values[0].(*ast.CompositeLit).Elts {
		ecgData := ecGroupData{
			PackageName:  r.packageName,
			Keys:         make([]key, 0),
			KeysPackages: map[string]bool{},
		}
		var mask int64
		for _, keyAst := range ecGroupAst.(*ast.CompositeLit).Elts {
			keySel := keyAst.(*ast.SelectorExpr)
			keyPackage := keySel.X.(*ast.Ident).Name

			ecgData.KeysPackages[keyPackage] = true

			k, ok := constant.Int64Val(r.info.Types[keyAst].Value)
			if !ok {
				panic("something went wrong")
			}
			mask = mask | k
			ecgData.Keys = append(ecgData.Keys, key{
				Value:            k,
				ShortStr:         keySel.Sel.Name,
				PackageName:      keyPackage,
				FullStr:          fmt.Sprintf("%s.%s", keyPackage, keySel.Sel.Name),
				StrWithoutPrefix: keySel.Sel.Name[len(keyPrefix):],
			})
		}
		entityMasks = append(entityMasks, mask)
		ecgData.Mask = mask
		ecgData.MaskName = fmt.Sprintf("Mask%d", mask)
		allECGroups.ECGroups[mask] = ecgData
		allECGroups.PackageName = ecgData.PackageName
		r.writeECGroupFile(newPath, ecgData, templates)
		r.writeChunkFile(newPath, ecgData, templates)
	}
	r.writeAllECGroupsFile(newPath, allECGroups, templates)

	for _, systemAst := range r.systemsValueSpec.Values[0].(*ast.CompositeLit).Elts {
		rData := repoData{
			PackageName:  r.packageName,
			Keys:         make([]key, 0),
			KeysPackages: map[string]bool{},
			ECGroups:     []string{},
		}
		var mask int64
		for _, keyAst := range systemAst.(*ast.CompositeLit).Elts {
			keySel := keyAst.(*ast.SelectorExpr)
			keyPackage := keySel.X.(*ast.Ident).Name

			rData.KeysPackages[keyPackage] = true

			k, ok := constant.Int64Val(r.info.Types[keyAst].Value)
			if !ok {
				panic("something went wrong")
			}
			mask = mask | k
			rData.Keys = append(rData.Keys, key{
				Value:            k,
				ShortStr:         keySel.Sel.Name,
				PackageName:      keyPackage,
				FullStr:          fmt.Sprintf("%s.%s", keyPackage, keySel.Sel.Name),
				StrWithoutPrefix: keySel.Sel.Name[len(keyPrefix):],
			})
		}
		for _, entityMask := range entityMasks {
			if entityMask&mask == mask {
				rData.ECGroups = append(rData.ECGroups, fmt.Sprintf("ECGroupMask%d", entityMask))
			}
		}
		rData.Mask = mask
		rData.MaskName = fmt.Sprintf("Mask%d", mask)
		r.writeRepoFile(newPath, rData, templates)
	}
}

func (r repositoryGenerator) writeECGroupFile(newPath string, ecgData ecGroupData, templates *template.Template) {
	filename := fmt.Sprintf("%s/ecgroup_%d.go", newPath, ecgData.Mask)
	writeToTemplate(filename, ecgData, templates, "ecgroup.go.tpl")
}

func (r repositoryGenerator) writeChunkFile(newPath string, ecgData ecGroupData, templates *template.Template) {
	filename := fmt.Sprintf("%s/chunk_%d.go", newPath, ecgData.Mask)
	writeToTemplate(filename, ecgData, templates, "chunk.go.tpl")
}

func (r repositoryGenerator) writeAllECGroupsFile(newPath string, ecGroups allECGroupData, templates *template.Template) {
	filename := fmt.Sprintf("%s/ecgroups.go", newPath)
	writeToTemplate(filename, ecGroups, templates, "all_ecgroups.go.tpl")
}

func (r repositoryGenerator) writeRepoFile(newPath string, rData repoData, templates *template.Template) {
	filename := fmt.Sprintf("%s/repo_%d.go", newPath, rData.Mask)
	writeToTemplate(filename, rData, templates, "repo_for_mask.go.tpl")
}

func writeToTemplate(filepath string, data interface{}, t *template.Template, tString string) {
	f, err := os.OpenFile(filepath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		panic(err)
	}
	defer mustCloseFile(f)
	if err = t.ExecuteTemplate(f, tString, data); err != nil {
		panic(err)
	}
}

func mustCloseFile(f *os.File) {
	if err := f.Close(); err != nil {
		panic(err)
	}
}

func main() {
	//cwd, err := os.Getwd()
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("  cwd = %s\n", cwd)
	//fmt.Printf("  os.Args = %#v\n", os.Args)
	//
	//for _, ev := range []string{"GOARCH", "GOOS", "GOFILE", "GOLINE", "GOPACKAGE"} {
	//	fmt.Println("  ", ev, "=", os.Getenv(ev))
	//}

	//path := "/Users/aaakimov/pet/fda/server/worldprocessor/ecs/declaration.go"

	if len(os.Args) < 2 {
		log.Fatal("Please specify new package name as an argument")
	}

	path := os.Getenv("GOFILE")
	if path == "" {
		log.Fatal("GOFILE must be set")
	}

	fileSet := token.NewFileSet()
	astInFile, err := parser.ParseFile(
		fileSet,
		path,
		nil,
		parser.ParseComments,
	)
	if err != nil {
		log.Fatalf("parse file: %v", err)
	}

	insp := inspector.New([]*ast.File{astInFile})

	// Obtain type information.
	conf := types.Config{Importer: importer.ForCompiler(fileSet, "source", nil)}
	info := &types.Info{
		Defs:  make(map[*ast.Ident]types.Object),
		Types: make(map[ast.Expr]types.TypeAndValue),
	}
	//_, err = conf.Check("github.com/justclimber/fda/server/worldprocessor/ecs", fileSet, []*ast.File{astInFile}, info)
	_, err = conf.Check(os.Getenv("GOPACKAGE"), fileSet, []*ast.File{astInFile}, info)
	if err != nil {
		log.Fatal(err) // type error
	}

	genTask := repositoryGenerator{
		info:        info,
		packageName: os.Args[1],
	}

	iFilter := []ast.Node{
		&ast.GenDecl{},
	}

	insp.Nodes(iFilter, func(node ast.Node, push bool) (proceed bool) {
		genDecl := node.(*ast.GenDecl)
		if genDecl.Doc == nil {
			return false
		}
		valueSpec, ok := genDecl.Specs[0].(*ast.ValueSpec)
		if !ok {
			return false
		}
		for _, comment := range genDecl.Doc.List {
			switch comment.Text {
			case "//generate:entities":
				genTask.entitiesValueSpec = valueSpec
			case "//generate:systems":
				genTask.systemsValueSpec = valueSpec
			}
		}
		return false
	})

	genTask.Generate()
}
