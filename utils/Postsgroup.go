package utils

import (
	"blogBack/common"
	uuid "github.com/satori/go.uuid"
	"time"
)

type QrArchive struct {
	ArchiveDate time.Time //month
	Total       int       //total
	Year        int       // year
	Month       int       // month
}

func ListPostsMonth() ([]*QrArchive, error) {
	DB := common.GetDB()
	var archives []*QrArchive
	querysql := `select DATE_FORMAT(created_at,'%Y-%m') as month,count(*) as total from posts group by month order by month desc`
	rows, err := DB.Raw(querysql).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var archive QrArchive
		var month string
		rows.Scan(&month, &archive.Total)
		archive.ArchiveDate, _ = time.Parse("2006-01", month)
		archive.Year = archive.ArchiveDate.Year()
		archive.Month = int(archive.ArchiveDate.Month())
		archives = append(archives, &archive)
	}
	return archives, nil
}

type PostInfo struct {
	ID        uuid.UUID
	Title     string
	CreatedAt string
}
type YearPost struct {
	Year     string
	PostInfo []*PostInfo
}

// 文章按年分组
func ListPostsYear() ([]*YearPost, error) {
	DB := common.GetDB()
	var years []*YearPost
	querysql := `select DATE_FORMAT(created_at,'%Y') as year from posts group by year order by year desc`
	rows, err := DB.Raw(querysql).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var y YearPost
		var year_str string
		rows.Scan(&year_str)
		y.Year = year_str
		yearquerysql := `select id,created_at,title from posts where year(created_at) = ?`
		yearrows, err := DB.Raw(yearquerysql, year_str).Rows()
		defer yearrows.Close()
		if err != nil {
			panic(err)
		}
		for yearrows.Next() {
			var postinfo PostInfo
			var id uuid.UUID
			var created_at time.Time
			var title string
			yearrows.Scan(&id, &created_at, &title)
			postinfo.ID = id
			postinfo.CreatedAt = created_at.Format("2006-01-02 15:04:05")
			postinfo.Title = title
			y.PostInfo = append(y.PostInfo, &postinfo)
		}
		years = append(years, &y)
	}
	return years, nil
}
