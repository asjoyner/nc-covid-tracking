// nc-covid-tracking exports three rows of CSV, suitable for graphing in Excel.
// The first row is a header, listing the counties in alpha-sorted order.
// The second row is the number of total cases per county.
// The third ros is the number of deaths per county.
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"

	"github.com/golang/glog"
)

const (
	// COVID19URL returns JSON encoded data from the NC DHHS about COVID-19 cases
	// and deaths, by county
	COVID19URL = "https://services.arcgis.com/iFBq2AW9XO0jYYF7/arcgis/rest/services/NCCovid19/FeatureServer/0/query?where=1%3D1&objectIds=&time=&geometry=&geometryType=esriGeometryEnvelope&inSR=&spatialRel=esriSpatialRelIntersects&resultType=none&distance=0.0&units=esriSRUnit_Meter&returnGeodetic=false&outFields=County%2C+Total%2C+Deaths&returnGeometry=false&returnCentroid=false&featureEncoding=esriDefault&multipatchOption=xyFootprint&maxAllowableOffset=&geometryPrecision=&outSR=102100&datumTransformation=&applyVCSProjection=false&returnIdsOnly=false&returnUniqueIdsOnly=false&returnCountOnly=false&returnExtentOnly=false&returnQueryGeometry=false&returnDistinctValues=false&cacheHint=true&orderByFields=&groupByFieldsForStatistics=&outStatistics=&having=&resultOffset=0&resultRecordCount=4000&returnZ=false&returnM=false&returnExceededLimitFeatures=true&quantizationParameters=%7B%22mode%22%3A%22view%22%2C%22originPosition%22%3A%22upperLeft%22%2C%22tolerance%22%3A0.26458386250105853%2C%22extent%22%3A%7B%22xmin%22%3A-84.32161913499993%2C%22ymin%22%3A33.834368624000035%2C%22xmax%22%3A-75.45998072699996%2C%22ymax%22%3A36.58841470600004%2C%22spatialReference%22%3A%7B%22wkid%22%3A4326%2C%22latestWkid%22%3A4326%7D%7D%7D&sqlFormat=none&f=pjson&token="
)

// This code extracts the attributes field from the JSON returned at COVID19URL
// {
//   "features" : [
//     {
//	     "attributes" : {
//	       "County" : "Forsyth",
//	       "Total" : 2,
//	       "Deaths" : 0
//	     }
//     },
//     ...
//   ]
// }
type covidJSON struct {
	Features []map[string]countyData
}

type countyData struct {
	County string
	Total  int
	Deaths int
}

func main() {
	resp, err := http.Get(COVID19URL)
	if err != nil {
		glog.Fatalf("fetching URL: %s\n", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	var cj covidJSON
	if err := json.Unmarshal(body, &cj); err != nil {
		glog.Fatalf("%s\n", err)
	}

	var results []countyData
	for _, feature := range cj.Features {
		results = append(results, feature["attributes"])
	}
	sort.Sort(ByCounty(results))

	for i, county := range results {
		if i+1 < len(results) {
			fmt.Printf("%s,", county.County)
			continue
		}
		fmt.Printf("%s\n", county.County)
	}
	for i, county := range results {
		if i+1 < len(results) {
			fmt.Printf("%d,", county.Total)
			continue
		}
		fmt.Printf("%d\n", county.Total)
	}
	for i, county := range results {
		if i+1 < len(results) {
			fmt.Printf("%d,", county.Deaths)
			continue
		}
		fmt.Printf("%d\n", county.Deaths)
	}
}

// ByCounty is a helper to sort the county data
type ByCounty []countyData

func (a ByCounty) Len() int           { return len(a) }
func (a ByCounty) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByCounty) Less(i, j int) bool { return a[i].County < a[j].County }
