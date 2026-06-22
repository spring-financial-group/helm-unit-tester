package pkg

import "testing"

func TestParseHelmVersion(t *testing.T) {
	testCases := []struct {
		name    string
		input   string
		want    string
		wantErr bool
	}{
		{
			name:  "helm 3 short",
			input: "v3.21.2+g1259634",
			want:  "v3",
		},
		{
			name:  "helm 4 short",
			input: "v4.34.5+g1259635",
			want:  "v4",
		},
		{
			name:  "helm 2 client short with prefix",
			input: "Client: v2.16.0+g4d4a929",
			want:  "v2",
		},
		{
			name:  "trailing newline",
			input: "v3.21.2+g1259634\n",
			want:  "v3",
		},
		{
			name:  "multi digit major",
			input: "v10.1.0+gabc",
			want:  "v10",
		},
		{
			name:    "empty",
			input:   "",
			wantErr: true,
		},
		{
			name:    "no version",
			input:   "could not find helm",
			wantErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := parseHelmVersion([]byte(tc.input))
			if tc.wantErr {
				if err == nil {
					t.Fatalf("expected an error for input %q, got version %q", tc.input, got)
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error for input %q: %v", tc.input, err)
			}
			if got != tc.want {
				t.Errorf("parseHelmVersion(%q) = %q, want %q", tc.input, got, tc.want)
			}
		})
	}
}
