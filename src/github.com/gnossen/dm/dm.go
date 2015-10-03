package dm

import (
    "io"
    "os"
    "fmt"
)

type DMController interface {
    io.ReadWriter
    ParseCmd(string) (string, error)
}

func ToFile(dm DMController, filename string) error {
    file, err := os.Create(filename)
    if err != nil {
        return err
    }
    defer file.Close()
    _, err = io.Copy(dm, file)
    return err
}

func FromFile(dm DMController, filename string) error {
    file, err := os.Open(filename)
    if err != nil {
        return err
    }
    defer file.Close()
    _, err = io.Copy(file, dm)
    return err
}

type DM struct {
    name string
}

func (dm DM) ParseCmd(cmd string) (string, error) {
    dm.name = cmd
    return fmt.Sprintf("My name is %s!", dm.name), nil
}

func (dm DM) Read(p []byte) (int, error) {
    return 0, nil
}

func (dm DM) Write(p []byte) (int, error) {
    return 0, nil
}
