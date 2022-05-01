package html

import (
	"fmt"
	"log"
	"net/http"

	"github.com/danel-alan/entropy/pkg/http/forms"
	"github.com/danel-alan/entropy/pkg/reporting"
	"github.com/gin-gonic/gin"
)

func EntropyPage() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Entropy Calculator",
		})
	}
}

func ReportFileEntropy(r *reporting.EntropyReporter) gin.HandlerFunc {
	return func(c *gin.Context) {
		var form forms.EntropyReport
		if err := c.ShouldBind(&form); err != nil {
			c.HTML(http.StatusInternalServerError, "error.tmpl", nil)
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
			c.HTML(http.StatusInternalServerError, "error.tmpl", nil)
			return
		}
		report, err := r.Report(file, form.BlockSize)
		if err != nil {
			c.HTML(http.StatusInternalServerError, "error.tmpl", nil)
			return
		}
		c.HTML(http.StatusOK, "resp.tmpl", gin.H{"title": "Entropy", "resp": report.String()})
	}
}
