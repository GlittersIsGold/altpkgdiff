package pkg

import (
	"fmt"

	"github.com/GlittersIsGold/altpkgdiff/api"
	rpmv "github.com/knqyf263/go-rpm-version"
)

type PackageDiff struct {
	OnlyInP10        map[string][]api.Package `json:"only_in_p10"`
	OnlyInSisyphus   map[string][]api.Package `json:"only_in_sisyphus"`
	HigherInSisyphus map[string][]api.Package `json:"higher_in_sisyphus"`
}

func DiffPkgs(sisyphusPackages, p10Packages []api.Package) PackageDiff {
	
	diffs := PackageDiff{
		OnlyInP10:        make(map[string][]api.Package),
		OnlyInSisyphus:   make(map[string][]api.Package),
		HigherInSisyphus: make(map[string][]api.Package),
	}

	sisyphusMap := make(map[string]api.Package)
	p10Map := make(map[string]api.Package)

	for _, pkg := range sisyphusPackages {
		key := pkg.Name + ":" + pkg.Arch
		sisyphusMap[key] = pkg
	}

	for _, pkg := range p10Packages {
		key := pkg.Name + ":" + pkg.Arch
		p10Map[key] = pkg
	}

	for key, p10Pkg := range p10Map {
		if sisPkg, found := sisyphusMap[key]; found {
			if CmpPkgVers(sisPkg, p10Pkg) > 0 {
				diffs.HigherInSisyphus[sisPkg.Arch] = append(diffs.HigherInSisyphus[sisPkg.Arch], sisPkg)
			}
			delete(sisyphusMap, key)
		} else {
			diffs.OnlyInP10[p10Pkg.Arch] = append(diffs.OnlyInP10[p10Pkg.Arch], p10Pkg)
		}
	}

	for _, pkg := range sisyphusMap {
		diffs.OnlyInSisyphus[pkg.Arch] = append(diffs.OnlyInSisyphus[pkg.Arch], pkg)
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