package main

import (
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"os"
	"path/filepath"
)

type K8sClient struct {
	ClientSet     *kubernetes.Clientset
	DynamicClient *dynamic.Interface
}

func BuildConfig() (*rest.Config, error) {
	var kubeconfig string
	kubeconfig = filepath.Join("/Users/mujian/GolandProjects/go-playground/frame/client-go")
	os.ReadFile(kubeconfig)
	cfg, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}

func NewClient() (*K8sClient, error) {
	config, err := BuildConfig()
	if err != nil {
		return nil, err
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}
	dynamicClient, err := dynamic.NewForConfig(config)
	if err != nil {
		return nil, err
	}
	return &K8sClient{clientset, &dynamicClient}, nil
}

func InitClient(clusterName string) (*kubernetes.Clientset, *rest.Config, error) {
	//数据库取出集群信息
	master, kubeconfig, err := GetClusterInfo(clusterName)
	if err != nil {
		logs.Error("get db for cluster kubeconfig error. %v ", err)
		return nil, nil, err
	}
	kubeconfigJson, err := yaml.YAMLToJSON([]byte(kubeconfig))
	if err != nil {
		logs.Error("yaml to json err")
	}
	configV1 := clientcmdapiv1.Config{}
	err = json.Unmarshal(kubeconfigJson, &configV1)
	if err != nil {
		logs.Error("json unmarshal kubeconfig error. %v ", err)
		return nil, nil, err
	}
	// 切换匹配的版本
	configObject, err := clientcmdlatest.Scheme.ConvertToVersion(&configV1, clientcmdapi.SchemeGroupVersion)
	if err != nil {
		logs.Error("ConvertToVersion error. %v ", err)
		return nil, nil, err
	}
	configInternal := configObject.(*clientcmdapi.Config)

	// 实例化配置信息
	clientConfig, err := clientcmd.NewDefaultClientConfig(*configInternal, &clientcmd.ConfigOverrides{
		ClusterDefaults: clientcmdapi.Cluster{Server: master},
	}).ClientConfig()

	if err != nil {
		logs.Error("build client config error. %v ", err)
		return nil, nil, err
	}
	clientConfig.QPS = defaultQPS
	clientConfig.Burst = defaultBurst
	// 实例化客户端
	clientSet, err := kubernetes.NewForConfig(clientConfig)

	if err != nil {
		logs.Error("(%s) kubernetes.NewForConfig(%v) error.%v", master, err, clientConfig)
		return nil, nil, err
	}
	return clientSet, clientConfig, nil

}

func GetOutClusterClient(name string) (*K8sClient, error) {
	clientSet, config, err := InitClient(name)
	if err != nil {
		return nil, err
	}
	return &K8sClient{Clientset: clientSet, Config: config}, nil
}
