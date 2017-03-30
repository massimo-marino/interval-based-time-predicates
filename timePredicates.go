package timePredicates

import (
	"fmt"
	"time"
)

// a location is a point in 3-D space with a heading, though not used in the
// interval-based temporal logic
type Location struct {
	// [x,y,z] coordinates
	x float64
	y float64
	z float64
	// heading: angle in radians/degrees
	theta float64
}

type Locationer interface {
	MakeNewLocation()
}

func (l Location) String() string {
	return fmt.Sprintf("<%.3f %.3f %.3f %.3f>", l.x, l.y, l.z, l.theta)
}

func NewLocation(x float64, y float64, z float64, theta float64) Location {
	return Location{x, y, z, theta}
}

func (l *Location) MakeNewLocation(x float64, y float64, z float64, theta float64) {
	*l = Location{x, y, z, theta}
	return
}

// Event
// An event is a timed location {time, Location}
type Event struct {
	t time.Time
	l Location
}

type Eventer interface {
	MakeNewEvent()
	GetEventTime()
	GetEventLocation()
	After()
	Before()
	Meets()
}

func (ev Event) String() string {
	return fmt.Sprintf("<%v %v>", ev.t, ev.l)
}

func NewEvent(t time.Time, l Location) Event {
	return Event{t, l}
}

func (ev *Event) MakeNewEvent(t time.Time, l Location) {
	*ev = Event{t, l}
	return
}

func (ev Event) GetEventTime() time.Time {
	return ev.t
}

func (ev Event) GetEventLocation() Location {
	return ev.l
}

// After: After reports whether the event ev1 is after ev2.
func (ev1 Event) After(ev2 Event) bool {
	return ev1.GetEventTime().After(ev2.GetEventTime())
}

// Before: Before reports whether the event ev1 is before ev2.
func (ev1 Event) Before(ev2 Event) bool {
	return ev1.GetEventTime().Before(ev2.GetEventTime())
}

// Meets: Meets reports whether the event ev1 equals ev2.
func (ev1 Event) Meets(ev2 Event) bool {
	return ev1.GetEventTime().Equal(ev2.GetEventTime())
}

// Event Interval
// An event interval is a 2-tuple that represents a temporal interval between
// two events: {ev1, ev2} = {{t1, l1}, {t2, l2}}
type EventInterval struct {
	ev1 Event
	ev2 Event
}

type EventIntervaller interface {
	MakeNewEventInterval()
	GetEventIntervalStartTime()
	GetEventIntervalEndTime()
	GetEventIntervalStartLocation()
	GetEventIntervalEndLocation()
	After()
	Before()
	Meets()
	Overlaps()
	Starts()
	Finishes()
	During()
}

func (ei EventInterval) String() string {
	return fmt.Sprintf("<%v %v>", ei.ev1, ei.ev2)
}

func NewEventInterval(ev1 Event, ev2 Event) EventInterval {
	return EventInterval{ev1, ev2}
}

func (ei *EventInterval) MakeNewEventInterval(ev1 Event, ev2 Event) {
	*ei = EventInterval{ev1, ev2}
	return
}

func (ei EventInterval) GetStartEvent() Event {
	return ei.ev1
}

func (ei EventInterval) GetEndEvent() Event {
	return ei.ev2
}

func (ei EventInterval) GetEventIntervalStartTime() time.Time {
	return ei.GetStartEvent().GetEventTime()
}

func (ei EventInterval) GetEventIntervalEndTime() time.Time {
	return ei.GetEndEvent().GetEventTime()
}

func (ei EventInterval) GetEventIntervalStartLocation() Location {
	return ei.GetStartEvent().GetEventLocation()
}

func (ei EventInterval) GetEventIntervalEndLocation() Location {
	return ei.GetEndEvent().GetEventLocation()
}

// After: After reports whether the event interval e11 is after ei2.
// [ev1----ei2---ev2]...[ev1----ei1----ev2]
func (ei1 EventInterval) After(ei2 EventInterval) bool {
	return ei1.GetEventIntervalStartTime().After(ei2.GetEventIntervalEndTime())
}

// Before: Before reports whether the event interval e11 is before ei2.
// [ev1----ei1---ev2]...[ev1----ei2----ev2]
func (ei1 EventInterval) Before(ei2 EventInterval) bool {
	return ei1.GetEventIntervalEndTime().Before(ei2.GetEventIntervalStartTime())
}

// Meets: Meets reports whether the event interval ei1 meets ei2: this happens
// when ei1.ev2 meets ei2.ev1
// [ev1----ei1---ev2][ev1----ei2----ev2]
func (ei1 EventInterval) Meets(ei2 EventInterval) bool {
	return ei1.GetEndEvent().Meets(ei2.GetStartEvent())
}

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
func (ei1 EventInterval) Overlaps(ei2 EventInterval) bool {
	return ((ei1.ev2.After(ei2.ev1) || ei1.ev2.Meets(ei2.ev1)) && (ei1.ev2.Before(ei2.ev2) || ei1.ev2.Meets(ei2.ev2)))
}

// Starts: Starts reports whether the event interval ei1 starts at the same time event of ei2:
// [ev1----ei1----ev2]
// [ev1--------ei2--------ev2]
func (ei1 EventInterval) Starts(ei2 EventInterval) bool {
	return ei1.GetStartEvent().Meets(ei2.GetStartEvent())
}

// Finishes: Finishes reports whether the event interval ei1 ends at the same time event of ei2:
//         [ev1----ei1----ev2]
// [ev1--------ei2--------ev2]
func (ei1 EventInterval) Finishes(ei2 EventInterval) bool {
	return ei1.GetEndEvent().Meets(ei2.GetEndEvent())
}

// During: During reports whether the event interval ei1 starts and ends, or is contained, within i2:
//     [ev1----ei1----ev2]
// [ev1--------ei2--------ev2]
func (ei1 EventInterval) During(ei2 EventInterval) bool {
	return ei1.GetStartEvent().After(ei2.GetStartEvent()) && ei1.GetEndEvent().Before(ei2.GetEndEvent())
}
