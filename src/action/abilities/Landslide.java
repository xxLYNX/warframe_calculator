package action.abilities;

import action.Action;

public class Landslide extends Action {

    public Landslide() {
        this.classifications = Set.of(Classification.ABILITY, Classification.MELEE);
        this.duration = 1.0; // Duration in seconds

    }

}
