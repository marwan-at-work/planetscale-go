package planetscale

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	qt "github.com/frankban/quicktest"
)

func TestOrganizations_List(t *testing.T) {
	c := qt.New(t)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		out := `{
  "data": [
    {
      "id": "my-cool-org",
      "type": "organization",
	  "name": "my-cool-org",
	  "created_at": "2021-01-14T10:19:23.000Z",
	  "updated_at": "2021-01-14T10:19:23.000Z"
    }
  ]
}`

		_, err := w.Write([]byte(out))
		c.Assert(err, qt.IsNil)
	}))

	client, err := NewClient(WithBaseURL(ts.URL))
	c.Assert(err, qt.IsNil)

	ctx := context.Background()

	orgs, err := client.Organizations.List(ctx)

	c.Assert(err, qt.IsNil)
	want := []*Organization{
		{
			Name:      "my-cool-org",
			CreatedAt: time.Date(2021, time.January, 14, 10, 19, 23, 000, time.UTC),
			UpdatedAt: time.Date(2021, time.January, 14, 10, 19, 23, 000, time.UTC),
		},
	}

	c.Assert(orgs, qt.DeepEquals, want)
}

func TestOrganizations_Get(t *testing.T) {
	c := qt.New(t)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		out := `{
      "id": "my-cool-org",
      "type": "organization",
      "name": "my-cool-org",
      "created_at": "2021-01-14T10:19:23.000Z",
      "updated_at": "2021-01-14T10:19:23.000Z"
}`

		_, err := w.Write([]byte(out))
		c.Assert(err, qt.IsNil)
	}))

	client, err := NewClient(WithBaseURL(ts.URL))
	c.Assert(err, qt.IsNil)

	ctx := context.Background()

	org, err := client.Organizations.Get(ctx, &GetOrganizationRequest{
		Organization: "my-cool-org",
	})

	c.Assert(err, qt.IsNil)
	want := &Organization{
		Name:      "my-cool-org",
		CreatedAt: time.Date(2021, time.January, 14, 10, 19, 23, 000, time.UTC),
		UpdatedAt: time.Date(2021, time.January, 14, 10, 19, 23, 000, time.UTC),
	}

	c.Assert(org, qt.DeepEquals, want)
}
