import PySimpleGUI as sg
import requests
import json
import base64

sg.theme('BlueMono')

layoutPrincipal = [
    [sg.Text('Se você tem o MicroSIP instalado, clique em Remover, se não em Instalar!')],
    [sg.Button('Instalar'), sg.Button('Remover')]
        
]
# Defina o layout da janela
windowPrincipal = sg.Window('Primeira Janela', layoutPrincipal)


while True:
    eventPrincipal, valoresPrincipais = windowPrincipal.read()
    
    
    if eventPrincipal == sg.WINDOW_CLOSED:
        break
    
    ## BOTÃO DE REMOVER DA TELA PRINCIPAL
    if eventPrincipal == 'Remover':
        
        url_excluir = 'http://localhost:8080/remover'
       
        try:
            
            response_excluir = requests.get(url_excluir)
            
            
            if response_excluir.status_code == 200:
                print('MicroSIP removido com sucesso.')
                [sg.Popup('MicroSIP removido')]
            else:
                print(f'Requisição falhou com status {response_excluir.status_code}')
        except Exception as e:
                print(f'Ocorreu um erro: {str(e)}')


      ## BOTÃO DE INSTALAR DA TELA PRINCIPAL
    elif eventPrincipal == 'Instalar':
        layoutInstalar = [
            [sg.Text('Digite a URL da Requisição:')],
            [sg.InputText(key='url_input')],
            [sg.Button('Enviar Requisição')],
            [sg.Text('Resposta da Requisição:')],
            [sg.Multiline('', size=(40, 10), key='response_text')],
            [sg.Text('Dados Mapeados:')],
            [sg.Text('', size=(40, 1), key='mapped_data')],
            [sg.Button('Fechar')]
        ]
        windowPrincipal.hide()
        windowInstalar = sg.Window('Requisição HTTP e Mapeamento JSON com PySimpleGUI', layoutInstalar, finalize=True)
        
        # Loop principal da aplicação
        while True:
                eventInstalar, valoresInstalar = windowInstalar.read()
                if eventInstalar == sg.WINDOW_CLOSED or eventPrincipal == 'Fechar':
                    break

                if eventInstalar == 'Fechar':
                    print('Clicou')
                
                if eventInstalar == 'Enviar Requisição':
                    
                    url_base = 'https://root:agner102030@basesip.makesystem.com.br/clientes/'
                    # Obtenha a URL da entrada
                    # Configura o cabeçalho da requisição
                
                    
                    path_ = valoresInstalar['url_input']
                    
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
                            
                            if inUse == False:
                            
                                windowInstalar['response_text'].update(ramalSelect)
                            
                            print(f'Cliente: {ramalSelect}') 
                            
                        else:
                            sg.popup_error(f'Requisição falhou com status {response.status_code}')
                    except Exception as e:
                        sg.popup_error(f'Ocorreu um erro: {str(e)}')

                        # Feche a janela e saia
                        windowInstalar.close()
    

 
        
