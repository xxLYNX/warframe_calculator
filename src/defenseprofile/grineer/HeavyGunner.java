package defenseprofile.grineer;

import defenseprofile.DefenseProfile;

/**
 * Heavy Gunner - High-level Grineer unit with heavy armor
 *
 * Typical stats for a level 30 Heavy Gunner: - High health for sustained combat
 * - Very high armor (damage reduction) - No shields (Grineer faction trait)
 */
public class HeavyGunner extends DefenseProfile {

    public HeavyGunner() {
        super("Heavy Gunner", 1800.0, 500.0, 0.0);
    }
}
