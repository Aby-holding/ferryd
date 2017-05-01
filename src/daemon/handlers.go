//
// Copyright © 2017 Ikey Doherty <ikey@solus-project.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package main

import (
	"bytes"
	"encoding/json"
	"ferry"
	"fmt"
	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
	"net/http"
)

// GetVersion will return the current version of the ferryd
func (s *Server) GetVersion(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// For now return nothing and default to 200 OK
	fmt.Printf("Got a version request: %v\n", r.URL.Path)

	vq := ferry.VersionRequest{Version: ferry.Version}
	buf := bytes.Buffer{}
	if err := json.NewEncoder(&buf).Encode(&vq); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(buf.Bytes())
}

// CreateRepo will handle remote requests for repository creation
func (s *Server) CreateRepo(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")
	log.WithFields(log.Fields{
		"id": id,
	}).Info("Repository creation requested")
	err := s.manager.CreateRepo(id)
	// TODO: Make this Moar Better..
	if err != nil {
		log.WithFields(log.Fields{
			"id":    id,
			"error": err,
		}).Error("Failed to create repository")
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
