package pkg

import (
	rpmv "github.com/knqyf263/go-rpm-version"
	"github.com/GlittersIsGold/altpkgdiff/api"
	"fmt"
)

type PackageDiff struct {
	OnlyInP10        []api.Package `json:"only_in_p10"`
	OnlyInSisyphus   []api.Package `json:"only_in_sisyphus"`
	HigherInSisyphus []api.Package `json:"higher_in_sisyphus"`
}

func DiffPkgs(sisyphusPackages, p10Packages []api.Package) PackageDiff {
	
	diffs := PackageDiff{}

	p10Map := make(map[string]api.Package)
	sisyphusMap := make(map[string]api.Package)

	for _, pkg := range p10Packages {
		p10Map[pkg.Name] = pkg
	}

	for _, pkg := range sisyphusPackages {
		sisyphusMap[pkg.Name] = pkg
	}

	for name, p10Pkg := range p10Map {
		if sisPkg, found := sisyphusMap[name]; found {
			if CmpPkgVers(p10Pkg, sisPkg) > 0 {
				diffs.HigherInSisyphus = append(diffs.HigherInSisyphus, sisPkg)
			}
			delete(sisyphusMap, name)
		} else {
			diffs.OnlyInP10 = append(diffs.OnlyInP10, p10Pkg)
		}
	}

	for _, pkg := range sisyphusMap {
		diffs.OnlyInSisyphus = append(diffs.OnlyInSisyphus, pkg)
	}

	return diffs
}

func CmpPkgVers(src, dst api.Package) int {
	vSrc := CreateVersion(src);
	vDst := CreateVersion(dst);
	return vSrc.Compare(vDst);
}

func CreateVersion(pkg api.Package) rpmv.Version {
    vStr := fmt.Sprintf("%d:%s-%s", pkg.Epoch, pkg.Version, pkg.Release)
    v := rpmv.NewVersion(vStr)
	return v;
}