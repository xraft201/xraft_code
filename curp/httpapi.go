// // Copyright 2015 The etcd Authors
// //
// // Licensed under the Apache License, Version 2.0 (the "License");
// // you may not use this file except in compliance with the License.
// // You may obtain a copy of the License at
// //
// //     http://www.apache.org/licenses/LICENSE-2.0
// //
// // Unless required by applicable law or agreed to in writing, software
// // distributed under the License is distributed on an "AS IS" BASIS,
// // WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// // See the License for the specific language governing permissions and
// // limitations under the License.
package curp

// package curp

// import (
// 	"io/ioutil"
// 	"log"
// 	"net/http"
// 	"strconv"

// 	"github/Fischer0522/xraft/curp/command"

// 	"go.etcd.io/etcd/raft/v3/raftpb"
// )

// // Handler for a http based key-value store backed by raft
// type httpKVAPI struct {
// 	store       *StateMachine
// 	confChangeC chan<- raftpb.ConfChange
// }

// func (h *httpKVAPI) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	key := r.RequestURI
// 	defer r.Body.Close()
// 	switch {
// 	case r.Method == "PUT":
// 		v, err := ioutil.ReadAll(r.Body)
// 		if err != nil {
// 			log.Printf("Failed to read on PUT (%v)\n", err)
// 			http.Error(w, "Failed on PUT", http.StatusBadRequest)
// 			return
// 		}
// 		proposeId := command.ProposeId{
// 			ClientId: 1,
// 			SeqId:    1,
// 		}
// 		cmd := command.ClientCommand{
// 			ProposeId: proposeId,
// 			Op:        command.PUT,
// 			Key:       key,
// 			Value:     string(v),
// 		}
// 		h.store.Propose(cmd)

// 		// Optimistic-- no waiting for ack from raft. Value is not yet
// 		// committed so a subsequent GET on the key may return old value
// 		w.WriteHeader(http.StatusNoContent)
// 	case r.Method == "GET":
// 		proposeId := command.ProposeId{
// 			ClientId: 1,
// 			SeqId:    1,
// 		}
// 		cmd := command.ClientCommand{
// 			ProposeId: proposeId,
// 			Op:        command.GET,
// 			Key:       key,
// 		}
// 		if reply := h.store.Lookup(cmd); reply.Status == command.ACCEPTED {
// 			w.Write([]byte(reply.Result))
// 		} else {
// 			http.Error(w, "Failed to GET", http.StatusNotFound)
// 		}
// 	case r.Method == "POST":
// 		url, err := ioutil.ReadAll(r.Body)
// 		if err != nil {
// 			log.Printf("Failed to read on POST (%v)\n", err)
// 			http.Error(w, "Failed on POST", http.StatusBadRequest)
// 			return
// 		}

// 		nodeId, err := strconv.ParseUint(key[1:], 0, 64)
// 		if err != nil {
// 			log.Printf("Failed to convert ID for conf change (%v)\n", err)
// 			http.Error(w, "Failed on POST", http.StatusBadRequest)
// 			return
// 		}

// 		cc := raftpb.ConfChange{
// 			Type:    raftpb.ConfChangeAddNode,
// 			NodeID:  nodeId,
// 			Context: url,
// 		}
// 		h.confChangeC <- cc

// 		// As above, optimistic that raft will apply the conf change
// 		w.WriteHeader(http.StatusNoContent)
// 	case r.Method == "DELETE":
// 		nodeId, err := strconv.ParseUint(key[1:], 0, 64)
// 		if err != nil {
// 			log.Printf("Failed to convert ID for conf change (%v)\n", err)
// 			http.Error(w, "Failed on DELETE", http.StatusBadRequest)
// 			return
// 		}

// 		cc := raftpb.ConfChange{
// 			Type:   raftpb.ConfChangeRemoveNode,
// 			NodeID: nodeId,
// 		}
// 		h.confChangeC <- cc

// 		// As above, optimistic that raft will apply the conf change
// 		w.WriteHeader(http.StatusNoContent)
// 	default:
// 		w.Header().Set("Allow", "PUT")
// 		w.Header().Add("Allow", "GET")
// 		w.Header().Add("Allow", "POST")
// 		w.Header().Add("Allow", "DELETE")
// 		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
// 	}
// }

// // serveHttpKVAPI starts a key-value server with a GET/PUT API and listens.
// func serveHttpKVAPI(kv *StateMachine, port int, confChangeC chan<- raftpb.ConfChange, errorC <-chan error) {
// 	srv := http.Server{
// 		Addr: ":" + strconv.Itoa(port),
// 		Handler: &httpKVAPI{
// 			store:       kv,
// 			confChangeC: confChangeC,
// 		},
// 	}
// 	go func() {
// 		if err := srv.ListenAndServe(); err != nil {
// 			log.Fatal(err)
// 		}
// 	}()

// 	// exit when raft goes down
// 	if err, ok := <-errorC; ok {
// 		log.Fatal(err)
// 	}
// }
