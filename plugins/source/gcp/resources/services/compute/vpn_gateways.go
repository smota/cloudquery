package compute

import (
	"context"

	"google.golang.org/api/iterator"

	pb "cloud.google.com/go/compute/apiv1/computepb"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/cloudquery/plugins/source/gcp/client"

	compute "cloud.google.com/go/compute/apiv1"
)

func VpnGateways() *schema.Table {
	return &schema.Table{
		Name:        "gcp_compute_vpn_gateways",
		Description: `https://cloud.google.com/compute/docs/reference/rest/v1/vpnGateways#VpnGateway`,
		Resolver:    fetchVpnGateways,
		Multiplex:   client.ProjectMultiplexEnabledServices("compute.googleapis.com"),
		Transform:   transformers.TransformWithStruct(&pb.VpnGateway{}, append(client.Options(), transformers.WithPrimaryKeys("SelfLink"))...),
		Columns: []schema.Column{
			{
				Name:     "project_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveProject,
			},
		},
	}
}

func fetchVpnGateways(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	req := &pb.AggregatedListVpnGatewaysRequest{
		Project: c.ProjectId,
	}
	gcpClient, err := compute.NewVpnGatewaysRESTClient(ctx, c.ClientOptions...)
	if err != nil {
		return err
	}
	it := gcpClient.AggregatedList(ctx, req, c.CallOptions...)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return err
		}

		res <- resp.Value.VpnGateways
	}
	return nil
}
