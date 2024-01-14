package utils

/*
import (

	"github.com/google/uuid"
)

type Image struct {
	Base
	Image      string    `json:"image"`
	UserId     uuid.UUID `json:"user_id"`
	EntityName string    `json:"entity_name"`
	ImageType  string    `json:"image_type"` // 1:profile, 2:post, 3:comment
}



func UploadImage(c *fiber.Ctx,tableName string) ([]string, error) {
	db := db.Client()
	form, _ := c.MultipartForm()
	files := form.File["file_name"]

	var images []string
	for _, file := range files {
		// Dosyayı okur
		src, err := file.Open()
		if err != nil {
			return images, errors.New("can't upload image")
		}
		defer src.Close()

		// Dosyadaki veriyi okur
		var data strings.Builder
		_, err = io.Copy(&data, src)
		if err != nil {
			return images, errors.New("can't upload image")
		}

		if

		// Veriyi base64 formatına çevirir
		base64Str := base64.StdEncoding.EncodeToString([]byte(data.String()))


		images = append(images, string(basbase64Str))
	}

	return images, nil
}

func ImageToBase64(filename string) (string, error) {
	// Dosyayı açar
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	// Dosyayı okur
	fileInfo, _ := file.Stat()
	size := fileInfo.Size()
	buffer := make([]byte, size)
	_, err = file.Read(buffer)
	if err != nil {
		return "", err
	}

	// Base64 formatına çevirir
	base64Str := base64.StdEncoding.EncodeToString(buffer)

	return base64Str, nil
}

func Base64ToImage(base64Str, filename string) error {
	// Base64 verisini decode eder
	decoded, err := base64.StdEncoding.DecodeString(base64Str)
	if err != nil {
		return err
	}

	// Dosyayı oluşturur
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// Dosyaya yazma işlemi
	_, err = io.Copy(file, strings.NewReader(string(decoded)))
	if err != nil {
		return err
	}

	return nil
}

*/
