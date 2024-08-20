package pkg

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/GlittersIsGold/altpkgdiff/api"
)

type PackageDiff struct {
	OnlyInP10        []api.Package `json:"only_in_p10"`
	OnlyInSisyphus   []api.Package `json:"only_in_sisyphus"`
	HigherInSisyphus []api.Package `json:"higher_in_sisyphus"`
}

func DiffPkgs(p10Packages, sisyphusPackages []api.Package) PackageDiff {
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
			if CmpPkgVers(p10Pkg, sisPkg) < 0 {
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

func CmpPkgVers(pkg1, pkg2 api.Package) int {

	if pkg1.Epoch != pkg2.Epoch {
		return pkg1.Epoch - pkg2.Epoch
	}

	verComp := CmpVerRel(pkg1.Version, pkg2.Version)
	if verComp != 0 {
		return verComp
	}

	return CmpVerRel(pkg1.Release, pkg2.Release)
}

func CmpVerRel(v1, v2 string) int {
	parts1 := SplitVer(v1)
	parts2 := SplitVer(v2)

	for i := 0; i < len(parts1) && i < len(parts2); i++ {
		comp := CmpVerParts(parts1[i], parts2[i])
		if comp != 0 {
			return comp
		}
	}

	return len(parts1) - len(parts2)
}

func SplitVer(version string) []string {
	re := regexp.MustCompile(`(\d+|\D+)`)
	return re.FindAllString(version, -1)
}

func CmpVerParts(part1, part2 string) int {
	isDigit1 := isNumeric(part1)
	isDigit2 := isNumeric(part2)

	if isDigit1 && isDigit2 {
		num1, _ := strconv.Atoi(part1)
		num2, _ := strconv.Atoi(part2)
		return num1 - num2
	}

	if isDigit1 {
		return -1
	}
	if isDigit2 {
		return 1
	}

	return strings.Compare(part1, part2)
}

func isNumeric(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}