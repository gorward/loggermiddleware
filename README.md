Logger Middleware for Accesss Logs
====

###How to import
```
import (
	"github.com/gorward/loggermiddleware"
)
```

###Sample Usage
```
func main() {
	http.ListenAndServe(":3000", loggermiddleware.LoggerMiddleware(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "helllo")
		})))
}
```