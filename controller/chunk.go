package controller

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

func ChunkUploadHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Unavailable method", http.StatusMethodNotAllowed)
		return
	}

	fmt.Println("Method: POST")
	fmt.Println("Starting file uploading")

	file, upload_file, err := r.FormFile("file1") // Get file from request

	if err != nil {
		http.Error(w, "Content-type should be multipart/form-data", http.StatusBadRequest) // Check content type
		return
	}

	file_size, err := strconv.Atoi(r.Header.Get("Content-Length"))

	fmt.Println("-----\n")

	fmt.Println("File info: ")
	fmt.Println("File size: ", file_size)
	fmt.Println("File name: ", upload_file.Filename)
	
	fmt.Println("\n-----\n")

	if err != nil {
		http.Error(w, "Missing file size in header", http.StatusBadRequest)
		return
	}

	if file_size > 10*1024*1024 {
		http.Error(w, "File size should be less than 10 MB", http.StatusBadRequest)
		return
	}

	f, err := os.Create("./temp-files/" + upload_file.Filename)

	if err != nil {
		http.Error(w, "Error creating a file", http.StatusInternalServerError)
		return
	}

	if _, err := io.Copy(f, file); err != nil {
		http.Error(w, "Error writing to a file", http.StatusInternalServerError)
		return
	}

	f.Close()

	if err != nil {
		http.Error(w, "Failed to upload file", http.StatusBadRequest)
		return
	}

	fmt.Println("File uploaded successfully!)")

	// response := &Response{
	// 	Success: true,
	// 	Message: "Uploaded successfully",
	// 	Name:    upload_file.Filename,
	// }

	// resp, err := json.Marshal(response)

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// w.Write(resp)
}
