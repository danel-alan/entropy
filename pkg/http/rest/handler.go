package rest

import (
	"fmt"
	"log"
	"net/http"

	"github.com/danel-alan/entropy/pkg/http/forms"
	"github.com/danel-alan/entropy/pkg/reporting"
	"github.com/gin-gonic/gin"
)

type ReponseError struct {
	Code  int    `json:"code"`
	Msg   string `json:"msg"`
	Error error  `json:"error"`
}

// Report reports the entropy of file
func ReportFileEntropy(r *reporting.EntropyReporter) gin.HandlerFunc {
	return func(c *gin.Context) {
		var form forms.EntropyReport
		if err := c.ShouldBind(&form); err != nil {
			c.JSON(http.StatusBadRequest, ReponseError{Code: http.StatusBadRequest, Msg: "bad request", Error: err})
			return
		}
		file, err := form.File.Open()
		defer func() {
			err := file.Close()
			if err != nil {
				log.Println(fmt.Sprintf("file close error: %v", err))
			}
		}()
		if err != nil {
			c.JSON(http.StatusInternalServerError, ReponseError{Code: http.StatusInternalServerError, Msg: "cannot open file", Error: err})
			return
		}
		report, err := r.Report(file, form.BlockSize)
		if err != nil {
			c.JSON(http.StatusInternalServerError, ReponseError{Code: http.StatusInternalServerError, Msg: "calculation error", Error: err})
			return
		}
		c.JSON(http.StatusOK, report)
	}
}
