package calculator;

import action.Action;
import action.ActionResult;
import action.abilities.Landslide;
import entity.Entity;
import entity.tenno.Atlas;
import entity.grineer.HeavyGunner;

import java.util.*;

public class main {

    public static void main(String[] args) {
        System.out.println("=== Warframe Damage Calculator ===\n");
        Scanner scanner = new Scanner(System.in);
        Simulation currentSim = null;

        while (true) {
            System.out.println("\n=== Main Menu ===");
            System.out.println("1. Create New Simulation");
            System.out.println("2. Load Simulation");
            System.out.println("3. Add Entity");
            System.out.println("4. Perform Action");
            System.out.println("5. Display Simulation");
            System.out.println("6. Advance Time");
            System.out.println("7. Save Simulation");
            System.out.println("8. Exit");
            System.out.print("\nChoice: ");

            int choice = scanner.nextInt();
            scanner.nextLine(); // Consume newline

            switch (choice) {
                case 1:
                    currentSim = createNewSimulation(scanner);
                    System.out.println("✓ New simulation created!");
                    break;

                case 2:
                    System.out.print("Enter filename: ");
                    String loadFile = scanner.nextLine();
                    currentSim = Simulation.load(loadFile);
                    if (currentSim != null) {
                        System.out.println("✓ Simulation loaded!");
                    } else {
                        System.out.println("✗ Failed to load simulation.");
                    }
                    break;

                case 3:
                    if (currentSim == null) {
                        System.out.println("✗ Create or load a simulation first!");
                    } else {
                        addEntity(scanner, currentSim);
                    }
                    break;

                case 4:
                    if (currentSim == null) {
                        System.out.println("✗ Create or load a simulation first!");
                    } else {
                        performAction(scanner, currentSim);
                    }
                    break;

                case 5:
                    if (currentSim == null) {
                        System.out.println("✗ No simulation to display!");
                    } else {
                        currentSim.display();
                    }
                    break;

                case 6:
                    if (currentSim == null) {
                        System.out.println("✗ Create or load a simulation first!");
                    } else {
                        System.out.print("Time to advance (seconds): ");
                        double deltaTime = scanner.nextDouble();
                        currentSim.tick(deltaTime);
                        System.out.println("✓ Advanced " + deltaTime + " seconds");
                    }
                    break;

                case 7:
                    if (currentSim == null) {
                        System.out.println("✗ No simulation to save!");
                    } else {
                        System.out.print("Enter filename: ");
                        String saveFile = scanner.nextLine();
                        currentSim.save(saveFile);
                        System.out.println("✓ Simulation saved!");
                    }
                    break;

                case 8:
                    System.out.println("Goodbye!");
                    scanner.close();
                    return;

                default:
                    System.out.println("✗ Invalid choice!");
            }
        }
    }

    private static Simulation createNewSimulation(Scanner scanner) {
        System.out.print("Simulation name: ");
        String name = scanner.nextLine();
        return new Simulation(name);
    }

    private static void addEntity(Scanner scanner, Simulation sim) {
        System.out.println("\n=== Add Entity ===");
        System.out.println("1. Atlas (Tenno)");
        System.out.println("2. Heavy Gunner (Grineer)");
        System.out.print("Choice: ");

        int choice = scanner.nextInt();
        scanner.nextLine();

        Entity entity = null;
        switch (choice) {
            case 1:
                entity = new Atlas();
                break;
            case 2:
                entity = new HeavyGunner();
                break;
            default:
                System.out.println("✗ Invalid choice!");
                return;
        }

        sim.addEntity(entity);
        System.out.println("✓ Added " + entity.getName());
    }

    private static void performAction(Scanner scanner, Simulation sim) {
        if (sim.getEntities().isEmpty()) {
            System.out.println("✗ No entities in simulation!");
            return;
        }

        // Select source entity
        System.out.println("\n=== Select Source Entity ===");
        List<Entity> entities = sim.getEntities();
        for (int i = 0; i < entities.size(); i++) {
            Entity e = entities.get(i);
            System.out.println((i + 1) + ". " + e.getName()
                    + " (HP: " + e.getState("health")
                    + ", Energy: " + e.getState("energy") + ")");
        }
        System.out.print("Choice: ");
        int sourceIdx = scanner.nextInt() - 1;
        scanner.nextLine();

        if (sourceIdx < 0 || sourceIdx >= entities.size()) {
            System.out.println("✗ Invalid choice!");
            return;
        }
        Entity source = entities.get(sourceIdx);

        // Select action
        System.out.println("\n=== Select Action ===");
        System.out.println("1. Landslide (Primary)");
        System.out.println("2. Landslide (Secondary)");
        System.out.println("3. Landslide (Tertiary)");
        System.out.print("Choice: ");
        int actionChoice = scanner.nextInt();
        scanner.nextLine();

        Action action = null;
        switch (actionChoice) {
            case 1:
                action = new Landslide();
                break;
            case 2:
                action = new Landslide.Secondary();
                break;
            case 3:
                action = new Landslide.Tertiary();
                break;
            default:
                System.out.println("✗ Invalid choice!");
                return;
        }

        // Select target entities
        System.out.println("\n=== Select Targets (comma-separated, or 'all') ===");
        for (int i = 0; i < entities.size(); i++) {
            Entity e = entities.get(i);
            System.out.println((i + 1) + ". " + e.getName()
                    + " (HP: " + e.getState("health") + ")");
        }
        System.out.print("Targets: ");
        String targetInput = scanner.nextLine();

        List<Entity> targets = new ArrayList<>();
        if (targetInput.equalsIgnoreCase("all")) {
            targets.addAll(entities);
            targets.remove(source); // Don't target self
        } else {
            String[] indices = targetInput.split(",");
            for (String idx : indices) {
                try {
                    int i = Integer.parseInt(idx.trim()) - 1;
                    if (i >= 0 && i < entities.size()) {
                        targets.add(entities.get(i));
                    }
                } catch (NumberFormatException e) {
                    System.out.println("✗ Invalid target: " + idx);
                }
            }
        }

        if (targets.isEmpty()) {
            System.out.println("✗ No valid targets selected!");
            return;
        }

        // Perform action
        ActionResult result = sim.performAction(source, action, targets);

        // Display result
        System.out.println("\n=== Action Result ===");
        System.out.println(source.getName() + " used " + action.getName());
        System.out.println("\nState Changes:");
        for (Map.Entry<Entity, Map<String, Double>> entry : result.getEntityStateChanges().entrySet()) {
            Entity e = entry.getKey();
            System.out.println("  " + e.getName() + ":");
            for (Map.Entry<String, Double> stateChange : entry.getValue().entrySet()) {
                double delta = stateChange.getValue();
                System.out.printf("    %s: %+.1f\n", stateChange.getKey(), delta);
            }
        }

        if (!result.getTriggeredEvents().isEmpty()) {
            System.out.println("\nTriggered Events:");
            for (ActionResult.GameEvent event : result.getTriggeredEvents()) {
                System.out.println("  - " + event.type);
            }
        }
    }
}
