package stats

import (
	"context"
	"errors"
	"fmt"
	"time"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

type KubernetesProvider struct {
	client kubernetes.Interface
}

func NewKubernetesProvider() (*KubernetesProvider, error) {
	config, err := buildConfig()
	if err != nil {
		return nil, err
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, fmt.Errorf("create kubernetes client: %w", err)
	}

	return &KubernetesProvider{client: clientset}, nil
}

func (p *KubernetesProvider) Collect(now time.Time) (Snapshot, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	nodes, err := p.client.CoreV1().Nodes().List(ctx, metav1.ListOptions{})
	if err != nil {
		return Snapshot{}, fmt.Errorf("list nodes: %w", err)
	}

	pods, err := p.client.CoreV1().Pods(v1.NamespaceAll).List(ctx, metav1.ListOptions{})
	if err != nil {
		return Snapshot{}, fmt.Errorf("list pods: %w", err)
	}

	namespaces, err := p.client.CoreV1().Namespaces().List(ctx, metav1.ListOptions{})
	if err != nil {
		return Snapshot{}, fmt.Errorf("list namespaces: %w", err)
	}

	latestPendingPod := latestPendingPod(pods.Items, now)
	latestPendingNode := latestPendingNode(nodes.Items, now)
	allocated, allocatable := aggregateResources(nodes.Items, pods.Items)

	return Snapshot{
		GeneratedAtEpoch:            now.UTC().Unix(),
		NodeCount:                   len(nodes.Items),
		PodCount:                    len(pods.Items),
		NamespaceCount:              len(namespaces.Items),
		LatestPendingPodEpoch:       latestPendingPod,
		LatestPendingPodAgeSeconds:  ageSeconds(now, latestPendingPod),
		LatestPendingNodeEpoch:      latestPendingNode,
		LatestPendingNodeAgeSeconds: ageSeconds(now, latestPendingNode),
		AllocatedResources:          allocated,
		AllocatableResources:        allocatable,
	}, nil
}

func buildConfig() (*rest.Config, error) {
	config, err := rest.InClusterConfig()
	if err == nil {
		return config, nil
	}

	loadingRules := clientcmd.NewDefaultClientConfigLoadingRules()
	clientConfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(loadingRules, &clientcmd.ConfigOverrides{})
	kubeConfig, kubeErr := clientConfig.ClientConfig()
	if kubeErr == nil {
		return kubeConfig, nil
	}

	return nil, errors.New("unable to load kubernetes config from in-cluster config or kubeconfig")
}

func latestPendingPod(pods []v1.Pod, _ time.Time) int64 {
	var newest *v1.Pod
	for i := range pods {
		pod := &pods[i]
		if pod.Status.Phase != v1.PodPending || pod.CreationTimestamp.IsZero() {
			continue
		}

		if newest == nil || pod.CreationTimestamp.Time.After(newest.CreationTimestamp.Time) {
			newest = pod
		}
	}

	if newest == nil {
		return 0
	}

	return newest.CreationTimestamp.Time.UTC().Unix()
}

func latestPendingNode(nodes []v1.Node, _ time.Time) int64 {
	var newest *v1.Node
	for i := range nodes {
		node := &nodes[i]
		if !isNodePending(node) || node.CreationTimestamp.IsZero() {
			continue
		}

		if newest == nil || node.CreationTimestamp.Time.After(newest.CreationTimestamp.Time) {
			newest = node
		}
	}

	if newest == nil {
		return 0
	}

	return newest.CreationTimestamp.Time.UTC().Unix()
}

func ageSeconds(now time.Time, epoch int64) int64 {
	if epoch <= 0 {
		return 0
	}

	age := now.UTC().Unix() - epoch
	if age < 0 {
		return 0
	}

	return age
}

func isNodePending(node *v1.Node) bool {
	for _, condition := range node.Status.Conditions {
		if condition.Type == v1.NodeReady {
			return condition.Status == v1.ConditionUnknown
		}
	}

	return false
}

func aggregateResources(nodes []v1.Node, pods []v1.Pod) (AllocatedResources, ClusterResourceTotals) {
	var allocatable ClusterResourceTotals
	for _, node := range nodes {
		allocatable.CPUMilli += node.Status.Allocatable.Cpu().MilliValue()
		allocatable.MemoryBytes += node.Status.Allocatable.Memory().Value()
		allocatable.Pods += node.Status.Allocatable.Pods().Value()
	}

	var allocated AllocatedResources
	for _, pod := range pods {
		if pod.Spec.NodeName == "" || pod.Status.Phase == v1.PodSucceeded || pod.Status.Phase == v1.PodFailed {
			continue
		}

		for _, container := range pod.Spec.Containers {
			requests := container.Resources.Requests
			allocated.CPUMilli += requests.Cpu().MilliValue()
			allocated.MemoryBytes += requests.Memory().Value()
		}

		for _, container := range pod.Spec.InitContainers {
			requests := container.Resources.Requests
			allocated.CPUMilli += requests.Cpu().MilliValue()
			allocated.MemoryBytes += requests.Memory().Value()
		}

		if _, exists := pod.Spec.Overhead[v1.ResourceCPU]; exists {
			allocated.CPUMilli += quantityMilli(pod.Spec.Overhead[v1.ResourceCPU])
		}
		if _, exists := pod.Spec.Overhead[v1.ResourceMemory]; exists {
			allocated.MemoryBytes += quantityValue(pod.Spec.Overhead[v1.ResourceMemory])
		}

		allocated.Pods++
	}

	return allocated, allocatable
}

func quantityMilli(q resource.Quantity) int64 {
	return q.MilliValue()
}

func quantityValue(q resource.Quantity) int64 {
	return q.Value()
}
