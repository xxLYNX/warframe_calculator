# Mod System Documentation

## Overview

The mod system is **fully data-driven via CSV** with no individual Go files needed. Mods are defined in [wfdata/modifiers/mods.csv](wfdata/modifiers/mods.csv) and loaded dynamically via [mods_loader.go](mods_loader.go).

## Architecture

### ModData Struct
Represents a mod definition with:
- **Name**: Mod identifier (e.g., "Pressure Point")
- **AppliesTo**: Classification tags required (e.g., `["melee"]`)
- **EffectType**: Type of effect (e.g., `flat_damage`, `flat_crit_chance`, `flat_armor`)
- **EffectValue**: Bonus per rank (e.g., `0.12` = 12% per rank)
- **MaxRank**: Maximum rank level (typically 10)
- **BaseDrain**: Mod capacity cost
- **Description**: Human-readable effect description

### LoadedMod Struct
Represents an installed mod on an entity with:
- **Mod**: The ModData definition
- **Rank**: Current rank (0 to MaxRank)

## CSV Format

```csv
name,applies_to,effect_type,effect_value,max_rank,base_drain,description
Pressure Point,melee,flat_damage,0.12,10,11,+12% melee damage per rank
Primed Pressure Point,melee,flat_damage,0.165,10,14,+16.5% melee damage per rank
Serration,primary,flat_damage,0.15,10,14,+15% primary damage per rank
Vitality,warframe,flat_health,0.40,10,6,+40% health per rank
```

### Columns
| Column | Type | Example | Notes |
|--------|------|---------|-------|
| name | string | Pressure Point | Unique mod identifier |
| applies_to | csv | melee,fists | Classification tags (comma-separated) |
| effect_type | string | flat_damage | Type of effect being modified |
| effect_value | float | 0.12 | Per-rank bonus (as decimal) |
| max_rank | int | 10 | Maximum rank level |
| base_drain | int | 11 | Capacity drain cost |
| description | string | +12% melee damage per rank | Effect description |

## API

### Load Mods
```go
mods, err := LoadModsFromCSV("./wfdata/modifiers/mods.csv")
```

### Check Compatibility
```go
mod := mods["Pressure Point"]
isCompatible := mod.CanApplyTo([]string{"melee", "fists", "landslide"})
// true - has all required tags
```

### Calculate Effects
```go
effects := mod.CalculateEffects(10) // rank 10
// Returns: map[string]float64{"flat_damage": 1.20}
```

## Examples

### Current Mods (14 Total)

#### Melee Mods
- **Pressure Point**: +12% melee damage/rank (max 120%)
- **Primed Pressure Point**: +16.5% melee damage/rank (max 165%)
- **Organ Shatter**: +30% melee crit damage/rank (max 150%)

#### Primary Mods
- **Serration**: +15% primary damage/rank (max 150%)
- **Split Chamber**: +20% multishot/rank (max 100%)
- **Point Strike**: +30% crit chance/rank (max 150%)
- **Vital Sense**: +30% crit damage/rank (max 150%)

#### Warframe Mods
- **Vitality**: +40% health/rank (max 400%)
- **Steel Fiber**: +27.5% armor/rank (max 275%)
- **Redirection**: +44% shields/rank (max 440%)
- **Streamline**: +6% ability efficiency/rank (max 30%)
- **Intensify**: +5% ability strength/rank (max 50%)
- **Constitution**: +10% ability duration/rank (max 100%)
- **Stretch**: +15% ability range/rank (max 150%)

## Adding New Mods

1. Add a row to [wfdata/modifiers/mods.csv](wfdata/modifiers/mods.csv):
   ```csv
   Blind Rage,warframe,flat_strength,0.06,10,6,+6% ability strength per rank
   ```

2. Mods are automatically available via `LoadModsFromCSV()` - no code changes needed.

## Effect Types

Currently supported:
- `flat_damage` - Flat damage bonus (as decimal percentage)
- `flat_crit_chance` - Critical chance bonus
- `flat_crit_damage` - Critical damage multiplier bonus
- `flat_armor` - Armor bonus
- `flat_health` - Health bonus
- `flat_shields` - Shield bonus
- `flat_multishot` - Multishot bonus
- `flat_efficiency` - Ability efficiency bonus
- `flat_strength` - Ability strength bonus
- `flat_duration` - Ability duration bonus
- `flat_range` - Ability range bonus

## Applying Mods to Entities

To apply a mod to a spawned entity (future implementation):
```go
config := spawnedEntities["atlas1"]
mod := mods["Vitality"]

loadedMod := LoadedMod{
    Mod:  mod,
    Rank: 10,
}

config.Mods = append(config.Mods, loadedMod)
```

The configuration can then be serialized to JSON and saved for persistence.

## Testing

Run the demo:
```bash
go run mods_loader.go mods_demo.go
```

Output shows:
- 14 mods loaded
- Example mod effects at rank 10
- Compatibility tests (melee/primary weapons)
