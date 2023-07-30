package ftp

import (
	"github.com/Deveimer/goofy/pkg/goofy"
	"github.com/jlaffaye/ftp"
	"time"
)

func NewClient(ctx *goofy.Context) (*ftp.ServerConn, error) {
	ftpHost := ctx.Config.Get("FTP_HOST")
	ftpUser := ctx.Config.Get("FTP_USER")
	ftpPassword := ctx.Config.Get("FTP_PASSWORD")

	dialOpts := []ftp.DialOption{ftp.DialWithTimeout(5 * time.Second)}

	dialOpts = append(dialOpts, ftp.DialWithDisabledEPSV(true))

	ftpObj, err := ftp.Dial(ftpHost+":21", dialOpts...)
	if err != nil {
		return nil, err
	}

	err = ftpObj.Login(ftpUser, ftpPassword)
	if err != nil {
		return nil, err
	}

	return ftpObj, nil
}
