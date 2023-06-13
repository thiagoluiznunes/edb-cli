package main

import (
	"edb-assessment/consts"
	"edb-assessment/service"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func main() {

	args := os.Args
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	k8sService := service.NewK8SService()
	if err := k8sService.InitClusterConfig(config); err == nil {
		fmt.Println(err)
	}

	if len(args) > 1 {
		switch args[1] {
		case "-h":
			fmt.Println(consts.HELP_INFO_TEXT)
		case "--help":
			fmt.Println(consts.HELP_INFO_TEXT)
		case "get":
			switch args[2] {
			case "pods":
				namespace := ""
				if args[3] == "-n" && args[4] != "" {
					namespace = args[4]
				}
				k8sService.GetPods(namespace)
			case "--help":
				fmt.Println(consts.HELP_INFO_GET_PODS)
			default:
				fmt.Println(consts.UNKOWN_COMMAND)
			}
		default:
			fmt.Println(consts.UNKOWN_COMMAND)
		}
	}

}
