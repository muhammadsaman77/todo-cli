package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"todo-cli/src"
)

func main() {
	next:= true
	for next {
		service:= src.NewService()
		todo:= *service.GetData()
		fmt.Println("===== Daftar Todo =====")
		for i:=1 ;i<= len(todo);i++{
			fmt.Println(strconv.Itoa(i),". ",todo[i-1].Task ," | ", todo[i-1].Status)
		}
		
		scanner := bufio.NewScanner(os.Stdin)

		fmt.Println("===== Pilih Aksi ====")
		fmt.Println("1. Tambah Todo")
		fmt.Println("2. Edit Todo ")
		fmt.Println("3. Update Status")
		fmt.Println("4. Hapus Todo")
		fmt.Println("5. Hapus Semua Todo")
		fmt.Println("6. Keluar Program")
		fmt.Println("Masukkan Pilhan Anda:")
		scanner.Scan()
		pilihan := scanner.Text()
		if pilihan == "1"{
			fmt.Println("Masukkan Todo Baru:")
			scanner.Scan()
			task := scanner.Text()
			service.AddData(task)
		} else if pilihan == "2"{
			fmt.Println("Pilih Nomor Todo yang mau diedit: ")
			scanner.Scan()
			index := scanner.Text()
			fmt.Println("Pilih Nomor Todo yang mau diedit: ")
			scanner.Scan()
			task := scanner.Text()
			indexInt,_ := strconv.Atoi(index)
			service.EditTodo(indexInt-1 ,task)
		}else if pilihan == "3"{
			fmt.Println("Masukkan Nomor Todo: ")
			scanner.Scan()
			index := scanner.Text()
			indexInt,_ := strconv.Atoi(index)
			service.UpdateStatus(indexInt-1 )
		}else if pilihan == "4"{
			fmt.Println("Masukkan Nomor Todo: ")
			scanner.Scan()
			index := scanner.Text()
			indexInt,_ := strconv.Atoi(index)
			service.Delete(indexInt-1)
		}else if pilihan == "5"{
			service.DeleteAll()
		}else if pilihan == "6"{
				next = false
		}else{
					fmt.Println("tes lag")
				}
			}
	
}