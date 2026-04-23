package stats

import "time"

type Snapshot struct {
	GeneratedAtEpoch            int64                 `json:"generatedAtEpoch"`
	NodeCount                   int                   `json:"nodeCount"`
	PodCount                    int                   `json:"podCount"`
	NamespaceCount              int                   `json:"namespaceCount"`
	LatestPendingPodEpoch       int64                 `json:"latestPendingPodEpoch"`
	LatestPendingPodAgeSeconds  int64                 `json:"latestPendingPodAgeSeconds"`
	LatestPendingNodeEpoch      int64                 `json:"latestPendingNodeEpoch"`
	LatestPendingNodeAgeSeconds int64                 `json:"latestPendingNodeAgeSeconds"`
	AllocatedResources          AllocatedResources    `json:"allocatedResources"`
	AllocatableResources        ClusterResourceTotals `json:"allocatableResources"`
}

type AllocatedResources struct {
	CPUMilli    int64 `json:"cpuMilli"`
	MemoryBytes int64 `json:"memoryBytes"`
	Pods        int64 `json:"pods"`
}

type ClusterResourceTotals struct {
	CPUMilli    int64 `json:"cpuMilli"`
	MemoryBytes int64 `json:"memoryBytes"`
	Pods        int64 `json:"pods"`
}

type Provider interface {
	Collect(now time.Time) (Snapshot, error)
}
