package pairing

import (
	"fmt"
	"strconv"
	"testing"
)

type TestCase struct {
	operations [][]string
}

func push(key string, priority int) []string {
	return []string{"push", key, strconv.Itoa(priority)}
}

func pop(expected string) []string {
	return []string{"pop", expected}
}

func update(key string, priority int) []string {
	return []string{"update", key, strconv.Itoa(priority)}
}

func empty() []string {
	return []string{"empty"}
}

func testCase(ops ...[]string) TestCase {
	return TestCase{ops}
}

var cases = []TestCase{
	testCase(
		push("a", 1), pop("a"), empty(),
	),
	testCase(
		push("a", 1), push("b", 2), pop("b"), pop("a"), empty(),
	),
	testCase(
		push("a", 10), push("b", 40), push("c", 30), push("d", 20), pop("b"), pop("c"), pop("d"), pop("a"), empty(),
	),
	testCase(
		push("a", 1), push("b", 2), update("a", 3), pop("a"), pop("b"), empty(),
	),
	testCase(
		push("a", 1), push("b", 3), update("a", 2), pop("b"), pop("a"), empty(),
	),
	testCase(
		push("a", 1), push("b", 3), update("a", 4), update("b", 2), pop("a"), pop("b"), empty(),
	),
	testCase(
		push("a", 10), push("b", 40), push("c", 30), push("d", 20), pop("b"), update("c", 0), pop("d"), pop("a"), pop("c"), empty(),
	),
	testCase(
		push("a", 10), push("b", 40), push("c", 30), push("d", 20), update("c", 0), pop("b"), pop("d"), pop("a"), pop("c"), empty(),
	),
	testCase(
		push("a", 10), push("b", 40), push("c", 30), update("c", 0), update("c", 50), update("c", 0), update("c", 50), pop("c"), pop("b"), pop("a"), empty(),
	),
}

func (t TestCase) String() string {
	return fmt.Sprint(t.operations)
}

func Test(t *testing.T) {
	for _, test := range cases {
		t.Run(test.String(), func(t *testing.T) {
			pq := New[string, int]()
			for _, op := range test.operations {
				switch op[0] {
				case "push":
					value, _ := strconv.Atoi(op[2])
					pq.Push(op[1], value)
				case "pop":
					actual := pq.Pop()
					if actual != op[1] {
						t.Errorf("got %v\nwant %v", actual, op[1])
					}
				case "update":
					value, _ := strconv.Atoi(op[2])
					pq.Update(op[1], value)
					fmt.Println(pq.Pretty())
				case "empty":
					actual := pq.Empty()
					if actual != true {
						t.Errorf("got %v\nwant %v", actual, true)
					}
				}
			}
		})
	}

}
