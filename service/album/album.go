package album

import (
	dal "job/err_no_rows/dal/album"
	model "job/err_no_rows/model/album"
)

func AlbumByID(id int64) (*model.Album, error) {
	return dal.AlbumByID(id)
}
