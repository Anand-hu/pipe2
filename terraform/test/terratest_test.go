package test

import (
  "testing"

  "github.com/gruntwork-io/terratest/modules/terraform"
  "github.com/gruntwork-io/terratest/modules/aws"
  "github.com/stretchr/testify/assert"
)

func TestVpcModule(t *testing.T) {
  t.Parallel()

  terraformOptions := &terraform.Options{
    TerraformDir: "../../modules/vpc",
  }

  defer terraform.Destroy(t, terraformOptions)
  terraform.InitAndApply(t, terraformOptions)

  vpcID := terraform.Output(t, terraformOptions, "vpc_id")
  assert.NotEmpty(t, vpcID)

  region := "us-east-1" // change if using different region
  actualCIDR := aws.GetVpcCidrBlock(t, vpcID, region)
  assert.Contains(t, actualCIDR, "10.0.0.0/16")
}

func TestAlbModule(t *testing.T) {
  t.Parallel()

  terraformOptions := &terraform.Options{
    TerraformDir: "../../modules/alb",
  }

  defer terraform.Destroy(t, terraformOptions)
  terraform.InitAndApply(t, terraformOptions)

  albDNS := terraform.Output(t, terraformOptions, "alb_dns_name")
  assert.NotEmpty(t, albDNS)
}

func TestEcsModule(t *testing.T) {
  t.Parallel()

  terraformOptions := &terraform.Options{
    TerraformDir: "../../modules/ecs",
  }

  defer terraform.Destroy(t, terraformOptions)
  terraform.InitAndApply(t, terraformOptions)

  serviceName := terraform.Output(t, terraformOptions, "ecs_service_name")
  assert.Contains(t, serviceName, "service")
}
