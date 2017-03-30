package timePredicates

import (
	"testing"
	"time"
)

var l Location = NewLocation(0, 0, 0, 0)

func Test_01(t *testing.T) {
	ev := NewEvent(time.Now(), l)

	t.Log("Event:", ev)
	t.Log("Event Time:", ev.GetEventTime())
	t.Log("Event Location:", ev.GetEventLocation())
}

func Test_02(t *testing.T) {
	ev1 := NewEvent(time.Now(), l)
	t.Log("Event ev1:", ev1, " - time:", ev1.GetEventTime())

	time.Sleep(time.Second)

	ev2 := NewEvent(time.Now(), l)
	t.Log("Event ev2:", ev2, " - time:", ev2.GetEventTime())

	t.Log("Is event ev2 after event ev1 ?:", ev2.After(ev1))
	t.Log("Is event ev1 after event ev2 ?:", ev1.After(ev2))
	t.Log("Is event ev1 after event ev1 ?:", ev1.After(ev1))

	t.Log("Is event ev2 before event ev1 ?:", ev2.Before(ev1))
	t.Log("Is event ev1 before event ev2 ?:", ev1.Before(ev2))
	t.Log("Is event ev1 before event ev1 ?:", ev1.Before(ev1))

	t.Log("Is event ev1 the same event ev1 (ev1 meets ev1) ?:", ev1.Meets(ev1))
	t.Log("Is event ev1 the same event ev2 (ev1 meets ev2) ?:", ev1.Meets(ev2))
}

func Test_03(t *testing.T) {
	ev1 := NewEvent(time.Now(), l)
	t.Log("Event ev1:", ev1, " - time:", ev1.GetEventTime())

	time.Sleep(time.Second)

	ev2 := NewEvent(time.Now(), l)
	t.Log("Event ev2:", ev2, " - time:", ev2.GetEventTime())

	time.Sleep(time.Second)

	ev3 := NewEvent(time.Now(), l)
	t.Log("Event ev3:", ev3, " - time:", ev3.GetEventTime())

	time.Sleep(time.Second)

	ev4 := NewEvent(time.Now(), l)
	t.Log("Event ev4:", ev4, " - time:", ev4.GetEventTime())

	ei1 := NewEventInterval(ev1, ev2)
	t.Log("Event Interval ei1:", ei1)

	ei2 := NewEventInterval(ev3, ev4)
	t.Log("Event Interval ei2:", ei2)

	ei3 := NewEventInterval(ev2, ev4)
	t.Log("Event Interval ei3:", ei3)

	ei4 := NewEventInterval(ev1, ev4)
	t.Log("Event Interval ei4:", ei4)

	ei5 := NewEventInterval(ev2, ev3)
	t.Log("Event Interval ei5:", ei5)

	ei6 := NewEventInterval(ev1, ev3)
	t.Log("Event Interval ei6:", ei6)

	// After()
	t.Log("Is event interval ei2 after event interval ei1 ?:", ei2.After(ei1))
	t.Log("Is event interval ei1 after event interval ei2 ?:", ei1.After(ei2))
	t.Log("Is event interval ei1 after event interval ei1 ?:", ei1.After(ei1))
	// Before()
	t.Log("Is event interval ei2 before event interval ei1 ?:", ei2.Before(ei1))
	t.Log("Is event interval ei1 before event interval ei2 ?:", ei1.Before(ei2))
	t.Log("Is event interval ei1 before event interval ei1 ?:", ei1.Before(ei1))
	// Meets()
	t.Log("Does event interval ei2 meet event interval ei1 ?:", ei2.Meets(ei1))
	t.Log("Does event interval ei1 meet event interval ei2 ?:", ei1.Meets(ei2))
	t.Log("Does event interval ei1 meet event interval ei1 ?:", ei1.Meets(ei1))
	t.Log("Does event interval ei1 meet event interval ei3 ?:", ei1.Meets(ei3))
	// Overlaps()
	t.Log("Does event interval ei2 overlap event interval ei1 ?:", ei2.Overlaps(ei1))
	t.Log("Does event interval ei1 overlap event interval ei2 ?:", ei1.Overlaps(ei2))
	t.Log("Does event interval ei1 overlap event interval ei1 ?:", ei1.Overlaps(ei1))
	t.Log("Does event interval ei6 overlap event interval ei3 ?:", ei6.Overlaps(ei3))
	// Starts()
	t.Log("Does event interval ei2 start event interval ei1 ?:", ei2.Starts(ei1))
	t.Log("Does event interval ei1 start event interval ei2 ?:", ei1.Starts(ei2))
	t.Log("Does event interval ei1 start event interval ei1 ?:", ei1.Starts(ei1))
	t.Log("Does event interval ei3 start event interval ei5 ?:", ei3.Starts(ei5))
	// Finishes()
	t.Log("Does event interval ei2 finish event interval ei1 ?:", ei2.Finishes(ei1))
	t.Log("Does event interval ei1 finish event interval ei2 ?:", ei1.Finishes(ei2))
	t.Log("Does event interval ei1 finish event interval ei1 ?:", ei1.Finishes(ei1))
	t.Log("Does event interval ei2 finish event interval ei4 ?:", ei2.Finishes(ei4))
	// During(): 'is-contained'
	t.Log("Is event interval ei2 during event interval ei1 ?:", ei2.During(ei1))
	t.Log("Is event interval ei1 during event interval ei2 ?:", ei1.During(ei2))
	t.Log("Is event interval ei1 during event interval ei1 ?:", ei1.During(ei1))
	t.Log("Is event interval ei5 during event interval ei4 ?:", ei5.During(ei4))
}
