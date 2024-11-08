package timeago

import "testing"

func TestTrans(t *testing.T) {
	cases := []struct {
		lang   string
		key    string
		result string
	}{
		{"ru", "online", "В сети"},
		{"ru", "second", "секунда"},
		{"ru", "hour", "час"},
		{"ru", "day", "день"},
		{"en", "online", "Online"},
		{"en", "second", "second"},
		{"en", "hour", "hour"},
		{"en", "day", "day"},
	}

	for _, tc := range cases {
		t.Run("returns "+tc.lang+" language", func(test *testing.T) {
			Set("language", tc.lang)

			if result := trans(tc.key); result != tc.result {
				test.Errorf("Result mast be %s but got %s", tc.result, result)
			}
		})
	}
}

func TestSet_for_language(t *testing.T) {
	cases := []struct {
		name  string
		value string
		err   string
	}{
		{"sets language to ru", "ru", "Set must set the `language` variable to `ru` but it didn't"},
		{"sets language to en", "en", "Set must set the `language` variable to `en` but it didn't"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(test *testing.T) {
			Set("language", tc.value)

			if language != tc.value {
				test.Error(tc.err)
			}
		})
	}
}

func TestSet_for_location(t *testing.T) {
	cases := []struct {
		name  string
		value string
		err   string
	}{
		{"sets location to India Delhi", "America/New_York", "Set must set the `location` variable to `America/New_York` but it didn't"},
		{"sets language to Europe/Kiev", "Europe/Kiev", "Set must set the `location` variable to `Europe/Kiev` but it didn't"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(test *testing.T) {
			err := Set("location", tc.value)
			if err != nil {
				panic(err)
			}

			if location != tc.value {
				test.Error(tc.err)
			}
		})
	}

	Set("location", ``)
}
