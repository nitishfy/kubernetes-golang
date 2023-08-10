package main

import (
	ctx "Kubernetes_Resource_Monitor/pkg/contexts"
	"Kubernetes_Resource_Monitor/pkg/kubectx"
	"Kubernetes_Resource_Monitor/pkg/namespace"
	"Kubernetes_Resource_Monitor/pkg/pod"
	"fmt"
)

func main() {
	fmt.Println("Listing all the contexts:")
	contexts, err := ctx.GetContexts()
	if err != nil {
		return
	}
	for _, context := range contexts {
		fmt.Println(context)
	}
	fmt.Println()
	fmt.Println("Listing all the namespaces:")
	namespaces, err := namespace.GetNamespace()
	if err != nil {
		return
	}
	for _, namespace := range namespaces {
		fmt.Println(namespace)
	}
	fmt.Println()
	fmt.Println("Listing all the Pods")
	pods, err := pod.GetPods()
	if err != nil {
		return
	}
	for _, pod := range pods {
		fmt.Println(pod)
	}
	fmt.Println()
	fmt.Println("Running kubectx command")
	results, err := kubectx.RunKubectx()
	if err != nil {
		return
	}
	for _, result := range results {
		fmt.Println(result)
	}
}
