package main

import (
	"compress/gzip"
	"io/ioutil"
	"net/http"

	sync_pb "sync-server/protos/components/sync/protocol"

	"github.com/golang/protobuf/proto"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetReportCaller(true)

	http.HandleFunc("/chrome-sync/dev/command/", func(rw http.ResponseWriter, r *http.Request) {
		var (
			msg sync_pb.ClientToServerMessage
		)

		zw, _ := gzip.NewReader(r.Body)
		defer zw.Close()

		data, err := ioutil.ReadAll(zw)
		if err := proto.Unmarshal(data, &msg); err != nil {
			logrus.Error(err)
		}

		logrus.Debugf("ClientToServerMessage msg is %s", msg.String())

		var resp sync_pb.ClientToServerResponse
		resp.StoreBirthday = proto.String("123")
		data, err = proto.Marshal(&resp)

		if err != nil {
			logrus.Error(err)
		}

		logrus.Debugf("ClientToServerMessage is %s", resp.String())
		rw.Write(data)
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		logrus.Fatal(err)
	}
}
