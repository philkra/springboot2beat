package beater

import(
    "fmt"
    "time"
    "strings"

    "net/http"
    "io/ioutil"
    "encoding/json"

    "github.com/elastic/beats/libbeat/beat"
    "github.com/elastic/beats/libbeat/logp"
    "github.com/elastic/beats/libbeat/common"
)

//
// Http Response Struct for Concurrent GET's
//
type HttpResponse struct {
    Status string
    Body   []byte
    Url    string
}

//
// Metrics List
//
type MetricsList struct {
    Names []string `json: "names"`
}

//
// Metric Endpoint Response
//
type Metric struct {
    Name string                `json: "statistic"`
    Measurements []Measurement `json: "measurements"`
}

//
// Sub Struct conatining micrometer.io Measurement struct
//
type Measurement struct {
    Statistic string `json: "statistic"`
    Value float32    `json: "value"`
}

//
// Perform a HTTP GET and put the response back on the channel
//
// @param url
// @param ch
//
func DoHttpGet(url string, ch chan<- HttpResponse) {
    client := http.Client{
        Timeout: time.Duration(10 * time.Second),
    }
    request, _ := http.NewRequest("GET", url, nil)
    // TODO -- add basic auth
//    request.SetBasicAuth(bt.config.Username, bt.config.Password)
    response, err := client.Do(request)
    if err != nil {
    	panic(err)
    }
    defer response.Body.Close()

    // Fetch "good" Response Body
    body, err := ioutil.ReadAll(response.Body)
    if err != nil {
        panic(err)
    }

    // Push back to Channel
    ch <- HttpResponse{response.Status, body, url}
}

//
// Process the Metrics Actuator
//
// @access /actuator/metrics
//
func (bt *Springboot2beat) ProcessMetricsActuator(b *beat.Beat) {
    var ch chan HttpResponse = make(chan HttpResponse)
    url := fmt.Sprintf("%s/actuator/metrics", bt.config.Host)

    // Fetch all available Metric Endpoints
    go DoHttpGet(url, ch)
    response := <-ch
    list     := MetricsList{}
    json.Unmarshal(response.Body, &list)

    // Commit Metric Requests
    for _, metric := range(list.Names) {
        if existsIn(metric, bt.config.Exclude) == false {
            go DoHttpGet(fmt.Sprintf("%s/%s", url, metric), ch)
        }
    }

    // Init Metrics Map
    fields := common.MapStr{
        "type": b.Info.Name,
    }

    // Join Results to Event
    for range(list.Names) {
        response := <-ch
        metric   := Metric{}
        json.Unmarshal(response.Body, &metric)
        prefix := strings.Replace(metric.Name, ".", "_", -1)
        for _, measurement := range(metric.Measurements) {
            key := fmt.Sprintf("%s_%s", prefix, strings.ToLower(measurement.Statistic))
            fields[key] = measurement.Value
        }
    }
    close(ch)

    // Push Event
    event := beat.Event{
        Timestamp: time.Now(),
        Fields: fields,
    }
    bt.client.Publish(event)
    logp.Info(fmt.Sprintf("Event with %d metrics sent.", len(fields)))
}
