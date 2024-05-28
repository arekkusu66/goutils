package fsys

import (
	"encoding/json"
	"io"
	"os"
);


type FileSys struct {
    Unm           any
};


func (fs *FileSys) AppendWrite(content []byte, fileName string) error {

    file, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, 0744);

    if err != nil {
        return err;
    };

    defer file.Close();


    if _, err = file.Write(content); err != nil {
        return err;
    };

    return nil;
};


func (fs *FileSys) Write(content []byte, fileName string) error {
    
    file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, 0744);


    if err != nil {
        return err;
    };


    defer file.Close();


    if err := file.Truncate(0); err != nil {
        return err;
    };

    if _, err := file.Seek(0, 0); err != nil {
        return err;
    };

    
    if _, err = file.Write(content); err != nil {
        return err;
    };


    return nil;
};


func (fs *FileSys) WriteString(content string, fileName string) error {
    
    file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, 0744);


    if err != nil {
        return err;
    };


    defer file.Close();


    if err := file.Truncate(0); err != nil {
        return err;
    };

    if _, err := file.Seek(0, 0); err != nil {
        return err;
    };

    
    if _, err = file.Write([]byte(content)); err != nil {
        return err;
    };


    return nil;
};


func (fs *FileSys) Read(fileName string) ([]byte, error) {
    file, err := os.OpenFile(fileName, os.O_RDONLY, 0744);

    if err != nil {
        return nil, err;
    };

    defer file.Close();

    content, err := io.ReadAll(file);

    if err != nil {
        return nil, err;
    };

    
    if fs.Unm != nil {
        if err = json.Unmarshal(content, &fs.Unm); err != nil {
            return nil, err;
        };
    };


    return content, nil;
};