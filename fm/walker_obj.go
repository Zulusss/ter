package main

const SCRN_W = 450
const SCRN_H = 120

type screen_t struct {
	screen [SCRN_H][SCRN_W]int
}

type player_t struct {
	pos             position_t
	inventory       inventory_t
	turns           int
	weapon          int
	health          int
	strength        int
	agility         int
	maxHealth       int
	potStrenght     int
	potAgility      int
	potMaxHealth    int
	gold            int
	isSleeped       bool
	gotBlue         bool
	gotMagenta      bool
	gotCyan         bool
	monsterKill     int
	foodConsumed    int
	potionsConsumed int
	scrollsRead     int
	strikesToEnemy  int
	strikesToPlayer int
	// exp          int
	// float view_angle;
	// float fov;
	// float view_distance;
}
