package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ModData represents a mod definition
type ModData struct {
	Name        string   `json:"name"`
	AppliesTo   []string `json:"applies_to"`   // Classification tags required
	EffectType  string   `json:"effect_type"`  // flat_damage, flat_crit_chance, etc.
	EffectValue float64  `json:"effect_value"` // Per rank value
	MaxRank     int      `json:"max_rank"`
	BaseDrain   int      `json:"base_drain"`
	Description string   `json:"description"`
}

// LoadedMod represents an installed mod with a rank
type LoadedMod struct {
	Mod  ModData `json:"mod"`
	Rank int     `json:"rank"` // Current rank (0 to MaxRank)
}

// CalculateEffects calculates the total effect at a given rank
func (m *ModData) CalculateEffects(rank int) map[string]float64 {
	if rank > m.MaxRank {
		rank = m.MaxRank
	}
	if rank < 0 {
		rank = 0
	}

	totalEffect := m.EffectValue * float64(rank)
	return map[string]float64{
		m.EffectType: totalEffect,
	}
}

// AppliesTo checks if this mod can apply to a weapon/entity with given classifications
func (m *ModData) CanApplyTo(classifications []string) bool {
	// Mod applies if target has ALL required tags
	for _, required := range m.AppliesTo {
		if !containsTag(classifications, required) {
			return false
		}
	}
	return true
}

// LoadModsFromCSV loads all mods from CSV
func LoadModsFromCSV(filepath string) (map[string]ModData, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	mods := make(map[string]ModData)
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	// Find column indices
	header := records[0]
	nameIdx := findCSVColumn(header, "name")
	appliesToIdx := findCSVColumn(header, "applies_to")
	effectTypeIdx := findCSVColumn(header, "effect_type")
	effectValueIdx := findCSVColumn(header, "effect_value")
	maxRankIdx := findCSVColumn(header, "max_rank")
	baseDrainIdx := findCSVColumn(header, "base_drain")
	descIdx := findCSVColumn(header, "description")

	for i := 1; i < len(records); i++ {
		record := records[i]
		if len(record) <= nameIdx {
			continue
		}

		name := record[nameIdx]
		appliesToStr := getCSVField(record, appliesToIdx)
		appliesTo := strings.Split(appliesToStr, ",")

		// Trim whitespace from each tag
		for i := range appliesTo {
			appliesTo[i] = strings.TrimSpace(appliesTo[i])
		}

		effectValue, _ := strconv.ParseFloat(getCSVField(record, effectValueIdx), 64)
		maxRank, _ := strconv.Atoi(getCSVField(record, maxRankIdx))
		baseDrain, _ := strconv.Atoi(getCSVField(record, baseDrainIdx))

		mod := ModData{
			Name:        name,
			AppliesTo:   appliesTo,
			EffectType:  getCSVField(record, effectTypeIdx),
			EffectValue: effectValue,
			MaxRank:     maxRank,
			BaseDrain:   baseDrain,
			Description: getCSVField(record, descIdx),
		}

		mods[name] = mod
	}

	return mods, nil
}

// Helper functions
func findCSVColumn(header []string, name string) int {
	for i, col := range header {
		if strings.EqualFold(col, name) {
			return i
		}
	}
	return -1
}

func getCSVField(record []string, idx int) string {
	if idx < 0 || idx >= len(record) {
		return ""
	}
	return strings.TrimSpace(record[idx])
}

func containsTag(tags []string, tag string) bool {
	for _, t := range tags {
		if strings.EqualFold(t, tag) {
			return true
		}
	}
	return false
}

// Example usage for testing
func testMods() {
	mods, err := LoadModsFromCSV("./wfdata/modifiers/mods.csv")
	if err != nil {
		fmt.Printf("Error loading mods: %v\n", err)
		return
	}

	fmt.Printf("Loaded %d mods\n\n", len(mods))

	// Show examples
	examples := []string{"Pressure Point", "Primed Pressure Point", "Serration", "Vitality"}
	for _, name := range examples {
		if mod, exists := mods[name]; exists {
			fmt.Printf("%s:\n", mod.Name)
			fmt.Printf("  Applies to: %v\n", mod.AppliesTo)
			fmt.Printf("  Effect: %s (+%.2f%% per rank)\n", mod.EffectType, mod.EffectValue*100)
			fmt.Printf("  Max Rank: %d (drain: %d)\n", mod.MaxRank, mod.BaseDrain)

			// Show effects at rank 10
			effects := mod.CalculateEffects(10)
			for effType, value := range effects {
				fmt.Printf("  At Rank 10: +%.2f%% %s\n", value*100, effType)
			}
			fmt.Println()
		}
	}

	// Test mod compatibility
	fmt.Println("Mod Compatibility Examples:")
	pressurePoint := mods["Pressure Point"]
	primedPressure := mods["Primed Pressure Point"]
	serration := mods["Serration"]

	meleeClassifications := []string{"melee", "fists", "landslide"}
	primaryClassifications := []string{"primary", "rifle", "braton"}

	fmt.Printf("Pressure Point (melee) applies to melee weapon: %v\n", pressurePoint.CanApplyTo(meleeClassifications))
	fmt.Printf("Pressure Point (melee) applies to primary weapon: %v\n", pressurePoint.CanApplyTo(primaryClassifications))
	fmt.Printf("Serration (primary) applies to primary weapon: %v\n", serration.CanApplyTo(primaryClassifications))
	fmt.Printf("Serration (primary) applies to melee weapon: %v\n", serration.CanApplyTo(meleeClassifications))
	fmt.Printf("Primed Pressure Point (melee) applies to melee weapon: %v\n", primedPressure.CanApplyTo(meleeClassifications))
}
