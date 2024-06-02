package decider

import (
	"testing"
)

func TestCreateParty(t *testing.T) {
	partyUid, err := CreateParty("test", "test")
	if err != nil {
		t.Error(err)
	}
	t.Log(partyUid)
}

func TestJoinParty(t *testing.T) {
	partyUid, err := CreateParty("deciderUid", "userUid")
	if err != nil {
		t.Error(err)
	}
	err = JoinParty("deciderUid", partyUid, "userUid", "user")
	if err != nil {
		t.Error(err)
	}
}

func TestLeaveParty(t *testing.T) {
	partyUid, err := CreateParty("deciderUid", "userUid")
	if err != nil {
		t.Error(err)
	}
	err = JoinParty("deciderUid", partyUid, "userUid", "user")
	if err != nil {
		t.Error(err)
	}
	party, err := GetParty("deciderUid", partyUid)
	if err != nil {
		t.Error(err)
	}
	if party.Users[0].Uid != "userUid" {
		t.Error("Expected userUid, got", party.Users[0].Uid)
	}
	err = LeaveParty("deciderUid", partyUid, "userUid")
	if err != nil {
		t.Error(err)
	}
}

func TestGetParty(t *testing.T) {
	partyUid, err := CreateParty("deciderUid", "userUid")
	if err != nil {
		t.Error(err)
	}
	JoinParty("deciderUid", partyUid, "userUid2", "user")
	party, err := GetParty("deciderUid", partyUid)
	if err != nil {
		t.Error(err)
	}
	if party == nil {
		t.Error("Expected party, got nil")
	}
	t.Log(party)
}
