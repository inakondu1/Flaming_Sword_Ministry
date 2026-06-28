package handlers

import (
	"fmt"
	"net/http"
)

// AddSermonHandler - Temporary test
func AddSermonHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Println("✅ AddSermonHandler reached")

	w.Header().Set("Content-Type", "text/html")

	fmt.Fprint(w, `
<!DOCTYPE html>
<html>
<head>
    <title>Add Sermon</title>
</head>
<body>

    <h1>🎉 Add Sermon Page is Working!</h1>

    <p>If you can see this page, your routing is correct.</p>

    <a href="/">Go Back Home</a>

</body>
</html>
`)
}

// ViewSermonsHandler - Temporary test
func ViewSermonsHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Println("✅ ViewSermonsHandler reached")

	w.Header().Set("Content-Type", "text/html")

	fmt.Fprint(w, `
<!DOCTYPE html>
<html>
<head>
    <title>Sermons</title>
</head>
<body>

    <h1>📖 Sermons Page</h1>

    <p>No sermons yet.</p>

    <a href="/admin/add-sermon">Add Sermon</a>

</body>
</html>
`)
}
