package game

var dataSprites = map[string]SpriteInfo{
	"background": {
		X: 1494,
		Y: 44,
		W: 320,
		H: 200,
	},
	"play_button": {
		X: 236,
		Y: 58,
		W: 61,
		H: 19,
	},
	"play_button_hover": {
		X: 298,
		Y: 58,
		W: 61,
		H: 19,
	},

	// Person stuff

	"person_player": {
		X: 443,
		Y: 31,
		W: 32,
		H: 48,
	},
	"person_marker": {
		X: 301,
		Y: 134,
		W: 32,
		H: 48,
	},
	"person_selection": {
		X: 301,
		Y: 82,
		W: 32,
		H: 48,
	},
	"person_alien_gray": {
		X: 443,
		Y: 79,
		W: 32,
		H: 48,
	},
	"person_alien_buddy": {
		X: 443,
		Y: 127,
		W: 32,
		H: 48,
	},
	"person_alien_ferengi": {
		X: 443,
		Y: 175,
		W: 32,
		H: 48,
	},
	"person_alien_robot": {
		X: 443,
		Y: 223,
		W: 32,
		H: 48,
	},

	// Cheat stuff

	"cheat_marker": {
		X: 14,
		Y: 154,
		W: 30,
		H: 30,
	},
	"cheat_fart": {
		X: 8,
		Y: 244,
		W: 24,
		H: 24,
	},
	"cheat_shove": {
		X: 8,
		Y: 268,
		W: 24,
		H: 24,
	},
	"cheat_persuade": {
		X: 8,
		Y: 292,
		W: 24,
		H: 24,
	},
	"cheat_bomb_threat": {
		X: 8,
		Y: 316,
		W: 24,
		H: 24,
	},
	"cheat_confusion": {
		X: 8,
		Y: 340,
		W: 24,
		H: 24,
	},
	"cheat_circuit_failure": {
		X: 8,
		Y: 364,
		W: 24,
		H: 24,
	},
	"cheat_bribe": {
		X: 8,
		Y: 388,
		W: 24,
		H: 24,
	},
	"button_discard": {
		X: 8,
		Y: 412,
		W: 24,
		H: 24,
	},
	"button_pass": {
		X: 8,
		Y: 436,
		W: 24,
		H: 24,
	},

	// Beam

	"beam_start": {
		X: 301,
		Y: 187,
		W: 32,
		H: 48,
	},
	"beam_middle": {
		X: 301,
		Y: 236,
		W: 32,
		H: 48,
	},
	"beam_end": {
		X: 301,
		Y: 285,
		W: 32,
		H: 48,
	},

	// Kill chamber

	"kill_prolog": {
		X: 301,
		Y: 342,
		W: 32,
		H: 48,
	},
	"kill_blazing": {
		X: 301,
		Y: 391,
		W: 32,
		H: 48,
	},
	"kill_fading": {
		X: 301,
		Y: 440,
		W: 32,
		H: 48,
	},
	"kill_chamber_ground": {
		X: 574,
		Y: 142,
		W: 20,
		H: 9,
	},
	"kill_chamber_ground_active": {
		X: 574,
		Y: 152,
		W: 20,
		H: 9,
	},
}
