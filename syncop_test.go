package syncop

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestE2E(t *testing.T) {

	// create engine
	cop := New()

	// create wg
	cop.WgAdd(1)
	go func() {
		defer cop.WgDone()
		// test err
		time.Sleep(1 / 2 * time.Second)
		cop.HandleError(errors.New("fake 1"))
	}()

	cop.WgAdd(1)
	go func() {
		defer cop.WgDone()
		// test no err
		fakeFunc := func() (string, error) {
			return "foo", nil
		}

		res, err := fakeFunc()
		if err != nil {
			cop.HandleError(err)
		}

		t.Log(res)
	}()

	err := cop.ListenForError()
	require.Error(t, err)

}
