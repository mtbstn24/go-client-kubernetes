package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	kubeconfig := flag.String("kubeconfig", "$USERPROFILE/.kube/config", "location to the Kube config file")
	image := flag.String("i", "", "image to be used")
	deployName := flag.String("d", "", "name of the deployment")
	replica := flag.String("r", "1", "number of replicas needed")
	flag.Parse()
	if *image == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}
	if *deployName == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}
	fmt.Printf("%s %s %s", *image, *deployName, *replica)

	configPath := filepath.Clean(os.ExpandEnv(*kubeconfig))
	config, err := clientcmd.BuildConfigFromFlags("", configPath)
	if err != nil {
		panic(err)
	}

	//Clientset is used access every resource individually. Eg: 'CoreV1' api in Clientset is used to access the Pods
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}
	pods, err := clientset.CoreV1().Pods("default").List(context.Background(), metav1.ListOptions{})
	if err != nil {
		panic(err)
	}
	fmt.Println("Default namespace Pods")
	for _, pod := range pods.Items {
		fmt.Printf("Pod name : %s \n", pod.Name)
	}

	deployments, err := clientset.AppsV1().Deployments("default").List(context.Background(), metav1.ListOptions{})
	if err != nil {
		panic(err)
	}
	fmt.Println(("\nDefault namespace deployments"))
	for _, deployment := range deployments.Items {
		fmt.Printf("Deployment name : %s \n", deployment.Name)
	}
}
