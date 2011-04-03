package inventory

import (
    "testing"

    .   "chunkymonkey/types"
)

func slotEq(s1, s2 *Slot) bool {
    return (s1.ItemType == s2.ItemType &&
        s1.Quantity == s2.Quantity &&
        s1.Uses == s2.Uses)
}

func TestSlot_Add(t *testing.T) {
    type Test struct {
        desc                 string
        initialA, initialB   Slot
        expectedA, expectedB Slot
    }

    tests := []Test{
        {
            "one empty slot added to another",
            Slot{ItemIDNull, 0, 0}, Slot{ItemIDNull, 0, 0},
            Slot{ItemIDNull, 0, 0}, Slot{ItemIDNull, 0, 0},
        },
        // Tests involving the same item types: (or empty plus an item)
        {
            "1 + 0 => 1 + 0",
            Slot{1, 1, 0}, Slot{ItemIDNull, 0, 0},
            Slot{1, 1, 0}, Slot{ItemIDNull, 0, 0},
        },
        {
            "1 + 1 => 2 + 0",
            Slot{1, 1, 0}, Slot{1, 1, 0},
            Slot{1, 2, 0}, Slot{ItemIDNull, 0, 0},
        },
        {
            "0 + 20 => 20 + 0",
            Slot{ItemIDNull, 0, 0}, Slot{1, 20, 0},
            Slot{1, 20, 0}, Slot{ItemIDNull, 0, 0},
        },
        {
            "0 + 64 => 64 + 0",
            Slot{ItemIDNull, 0, 0}, Slot{1, 64, 0},
            Slot{1, 64, 0}, Slot{ItemIDNull, 0, 0},
        },
        {
            "32 + 33 => 64 + 1 (hitting max quantity)",
            Slot{1, 32, 0}, Slot{1, 33, 0},
            Slot{1, 64, 0}, Slot{1, 1, 0},
        },
        {
            "65 + 1 => 65 + 1 (already above max quantity)",
            Slot{1, 65, 0}, Slot{1, 1, 0},
            Slot{1, 65, 0}, Slot{1, 1, 0},
        },
        {
            "64 + 64 => 64 + 64",
            Slot{1, 64, 0}, Slot{1, 64, 0},
            Slot{1, 64, 0}, Slot{1, 64, 0},
        },
        {
            "1 + 1 => 1 + 1 where items' \"Uses\" value differs",
            Slot{1, 1, 5}, Slot{1, 1, 6},
            Slot{1, 1, 5}, Slot{1, 1, 6},
        },
        {
            "1 + 1 => 2 + 0 where items' \"Uses\" value is the same",
            Slot{1, 1, 5}, Slot{1, 1, 5},
            Slot{1, 2, 5}, Slot{ItemIDNull, 0, 0},
        },
        {
            "0 + 1 => 1 + 0 - carrying the \"use\" value",
            Slot{ItemIDNull, 0, 0}, Slot{1, 1, 5},
            Slot{1, 1, 5}, Slot{ItemIDNull, 0, 0},
        },
        // Tests involving different item types:
        {
            "different item types don't mingle",
            Slot{1, 5, 0}, Slot{2, 5, 0},
            Slot{1, 5, 0}, Slot{2, 5, 0},
        },
    }

    var a, b Slot
    for _, test := range tests {
        t.Logf(
            "Test %s: initial a=%+v, b=%+v - expecting a=%+v, b=%+v",
            test.desc,
            test.initialA, test.initialB,
            test.expectedA, test.expectedB)
        // Sanity check the test itself. Sum of inputs must equal sum of
        // outputs.
        sumInput := test.initialA.Quantity + test.initialB.Quantity
        sumExpectedOutput := test.expectedA.Quantity + test.expectedB.Quantity
        if sumInput != sumExpectedOutput {
            t.Errorf(
                "    Test incorrect: sum of inputs %d != sum of expected outputs %d",
                sumInput, sumExpectedOutput)
            continue
        }

        a = test.initialA
        b = test.initialB
        a.Add(&b)
        if !slotEq(&test.expectedA, &a) || !slotEq(&test.expectedB, &b) {
            t.Errorf("    Fail: got a=%+v, b=%+v", a, b)
        }
    }
}

func TestSlot_Split(t *testing.T) {
    // TODO
}

func TestSlot_Take(t *testing.T) {
    // TODO
}
