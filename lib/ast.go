package lib

import (
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"path/filepath"
	"strings"
)

// ReferencedAsset holds the information for an asset referenced
// by the user
type ReferencedAsset struct {
	Name      string
	AssetPath string
}

// ReferencedAssets is a collection of assets referenced from a file
type ReferencedAssets struct {
	Caller      string
	PackageName string
	BaseDir     string
	Assets      []*ReferencedAsset
}

// HasAsset returns true if the given asset name has already been processed
// for this asset group
func (r *ReferencedAssets) HasAsset(name string) bool {
	for _, asset := range r.Assets {
		if asset.Name == name {
			return true
		}
	}
	return false
}

// GetReferencedAssets gets a list of referenced assets from the AST
func GetReferencedAssets(filenames []string) ([]*ReferencedAssets, error) {

	var result []*ReferencedAssets
	assetMap := make(map[string]*ReferencedAssets)

	for _, filename := range filenames {
		fset := token.NewFileSet()
		node, err := parser.ParseFile(fset, filename, nil, parser.AllErrors)
		if err != nil {
			return nil, err
		}

		var packageName string

		// Normalise per directory imports
		var baseDir = filepath.Dir(filename)
		var thisAssetBundle = assetMap[baseDir]
		if thisAssetBundle == nil {
			thisAssetBundle = &ReferencedAssets{Caller: filename, BaseDir: baseDir}
			assetMap[baseDir] = thisAssetBundle
		}

		ast.Inspect(node, func(node ast.Node) bool {
			switch x := node.(type) {
			case *ast.File:
				packageName = x.Name.Name
				thisAssetBundle.PackageName = packageName
			case *ast.CallExpr:
				// fmt.Printf("Call Expr = %#v\n", x)
				// fmt.Printf("Selector = %+v\n", x.Fun)
				switch x.Fun.(type) {
				case *ast.SelectorExpr:
					fn := x.Fun.(*ast.SelectorExpr)
					// fmt.Printf("x = %#v\n", fn.X)
					switch y := fn.X.(type) {
					case *ast.Ident:
						// fmt.Printf("Ident name = " + y.Name)
						if y.Name == "mewn" {
							if len(x.Args) == 1 {
								switch y := x.Args[0].(type) {
								case *ast.BasicLit:
									// fmt.Printf("argname = %s\n", y.Value)
									referencedFile := strings.Replace(y.Value, "\"", "", -1)

									// Only add the asset once
									if !thisAssetBundle.HasAsset(referencedFile) {
										// Get full asset filename
										baseDir := filepath.Dir(filename)
										assetFile, err := filepath.Abs(filepath.Join(baseDir, referencedFile))
										if err != nil {
											log.Fatal(err)
										}
										thisAsset := &ReferencedAsset{Name: referencedFile, AssetPath: assetFile}
										thisAssetBundle.Assets = append(thisAssetBundle.Assets, thisAsset)
									}
								}
							}
						}
					}
				}
			}
			return true
		})
		result = append(result, thisAssetBundle)
	}
	return result, nil
}
