# Architectural drivers for GearBoxDriver

## Functional

- when in M-Dynamic mode and drifting, no gear change must occur
- when towing a trailor and driving down the slope, slowing down must be aided by engine
- when in comfort mode, pressing gas firmly results in kick down (gear down)
- when in sport mode, pressing gas results in kick down 1 or 2, depending on threshold
- when gear was changed manually, the driver doesnt change gear for some period of time
- when in comfort mode, optimal RPM range is affectedd by aggressiveness RPM multiplier
- when in sport mode, it's like comfort mode + soundmoule pipe blast is heard on gear down

## Quality

- adhere to company patterns, eg. "metoda kopiego pasty" ;)

## Project goals

- this is a quick demo for client

## Project restrictions

- must be ready in next 30min
- gearbox driver means embedded. JVM may not fit here; lets go with GO

## Conventions

- everything goes to a single source file
- avoid unit tests; they are waste of time
- least amount of classes...
