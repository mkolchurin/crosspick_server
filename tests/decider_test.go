package tests

import (
	"testing"

	"github.com/mkolchurin/crosspick_server/logic"
)

func TestCreateParty(t *testing.T) {
	partyUid, err := logic.CreateParty("test", "test")
	if err != nil {
		t.Error(err)
	}
	t.Log(partyUid)
}

func TestJoinParty(t *testing.T) {
	partyUid, err := logic.CreateParty("deciderUid", "userUid")
	if err != nil {
		t.Error(err)
	}
	err = logic.JoinParty("deciderUid", partyUid, "userUid", "user")
	if err != nil {
		t.Error(err)
	}
}

func TestLeaveParty(t *testing.T) {
	partyUid, err := logic.CreateParty("deciderUid", "userUid")
	if err != nil {
		t.Error(err)
	}
	err = logic.JoinParty("deciderUid", partyUid, "userUid", "user")
	if err != nil {
		t.Error(err)
	}
	party, err := logic.GetParty("deciderUid", partyUid)
	if err != nil {
		t.Error(err)
	}
	if party.Users[0].Uid != "userUid" {
		t.Error("Expected userUid, got", party.Users[0].Uid)
	}
	err = logic.LeaveParty("deciderUid", partyUid, "userUid")
	if err != nil {
		t.Error(err)
	}
}

func TestGetParty(t *testing.T) {
	partyUid, err := logic.CreateParty("deciderUid", "userUid")
	if err != nil {
		t.Error(err)
	}
	logic.JoinParty("deciderUid", partyUid, "userUid2", "user")
	party, err := logic.GetParty("deciderUid", partyUid)
	if err != nil {
		t.Error(err)
	}
	if party == nil {
		t.Error("Expected party, got nil")
	}
	t.Log(party)
}
