package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	appsv1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	kubeconfig := flag.String("kubeconfig", "$USERPROFILE/.kube/config", "location to the Kube config file")
	image := flag.String("i", "", "image to be used")
	deployName := flag.String("d", "", "name of the deployment")
	replica := flag.Int("r", 1, "number of replicas needed")
	port := flag.Int("p", 80, "port number to be used")
	flag.Parse()
	if *image == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}
	if *deployName == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}
	rep := int32(*replica)
	fmt.Printf("%v %v %v %d", *image, *deployName, *replica, *port)

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
	fmt.Println("\nDefault namespace Pods")
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

	deployClient := clientset.AppsV1().Deployments("default")
	deployment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name: *deployName,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: &rep,
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": *deployName,
				},
			},
			Template: apiv1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": *deployName,
					},
				},
				Spec: apiv1.PodSpec{
					Containers: []apiv1.Container{
						{
							Name:  "web",
							Image: *image,
							Ports: []apiv1.ContainerPort{
								{
									Name:          "http",
									Protocol:      apiv1.ProtocolTCP,
									ContainerPort: int32(*port),
								},
							},
						},
					},
				},
			},
		},
	}

	serviceName := *deployName + "svc"
	serviceClient := clientset.CoreV1().Services("default")
	service := &apiv1.Service{
		TypeMeta: metav1.TypeMeta{
			Kind: "Service",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: serviceName,
		},
		Spec: apiv1.ServiceSpec{
			Selector: map[string]string{
				"app": *deployName,
			},
			Ports: []apiv1.ServicePort{
				{
					Name:       "http",
					Protocol:   apiv1.ProtocolTCP,
					Port:       int32(*port),
					TargetPort: intstr.FromInt(*port),
				},
			},
		},
	}

	fmt.Println("Creating Deployment....")
	result, err := deployClient.Create(context.TODO(), deployment, metav1.CreateOptions{})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Created Deployment %q \n", result.GetObjectMeta().GetName())

	fmt.Println("Creating Service....")
	result1, err := serviceClient.Create(context.TODO(), service, metav1.CreateOptions{})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Created Service %q \n", result1.GetName())

}
