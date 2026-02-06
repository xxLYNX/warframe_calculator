# Weapon Classifications - Hierarchical Tag System

## How It Works

Weapons have **ordered classifications** from general → specific:

```
["exalted", "melee", "fists", "landslide"]
 ↑ general                    specific ↑
```

## Example: Landslide (Atlas ability)

```go
Landslide: {
    Classifications: ["exalted", "melee", "fists", "landslide"]
}
```

### Mod Matching Examples

```go
// Broad mod - affects ALL melee
PressurePoint := WeaponMod{
    Name: "Pressure Point",
    AppliesTo: []string{"melee"}  // Landslide HAS "melee" → ✓ applies
}

// Narrower - affects ALL exalted melee
PowerStrength := WeaponMod{
    Name: "Intensify",
    AppliesTo: []string{"exalted", "melee"}  // Landslide HAS both → ✓ applies
}

// Very specific - affects ONLY fists
GeminiCross := WeaponMod{
    Name: "Gemini Cross",
    AppliesTo: []string{"fists"}  // Landslide HAS "fists" → ✓ applies
}

// Ultra specific - affects ONLY landslide family
LandslideAugment := WeaponMod{
    Name: "Path of Statues",
    AppliesTo: []string{"landslide"}  // Landslide HAS "landslide" → ✓ applies
}

// Won't apply - wrong weapon type
Serration := WeaponMod{
    Name: "Serration",
    AppliesTo: []string{"primary"}  // Landslide does NOT have "primary" → ✗ fails
}
```

## Braton Family Example

```go
// All three Bratons share family tag
Braton:       ["primary", "rifle", "braton", "braton"]
BratonPrime:  ["primary", "rifle", "braton", "braton prime"]
BratonVandal: ["primary", "rifle", "braton", "braton vandal"]

// Family-wide augment
BratonAugment := WeaponMod{
    AppliesTo: []string{"braton"}  // All 3 Bratons have "braton" tag → ✓ applies to all
}

// Variant-specific
PrimeAugment := WeaponMod{
    AppliesTo: []string{"braton prime"}  // Only Braton Prime → ✓ specific variant
}
```

## CSV Structure (Extended)

For manual data entry, you can extend weapons.csv with a classification column:

```csv
Name,Classifications,Damage,...
Landslide,"exalted,melee,fists,landslide",800,...
Braton,"primary,rifle,braton",24,...
Braton Prime,"primary,rifle,braton,braton prime",35,...
```

Or let it auto-generate from existing columns (Slot, Class, Family, Name) as the current implementation does.

## Benefits

1. **Flexible hierarchy**: Add as many levels as needed
2. **Simple matching**: Weapon must have ALL mod's required tags
3. **No duplicates**: Each weapon/mod combo checked once
4. **Inheritance**: Broad mods automatically work on specific weapons
5. **Go-idiomatic**: Just string slices and simple loops
