package main
import "C"
import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type livro struct {
	ID         string `json:"id"`
	Titulo     string `json:"titulo"`
	Autor      string `json:"autor"`
	Quantidade int    `json:"quantidade"`
}

// Estrutura de dados para guardar os livros, quantidade e os Autores
var livros = []livro{
	{ID: "1", Titulo: "Clean Code", Autor: "Robert Cecil Martin", Quantidade: 5},
	{ID: "2", Titulo: "Learning Go", Autor: "Jon Bodner", Quantidade: 10},
	{ID: "3", Titulo: "Programming Rust: Fast, Safe Systems Develop", Autor: "Jim Blandy, Jason Orendorff", Quantidade: 2},
}

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, livros)
}

func BookById(c *gin.Context) {
	id := c.Param("id")
	livro, err := getBookByID(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Livro não encontrado"})
		return
	}

	c.IndentedJSON(http.StatusOK, livro)
}

func checkOutBook(c *gin.Context) {
	id, ok := c.GetQuery("id")

	// Verificar se o ID existe
	if ok == false {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Não existe esse ID"})
		return
	}

	// Verificar se o livro está disponível
	livro, err := getBookByID(id)

	// Verificar se o livro está disponível
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Livro não encontrado"})
		return
	}

	// Verificar se o livro está disponível
	if livro.Quantidade <= 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Livro não disponível"})
		return
	}

	// Decrementar a quantidade do livro
	livro.Quantidade -= 1
	c.IndentedJSON(http.StatusOK, livro)
}

// Função para retornar o livro
func ReturnBook(c *gin.Context) {
	id, ok := c.GetQuery("id")

	// Verificar se o ID existe
	if ok == false {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Não existe esse ID"})
		return
	}

	// Verificar se o livro está disponível
	livro, err := getBookByID(id)

	// Verificar se o livro está disponível
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Livro não encontrado"})
		return
	}

	// Incrementar a quantidade do livro
	livro.Quantidade += 1
	c.IndentedJSON(http.StatusOK, livro)
}

// Função para ver o livro pelo ID
func getBookByID(id string) (*livro, error) {
	for i, b := range livros {
		if b.ID == id {
			return &livros[i], nil
		}
	}
	return nil, errors.New("Livro não encontrado")
}

// Função para criar um novo livro
func CriarLivro(c *gin.Context) {
	var novoLivro livro

	if err := c.BindJSON(&novoLivro); err != nil {
		return // se tivermos erro, iremos da return
	}
	// se nao tivermos erro iremos poder adicionar um novo livro
	livros = append(livros, novoLivro)
	c.IndentedJSON(http.StatusCreated, novoLivro)
}

// Função principal
func main() {
	router := gin.Default()
	router.GET("/livros", getBooks) 
	router.GET("/livros/:id", BookById)
	router.POST("/livros", CriarLivro)
	router.PATCH("/livros", checkOutBook)
	router.PATCH("/devolver", ReturnBook)
	router.Run("localhost:8080")
}
