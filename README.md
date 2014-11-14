# RSS Feed Reader Library

Returns Headline objects for a given RSS or Atom feed. It's not perfect, but
it's what I needed for a project, so there. Feel free to clone and do what you
want with it (if it's even worth it). I likely won't care for pull requests.

### Usage

```
import "github.com/renquinn/rss/rss"
```

See example.go for full usage.

```
headlines, err := rss.Get("http://mycoolrssfeed.com/rss.xml")
```

Or when using on Google's AppEngine:

```
c := appengine.NewContext(r)
headlines, err := rss.GetAE(c, "http://mycoolrssfeed.com/rss.xml")
```

A headline object is as follows:

```
type Headline struct {
    Source string
    Title string
    Link string
    Since time.Duration
}
```
