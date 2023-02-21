package mystr

import (
	"fmt"
	"strings"
	"testing"
)

const s = "We ask ourselves, who am I to be brilliant, gorgeous, talented, " +
	"fabulous? Actually, who are you not to be? Your playing small soes not " +
	"server the world. There is nothing enlightened about shrinking so that " +
	"other people won't feel insecure around you. We are all meant to shine" +
	", as children do. We were born to make manifest the glory that is " +
	"within us. It's not in some of usl it's in everyone. ANd as we let our " +
	"own light shine, we unconsciously give other people permission to do " +
	"the same. As we are liberated from our own fear, our presence " +
	"automatically liberates others. - Marianne Williamson"
var xs = strings.Split(s, " ")

func TestCat1(t *testing.T) {
	xs := []string{"letters"}
	v := Cat(xs)
	if v != "letters" {
		t.Errorf("Expected 'letters' got '%s'", v)
	}
}
func TestCat2(t *testing.T) {
	xs := []string{"letters", "are"}
	v := Cat(xs)
	if v != "letters are" {
		t.Errorf("Expected 'letters are' got '%s'", v)
	}
}
func TestCat3(t *testing.T) {
	xs := []string{"letters", "are", "what", "make", "words"}
	v := Cat(xs)
	if v != "letters are what make words" {
		t.Errorf("Expected 'letters are what make words' got '%s'", v)
	}
}

func ExampleCat() {
	fmt.Println(Cat([]string{"hi", "there"}))
	// Output:
	// hi there
}

func BenchmarkCat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Cat(xs)
	}
}

func TestJoin1(t *testing.T) {
	xs := []string{"letters"}
	v := Join(xs)
	if v != "letters" {
		t.Errorf("Expected 'letters' got '%s'", v)
	}
}
func TestJoin2(t *testing.T) {
	xs := []string{"letters", "are"}
	v := Join(xs)
	if v != "letters are" {
		t.Errorf("Expected 'letters are' got '%s'", v)
	}
}
func TestJoin3(t *testing.T) {
	xs := []string{"letters", "are", "what", "make", "words"}
	v := Join(xs)
	if v != "letters are what make words" {
		t.Errorf("Expected 'letters are what make words' got '%s'", v)
	}
}

func ExampleJoin() {
	fmt.Println(Join([]string{"hi", "there"}))
	// Output:
	// hi there
}

func BenchmarkJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Join(xs)
	}
}