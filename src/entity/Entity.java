package entity;

public abstract class Entity {

    protected String name;
    protected String faction;
    protected long health;
    protected long armor;
    protected long shields;
    protected List<Modifier> modifiers = new LinkedList<>();

    public Result update(Action action, Entity source) {
        //applying action taking into account prexisting modifiers
    }
}
