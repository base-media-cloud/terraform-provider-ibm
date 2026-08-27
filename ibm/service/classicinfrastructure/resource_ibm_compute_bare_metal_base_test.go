// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package classicinfrastructure_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccIBMComputeBareMetalDataSource_base(t *testing.T) {
	//configName := "data.ibm_compute_bare_metal.tf-bm-ds-acc-test"
	//hostname := acctest.RandString(16)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				ResourceName:  "ibm_compute_bare_metal.edge_transcoder",
				ImportState:   true,
				ImportStateId: "gid:93e5bfe6-5cc0-4ea2-a5f7-49237d28c13d",
				Config:        testAccCheckIBMComputeBareMetalResourceConfigBase(),
				ImportStateCheck: func(s []*terraform.InstanceState) error {
					if len(s) != 1 {
						return fmt.Errorf("expected 1 state, got %d", len(s))
					}
					state := s[0]
					if state.Attributes["global_identifier"] != "93e5bfe6-5cc0-4ea2-a5f7-49237d28c13d" {
						return fmt.Errorf("expected global_identifier '93e5bfe6-5cc0-4ea2-a5f7-49237d28c13d', got %q", state.Attributes["global_identifier"])
					}
					if state.Attributes["id"] != "1900048" {
						return fmt.Errorf("expected id '1900048', got %q", state.Attributes["id"])
					}
					if state.Attributes["hostname"] != "base-poc-edge-001" {
						return fmt.Errorf("expected hostname 'base-poc-edge-001', got %q", state.Attributes["hostname"])
					}
					if state.Attributes["domain"] != "basemediacloud.com" {
						return fmt.Errorf("expected domain 'basemediacloud.com', got %q", state.Attributes["domain"])
					}
					return nil
				},
			},
		},
	})
}

func testAccCheckIBMComputeBareMetalResourceConfigBase() string {
	return fmt.Sprintf(`
		resource "ibm_compute_bare_metal" "edge_transcoder" {
		  package_key_name = "SINGLE_E3_1270_V6_CL"
		  process_key_name = "INTEL_INTEL_XEON_3_80"
		  memory           = 16
		  os_key_name      = "OS_UBUNTU_24_04_LTS_NOBLE_NUMBAT_64_BIT"
		  hostname         = "base-poc-edge-001"
		  domain           = "basemediacloud.com"
		  datacenter       = "lon04"
		  network_speed    = 1000  # 1 Gbps
		  public_bandwidth = 500
		  disk_key_names   = ["HARD_DRIVE_960GB_SSD"]
		  hourly_billing   = false
		
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
