package main

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func DeployInfra(ctx *pulumi.Context) error {
	infra := Infra{
		Ctx:  ctx,
		Name: "demo-daveops-pulumi",
		FQDN: "demo.papelon.store",
	}

	s3, err := infra.bucket()
	if err != nil {
		return err
	}

	ctx.Export(infra.Name, s3.ID())
	return nil
}

/*
func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Create an AWS resource (S3 Bucket)
		err := DeployInfra(ctx)
		if err != nil {
			return err
		}
		return nil
	})
}*/
