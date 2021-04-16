package basic

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceAviatrixFireNet() *schema.Resource {
	return &schema.Resource{

		Schema: map[string]*schema.Schema{
			"vpc_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "VPC ID.",
			},
			"firewall_instance_association": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "List of firewall instances associated with fireNet.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"firenet_gw_name": {
							Type:        schema.TypeString,
							Description: "Name of the gateway to launch the firewall instance.",
							Computed:    true,
						},
						"instance_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "ID of Firewall instance, or FQDN Gateway's gw_name.",
						},
						"vendor_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Indication it is a firewall instance or FQDN gateway to be associated to fireNet. Valid values: 'Generic', 'fqdn_gateway'. Value 'fqdn_gateway' is required for FQDN gateway.",
						},
						"firewall_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Firewall instance name, or FQDN Gateway's gw_name, required if it is a firewall instance.",
						},
						"lan_interface": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Lan interface ID, required if it is a firewall instance.",
						},
						"management_interface": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Management interface ID, required if it is a firewall instance.",
						},
						"egress_interface": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Egress interface ID, required if it is a firewall instance.",
						},
						"attached": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Switch to attach/detach firewall instance to/from fireNet.",
						},
					},
				},
			},
			"inspection_enabled": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Enable/Disable traffic inspection.",
			},
			"egress_enabled": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Enable/Disable egress through firewall.",
			},
			"hashing_algorithm": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Hashing algorithm to load balance traffic across the firewall.",
			},
			"keep_alive_via_lan_interface_enabled": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Enable Keep Alive via Firewall LAN Interface.",
			},
			"tgw_segmentation_for_egress_enabled": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Enable TGW segmentation for egress.",
			},
			"egress_static_cidrs": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "List of egress static cidrs.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}