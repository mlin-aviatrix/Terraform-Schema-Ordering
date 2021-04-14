package basic

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceAviatrixGateway() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		SchemaVersion: 1,
		Schema: map[string]*schema.Schema{
			"cloud_type": {
				Type:         schema.TypeInt,
				Required:     true,
				Description:  "Type of cloud service provider.",
				ValidateFunc: validateCloudType, // want "schema ordering"
			},
			"account_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Account name. This account will be used to launch Aviatrix gateway.",
			},
			"gw_name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Aviatrix gateway unique name.",
			},
			"vpc_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "ID of legacy VPC/Vnet to be connected.",
			},
			"vpc_reg": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Region where this gateway will be launched.",
			},
			"gw_size": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Size of Gateway Instance.",
			},
			"subnet": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "A VPC Network address range selected from one of the available network ranges.",
			},
			"zone": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Availability Zone. Only available for Azure and Public Subnet Filtering gateway",
			},
			"insane_mode_az": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "AZ of subnet being created for Insane Mode Gateway. Required if insane_mode is set.",
			},
			"single_ip_snat": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Enable Source NAT for this container.",
			},
			"vpn_access": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Enable user access through VPN to this container.",
			},
			"vpn_cidr": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "VPN CIDR block for the container.",
			},
			"enable_elb": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Specify whether to enable ELB or not.",
			},
			"split_tunnel": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				Description: "Specify split tunnel mode.",
			},
			"max_vpn_conn": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "Maximum connection of VPN access.",
			},
			"name_servers": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
				Description: "A list of DNS servers used to resolve domain names by " +
					"a connected VPN user when Split Tunnel Mode is enabled.",
			},
			"search_domains": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
				Description: "A list of domain names that will use the NameServer " +
					"when a specific name is not in the destination when Split Tunnel Mode is enabled.",
			},
			"additional_cidrs": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
				Description: "A list of destination CIDR ranges that will also go through the VPN tunnel " +
					"when Split Tunnel Mode is enabled.",
			},
			"otp_mode": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "Two step authentication mode.",
			},
			"saml_enabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "This field indicates whether to enable SAML or not.",
			},
			"enable_vpn_nat": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				Description: "This field indicates whether to enable VPN NAT or not. Only supported for VPN gateway. Valid values: true, false. Default value: true.",
			},
			"okta_token": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Sensitive:   true,
				Description: "Token for Okta auth mode.",
			},
			"okta_url": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "URL for Okta auth mode.",
			},
			"okta_username_suffix": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "Username suffix for Okta auth mode.",
			},
			"duo_integration_key": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "Integration key for DUO auth mode.",
			},
			"duo_secret_key": {
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				Default:     "",
				Description: "Secret key for DUO auth mode.",
			},
			"duo_api_hostname": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "API hostname for DUO auth mode.",
			},
			"duo_push_mode": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "Push mode for DUO auth.",
			},
			"enable_ldap": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Specify whether to enable LDAP or not. Supported values: 'yes' and 'no'.",
			},
			"ldap_server": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "LDAP server address. Required: Yes if enable_ldap is 'yes'.",
			},
			"ldap_bind_dn": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "LDAP bind DN. Required: Yes if enable_ldap is 'yes'.",
			},
			"ldap_password": {
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				Default:     "",
				Description: "LDAP password. Required: Yes if enable_ldap is 'yes'.",
			},
			"ldap_base_dn": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "LDAP base DN. Required: Yes if enable_ldap is 'yes'.",
			},
			"ldap_username_attribute": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "LDAP user attribute. Required: Yes if enable_ldap is 'yes'.",
			},
			"peering_ha_subnet": {
				Type:     schema.TypeString,
				Optional: true,
				Description: "Public Subnet Information while creating Peering HA Gateway, only subnet is accepted. " +
					"Required to create peering ha gateway if cloud_type = 1 or 8 (AWS or AZURE). Optional if cloud_type = 4 (GCP)",
			},
			"peering_ha_zone": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "Zone information for creating Peering HA Gateway. Required to create peering ha gateway if cloud_type = 4 (GCP). Optional for cloud_type = 8 (AZURE).",
			},
			"peering_ha_insane_mode_az": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "AZ of subnet being created for Insane Mode Peering HA Gateway. Required if insane_mode is set.",
			},
			"peering_ha_gw_size": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "Peering HA Gateway Size.",
			},
			"single_az_ha": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Set to true if this feature is desired.",
			},
			"allocate_new_eip": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
				Description: "When value is false, reuse an idle address in Elastic IP pool for this gateway. " +
					"Otherwise, allocate a new Elastic IP and use it for this gateway.",
			},
			"tag_list": {
				Type:        schema.TypeList,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Optional:    true,
				Default:     nil,
				Deprecated:  "Use tags instead.",
				Description: "Instance tag of cloud provider.",
			},
			"insane_mode": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Enable Insane Mode for Gateway. Valid values: true, false.",
			},
			"enable_vpc_dns_server": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Enable vpc_dns_server for Gateway. Only supports AWS. Valid values: true, false.",
			},
			"enable_designated_gateway": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Enable 'designated_gateway' feature for Gateway. Only supports AWS and AWSGOV. Valid values: true, false.",
			},
			"additional_cidrs_designated_gateway": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "A list of CIDR ranges separated by comma to configure when 'designated_gateway' feature is enabled.",
			},
			"enable_encrypt_volume": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Enable encrypt gateway EBS volume. Only supported for AWS provider. Valid values: true, false. Default value: false.",
			},
			"customer_managed_keys": {
				Type:         schema.TypeString,
				Optional:     true,
				Sensitive:    true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "Customer managed key ID.",
			},
			"enable_monitor_gateway_subnets": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Enable monitor gateway subnets. Valid values: true, false. Default value: false.",
			},
			"monitor_exclude_list": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "A set of monitored instance ids. Only valid when 'enable_monitor_gateway_subnets' = true.",
			},
			"idle_timeout": {
				Type:         schema.TypeInt,
				Optional:     true,
				Default:      -1,
				ValidateFunc: validation.IntAtLeast(301),
				Description:  "Typed value when modifying idle_timeout. If it's -1, this feature is disabled.",
			},
			"renegotiation_interval": {
				Type:         schema.TypeInt,
				Optional:     true,
				Default:      -1,
				ValidateFunc: validation.IntAtLeast(301),
				Description:  "Typed value when modifying renegotiation_interval. If it's -1, this feature is disabled.",
			},
			"fqdn_lan_cidr": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				ForceNew:    true,
				Description: "FQDN gateway lan interface cidr.",
			},
			"fqdn_lan_vpc_id": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "LAN VPC ID. Only used for GCP FQDN Gateway.",
			},
			"enable_public_subnet_filtering": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
				RequiredWith: []string{
					"public_subnet_filtering_route_tables",
					"public_subnet_filtering_guard_duty_enforced",
				},
				Description:   "Create a [Public Subnet Filtering gateway](https://docs.aviatrix.com/HowTos/public_subnet_filtering_faq.html).",
			},
			"public_subnet_filtering_route_tables": {
				Type:        schema.TypeSet,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Optional:    true,
				Description: "Route tables whose associated public subnets are protected. Required when `enable_public_subnet_filtering` attribute is true.",
			},
			"public_subnet_filtering_ha_route_tables": {
				Type:        schema.TypeSet,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Optional:    true,
				Description: "Route tables whose associated public subnets are protected for the HA PSF gateway. Required when enable_public_subnet_filtering and peering_ha_subnet are set.",
			},
			"public_subnet_filtering_guard_duty_enforced": {
				Type:        schema.TypeBool,
				Default:     true,
				Optional:    true,
				Description: "Whether to enforce Guard Duty IP blocking. Required when `enable_public_subnet_filtering` attribute is true. Valid values: true or false. Default value: true.",
			},
			"enable_jumbo_frame": {
				Type:        schema.TypeBool,
				Default:     true,
				Optional:    true,
				Description: "Enable jumbo frame support for Gateway. Valid values: true or false. Default value: true.",
			},
			"tags": {
				Type:          schema.TypeMap,
				Elem:          &schema.Schema{Type: schema.TypeString},
				Optional:      true,
				Description:   "A map of tags to assign to the gateway.",
				ConflictsWith: []string{"tag_list"},
			},
			"storage_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Name of storage account with gateway images. Only valid for Azure China (2048)",
			},
			"eip": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Required when allocate_new_eip is false. It uses specified EIP for this gateway.",
			},
			"peering_ha_eip": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Public IP address that you want assigned to the HA peering instance.",
			},
			"elb_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "A name for the ELB that is created.",
			},
			"vpn_protocol": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validation.StringInSlice([]string{"TCP", "UDP"}, false),
				Description: "Elb protocol for VPN gateway with elb enabled. Only supports AWS provider. " +
					"Valid values: 'TCP', 'UDP'. If not specified, 'TCP'' will be used.",
			},
			"tunnel_detection_time": {
				Type:         schema.TypeInt,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validation.IntBetween(20, 600),
				Description:  "The IPSec tunnel down detection time for the Gateway.",
			},
			"elb_dns_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "ELB DNS Name.",
			},
			"security_group_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Security group used for the gateway.",
			},
			"public_dns_server": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "NS server used by the gateway.",
			},
			"cloud_instance_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Instance ID of the gateway.",
			},
			"private_ip": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Private IP address of the Gateway created.",
			},
			"peering_ha_cloud_instance_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Instance ID of the peering HA gateway.",
			},
			"peering_ha_gw_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Aviatrix gateway unique name of HA gateway.",
			},
			"peering_ha_private_ip": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Private IP address of HA gateway.",
			},
			"fqdn_lan_interface": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "FQDN gateway lan interface id.",
			},
		},
	}
}

func validateCloudType(i interface{}, k string) (warnings []string, errors []error) {
	return validation.IntInSlice([]int{1})(i, k)
}
