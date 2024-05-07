package tests

import (
	"testing"

	"github.com/mkolchurin/crosspick_server/db"
)

// {
// 	"maps": [
// 	  "2p_abandon_all_hope.jpg",
// 	  "2p_absolute_zero.jpg",
// 	  "2p_antiga_bay.jpg",
// 	  "2p_battle_marshes.jpg",
// 	  "2p_blood_river.jpg",
// 	  "2p_blood_river_[Rem].jpg",
// 	  "2p_blood_river_remix.jpg",
// 	  "2p_bloody hell.jpg",
// 	  "2p_bloody_hell.jpg",
// 	  "2p_bloody_hell_[Ed].jpg",
// 	  "2p_bridge_too_far.jpg",
// 	  "2p_chaos_gate.jpg",
// 	  "2p_colosseum.jpg",
// 	  "2p_deadly_fun_archeology.jpg",
// 	  "2p_deadmans_crossing.jpg",
// 	  "2p_edemus_gamble.jpg",
// 	  "2p_eden.jpg",
// 	  "2p_emerald_river.jpg",
// 	  "2p_emperors_valley.jpg",
// 	  "2p_faceoff.jpg",
// 	  "2p_fallen_city.jpg",
// 	  "2p_fallen_city_[Rem].jpg",
// 	  "2p_fata_morgana.jpg",
// 	  "2p_fata_morgana_[Rem].jpg",
// 	  "2p_fear.jpg",
// 	  "2p_fraziersdemise.jpg",
// 	  "2p_frostbite_river.jpg",
// 	  "2p_galenas_crusade.jpg",
// 	  "2p_haines_demise.jpg",
// 	  "2p_hellfire_canyon.jpg",
// 	  "2p_jungle_morning.jpg",
// 	  "2p_light_brigade.jpg",
// 	  "2p_light_brigade_pro.jpg",
// 	  "2p_marsh_of_rakat.jpg",
// 	  "2p_meeting_of_minds.jpg",
// 	  "2p_meeting_of_minds_pro.jpg",
// 	  "2p_meeting_of_minds_pro_5p.jpg",
// 	  "2p_meeting_of_minds_pro_lis.jpg",
// 	  "2p_meeting_of_minds_pro_lis_without_1p.jpg",
// 	  "2p_moonbase.jpg",
// 	  "2p_moonbase_[Ed].jpg",
// 	  "2p_outer_reaches.jpg",
// 	  "2p_quests_triumph.jpg",
// 	  "2p_quests_triumph_pro.jpg",
// 	  "2p_railway.jpg",
// 	  "2p_river_bed.jpg",
// 	  "2p_short_below_zero.jpg",
// 	  "2p_shrine_of_excellion.jpg",
// 	  "2p_shrine_of_excellion_[Rem].jpg",
// 	  "2p_sugaroasis.jpg",
// 	  "2p_tainted_pair.jpg",
// 	  "2p_tazins_folly.jpg",
// 	  "2p_terror_psyclaw.jpg",
// 	  "2p_tiboraxx.jpg",
// 	  "2p_titan_fall.jpg",
// 	  "2p_titan_fall_[Rem].jpg",
// 	  "2p_tower_ruins.jpg",
// 	  "2p_tranquilitys_end.jpg",
// 	  "2p_tranquilitys_end_[Rem].jpg",
// 	  "2p_tranquilitys_end_pro.jpg",
// 	  "2p_valley_of_khorne.jpg",
// 	  "2p_velvet_duress.jpg",
// 	  "2p_vile_reef.jpg",
// 	  "2p_volatile_ground.jpg",
// 	  "2p_vortex_plateau.jpg",
// 	  "2p_winter_confortation.jpg",
// 	  "3p_coral_reef.jpg",
// 	  "3p_faded_dreams.jpg",
// 	  "3p_fortress.jpg",
// 	  "3p_heat_wave.jpg",
// 	  "4p_antiga_bay.jpg",
// 	  "4p_ariel_highlands.jpg",
// 	  "4p_biffys_peril.jpg",
// 	  "4p_boramus.jpg",
// 	  "4p_broken_lands.jpg",
// 	  "4p_cape_of_despair.jpg",
// 	  "4p_chaos_platenau.jpg",
// 	  "4p_cold_war.jpg",
// 	  "4p_colosseum_of_deadman.jpg",
// 	  "4p_doom_spiral.jpg",
// 	  "4p_dread_peak.jpg",
// 	  "4p_eres_badlands.jpg",
// 	  "4p_forgotten_isles.jpg",
// 	  "4p_gorhael_crater.jpg",
// 	  "4p_gurmuns_pass.jpg",
// 	  "4p_ice_flow.jpg",
// 	  "4p_imperial_area.jpg",
// 	  "4p_into_the_breach.jpg",
// 	  "4p_janus_savannah.jpg",
// 	  "4p_janus_savannah_pro.jpg",
// 	  "4p_lost_relic.jpg",
// 	  "4p_marconia.jpg",
// 	  "4p_mariza.jpg",
// 	  "4p_mountain_trail-1.jpg",
// 	  "4p_mountain_trail.jpg",
// 	  "4p_murad_swamplands.jpg",
// 	  "4p_oderas valley.jpg",
// 	  "4p_oderas_valley.jpg",
// 	  "4p_panrea_lowlands.jpg",
// 	  "4p_quatra.jpg",
// 	  "4p_refinery.jpg",
// 	  "4p_rokclaw_foothills.jpg",
// 	  "4p_sad_place.jpg",
// 	  "4p_saints_square.jpg",
// 	  "4p_sands_of_victory.jpg",
// 	  "4p_skerries.jpg",
// 	  "4p_skerries_pro.jpg",
// 	  "4p_snowblind.jpg",
// 	  "4p_st_mathias_bridge_se.jpg",
// 	  "4p_starving_equator.jpg",
// 	  "4p_tainted_soul.jpg",
// 	  "4p_tartarus_center.jpg",
// 	  "4p_temple_of_arcanum.jpg",
// 	  "4p_testcake.jpg",
// 	  "4p_tiboraxx.jpg",
// 	  "4p_torrents.jpg",
// 	  "4p_van_de_mar_mountains.jpg",
// 	  "4p_volcanic reaction.jpg",
// 	  "4p_volcanic_reaction.jpg",
// 	  "4p_white_silence.jpg",
// 	  "5p_aceria_forests.jpg",
// 	  "5p_eye_of_gorgon.jpg",
// 	  "5p_gorhjan_jungle.jpg",
// 	  "5p_red_jungle.jpg",
// 	  "5p_totmachers_prison.jpg",
// 	  "5p_vyasastan.jpg",
// 	  "6p_agamar_desert.jpg",
// 	  "6p_alvarus.jpg",
// 	  "6p_aqueduct.jpg",
// 	  "6p_bloodshed_alley.jpg",
// 	  "6p_confrontation.jpg",
// 	  "6p_crossroads.jpg",
// 	  "6p_crozius_arcanum.jpg",
// 	  "6p_desiderata_highlands.jpg",
// 	  "6p_dread_alley.jpg",
// 	  "6p_fury_island.jpg",
// 	  "6p_gear.jpg",
// 	  "6p_hyperion_peaks.jpg",
// 	  "6p_irridene.jpg",
// 	  "6p_irridene_v2.jpg",
// 	  "6p_jungle_walls.jpg",
// 	  "6p_kasyr_lutien.jpg",
// 	  "6p_kaurav_city.jpg",
// 	  "6p_kocapb.jpg",
// 	  "6p_maiden_world.jpg",
// 	  "6p_mortalis.jpg",
// 	  "6p_nirraein.jpg",
// 	  "6p_orestan_plains.jpg",
// 	  "6p_parmenian_heath.jpg",
// 	  "6p_parmenie.jpg",
// 	  "6p_pavonian_heartland.jpg",
// 	  "6p_pavonis.jpg",
// 	  "6p_paynes_retribution.jpg",
// 	  "6p_platform.jpg",
// 	  "6p_principian_badlands.jpg",
// 	  "6p_principian_badlands_BM.jpg",
// 	  "6p_rhean_floodlands.jpg",
// 	  "6p_ruined_greatway.jpg",
// 	  "6p_shakun_coast.jpg",
// 	  "6p_snowblind.jpg",
// 	  "6p_tavesta_wastelands.jpg",
// 	  "6p_team_streets_of_vogen.jpg",
// 	  "6p_temple_cyrene.jpg",
// 	  "6p_temple_cyrene_pro.jpg",
// 	  "6p_testing_grounds.jpg",
// 	  "6p_thargorum.jpg",
// 	  "6p_tristram_plains.jpg",
// 	  "6p_trivian_groves.jpg",
// 	  "6p_trivian_groves_v2.jpg",
// 	  "6p_vandea_coast.jpg",
// 	  "6p_western_barrens.jpg",
// 	  "6pteam_streets_of_vogen.jpg",
// 	  "8p_burial_grounds.jpg",
// 	  "8p_cerulea.jpg",
// 	  "8p_dantoonvale.jpg",
// 	  "8p_daturias_pits.jpg",
// 	  "8p_demes_northlands.jpg",
// 	  "8p_doom_chamber.jpg",
// 	  "8p_forbidden_jungle.jpg",
// 	  "8p_fort atlantis.jpg",
// 	  "8p_glacier.jpg",
// 	  "8p_jalaganda_lowlands.jpg",
// 	  "8p_kier_harrad.jpg",
// 	  "8p_lost_hope.jpg",
// 	  "8p_marinus.jpg",
// 	  "8p_monse.jpg",
// 	  "8p_morholt_range.jpg",
// 	  "8p_morriah_coast.jpg",
// 	  "8p_oasis_of_sharr.jpg",
// 	  "8p_penal_colony.jpg",
// 	  "8p_rhean_jungle.jpg",
// 	  "8p_thurabis_plateau.jpg",
// 	  "8p_verdant_isles.jpg",
// 	  "8p_verdant_isles_ed.jpg",
// 	  "IM_MAP_IM_MAP (1).docx",
// 	  "IM_MAP_IM_MAP (2).docx",
// 	  "IM_MAP_IM_MAP.docx",
// 	  "[tp_mod]4p_ice_flow.jpg",
// 	  "[tp_mod]edemus_gamble.jpg",
// 	  "[tp_mod]galenas_crusade.jpg",
// 	  "[tp_mod]jungle_morning.jpg",
// 	  "[tp_mod]light_brigade.jpg",
// 	  "[tp_mod]meeting_of_minds.jpg",
// 	  "[tp_mod]quests_triumph.jpg",
// 	  "[tp_mod]shrine_of_excellion.jpg",
// 	  "allMaps.json",
// 	  "default.jpg",
// 	  "img.png",
// 	  "img_1.png",
// 	  "img_2.png",
// 	  "img_3.png",
// 	  "img_4.png",
// 	  "img_5.png",
// 	  "img_6.png",
// 	  "винегрет.jpg",
// 	  "гнездо глухаря.jpg",
// 	  "грибная поляна.jpg",
// 	  "из пекинской капусты.jpg",
// 	  "крабовый.jpg",
// 	  "летний.jpg",
// 	  "мимоза.jpg",
// 	  "оливье.jpg",
// 	  "орлиное гнездо.jpg",
// 	  "селедка под шубой.jpg",
// 	  "цезарь.jpg"
// 	]
//   }

func TestInsertMaps(t *testing.T) {
	err := db.InsertMap("Abandon All Hope", 2, "2p_abandon_all_hope.jpg", "", "")
	if err != nil {
		t.Error(err)
	}
	err = db.InsertMap("Absolute Zero", 2, "2p_absolute_zero.jpg", "", "")
	if err != nil {
		t.Error(err)
	}
	err = db.InsertMap("Antiga Bay", 2, "2p_antiga_bay.jpg", "", "")
	if err != nil {
		t.Error(err)
	}
	err = db.InsertMap("Battle Marshes", 2, "2p_battle_marshes.jpg", "", "")
	if err != nil {
		t.Error(err)
	}
	err = db.InsertMap("Blood River", 2, "2p_blood_river.jpg", "", "")
	if err != nil {
		t.Error(err)
	}
	err = db.InsertMap("Blood River [Rem]", 2, "2p_blood_river_[Rem].jpg", "", "")
	if err != nil {
		t.Error(err)
	}
	err = db.InsertMap("Blood River Remix", 2, "2p_blood_river_remix.jpg", "", "")
	if err != nil {
		t.Error(err)
	}
	err = db.InsertMap("Bloody Hell", 2, "2p_bloody hell.jpg", "", "")
	if err != nil {
		t.Error(err)
	}
	err = db.InsertMap("Bloody Hell", 2, "2p_bloody_hell.jpg", "", "")
	if err != nil {
		t.Error(err)
	}
	err = db.InsertMap("Bloody Hell [Ed]", 2, "2p_bloody_hell_[Ed].jpg", "", "")
	if err != nil {
		t.Error(err)
	}
	err = db.InsertMap("Bridge Too Far", 2, "2p_bridge_too_far.jpg", "", "")
	if err != nil {
		t.Error(err)
	}
	err = db.InsertMap("Chaos Gate", 2, "2p_chaos_gate.jpg", "", "")
	if err != nil {
		t.Error(err)
	}
	err = db.InsertMap("Colosseum", 2, "2p_colosseum.jpg", "", "")
	if err != nil {
		t.Error(err)
	}
	err = db.InsertMap("Deadly Fun Archeology", 2, "2p_deadly_fun_archeology.jpg", "", "")
	if err != nil {
		t.Error(err)
	}
	err = db.InsertMap("Deadmans Crossing", 2, "2p_deadmans_crossing.jpg", "", "")
	if err != nil {
		t.Error(err)
	}
	err = db.InsertMap("Edemus Gamble", 2, "2p_edemus_gamble.jpg", "", "")
	if err != nil {
		t.Error(err)
	}
	err = db.InsertMap("Eden", 2, "2p_eden.jpg", "", "")
	if err != nil {
		t.Error(err)
	}
	err = db.InsertMap("Emerald River", 2, "2p_emerald_river.jpg", "", "")
	if err != nil {
		t.Error(err)
	}
	err = db.InsertMap("Emperors Valley", 2, "2p_emperors_valley.jpg", "", "")
	if err != nil {
		t.Error(err)
	}
	err = db.InsertMap("Faceoff", 2, "2p_faceoff.jpg", "", "")
	if err != nil {
		t.Error(err)
	}
	err = db.InsertMap("Fallen City", 2, "2p_fallen_city.jpg", "", "")
	if err != nil {
		t.Error(err)
	}
	err = db.InsertMap("Fallen City [Rem]", 2, "2p_fallen_city_[Rem].jpg", "", "")
	if err != nil {
		t.Error(err)
	}
	err = db.InsertMap("Fata Morgana", 2, "2p_fata_morgana.jpg", "", "")
	if err != nil {
		t.Error(err)
	}
	err = db.InsertMap("Fata Morgana [Rem]", 2, "2p_fata_morgana_[Rem].jpg", "", "")
	if err != nil {
		t.Error(err)
	}
	err = db.InsertMap("Fear", 2, "2p_fear.jpg", "", "")
	if err != nil {
		t.Error(err)
	}
	err = db.InsertMap("Fraziersdemise", 2, "2p_fraziersdemise.jpg", "", "")
	if err != nil {
		t.Error(err)
	}
	err = db.InsertMap("Frostbite River", 2, "2p_frostbite_river.jpg", "", "")
	if err != nil {
		t.Error(err)
	}
	err = db.InsertMap("Galenas Crusade", 2, "2p_galenas_crusade.jpg", "", "")
	if err != nil {
		t.Error(err)
	}
	err = db.InsertMap("Haines Demise", 2, "2p_haines_demise.jpg", "", "")
	if err != nil {
		t.Error(err)
	}
	err = db.InsertMap("Hellfire Canyon", 2, "2p_hellfire_canyon.jpg", "", "")
	if err != nil {
		t.Error(err)
	}
	err = db.InsertMap("Jungle Morning", 2, "2p_jungle_morning.jpg", "", "")
	if err != nil {
		t.Error(err)
	}
	err = db.InsertMap("Light Brigade", 2, "2p_light_brigade.jpg", "", "")
	if err != nil {
		t.Error(err)
	}
	err = db.InsertMap("Light Brigade Pro", 2, "2p_light_brigade_pro.jpg", "", "")
	if err != nil {
		t.Error(err)
	}
	err = db.InsertMap("Marsh Of Rakat", 2, "2p_marsh_of_rakat.jpg", "", "")
	if err != nil {
		t.Error(err)
	}
	err = db.InsertMap("Meeting Of Minds", 2, "2p_meeting_of_minds.jpg", "", "")
	if err != nil {
		t.Error(err)
	}
	err = db.InsertMap("Meeting Of Minds Pro", 2, "2p_meeting_of_minds_pro.jpg", "", "")
	if err != nil {
		t.Error(err)
	}
	err = db.InsertMap("Meeting Of Minds Pro 5p", 2, "2p_meeting_of_minds_pro_5p.jpg", "", "")
	if err != nil {
		t.Error(err)
	}
	err = db.InsertMap("Meeting Of Minds Pro Lis", 2, "2p_meeting_of_minds_pro_lis.jpg", "", "")
	if err != nil {
		t.Error(err)
	}
	err = db.InsertMap("Meeting Of Minds Pro Lis Without 1p", 2, "2p_meeting_of_minds_pro_lis_without_1p.jpg", "", "")
	if err != nil {
		t.Error(err)
	}

}
