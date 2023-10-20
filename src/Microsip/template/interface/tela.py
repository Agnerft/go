import PySimpleGUI as sg
from gui_handler import create_main_window, create_install_window, execute, create_user, command
from requests_handler import make_request, atualizar_inUse
import os
import threading
import requests
import subprocess
import base64


sg.theme('BlueMono')


def main():
    resultado = command()
    windowPrincipal = create_main_window()
    contador_cliques = 0
    thread = threading.Thread(target=execute)
    #resultado.stderr
    thread.start()
    username = 'root'
    password = 'agner102030'
       
    

    while True:
        eventPrincipal, valorPrincipal = windowPrincipal.read()
        
        if eventPrincipal == sg.WINDOW_CLOSED:
            print('check')
            #comando = 'taskkill /F /IM executavel.exe'
            #resultado = subprocess.run(comando, shell=True, stdout=subprocess.PIPE, stderr=subprocess.PIPE, text=True)
            command()
            
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
                sg.Popup(response_text.text)
            else:
                print('Requisição falhou')

        elif eventPrincipal == 'Usuários':
            print('clicou')

        elif eventPrincipal == 'Instalar':
            windowPrincipal.hide()
            install_window = create_install_window()
            while True:
                eventInstalar, valoresInstalar = install_window.read()
                #thread.join()
                if eventInstalar == sg.WINDOW_CLOSED or eventInstalar == 'Fechar':
                    resultado = command()
                    resultado.stdout
                    windowPrincipal.un_hide()
                    #thread.join()
                    break
                    
                   
                
                select_option = valoresInstalar['inputChoise']
                if select_option == 'CNPJ':
                    
                    url_base = 'https://basesip.makesystem.com.br/clientes?doc='
                    print('Passei aqui!')
                    
                    
                    
                elif select_option == 'Nome da Empresa':
                
                    url_base = 'https://basesip.makesystem.com.br/clientes?cliente='
                    print('Agora to aqui!')  
                
                
                if eventInstalar == 'Pesquise':
                    
                    path_ = valoresInstalar['url_input']
                    url = url_base + path_
                    response_text = make_request(url)
                    data = response_text.json()
                    
                    quantRamaisOpen = data[0]['quantRamaisOpen']
                    
                    print(quantRamaisOpen)
                                           
                    lista_ramais = [f"{ramal['ramalSelect']}" for ramal in quantRamaisOpen if not ramal['inUse']]
                    
                                                        
                    install_window['response_text'].update(values=lista_ramais)
                    
                       
                    # lista_ramais = [str(ramal_info['ramalSelect']) for ramal_info in quantRamaisOpen]
                            
                if eventInstalar == 'Adicionar':
                    
                    select = valoresInstalar['response_text']
                    
                
                    print(quantRamaisOpen)
                    
             
                    if select:
                            print(f"Selecionei o ramal {select[0]}")
                            
                            atualizar_inUse(select[0], True, quantRamaisOpen)
                         
                            print(data[0]['id'])
                    
                    print(quantRamaisOpen)
                    
                    url_do_servidor = 'https://basesip.makesystem.com.br/clientes/'+str(data[0]['id'])
                    
                    print(url_do_servidor)
                    headers = {
                        'Authorization': f'Basic {base64.b64encode(f"{username}:{password}".encode("utf-8")).decode("utf-8")}'
                    }
                    print(data[0])
                    
                    response_put = requests.put(url_do_servidor,json=data[0], headers=headers)
                    
                    if response_put.status_code == 200:
                        print('Atualizado com sucesso.')
                    else:
                        print(f'Erro na atualização do cliente {data[0]["cliente"]} {response_put.status_code}')
                       
                    install_window['Instalar Microsip'].update(disabled=False)    
                
                
                if eventInstalar == 'Instalar Microsip':
                    print()
                    
                    url_instalador_1 = 'http://localhost:8080/documento/'+str(data[0]['doc'])
                    url_instalador_2 = 'http://localhost:8080/execute'
                    response_install_1 = make_request(url_instalador_1)
                    response_install_2 = make_request(url_instalador_2)
                    print(response_install_1)
                    print(response_install_2)              
                
                        

                            
            install_window.close()
        
        elif eventPrincipal == 'Executar':
            windowPrincipal.hide()
            create_user_window = create_user()
            while True:
                eventCreateUser, valoresCreateUser = create_user_window.read()
                
                
                if eventCreateUser == sg.WINDOW_CLOSED:
                    #command()
                    windowPrincipal.un_hide()
                    
                    break
 
                if eventCreateUser == 'Salvar':
                    documento = valoresCreateUser['documento']
                    cliente = valoresCreateUser['cliente']
                    grupoRecurso = valoresCreateUser['grupoRecurso']
                    link = valoresCreateUser['link']
                    porta = valoresCreateUser['porta']
                    quantidade_ramais = valoresCreateUser['quantRamaisOpen']
                    
                    ramais = []
                    
                    for i in range(7801, 7801 + int(quantidade_ramais)):
                        
                        ramal = {
                            "ramalSelect" : str(i),
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