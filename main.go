package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("FileConfig.txt") // Le programme ci-dessus ouvre un fichier nommé "FileConfig.txt" en utilisant la fonction os.Open().
	if err != nil {
		fmt.Println("File reading error", err) // S'il y a une erreur lors de l'ouverture du fichier, une erreur est affichée et le programme se termine.
		return
	}
	defer file.Close() // Sinon, le fichier est fermé à la fin de la fonction main() en utilisant defer file.Close().

	data := make([]byte, 1024) // Ensuite, une boucle est utilisée pour lire le contenu du fichier par blocs de 1024 octets.
	for {
		n, err := file.Read(data) // La fonction file.Read() est utilisée pour lire les données du fichier dans le tableau de bytes "data".
		if err == io.EOF {        // Si la fin du fichier est atteinte (io.EOF), la boucle se termine.
			break
		}
		if err != nil {
			fmt.Println("File reading error", err) // Si une autre erreur se produit lors de la lecture du fichier, une erreur est affichée et le programme se termine.
			return
		}
		fmt.Println("Read", n, "bytes:", string(data[:n])) // Sinon, le nombre de bytes lus et les données lues sont affichés à l'aide de fmt.Println().
	}
}
