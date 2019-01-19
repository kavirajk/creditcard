package creditcard

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestType(t *testing.T) {
	cases := []struct {
		name    string
		card    string
		expType string
	}{
		{"Visa", "4111111111111111", "Visa"},
		{"MasterCard", "5500-0000-0000-0004", "MasterCard"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			card, err := Parse(c.card)
			require.NoError(t, err)
			require.Equal(t, c.expType, card.Type())
		})
	}
}

func TestParse(t *testing.T) {
	cases := []struct {
		name       string
		card       string
		shouldFail bool
		err        error
	}{
		{"valid-card", "79927398713", false, nil},
		{"invalid-card", "79927398711", true, ErrInvalid},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			card, err := Parse(c.card)
			if c.shouldFail {
				require.Equal(t, c.err, err)
				return
			}
			require.NoError(t, err)
			require.NotNil(t, card)
		})
	}
}
