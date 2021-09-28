package iteration

import "testing"

func TestRepeat(t *testing.T) {

  repeated := Repeat("a")
  expected := "aaaaa"

  if repeated != expected {
    t.Errorf("expected %q and got %q", expected, repeated)

  }


}
