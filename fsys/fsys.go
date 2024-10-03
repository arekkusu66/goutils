package fsys

import (
	"encoding/json"
	"io"
	"io/fs"
	"os"
)


type FileSys struct {
    Unm           any
}

//AppendWrite appends content to an exisiting file, or creates it if it doesn't exist
func (*FileSys) AppendWrite(content []byte, fileName string, perms fs.FileMode) error {

    file, err := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, perms)

    if err != nil {
        return err
    }

    defer file.Close()


    if _, err = file.Write(content); err != nil {
        return err
    }

    return nil
}


func (*FileSys) Write(content []byte, fileName string, perms fs.FileMode) error {
    
    file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, perms)

    if err != nil {
        return err
    }

    defer file.Close()


    if err := file.Truncate(0); err != nil {
        return err
    }

    if _, err := file.Seek(0, 0); err != nil {
        return err
    }

    
    if _, err = file.Write(content); err != nil {
        return err
    }


    return nil
}


func (*FileSys) WriteString(content string, fileName string, perms fs.FileMode) error {
    
    file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, perms)


    if err != nil {
        return err
    }

    defer file.Close()


    if err := file.Truncate(0); err != nil {
        return err
    }

    if _, err := file.Seek(0, 0); err != nil {
        return err
    }

    
    if _, err = file.Write([]byte(content)); err != nil {
        return err
    }


    return nil
}


func (fsys *FileSys) Read(fileName string, perms fs.FileMode) ([]byte, error) {

    file, err := os.OpenFile(fileName, os.O_RDONLY, perms)

    if err != nil {
        return nil, err
    }

    defer file.Close()


    content, err := io.ReadAll(file)

    if err != nil {
        return nil, err
    }

    
    if fsys.Unm != nil {
        if err = json.Unmarshal(content, &fsys.Unm); err != nil {
            return nil, err
        }
    }


    return content, nil
}