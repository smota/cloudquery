package appstream

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/appstream"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchAppstreamImages(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	input := appstream.DescribeImagesInput{MaxResults: aws.Int32(25)}
	c := meta.(*client.Client)
	svc := c.Services().Appstream
	for {
		response, err := svc.DescribeImages(ctx, &input)
		if err != nil {
			return err
		}
		res <- response.Images

		if aws.ToString(response.NextToken) == "" {
			break
		}
		input.NextToken = response.NextToken
	}
	return nil
}
