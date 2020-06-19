package main
 
import (
    "encoding/json"
    "fmt"
    "net/http"
	"time"
	"os"
)
 
// Data
type Data struct {
    ID int `json:"id"`
    Name string `json:"name"`
    Age int `json:"age"`
}
 
// newData
func newData() []Data {
    dt := []Data{
        Data{
            ID: 1,
        	Name: "Udin",
        	Age: 12,
        },
        Data{
            ID: 2,
        	Name: "Reane",
        	Age: 51,
        },
        Data{
            ID: 3,
        	Name: "Budi",
        	Age: 34,
        },
		Data{
            ID: 4,
        	Name: "Agus",
        	Age: 16,
        },
		Data{
            ID: 5,
        	Name: "Sari",
        	Age: 19,
        },
    }
    return dt
}
 
// getData
func getData(w http.ResponseWriter, r *http.Request) {
    if r.Method == "GET" {
        dt := newData()
        data, err := json.Marshal(dt)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
        }
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusOK)
        w.Write(data)
        return
    }
 
    http.Error(w, "gagal", http.StatusNotFound)
}

func handleTime(w http.ResponseWriter, r *http.Request){
	response := "Halo Keindahan Dunia \t->\t" + time.Now().String()
	fmt.Fprintln(w, response)
}

func getIPAddress(r *http.Request) string {

  ipAddress := r.RemoteAddr
	fwdAddress := r.Header.Get("X-Forwarded-For")
	if fwdAddress != "" {
		ipAddress = fwdAddress
	}
  return ipAddress
}

func handlerPing(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ping listening on 0.0.0.0, port 7000")
	w.Write([]byte("pong"))
	fmt.Println(time.Now(), getIPAddress(r), r.Method, r.RequestURI, r.UserAgent())
}
 
func main() {
	// Set an Environment Variable
	os.Setenv("DEVELOPMENT_PORT", ":7000")

    http.HandleFunc("/", getData)
	http.HandleFunc("/time", handleTime)
	http.HandleFunc("/ping", handlerPing)
    fmt.Println("server running...")
	err  :=  http.ListenAndServe(os.Getenv("DEVELOPMENT_PORT"), nil)

	if  err != nil {
		fmt.Println( "Galat memulai server" , err )
	}
}