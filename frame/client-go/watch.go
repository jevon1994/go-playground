package main

import (
	"context"
	"flag"
	"fmt"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/dynamic/dynamicinformer"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"path/filepath"
)

func main() {
	//if len(os.Args) < 2 {
	//	fmt.Println("resource missing")
	//	return
	//}
	// 资源，比如 "configmaps.v1.", "deployments.v1.apps", "rabbits.v1.stable.wbsnail.com"
	resource := "deployments.v1.apps"

	//kubeconfig := os.Getenv("KUBECONFIG")

	var kubeconfig *string

	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "[可选] kubeconfig 绝对路径")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "kubeconfig 绝对路径")
	}

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	Must(err)
	// 注意创建了 dynamicClient, 而不是 clientset
	dynamicClient, err := dynamic.NewForConfig(config)

	// 实例化clientset对象
	clientset, err := kubernetes.NewForConfig(config)

	ListNameSpaces(clientset)
	Must(err)
	// 同样这里也是 DynamicSharedInformerFactory, 而不是 SharedInformerFactory
	informerFactory := dynamicinformer.NewFilteredDynamicSharedInformerFactory(
		dynamicClient, 0, "default", nil)

	// 通过 schema 包提供的 ParseResourceArg 由资源描述字符串解析出 GroupVersionResource
	gvr, _ := schema.ParseResourceArg(resource)
	if gvr == nil {
		fmt.Println("cannot parse gvr")
		return
	}
	// 使用 gvr 动态生成 Informer
	informer := informerFactory.ForResource(*gvr).Informer()
	// 熟悉的代码，熟悉的味道，只是收到的 obj 类型好像不太一样
	informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		onAdd,
		onUpdate,
		onDelete,
	})

	stopCh := make(chan struct{})
	defer close(stopCh)

	fmt.Println("Start syncing....")

	go informerFactory.Start(stopCh)

	<-stopCh

}

func onAdd(obj interface{}) {
	// *unstructured.Unstructured 类是所有 Kubernetes 资源类型公共方法的抽象，
	// 提供所有对公共属性的访问方法，像 GetName, GetNamespace, GetLabels 等等，
	s, ok := obj.(*unstructured.Unstructured)
	if !ok {
		return
	}
	fmt.Printf("created: %s\n", s.GetName())
}

func onUpdate(oldObj, newObj interface{}) {
	oldS, ok1 := oldObj.(*unstructured.Unstructured)
	newS, ok2 := newObj.(*unstructured.Unstructured)
	if !ok1 || !ok2 {
		return
	}
	// 要访问公共属性外的字段，可以借助 unstructured 包提供的一些助手方法：
	oldColor, ok1, err1 := unstructured.NestedSlice(oldS.Object, "status", "conditions")
	newColor, ok2, err2 := unstructured.NestedSlice(newS.Object, "status", "conditions")
	if !ok1 || !ok2 || err1 != nil || err2 != nil {
		fmt.Printf("updated: %s\n", newS.GetName())
		return
	}
	for i, old := range oldColor {
		m := old.(map[string]interface{})
		n := newColor[i].(map[string]interface{})
		oldVal := m["type"]
		newVal := n["type"]
		fmt.Printf("updated: %s,instance: %d, old: %s, new: %s\n", newS.GetName(), i+1, oldVal, newVal)
	}

}

func onDelete(obj interface{}) {
	s, ok := obj.(*unstructured.Unstructured)
	if !ok {
		return
	}
	fmt.Printf("deleted: %s\n", s.GetName())
}

func Must(e interface{}) {
	if e != nil {
		panic(e)
	}
}

func buildConfigFromFlags(context, kubeconfigPath string) (*rest.Config, error) {
	return clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		&clientcmd.ClientConfigLoadingRules{ExplicitPath: kubeconfigPath},
		&clientcmd.ConfigOverrides{
			CurrentContext: context,
		}).ClientConfig()
}

func ListNameSpaces(coreClient kubernetes.Interface) {

	nsList, err := coreClient.CoreV1().
		Namespaces().
		List(context.Background(), v1.ListOptions{})
	//checkErr(err)
	fmt.Println(err)

	for _, n := range nsList.Items {
		fmt.Println(n.Name)
	}
}
