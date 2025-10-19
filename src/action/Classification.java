package action;

public enum Classification {
    //Action Equip Region
    PRIMARY,
    SECONDARY,
    MELEE,
    AMP,
    ARCHWING_PRIMARY,
    ARCHWING_MELEE,
    SENTINEL,
    BEAST,
    MOA,
    //Action Subtype
    RIFLE,
    SHOTGUN,
    SNIPER,
    BOW,
    KITGUN,
    ZAW,
    AUGMENT, //this is special where if an action has this classification a name check is performed to give it access to an augment mod corresponding to its name

    //Abilities
    ABILITY,
    EXALTED,
    PSEUDO_EXALTED,
    CHANNELED,
    //Type
    EVENT,
    MANUAL

}
