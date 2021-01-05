package main

import (
	"fmt"
	"github.com/barnett-yuxiang/perfetto_protos/gen/protos/perfetto/trace_processor"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
	"os"
)

func main() {
	cmr := &trace_processor.ComputeMetricResult{
		Result: &trace_processor.ComputeMetricResult_MetricsAsPrototext{
			MetricsAsPrototext: "metrics_as_proto_text",
		},
		Error: nil,
	}
	fmt.Printf("1. %+v\n", cmr)

	data, _ := proto.Marshal(cmr)
	_ = ioutil.WriteFile("./example/compute_metric_result.txt", data, os.ModePerm)

	newData, _ := ioutil.ReadFile("./example/compute_metric_result.txt")
	newCmr := &trace_processor.ComputeMetricResult{}
	_ = proto.Unmarshal(newData, newCmr)

	fmt.Printf("2. %+v\n", newCmr)
}
