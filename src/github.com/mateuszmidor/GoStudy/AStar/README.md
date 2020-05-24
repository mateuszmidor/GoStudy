# A* (A star)

This is an attempt to implement ```A* search algorithm```

See: <https://en.wikipedia.org/wiki/A*_search_algorithm>

## Subprojects

- wikiastar - implementation from wikipedia; graph node is specifically: x,y, name
- genericastar - node is just and ID; can handle any graph
- multipathastar - search all possible paths, not just the cheapest one

## Example graph - a flight map

```C
                          +-----+
                          | GDA |<------
                       /--+-----+\      \------------
                 /-----    /      \                  \-----
+-----+    /-----         /        \                    +-----+
| SZC | <--              v          \                   | BYK |
+-----+              +-----+         |                  +-----+
   \                 | BDG |         \                    -> ^
    \                +-----+          \               ---/  /
     |                  /              \          ---/      |
     \                /-                \     ---/         /
      \              /                   v  -/             |
       \           /-                 +-----+             /
        |         /                   | WAW |             |
        \        /                    +-----+<           /
         \     /-                        /    \          |
          v   <                         /      \        /
        +-----+                        |        \-      |
        | WRO |                        /          \    /
        +-----+--\                    /            \   |
                  ---\               /              \ /
                      ---\          v             +-----+
                          -->+-----+       ------>| RZE |
                             | KRK |------/       +-----+
                             +-----+
```
