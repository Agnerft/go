import sys
from PyQt5.QtWidgets import QApplication, QMainWindow, QPushButton, QGraphicsView, QGraphicsScene
from PyQt5.QtCore import Qt

class MyMainWindow(QMainWindow):
    def __init__(self):
        super().__init__()

        # Crie a cena e a visualização
        self.scene = QGraphicsScene()
        self.view = QGraphicsView(self.scene)
        self.setCentralWidget(self.view)

        # Crie um botão
        self.button = QPushButton("Botão")
        self.button.setFlag(QPushButton.ItemIsMovable)
        self.button.setPos(50, 50)
        self.scene.addWidget(self.button)

        # Configure a janela principal
        self.setGeometry(100, 100, 800, 600)
        self.setWindowTitle("Arrastar e Soltar Botões")
        self.show()

if __name__ == '__main__':
    app = QApplication(sys.argv)
    window = MyMainWindow()
    sys.exit(app.exec_())
