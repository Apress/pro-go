package main

import "io"

type CustomReader struct {
    reader io.Reader
    readCount int
}

func NewCustomReader(reader io.Reader) *CustomReader {
    return &CustomReader { reader, 0 }
}

func (cr *CustomReader) Read(slice []byte) (count int, err error) {
    count, err = cr.reader.Read(slice)
    cr.readCount++
    Printfln("Custom Reader: %v bytes", count)
    if (err == io.EOF) {
        Printfln("Total Reads: %v", cr.readCount)
    }
    return
}

type CustomWriter struct {
    writer io.Writer
    writeCount int
}

func NewCustomWriter(writer io.Writer) * CustomWriter {
    return &CustomWriter{ writer, 0}
}

func (cw *CustomWriter) Write(slice []byte) (count int, err error) {
    count, err = cw.writer.Write(slice)
    cw.writeCount++
    Printfln("Custom Writer: %v bytes", count)
    return
}

func (cw *CustomWriter) Close() (err error) {
    if closer, ok := cw.writer.(io.Closer); ok {
        closer.Close()
    }
    Printfln("Total Writes: %v", cw.writeCount)
    return
}
