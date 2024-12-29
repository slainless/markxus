package main

import (
	"context"
	"log"
	"os"

	"github.com/slainless/markxus/cli/command"
)

func main() {
	if err := command.Main.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}

// func main() {
// 	nexusKey := nexus.WithApiKey("75K+pcclj+XVwg8GM5jTrMLXcHMBARdZ9RcDbL3pb+nCvfyS48D6--KlL2Ej4HSG68hbh2--4tVrTFa6dIytXQex2/K0Gw==")
// 	nexus, err := nexus.NewClient(nexusKey, nexus.WithHTTPDriver(resty.NewRestyClient()))
// 	if err != nil {
// 		panic(err)
// 	}

// 	genaiKey := genai.WithApiKey("AIzaSyAgMo6REA1M37g-GclUt9NXYcezGkRNQqw")
// 	genai, err := genai.NewGenAiClient(context.TODO(), genaiKey)
// 	if err != nil {
// 		panic(err)
// 	}

// 	markxus := markxus.NewMarkxus(nexus, genai)
// 	mod, err := markxus.Generate(context.TODO(), "skyrimspecialedition", "68068")
// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Printf("%+v", mod)
// }
