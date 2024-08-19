package pkg

import (
	"regexp"
	"strconv"
	"strings"
	
    "github.com/GlittersIsGold/altpkgdiff/api"
)

type PackageDiff struct {
    OnlyInP10   []api.Package `json:"only_in_p10"`
    OnlyInSisyphus []api.Package `json:"only_in_sisyphus"`
    HigherInSisyphus []api.Package `json:"higher_in_sisyphus"`
}

func ComparePackages(p10Packages, sisyphusPackages []api.Package) PackageDiff {
	diffs := PackageDiff{}

	// Создаем карты для быстрого поиска по именам пакетов и их версиям
	p10Map := make(map[string]api.Package)
	sisyphusMap := make(map[string]api.Package)

	for _, pkg := range p10Packages {
		p10Map[pkg.Name] = pkg
	}

	for _, pkg := range sisyphusPackages {
		sisyphusMap[pkg.Name] = pkg
	}

	// Поиск пакетов только в p10
	for name, p10Pkg := range p10Map {
		if sisPkg, found := sisyphusMap[name]; found {
			// Сравнение версий
			if compareVersions(p10Pkg.Version, sisPkg.Version) < 0 {
				diffs.HigherInSisyphus = append(diffs.HigherInSisyphus, sisPkg)
			}
			// Удаляем найденный пакет из sisyphusMap, чтобы не повторять его в последующей проверке
			delete(sisyphusMap, name)
		} else {
			diffs.OnlyInP10 = append(diffs.OnlyInP10, p10Pkg)
		}
	}

	// Поиск пакетов только в sisyphus
	for _, pkg := range sisyphusMap {
		diffs.OnlyInSisyphus = append(diffs.OnlyInSisyphus, pkg)
	}

	return diffs
}

func compareVersions(v1, v2 string) int {
	// Преобразование версий в массив чисел и строк для сравнения
	parts1 := splitVersion(v1)
	parts2 := splitVersion(v2)

	for i := 0; i < len(parts1) && i < len(parts2); i++ {
		part1 := parts1[i]
		part2 := parts2[i]
		comp := comparePart(part1, part2)
		if comp != 0 {
			return comp
		}
	}
	return len(parts1) - len(parts2)
}

func comparePart(part1, part2 string) int {
	if isNumeric(part1) && isNumeric(part2) {
		num1, _ := strconv.Atoi(part1)
		num2, _ := strconv.Atoi(part2)
		return num1 - num2
	}
	return strings.Compare(part1, part2)
}

func isNumeric(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

func splitVersion(version string) []string {
	// Разделение версии на числа и строки
	re := regexp.MustCompile(`(\d+|\D+)`)
	return re.FindAllString(version, -1)
}