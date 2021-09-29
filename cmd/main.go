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
		rw.Write([]byte(r.Method))

		var (
			msg sync_pb.ClientToServerMessage
		)

		zw, _ := gzip.NewReader(r.Body)
		data, err := ioutil.ReadAll(zw)
		if err := proto.Unmarshal(data, &msg); err != nil {
			logrus.Error(err)
		}

		logrus.Debugf("ClientToServerMessage msg is %s", msg.String())

		var resp sync_pb.ClientToServerResponse
		var birthday string = "123"
		resp.StoreBirthday = &birthday
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
