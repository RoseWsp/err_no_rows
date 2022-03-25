package album

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"job/err_no_rows/model"
	service "job/err_no_rows/service/album"
	"log"
	"net/http"
	"strconv"
)

type writeError struct {
	io.Writer
	err error
}

func (w *writeError) Write(buf []byte) (int, error) {
	if w.err != nil {
		return 0, w.err
	}
	n := 0
	n, w.err = w.Writer.Write(buf)
	return n, w.err
}

func Album(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	fmt.Printf("id=%s\r\n", id)
	idInt, err := strconv.ParseInt(id, 10, 64) //string to int64
	sw := writeError{Writer: w}
	defer func() {
		if sw.err != nil {
			log.Fatal(sw.err)
		}
	}()
	if err != nil {
		sw.Write([]byte(fmt.Sprintf("{Code:-1;Msg:\"ID:%s is parse error\"}", id)))
		return
	}
	data, err := service.AlbumByID(idInt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			sw.Write([]byte(fmt.Sprintf("{Code:-1;Msg:\"id=%s have no found \"}", id)))
			return
		}
		b, err := json.Marshal(model.ResponseModel{
			Code: -1,
			Data: nil,
			Msg:  fmt.Sprintf("%+v", err),
		})
		if err != nil {
			sw.Write([]byte("{Code:-1;Msg:\"Json Marsh Error\"}"))
			return
		}
		sw.Write(b)
		return
	}

	b, err := json.Marshal(model.ResponseModel{
		Code: 1,
		Data: data,
		Msg:  "",
	})

	sw.Write(b)

}
