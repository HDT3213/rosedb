package set

import (
	"fmt"
	"testing"
)

var key = "my_set"

func InitSet() *Set {
	set := New()

	set.SAdd(key, []byte("a"))
	set.SAdd(key, []byte("b"))
	set.SAdd(key, []byte("c"))
	set.SAdd(key, []byte("d"))
	set.SAdd(key, []byte("e"))
	set.SAdd(key, []byte("f"))

	return set
}

func PrintSetData(s *Set) {
	for k, v := range s.record {
		fmt.Printf("%9s -> ", k)
		for val := range v {
			fmt.Print(val, " ")
		}
		fmt.Println()
	}
	fmt.Println()
}

func TestSet_SAdd(t *testing.T) {
	set := InitSet()

	n := set.SAdd(key, []byte("abcd"))
	t.Log(n)
	PrintSetData(set)
}

func TestSet_SPop(t *testing.T) {
	set := InitSet()
	res := set.SPop(key, 3)
	for _, v := range res {
		t.Log(string(v))
	}

	PrintSetData(set)
}

func TestSet_SIsMember(t *testing.T) {
	set := InitSet()

	isMember := set.SIsMember(key, []byte("a"))
	t.Log(isMember)

	isMember = set.SIsMember(key, []byte("123"))
	t.Log(isMember)
}

func TestSet_SRandMember(t *testing.T) {
	set := InitSet()

	t.Run("normal situation", func(t *testing.T) {
		members := set.SRandMember(key, 4)
		for _, m := range members {
			t.Log(string(m))
		}
	})

	t.Run("count larger than the set card", func(t *testing.T) {
		members := set.SRandMember(key, 12)
		for _, m := range members {
			t.Log(string(m))
		}
	})

	t.Run("count is an negative number", func(t *testing.T) {
		members := set.SRandMember(key, -2)
		for _, m := range members {
			t.Log(string(m))
		}
		members = set.SRandMember(key, -10)
		for _, m := range members {
			t.Log(string(m))
		}
	})
}

func TestSet_SRem(t *testing.T) {
	set := InitSet()

	n := set.SRem(key, []byte("a"))
	n = set.SRem(key, []byte("a"))
	n = set.SRem(key, []byte("c"))
	t.Log(n)
	PrintSetData(set)

	n = set.SRem(key, []byte("ss"))
	n = set.SRem(key, []byte("d"))
	t.Log(n)
	PrintSetData(set)

	n = set.SRem(key, []byte("e"))
	n = set.SRem(key, []byte("x"))
	t.Log(n)
	PrintSetData(set)
}

func TestSet_SMove(t *testing.T) {
	set := InitSet()

	move := set.SMove(key, "set2", []byte("a"))
	t.Log(move)
	move = set.SMove(key, "set2", []byte("f"))
	t.Log(move)
	move = set.SMove(key, "set2", []byte("12332"))
	t.Log(move)

	PrintSetData(set)
}

func TestSet_SCard(t *testing.T) {
	set := InitSet()
	card := set.SCard(key)

	t.Log(card)

	t.Log(set.SCard("aaa"))
}

func TestSet_SMembers(t *testing.T) {
	set := InitSet()

	members := set.SMembers(key)
	for _, m := range members {
		t.Log(string(m))
	}
}

func TestSet_SUnion(t *testing.T) {
	set := InitSet()

	set.SAdd("set2", []byte("h"))
	set.SAdd("set2", []byte("f"))
	set.SAdd("set2", []byte("g"))
	members := set.SUnion(key, "set2")

	for _, m := range members {
		t.Log(string(m))
	}
}

func TestSet_SDiff(t *testing.T) {
	set := InitSet()
	set.SAdd("set2", []byte("a"))
	set.SAdd("set2", []byte("f"))
	set.SAdd("set2", []byte("g"))
	t.Run("normal situation", func(t *testing.T) {
		members := set.SDiff(key, "set2")
		for _, m := range members {
			t.Log(string(m))
		}
	})
	t.Run("one key", func(t *testing.T) {
		members := set.SDiff(key)
		for _, m := range members {
			t.Log(string(m))
		}
	})
}
