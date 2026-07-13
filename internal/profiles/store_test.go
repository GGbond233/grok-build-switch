package profiles

import (
	"path/filepath"
	"testing"
)

func TestStorePreservesProfileTemplate(t *testing.T) {
	store := NewStore(filepath.Join(t.TempDir(), "profiles.json"))
	created, err := store.Create(Profile{
		Name:           "Responses Provider",
		Template:       "responses",
		UpstreamFormat: "openai_responses",
		BaseURL:        "https://api.example.com/v1",
	})
	if err != nil {
		t.Fatalf("Create() error = %v", err)
	}

	profiles, err := store.List()
	if err != nil {
		t.Fatalf("List() error = %v", err)
	}
	if len(profiles) != 1 {
		t.Fatalf("List() returned %d profiles, want 1", len(profiles))
	}
	if profiles[0].ID != created.ID {
		t.Fatalf("List() profile ID = %q, want %q", profiles[0].ID, created.ID)
	}
	if profiles[0].Template != "responses" {
		t.Fatalf("List() template = %q, want responses", profiles[0].Template)
	}
}
