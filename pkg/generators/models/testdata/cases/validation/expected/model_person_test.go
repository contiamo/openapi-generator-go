package generatortest

import (
	"testing"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func TestPerson(t *testing.T) {
	name := "foo"
	t.Run("validation passes on valid object", func(t *testing.T) {
		err := validation.Validate(&Person{
			Cron:           "@daily",
			Age:            18,
			Gender:         GenderDefault,
			Name:           "Foo",
			Hostname:       "foobar.com",
			Email:          "foo@localhost.de",
			Uuid:           uuid.NewV4().String(),
			Base64:         "Zm9vYmFyCg==",
			FavoriteColors: []Color{ColorBlue},
			Ip:             "127.0.0.1",
			Ipv4:           "127.0.0.1",
			Ipv6:           "::1",
			Url:            "https://www.google.com",
			Address: Address{
				Street: "foo",
				Name:   &name,
			},
		})
		require.NoError(t, err)
	})

	t.Run("validation throws expected error messages for patterns", func(t *testing.T) {
		err := validation.Validate(&Person{
			Cron:           "@never",
			Pomodoro:       "abc",
			Age:            18,
			Gender:         GenderDefault,
			Name:           "Foo",
			Hostname:       "foobar.com",
			Email:          "foo@localhost.de",
			Uuid:           uuid.NewV4().String(),
			Base64:         "Zm9vYmFyCg==",
			FavoriteColors: []Color{ColorBlue},
			Ip:             "127.0.0.1",
			Ipv4:           "127.0.0.1",
			Ipv6:           "::1",
			Url:            "https://www.google.com",
			Address: Address{
				Street: "foo",
				Name:   &name,
			},
		})
		expected := validation.Errors{"cron": PersonCronPatternError, "pomodoro": validation.ErrMatchInvalid}
		require.Equal(t, err, expected)
	})
}
