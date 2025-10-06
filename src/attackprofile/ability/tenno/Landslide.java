package attackprofile.ability.tenno;

import attackprofile.AttackProfile;
import attackprofile.Classification;

public class Landslide {

    public static class Primary extends AttackProfile {

        public Primary() {
            super("Landslide (Primary)",
                    350.0, 0, 0, // Impact: 350, Puncture: 0, Slash: 0
                    1.0, 0.35, 2.0, 0.1, // FireRate: 1.0, Crit: 35%, CritMult: 2.0x, Status: 10%
                    true, // isAOE: true
                    Classification.ABILITY, Classification.MELEE);
        }
    }

    public static class Secondary extends AttackProfile {

        public Secondary() {
            super("Landslide (Secondary)",
                    525.0, 0, 0, // Impact: 525 (1.5x multiplier)
                    1.0, 0.35, 2.0, 0.1, // Same stats as Primary
                    true, // isAOE: true
                    Classification.ABILITY, Classification.MELEE);
        }
    }

    public static class Tertiary extends AttackProfile {

        public Tertiary() {
            super("Landslide (Tertiary)",
                    700.0, 0, 0, // Impact: 700 (2.0x multiplier)
                    1.0, 0.35, 2.0, 0.1, // Same stats as Primary
                    true, // isAOE: true
                    Classification.ABILITY, Classification.MELEE);
        }
    }
}
