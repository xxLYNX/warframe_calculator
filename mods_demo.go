//package main

import (
	"fmt"
)

func main() {
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
