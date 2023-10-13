import requests
import PySimpleGUI as sg

def make_request(url):
    try:
        response = requests.get(url)
        if response.status_code == 200:
            return response.text
        else:
            sg.popup_error(f'Requisição falhou com status {response.status_code}')
    except Exception as e:
        sg.popup_error(f'Ocorreu um erro: {str(e)}')
    return None
