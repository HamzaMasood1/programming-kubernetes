package ch1

import (
	"context"
	"flag"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func GetPodInfo() {
	ctx := context.TODO()
	kubeconfig := flag.String("kubeconfig", "/Users/hmasood/.kube/kind", "kubeconfig file")
	flag.Parse()
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		fmt.Println(err)
	}
	clientset, _ := kubernetes.NewForConfig(config)
	pod, err := clientset.CoreV1().Pods("default").Get(ctx, "nginx", metav1.GetOptions{})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(pod.ObjectMeta.Labels)

}
