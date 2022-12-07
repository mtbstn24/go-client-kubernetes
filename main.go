package main

import (
	"context"
	"flag"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	kubeconfig := flag.String("kubeconfig", "C:/Users/WSO2/.kube/config", "location to the Kube config file")
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		// error handle
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		// error handle
	}
	pods, err := clientset.CoreV1().Pods("default").List(context.Background(), metav1.ListOptions{})
	if err != nil {
		// error handle
	}
	fmt.Println("Default namespace Pods")
	for _, pod := range pods.Items {
		fmt.Printf("Pod name : %s \n", pod.Name)
	}

	deployments, err := clientset.AppsV1().Deployments("default").List(context.Background(), metav1.ListOptions{})
	fmt.Println(("\nDefault namespace deployments"))
	for _, deployment := range deployments.Items {
		fmt.Printf("Deployment name : %s \n", deployment.Name)
	}
}
