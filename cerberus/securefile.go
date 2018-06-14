/*
Copyright 2017 Nike Inc.

Licensed under the Apache License, Version 2.0 (the License);
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an AS IS BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cerberus

import (
	"fmt"
	"net/http"

	"github.com/Nike-Inc/cerberus-go-client/api"
)

// SecureFile is a subclient for secure files
type SecureFile struct {
	c *Client
}

var secureFileBasePath = "/v1/secure-file"
var secureFileListBasePath = "/v1/secure-files"

// List returns a list of secure files
func (r *SecureFile) List() (*api.SecureFilesResponse, error) {
	resp, err := r.c.DoRequest(http.MethodGet, secureFileListBasePath, map[string]string{}, nil)
	if err != nil {
		return nil, fmt.Errorf("error while trying to get secure files: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error while trying to list secure f. Got HTTP status code %d", resp.StatusCode)
	}
	sfr := &api.SecureFilesResponse{}
	err = parseResponse(resp.Body, sfr)
	if err != nil {
		return nil, err
	}
	return sfr, nil
}
