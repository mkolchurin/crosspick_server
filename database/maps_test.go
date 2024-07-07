package database

import (
	"testing"
)

func TestGetMaps(t *testing.T) {
	maps, err := GetMaps(0, 1)
	if err != nil {
		t.Error(err)
	}
	if len(maps) != 1 {
		t.Error("Expected 1 map, got", len(maps))
	}
	t.Log(maps)
	maps, err = GetMaps(1, 3)
	if err != nil {
		t.Error(err)
	}
	if len(maps) != 3 {
		t.Error("Expected 3 maps, got", len(maps))
	}
	t.Log(maps)
}

func TestInsertMap(t *testing.T) {
	user, err := GetUser("admin")
	if err != nil {
		t.Error(err)
	}
	err = InsertMap("test", 2, "test", "test", "test", user.Id)
	if err != nil {
		t.Error(err)
	}
	RemoveMap("test")
}

func TestRemoveMap(t *testing.T) {
	err := RemoveMap("test")
	if err != nil {
		t.Error(err)
	}
}

func TestCreateListOfMaps(t *testing.T) {
	user, err := GetUser("admin")
	if err != nil {
		t.Error(err)
	}
	InsertMap("Vile Reef (2)", 2, "2p_vile_reef.jpg", "Описание/Description", "dowss", user.Id)
	InsertMap("Eden (2)", 2, "2p_eden.jpg", "Описание/Description", "dowss", user.Id)
	InsertMap("Doom Spiral (4)", 4, "4p_doom_spiral.jpg", "Описание/Description", "dowss", user.Id)
	InsertMap("GorHael Crater (4)", 4, "4p_gorhael_crater.jpg", "Описание/Description", "dowss", user.Id)
	InsertMap("Ariel Highlands (4)", 4, "4p_ariel_highlands.jpg", "Описание/Description", "dowss", user.Id)
	InsertMap("Forgotten Isles (4)", 4, "4p_forgotten_isles.jpg", "Описание/Description", "dowss", user.Id)
	InsertMap("Gurmuns Pass (4)", 4, "4p_gurmuns_pass.jpg", "Описание/Description", "dowss", user.Id)
	InsertMap("Temple of Cyrene (6)", 6, "4p_temple_of_arcanum.jpg", "Описание/Description", "dowss", user.Id)
	InsertMap("Murad Swamplands (4)", 4, "4p_murad_swamplands.jpg", "Описание/Description", "dowss", user.Id)
	InsertMap("Dread Peak (4)", 4, "4p_dread_peak.jpg", "Описание/Description", "dowss", user.Id)
	InsertMap("Eres Badlands (4)", 4, "4p_eres_badlands.jpg", "Описание/Description", "dowss", user.Id)
	InsertMap("Janus Savannah (4)", 4, "4p_janus_savannah.jpg", "Описание/Description", "dowss", user.Id)
	InsertMap("Janus Savannah Pro (4)", 4, "4p_janus_savannah_pro.jpg", "Описание/Description", "dowss", user.Id)
	InsertMap("Panrea Lowlands (4)", 4, "4p_panrea_lowlands.jpg", "Описание/Description", "dowss", user.Id)
	InsertMap("The Skerries (4)", 4, "4p_skerries.jpg", "Описание/Description", "dowss", user.Id)
	InsertMap("The Skerries Pro (4)", 4, "4p_skerries_pro.jpg", "Описание/Description", "dowss", user.Id)
	InsertMap("Torrents (4)", 4, "4p_torrents.jpg", "Описание/Description", "dowss", user.Id)
	InsertMap("Van de Mar Mountains (4)", 4, "4p_van_de_mar_mountains.jpg", "Описание/Description", "dowss", user.Id)
	InsertMap("Sands of Victory (4)", 4, "4p_sands_of_victory.jpg", "Описание/Description", "dowss", user.Id)
	InsertMap("Antiga Bay (4)", 4, "4p_antiga_bay.jpg", "Описание/Description", "dowss", user.Id)
	InsertMap("Testcake [Rem] (4)", 4, "4p_testcake.jpg", "Описание/Description", "dowss", user.Id)
	InsertMap("Imperial Area [Ed] (4)", 4, "4p_imperial_area.jpg", "Описание/Description", "dowss", user.Id)
	InsertMap("Volcanic Reaction (4)", 4, "4p_volcanic_reaction.jpg", "Описание/Description", "dowss", user.Id)
	InsertMap("Broken Lands (4)", 4, "4p_broken_lands.jpg", "Описание/Description", "dowss", user.Id)
	InsertMap("Vyasastan (5)", 5, "5p_vyasastan.jpg", "Описание/Description", "dowss", user.Id)
	InsertMap("Totmachers Prison (5)", 5, "5p_totmachers_prison.jpg", "Описание/Description", "dowss", user.Id)
	InsertMap("Red Jungle (5)", 5, "5p_red_jungle.jpg", "Описание/Description", "dowss", user.Id)
	InsertMap("Aceria Forests (5)", 5, "5p_aceria_forests.jpg", "Описание/Description", "dowss", user.Id)
	InsertMap("Eye of gorgon (5)", 5, "5p_eye_of_gorgon.jpg", "Описание/Description", "dowss", user.Id)
	InsertMap("Parmenie (6)", 6, "6p_parmenie.jpg", "Описание/Description", "dowss", user.Id)
	InsertMap("Temple of Cyrene (6)", 6, "6p_temple_cyrene.jpg", "Описание/Description", "dowss", user.Id)
	InsertMap("Nirraein (6)", 6, "6p_nirraein.jpg", "Описание/Description", "dowss", user.Id)
	InsertMap("Shakun Coast (6)", 6, "6p_shakun_coast.jpg", "Описание/Description", "dowss", user.Id)
	InsertMap("Streets of Vogen (6)", 6, "6pteam_streets_of_vogen.jpg", "Описание/Description", "dowss", user.Id)

}
