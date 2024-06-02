package decider

import (
	"sync"
	"time"

	"github.com/pkg/errors"
)

type Map struct {
	Name    string `json:"name"`
	Img     string `json:"img"`
	Checked bool   `json:"checked"`
}

type MapPool struct {
	Maps []Map `json:"maps"`
}

type DeciderUser struct {
	Uid      string `json:"uid"`
	LastPing int64  `json:"lastPing"`
}

type Decider struct {
	Users     []DeciderUser `json:"users"`
	Observers []DeciderUser `json:"observers"`
	Pool      MapPool       `json:"pool"`
	mux       sync.Mutex
	/* rules TODO: */
}

var deciders = make(map[string][]*Decider)

const (
	roleUser     = "user"
	roleObserver = "observer"
	timeToLive   = 5
)

func init() {
	go func() {
		for decider := range deciders {
			for _, dec := range deciders[decider] {
				dec.mux.Lock()
				for index, user := range dec.Users {
					if time.Now().Unix()-user.LastPing > timeToLive {
						// remove user
						dec.Users = append(dec.Users[:index], dec.Users[index+1:]...)
					}
				}
				for index, observer := range dec.Observers {
					if time.Now().Unix()-observer.LastPing > timeToLive {
						// remove observer
						dec.Observers = append(dec.Observers[:index], dec.Observers[index+1:]...)
					}
				}
				dec.mux.Unlock()
			}
		}
		time.Sleep(timeToLive * time.Second)
	}()
}

func CreateParty(deciderUid string, userUid string) (partyUid int, err error) {
	deciders[deciderUid] = append(deciders[deciderUid], &Decider{
		Users:     []DeciderUser{{Uid: userUid, LastPing: time.Now().Unix()}},
		Observers: []DeciderUser{},
		Pool:      MapPool{},
	})

	return
}

func JoinParty(deciderUid string, partyUid int, user_uid string, role string) error {
	if deciders[deciderUid] == nil {
		return errors.Errorf("Not Found")
	}
	deciders[deciderUid][partyUid].mux.Lock()
	defer deciders[deciderUid][partyUid].mux.Unlock()
	switch role {
	case roleUser:
		if len(deciders[deciderUid][partyUid].Users) == 2 {
			return errors.Errorf("Full")
		}
		deciders[deciderUid][partyUid].Users =
			append(deciders[deciderUid][partyUid].Users, DeciderUser{Uid: user_uid, LastPing: time.Now().Unix()})
	case roleObserver:
		deciders[deciderUid][partyUid].Observers =
			append(deciders[deciderUid][partyUid].Observers, DeciderUser{Uid: user_uid, LastPing: time.Now().Unix()})
	default:
		return errors.Errorf("Invalid Role")
	}
	return nil
}

func LeaveParty(deciderUid string, partyUid int, user_uid string) error {
	if deciders[deciderUid] == nil {
		return errors.Errorf("Not Found")
	}
	deciders[deciderUid][partyUid].mux.Lock()
	defer deciders[deciderUid][partyUid].mux.Unlock()
	for index, user := range deciders[deciderUid][partyUid].Users {
		if user.Uid == user_uid {
			deciders[deciderUid][partyUid].Users =
				append(deciders[deciderUid][partyUid].Users[:index], deciders[deciderUid][partyUid].Users[index+1:]...)
			return nil
		}
	}
	for index, observer := range deciders[deciderUid][partyUid].Observers {
		if observer.Uid == user_uid {
			deciders[deciderUid][partyUid].Observers =
				append(deciders[deciderUid][partyUid].Observers[:index], deciders[deciderUid][partyUid].Observers[index+1:]...)
			return nil
		}
	}
	return errors.Errorf("Not Found")
}

func GetParty(deciderUid string, partyUid int) (*Decider, error) {
	if deciders[deciderUid] == nil {
		return nil, errors.Errorf("Not Found")
	}
	return deciders[deciderUid][partyUid], nil
}

func GetParties(deciderUid string) ([]*Decider, error) {
	if deciders[deciderUid] == nil {
		return nil, errors.Errorf("Not Found")
	}
	return deciders[deciderUid], nil
}
