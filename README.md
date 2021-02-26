# go-snapas

[![godoc](https://godoc.org/github.com/snapas/go-snapas?status.svg)](https://godoc.org/github.com/snapas/go-snapas)

Official Snap.as Go client library.

## Installation

```bash
go get github.com/snapas/go-snapas
```

## Documentation

See all functionality and usages in the [API documentation](https://developers.snap.as/docs/api/).

### Example usage

```go
import (
	"github.com/snapas/go-snapas"
	"github.com/writeas/go-writeas/v2"
	"log"
)

func main() {
	// Authenticate with Write.as
	wc := writeas.NewClient()
	u, err := wc.LogIn("demo", "demo")
	if err != nil {
		log.Fatal(err)
	}

	// Upload to Snap.as
	sc := snapas.NewClient(u.AccessToken)
	p, err := sc.UploadPhoto(&snapas.PhotoParams{
		FileName: "image.jpg",
	})
	if err != nil {
		wc.LogOut()
		log.Fatal(err)
	}

	// Output final domain
	log.Println(p.URL)

	// Clean up
	wc.LogOut()
}
```

## License

MIT
