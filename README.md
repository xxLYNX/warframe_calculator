# MathFrame - the next gen WarFrameCalculator
A comprehensive mathmatical understanding of all WarFrame damage sources. 
## Introduction
MathFrame works off a very simple premise - every possible source of damage be it a bullet from a Lex or a Rhino's Stomp ability are all subject to a specific subset of the possible buffs/debuffs in WarFrame. All of these sources of damage are what MathFrame refers to as __attack profiles__. Attack profiles must have a target to deal damage against, MathFrame refers to a given recipient of damage as the __defense profile__. Effects that modify this calculation are referred to as __modifiers__ by MathFrame. As such __modifiers__ will be the primary way you will be modifying your build. Common __modifiers__ are damage mods like _Serration_, _Arcanes_, _Armor_, and more. Modifiers fall under one of three categories:
### Static
Static modifiers are modifiers that are always active reguardless of attacker or defender behavior. These include most raw damage mods like _Serration_ or damage resistances like _Armor_. Some modifiers supplied by _Arcanes_ such as _Molt Augmented_ (+0.24% strength per kill up to 250x stacks) as based on cumulative actions up to a cap. For simplicity MathFrame assumes you have reached this cap and therefore treats this and similar effects such as the _Melee Combo Counter_ as being a static modifier. Notably a warning is given for melee sources with unchanged combo decay rates.
### Conditional
Conditional modifiers are modifiers that require a non deterministic (aka not guaranteed) pre-condition to occur that cannot be "forced" and we have no control over. For example _Arcane Avenger_ sometimes gives a crit chance buff upon the user taking damage BUT there is no way to say for sure how often this will occur. Both the probability of enemies hitting the player and the probability of the _Arcane_ triggering cannot be controlled. This is not the same as crit chance as we can easily calculate the average number of crits will occur.
### Active
Active modifiers are modifiers that we have conplete control over and they will never activate themselves without some player action being taken. For example _Ability Roar_ from Rhino provides a flat damage buff to everything while active. If its cast we benefit if not we never benefit. Splitting these catagorically helps us understand what we are reliant on while also finding the amount of _energy_ a set of modifiers may require per minute. 
## This Seems Complex?
Compared to more traditional builders like WarframeBuilder or Overframe MathFrame takes things in a very different direction. Historically these builders were simply an out of game recreation of the in-game mod selection screen and did very little to meaningfuly communicate to the user the sum total in-game effect of their combined modifiers. Questions like - "How much does this Slash Proc do?" and "What is my probability of one-shotting a lvl 100 Corrupted Bombard?" are all impossible to answer with these builders. MathFrame can answer these and so much more. Not only that but unlike these builders MathFrame helps identity issues in your configurations before you realize it yourself and inadvertently waste valuable resources (and your time!) in-game. And while some may be unused to programs with no graphical user interface I promise MathFrame is easy to use should you give it a chance!
## Impossible Builds
In an effort to facilitate changes should DE decide to rework things AGAIN everything is highly flexible in how its used. For now there are zero restrictions on if a given series of modifiers is possible in an actual in-game scinario. For example having both a primed fury and fury mod equiped is impossible in-game but can be done here. This is done to give you freedom to experiment "what-if" scinarios but should be kept in mind. I may add checks for these but overall functionality takes absolute priority.
# Getting Started



# Development
## Architecture
- **Attack Profiles**: Data containers for weapon/damage source specifications
- **Defense Profiles**: Data containers for target enemy/warframe specifications  
- **Modifiers**: Logic components that transform the damage calculation
- **Calculator**: Orchestrates AttackProfile + Modifiers + DefenseProfile → Results

## Planned Features
- Basic Attack Profile Framework (placeholder custom ability)
- Basic Defense Profile Framework
- Melee Modifiers (mods, arcanes)
- Primary Modifiers (mods, arcanes)
- Secondary Modifiers (mods, arcanes)
- Operator Modifiers (arcanes, abilities, schools)
- Warframe Modifiers (abilities, passives, shards)
## Bugs
See something that doesn't match up in-game? Well please submit a bug report that includes the following: all mods, shards, companions, warframe, weapons, and school equipped along with the mission in question. I'll do my best to isolate any issues as they crop up but be aware of teammates buffs etc. If possible please attempt to replicate the issue in the Simulacrum before reporting issues with Simulacrum proof are my top priority!
## Conclave PVP
No I will not add support for the conclave game mode not enough people play it to be worthwhile.