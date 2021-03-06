package handler

import (
	"fmt"
	"net/http"
)

// Root is the homepage for the app.
func Root(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	fmt.Fprint(w, `
<!DOCTYPE html>
<html>
  <head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1">    
    <link href="css/style.css" rel="stylesheet" type="text/css">
  </head>
  <body>
    <div id="timer"></div>
    <div id="menu"></div>
    <script src="js/compiled/ui.js" type="text/javascript"></script>
  </body>
</html>
`)
}
