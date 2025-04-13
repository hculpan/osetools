// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.857
package templates

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import (
	"github.com/hculpan/osetools/cmd/web/templates/components"
)

func AddCharacterForm(appTitle string,
	username string,
	campaignId string,
	authorized bool) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "<!doctype html><html lang=\"en\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = components.Head(appTitle).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 2, "<body><div class=\"d-flex\" style=\"min-height: 100vh;\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = components.Sidebar("Campaigns", username, authorized).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 3, "<div class=\"container mt-5\"><div><a href=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var2 templ.SafeURL = templ.SafeURL("/characters/" + campaignId)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(string(templ_7745c5c3_Var2)))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 4, "\" class=\"btn btn-sm button mb-3\"><svg xmlns=\"http://www.w3.org/2000/svg\" width=\"16\" height=\"16\" fill=\"currentColor\" class=\"bi bi-arrow-return-left\" viewBox=\"0 0 16 16\"><path fill-rule=\"evenodd\" d=\"M14.5 1.5a.5.5 0 0 1 .5.5v4.8a2.5 2.5 0 0 1-2.5 2.5H2.707l3.347 3.346a.5.5 0 0 1-.708.708l-4.2-4.2a.5.5 0 0 1 0-.708l4-4a.5.5 0 1 1 .708.708L2.707 8.3H12.5A1.5 1.5 0 0 0 14 6.8V2a.5.5 0 0 1 .5-.5\"></path></svg> Back to Characters List</a></div><h1>Add New Character</h1><form action=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var3 templ.SafeURL = templ.SafeURL("/add-character")
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(string(templ_7745c5c3_Var3)))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 5, "\" method=\"POST\"><input type=\"number\" class=\"form-control\" id=\"campaignId\" name=\"campaignId\" value=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var4 string
		templ_7745c5c3_Var4, templ_7745c5c3_Err = templ.JoinStringErrs(campaignId)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `cmd/web/templates/addcharacterform.templ`, Line: 31, Col: 110}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var4))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 6, "\" hidden><div class=\"mb-3\"><label for=\"characterName\" class=\"form-label\">Character Name</label> <input type=\"text\" class=\"form-control\" id=\"characterName\" name=\"characterName\" required></div><div class=\"mb-3\"><label for=\"playerName\" class=\"form-label\">Player Name</label> <input type=\"text\" class=\"form-control\" id=\"playerName\" name=\"playerName\" required></div><div class=\"mb-3\"><label for=\"retainer\" class=\"form-label\">Retainer</label> <input type=\"checkbox\" class=\"form-check-input\" id=\"retainer\" name=\"retainer\" value=\"1\"></div><div class=\"mb-3\"><label for=\"totalXp\" class=\"form-label\">Starting XP</label> <input type=\"number\" class=\"form-control\" style=\"max-width: 200px;\" id=\"totalXp\" name=\"totalXp\" value=\"0\" required></div><div class=\"mb-3\"><label for=\"xpBonus\" class=\"form-label\">XP Bonuses</label><div id=\"xpBonusContainer\"><div class=\"input-group mb-3\"><!--                            <input type=\"number\" class=\"form-control me-4\" style=\"max-width: 200px;\" name=\"xpBonuses[]\"\n                                placeholder=\"XP Bonus (%)\" value=\"0\" required>\n                            <input type=\"text\" class=\"form-control\" name=\"xpBonusReasons[]\" placeholder=\"Reason\"\n                                required>\n                            <button type=\"button\" class=\"btn btn-danger\" onclick=\"removeXpBonus(this)\">Remove</button> --></div></div><button type=\"button\" class=\"btn btn-secondary\" onclick=\"addXpBonus()\">Add XP Bonus</button></div><hr><button type=\"submit\" class=\"btn btn-primary\">Add Character</button></form></div></div><script src=\"https://cdn.jsdelivr.net/npm/bootstrap@5.0.2/dist/js/bootstrap.bundle.min.js\" integrity=\"sha384-MrcW6ZMFYlzcLA8Nl+NtUVF0sA7MsXsP1UyJoMp4YLEuNSfAP+JcXn/tWtIaxVXM\" crossorigin=\"anonymous\"></script><script>\n        function addXpBonus() {\n            var container = document.getElementById(\"xpBonusContainer\");\n            var div = document.createElement(\"div\");\n            div.className = \"input-group mb-3\";\n            div.innerHTML = `\n                    <input type=\"number\" class=\"form-control me-4\" style=\"max-width: 200px;\" name=\"xpBonuses[]\"\n                                placeholder=\"XP Bonus (%)\" value=\"0\" required>\n                    <input type=\"text\" class=\"form-control\" name=\"xpBonusReasons[]\" placeholder=\"Reason\"\n                                required>\n                    <button type=\"button\" class=\"btn btn-danger\" onclick=\"removeXpBonus(this)\">Remove</button>\n                `;\n            container.appendChild(div);\n        }\n\n        function removeXpBonus(button) {\n            button.parentElement.remove();\n        }\n    </script></body></html>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

var _ = templruntime.GeneratedTemplate
