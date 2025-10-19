package action;

public abstract class Action {

    protected double duration;
    protected Set<Classification> classifications;
    protected Map<String, Double> damageTypes;

    public abstract Result execute(Entity source, Entity target);

    public List<Modifier> generatesModifiers() {
        return Collections.emptyList(); // Override if action generates modifiers
    }
}
