package defenseprofile;

/**
 * Abstract base class for all defense profile stat blocks in Warframe. Each
 * defense profile is a simple data container - no complex logic, just stats.
 * Modifiers handle all the damage calculations and interactions.
 */
public abstract class DefenseProfile {

    // Core defensive stats - final because they're fixed stat blocks
    public final double health;
    public final double armor;
    public final double shields;
    public final String name;

    /**
     * Constructor for defense profile stat blocks
     */
    protected DefenseProfile(String name, double health, double armor, double shields) {
        this.name = name;
        this.health = health;
        this.armor = armor;
        this.shields = shields;
    }

    /**
     * Simple string representation for debugging/display
     */
    @Override
    public String toString() {
        return String.format("%s - Health: %.0f, Armor: %.0f, Shields: %.0f",
                name, health, armor, shields);
    }
}
