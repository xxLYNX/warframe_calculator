package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Weapon represents a weapon with hierarchical classifications
type Weapon struct {
	Name            string             `json:"name"`
	Classifications []string           `json:"classifications"` // ["exalted", "melee", "fists", "landslide"]
	BaseDamage      map[string]float64 `json:"base_damage"`
	CritChance      float64            `json:"crit_chance"`
	CritMultiplier  float64            `json:"crit_multiplier"`
	StatusChance    float64            `json:"status_chance"`
	FireRate        float64            `json:"fire_rate,omitempty"`
	Magazine        int                `json:"magazine,omitempty"`
	Slot            string             `json:"slot"`   // Primary, Secondary, Melee
	Family          string             `json:"family"` // Braton, Tetra, etc.
}

// WeaponMod represents a mod with classification requirements
type WeaponMod struct {
	Name      string             `json:"name"`
	AppliesTo []string           `json:"applies_to"` // Tags required for mod to apply
	Effects   map[string]float64 `json:"effects"`
	MaxRank   int                `json:"max_rank"`
}

// CanEquipMod checks if this weapon can equip a given mod
func (w *Weapon) CanEquipMod(mod WeaponMod) bool {
	// Weapon must have ALL the tags the mod requires
	for _, requiredTag := range mod.AppliesTo {
		if !w.HasClassification(requiredTag) {
			return false
		}
	}
	return true
}

// HasClassification checks if weapon has a specific classification tag
func (w *Weapon) HasClassification(tag string) bool {
	for _, c := range w.Classifications {
		if strings.EqualFold(c, tag) {
			return true
		}
	}
	return false
}

// LoadWeaponsFromCSV loads weapons from the CSV file
func LoadWeaponsFromCSV(filepath string) (map[string]Weapon, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	weapons := make(map[string]Weapon)
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	// Find column indices
	header := records[0]
	nameIdx := findColumn(header, "Name")
	impactIdx := findColumn(header, "Impact")
	punctureIdx := findColumn(header, "Puncture")
	slashIdx := findColumn(header, "Slash")
	critChanceIdx := findColumn(header, "CritChance")
	critMultIdx := findColumn(header, "CritMultiplier")
	statusChanceIdx := findColumn(header, "StatusChance")
	fireRateIdx := findColumn(header, "FireRate")
	magazineIdx := findColumn(header, "Magazine")
	slotIdx := findColumn(header, "Slot")
	classIdx := findColumn(header, "Class")
	familyIdx := findColumn(header, "Family")

	for i := 1; i < len(records); i++ {
		record := records[i]
		if len(record) <= nameIdx {
			continue
		}

		name := record[nameIdx]
		
		// Build base damage map
		baseDamage := make(map[string]float64)
		if impactIdx >= 0 {
			if val, err := strconv.ParseFloat(record[impactIdx], 64); err == nil && val > 0 {
				baseDamage["impact"] = val
			}
		}
		if punctureIdx >= 0 {
			if val, err := strconv.ParseFloat(record[punctureIdx], 64); err == nil && val > 0 {
				baseDamage["puncture"] = val
			}
		}
		if slashIdx >= 0 {
			if val, err := strconv.ParseFloat(record[slashIdx], 64); err == nil && val > 0 {
				baseDamage["slash"] = val
			}
		}

		// Build classifications array from general to specific
		classifications := buildClassifications(
			getField(record, slotIdx),
			getField(record, classIdx),
			getField(record, familyIdx),
			name,
		)

		weapon := Weapon{
			Name:            name,
			Classifications: classifications,
			BaseDamage:      baseDamage,
			CritChance:      parseFloat(record, critChanceIdx),
			CritMultiplier:  parseFloat(record, critMultIdx),
			StatusChance:    parseFloat(record, statusChanceIdx),
			FireRate:        parseFloat(record, fireRateIdx),
			Magazine:        parseInt(record, magazineIdx),
			Slot:            getField(record, slotIdx),
			Family:          getField(record, familyIdx),
		}

		weapons[name] = weapon
	}

	return weapons, nil
}

// buildClassifications creates hierarchical classification tags
func buildClassifications(slot, class, family, name string) []string {
	var tags []string

	// Add slot (Primary, Secondary, Melee)
	if slot != "" {
		tags = append(tags, strings.ToLower(slot))
	}

	// Add class (Rifle, Shotgun, Bow, etc.)
	if class != "" {
		tags = append(tags, strings.ToLower(class))
	}

	// Add family (Braton, Tetra, etc.) - only if different from name
	if family != "" && !strings.EqualFold(family, name) {
		tags = append(tags, strings.ToLower(family))
	}

	// Add specific weapon name
	tags = append(tags, strings.ToLower(name))

	return tags
}

// Helper functions
func findColumn(header []string, name string) int {
	for i, col := range header {
		if strings.EqualFold(col, name) {
			return i
		}
	}
	return -1
}

func getField(record []string, idx int) string {
	if idx < 0 || idx >= len(record) {
		return ""
	}
	return record[idx]
}

func parseFloat(record []string, idx int) float64 {
	if idx < 0 || idx >= len(record) {
		return 0
	}
	val, _ := strconv.ParseFloat(record[idx], 64)
	return val
}

func parseInt(record []string, idx int) int {
	if idx < 0 || idx >= len(record) {
		return 0
	}
	val, _ := strconv.Atoi(record[idx])
	return val
}

// Example usage for testing
//func main() {
	weapons, err := LoadWeaponsFromCSV("./wfdata/weapons.csv")
	if err != nil {
		fmt.Printf("Error loading weapons: %v\n", err)
		return
	}

	fmt.Printf("Loaded %d weapons\n\n", len(weapons))

	// Show some examples
	examples := []string{"Braton", "Kuva Bramma", "Zhuge"}
	for _, name := range examples {
		if weapon, exists := weapons[name]; exists {
			fmt.Printf("%s:\n", weapon.Name)
			fmt.Printf("  Classifications: %v\n", weapon.Classifications)
			fmt.Printf("  Slot: %s, Family: %s\n", weapon.Slot, weapon.Family)
			fmt.Printf("  Damage: %v\n", weapon.BaseDamage)
			fmt.Printf("  Crit: %.2f%% x%.1f\n", weapon.CritChance*100, weapon.CritMultiplier)
			fmt.Println()
		}
	}

	// Test mod matching example
	fmt.Println("Example Mod Matching:")
	testWeapon := weapons["Braton"]
	
	serration := WeaponMod{Name: "Serration", AppliesTo: []string{"primary"}}
	splitChamber := WeaponMod{Name: "Split Chamber", AppliesTo: []string{"primary"}}
	bratonAugment := WeaponMod{Name: "Braton Augment", AppliesTo: []string{"braton"}}
	pressurePoint := WeaponMod{Name: "Pressure Point", AppliesTo: []string{"melee"}}
	
	fmt.Printf("Braton can equip Serration (primary): %v\n", testWeapon.CanEquipMod(serration))
	fmt.Printf("Braton can equip Split Chamber (primary): %v\n", testWeapon.CanEquipMod(splitChamber))
	fmt.Printf("Braton can equip Braton Augment (braton): %v\n", testWeapon.CanEquipMod(bratonAugment))
	fmt.Printf("Braton can equip Pressure Point (melee): %v\n", testWeapon.CanEquipMod(pressurePoint))
}
