package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"github.com/olekukonko/tablewriter"
	"io"
	"log"
	"net"
	"os"
	"strings"
)

var keys = []string{
	"pxname", "svname", "qcur", "qmax", "scur", "smax", "slim", "stot", "bin", "bout",
	"ereq", "econ", "eresp", "status", "rate", "qtime", "ctime", "rtime",
}

var ultraReadableKeys = map[string]string{
	"pxname": "Proxy Name",
	"svname": "Service Name",
	"qcur":   "Current Queue",
	"qmax":   "Max Queue",
	"scur":   "Current Sessions",
	"smax":   "Max Sessions",
	"slim":   "Session Limit",
	"stot":   "Cumulative Connections",
	"bin":    "Bytes In",
	"bout":   "Bytes Out",
	"ereq":   "Request Errors",
	"econ":   "Connection Errors",
	"eresp":  "Response Errors",
	"status": "Status",
	"rate":   "Request Rate",
	"qtime":  "Avg Queue Time",
	"ctime":  "Avg Connect Time",
	"rtime":  "Avg Response Time",
}

// Might use these later
var readableKeys = map[string]string{
	"pxname":         "proxy name",
	"svname":         "service name (FRONTEND for frontend, BACKEND for backend, any name for server/listener)",
	"qcur":           "current queued requests. For the backend this reports the number queued without a server assigned.",
	"qmax":           "max value of qcur",
	"scur":           "current sessions",
	"smax":           "max sessions",
	"slim":           "configured session limit",
	"stot":           "cumulative number of connections",
	"bin":            "bytes in",
	"bout":           "bytes out",
	"dreq":           "requests denied because of security concerns.",
	"dresp":          "responses denied because of security concerns.",
	"ereq":           "request errors. Some of the possible causes are:",
	"econ":           "number of requests that encountered an error trying to connect to a backend server.",
	"eresp":          "response errors. srv_abrt will be counted here also.",
	"wretr":          "number of times a connection to a server was retried.",
	"wredis":         "number of times a request was redispatched to another server.",
	"status":         "status (UP/DOWN/NOLB/MAINT/MAINT(via)...)",
	"weight":         "total weight (backend), server weight (server)",
	"act":            "number of active servers (backend), server is active (server)",
	"bck":            "number of backup servers (backend), server is backup (server)",
	"chkfail":        "number of failed checks.",
	"chkdown":        "number of UP->DOWN transitions.",
	"lastchg":        "number of seconds since the last UP<->DOWN transition",
	"downtime":       "total downtime (in seconds).",
	"qlimit":         "configured maxqueue for the server, or nothing in the value is 0",
	"pid":            "process id (0 for first instance, 1 for second, ...)",
	"iid":            "unique proxy id",
	"sid":            " server id (unique inside a proxy)",
	"throttle":       "current throttle percentage for the server, when slowstart is active.",
	"lbtot":          "total number of times a server was selected, either for new sessions, or when re-dispatching. ",
	"tracked":        "id of proxy/server if tracking is enabled.",
	"type":           "(0=frontend, 1=backend, 2=server, 3=socket/listener)",
	"rate":           "number of sessions per second over last elapsed second",
	"rate_lim":       " configured limit on new sessions per second",
	"rate_max":       "max number of new sessions per second",
	"check_status":   "status of last health check, one of:",
	"check_code":     "layer5-7 code, if available",
	"check_duration": "time in ms took to finish last health check",
	"hrsp_1xx":       "http responses with 1xx code",
	"hrsp_2xx":       "http responses with 2xx code",
	"hrsp_3xx":       "http responses with 3xx code",
	"hrsp_4xx":       "http responses with 4xx code",
	"hrsp_5xx":       "http responses with 5xx code",
	"hrsp_other":     "http responses with other codes (protocol error)",
	"hanafail":       "failed health checks details",
	"req_rate":       "HTTP requests per second over last elapsed second",
	"req_rate_max":   "max number of HTTP requests per second observed",
	"req_tot":        "total number of HTTP requests received",
	"cli_abrt":       "number of data transfers aborted by the client",
	"srv_abrt":       "number of data transfers aborted by the server",
	"comp_in":        "number of HTTP response bytes fed to the compressor",
	"comp_out":       "number of HTTP response bytes emitted by the compressor",
	"comp_byp":       "number of bytes that bypassed the HTTP compressor",
	"comp_rsp":       "number of HTTP responses that were compressed",
	"lastsess":       "number of seconds since last session assigned to server/backend",
	"last_chk":       "last health check contents or textual error",
	"last_agt":       "last agent check contents or textual error",
	"qtime":          "the average queue time in ms over the 1024 last requests",
	"ctime":          "the average connect time in ms over the 1024 last requests",
	"rtime":          "the average response time in ms over the 1024 last requests (0 for TCP)",
	"ttime":          "the average total session time in ms over the 1024 last equests",
}

func read(r io.Reader) string {
	buf := make([]byte, 32768)
	for {
		n, err := r.Read(buf[:])
		if err != nil {
			if err.Error() == "EOF" {
				log.Fatal("EOF error on socket")
			}
			log.Fatal("Error Reading socket: ", err)
		}
		return string(buf[0:n])
	}
}

func process(raw string) {
	in := strings.Trim(raw, "# ")
	r := csv.NewReader(strings.NewReader(in))

	rows, err := r.ReadAll()
	if err != nil {
		log.Fatal("Error processing data: ", err)
	}

	keyMap := getKeys(rows[0])

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(keys)
	tableData := [][]string{}

	for _, row := range rows[1:] {
		tableData = append(tableData, buildRow(row, keyMap))
	}

	for _, entry := range tableData {
		table.Append(entry)
	}

	table.Render()
}

func getKeys(row []string) map[string]int {
	result := make(map[string]int)

	for _, key := range keys {
		result[key] = sliceFind(row, key)
	}
	return result
}

func sliceFind(slice []string, key string) int {
	for i, k := range slice {
		if k == key {
			return i
		}
	}

	return -1
}

func buildRow(row []string, keyMap map[string]int) []string {
	result := []string{}

	for _, key := range keys {
		result = append(result, row[keyMap[key]])
	}

	return result
}

func printExp() {
	for _, k := range keys {
		fmt.Println(fmt.Sprintf("%s: %s", k, ultraReadableKeys[k]))
	}
}

func main() {
	info := flag.Bool("info", false, "Prints out helpful definitions")
	socket := flag.String("socket", "/var/run/haproxy", "Haproxy stats socket to connect to")
	flag.Parse()

	if *info == true {
		printExp()
	}

	c, err := net.Dial("unix", *socket)
	if err != nil {
		log.Fatal("Error connecting to socket: ", err)
	}
	defer c.Close()

	_, err = c.Write([]byte("show stat\n"))
	if err != nil {
		log.Fatal("Error writing to socket: ", err)
	}

	raw := read(c)
	process(raw)
}
