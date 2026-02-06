package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/xxLYNX/warframe_calculator/wfdata"
)

// Entity represents any entity in the game
type Entity struct {
	Name      string  `json:"name"`
	Faction   string  `json:"faction"`
	Health    float64 `json:"health"`
	Shields   float64 `json:"shields"`
	Armor     float64 `json:"armor"`
	Energy    float64 `json:"energy,omitempty"`
	Overguard float64 `json:"overguard,omitempty"`
	Sprint    float64 `json:"sprint,omitempty"`
	Level     int     `json:"level"`
}

// Configuration represents a saved entity configuration
type Configuration struct {
	ID                    string        `json:"id"`
	Entity                Entity        `json:"entity"`
	Stats                 Entity        `json:"stats"`
	EntityConfig          []interface{} `json:"entity_config"`
	AbilityConfig         []interface{} `json:"ability_config"`
	PrimaryConfig         []interface{} `json:"primary_config"`
	SecondaryConfig       []interface{} `json:"secondary_config"`
	MeleeConfig           []interface{} `json:"melee_config"`
	ArchgunConfig         []interface{} `json:"archgun_config"`
	OperatorConfig        []interface{} `json:"operator_config"`
	ArchwingConfig        []interface{} `json:"archwing_config"`
	ArchwinggunConfig     []interface{} `json:"archwinggun_config"`
	ArchwingmeleeConfig   []interface{} `json:"archwingmelee_config"`
	CompanionConfig       []interface{} `json:"companion_config"`
	CompanionweaponConfig []interface{} `json:"companionweapon_config"`
}

// Currently spawned entities (in scope for simulation)
var spawnedEntities map[string]Configuration

var entitiesCache map[string]Entity
var cacheTime time.Time

// WarframeManifestItem from official export (simplified)
type WarframeManifestItem struct {
	UniqueName  string  `json:"uniqueName"`
	Name        string  `json:"name"`
	Health      float64 `json:"health"`
	Shield      float64 `json:"shield"`
	Armor       float64 `json:"armor"`
	Power       float64 `json:"power"` // energy at r0
	SprintSpeed float64 `json:"sprintSpeed"`
}

func main() {
	spawnedEntities = make(map[string]Configuration)

	fmt.Println("MathFrame - Warframe Event Calculator")
	fmt.Println("======================================")
	fmt.Println(" - new: create a new configuration")
	fmt.Println(" - spawn: load a saved config into simulation scope")
	fmt.Println(" - despawn: remove a config from simulation scope")
	fmt.Println(" - spawned: list currently spawned entities")
	fmt.Println(" - list: list available entity types")
	fmt.Println(" - saves: list saved configurations")
	fmt.Println(" - exit: quit")
	fmt.Println()

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Command: ")
		if !scanner.Scan() {
			break
		}

		command := strings.TrimSpace(strings.ToLower(scanner.Text()))

		switch command {
		case "new":
			createNewConfiguration(scanner)
		case "spawn":
			spawnConfiguration(scanner)
		case "despawn":
			despawnConfiguration(scanner)
		case "spawned":
			listSpawned()
		case "list":
			listEntities()
		case "saves":
			listSavedConfigurations()
		case "exit":
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Unknown command")
		}
	}
}

func loadEntities() map[string]Entity {
	// Return cached result if still valid
	if entitiesCache != nil && time.Since(cacheTime) < 24*time.Hour {
		return entitiesCache
	}

	// Try local cache first
	filename := "./wfdata/warframes.json"
	if data, err := os.ReadFile(filename); err == nil {
		var items []WarframeManifestItem
		if err := json.Unmarshal(data, &items); err == nil && len(items) > 0 {
			m := parseWarframes(items)
			if len(m) > 0 {
				entitiesCache = m
				cacheTime = time.Now()
				return m
			}
		}
	}

	// Fetch fresh data
	const url = "https://raw.githubusercontent.com/calamity-inc/warframe-public-export/senpai/ExportWarframes_en.json"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Fetch error: %v\n", err)
		return make(map[string]Entity)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Read error: %v\n", err)
		return make(map[string]Entity)
	}

	// Parse the actual structure: { "ExportWarframes": [ ... ] }
	type wrapper struct {
		ExportWarframes []WarframeManifestItem `json:"ExportWarframes"`
	}
	var w wrapper
	if err := json.Unmarshal(body, &w); err != nil {
		fmt.Printf("JSON unmarshal failed: %v\n", err)
		// Optional: print first ~200 bytes for debugging
		if len(body) > 200 {
			fmt.Printf("JSON starts with: %s...\n", string(body[:200]))
		}
		return make(map[string]Entity)
	}

	items := w.ExportWarframes
	if len(items) == 0 {
		fmt.Println("No warframes found in data")
		return make(map[string]Entity)
	}

	m := parseWarframes(items)

	// Cache the parsed items (array)
	if err := os.MkdirAll("./wfdata", 0755); err == nil {
		cacheData, _ := json.MarshalIndent(items, "", "  ")
		os.WriteFile(filename, cacheData, 0644)
	}

	entitiesCache = m
	cacheTime = time.Now()
	return m
}

func parseWarframes(items []WarframeManifestItem) map[string]Entity {
	m := make(map[string]Entity)
	for _, item := range items {
		if !strings.Contains(item.UniqueName, "/Powersuits/") {
			continue
		}

		// Take the last segment after final /
		last := item.UniqueName[strings.LastIndex(item.UniqueName, "/")+1:]
		key := strings.ToLower(last) // "ash", "ashprime", "garuda", "garudaprime", etc.

		// Skip Archwings / Necramechs if needed (filter by productCategory == "Suits" if field exists)
		// if item.ProductCategory != "Suits" { continue }

		e := Entity{
			Name:      key, // store the key as Name for now; or use item.Name later for display
			Faction:   "tenno",
			Health:    item.Health,
			Shields:   item.Shield,
			Armor:     item.Armor,
			Energy:    item.Power,
			Overguard: 0,
			Sprint:    item.SprintSpeed,
			Level:     0,
		}
		m[key] = e
	}
	return m
}

func listEntities() {
	entities := loadEntities()
	if len(entities) == 0 {
		fmt.Println("No entities loaded (check network/cache)")
		return
	}

	fmt.Println("\nAvailable entities (Tenno Warframes - base/rank 0 stats):")
	for name, entity := range entities {
		fmt.Printf("  %-20s [%s] H:%.0f A:%.0f S:%.0f\n",
			name, entity.Faction, entity.Health, entity.Armor, entity.Shields)
	}
	fmt.Println()
}

func createNewConfiguration(scanner *bufio.Scanner) {
	fmt.Print("Configuration name: ")
	if !scanner.Scan() {
		return
	}
	configName := strings.TrimSpace(scanner.Text())
	if configName == "" {
		fmt.Println("Configuration name cannot be empty.")
		return
	}

	fmt.Print("Entity name (type 'list' to see available): ")
	if !scanner.Scan() {
		return
	}
	entityInput := strings.TrimSpace(scanner.Text())

	if strings.ToLower(entityInput) == "list" {
		listEntities()
		fmt.Print("Entity name: ")
		if !scanner.Scan() {
			return
		}
		entityInput = strings.TrimSpace(scanner.Text())
	}

	// Normalize input
	entityInput = strings.ToLower(entityInput)

	entities := loadEntities()

	// Try direct match first
	baseEntity, exists := entities[entityInput]

	// If not found, try translation / alias map
	if !exists {
		if translated, ok := wfdata.WarframeNameMap[entityInput]; ok {
			baseEntity, exists = entities[translated]
		}
	}

	if !exists {
		fmt.Printf("Entity '%s' not found. Try 'list' to see available names.\n", entityInput)
		return
	}

	// Ask for level (required, no default)
	fmt.Print("Level: ")
	if !scanner.Scan() {
		return
	}
	levelStr := strings.TrimSpace(scanner.Text())
	level, err := strconv.Atoi(levelStr)
	if err != nil || level < 1 {
		fmt.Println("Invalid level — must be a positive integer.")
		return // ← crucial: abort here, do NOT save
	}

	// All inputs valid → proceed to create config
	currentStats := baseEntity
	currentStats.Level = level
	// Future: scaling logic can go here or in a separate step

	config := Configuration{
		ID:                    configName,
		Entity:                baseEntity,   // original base stats & name (level 0)
		Stats:                 currentStats, // copy with user-chosen level
		EntityConfig:          []interface{}{},
		AbilityConfig:         []interface{}{},
		PrimaryConfig:         []interface{}{},
		SecondaryConfig:       []interface{}{},
		MeleeConfig:           []interface{}{},
		ArchgunConfig:         []interface{}{},
		OperatorConfig:        []interface{}{},
		ArchwingConfig:        []interface{}{},
		ArchwinggunConfig:     []interface{}{},
		ArchwingmeleeConfig:   []interface{}{},
		CompanionConfig:       []interface{}{},
		CompanionweaponConfig: []interface{}{},
	}

	if err := saveConfiguration(configName, config); err != nil {
		fmt.Printf("Error saving configuration: %v\n", err)
		return
	}

	fmt.Printf("✓ Configuration '%s' created and saved.\n", configName)
	fmt.Printf("  Entity: %s (key: %s)\n", baseEntity.Name, entityInput)
	fmt.Printf("  Level: %d\n", level)
	fmt.Println("  (Scaling based on faction and level will be applied later.)")
}

func spawnConfiguration(scanner *bufio.Scanner) {
	fmt.Print("Configuration name to spawn (or 'list' to see available): ")
	if !scanner.Scan() {
		return
	}
	configName := strings.TrimSpace(scanner.Text())

	if configName == "list" {
		listSavedConfigurations()
		fmt.Print("Configuration name to spawn: ")
		if !scanner.Scan() {
			return
		}
		configName = strings.TrimSpace(scanner.Text())
	}

	config, err := loadConfigurationFromFile(configName)
	if err != nil {
		fmt.Printf("Error loading configuration: %v\n", err)
		return
	}

	baseName := config.Entity.Name
	spawnedName := findNextSpawnedName(baseName)

	spawnedEntities[spawnedName] = *config
	fmt.Printf("✓ Spawned '%s' as '%s' (Level %d)\n", configName, spawnedName, config.Stats.Level)
}

func findNextSpawnedName(baseName string) string {
	maxNum := 0
	for name := range spawnedEntities {
		if strings.HasPrefix(name, baseName) {
			numStr := strings.TrimPrefix(name, baseName)
			if num, err := strconv.Atoi(numStr); err == nil && num > maxNum {
				maxNum = num
			}
		}
	}
	return fmt.Sprintf("%s%d", baseName, maxNum+1)
}

func despawnConfiguration(scanner *bufio.Scanner) {
	if len(spawnedEntities) == 0 {
		fmt.Println("No entities currently spawned")
		return
	}

	fmt.Println("Currently spawned:")
	i := 1
	names := make([]string, 0, len(spawnedEntities))
	for name := range spawnedEntities {
		fmt.Printf("  [%d] %s\n", i, name)
		names = append(names, name)
		i++
	}

	fmt.Print("Choose entity to despawn (number or name): ")
	if !scanner.Scan() {
		return
	}
	input := strings.TrimSpace(scanner.Text())

	var choice int
	var configName string
	if _, err := fmt.Sscanf(input, "%d", &choice); err == nil {
		if choice < 1 || choice > len(names) {
			fmt.Println("Invalid choice")
			return
		}
		configName = names[choice-1]
	} else {
		configName = input
	}

	if _, exists := spawnedEntities[configName]; !exists {
		fmt.Printf("'%s' is not currently spawned\n", configName)
		return
	}

	delete(spawnedEntities, configName)
	fmt.Printf("✓ Despawned '%s'\n", configName)
}

func listSpawned() {
	if len(spawnedEntities) == 0 {
		fmt.Println("No entities currently spawned")
		return
	}

	fmt.Println("\nCurrently spawned entities:")
	for name, config := range spawnedEntities {
		fmt.Printf("  %-20s %s (Lv %d)  H:%.0f A:%.0f S:%.0f\n",
			name, config.Entity.Name, config.Stats.Level,
			config.Stats.Health, config.Stats.Armor, config.Stats.Shields)
	}
	fmt.Println()
}

func listSavedConfigurations() {
	entries, err := os.ReadDir("saves/entities")
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("No saved configurations found")
		} else {
			fmt.Printf("Error reading configurations: %v\n", err)
		}
		return
	}

	var jsonFiles []string
	for _, entry := range entries {
		if !entry.IsDir() && filepath.Ext(entry.Name()) == ".json" {
			name := entry.Name()[:len(entry.Name())-5]
			jsonFiles = append(jsonFiles, name)
		}
	}

	if len(jsonFiles) == 0 {
		fmt.Println("No saved configurations found")
		return
	}

	fmt.Println("\nSaved configurations:")
	for _, name := range jsonFiles {
		fmt.Printf("  - %s\n", name)
	}
	fmt.Println()
}

func loadConfigurationFromFile(name string) (*Configuration, error) {
	filename := filepath.Join("saves", "entities", name+".json")
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var config Configuration
	if err := json.Unmarshal(data, &config); err != nil {
		return nil, err
	}

	return &config, nil
}

func saveConfiguration(name string, config Configuration) error {
	if err := os.MkdirAll("saves/entities", 0755); err != nil {
		return err
	}

	data, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return err
	}

	filename := filepath.Join("saves", "entities", name+".json")
	return os.WriteFile(filename, data, 0644)
}
