package method

import (
	"archive/zip"
	"bytes"
	"fmt"
	"github.com/mholt/archiver"
	"github.com/nwaples/rardecode"
	log "github.com/sirupsen/logrus"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io"
	"os"
	"path/filepath"
	"unicode/utf8"
)

type ZipFile struct {
	ZipFilePath     string
	ZipType         string
	ZipOutputFolder string
	ZipFileNameList []string
}
type UnZipfile struct {
	ZipFilePath     string
	ZipType         string
	ZipOutputFolder string
	ZipFileNameList []string
}

func Utf8ToGBK(text string) (string, error) {
	dst := make([]byte, len(text)*2)
	tr := simplifiedchinese.GB18030.NewEncoder()
	nDst, _, err := tr.Transform(dst, []byte(text), true)
	if err != nil {
		return text, err
	}
	return string(dst[:nDst]), nil
}

func (zipF *ZipFile) Unzip() {

	zipF.ZipFileNameList = make([]string, 0)
	zipF.ZipType = filepath.Ext(zipF.ZipFilePath)
	Mkdir(zipF.ZipOutputFolder)
	if zipF.ZipType == ".zip" {
		// 第一步，打开 zipF 文件
		zipFile, err := zip.OpenReader(zipF.ZipFilePath)
		if err != nil {
			panic(err)
		}
		defer func(zipFile *zip.ReadCloser) {
			err = zipFile.Close()
			if err != nil {
				log.Error(err)
			}
		}(zipFile)
		// 第二步，遍历 zipF 中的文件
		for _, f := range zipFile.File {
			decodeName := ""
			if utf8.Valid([]byte(f.Name)) {
				decodeName = f.Name
			} else {
				i := bytes.NewReader([]byte(f.Name))
				decoder := transform.NewReader(i, simplifiedchinese.GB18030.NewDecoder())
				content, _ := io.ReadAll(decoder)
				decodeName = string(content)
			}
			decodeName = filepath.Base(decodeName)
			filePath := filepath.Join(zipF.ZipOutputFolder, decodeName)
			zipF.ZipFileNameList = append(zipF.ZipFileNameList, decodeName)
			if f.FileInfo().IsDir() {
				// _ = os.MkdirAll(filePath, os.ModePerm)
				continue
			}
			// 创建对应文件夹
			if err = os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
				panic(err)
			}
			// 解压到的目标文件
			dstFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				panic(err)
			}
			file, err := f.Open()
			if err != nil {
				panic(err)
			}
			// 写入到解压到的目标文件
			if _, err := io.Copy(dstFile, file); err != nil {
				panic(err)
			}
			err = dstFile.Close()
			if err != nil {
				return
			}
			err = file.Close()
			if err != nil {
				log.Error(err)
				return
			}
		}
	} else if zipF.ZipType == ".rar" {
		r := archiver.NewRar()
		rarFile := zipF.ZipFilePath
		destination := zipF.ZipOutputFolder + "\\"
		count := 0
		err := r.Walk(rarFile, func(f archiver.File) error {
			if !f.IsDir() {
				count += 1
			}
			return nil
		})
		if err != nil {
			panic(err)
		}
		log.Debug("File count: ", count)

		err = r.Walk(rarFile, func(f archiver.File) error {
			rh, ok := f.Header.(*rardecode.FileHeader)

			if !ok {
				return fmt.Errorf("expected header")
			}
			to := destination + filepath.Base(rh.Name)
			log.Debug("解压出文件:", filepath.Base(rh.Name))
			zipF.ZipFileNameList = append(zipF.ZipFileNameList, filepath.Base(to))
			if !f.IsDir() && !r.OverwriteExisting && IsPathExist(to) {
				return fmt.Errorf("file already exists: %s", to)
			}
			Mkdir(filepath.Dir(to))
			if f.IsDir() {
				return nil
			}
			return WriteNewFile(to, f.ReadCloser, rh.Mode())
		})
	}
}

func (unzip *UnZipfile) Unzip() {
	unzip.ZipFileNameList = make([]string, 0)
	unzip.ZipType = filepath.Ext(unzip.ZipFilePath)
	Mkdir(unzip.ZipOutputFolder)
	if unzip.ZipType == ".zip" {
		// 第一步，打开 unzip 文件
		zipFile, err := zip.OpenReader(unzip.ZipFilePath)
		if err != nil {
			panic(err)
		}
		defer func(zipFile *zip.ReadCloser) {
			_ = zipFile.Close()
		}(zipFile)
		// 第二步，遍历 unzip 中的文件
		for _, f := range zipFile.File {
			decodeName := ""
			if utf8.Valid([]byte(f.Name)) {
				decodeName = f.Name
			} else {
				i := bytes.NewReader([]byte(f.Name))
				decoder := transform.NewReader(i, simplifiedchinese.GB18030.NewDecoder())
				content, _ := io.ReadAll(decoder)
				decodeName = string(content)
			}
			decodeName = filepath.Base(decodeName)
			filePath := filepath.Join(unzip.ZipOutputFolder, decodeName)
			unzip.ZipFileNameList = append(unzip.ZipFileNameList, decodeName)
			if f.FileInfo().IsDir() {
				//_ = os.MkdirAll(filePath, os.ModePerm)
				continue
			}
			log.Debug(filePath)
			// 创建对应文件夹
			if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
				panic(err)
			}
			// 解压到的目标文件
			dstFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				panic(err)
			}
			file, err := f.Open()
			if err != nil {
				panic(err)
			}
			// 写入到解压到的目标文件
			if _, err := io.Copy(dstFile, file); err != nil {
				panic(err)
			}
			_ = dstFile.Close()
			err = file.Close()
			if err != nil {
				log.Error(err)
				return
			}
		}
	} else if unzip.ZipType == ".rar" {
		r := archiver.NewRar()
		rarFile := unzip.ZipFilePath
		destination := unzip.ZipOutputFolder + "\\"
		count := 0
		err := r.Walk(rarFile, func(f archiver.File) error {
			if !f.IsDir() {
				count += 1
			}
			return nil
		})
		if err != nil {
			panic(err)
		}
		log.Debug("File count: ", count)

		err = r.Walk(rarFile, func(f archiver.File) error {
			rh, ok := f.Header.(*rardecode.FileHeader)

			if !ok {
				return fmt.Errorf("expected header")
			}
			to := destination + filepath.Base(rh.Name)
			log.Debug("解压出文件:", filepath.Base(rh.Name))
			unzip.ZipFileNameList = append(unzip.ZipFileNameList, filepath.Base(to))
			if !f.IsDir() && !r.OverwriteExisting && IsPathExist(to) {
				return fmt.Errorf("file already exists: %s", to)
			}
			Mkdir(filepath.Dir(to))

			if f.IsDir() {
				return nil
			}

			return WriteNewFile(to, f.ReadCloser, rh.Mode())
		})
	}
}
