// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package classicinfrastructure_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMComputeBareMetalDataSource_basic(t *testing.T) {
	configName := "data.ibm_compute_bare_metal.tf-bm-ds-acc-test"
	hostname := acctest.RandString(16)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMComputeBareMetalDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						configName, "hostname", hostname),
					resource.TestCheckResourceAttr(
						configName, "domain", "terraformuat.ibm.com"),
					resource.TestCheckResourceAttr(
						configName, "os_reference_code", "UBUNTU_16_64"),
					resource.TestCheckResourceAttr(
						configName, "datacenter", "dal01"),
					resource.TestCheckResourceAttr(
						configName, "network_speed", "100"),
					resource.TestCheckResourceAttr(
						configName, "hourly_billing", "true"),
					resource.TestCheckResourceAttr(
						configName, "private_network_only", "false"),
					resource.TestCheckResourceAttr(
						configName, "ipv6_enabled", "true"),
					resource.TestCheckResourceAttr(
						configName, "secondary_ip_count", "4"),
					resource.TestCheckResourceAttrSet(
						configName, "secondary_ip_addresses.0"),
					resource.TestCheckResourceAttrSet(
						configName, "secondary_ip_addresses.1"),
					resource.TestCheckResourceAttrSet(
						configName, "secondary_ip_addresses.2"),
					resource.TestCheckResourceAttrSet(
						configName, "secondary_ip_addresses.3"),
					resource.TestCheckResourceAttr(
						configName, "user_metadata", "{\"value\":\"newvalue\"}"),
					resource.TestCheckResourceAttr(
						configName, "notes", "baremetal notes"),
					CheckStringSet(
						configName,
						"tags", []string{"collectd"},
					),
				),
			},
		},
	})
}

func testAccCheckIBMComputeBareMetalDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		resource "ibm_compute_bare_metal" "edge_transcoder" {
		  hostname       = "base-poc-edge-001"
		  domain         = "basemediacloud.com"
		  os_reference_code = "OS_UBUNTU_24_04_LTS_64_BIT"
		  datacenter     = "lon04"
		
		  hourly_billing       = true
		  fixed_config_preset  = "1U_2174S_64GB_2X2TB_RAID_1"
		
		  network_speed = 1000
		
		  tcp_monitoring = false
		}`)

}
