package admin

import (
    "platform/http/actionresults"
    "platform/http/handling"
)

var sectionNames = []string { "Products", "Categories", "Orders", "Database"}

type AdminHandler struct {
    handling.URLGenerator
}

type AdminTemplateContext struct {
    Sections []string
    ActiveSection string
    SectionUrlFunc func(string) string
}

func (handler AdminHandler) GetSection(section string) actionresults.ActionResult {
    return actionresults.NewTemplateAction("admin.html", AdminTemplateContext {
        Sections: sectionNames,
        ActiveSection: section,
        SectionUrlFunc: func(sec string) string {
            sectionUrl, _ := handler.GenerateUrl(AdminHandler.GetSection, sec)
            return sectionUrl
        },
    })
}
