from tkinter import *
import tkinter as tk
from typing import Self
import requests
import json as json
import pandas as pd


def armazenar_id():
    global id_var  # Declara a variável global para armazenar o ID
    id_var = id_entry.get()  # Obtém o valor da caixa de entrada
    id_label.config(text=f"ID Armazenado: {id_var}")
    
    
    url = f"https://root:agner102030@basesip.makesystem.com.br/clientes?doc={id_var}"  # Substitua pela URL da API que você deseja acessar
    response = requests.get(url)
   
    if response.status_code == 200:
      
      
      # Instanciando o json
      data = response.json()
     
     
     # Meu parametro
      doc = data[0].get('doc')
     
      id_label_doc = tk.Label(root, text=doc)
      
      
      id_label_doc.pack()
     
      
      #doc = data.get('doc', 'não encontrado')
      
      print(f"{data} DEU CERTO")
      print(f"{doc}")
      
    else:
        print(F"Não encontrado {response.status_code}")
        
        
        
        
  
# Crie a janela principal
root = tk.Tk()

root.geometry("600x400")
root.title("Armazenamento de ID")

# Crie uma caixa de entrada para inserir o ID
id_entry = tk.Entry(root)
id_entry.pack(pady=20)

# Crie um botão para armazenar o ID
armazenar_botao = tk.Button(root, text="Armazenar ID",fg="red",bg="blue", command=armazenar_id)
armazenar_botao.pack()



# Crie uma etiqueta para exibir o ID armazenado
id_label = tk.Label(root, text="")
id_label.pack()

quit_botao = tk.Button(root, text="Quit", command=root.destroy)
quit_botao.pack()
# Inicie a aplicação
root.mainloop()
