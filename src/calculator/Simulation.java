package calculator;

import entity.Entity;
import java.util.*;

/**
 * Simulation - manages entities and action execution
 */
public class Simulation {

    private String name;
    private List<Entity> entities = new ArrayList<>();

    public Simulation(String name) {
        this.name = name;
    }

    public void addEntity(Entity entity) {
        entities.add(entity);
    }

    public void removeEntity(Entity entity) {
        entities.remove(entity);
    }

    public List<Entity> getEntities() {
        return entities;
    }

    public String getName() {
        return name;
    }

    /**
     * Display current simulation state
     */
    public void display() {
        System.out.println("\n=== Simulation: " + name + " ===");
        System.out.println("Entities: " + entities.size());
        for (Entity e : entities) {
            System.out.println("  - " + e.toString());
        }
    }

    /**
     * Save simulation (placeholder)
     */
    public void save(String filename) {
        System.out.println("Saving to " + filename + " (not yet implemented)");
    }

    /**
     * Load simulation (placeholder)
     */
    public static Simulation load(String filename) {
        System.out.println("Loading from " + filename + " (not yet implemented)");
        return null;
    }
}
