package pkg

import (
	"fmt"

	"github.com/GlittersIsGold/altpkgdiff/api"
	rpmv "github.com/knqyf263/go-rpm-version"
)

type PackageDiff struct {
	OnlyInDst        map[string][]api.Package `json:"only_in_dst"`
	OnlyInSrc   map[string][]api.Package `json:"only_in_src"`
	HigherInSrc map[string][]api.Package `json:"higher_in_src"`
}

func DiffPkgs(sisyphusPackages, p10Packages []api.Package) PackageDiff {
	
	diffs := PackageDiff{
		OnlyInDst:        make(map[string][]api.Package),
		OnlyInSrc:   make(map[string][]api.Package),
		HigherInSrc: make(map[string][]api.Package),
	}

	srcMap := make(map[string]api.Package)
	dstMap := make(map[string]api.Package)

	for _, pkg := range sisyphusPackages {
		key := pkg.Name + ":" + pkg.Arch
		srcMap[key] = pkg
	}

	for _, pkg := range p10Packages {
		key := pkg.Name + ":" + pkg.Arch
		dstMap[key] = pkg
	}

	for key, dstPkg := range dstMap {
		if srcPkg, found := srcMap[key]; found {
			if CmpPkgVers(srcPkg, dstPkg) > 0 {
				diffs.HigherInSrc[srcPkg.Arch] = append(diffs.HigherInSrc[srcPkg.Arch], srcPkg)
			}
			delete(srcMap, key)
		} else {
			diffs.OnlyInDst[dstPkg.Arch] = append(diffs.OnlyInDst[dstPkg.Arch], dstPkg)
		}
	}

	for _, pkg := range srcMap {
		diffs.OnlyInSrc[pkg.Arch] = append(diffs.OnlyInSrc[pkg.Arch], pkg)
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

func FilterByArch(packages []api.Package, arch string) []api.Package {
    var filtered []api.Package
    for _, p := range packages {
        if p.Arch == arch {
            filtered = append(filtered, p)
        }
    }
    return filtered
}