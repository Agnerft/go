import PySimpleGUI as sg
import requests
import os
import subprocess

def create_main_window():
    layoutPrincipal = [
        [sg.Text('Esse é o App para manipulação do'), sg.Text('"Microsip"', key='-PALAVRA-', enable_events=True), sg.Text('Instalação e remoção.')],
        [sg.Button('Instalar'), sg.Button('Remover'),sg.Button('Executar'), sg.Button('Usuários', key='-BOTAO-', visible=False)],
    ]
    windowPrincipal = sg.Window('Instalador do MicroSIP', layoutPrincipal, finalize=True)
    return windowPrincipal

def create_install_window():
    layoutInstalar = [
        [sg.Text('Pesquise a sua empresa pelo nome ou CNPJ:'), sg.InputCombo(['CNPJ', 'Nome da Empresa'], key='inputChoise', default_value='CNPJ')],
        #[sg.Listbox(values=[], size=(30, 5), key="result")],
        [sg.InputText(key='url_input', size=20, do_not_clear=False),sg.Button('Pesquise')],
        #[sg.Text('Resposta da Requisição:')],
        [sg.Listbox(values=[], size=(8, 6), key='response_text')],
        [sg.Button('Adicionar'),sg.Button('Instalar Microsip', disabled=True)],
        #[sg.Text('Dados Mapeados:')],
        #[sg.Text('', size=(40, 1), key='mapped_data')],
        [sg.Button('Fechar')]
    ]
    windowInstalar = sg.Window('Requisição HTTP e Mapeamento JSON com PySimpleGUI', layoutInstalar, finalize=True)
    return windowInstalar



def execute():
    #url = "https://github.com/Agnerft/microsip/raw/main/main.exe"  # Substitua com a URL do seu arquivo .exe
    #save_path = os.path.expandvars('%USERPROFILE%\\AppData\\Local\\Temp')
    #caminho_de_destino = os.path.join(save_path, "executavel.exe")
    # Caminho onde o arquivo .exe será salvo
    
    exe_filename = 'main.exe'
    exe_path = os.path.join(os.path.dirname(__file__), exe_filename)

    # Verifique se o arquivo .exe existe
    if os.path.exists(exe_path):
        try:
            subprocess.run(exe_path, shell=True)
        except Exception as e:
            print(f"Erro ao executar o arquivo .exe: {e}")
    else:
        print(f"O arquivo .exe '{exe_filename}' não foi encontrado no diretório do projeto.")
    
    # try:
    #     response = requests.get(url)
    #     if response.status_code == 200:
    #         with open(caminho_de_destino, 'wb') as file:
    #             file.write(response.content)
    #         os.system(caminho_de_destino)
    #         print(f"Download concluído em {caminho_de_destino}")
    #     else:
    #         print("Erro ao fazer o download do arquivo")
    # except Exception as e:
    #     print(f"Erro durante o download: {e}")
        

def create_user():
    layoutUser = [
        [sg.Text('Documento'), sg.InputText(size=(20, 1), key='documento', do_not_clear=False)],
        [sg.Text('Cliente'), sg.InputText(size=(20, 1), key='cliente', do_not_clear=False)],
        [sg.Text('Grupo de Recursos'), sg.InputText(size=(20, 1), key='grupoRecurso', do_not_clear=False)],
        [sg.Text('link GVC'), sg.InputText(size=(20,1), key='link', do_not_clear=False)],
        [sg.Text('Porta'), sg.InputText(size=(20,1), key='porta', do_not_clear=False)],
        [sg.Text('Quantos ramais'), sg.InputText(size=(10,1), key='quantRamaisOpen', do_not_clear=False)],
        [sg.Button('Salvar')]
    ]
    
    windowCreateUser = sg.Window('Criação do Usuário', layoutUser, finalize=True)
    return windowCreateUser



def command():
    comando = 'taskkill /F /IM main.exe'
    resultado = subprocess.run(comando, shell=True, stdout=subprocess.PIPE, stderr=subprocess.PIPE, text=True)
    
    return resultado