package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Infra struct {
	Ctx  *pulumi.Context
	Name string
	FQDN string
}

func NewInfra(context *pulumi.Context, name, fqdn string) Infra {
	return Infra{
		Ctx:  context,
		Name: name,
		FQDN: fqdn,
	}
}

func (i *Infra) bucket() (*s3.BucketV2, error) {
	s3, err := s3.NewBucketV2(i.Ctx, i.Name, &s3.BucketV2Args{
		Bucket: pulumi.String(i.Name),
	})
	if err != nil {
		return nil, err
	}
	return s3, nil
}
