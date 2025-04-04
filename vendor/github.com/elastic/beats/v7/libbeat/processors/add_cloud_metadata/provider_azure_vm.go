// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package add_cloud_metadata

import (
	"github.com/elastic/beats/v7/libbeat/common"
	s "github.com/elastic/beats/v7/libbeat/common/schema"
	c "github.com/elastic/beats/v7/libbeat/common/schema/mapstriface"
)

// Azure VM Metadata Service
var azureVMMetadataFetcher = provider{
	Name: "azure-compute",

	Local: true,

	Create: func(_ string, config *common.Config) (metadataFetcher, error) {
		azMetadataURI := "/metadata/instance/compute?api-version=2021-02-01"
		azHeaders := map[string]string{"Metadata": "true"}
		azSchema := func(m map[string]interface{}) common.MapStr {
			m["serviceName"] = "Virtual Machines"
			out, _ := s.Schema{
				"account": s.Object{
					"id": c.Str("subscriptionId"),
				},
				"instance": s.Object{
					"id":   c.Str("vmId"),
					"name": c.Str("name"),
				},
				"machine": s.Object{
					"type": c.Str("vmSize"),
				},
				"service": s.Object{
					"name": c.Str("serviceName"),
				},
				"region": c.Str("location"),
			}.Apply(m)
			return common.MapStr{"cloud": out}
		}

		fetcher, err := newMetadataFetcher(config, "azure", azHeaders, metadataHost, azSchema, azMetadataURI)
		return fetcher, err
	},
}
