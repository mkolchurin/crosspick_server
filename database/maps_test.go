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
	err = InsertMap("test", 2, "test", "test", "test", user.ID)
	if err != nil {
		t.Error(err)
	}
	DeleteMap("test")
}

func TestRemoveMap(t *testing.T) {
	err := DeleteMap("test")
	if err != nil {
		t.Error(err)
	}
}

func TestCreateListOfMaps(t *testing.T) {
	user, err := GetUser("admin")
	if err != nil {
		t.Error(err)
	}
	InsertMap("Vile Reef (2)", 2, "2p_vile_reef.jpg", "Описание/Description", "dowss", user.ID)
	InsertMap("Eden (2)", 2, "2p_eden.jpg", "Описание/Description", "dowss", user.ID)
	InsertMap("Doom Spiral (4)", 4, "4p_doom_spiral.jpg", "Описание/Description", "dowss", user.ID)
	InsertMap("GorHael Crater (4)", 4, "4p_gorhael_crater.jpg", "Описание/Description", "dowss", user.ID)
	InsertMap("Ariel Highlands (4)", 4, "4p_ariel_highlands.jpg", "Описание/Description", "dowss", user.ID)
	InsertMap("Forgotten Isles (4)", 4, "4p_forgotten_isles.jpg", "Описание/Description", "dowss", user.ID)
	InsertMap("Gurmuns Pass (4)", 4, "4p_gurmuns_pass.jpg", "Описание/Description", "dowss", user.ID)
	InsertMap("Temple of Cyrene (6)", 6, "4p_temple_of_arcanum.jpg", "Описание/Description", "dowss", user.ID)
	InsertMap("Murad Swamplands (4)", 4, "4p_murad_swamplands.jpg", "Описание/Description", "dowss", user.ID)
	InsertMap("Dread Peak (4)", 4, "4p_dread_peak.jpg", "Описание/Description", "dowss", user.ID)
	InsertMap("Eres Badlands (4)", 4, "4p_eres_badlands.jpg", "Описание/Description", "dowss", user.ID)
	InsertMap("Janus Savannah (4)", 4, "4p_janus_savannah.jpg", "Описание/Description", "dowss", user.ID)
	InsertMap("Janus Savannah Pro (4)", 4, "4p_janus_savannah_pro.jpg", "Описание/Description", "dowss", user.ID)
	InsertMap("Panrea Lowlands (4)", 4, "4p_panrea_lowlands.jpg", "Описание/Description", "dowss", user.ID)
	InsertMap("The Skerries (4)", 4, "4p_skerries.jpg", "Описание/Description", "dowss", user.ID)
	InsertMap("The Skerries Pro (4)", 4, "4p_skerries_pro.jpg", "Описание/Description", "dowss", user.ID)
	InsertMap("Torrents (4)", 4, "4p_torrents.jpg", "Описание/Description", "dowss", user.ID)
	InsertMap("Van de Mar Mountains (4)", 4, "4p_van_de_mar_mountains.jpg", "Описание/Description", "dowss", user.ID)
	InsertMap("Sands of Victory (4)", 4, "4p_sands_of_victory.jpg", "Описание/Description", "dowss", user.ID)
	InsertMap("Antiga Bay (4)", 4, "4p_antiga_bay.jpg", "Описание/Description", "dowss", user.ID)
	InsertMap("Testcake [Rem] (4)", 4, "4p_testcake.jpg", "Описание/Description", "dowss", user.ID)
	InsertMap("Imperial Area [Ed] (4)", 4, "4p_imperial_area.jpg", "Описание/Description", "dowss", user.ID)
	InsertMap("Volcanic Reaction (4)", 4, "4p_volcanic_reaction.jpg", "Описание/Description", "dowss", user.ID)
	InsertMap("Broken Lands (4)", 4, "4p_broken_lands.jpg", "Описание/Description", "dowss", user.ID)
	InsertMap("Vyasastan (5)", 5, "5p_vyasastan.jpg", "Описание/Description", "dowss", user.ID)
	InsertMap("Totmachers Prison (5)", 5, "5p_totmachers_prison.jpg", "Описание/Description", "dowss", user.ID)
	InsertMap("Red Jungle (5)", 5, "5p_red_jungle.jpg", "Описание/Description", "dowss", user.ID)
	InsertMap("Aceria Forests (5)", 5, "5p_aceria_forests.jpg", "Описание/Description", "dowss", user.ID)
	InsertMap("Eye of gorgon (5)", 5, "5p_eye_of_gorgon.jpg", "Описание/Description", "dowss", user.ID)
	InsertMap("Parmenie (6)", 6, "6p_parmenie.jpg", "Описание/Description", "dowss", user.ID)
	InsertMap("Temple of Cyrene (6)", 6, "6p_temple_cyrene.jpg", "Описание/Description", "dowss", user.ID)
	InsertMap("Nirraein (6)", 6, "6p_nirraein.jpg", "Описание/Description", "dowss", user.ID)
	InsertMap("Shakun Coast (6)", 6, "6p_shakun_coast.jpg", "Описание/Description", "dowss", user.ID)
	InsertMap("Streets of Vogen (6)", 6, "6pteam_streets_of_vogen.jpg", "Описание/Description", "dowss", user.ID)

}
