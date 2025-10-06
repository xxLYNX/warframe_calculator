package attackprofile;

import java.util.Set;

/**
 * Abstract base class for all attack profile stat blocks in Warframe. Each
 * attack profile is a simple data container - no complex logic, just stats.
 * Modifiers handle all the damage calculations and interactions.
 */
public abstract class AttackProfile {

    // Core damage types - final because they're fixed stat blocks
    public final double DT_IMPACT;
    public final double DT_PUNCTURE;
    public final double DT_SLASH;
    public final double DT_FIRE;
    public final double DT_FREEZE;
    public final double DT_ELECTRICITY;
    public final double DT_POISON;
    public final double DT_EXPLOSION;
    public final double DT_CORROSIVE;
    public final double DT_GAS;
    public final double DT_MAGNETIC;
    public final double DT_RADIATION;
    public final double DT_VIRAL;
    // Exotic damage types - less commonly used
    public final double DT_FINISHER;
    public final double DT_RADIANT;
    public final double DT_SENTIENT;
    public final double DT_CINEMATIC;
    public final double DT_SHIELD_DRAIN;
    public final double DT_TRUE;
    public final double DT_ENERGY_DRAIN;

    // Core attack stats
    public final double fireRate;
    public final double criticalChance;
    public final double criticalMultiplier;
    public final double statusChance;
    public final String name;
    public final String property;
    public final boolean isAOE;  // Area of Effect attacks

    // Multiple classifications support
    public final Set<Classification> classifications;

    /**
     * Full constructor for complex attack profiles
     */
    protected AttackProfile(String name,
            double dtImpact, double dtPuncture, double dtSlash,
            double dtFire, double dtFreeze, double dtElectricity, double dtPoison,
            double dtExplosion, double dtCorrosive, double dtGas,
            double dtMagnetic, double dtRadiation, double dtViral,
            double dtFinisher, double dtRadiant, double dtSentient,
            double dtCinematic, double dtShieldDrain, double dtTrue, double dtEnergyDrain,
            double fireRate, double criticalChance, double criticalMultiplier,
            double statusChance, String property, boolean isAOE,
            Classification... classifications) {
        this.name = name;
        this.DT_IMPACT = dtImpact;
        this.DT_PUNCTURE = dtPuncture;
        this.DT_SLASH = dtSlash;
        this.DT_FIRE = dtFire;
        this.DT_FREEZE = dtFreeze;
        this.DT_ELECTRICITY = dtElectricity;
        this.DT_POISON = dtPoison;
        this.DT_EXPLOSION = dtExplosion;
        this.DT_CORROSIVE = dtCorrosive;
        this.DT_GAS = dtGas;
        this.DT_MAGNETIC = dtMagnetic;
        this.DT_RADIATION = dtRadiation;
        this.DT_VIRAL = dtViral;
        this.DT_FINISHER = dtFinisher;
        this.DT_RADIANT = dtRadiant;
        this.DT_SENTIENT = dtSentient;
        this.DT_CINEMATIC = dtCinematic;
        this.DT_SHIELD_DRAIN = dtShieldDrain;
        this.DT_TRUE = dtTrue;
        this.DT_ENERGY_DRAIN = dtEnergyDrain;
        this.fireRate = fireRate;
        this.criticalChance = criticalChance;
        this.criticalMultiplier = criticalMultiplier;
        this.statusChance = statusChance;
        this.property = property;
        this.isAOE = isAOE;
        this.classifications = Set.of(classifications);  // Modern Java 21 syntax
    }

    /**
     * Physical damage constructor with smart defaults
     */
    protected AttackProfile(String name, double impact, double puncture, double slash,
            double fireRate, double criticalChance, double criticalMultiplier,
            double statusChance, boolean isAOE, Classification... classifications) {
        this(name, impact, puncture, slash,
                0, 0, 0, 0, // All elemental = 0
                0, 0, 0, 0, 0, 0, // All combined = 0
                0, 0, 0, 0, 0, 0, 0, // All exotic = 0
                fireRate, criticalChance, criticalMultiplier, statusChance, "", isAOE,
                classifications);
    }

    /**
     * Minimal constructor for simple cases with smart defaults
     */
    protected AttackProfile(String name, double primaryDamage, String damageType,
            boolean isAOE, Classification... classifications) {
        this(name,
                damageType.equals("Impact") ? primaryDamage : 0,
                damageType.equals("Puncture") ? primaryDamage : 0,
                damageType.equals("Slash") ? primaryDamage : 0,
                1.0, 0.0, 1.0, 0.0, isAOE, // Smart defaults: fireRate=1.0, no crits, no status
                classifications);
    }

    // Helper methods for classification checking
    public boolean hasClassification(Classification classification) {
        return classifications.contains(classification);
    }

    public boolean isAbility() {
        return hasClassification(Classification.ABILITY);
    }

    public boolean isMelee() {
        return hasClassification(Classification.MELEE);
    }

    public boolean isPrimary() {
        return hasClassification(Classification.PRIMARY);
    }

    public boolean isSecondary() {
        return hasClassification(Classification.SECONDARY);
    }

    /**
     * Simple string representation for debugging/display
     */
    @Override
    public String toString() {
        double totalDamage = DT_IMPACT + DT_PUNCTURE + DT_SLASH + DT_FIRE + DT_FREEZE
                + DT_ELECTRICITY + DT_POISON + DT_EXPLOSION + DT_CORROSIVE
                + DT_GAS + DT_MAGNETIC + DT_RADIATION + DT_VIRAL + DT_FINISHER
                + DT_RADIANT + DT_SENTIENT + DT_CINEMATIC + DT_SHIELD_DRAIN
                + DT_TRUE + DT_ENERGY_DRAIN;

        return String.format("%s (%s) - Total: %.1f, Fire Rate: %.1f, Crit: %.1f%% x%.1fx",
                name, classifications, totalDamage, fireRate,
                criticalChance * 100, criticalMultiplier);
    }
}
