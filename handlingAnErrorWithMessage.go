package itswizard_m_handlingerrors

import (
	"fmt"
	"github.com/itslearninggermany/itswizard_m_basic"
	"github.com/jinzhu/gorm"
	"html/template"
	"net/http"
)

func handlingAnErrorWithMessage(tpl *template.Template, webserverdatabase *gorm.DB, err error, sitename string, headline string, message string, targelclose string, targetSubmit string, buttontext string, user string, w http.ResponseWriter) bool {
	if err != nil {
		site := itswizard_m_basic.Site{
			Sitename: sitename,
			Special: itswizard_m_basic.MessageStruct{
				Headline:     headline,
				Message:      message,
				TargetClose:  targelclose,
				TargetSubmit: targetSubmit,
				Buttontext:   buttontext,
			},
		}

		er := tpl.ExecuteTemplate(w, "message.html", site)
		if er != nil {
			WritingToErrorLog(webserverdatabase, "", fmt.Sprint(er))
			fmt.Println(err)
		}
		return true
	} else {
		return false
	}
}
