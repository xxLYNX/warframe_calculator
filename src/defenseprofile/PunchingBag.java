package defenseprofile;

/**
 * PunchingBag - A neutral testing target with infinite health.
 *
 * This defense profile is used as a baseline for damage testing: - Infinite
 * health so you can see pure DPS numbers - Moderate armor value for testing
 * armor calculations - No shields to keep it simple - No faction-specific
 * modifiers (handled by modifier system)
 */
public class PunchingBag extends DefenseProfile {

    /**
     * Create a PunchingBag with infinite health for pure damage testing
     */
    public PunchingBag() {
        super("Punching Bag", Double.POSITIVE_INFINITY, 0, 0.0);
    }
}
