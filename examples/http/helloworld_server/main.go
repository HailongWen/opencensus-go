// Copyright 2018, OpenCensus Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"go.opencensus.io/zpages"

	"go.opencensus.io/examples/exporter"
	"go.opencensus.io/plugin/ochttp"
	"go.opencensus.io/stats/view"
	"go.opencensus.io/trace"
)

func main() {
	// Register stats and trace exporters to export the collected data.
	exporter := &exporter.Exporter{}
	view.RegisterExporter(exporter)
	trace.RegisterExporter(exporter)

	// Always trace for this demo.
	trace.SetDefaultSampler(trace.AlwaysSample())

	// Report stats at every second.
	view.SetReportingPeriod(1 * time.Second)

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "hello world")
	})
	http.HandleFunc("/rpcz", zpages.RpczHandler)
	http.HandleFunc("/tracez", zpages.TracezHandler)
	log.Fatal(http.ListenAndServe(":50030", &ochttp.Handler{}))
}
