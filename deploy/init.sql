-- create database with name test and create table maps with columns id, name, description
		-- "ID": 1,
		-- "UpdatedAt": "2024-07-13T22:14:29.725288+03:00",
		-- "DeletedAt": null,
		-- "Name": "Vile Reef (2)",
		-- "Mode": 2,
		-- "Icon": "2p_vile_reef.jpg",
		-- "Description": "Описание/Description",
		-- "Author": "dowss",
		-- "Order": 0,
		-- "CreatedAt": 1720898069,
		-- "UploaderId": 0

CREATE TABLE maps (
	ID          INT PRIMARY KEY,
	Name        VARCHAR(255) NOT NULL,
	Description TEXT,
	Mode        INT,
	Icon        VARCHAR(255),
	Author      VARCHAR(255),
	"Order"     INT,
	CreatedAt   TIMESTAMP DEFAULT NULL,
	UpdatedAt   TIMESTAMP DEFAULT NULL,
	DeletedAt   TIMESTAMP DEFAULT NULL,
	UploaderId  INT
);

INSERT INTO public.maps (ID, Name, Mode, Icon, Description, Author, "Order") VALUES (1, 'Vile Reef (2)', 2, '2p_vile_reef', 'Описание', 'dowss', 1);
INSERT INTO public.maps (ID, Name, Mode, Icon, Description, Author, "Order") VALUES (2, 'Eden (2)', 2, '2p_eden', 'Описание', 'dowss', 1);
INSERT INTO public.maps (ID, Name, Mode, Icon, Description, Author, "Order") VALUES (3, 'Frostbite River (2)', 2, '2p_frostbite_river', 'Описание', 'dowss', 1);
INSERT INTO public.maps (ID, Name, Mode, Icon, Description, Author, "Order") VALUES (4, 'Quests Triumph Pro (2)', 2, '2p_quests_triumph_pro', 'Описание', 'dowss', 1);
INSERT INTO public.maps (ID, Name, Mode, Icon, Description, Author, "Order") VALUES (5, 'Light Brigade (2)', 2, '2p_light_brigade', 'Описание', 'dowss', 1);
INSERT INTO public.maps (ID, Name, Mode, Icon, Description, Author, "Order") VALUES (6, 'Quests Triumph (2)', 2, '2p_quests_triumph', 'Описание', 'dowss', 1);
INSERT INTO public.maps (ID, Name, Mode, Icon, Description, Author, "Order") VALUES (7, 'Light Brigade Pro (2)', 2, '2p_light_brigade_pro', 'Описание', 'dowss', 1);
INSERT INTO public.maps (ID, Name, Mode, Icon, Description, Author, "Order") VALUES (8, 'Bridge Too Far (2)', 2, '2p_bridge_too_far', 'Описание', 'dowss', 1);
INSERT INTO public.maps (ID, Name, Mode, Icon, Description, Author, "Order") VALUES (9, 'Railway (2)', 2, '2p_railway', 'Описание', 'dowss', 1);
INSERT INTO public.maps (ID, Name, Mode, Icon, Description, Author, "Order") VALUES (10, 'Battle Marshes (2)', 2, '2p_battle_marshes', 'Описание', 'dowss', 1);
INSERT INTO public.maps (ID, Name, Mode, Icon, Description, Author, "Order") VALUES (11, 'Colosseum Suicide Pro (2)', 2, '2p_colosseum_suicide_pro', 'Описание', 'dowss', 1);
INSERT INTO public.maps (ID, Name, Mode, Icon, Description, Author, "Order") VALUES (12, 'Blood River (2)', 2, '2p_blood_river', 'Описание', 'dowss', 1);
INSERT INTO public.maps (ID, Name, Mode, Icon, Description, Author, "Order") VALUES (13, 'Blood River Remix (2)', 2, '2p_blood_river_remix', 'Описание', 'dowss', 1);
INSERT INTO public.maps (ID, Name, Mode, Icon, Description, Author, "Order") VALUES (14, 'Antiga Bay (2)', 2, '2p_antiga_bay', 'Описание', 'dowss', 1);
INSERT INTO public.maps (ID, Name, Mode, Icon, Description, Author, "Order") VALUES (15, 'Terror Psyclaw (2)', 2, '2p_terror_psyclaw', 'Описание', 'dowss', 1);
INSERT INTO public.maps (ID, Name, Mode, Icon, Description, Author, "Order") VALUES (16, 'Outer Reaches (2)', 2, '2p_outer_reaches', 'Описание', 'dowss', 1);
INSERT INTO public.maps (ID, Name, Mode, Icon, Description, Author, "Order") VALUES (17, 'Moonbase (2)', 2, '2p_moonbase', 'Описание', 'dowss', 1);
INSERT INTO public.maps (ID, Name, Mode, Icon, Description, Author, "Order") VALUES (18, 'Deadly Fun Archeology (2)', 2, '2p_deadly_fun_archeology', 'Описание', 'dowss', 1);
INSERT INTO public.maps (ID, Name, Mode, Icon, Description, Author, "Order") VALUES (19, 'SugarOasis (2)', 2, '2p_sugaroasis', 'Описание', 'dowss', 1);
INSERT INTO public.maps (ID, Name, Mode, Icon, Description, Author, "Order") VALUES (20, 'Tower Ruins (2)', 2, '2p_tower_ruins', 'Описание', 'dowss', 1);
INSERT INTO public.maps (ID, Name, Mode, Icon, Description, Author, "Order") VALUES (21, 'Meeting of Minds Pro (2)', 2, '2p_meeting_of_minds_pro_lis', 'Описание', 'dowss', 1);
INSERT INTO public.maps (ID, Name, Mode, Icon, Description, Author, "Order") VALUES (22, 'Jungle Morning (2)', 2, '2p_jungle_morning', 'Описание', 'dowss', 1);
INSERT INTO public.maps (ID, Name, Mode, Icon, Description, Author, "Order") VALUES (23, 'Colosseum Suicide (2)', 2, '2p_colosseum_suicide', 'Описание', 'dowss', 1);
INSERT INTO public.maps (ID, Name, Mode, Icon, Description, Author, "Order") VALUES (24, 'Shrine of Excellion (2)', 2, '2p_shrine_of_excellion', 'Описание', 'dowss', 1);
INSERT INTO public.maps (ID, Name, Mode, Icon, Description, Author, "Order") VALUES (25, 'Shrine of excellion [Rem] (2)', 2, '2p_shrine_of_excellion_[Rem]', 'Описание', 'dowss', 1);
INSERT INTO public.maps (ID, Name, Mode, Icon, Description, Author, "Order") VALUES (26, 'Meeting of Minds Pro (2)', 2, '2p_meeting_of_minds_pro', 'Описание', 'dowss', 1);
INSERT INTO public.maps (ID, Name, Mode, Icon, Description, Author, "Order") VALUES (27, 'Emerald River (2)', 2, '2p_emerald_river', 'Описание', 'dowss', 1);
INSERT INTO public.maps (ID, Name, Mode, Icon, Description, Author, "Order") VALUES (28, 'Haines Demise (2)', 2, '2p_haines_demise', 'Описание', 'dowss', 1);
INSERT INTO public.maps (ID, Name, Mode, Icon, Description, Author, "Order") VALUES (29, 'Meeting of Minds Pro (2)', 2, '2p_meeting_of_minds_pro_lis_without_1p', 'Описание', 'dowss', 1);
INSERT INTO public.maps (ID, Name, Mode, Icon, Description, Author, "Order") VALUES (30, 'Meeting of Minds (2)', 2, '2p_meeting_of_minds', 'Описание', 'dowss', 1);
INSERT INTO public.maps (ID, Name, Mode, Icon, Description, Author, "Order") VALUES (31, 'Bloody Hell (2)', 2, '2p_bloody hell', 'Описание', 'dowss', 1);
INSERT INTO public.maps (ID, Name, Mode, Icon, Description, Author, "Order") VALUES (32, 'Bloody Hell (2)', 2, '2p_bloody_hell', 'Описание', 'dowss', 1);
INSERT INTO public.maps (ID, Name, Mode, Icon, Description, Author, "Order") VALUES (33, 'Abandon All Hope (2)', 2, '2p_abandon_all_hope', 'Описание', 'dowss', 1);
INSERT INTO public.maps (ID, Name, Mode, Icon, Description, Author, "Order") VALUES (34, 'Faceoff (2)', 2, '2p_faceoff', 'Описание', 'dowss', 1);
INSERT INTO public.maps (ID, Name, Mode, Icon, Description, Author, "Order") VALUES (35, 'Riverbed (2)', 2, '2p_river_bed', 'Описание', 'dowss', 1);
INSERT INTO public.maps (ID, Name, Mode, Icon, Description, Author, "Order") VALUES (36, 'Tranquilitys End (2)', 2, '2p_tranquilitys_end', 'Описание', 'dowss', 1);
INSERT INTO public.maps (ID, Name, Mode, Icon, Description, Author, "Order") VALUES (37, 'Tranquility''s End [Rem] (2)', 2, '2p_tranquilitys_end_[Rem]', 'Описание', 'dowss', 1);
INSERT INTO public.maps (ID, Name, Mode, Icon, Description, Author, "Order") VALUES (38, 'Tranquilitys End Pro (2)', 2, '2p_tranquilitys_end_pro', 'Описание', 'dowss', 1);
INSERT INTO public.maps (ID, Name, Mode, Icon, Description, Author, "Order") VALUES (39, 'Deadman''s Crossing (2)', 2, '2p_deadmans_crossing', 'Описание', 'dowss', 1);
INSERT INTO public.maps (ID, Name, Mode, Icon, Description, Author, "Order") VALUES (40, 'Tainted Pair (2)', 2, '2p_tainted_pair', 'Описание', 'dowss', 1);
INSERT INTO public.maps (ID, Name, Mode, Icon, Description, Author, "Order") VALUES (41, 'Emperors Valley (2)', 2, '2p_emperors_valley', 'Описание', 'dowss', 1);
INSERT INTO public.maps (ID, Name, Mode, Icon, Description, Author, "Order") VALUES (42, 'Titans Fall (2)', 2, '2p_titan_fall', 'Описание', 'dowss', 1);
INSERT INTO public.maps (ID, Name, Mode, Icon, Description, Author, "Order") VALUES (43, 'Titan''s Fall [Rem] (2)', 2, '2p_titan_fall_[Rem]', 'Описание', 'dowss', 1);
INSERT INTO public.maps (ID, Name, Mode, Icon, Description, Author, "Order") VALUES (44, 'Fear (2)', 2, '2p_fear', 'Описание', 'dowss', 1);
INSERT INTO public.maps (ID, Name, Mode, Icon, Description, Author, "Order") VALUES (45, 'Edemus Gamble (2)', 2, '2p_edemus_gamble', 'Описание', 'dowss', 1);
INSERT INTO public.maps (ID, Name, Mode, Icon, Description, Author, "Order") VALUES (46, 'Tazins Folly (2)', 2, '2p_tazins_folly', 'Описание', 'dowss', 1);
INSERT INTO public.maps (ID, Name, Mode, Icon, Description, Author, "Order") VALUES (47, 'Volatile Ground (2)', 2, '2p_volatile_ground', 'Описание', 'dowss', 1);
INSERT INTO public.maps (ID, Name, Mode, Icon, Description, Author, "Order") VALUES (48, 'Hellfire Canyon (2)', 2, '2p_hellfire_canyon', 'Описание', 'dowss', 1);
INSERT INTO public.maps (ID, Name, Mode, Icon, Description, Author, "Order") VALUES (49, 'Absolute Zero (2)', 2, '2p_absolute_zero', 'Описание', 'dowss', 1);
INSERT INTO public.maps (ID, Name, Mode, Icon, Description, Author, "Order") VALUES (50, 'Meeting of Minds Pro (2)', 2, '2p_meeting_of_minds_pro_5p', 'Описание', 'dowss', 1);
INSERT INTO public.maps (ID, Name, Mode, Icon, Description, Author, "Order") VALUES (51, 'Short Below Zero (2)', 2, '2p_short_below_zero', 'Описание', 'dowss', 1);
INSERT INTO public.maps (ID, Name, Mode, Icon, Description, Author, "Order") VALUES (52, 'Fata Morga (2)', 2, '2p_fata_morgana', 'Описание', 'dowss', 1);
INSERT INTO public.maps (ID, Name, Mode, Icon, Description, Author, "Order") VALUES (53, 'Fallen City (2)', 2, '2p_fallen_city', 'Описание', 'dowss', 1);
INSERT INTO public.maps (ID, Name, Mode, Icon, Description, Author, "Order") VALUES (54, 'Fallen City [Rem] (2)', 2, '2p_fallen_city_[Rem]', 'Описание', 'dowss', 1);
INSERT INTO public.maps (ID, Name, Mode, Icon, Description, Author, "Order") VALUES (55, 'Velvet Duress (2)', 2, '2p_velvet_duress', 'Описание', 'dowss', 1);
INSERT INTO public.maps (ID, Name, Mode, Icon, Description, Author, "Order") VALUES (56, 'Fraziers Demise (2)', 2, '2p_fraziersdemise', 'Описание', 'dowss', 1);
INSERT INTO public.maps (ID, Name, Mode, Icon, Description, Author, "Order") VALUES (57, 'Vortex Plateau (2)', 2, '2p_vortex_plateau', 'Описание', 'dowss', 1);
INSERT INTO public.maps (ID, Name, Mode, Icon, Description, Author, "Order") VALUES (58, 'Chaos gate (2)', 2, '2p_chaos_gate', 'Описание', 'dowss', 1);
INSERT INTO public.maps (ID, Name, Mode, Icon, Description, Author, "Order") VALUES (59, 'Valley of Khorne (2)', 2, '2p_valley_of_khorne', 'Описание', 'dowss', 1);
INSERT INTO public.maps (ID, Name, Mode, Icon, Description, Author, "Order") VALUES (60, 'Tiboraxx (2)', 2, '2p_tiboraxx', 'Описание', 'dowss', 1);
INSERT INTO public.maps (ID, Name, Mode, Icon, Description, Author, "Order") VALUES (61, 'Fortress (3)', 3, '3p_fortress', 'Описание', 'dowss', 1);
INSERT INTO public.maps (ID, Name, Mode, Icon, Description, Author, "Order") VALUES (62, 'Marconia (4)', 4, '4p_marconia', 'Описание', 'dowss', 1);
INSERT INTO public.maps (ID, Name, Mode, Icon, Description, Author, "Order") VALUES (63, 'Quatra (4)', 4, '4p_quatra', 'Описание', 'dowss', 1);
INSERT INTO public.maps (ID, Name, Mode, Icon, Description, Author, "Order") VALUES (64, 'Tartarus Center (4)', 4, '4p_tartarus_center', 'Описание', 'dowss', 1);
INSERT INTO public.maps (ID, Name, Mode, Icon, Description, Author, "Order") VALUES (65, 'Rokclaw Foothills (4)', 4, '4p_rokclaw_foothills', 'Описание', 'dowss', 1);
INSERT INTO public.maps (ID, Name, Mode, Icon, Description, Author, "Order") VALUES (66, 'Saint''s Square (4)', 4, '4p_saints_square', 'Описание', 'dowss', 1);
INSERT INTO public.maps (ID, Name, Mode, Icon, Description, Author, "Order") VALUES (67, 'Tiboraxx (4)', 4, '4p_tiboraxx', 'Описание', 'dowss', 1);
INSERT INTO public.maps (ID, Name, Mode, Icon, Description, Author, "Order") VALUES (68, 'Sad Place (4)', 4, '4p_sad_place', 'Описание', 'dowss', 1);
INSERT INTO public.maps (ID, Name, Mode, Icon, Description, Author, "Order") VALUES (69, 'Biffy''s Peril (4)', 4, '4p_biffys_peril', 'Описание', 'dowss', 1);
INSERT INTO public.maps (ID, Name, Mode, Icon, Description, Author, "Order") VALUES (70, 'Mountain Trail (4)', 4, '4p_mountain_trail', 'Описание', 'dowss', 1);
INSERT INTO public.maps (ID, Name, Mode, Icon, Description, Author, "Order") VALUES (71, 'Tainted Place (4)', 4, '4p_tainted_soul', 'Описание', 'dowss', 1);
INSERT INTO public.maps (ID, Name, Mode, Icon, Description, Author, "Order") VALUES (72, 'Into the Breach (4)', 4, '4p_into_the_breach', 'Описание', 'dowss', 1);
INSERT INTO public.maps (ID, Name, Mode, Icon, Description, Author, "Order") VALUES (73, 'Cape of Despair (4)', 4, '4p_cape_of_despair', 'Описание', 'dowss', 1);
INSERT INTO public.maps (ID, Name, Mode, Icon, Description, Author, "Order") VALUES (74, 'Ice Flow (4)', 4, '4p_ice_flow', 'Описание', 'dowss', 1);
INSERT INTO public.maps (ID, Name, Mode, Icon, Description, Author, "Order") VALUES (75, 'Doom Spiral (4)', 4, '4p_doom_spiral', 'Описание', 'dowss', 1);
INSERT INTO public.maps (ID, Name, Mode, Icon, Description, Author, "Order") VALUES (76, 'GorHael Crater (4)', 4, '4p_gorhael_crater', 'Описание', 'dowss', 1);
INSERT INTO public.maps (ID, Name, Mode, Icon, Description, Author, "Order") VALUES (77, 'Ariel Highlands (4)', 4, '4p_ariel_highlands', 'Описание', 'dowss', 1);
INSERT INTO public.maps (ID, Name, Mode, Icon, Description, Author, "Order") VALUES (78, 'Forgotten Isles (4)', 4, '4p_forgotten_isles', 'Описание', 'dowss', 1);
INSERT INTO public.maps (ID, Name, Mode, Icon, Description, Author, "Order") VALUES (79, 'Gurmuns Pass (4)', 4, '4p_gurmuns_pass', 'Описание', 'dowss', 1);
INSERT INTO public.maps (ID, Name, Mode, Icon, Description, Author, "Order") VALUES (80, 'Temple of Cyrene (6)', 6, '4p_temple_of_arcanum', 'Описание', 'dowss', 1);
INSERT INTO public.maps (ID, Name, Mode, Icon, Description, Author, "Order") VALUES (81, 'Murad Swamplands (4)', 4, '4p_murad_swamplands', 'Описание', 'dowss', 1);
INSERT INTO public.maps (ID, Name, Mode, Icon, Description, Author, "Order") VALUES (82, 'Dread Peak (4)', 4, '4p_dread_peak', 'Описание', 'dowss', 1);
INSERT INTO public.maps (ID, Name, Mode, Icon, Description, Author, "Order") VALUES (83, 'Eres Badlands (4)', 4, '4p_eres_badlands', 'Описание', 'dowss', 1);
INSERT INTO public.maps (ID, Name, Mode, Icon, Description, Author, "Order") VALUES (84, 'Janus Savannah (4)', 4, '4p_janus_savannah', 'Описание', 'dowss', 1);
INSERT INTO public.maps (ID, Name, Mode, Icon, Description, Author, "Order") VALUES (85, 'Janus Savannah Pro (4)', 4, '4p_janus_savannah_pro', 'Описание', 'dowss', 1);
INSERT INTO public.maps (ID, Name, Mode, Icon, Description, Author, "Order") VALUES (86, 'Panrea Lowlands (4)', 4, '4p_panrea_lowlands', 'Описание', 'dowss', 1);
INSERT INTO public.maps (ID, Name, Mode, Icon, Description, Author, "Order") VALUES (87, 'The Skerries (4)', 4, '4p_skerries', 'Описание', 'dowss', 1);
INSERT INTO public.maps (ID, Name, Mode, Icon, Description, Author, "Order") VALUES (88, 'The Skerries Pro (4)', 4, '4p_skerries_pro', 'Описание', 'dowss', 1);
INSERT INTO public.maps (ID, Name, Mode, Icon, Description, Author, "Order") VALUES (89, 'Torrents (4)', 4, '4p_torrents', 'Описание', 'dowss', 1);
INSERT INTO public.maps (ID, Name, Mode, Icon, Description, Author, "Order") VALUES (90, 'Van de Mar Mountains (4)', 4, '4p_van_de_mar_mountains', 'Описание', 'dowss', 1);
INSERT INTO public.maps (ID, Name, Mode, Icon, Description, Author, "Order") VALUES (91, 'Sands of Victory (4)', 4, '4p_sands_of_victory', 'Описание', 'dowss', 1);
INSERT INTO public.maps (ID, Name, Mode, Icon, Description, Author, "Order") VALUES (92, 'Antiga Bay (4)', 4, '4p_antiga_bay', 'Описание', 'dowss', 1);
INSERT INTO public.maps (ID, Name, Mode, Icon, Description, Author, "Order") VALUES (93, 'Testcake [Rem] (4)', 4, '4p_testcake', 'Описание', 'dowss', 1);
INSERT INTO public.maps (ID, Name, Mode, Icon, Description, Author, "Order") VALUES (94, 'Imperial Area [Ed] (4)', 4, '4p_imperial_area', 'Описание', 'dowss', 1);
INSERT INTO public.maps (ID, Name, Mode, Icon, Description, Author, "Order") VALUES (95, 'Volcanic Reaction (4)', 4, '4p_volcanic_reaction', 'Описание', 'dowss', 1);
INSERT INTO public.maps (ID, Name, Mode, Icon, Description, Author, "Order") VALUES (96, 'Broken Lands (4)', 4, '4p_broken_lands', 'Описание', 'dowss', 1);
INSERT INTO public.maps (ID, Name, Mode, Icon, Description, Author, "Order") VALUES (97, 'Vyasastan (5)', 5, '5p_vyasastan', 'Описание', 'dowss', 1);
INSERT INTO public.maps (ID, Name, Mode, Icon, Description, Author, "Order") VALUES (98, 'Totmachers Prison (5)', 5, '5p_totmachers_prison', 'Описание', 'dowss', 1);
INSERT INTO public.maps (ID, Name, Mode, Icon, Description, Author, "Order") VALUES (99, 'Red Jungle (5)', 5, '5p_red_jungle', 'Описание', 'dowss', 1);
INSERT INTO public.maps (ID, Name, Mode, Icon, Description, Author, "Order") VALUES (100, 'Aceria Forests (5)', 5, '5p_aceria_forests', 'Описание', 'dowss', 1);
INSERT INTO public.maps (ID, Name, Mode, Icon, Description, Author, "Order") VALUES (101, 'Eye of gorgon (5)', 5, '5p_eye_of_gorgon.jpg', 'Описание', 'dowss', 1);
INSERT INTO public.maps (ID, Name, Mode, Icon, Description, Author, "Order") VALUES (102, 'Parmenie (6)', 6, '6p_parmenie', 'Описание', 'dowss', 1);
INSERT INTO public.maps (ID, Name, Mode, Icon, Description, Author, "Order") VALUES (103, 'Temple of Cyrene (6)', 6, '6p_temple_cyrene', 'Описание', 'dowss', 1);
INSERT INTO public.maps (ID, Name, Mode, Icon, Description, Author, "Order") VALUES (104, 'Nirraein (6)', 6, '6p_nirraein', 'Описание', 'dowss', 1);
INSERT INTO public.maps (ID, Name, Mode, Icon, Description, Author, "Order") VALUES (105, 'Shakun Coast (6)', 6, '6p_shakun_coast', 'Описание', 'dowss', 1);
INSERT INTO public.maps (ID, Name, Mode, Icon, Description, Author, "Order") VALUES (106, 'Streets of Vogen (6)', 6, '6pteam_streets_of_vogen', 'Описание', 'dowss', 1);
INSERT INTO public.maps (ID, Name, Mode, Icon, Description, Author, "Order") VALUES (107, 'Tristram Plains (6)', 6, '6p_tristram_plains', 'Описание', 'dowss', 1);
INSERT INTO public.maps (ID, Name, Mode, Icon, Description, Author, "Order") VALUES (108, 'Pavonis (6)', 6, '6p_pavonis', 'Описание', 'dowss', 1);
INSERT INTO public.maps (ID, Name, Mode, Icon, Description, Author, "Order") VALUES (109, 'Pavonian Heartland (6)', 6, '6p_pavonian_heartland', 'Описание', 'dowss', 1);
INSERT INTO public.maps (ID, Name, Mode, Icon, Description, Author, "Order") VALUES (110, 'Paynes Retribution (6)', 6, '6p_paynes_retribution', 'Описание', 'dowss', 1);
INSERT INTO public.maps (ID, Name, Mode, Icon, Description, Author, "Order") VALUES (111, 'Thargorum (6)', 6, '6p_thargorum', 'Описание', 'dowss', 1);
INSERT INTO public.maps (ID, Name, Mode, Icon, Description, Author, "Order") VALUES (112, 'Parmenian Heath (6)', 6, '6p_parmenian_heath', 'Описание', 'dowss', 1);
INSERT INTO public.maps (ID, Name, Mode, Icon, Description, Author, "Order") VALUES (113, 'Jungle Walls (6)', 6, '6p_jungle_walls', 'Описание', 'dowss', 1);
INSERT INTO public.maps (ID, Name, Mode, Icon, Description, Author, "Order") VALUES (114, 'Trivian Groves (6)', 6, '6p_trivian_groves', 'Описание', 'dowss', 1);
INSERT INTO public.maps (ID, Name, Mode, Icon, Description, Author, "Order") VALUES (115, 'Kasyr Lutien (6)', 6, '6p_kasyr_lutien', 'Описание', 'dowss', 1);
INSERT INTO public.maps (ID, Name, Mode, Icon, Description, Author, "Order") VALUES (116, 'Mortalis (6)', 6, '6p_mortalis', 'Описание', 'dowss', 1);
INSERT INTO public.maps (ID, Name, Mode, Icon, Description, Author, "Order") VALUES (117, 'Irridene (6)', 6, '6p_irridene', 'Описание', 'dowss', 1);
INSERT INTO public.maps (ID, Name, Mode, Icon, Description, Author, "Order") VALUES (118, 'Fury Island (6)', 6, '6p_fury_island', 'Описание', 'dowss', 1);
INSERT INTO public.maps (ID, Name, Mode, Icon, Description, Author, "Order") VALUES (119, 'Confrontation (6)', 6, '6p_confrontation', 'Описание', 'dowss', 1);
INSERT INTO public.maps (ID, Name, Mode, Icon, Description, Author, "Order") VALUES (120, 'Crozius Arcanum (6)', 6, '6p_crozius_arcanum', 'Описание', 'dowss', 1);
INSERT INTO public.maps (ID, Name, Mode, Icon, Description, Author, "Order") VALUES (121, 'Dread Alley (6)', 6, '6p_dread_alley', 'Описание', 'dowss', 1);
INSERT INTO public.maps (ID, Name, Mode, Icon, Description, Author, "Order") VALUES (122, 'Vandean Coast (6)', 6, '6p_vandea_coast', 'Описание', 'dowss', 1);
INSERT INTO public.maps (ID, Name, Mode, Icon, Description, Author, "Order") VALUES (123, 'Orestan Plains (6)', 6, '6p_orestan_plains', 'Описание', 'dowss', 1);
INSERT INTO public.maps (ID, Name, Mode, Icon, Description, Author, "Order") VALUES (124, 'Western Barrens (6)', 6, '6p_western_barrens', 'Описание', 'dowss', 1);
INSERT INTO public.maps (ID, Name, Mode, Icon, Description, Author, "Order") VALUES (125, 'Principian Badlands (6)', 6, '6p_principian_badlands', 'Описание', 'dowss', 1);
INSERT INTO public.maps (ID, Name, Mode, Icon, Description, Author, "Order") VALUES (126, 'Rhean Floodlands (6)', 6, '6p_rhean_floodlands', 'Описание', 'dowss', 1);
INSERT INTO public.maps (ID, Name, Mode, Icon, Description, Author, "Order") VALUES (127, 'Hyperion Peaks (6)', 6, '6p_hyperion_peaks', 'Описание', 'dowss', 1);
INSERT INTO public.maps (ID, Name, Mode, Icon, Description, Author, "Order") VALUES (128, 'Crossroads (6)', 6, '6p_crossroads', 'Описание', 'dowss', 1);
INSERT INTO public.maps (ID, Name, Mode, Icon, Description, Author, "Order") VALUES (129, 'Testing Grounds (6)', 6, '6p_testing_grounds', 'Описание', 'dowss', 1);
INSERT INTO public.maps (ID, Name, Mode, Icon, Description, Author, "Order") VALUES (130, 'Bloodshed Alley (6)', 6, '6p_bloodshed_alley', 'Описание', 'dowss', 1);
INSERT INTO public.maps (ID, Name, Mode, Icon, Description, Author, "Order") VALUES (131, 'Agamar Desert (6)', 6, '6p_agamar_desert', 'Описание', 'dowss', 1);
INSERT INTO public.maps (ID, Name, Mode, Icon, Description, Author, "Order") VALUES (132, 'Alvarus (6)', 6, '6p_alvarus', 'Описание', 'dowss', 1);
INSERT INTO public.maps (ID, Name, Mode, Icon, Description, Author, "Order") VALUES (133, 'Kaurav City (6)', 6, '6p_kaurav_city', 'Описание', 'dowss', 1);
INSERT INTO public.maps (ID, Name, Mode, Icon, Description, Author, "Order") VALUES (134, 'Snowblind (6)', 6, '6p_snowblind', 'Описание', 'dowss', 1);
INSERT INTO public.maps (ID, Name, Mode, Icon, Description, Author, "Order") VALUES (135, 'Ruined Greatway (6)', 6, '6p_ruined_greatway', 'Описание', 'dowss', 1);
INSERT INTO public.maps (ID, Name, Mode, Icon, Description, Author, "Order") VALUES (136, 'Forbidden Jungle (8)', 8, '8p_forbidden_jungle', 'Описание', 'dowss', 1);
INSERT INTO public.maps (ID, Name, Mode, Icon, Description, Author, "Order") VALUES (137, 'Rhean Jungle (8)', 8, '8p_rhean_jungle', 'Описание', 'dowss', 1);
INSERT INTO public.maps (ID, Name, Mode, Icon, Description, Author, "Order") VALUES (138, 'ThurAbis Plateau (8)', 8, '8p_thurabis_plateau', 'Описание', 'dowss', 1);
INSERT INTO public.maps (ID, Name, Mode, Icon, Description, Author, "Order") VALUES (139, 'Penal Colony (8)', 8, '8p_penal_colony', 'Описание', 'dowss', 1);
INSERT INTO public.maps (ID, Name, Mode, Icon, Description, Author, "Order") VALUES (140, 'Morriah Coast (8)', 8, '8p_morriah_coast', 'Описание', 'dowss', 1);
INSERT INTO public.maps (ID, Name, Mode, Icon, Description, Author, "Order") VALUES (141, 'Demes Northlands (8)', 8, '8p_demes_northlands', 'Описание', 'dowss', 1);
INSERT INTO public.maps (ID, Name, Mode, Icon, Description, Author, "Order") VALUES (142, 'Morholt Range (8)', 8, '8p_morholt_range', 'Описание', 'dowss', 1);
INSERT INTO public.maps (ID, Name, Mode, Icon, Description, Author, "Order") VALUES (143, 'Monse (8)', 8, '8p_clooth-na-bare', 'Описание', 'dowss', 1);
INSERT INTO public.maps (ID, Name, Mode, Icon, Description, Author, "Order") VALUES (144, 'Oasis of Sharr (8)', 8, '8p_oasis_of_sharr', 'Описание', 'dowss', 1);
INSERT INTO public.maps (ID, Name, Mode, Icon, Description, Author, "Order") VALUES (145, 'Doom Chamber (8)', 8, '8p_doom_chamber', 'Описание', 'dowss', 1);
INSERT INTO public.maps (ID, Name, Mode, Icon, Description, Author, "Order") VALUES (146, 'Daturias Pits (8)', 8, '8p_daturias_pits', 'Описание', 'dowss', 1);
INSERT INTO public.maps (ID, Name, Mode, Icon, Description, Author, "Order") VALUES (147, 'Burial Grounds (8)', 8, '8p_burial_grounds', 'Описание', 'dowss', 1);
INSERT INTO public.maps (ID, Name, Mode, Icon, Description, Author, "Order") VALUES (148, 'Jalaganda Lowlands (8)', 8, '8p_jalaganda_lowlands', 'Описание', 'dowss', 1);
INSERT INTO public.maps (ID, Name, Mode, Icon, Description, Author, "Order") VALUES (149, 'Lost Hope (8)', 8, '8p_lost_hope', 'Описание', 'dowss', 1);
INSERT INTO public.maps (ID, Name, Mode, Icon, Description, Author, "Order") VALUES (150, 'Cerulea (8)', 8, '8p_cerulea', 'Описание', 'dowss', 1);
INSERT INTO public.maps (ID, Name, Mode, Icon, Description, Author, "Order") VALUES (151, 'Kier Harrad (8)', 8, '8p_kier_harrad', 'Описание', 'dowss', 1);
INSERT INTO public.maps (ID, Name, Mode, Icon, Description, Author, "Order") VALUES (152, '[TP MOD]Ice Flow (4)', 4, '[tp_mod]4p_ice_flow', 'Описание', 'dowss', 1);
INSERT INTO public.maps (ID, Name, Mode, Icon, Description, Author, "Order") VALUES (153, '[TP MOD] Edemus Gamble (2)', 2, '[tp_mod]edemus_gamble', 'Описание', 'dowss', 1);
INSERT INTO public.maps (ID, Name, Mode, Icon, Description, Author, "Order") VALUES (154, '[TP MOD] Galenas Crusade (2)', 2, '[tp_mod]galenas_crusade', 'Описание', 'dowss', 1);
INSERT INTO public.maps (ID, Name, Mode, Icon, Description, Author, "Order") VALUES (155, '[TP MOD] Jungle Morning (2)', 2, '[tp mod]jungle_morning', 'Описание', 'dowss', 1);
INSERT INTO public.maps (ID, Name, Mode, Icon, Description, Author, "Order") VALUES (156, '[TP MOD] Light Brigade (2)', 2, '[tp_mod]light_brigade', 'Описание', 'dowss', 1);
INSERT INTO public.maps (ID, Name, Mode, Icon, Description, Author, "Order") VALUES (157, '[TP MOD] Quests Triumph (2)', 2, '[tp_mod]quests_triumph', 'Описание', 'dowss', 1);
INSERT INTO public.maps (ID, Name, Mode, Icon, Description, Author, "Order") VALUES (158, '[TP MOD] Shrine of Excellion (2)', 2, '[tp_mod]shrine_of_excellion', 'Описание', 'dowss', 1);
INSERT INTO public.maps (ID, Name, Mode, Icon, Description, Author, "Order") VALUES (159, '[TP MOD] Shrine of Excellion (2)', 2, '[tp_mod]shrine of excellion', 'Описание', 'dowss', 1);
INSERT INTO public.maps (ID, Name, Mode, Icon, Description, Author, "Order") VALUES (160, '[TP MOD] Meeting of Minds (2)', 2, '[tp_mod]meeting_of_minds', 'Описание', 'dowss', 1);