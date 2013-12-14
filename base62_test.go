package base62

import (
  "fmt"
  "testing"
  assert "github.com/pilu/miniassert"
)

func TestEncode(t *testing.T) {
  assert.Equal(t, "0", Encode(0))
  assert.Equal(t, "1B", Encode(99))
}

func TestDecode(t *testing.T) {
  assert.Equal(t, 0,  Decode("0"))
  assert.Equal(t, 99, Decode("1B"))
}

func ExampleEncode() {
  fmt.Println(Encode(99))
  // Output:
  // 1B
}

func ExampleDecode() {
  fmt.Println(Decode("1B"))
  // Output:
  // 99
}
