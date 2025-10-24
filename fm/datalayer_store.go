package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"sort"
	"time"
)

const (
	SaveDirName     = "rogue_saves"
	LastSessionFile = "last_session.json"
	AttemptsFile    = "game_attempts.json"
)

type Store struct {
	saveDir string
}

func NewStore() (*Store, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("failed to get home directory: %w", err)
	}

	saveDir := filepath.Join(homeDir, SaveDirName)

	if err := os.MkdirAll(saveDir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create save directory: %w", err)
	}

	return &Store{saveDir: saveDir}, nil
}

func (s *Store) SaveSession(session *Session) error {
	session.Timestamp = time.Now().Unix()

	data, err := json.MarshalIndent(session, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal session: %w", err)
	}

	filePath := filepath.Join(s.saveDir, LastSessionFile)
	if err := os.WriteFile(filePath, data, 0644); err != nil {
		return fmt.Errorf("failed to write session file: %w", err)
	}

	return nil
}

func (s *Store) LoadSession() (*Session, error) {
	filePath := filepath.Join(s.saveDir, LastSessionFile)

	data, err := os.ReadFile(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to read session file: %w", err)
	}

	var session Session
	if err := json.Unmarshal(data, &session); err != nil {
		return nil, fmt.Errorf("failed to unmarshal session: %w", err)
	}

	return &session, nil
}

func (s *Store) SaveGameAttempt(attempt SavedGames) error {
	attemptsData, err := s.loadAttempts()
	if err != nil {
		return fmt.Errorf("failed to load existing attempts: %w", err)
	}

	attemptsData.Attempts = append(attemptsData.Attempts, attempt)
	attemptsData.NextID++

	sort.Slice(attemptsData.Attempts, func(i, j int) bool {
		return attemptsData.Attempts[i].Gold > attemptsData.Attempts[j].Gold
	})

	return s.saveAttempts(attemptsData)
}

func (s *Store) loadAttempts() (*AttemptsData, error) {
	filePath := filepath.Join(s.saveDir, AttemptsFile)

	data, err := os.ReadFile(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return &AttemptsData{
				Attempts: make([]SavedGames, 0),
				NextID:   1,
			}, nil
		}
		return nil, fmt.Errorf("failed to read attempts file: %w", err)
	}

	var attemptsData AttemptsData
	if err := json.Unmarshal(data, &attemptsData); err != nil {
		return nil, fmt.Errorf("failed to unmarshal attempts: %w", err)
	}

	return &attemptsData, nil
}

func (s *Store) saveAttempts(attemptsData *AttemptsData) error {
	data, err := json.MarshalIndent(attemptsData, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal attempts: %w", err)
	}

	filePath := filepath.Join(s.saveDir, AttemptsFile)
	if err := os.WriteFile(filePath, data, 0644); err != nil {
		return fmt.Errorf("failed to write attempts file: %w", err)
	}

	return nil
}

func recordAttempt(dungeon *dungeon_t, player *player_t) {
	var AttemptToSave SavedGames
	AttemptToSave.FoodConsumed = player.foodConsumed
	AttemptToSave.Gold = player.gold
	AttemptToSave.Level = dungeon.level
	AttemptToSave.MonsterKill = player.monsterKill
	AttemptToSave.PotionsConsumed = player.potionsConsumed
	AttemptToSave.ScrollsRead = player.scrollsRead
	AttemptToSave.StrikesToEnemy = player.strikesToEnemy
	AttemptToSave.StrikesToPlayer = player.strikesToPlayer
	AttemptToSave.Turns = player.turns

	if store, err := NewStore(); err == nil {
		store.SaveGameAttempt(AttemptToSave)
	}
}

func saveSessionFromGame(dungeon *dungeon_t, field *map_t, player *player_t) *Session {
	var SessionToSave Session
	var Dungeon Dungeon_t
	var Field Map_t
	var Player Player_t
	SessionToSave.Dungeon = &Dungeon
	SessionToSave.Field = &Field
	SessionToSave.Player = &Player
	SessionToSave.Dungeon.Room_cnt = dungeon.room_cnt
	SessionToSave.Dungeon.Corridors_cnt = dungeon.corridors_cnt
	SessionToSave.Dungeon.Level = dungeon.level

	var cnt int
	for i := 1; i < 4; i++ {
		for j := 1; j < 4; j++ {
			convertRoom(&dungeon.rooms[i][j], &SessionToSave.Dungeon.Rooms[i][j])
			Room := SessionToSave.Dungeon.Rooms[i][j]
			SessionToSave.Dungeon.Sequence[cnt] = Room
			cnt++
		}
	}

	for i := range dungeon.corridors {
		convertCorridor(&dungeon.corridors[i], &SessionToSave.Dungeon.Corridors[i])
	}

	for i := range dungeon.lockedDoors {
		SessionToSave.Dungeon.LockedDoors[i].Y, SessionToSave.Dungeon.LockedDoors[i].X = dungeon.lockedDoors[i].y, dungeon.lockedDoors[i].x
	}

	convertField(field, SessionToSave.Field)
	convertPlayer(player, SessionToSave.Player)

	return &SessionToSave

}

func convertField(field *map_t, Field *Map_t) {
	Field.Playground = field.playground
	Field.Items_cnt = field.items_cnt
	Field.Enemies_cnt = field.enemies_cnt
	convertEntity(&field.player_spawn, &Field.Player_spawn)
	convertEntity(&field.exit, &Field.Exit)
	for i := range field.items {
		convertEntity(&field.items[i], &Field.Items[i])
	}
	for i := range field.enemies {
		convertEntity(&field.enemies[i], &Field.Enemies[i])
	}

	for i := range field.visited {
		var pos Position_t
		pos.Y, pos.X = i.y, i.x
		Field.VisitedSlice = append(Field.VisitedSlice, pos)
	}
}

func convertPlayer(player *player_t, Player *Player_t) {
	Player.Agility = player.agility
	Player.FoodConsumed = player.foodConsumed
	Player.Gold = player.gold
	Player.GotBlue = player.gotBlue
	Player.GotCyan = player.gotCyan
	Player.GotMagenta = player.gotMagenta
	Player.Health = player.health
	Player.IsSleeped = player.isSleeped
	Player.MaxHealth = player.maxHealth
	Player.MonsterKill = player.monsterKill
	Player.Pos.Y, Player.Pos.X = player.pos.y, player.pos.x
	Player.PotAgility = player.potAgility
	Player.PotMaxHealth = player.potMaxHealth
	Player.PotStrenght = player.potStrenght
	Player.PotionsConsumed = player.potionsConsumed
	Player.FoodConsumed = player.foodConsumed
	Player.ScrollsRead = player.scrollsRead
	Player.Strength = player.strength
	Player.StrikesToEnemy = player.strikesToEnemy
	Player.StrikesToPlayer = player.strikesToPlayer
	Player.Turns = player.turns
	Player.Weapon = player.weapon
	Player.Inventory.Food = player.inventory.food
	Player.Inventory.Potion = player.inventory.potion
	Player.Inventory.Scroll = player.inventory.scroll
	Player.Inventory.Weapon = player.inventory.weapon
}

func convertCorridor(corridor *corridor_t, Corridor *Corridor_t) {
	Corridor.TypeE = corridor.typeE
	Corridor.Points_cnt = corridor.points_cnt
	for i := range corridor.points {
		Corridor.Points[i].Y, Corridor.Points[i].X = corridor.points[i].y, corridor.points[i].x

	}
}

func convertRoom(room *room_t, Room *Room_t) {
	Room.Sector = room.sector
	Room.Grid_i = room.grid_i
	Room.Grid_j = room.grid_j
	Room.Top_left.Y, Room.Top_left.X, Room.Bot_right.Y, Room.Bot_right.X = room.top_left.y, room.top_left.x, room.bot_right.y, room.bot_right.x
	for i := range room.doors {
		Room.Doors[i].Y, Room.Doors[i].X = room.doors[i].y, room.doors[i].x
	}
	for i := range room.entities {
		convertEntity(&room.entities[i], &Room.Entities[i])
	}
	Room.Entities_cnt = room.entities_cnt
	Room.IsStart = room.isStart
	Room.IsLocked = room.isLocked
	Room.BlueLock = room.blueLock
	Room.MagentaLock = room.magentaLock
	Room.CyanLock = room.cyanLock
}

func convertEntity(entity *entity_t, Entity *Entity_t) {
	Entity.TypeE = entity.typeE
	Entity.TypeI = entity.typeI
	Entity.Symbol = entity.symbol
	Entity.Pos.Y, Entity.Pos.X = entity.pos.y, entity.pos.x
	convertStats(entity, Entity)
}

func convertStats(entity *entity_t, Entity *Entity_t) {
	Entity.Stats.Strength = entity.stats.strength
	Entity.Stats.Agility = entity.stats.agility
	Entity.Stats.Hp = entity.stats.hp
	Entity.Stats.Aggression = entity.stats.aggression
	Entity.Stats.ExtraSym = entity.stats.extraSym
	Entity.Stats.HpStealed = entity.stats.hpStealed
	Entity.Stats.IsChasing = entity.stats.isChasing
	Entity.Stats.IsFirst = entity.stats.isFirst
	Entity.Stats.LastMove = entity.stats.lastMove

}

func LoadGame() (*dungeon_t, *map_t, *player_t, error) {
	store, err := NewStore()
	if err != nil {
		return nil, nil, nil, err
	}

	savedSession, err := store.LoadSession()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
		return nil, nil, nil, err
	}

	if savedSession == nil {
		return nil, nil, nil, nil
	}

	dungeon, field, player := loadSessionToGame(savedSession)
	return dungeon, field, player, err

}

func loadSessionToGame(session *Session) (*dungeon_t, *map_t, *player_t) {
	var dungeon dungeon_t
	var field map_t
	var player player_t
	dungeon.level = session.Dungeon.Level
	dungeon.room_cnt = session.Dungeon.Room_cnt
	dungeon.corridors_cnt = session.Dungeon.Corridors_cnt
	var cnt int

	for i := 1; i < 4; i++ {
		for j := 1; j < 4; j++ {
			recoverRoom(&dungeon.rooms[i][j], &session.Dungeon.Rooms[i][j])
			dungeon.sequence[cnt] = &dungeon.rooms[i][j]
			cnt++
		}
	}

	for i := range session.Dungeon.Corridors {
		recoverCorridor(&dungeon.corridors[i], &session.Dungeon.Corridors[i])
	}

	for i := range session.Dungeon.LockedDoors {
		dungeon.lockedDoors[i].y, dungeon.lockedDoors[i].x = session.Dungeon.LockedDoors[i].Y, session.Dungeon.LockedDoors[i].X
	}

	recoverField(&field, session.Field)
	recoverPlayer(&player, session.Player)

	return &dungeon, &field, &player

}

func recoverRoom(room *room_t, Room *Room_t) {
	room.sector = Room.Sector
	room.grid_i = Room.Grid_i
	room.grid_j = Room.Grid_j
	room.top_left.y, room.top_left.x, room.bot_right.y, room.bot_right.x = Room.Top_left.Y, Room.Top_left.X, Room.Bot_right.Y, Room.Bot_right.X
	for i := range Room.Doors {
		room.doors[i].y, room.doors[i].x = Room.Doors[i].Y, Room.Doors[i].X
	}
	for i := range Room.Entities {
		recoverEntity(&room.entities[i], &Room.Entities[i])
	}
	room.entities_cnt = Room.Entities_cnt
	room.isStart = Room.IsStart
	room.isLocked = Room.IsLocked
	room.blueLock = Room.BlueLock
	room.magentaLock = Room.MagentaLock
	room.cyanLock = Room.CyanLock
}

func recoverEntity(entity *entity_t, Entity *Entity_t) {
	entity.typeE = Entity.TypeE
	entity.typeI = Entity.TypeI
	entity.symbol = Entity.Symbol
	entity.pos.y, entity.pos.x = Entity.Pos.Y, Entity.Pos.X
	recoverStats(entity, Entity)
}

func recoverStats(entity *entity_t, Entity *Entity_t) {
	entity.stats.strength = Entity.Stats.Strength
	entity.stats.agility = Entity.Stats.Agility
	entity.stats.hp = Entity.Stats.Hp
	entity.stats.aggression = Entity.Stats.Aggression
	entity.stats.extraSym = Entity.Stats.ExtraSym
	entity.stats.hpStealed = Entity.Stats.HpStealed
	entity.stats.isChasing = Entity.Stats.IsChasing
	entity.stats.isFirst = Entity.Stats.IsFirst
	entity.stats.lastMove = Entity.Stats.LastMove
}

func recoverCorridor(corridor *corridor_t, Corridor *Corridor_t) {
	corridor.typeE = Corridor.TypeE
	corridor.points_cnt = Corridor.Points_cnt
	for i := range Corridor.Points {
		corridor.points[i].y, corridor.points[i].x = Corridor.Points[i].Y, Corridor.Points[i].X
	}
}

func recoverField(field *map_t, Field *Map_t) {
	field.playground = Field.Playground
	field.items_cnt = Field.Items_cnt
	field.enemies_cnt = Field.Enemies_cnt
	recoverEntity(&field.player_spawn, &Field.Player_spawn)
	recoverEntity(&field.exit, &Field.Exit)
	for i := range Field.Items {
		recoverEntity(&field.items[i], &Field.Items[i])
	}
	for i := range Field.Enemies {
		recoverEntity(&field.enemies[i], &Field.Enemies[i])
	}
	field.visited = make(map[position_t]bool)

	for i := range Field.VisitedSlice {
		var pos position_t
		pos.y, pos.x = Field.VisitedSlice[i].Y, Field.VisitedSlice[i].X
		field.visited[pos] = true
	}
}

func recoverPlayer(player *player_t, Player *Player_t) {
	player.agility = Player.Agility
	player.foodConsumed = Player.FoodConsumed
	player.gold = Player.Gold
	player.gotBlue = Player.GotBlue
	player.gotCyan = Player.GotCyan
	player.gotMagenta = Player.GotMagenta
	player.health = Player.Health
	player.isSleeped = Player.IsSleeped
	player.maxHealth = Player.MaxHealth
	player.monsterKill = Player.MonsterKill
	player.pos.y, player.pos.x = Player.Pos.Y, Player.Pos.X
	player.potAgility = Player.PotAgility
	player.potMaxHealth = Player.PotMaxHealth
	player.potStrenght = Player.PotStrenght
	player.potionsConsumed = Player.PotionsConsumed
	player.foodConsumed = Player.FoodConsumed
	player.scrollsRead = Player.ScrollsRead
	player.strength = Player.Strength
	player.strikesToEnemy = Player.StrikesToEnemy
	player.strikesToPlayer = Player.StrikesToPlayer
	player.turns = Player.Turns
	player.weapon = Player.Weapon
	player.inventory.food = Player.Inventory.Food
	player.inventory.potion = Player.Inventory.Potion
	player.inventory.scroll = Player.Inventory.Scroll
}
