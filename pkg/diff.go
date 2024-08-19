package pkg

import (
    "github.com/blang/semver/v4"
    "github.com/GlittersIsGold/altpkgdiff/api"
)

type PackageDiff struct {
    OnlyInP10   []api.Package `json:"only_in_p10"`
    OnlyInSisyphus []api.Package `json:"only_in_sisyphus"`
    HigherInSisyphus []api.Package `json:"higher_in_sisyphus"`
}

func ComparePackages(p10Packages, sisyphusPackages []api.Package) PackageDiff {
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

            p10Version := semver.MustParse(p10Pkg.Version)
			sisVersion := semver.MustParse(sisPkg.Version)
			
			if p10Version.GT(sisVersion) {
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