// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package classicinfrastructure_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMComputeBareMetalDataSource_socket(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMComputeBareMetalDestroy,
		Steps: []resource.TestStep{
			{
				Config:  testAccCheckIBMComputeBareMetalResourceConfigSocket(),
				Destroy: false,
				Check:   resource.ComposeTestCheckFunc(),
			},
		},
	})
}

func testAccCheckIBMComputeBareMetalResourceConfigSocket() string {
	return fmt.Sprintf(`
		resource "ibm_compute_bare_metal" "edge_transcoder" {
		  fixed_config_preset    = "1U_2174S_64GB_2X4TB_RAID_1"
		  os_key_name            = "OS_UBUNTU_24_04_LTS_NOBLE_NUMBAT_64_BIT"
		  hostname               = "base-socket-test-001"
		  domain                 = "basemediacloud.com"
		  datacenter             = "lon04"
		  network_speed          = 1000  # 1 Gbps
		  public_bandwidth       = 500
		  post_install_script_uri = "https://raw.githubusercontent.com/base-media-cloud/base-ibm-scripts/main/bare-metal-init.sh"
		  hourly_billing         = false
		
		  private_network_only = false
		  redundant_power_supply = false
		  tags = [
			"edge-transcoder",
			"mediahub",
		  ]
		
		  # Monitoring
		  tcp_monitoring = false
		}`)

}
