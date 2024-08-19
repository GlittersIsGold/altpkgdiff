package main


import (
    "encoding/json"
    "fmt"
    "os"

    "github.com/GlittersIsGold/altpkgdiff/api"
    "github.com/GlittersIsGold/altpkgdiff/pkg"
)

func main() {
    
    sisyphus_response, err := api.FetchPackages("sisyphus")
    if err != nil {
        fmt.Println("Error fetching sisyphus packages:", err)
        os.Exit(1)
    }
    sisyphusPackages := sisyphus_response.Packages;

    p10_res, err := api.FetchPackages("p10")
    if err != nil {
        fmt.Println("Error fetching p10 packages:", err)
        os.Exit(1)
    }

    p10Packages := p10_res.Packages;

    fmt.Printf("N of sisy: %d\n", len(sisyphus_response.Packages))
    fmt.Printf("N of p10: %d\n", len(p10_res.Packages))

    diffs := pkg.ComparePackages(sisyphusPackages, p10Packages)

    output, err := json.MarshalIndent(diffs, "", "  ")
    if err != nil {
        fmt.Println("Error generating JSON output:", err)
        os.Exit(1)
    }

    fmt.Println(string(output))
}