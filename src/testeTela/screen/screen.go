package screen

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"log"
	"testeTela/data"
)

func ScreenPrincipal() {
	var input1, input2 *walk.LineEdit
	//var info1, info2 string

	MainWindow{
		Title:  "Formulário de Informação",
		Size:   Size{400, 200},
		Layout: VBox{},
		Children: []Widget{
			Composite{
				Layout: HBox{},
				Children: []Widget{
					Label{
						Text: "Informação 1:",
					},
					LineEdit{
						AssignTo: &input1,
					},
				},
			},
			Composite{
				Layout: HBox{},
				Children: []Widget{
					Label{
						Text: "Informação 2:",
					},
					LineEdit{
						AssignTo: &input2,
					},
				},
			},
			PushButton{
				Text: "Salvar",
				OnClicked: func() {
					data.Info1 = input1.Text()
					data.Info2 = input2.Text()
					ScreenSecundaria()

					if data.Info1 == "" || data.Info2 == "" {
						walk.MsgBox(nil, "Erro", "Preencha ambos os campos.", walk.MsgBoxIconError)
						return
					}

					// Aqui você pode usar as variáveis info1 e info2 no seu código.
					log.Printf("Informação 1: %s\n", data.Info1)
					log.Printf("Informação 2: %s\n", data.Info2)

					walk.MsgBox(nil, "Sucesso", "Informações salvas com sucesso!", walk.MsgBoxIconInformation)
				},
			},
		},
	}.Run()

}
