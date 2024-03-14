package formatter

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"os"
)

func (fm *Formatter) AddPhotoFromStorage(path string) {
	fm.Message.Photo = []string{path}
	fm.kindofmedia = []int{fromStorage}
	fm.mediatype = []string{"photo"}
}

func (fm *Formatter) AddPhotoFromTG(path string) {
	fm.Message.Photo = []string{path}
	fm.kindofmedia = []int{fromTelegram}
	fm.mediatype = []string{"photo"}
}

func (fm *Formatter) AddPhotoFromInternet(path string) {
	fm.Message.Photo = []string{path}
	fm.kindofmedia = []int{fromInternet}
	fm.mediatype = []string{"photo"}
}
func (fm *Formatter) AddVideoFromStorage(path string) {
	fm.Message.Video = []string{path}
	fm.kindofmedia = []int{fromStorage}
	fm.mediatype = []string{"video"}
}

func (fm *Formatter) AddVideoFromTG(path string) {
	fm.Message.Video = []string{path}
	fm.kindofmedia = []int{fromTelegram}
	fm.mediatype = []string{"video"}
}

func (fm *Formatter) AddVideoFromInternet(path string) {
	fm.Message.Video = []string{path}
	fm.kindofmedia = []int{fromInternet}
	fm.mediatype = []string{"video"}
}

func (fm *Formatter) AddMapOfMedia(arr map[string]string) {
	i := 0
	for key := range arr {
		val, ok := arr[key]
		if ok {
			if val == "photo" {
				fm.Message.Photo = append(fm.Message.Photo, key)
			} else if val == "video" {
				fm.Message.Video = append(fm.Message.Video, key)
			}
		}
		i++
	}
}

func (fm *Formatter) createMediaGroup(buf *bytes.Buffer) (string, error) {
	var (
		file   *os.File
		part   io.Writer
		writer *multipart.Writer
		err    error
		media  string
		p      bool
	)
	writer = multipart.NewWriter(buf)
	for i := 0; i < len(fm.Message.Photo) || i < len(fm.Message.Video); i++ {
		if i < len(fm.Message.Photo) && !p {
			media = fm.Message.Photo[i]
		} else if i < len(fm.Message.Video) && p {
			media = fm.Message.Video[i]
		}
		file, err = os.Open(media)
		if err == nil {
			part, err = writer.CreateFormFile(fm.mediatype[i], media)
		}
		if err == nil {
			_, err = io.Copy(part, file)
		}
		file.Close()
		if i == len(fm.Message.Photo) {
			i = 0
			p = true
		}
	}
	return writer.FormDataContentType(), err
}

func (fm *Formatter) PrepareMedia(buf *bytes.Buffer) (string, error) {
	var (
		file   *os.File
		part   io.Writer
		writer *multipart.Writer
		err    error
		media  string
	)

	writer = multipart.NewWriter(buf)
	if fm.mediatype[0] == "photo" {
		media = fm.Message.Photo[0]
	} else if fm.mediatype[0] == "video" {
		media = fm.Message.Video[0]
	}

	file, err = os.Open(media)
	if err == nil {
		part, err = writer.CreateFormFile(fm.mediatype[0], media)
	}
	if err == nil {
		_, err = io.Copy(part, file)
	}
	if err == nil {
		err = writer.WriteField("chat_id", fmt.Sprintf("%d", fm.Message.ChatID))
	}
	if err == nil {
		err = writer.WriteField("caption", fm.Message.Text)
	}
	if err == nil {
		err = writer.WriteField("reply_markup", fm.Message.ReplyMarkup)
	}
	if err == nil {
		err = writer.WriteField("parse_mode", fm.Message.ParseMode)
	}
	if err == nil {
		err = writer.Close()
	}
	file.Close()

	return writer.FormDataContentType(), err

}

func (fm *Formatter) PrepareMediaForEdit(buf *bytes.Buffer) (string, error) {
	var (
		file   *os.File
		part   io.Writer
		writer *multipart.Writer
		err    error
		media  string
	)
	writer = multipart.NewWriter(buf)
	if fm.mediatype[0] == "photo" {
		media = fm.Message.Photo[0]
	} else if fm.mediatype[0] == "video" {
		media = fm.Message.Video[0]
	}

	file, err = os.Open(media)
	if err == nil {
		part, err = writer.CreateFormFile(fm.mediatype[0], media)
	}
	if err == nil {
		_, err = io.Copy(part, file)
	}
	if err == nil {
		err = writer.WriteField("caption", fm.Message.Text)
	}
	if err == nil {
		err = writer.WriteField("parse_mode", fm.Message.ParseMode)
	}
	if err == nil {
		err = writer.Close()
	}
	file.Close()
	return writer.FormDataContentType(), err
}
