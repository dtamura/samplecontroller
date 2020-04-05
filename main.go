package main

import (
	clientset "github.com/dtamura/samplecontroller/pkg/client/clientset/versioned"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"

	// Uncomment the following line to load the gcp plugin (only required to authenticate against GKE clusters).
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"

	"fmt"
	"log"
	"os"
	"time"

	"github.com/urfave/cli"
)

func main() {
	log.SetFlags(log.Lshortfile | log.Ldate | log.Ltime)

	app := cli.NewApp()
	app.Flags = []cli.Flag{}

	app.Action = func(c *cli.Context) error {
		return action(c)
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func action(c *cli.Context) error {
	cfg, err := rest.InClusterConfig()
	if err != nil {
		log.Printf(err.Error())
	}
	client, err := clientset.NewForConfig(cfg)
	if err != nil {
		log.Printf(err.Error())
	}

	for {
		foos, err := client.SamplecontrollerV1alpha().Foos("default").List(c.Context, v1.ListOptions{})
		if err != nil {
			log.Printf(err.Error())
			continue
		}

		fmt.Printf("%+v\n", foos)

		time.Sleep(10 * time.Second)
	}
	return nil
}
