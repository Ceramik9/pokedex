package pokeapi

import(
	"testing"
	"github.com/google/go-cmp/cmp"
)

func TestUpdate(t *testing.T) {
	tests := map[string]struct {
		input string
		want []string
	} {
		"first page": {input: "https://pokeapi.co/api/v2/location-area", want: []string{"canalave-city-area"}},
		"second page": {input: "https://pokeapi.co/api/v2/location-area?offset=20&limit=20", want: []string{"mt-coronet-1f-route-216"}},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			var got []string
			for i := 0; i < 1; i++ {
				MapLocations.Update(tc.input)
				got = append(got, MapLocations.Results[0].Name)
			}
			for i := 0; i < 1; i++ {
					diff := cmp.Diff(tc.want, got)
					if diff != "" {
						t.Fatalf("%s", diff)
					}
				}
		})
	}
}
