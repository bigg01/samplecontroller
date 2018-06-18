package main

import (
	"crypto/tls"
	"encoding/json"
	"html/template"
	"io/ioutil"
	"os"
	"time"
	//"io"

	//"log"
	log "github.com/sirupsen/logrus"
	//"log/syslog"
	"net/http"
	//"net"
)

// https://mholt.github.io/json-to-go/
type netGates struct {
	APIVersion string `json:"apiVersion"`
	Items      []struct {
		APIVersion string `json:"apiVersion"`
		Kind       string `json:"kind"`
		Metadata   struct {
			ClusterName       string    `json:"clusterName"`
			CreationTimestamp time.Time `json:"creationTimestamp"`
			Labels            struct {
				Netzone string `json:"netzone"`
			} `json:"labels"`
			Name            string `json:"name"`
			Namespace       string `json:"namespace"`
			ResourceVersion string `json:"resourceVersion"`
			SelfLink        string `json:"selfLink"`
			UID             string `json:"uid"`
		} `json:"metadata"`
		Spec struct {
			Hostname  string `json:"hostname"`
			Namespace string `json:"namespace"`
			Netzone   string `json:"netzone"`
			Port      int    `json:"port"`
		} `json:"spec"`
	} `json:"items"`
	Kind     string `json:"kind"`
	Metadata struct {
		Continue        string `json:"continue"`
		ResourceVersion string `json:"resourceVersion"`
		SelfLink        string `json:"selfLink"`
	} `json:"metadata"`
}

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.InfoLevel)
}

func main() {

	apiAddr := "https://192.168.64.2:8443"
	//apiToken := "qpFCYXtndlefdk2Z7jRWnNj7A3p27KImrOQpxCMiZUI"
	apiToken := os.Getenv("OPENSHIFT_TOKEN")

	ignoreSSL := "TRUE"

	var client http.Client
	if ignoreSSL == "TRUE" {
		tr := &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}
		client = http.Client{Transport: tr}
	} else {
		client = http.Client{}
	}
	req, err := http.NewRequest("GET", apiAddr+"/apis/o.guggenbuehl.local/v1/namespaces/default/netgates", nil)
	log.Println(req.URL)
	if err != nil {
		log.Fatal("## Error while opening connection to openshift api", err)
	}
	req.Header.Add("Authorization", "Bearer "+apiToken)
	req.Header.Add("Accept", "application/json")

	resp, err := client.Do(req)

	if err != nil {
		log.Fatal("## Error while connecting to:", apiAddr, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		log.Println("Sucessfullty get Reponse")
	}

	log.Println("--> Status respone:", resp.StatusCode)

	temp, _ := ioutil.ReadAll(resp.Body)

	var netgate netGates
	decErr := json.Unmarshal(temp, &netgate)
	if decErr != nil {
		log.Println("## Error decoding json.", err)

	}
	log.Println("found :", len(netgate.Items))

	//fmt.Println(netgate.APIVersion)

	tmpl := template.Must(template.ParseFiles("templates/layout.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, r.RequestURI, r.UserAgent())
		data := netgate
		tmpl.Execute(w, data)
	})

	http.ListenAndServe(":8080", nil)

}
