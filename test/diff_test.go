package test

import (
	"testing"

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

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := pkg.CmpPkgVers(tt.src, tt.dst)
			if result != tt.expected {
				t.Errorf("CmpPkgVers(%v, %v) = %d; expected %d", tt.src, tt.dst, result, tt.expected)
			}
		})
	}
}
