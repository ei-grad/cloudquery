// Code generated by codegen; DO NOT EDIT.

package domains

import (
	"context"
	"github.com/pkg/errors"
	"google.golang.org/api/iterator"

	pb "google.golang.org/genproto/googleapis/cloud/domains/v1beta1"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugins/source/gcp/client"

	"fmt"
)

func Registrations() *schema.Table {
	return &schema.Table{
		Name:      "gcp_domains_registrations",
		Resolver:  fetchRegistrations,
		Multiplex: client.ProjectMultiplex,
		Columns: []schema.Column{
			{
				Name:     "project_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveProject,
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "domain_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DomainName"),
			},
			{
				Name:     "create_time",
				Type:     schema.TypeTimestamp,
				Resolver: client.ResolveProtoTimestamp("CreateTime"),
			},
			{
				Name:     "expire_time",
				Type:     schema.TypeTimestamp,
				Resolver: client.ResolveProtoTimestamp("ExpireTime"),
			},
			{
				Name:     "state",
				Type:     schema.TypeString,
				Resolver: client.ResolveProtoEnum("State"),
			},
			{
				Name:     "issues",
				Type:     schema.TypeIntArray,
				Resolver: schema.PathResolver("Issues"),
			},
			{
				Name:     "labels",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Labels"),
			},
			{
				Name:     "management_settings",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ManagementSettings"),
			},
			{
				Name:     "dns_settings",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("DnsSettings"),
			},
			{
				Name:     "contact_settings",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ContactSettings"),
			},
			{
				Name:     "pending_contact_settings",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("PendingContactSettings"),
			},
			{
				Name:     "supported_privacy",
				Type:     schema.TypeIntArray,
				Resolver: schema.PathResolver("SupportedPrivacy"),
			},
		},
	}
}

func fetchRegistrations(ctx context.Context, meta schema.ClientMeta, r *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	req := &pb.ListRegistrationsRequest{
		Parent: fmt.Sprintf("projects/%s/locations/-", c.ProjectId),
	}
	it := c.Services.DomainsClient.ListRegistrations(ctx, req)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return errors.WithStack(err)
		}

		res <- resp

	}
	return nil
}
