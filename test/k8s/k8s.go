package main

import (
	"flag"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"log"
)

var clientset *kubernetes.Clientset

func main() {
	K8sClient()
}

func K8sClient() {
	log.Println("TestK8sClient success")

	k8sconfig := flag.String("k8sconfig", "./config", "kubernetes config file path")
	flag.Parse()
	config, err := clientcmd.BuildConfigFromFlags("", *k8sconfig)
	if err != nil {
		log.Println(err)
	}
	clientset, err = kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatalln(err)
	} else {
		log.Println("connect k8s success")
	}

	//获取POD
	pods, err := clientset.CoreV1().Pods("kube-system").List(metav1.ListOptions{})

	for _, val := range pods.Items {
		log.Println("pods：" + val.Name)
	}

	namespaces, err := clientset.CoreV1().Namespaces().List(metav1.ListOptions{})
	for _, val := range namespaces.Items {
		log.Println("namespaces：" + val.Name)
	}
}
