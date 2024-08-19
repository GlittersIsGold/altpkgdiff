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

func ComparePackages(p10Packages, sisyphusPackages map[string][]api.Package) map[string]PackageDiff {
    diffs := make(map[string]PackageDiff)

    for arch, p10List := range p10Packages {
        sisyphusList := sisyphusPackages[arch]

        onlyInP10 := []api.Package{}
        onlyInSisyphus := []api.Package{}
        higherInSisyphus := []api.Package{}

        sisyphusMap := make(map[string]api.Package)
        for _, pkg := range sisyphusList {
            sisyphusMap[pkg.Name] = pkg
        }

        for _, pkg := range p10List {
            if sisPkg, found := sisyphusMap[pkg.Name]; found {
                if semver.MustParse(sisPkg.Version).GT(semver.MustParse(pkg.Version)) {
                    higherInSisyphus = append(higherInSisyphus, sisPkg)
                }
                delete(sisyphusMap, pkg.Name)
            } else {
                onlyInP10 = append(onlyInP10, pkg)
            }
        }

        for _, pkg := range sisyphusMap {
            onlyInSisyphus = append(onlyInSisyphus, pkg)
        }

        diffs[arch] = PackageDiff{
            OnlyInP10:     onlyInP10,
            OnlyInSisyphus: onlyInSisyphus,
            HigherInSisyphus: higherInSisyphus,
        }
    }
    return diffs
}