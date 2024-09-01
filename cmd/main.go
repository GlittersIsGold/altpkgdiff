package main

import (
    "encoding/json"
    "bufio"
    "flag"
    "fmt"
    "os"

    "github.com/GlittersIsGold/altpkgdiff/api"
    "github.com/GlittersIsGold/altpkgdiff/pkg"
)

func main() {

    src := flag.String("src","","sisyphus");
    dst := flag.String("dst","","p10");

    flag.Parse()

    if *src == "" || *dst == "" {
        fmt.Println("Both flags -src & -dst should be used.")
        flag.Usage() 
        os.Exit(1)
    }
    
    srcResponce, err := api.FetchPackages(*src)
    if err != nil {
        fmt.Printf("Error fetching packages for %s: %v\n", *src, err)
        os.Exit(1)
    }
    
    srcPackages := srcResponce.Packages;

    dstRespone, err := api.FetchPackages(*dst)
    if err != nil {
        fmt.Printf("Error fetching packages for %s: %v\n", *dst, err)
        os.Exit(1)
    }

    dstPackages := dstRespone.Packages;

    diffs := pkg.DiffPkgs(srcPackages, dstPackages)

    output, err := json.MarshalIndent(diffs, "", "  ")
    if err != nil {
        fmt.Println("Error generating JSON output:", err)
        os.Exit(1)
    }

    fmt.Println(string(output))
    
    fmt.Print("Press enter to exit")
    reader := bufio.NewReader(os.Stdin)
    _, _ = reader.ReadString('\n') 

    fmt.Println("Key pressed")
}