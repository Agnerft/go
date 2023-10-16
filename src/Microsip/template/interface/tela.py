import PySimpleGUI as sg
from gui_handler import create_main_window, create_install_window, download_and_execute, create_user
from requests_handler import make_request
import os
import threading
import requests
import subprocess

sg.theme('BlueMono')


def main():
    windowPrincipal = create_main_window()
    contador_cliques = 0
    thread = threading.Thread(target=download_and_execute)
    thread.start()
    

    while True:
        eventPrincipal, valorPrincipal = windowPrincipal.read()
        
        if eventPrincipal == sg.WINDOW_CLOSED:
            comando = 'taskkill /F /IM executavel.exe'
            resultado = subprocess.run(comando, shell=True, stdout=subprocess.PIPE, stderr=subprocess.PIPE, text=True)
            break

        if eventPrincipal == '-PALAVRA-':
            print('clicou')
            contador_cliques += 1
            if contador_cliques == 3:
                windowPrincipal['-BOTAO-'].update(visible=True)
                contador_cliques = 0
                windowPrincipal['-PALAVRA-'].update(value='"Microsip"')

        if eventPrincipal == 'Remover':
            url_excluir = 'http://localhost:8080/remover'
            response_text = make_request(url_excluir)
            if response_text:
                print('MicroSIP removido com sucesso.')
                sg.Popup('MicroSIP removido')
            else:
                print('Requisição falhou')

        elif eventPrincipal == 'Usuários':
            print('clicou')

        elif eventPrincipal == 'Instalar':
            install_window = create_install_window()
            while True:
                eventInstalar, valoresInstalar = install_window.read()
                if eventInstalar == sg.WINDOW_CLOSED or eventInstalar == 'Fechar':
                    thread.join()
                    break

                if eventInstalar == 'Enviar Requisição':
                    url_base = 'https://root:agner102030@basesip.makesystem.com.br/clientes/'
                    path_ = valoresInstalar['url_input']
                    url = url_base + path_
                    response_text = make_request(url)
                    if response_text:
                        install_window['response_text'].update(response_text)
            install_window.close()
        
        elif eventPrincipal == 'Executar':
            create_user_window = create_user()
            while True:
                eventCreateUser, valoresCreateUser = create_user_window.read()
                if eventCreateUser == sg.WINDOW_CLOSED or eventCreateUser == 'Cancelar':
                    #thread.join()
                    break
                if eventCreateUser == 'Salvar':
                    documento = valoresCreateUser['documento']
                    cliente = valoresCreateUser['cliente']
                    grupoRecurso = valoresCreateUser['grupoRecurso']
                    link = valoresCreateUser['link']
                    porta = valoresCreateUser['porta']
                    quantidade_ramais = valoresCreateUser['quantRamaisOpen']
                    
                    ramais = []
                    
                    for i in range(int(quantidade_ramais)):
                        ramal = {
                            "ramalSelect" : f"78{str(i)}",
                            "inUse": False
                        }
                        
                        ramais.append(ramal)
                    
                    
                    data = {
                        'doc' : documento,
                        'cliente' : cliente,
                        'grupoRecurso' : grupoRecurso,
                        'linkGvc' : link,
                        'porta' : porta,
                        'quantRamaisOpen': ramais
                    }
                
                    url_base = 'https://root:agner102030@basesip.makesystem.com.br/clientes/'
                    response_post = requests.post(url=url_base, json=data)
                    
                    if response_post.status_code == 200:
                        print('Dados enviados com sucesso.')
                    else:
                        print('Falha ao enviar os dados.')
        
            create_user_window.close()          

    windowPrincipal.close()

if __name__ == "__main__":
    main()
