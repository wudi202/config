//
//Copyright [2016] [SnapRoute Inc]
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
//	 Unless required by applicable law or agreed to in writing, software
//	 distributed under the License is distributed on an "AS IS" BASIS,
//	 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//	 See the License for the specific language governing permissions and
//	 limitations under the License.
//
// _______  __       __________   ___      _______.____    __    ____  __  .___________.  ______  __    __
// |   ____||  |     |   ____\  \ /  /     /       |\   \  /  \  /   / |  | |           | /      ||  |  |  |
// |  |__   |  |     |  |__   \  V  /     |   (----` \   \/    \/   /  |  | `---|  |----`|  ,----'|  |__|  |
// |   __|  |  |     |   __|   >   <       \   \      \            /   |  |     |  |     |  |     |   __   |
// |  |     |  `----.|  |____ /  .  \  .----)   |      \    /\    /    |  |     |  |     |  `----.|  |  |  |
// |__|     |_______||_______/__/ \__\ |_______/        \__/  \__/     |__|     |__|      \______||__|  |__|
//

package clients

var ClientInterfaces = map[string]ClientIf{"ribd": &RIBDClient{},
	"asicd":      &ASICDClient{},
	"arpd":       &ARPDClient{},
	"bgpd":       &BGPDClient{},
	"lacpd":      &LACPDClient{},
	"dhcprelayd": &DHCPRELAYDClient{},
	"local":      &LocalClient{},
	"ospfd":      &OSPFDClient{},
	"stpd":       &STPDClient{},
	"bfdd":       &BFDDClient{},
	"vrrpd":      &VRRPDClient{},
	"sysd":       &SYSDClient{},
	"vxland":     &VXLANDClient{},
	"lldpd":      &LLDPDClient{},
	"dhcpd":      &DHCPDClient{},
	"fMgrd":      &FMGRDClient{},
	"ndpd":       &NDPDClient{},
	"platformd":  &PLATFORMDClient{},
	"notifierd":  &NOTIFIERDClient{},
	"isisd":      &ISISDClient{},
}
