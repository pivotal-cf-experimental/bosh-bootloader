package commands_test

import (
	"fmt"

	"github.com/cloudfoundry/bosh-bootloader/commands"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("Commands Usage", func() {
	Describe("Up", func() {
		Describe("Usage", func() {
			It("returns string describing usage", func() {
				upCmd := commands.Up{}
				usageText := upCmd.Usage()
				Expect(usageText).To(Equal(`Deploys BOSH director on an IAAS

  --iaas                     IAAS to deploy your BOSH director onto: "aws", "azure", "gcp", "vsphere"   env: $BBL_IAAS
  --name                     Name to assign to your BOSH director (optional)                            env: $BBL_ENV_NAME

  --aws-access-key-id                AWS Access Key ID                env: $BBL_AWS_ACCESS_KEY_ID
  --aws-secret-access-key            AWS Secret Access Key            env: $BBL_AWS_SECRET_ACCESS_KEY
  --aws-region                       AWS Region                       env: $BBL_AWS_REGION

  --gcp-service-account-key          GCP Service Access Key to use    env: $BBL_GCP_SERVICE_ACCOUNT_KEY
  --gcp-region                       GCP Region to use                env: $BBL_GCP_REGION

  --azure-subscription-id            Azure Subscription ID            env: $BBL_AZURE_SUBSCRIPTION_ID
  --azure-tenant-id                  Azure Tenant ID                  env: $BBL_AZURE_TENANT_ID
  --azure-client-id                  Azure Client ID                  env: $BBL_AZURE_CLIENT_ID
  --azure-client-secret              Azure Client Secret              env: $BBL_AZURE_CLIENT_SECRET
  --azure-region                     Azure Region                     env: $BBL_AZURE_REGION
  --azure-resource-group-name        Azure Resource Group Name        env: $BBL_AZURE_RESOURCE_GROUP_NAME
  --azure-vnet-resource-group-name   Azure Vnet Resource Group Name   env: $BBL_AZURE_VNET_RESOURCE_GROUP_NAME
  --azure-vnet-name                  Azure Vnet Name                  env: $BBL_AZURE_VNET_NAME
  --azure-subnet-name                Azure Subnet Name                env: $BBL_AZURE_SUBNET_NAME
  --azure-disable-publicip           Azure Do not use public IP       env: $BBL_AZURE_DISABLE_PUBLICIP
  --azure-cidr                       Azure CIDR                       env: $BBL_AZURE_CIDR

  --vsphere-vcenter-user             vSphere vCenter User             env: $BBL_VSPHERE_VCENTER_USER
  --vsphere-vcenter-password         vSphere vCenter Password         env: $BBL_VSPHERE_VCENTER_PASSWORD
  --vsphere-vcenter-ip               vSphere vCenter IP               env: $BBL_VSPHERE_VCENTER_IP
  --vsphere-vcenter-dc               vSphere vCenter Datacenter       env: $BBL_VSPHERE_VCENTER_DC
  --vsphere-vcenter-cluster          vSphere vCenter Cluster          env: $BBL_VSPHERE_VCENTER_CLUSTER
  --vsphere-vcenter-rp               vSphere vCenter Resource Pool    env: $BBL_VSPHERE_VCENTER_RP
  --vsphere-network                  vSphere Network                  env: $BBL_VSPHERE_NETWORK
  --vsphere-vcenter-ds               vSphere vCenter Datastore        env: $BBL_VSPHERE_VCENTER_DS
  --vsphere-subnet                   vSphere Subnet                   env: $BBL_VSPHERE_SUBNET

  --openstack-internal-cidr          OpenStack Internal CIDR          env: $BBL_OPENSTACK_INTERNAL_CIDR
  --openstack-external-ip            OpenStack External IP            env: $BBL_OPENSTACK_EXTERNAL_IP
  --openstack-auth-url               OpenStack Auth URL               env: $BBL_OPENSTACK_AUTH_URL
  --openstack-az                     OpenStack Availability Zone      env: $BBL_OPENSTACK_AZ
  --openstack-default-key-name       OpenStack Default Key Name       env: $BBL_OPENSTACK_DEFAULT_KEY_NAME
  --openstack-default-security-group OpenStack Default Security Group env: $BBL_OPENSTACK_DEFAULT_SECURITY_GROUP
  --openstack-network-id             OpenStack Network ID             env: $BBL_OPENSTACK_NETWORK_ID
  --openstack-password               OpenStack Password               env: $BBL_OPENSTACK_PASSWORD
  --openstack-username               OpenStack Username               env: $BBL_OPENSTACK_USERNAME
  --openstack-project                OpenStack Project                env: $BBL_OPENSTACK_PROJECT
  --openstack-domain                 OpenStack Domain                 env: $BBL_OPENSTACK_DOMAIN
  --openstack-region                 OpenStack Region                 env: $BBL_OPENSTACK_REGION
  --openstack-private-key            OpenStack Private Key            env: $BBL_OPENSTACK_PRIVATE_KEY

  Load Balancer options:
  --lb-type                  Load balancer(s) type: "concourse" or "cf"
  --lb-cert                  Path to SSL certificate (supported when type="cf")
  --lb-key                   Path to SSL certificate key (supported when type="cf")
  --lb-chain                 Path to SSL certificate chain (supported when iaas="aws")
  --lb-domain                Creates a DNS zone and records for the given domain (supported when type="cf")`))
			})
		})
	})

	Describe("Plan", func() {
		Describe("Usage", func() {
			It("returns string describing usage", func() {
				planCmd := commands.Plan{}
				usageText := planCmd.Usage()
				Expect(usageText).To(Equal(fmt.Sprintf(`Populates a state directory with the latest config without applying it

  --iaas                     IAAS to deploy your BOSH director onto: "aws", "azure", "gcp", "vsphere"   env: $BBL_IAAS
  --name                     Name to assign to your BOSH director (optional)                            env: $BBL_ENV_NAME
%s%s`, commands.Credentials, commands.LBUsage)))
			})
		})
	})

	Describe("Destroy", func() {
		Describe("Usage", func() {
			It("returns string describing usage", func() {
				command := commands.Destroy{}
				usageText := command.Usage()
				Expect(usageText).To(Equal(fmt.Sprintf(`Tears down BOSH director infrastructure

  [--no-confirm]       Do not ask for confirmation (optional)

  Credentials for your IaaS are required:%s`, commands.Credentials)))
			})
		})
	})

	Describe("CleanupLeftovers", func() {
		It("returns string describing usage", func() {
			command := commands.CleanupLeftovers{}
			usageText := command.Usage()
			Expect(usageText).To(Equal(fmt.Sprintf(`Cleans up orphaned IAAS resources

  --filter            Only delete resources with this string in their name

  Credentials for your IaaS are required:%s`, commands.Credentials)))
		})
	})

	Describe("Rotate", func() {
		Describe("Usage", func() {
			It("returns string describing usage", func() {
				command := commands.Rotate{}
				usageText := command.Usage()
				Expect(usageText).To(Equal(fmt.Sprintf(`Rotates SSH key for the jumpbox user.

  Credentials for your IaaS are required:%s`, commands.Credentials)))
			})
		})
	})

	Describe("SSH", func() {
		Describe("Usage", func() {
			It("returns string describing usage", func() {
				command := commands.SSH{}
				usageText := command.Usage()
				Expect(usageText).To(Equal(`Opens an SSH connection to the director or the jumpbox.

  --jumpbox                Open a connection to the jumpbox
  --director               Open a connection to the director
`))
			})
		})
	})

	Describe("Usage", func() {
		Describe("Usage", func() {
			It("returns string describing usage", func() {
				command := commands.Usage{}
				usageText := command.Usage()
				Expect(usageText).To(Equal(`Prints helpful message for the given command`))
			})
		})
	})

	DescribeTable("command description", func(command commands.Command, expectedDescription string) {
		usageText := command.Usage()
		Expect(usageText).To(Equal(expectedDescription))
	},
		Entry("LBs", commands.LBs{}, "Prints attached load balancer(s)"),
		Entry("outputs", commands.Outputs{}, "Prints the outputs from terraform."),
		Entry("jumpbox-address", newStateQuery("jumpbox address"), "Prints BOSH jumpbox address"),
		Entry("director-address", newStateQuery("director address"), "Prints BOSH director address"),
		Entry("director-password", newStateQuery("director password"), "Prints BOSH director password"),
		Entry("director-username", newStateQuery("director username"), "Prints BOSH director username"),
		Entry("director-ca-cert", newStateQuery("director ca cert"), "Prints BOSH director CA certificate"),
		Entry("env-id", newStateQuery("environment id"), "Prints environment ID"),
		Entry("ssh-key", commands.SSHKey{}, "Prints SSH private key for the jumpbox."),
		Entry("director-ssh-key", commands.SSHKey{Director: true}, "Prints SSH private key for the director."),
		Entry("print-env", commands.PrintEnv{}, "Prints required BOSH environment variables"),
		Entry("latest-error", commands.LatestError{}, "Prints the output from the latest call to terraform"),
		Entry("version", commands.Version{}, "Prints version"),
	)
})

func newStateQuery(propertyName string) commands.StateQuery {
	return commands.NewStateQuery(nil, nil, nil, propertyName)
}
