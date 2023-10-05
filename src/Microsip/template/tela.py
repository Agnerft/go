import PySimpleGUI as sg
import requests
import json
import base64

# Defina o layout da janela
layout = [
    [sg.Text('Digite a URL da Requisição:')],
    [sg.InputText(key='url_input')],
    [sg.Button('Enviar Requisição')],
    [sg.Text('Resposta da Requisição:')],
    [sg.Multiline('', size=(40, 10), key='response_text')],
    [sg.Text('Dados Mapeados:')],
    [sg.Text('', size=(40, 1), key='mapped_data')]
]

# Crie a janela
window = sg.Window('Requisição HTTP e Mapeamento JSON com PySimpleGUI', layout)

# Loop principal da aplicação
while True:
    event, values = window.read()

    # Feche a janela se o usuário clicar no botão Fechar ou sair
    if event == sg.WINDOW_CLOSED:
        break

    if event == 'Enviar Requisição':
        
        url_base = 'https://root:agner102030@basesip.makesystem.com.br/clientes/'
        # Obtenha a URL da entrada
        # Configura o cabeçalho da requisição
      
        
        path_ = values['url_input']
        
        url = url_base + path_
        print(url)

        try:
            # Faça a requisição HTTP
            response = requests.get(url)

            # Verifique se a resposta foi bem-sucedida
            if response.status_code == 200:
                # Exiba a resposta na janela
                response_text = response.text
                
                
                json_string = response.text
                
                json_data = json.loads(json_string)
                
                id = json_data['id']
                doc = json_data['doc']
                cliente = json_data['cliente']
                grupoRecurso = json_data['grupoRecurso']
                linkGvc = json_data['linkGvc']
                porta = json_data['porta']
                ramal = json_data['ramal']
                senha = json_data['senha']
                quant_ramais_open = json_data.get("quantRamaisOpen", [])
                
                for ramal_info in quant_ramais_open:
                    ramalSelect = ramal_info.get("ramalSelect")
                    inUse = ramal_info.get("inUse")
                
                
                
                window['response_text'].update(cliente)
                
                print(f'Cliente: {cliente}')

                
                
            else:
                sg.popup_error(f'Requisição falhou com status {response.status_code}')
        except Exception as e:
            sg.popup_error(f'Ocorreu um erro: {str(e)}')

# Feche a janela e saia
window.close()
