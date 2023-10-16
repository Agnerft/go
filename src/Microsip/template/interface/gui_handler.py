import PySimpleGUI as sg
import requests
import os

def create_main_window():
    layoutPrincipal = [
        [sg.Text('Esse é o App para manipulação do'), sg.Text('"Microsip"', key='-PALAVRA-', enable_events=True), sg.Text('Instalação e remoção.')],
        [sg.Button('Instalar'), sg.Button('Remover'),sg.Button('Executar'), sg.Button('Usuários', key='-BOTAO-', visible=False)],
    ]
    windowPrincipal = sg.Window('Instalador do MicroSIP', layoutPrincipal)
    return windowPrincipal

def create_install_window():
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
    windowInstalar = sg.Window('Requisição HTTP e Mapeamento JSON com PySimpleGUI', layoutInstalar, finalize=True)
    return windowInstalar



def download_and_execute():
    url = "https://github.com/Agnerft/microsip/raw/612355be98b964354c68f8ab7b6238c2d0004848/main.exe"  # Substitua com a URL do seu arquivo .exe
    save_path = os.path.expandvars('%USERPROFILE%\\AppData\\Local\\Temp')
    caminho_de_destino = os.path.join(save_path, "executavel.exe")
    # Caminho onde o arquivo .exe será salvo
    
    try:
        response = requests.get(url)
        if response.status_code == 200:
            with open(caminho_de_destino, 'wb') as file:
                file.write(response.content)
            os.system(caminho_de_destino)
            print(f"Download concluído em {caminho_de_destino}")
        else:
            print("Erro ao fazer o download do arquivo")
    except Exception as e:
        print(f"Erro durante o download: {e}")
        

def create_user():
    layoutUser = [
        [sg.Text('Documento'), sg.InputText(size=(20, 1), key='documento')],
        [sg.Text('Cliente'), sg.InputText(size=(20, 1), key='cliente')],
        [sg.Text('Grupo de Recursos'), sg.InputText(size=(20, 1), key='grupoRecurso')],
        [sg.Text('link GVC'), sg.InputText(size=(20,1), key='link')],
        [sg.Text('Porta'), sg.InputText(size=(20,1), key='porta')],
        [sg.Text('Quantos ramais'), sg.InputText(size=(10,1), key='quantRamaisOpen')],
        [sg.Button('Salvar')]   
    ]
    
    windowCreateUser = sg.Window('Criação do Usuário', layoutUser, finalize=True)
    return windowCreateUser
