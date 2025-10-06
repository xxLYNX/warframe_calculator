package modifier.mod;

public class SerrationMod implements Modifier {

    private double damageMultiplier = 1.65; // +165% damage

    public DamageResult apply(DamageResult damage, Defender target) {
        // Additive with other damage mods, then multiplicative
        return damage.multiplyBaseDamage(damageMultiplier);
    }

}
