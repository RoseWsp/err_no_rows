package album

import (
	"database/sql"
	"job/err_no_rows/dal"
	"job/err_no_rows/model/album"

	"github.com/pkg/errors"
)

func AlbumByID(id int64) (*album.Album, error) {
	querySql := "SELECT * FROM album WHERE id=?"
	var album album.Album
	row := dal.DB.QueryRow(querySql, id)
	if err := row.Scan(&album.ID, &album.Title, &album.Artist, &album.Price); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.Wrapf(err, "albumByID(id=%d) no found album", id)
		}
		return nil, errors.Wrapf(err, "albumByID(id=%d) has error", id)
	}
	return &album, nil
}
