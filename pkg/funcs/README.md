# Functions

## DefaultFunction

Function name: `default`

if function-name not match any func, use it as default function.

### Params

no param

### Expression

`y = 0`

---

## LinearFunction

Function name: `StandardLinearFunction`

### Params

- slope
- offsetX
- offsetY

### Expression

Input: x (`int64`)

`y = slope * (x + offsetX) + offsetY`

### Inner set:

`DefaultLinearFunction`

- slope: 1
- offsetX: 0
- offsetY: 0

`ReverseLinearFunction`

- slope: -1
- offsetX: 0
- offsetY: 0

---

## RandomFunction

Function name `StandardRandomFunction`

### Params:

- range
- seed
- base-point

### Expression

rand generate by seed

y = base-point + rand()

base-point <= y <= base-point + range(ceil)