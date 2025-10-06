package calculator;

import attackprofile.AttackProfile;
import attackprofile.Classification;
import attackprofile.ability.tenno.Landslide;
import defenseprofile.DefenseProfile;
import defenseprofile.PunchingBag;

public class main {

    public static void main(String[] args) {
        System.out.println("=== Warframe Attack & Defense Profile Test ===\n");

        // Test Defense Profiles
        System.out.println("=== Defense Profiles ===");
        DefenseProfile target = new PunchingBag();
        System.out.println("Target: " + target.toString());
        System.out.println();

        // Test Attack Profiles
        System.out.println("=== Attack Profiles ===");
        AttackProfile landslide1 = new Landslide.Primary();
        AttackProfile landslide2 = new Landslide.Secondary();
        AttackProfile landslide3 = new Landslide.Tertiary();

        System.out.println("Attack 1: " + landslide1.toString());
        System.out.println("Attack 2: " + landslide2.toString());
        System.out.println("Attack 3: " + landslide3.toString());
        System.out.println();

        // Test Classification System
        System.out.println("=== Classification Tests ===");
        System.out.println("Landslide Primary classifications: " + landslide1.classifications);
        System.out.println("Is Ability? " + landslide1.isAbility());
        System.out.println("Is Melee? " + landslide1.isMelee());
        System.out.println("Is Primary Weapon? " + landslide1.isPrimary());
        System.out.println("Has ABILITY classification? " + landslide1.hasClassification(Classification.ABILITY));
        System.out.println("Has MELEE classification? " + landslide1.hasClassification(Classification.MELEE));
        System.out.println();

        // Test Damage Type Breakdown
        System.out.println("=== Damage Type Breakdown ===");
        System.out.println("Landslide Primary damage types:");
        System.out.println("  Impact: " + landslide1.DT_IMPACT);
        System.out.println("  Puncture: " + landslide1.DT_PUNCTURE);
        System.out.println("  Slash: " + landslide1.DT_SLASH);
        System.out.println("  Fire: " + landslide1.DT_FIRE);
        System.out.println();

        // Future: Calculator test
        System.out.println("=== Future Calculator Integration ===");
        System.out.println("Ready for: DamageResult result = calculator.calculate(attack, modifiers, target);");
        System.out.println("Attack: " + landslide1.name);
        System.out.println("Target: " + target.name);
        System.out.println("Modifiers: [None yet - to be implemented]");
    }
}
