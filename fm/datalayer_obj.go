package main

type Player_t struct {
	Pos             Position_t
	Inventory       Inventory_t
	Turns           int
	Weapon          int
	Health          int
	Strength        int
	Agility         int
	MaxHealth       int
	PotStrenght     int
	PotAgility      int
	PotMaxHealth    int
	Gold            int
	IsSleeped       bool
	GotBlue         bool
	GotMagenta      bool
	GotCyan         bool
	MonsterKill     int
	FoodConsumed    int
	PotionsConsumed int
	ScrollsRead     int
	StrikesToEnemy  int
	StrikesToPlayer int
}

type Position_t struct {
	X int
	Y int
}

type Entity_t struct {
	TypeE  int
	Symbol int
	TypeI  int
	Pos    Position_t
	Stats  Stats_t
}

type Map_t struct {
	Playground   [MAP_HEIGHT][MAP_WIDTH]int
	Player_spawn Entity_t
	Exit         Entity_t
	Items        [MAX_ITEMS_TOTAL]Entity_t
	Items_cnt    int
	Enemies      [MAX_ENEMIES_TOTAL]Entity_t
	Enemies_cnt  int
	// Visited      map[Position_t]bool
	VisitedSlice []Position_t
}

type Room_t struct {
	Sector       int
	Grid_i       int
	Grid_j       int
	Top_left     Position_t
	Bot_right    Position_t
	Doors        [4]Position_t
	Connections  [4]*Room_t
	Entities     [MAX_ENTITIES_PER_ROOM]Entity_t
	Entities_cnt int
	IsStart      bool
	IsLocked     bool
	BlueLock     bool
	MagentaLock  bool
	CyanLock     bool
}

type Corridor_t struct {
	TypeE      int
	Points     [4]Position_t
	Points_cnt int
}

type Dungeon_t struct {
	Room_cnt      int
	Corridors_cnt int
	Rooms         [ROOMS_PER_SIDE + 2][ROOMS_PER_SIDE + 2]Room_t
	Sequence      [MAX_ROOMS_NUMBER]Room_t
	Corridors     [MAX_CORRIDORS_NUMBER]Corridor_t
	Level         int
	LockedDoors   [3]Position_t
}

type Inventory_t struct {
	Weapon []int
	Scroll []int
	Potion []int
	Food   []int
}

type Stats_t struct {
	Strength   int
	Agility    int
	Hp         int
	IsChasing  bool
	IsFirst    bool
	Aggression int
	LastMove   int
	HpStealed  int
	ExtraSym   int
}

type Session struct {
	Dungeon   *Dungeon_t
	Field     *Map_t
	Player    *Player_t
	Timestamp int64
}

type SavedGames struct {
	Gold            int
	Level           int
	Turns           int
	MonsterKill     int
	FoodConsumed    int
	ScrollsRead     int
	PotionsConsumed int
	StrikesToEnemy  int
	StrikesToPlayer int
}

type AttemptsData struct {
	Attempts []SavedGames
	NextID   int
}
