//go:build 2405

package api

import (
	"github.com/lcrownover/prometheus-slurm-exporter/internal/types"
)

var versionedEndpoints = []endpoint{
	{types.ApiJobsEndpointKey, "/slurm/v0.0.41/jobs"},
	{types.ApiNodesEndpointKey, "/slurm/v0.0.41/nodes"},
	{types.ApiPartitionsEndpointKey, "/slurm/v0.0.41/partitions"},
	{types.ApiDiagEndpointKey, "/slurm/v0.0.41/diag"},
	{types.ApiSharesEndpointKey, "/slurm/v0.0.41/shares"},
}
