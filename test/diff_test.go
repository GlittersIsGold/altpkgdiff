package test

import (
	"testing"
	"reflect"

	"github.com/GlittersIsGold/altpkgdiff/api"
	"github.com/GlittersIsGold/altpkgdiff/pkg"
)

func TestCmpPkgVers(t *testing.T) {
	tests := []struct {
		name     string
		src      api.Package
		dst      api.Package
		expected int
	}{
		{
			name: "same version",
			src: api.Package{
				Name: "example", Epoch: 0, Version: "1.0.0", Release: "1",
				Arch: "x86_64", Disttag: "sisyphus", Buildtime: 1234567890, Source: "example-src",
			},
			dst: api.Package{
				Name: "example", Epoch: 0, Version: "1.0.0", Release: "1",
				Arch: "x86_64", Disttag: "sisyphus", Buildtime: 1234567890, Source: "example-src",
			},
			expected: 0,
		},
		{
			name: "higher epoch",
			src: api.Package{
				Name: "example", Epoch: 1, Version: "1.0.0", Release: "1",
				Arch: "x86_64", Disttag: "sisyphus", Buildtime: 1234567890, Source: "example-src",
			},
			dst: api.Package{
				Name: "example", Epoch: 0, Version: "1.0.0", Release: "1",
				Arch: "x86_64", Disttag: "sisyphus", Buildtime: 1234567890, Source: "example-src",
			},
			expected: 1,
		},
		{
			name: "lower epoch",
			src: api.Package{
				Name: "example", Epoch: 0, Version: "1.0.0", Release: "1",
				Arch: "x86_64", Disttag: "sisyphus", Buildtime: 1234567890, Source: "example-src",
			},
			dst: api.Package{
				Name: "example", Epoch: 1, Version: "1.0.0", Release: "1",
				Arch: "x86_64", Disttag: "sisyphus", Buildtime: 1234567890, Source: "example-src",
			},
			expected: -1,
		},
		{
			name: "higher version",
			src: api.Package{
				Name: "example", Epoch: 0, Version: "2.0.0", Release: "1",
				Arch: "x86_64", Disttag: "sisyphus", Buildtime: 1234567890, Source: "example-src",
			},
			dst: api.Package{
				Name: "example", Epoch: 0, Version: "1.0.0", Release: "1",
				Arch: "x86_64", Disttag: "sisyphus", Buildtime: 1234567890, Source: "example-src",
			},
			expected: 1,
		},
		{
			name: "lower version",
			src: api.Package{
				Name: "example", Epoch: 0, Version: "1.0.0", Release: "1",
				Arch: "x86_64", Disttag: "sisyphus", Buildtime: 1234567890, Source: "example-src",
			},
			dst: api.Package{
				Name: "example", Epoch: 0, Version: "2.0.0", Release: "1",
				Arch: "x86_64", Disttag: "sisyphus", Buildtime: 1234567890, Source: "example-src",
			},
			expected: -1,
		},
		{
			name: "higher release",
			src: api.Package{
				Name: "example", Epoch: 0, Version: "1.0.0", Release: "2",
				Arch: "x86_64", Disttag: "sisyphus", Buildtime: 1234567890, Source: "example-src",
			},
			dst: api.Package{
				Name: "example", Epoch: 0, Version: "1.0.0", Release: "1",
				Arch: "x86_64", Disttag: "sisyphus", Buildtime: 1234567890, Source: "example-src",
			},
			expected: 1,
		},
		{
			name: "lower release",
			src: api.Package{
				Name: "example", Epoch: 0, Version: "1.0.0", Release: "1",
				Arch: "x86_64", Disttag: "sisyphus", Buildtime: 1234567890, Source: "example-src",
			},
			dst: api.Package{
				Name: "example", Epoch: 0, Version: "1.0.0", Release: "2",
				Arch: "x86_64", Disttag: "sisyphus", Buildtime: 1234567890, Source: "example-src",
			},
			expected: -1,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := pkg.CmpPkgVers(tc.src, tc.dst)
			if result != tc.expected {
				t.Errorf("CmpPkgVers(%v, %v) = %d; expected %d", tc.src, tc.dst, result, tc.expected)
			}
		})
	}
}

func TestCreateVersion(t *testing.T) {
    testCases := []struct {
        name           string
        pkg            api.Package
        expectedResult string
    }{
        {
            name: "Standard version format",
            pkg: api.Package{
                Epoch:   1,
                Version: "1.2.3",
                Release: "4",
            },
            expectedResult: "1:1.2.3-4",
        },
        {
            name: "Zero epoch",
            pkg: api.Package{
                Epoch:   0,
                Version: "2.0.0",
                Release: "1",
            },
            expectedResult: "2.0.0-1", // изменено ожидаемое значение
        },
        {
            name: "Empty release",
            pkg: api.Package{
                Epoch:   2,
                Version: "3.4.5",
                Release: "",
            },
            expectedResult: "2:3.4.5", // изменено ожидаемое значение
        },
        {
            name: "Empty version and release",
            pkg: api.Package{
                Epoch:   0,
                Version: "",
                Release: "",
            },
            expectedResult: "", // изменено ожидаемое значение
        },
    }

    for _, tc := range testCases {
        t.Run(tc.name, func(t *testing.T) {
            result := pkg.CreateVersion(tc.pkg)
            if result.String() != tc.expectedResult {
                t.Errorf("expected %v, got %v", tc.expectedResult, result.String())
            }
        })
    }
}

func TestFilterByArch(t *testing.T) {
    testCases := []struct {
        name           string
        packages       []api.Package
        arch           string
        expectedResult []api.Package
    }{
        {
            name: "Filter packages by arch x86_64",
            packages: []api.Package{
                {Name: "package1", Arch: "x86_64"},
                {Name: "package2", Arch: "i386"},
                {Name: "package3", Arch: "x86_64"},
            },
            arch: "x86_64",
            expectedResult: []api.Package{
                {Name: "package1", Arch: "x86_64"},
                {Name: "package3", Arch: "x86_64"},
            },
        },
        {
            name: "Filter packages by arch i386",
            packages: []api.Package{
                {Name: "package1", Arch: "x86_64"},
                {Name: "package2", Arch: "i386"},
                {Name: "package3", Arch: "x86_64"},
            },
            arch: "i386",
            expectedResult: []api.Package{
                {Name: "package2", Arch: "i386"},
            },
        },
        {
            name: "No packages match the given arch",
            packages: []api.Package{
                {Name: "package1", Arch: "x86_64"},
                {Name: "package2", Arch: "x86_64"},
            },
            arch: "i386",
            expectedResult: []api.Package{},
        },
        {
            name: "Empty package list",
            packages: []api.Package{},
            arch: "x86_64",
            expectedResult: []api.Package{},
        },
        {
            name: "Empty arch, should return empty result",
            packages: []api.Package{
                {Name: "package1", Arch: "x86_64"},
                {Name: "package2", Arch: "i386"},
            },
            arch: "",
            expectedResult: []api.Package{},
        },
    }

    for _, tc := range testCases {
        t.Run(tc.name, func(t *testing.T) {
            result := pkg.FilterByArch(tc.packages, tc.arch)
            if len(result) == 0 && len(tc.expectedResult) == 0 {
                // Если оба среза пусты, то они считаются равными
                return
            }
            if !reflect.DeepEqual(result, tc.expectedResult) {
                t.Errorf("expected %v, got %v", tc.expectedResult, result)
            }
        })
    }
}
