// Copyright 2018 Elasticsearch Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package types

import "time"

type Host interface {
	Info() HostInfo
}

type HostInfo struct {
	Architecture      string    `json:"architecture"`            // Hardware architecture (e.g. x86_64, arm, ppc, mips).
	BootTime          time.Time `json:"boot_time"`               // Host boot time.
	Containerized     *bool     `json:"containerized,omitempty"` // Is the process containerized.
	Hostname          string    `json:"hostname"`                // Hostname
	IPs               []string  `json:"ips,omitempty"`           // List of all IPs.
	KernelVersion     string    `json:"kernel_version"`          // Kernel version.
	MACs              []string  `json:"mac_addresses"`           // List of MAC addresses.
	OS                *OSInfo   `json:"os"`                      // OS information.
	Timezone          string    `json:"timezone"`                // System timezone.
	TimezoneOffsetSec int       `json:"timezone_offset_sec"`     // Timezone offset (seconds from UTC).
	UniqueID          string    `json:"id,omitempty"`            // Unique ID of the host (optional).
}

func (host HostInfo) Uptime() time.Duration {
	return time.Since(host.BootTime)
}

type OSInfo struct {
	Family   string `json:"family"`             // OS Family (e.g. redhat, debian, freebsd, windows).
	Platform string `json:"platform"`           // OS platform (e.g. centos, ubuntu, windows).
	Name     string `json:"name"`               // OS Name (e.g. Mac OS X, CentOS).
	Version  string `json:"version"`            // OS version (e.g. 10.12.6).
	Major    int    `json:"major"`              // Major release version.
	Minor    int    `json:"minor"`              // Minor release version.
	Patch    int    `json:"patch"`              // Patch release version.
	Build    string `json:"build,omitempty"`    // Build (e.g. 16G1114).
	Codename string `json:"codename,omitempty"` // OS codename (e.g. jessie).
}

type LoadAverager interface {
	LoadAverage() LoadAverage
}

type LoadAverage struct {
	One     float64 `json:"one_min"`
	Five    float64 `json:"five_min"`
	Fifteen float64 `json:"fifteen_min"`
}
