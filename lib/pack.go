package lib

import (
	"fmt"
	"log"
)

// GeneratePackFileString creates the contents of a pack file
func GeneratePackFileString(assetBundle *ReferencedAssets) string {

	result := fmt.Sprintf("package %s\n", assetBundle.PackageName)
	result += "// Autogenerated by Mewn - Do not alter\n\n"
	result += "import \"github.com/leaanthony/mewn\"\n"
	result += "func init() {\n"
	for _, asset := range assetBundle.Assets {
		packedData, err := CompressFile(asset.AssetPath)
		if err != nil {
			log.Fatal(err)
		}
		result += fmt.Sprintf("  mewn.AssetDirectory[\"%s\"] = \"%s\"\n", asset.Name, packedData)
	}
	result += "}\n"

	return result
}
