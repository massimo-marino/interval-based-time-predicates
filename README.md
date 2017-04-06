# Interval-Based Temporal Predicates

**Files:** *timePredicates.go, timePredicates_test.go*

A very simple implementation of some methods for an interval-based temporal logic based on [James F. Allen's work](http://web.mit.edu/larsb/Public/16.412/pset%204/allen94actions.pdf).  
This code can be used to implement scheduling/planning algorithms.

## Location

A location `l` is a point in space with a heading:

`l = {x, y, z, theta}`

#### Location methods

Constructors:  
```go
func NewLocation(x float64, y float64, z float64, theta float64) Location

func (l *Location) MakeNewLocation(x float64, y float64, z float64, theta float64)
```
### Event

An event `ev` is a timed location:

`ev = {t, l}`

#### Event Methods

Constructors:  
```go
func NewEvent(t time.Time, l Location) Event

func (ev *Event) MakeNewEvent(t time.Time, l Location)
```

Getters:  
```go
func (ev Event) GetEventTime() time.Time

func (ev Event) GetEventLocation() Location
```

Event Time Predicates:  

```go
// After: After reports whether the event ev1 is after ev2.
func (ev1 Event) After(ev2 Event) bool

// Before: Before reports whether the event ev1 is before ev2.
func (ev1 Event) Before(ev2 Event) bool

// Meets: Meets reports whether the event ev1 equals ev2.
func (ev1 Event) Meets(ev2 Event) bool
```

## Event Interval

An event interval `ei` is a pair of events that identify a time interval in space-time between two events `ev1`, `ev2`:

`ei = {ev1, ev2}`


#### Event Interval Methods

Constructors:  
```go
func NewEventInterval(ev1 Event, ev2 Event) EventInterval

func (ei *EventInterval) MakeNewEventInterval(ev1 Event, ev2 Event)
```

Getters:  
```go
func (ei EventInterval) GetStartEvent() Event

func (ei EventInterval) GetEndEvent() Event

func (ei EventInterval) GetEventIntervalStartTime() time.Time

func (ei EventInterval) GetEventIntervalEndTime() time.Time

func (ei EventInterval) GetEventIntervalStartLocation() Location

func (ei EventInterval) GetEventIntervalEndLocation() Location
```

Event Interval Time Predicates:  

```go
// After: After reports whether the event interval e11 is after ei2.
// [ev1----ei2---ev2]...[ev1----ei1----ev2]
func (ei1 EventInterval) After(ei2 EventInterval) bool

// Before: Before reports whether the event interval e11 is before ei2.
// [ev1----ei1---ev2]...[ev1----ei2----ev2]
func (ei1 EventInterval) Before(ei2 EventInterval) bool

// Meets: Meets reports whether the event interval ei1 meets ei2: this happens
// when ei1.ev2 meets ei2.ev1
// [ev1----ei1---ev2][ev1----ei2----ev2]
func (ei1 EventInterval) Meets(ei2 EventInterval) bool

// Overlaps: Overlaps reports whether the event interval ei1 overlaps ei2:
// 1.
// [ev1----ei1----ev2]
//           [ev1----ei2----ev2]
// 2.
// [ev1----ei1----ev2]
//                   [ev1----ei2----ev2]
// 3.
// [ev1----ei1----ev2]
//     [ev1--ei2--ev2]
// 4.
// [ev1----ei1----ev2]
// [ev1----ei2----ev2]
func (ei1 EventInterval) Overlaps(ei2 EventInterval) bool

// Starts: Starts reports whether the event interval ei1 starts at the same time event of ei2:
// [ev1----ei1----ev2]
// [ev1--------ei2--------ev2]
func (ei1 EventInterval) Starts(ei2 EventInterval) bool

// Finishes: Finishes reports whether the event interval ei1 ends at the same time event of ei2:
//         [ev1----ei1----ev2]
// [ev1--------ei2--------ev2]
func (ei1 EventInterval) Finishes(ei2 EventInterval) bool

// During: During reports whether the event interval ei1 starts and ends, or is contained, within i2:
//     [ev1----ei1----ev2]
// [ev1--------ei2--------ev2]
func (ei1 EventInterval) During(ei2 EventInterval) bool
```

## Examples

See the test file *timePredicates_test.go*
