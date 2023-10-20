import base64
import requests
import PySimpleGUI as sg

def make_request(url):
    try:
        username = 'root'
        password = 'agner102030'
        headers = {
            'Authorization': f'Basic {base64.b64encode(f"{username}:{password}".encode("utf-8")).decode("utf-8")}'
        }
        
        response = requests.get(url, headers=headers)
        if response.status_code == 200:
            return response
        else:
            sg.popup_error(response.text)
    except Exception as e:
        sg.popup_error(f'Ocorreu um erro: {str(e)}')
    return None

# def search_clients():
    
#     try:
#         username = 'root'
#         password = 'agner102030'
#         headers = {
#             'Authorization': f'Basic {base64.b64encode(f"{username}:{password}".encode("utf-8")).decode("utf-8")}'
#         }
    
#         url = 'https://basesip.makesystem.com.br/clientes'
        
        
#         response = requests.get(url, headers=headers)
#         if response.status_code == 200:
#             data = response.json()
#         # Suponhamos que você deseje pegar o primeiro nome da lista de resultados
#             if data and 'nome' in data[0]:
#                 return data[0]['nome']
#             else:
#                 texto = "Não encontramos o cliente."
#                 return texto
#         else:
#             return "Erro na solicitação"
    
#     except Exception as e:
#         sg.popup_error(f'Ocorreu um erro: {str(e)}')
#     return None




def atualizar_inUse(numero_ramal, novo_valor,quantRamaisOpen):
    for ramal_info in quantRamaisOpen:
        if ramal_info['ramalSelect'] == numero_ramal:
            ramal_info['inUse'] = novo_valor
            
            print(novo_valor)        
    