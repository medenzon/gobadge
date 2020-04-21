# GoBadge

GoBadge is an open source project for displaying information via SVG "badges". GoBadge was inspired by the open source project, **[shields.io](https://shields.io)**, however, this project was developed to solve for one of shields.io's biggest flaws: the one badge per URL limit. GoBadge was engineered specifically for SVG dashboarding, which means that just one URL can dynamically generate a table-style dashboard of SVG badges.

&nbsp;

Dashboards can be dynamically configured to fit your perfect look.

![demo](docs/demo.gif)

## Getting Started

```bash
go get github.com/medenzon/gobadge
```

```go
package main

import (
    "github.com/medenzon/gobadge"
)
```

## Example

Build and run the sample server in `github.com/medenzon/gobadge/examples`.


### HTTP Endpoint

```go
var dashboard = func(writer http.ResponseWriter, request *http.Request) {

	writer.Header().Set("Content-Type", "image/svg+xml")
	writer.Header().Set("Cache-Control", "no-store, no-cache, must-revalidate, post-check=0, pre-check=0")

	param := request.FormValue("columns")
	cols, err := strconv.ParseInt(param, 10, 64)

	if err != nil {
		cols = 1
	}

	canvas := svg.New(writer)       // create new canvas
	view := new(canvas, int(cols))  // create view
	view.Draw()                     // draw view
}
```

### Sample Data

```json
{
    "badges": [
        {
            "label": "stable release",
            "tag": "v2.8.8",
            "color": "green"
        },
        {
            "label": "build",
            "tag": "passsing",
            "color": "green"
        }
    ]
}
```

### Load Sample Data

```go
type data struct {
	Badges []dash.Badge `json:"badges"`
}

func new(canvas *svg.Canvas, columns int) dash.View {

	file, err := ioutil.ReadFile("data.json")

	if err != nil {
		log.Printf("readfile error %v")
	}

	var data data

	log.Printf("%s", file)

	err = json.Unmarshal(file, &data)

	if err != nil {
		log.Printf("unmarshal error %v")
	}

	return dash.View{
		Canvas: canvas,
		Header: dash.Header{
			Title: "App.FrontEnd | April 1, 2020 10:18.06",
		},
		Columns: columns,
		Badges:  data.Badges,
	}
}
```
