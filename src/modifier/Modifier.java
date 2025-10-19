package modifier;

public interface Modifier {

    void applyToEntity(Entity entity);

    void applyToAction(Action action);

    boolean isExpired();

    String getDescription();
}
