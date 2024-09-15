package curp

import (
	"fmt"
	"log"
	"os"
	"time"

	"go.etcd.io/etcd/raft/v3"
	"go.etcd.io/etcd/raft/v3/raftpb"
)

func getSnapshotFn2() (func() ([]byte, error), <-chan struct{}) {
	snapshotTriggeredC := make(chan struct{})
	return func() ([]byte, error) {
		snapshotTriggeredC <- struct{}{}
		return nil, nil
	}, snapshotTriggeredC
}

type peerInfo struct {
	peers              []string
	grpcservers        string
	commitC            <-chan *commit
	errorC             <-chan error
	proposeC           chan string
	stateC             chan raft.StateType
	confChangeC        chan raftpb.ConfChange
	snapshotTriggeredC <-chan struct{}
	state              *StateMachine
	id                 int
}

func (p *peerInfo) close() (err error) {

	time.Sleep(time.Second)
	go func() {
		for range p.commitC {
			// drain pending commits
		}
	}()
	close(p.proposeC)
	// wait for channel to close
	if erri := <-p.errorC; erri != nil {
		err = erri
		log.Println(err)
	}

	err = os.RemoveAll(fmt.Sprintf("./raftexample-%d", p.id))
	if err != nil {
		log.Println(err)
	}
	err = os.RemoveAll(fmt.Sprintf("./raftexample-%d-snap", p.id))
	if err != nil {
		log.Println(err)
	}
	return err
}

func RunStateMachine(peers []string, grpcserver string, id int) (*StateMachine, func() error) {
	peer := &peerInfo{
		peers:       peers,
		grpcservers: grpcserver,
		proposeC:    make(chan string, 1),
		stateC:      make(chan raft.StateType, 1),
		confChangeC: make(chan raftpb.ConfChange, 1),
		id:          id,
	}
	fn, snapshotTriggeredC := getSnapshotFn2()
	peer.snapshotTriggeredC = snapshotTriggeredC
	commitC, errorC, snapReady := newRaftNode(id, peer.peers, false, fn, peer.proposeC, peer.stateC, peer.confChangeC)
	peer.commitC = commitC
	peer.errorC = errorC
	peer.state = newGrpcState(peer.grpcservers, id, <-snapReady, peer.proposeC, peer.stateC, peer.commitC, peer.errorC)
	return peer.state, peer.close
}
