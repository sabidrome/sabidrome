package files

import (
    "log/slog"
    "archive/zip"
    "encoding/xml"
    "io"
    "strings"
    "errors"
)


/*
 * EPUB METADATA STRUCTS
 *
 * Read each file only once and store its metadata
 * in a struct that follows EPUB specification
 *
 * The first important file is META-INF/container.xml
 * that is standarized across all EPUB versions
 * and points to the location of the OPF file
 *
 * The second important file is OPF that holds
 * the metadata of the book, this implementation
 * tries to hold only the most relevant fields
 *
 */

// .epub::META-INF/container.xml::container
type Container struct {
    XMLName         xml.Name            `xml:"container"`
    Rootfiles       RootFileListStruct  `xml:"rootfiles"`
}

// .epub::META-INF/container.xml::container/rootfiles
type RootFileListStruct struct {
    Rootfile        RootFileStruct      `xml:"rootfile"`
}

// .epub::META-INF/container.xml::container/rootfiles/rootfile
type RootFileStruct struct {
    FullPath        string              `xml:"full-path,attr"`
    MediaType       string              `xml:"media-type,attr"`
}

// .opf::package
type OpfPackage struct {
    XMLName         xml.Name            `xml:"package"`
    Metadata        MetadataStruct      `xml:"metadata"`
    Manifest        ManifestStruct      `xml:"manifest"`
    Spine           SpineStruct         `xml:"spine"`
}

// .opf::package/metadata
type MetadataStruct struct {
    Title           string      `xml:"title"`
    Language        string      `xml:"language"`
    Creator         string      `xml:"creator"`
    Publisher       string      `xml:"publisher"`
    Description     string      `xml:"description"`
    Date            string      `xml:"date"`
}

// .opf::package/manifest
type ManifestStruct struct {
    Items           []Item      `xml:"item"`
}

// .opf::package/manifest/item
type Item struct {
    Id              string      `xml:"id,attr"`
    Href            string      `xml:"href,attr"`
    MediaType       string      `xml:"media-type,attr"`
}

// .opf::package/spine
type SpineStruct struct {
    Itemrefs        []Itemref   `xml:"itemref"`
}

// .opf::package/spine/itemref
type Itemref struct {
    Idref           string      `xml:"idref,attr"`
}


// Only for testing purposes
func ListEpubFileContent(path string) {

    zipListing, err := zip.OpenReader(path)
    if err != nil {
        slog.Error("Could not open zip file", "path", path)
    }
    defer zipListing.Close()
    for _, file := range zipListing.File {
        slog.Debug("File inside zip", "file", file.Name)
    }

}

/*
 * EPUB METADATA (UN)MARSHALLING
 *
 * To handle XML inside golang we must rely on marshalled data
 *
 * The following functions hide that complexity to only
 * release ready-to-use functions that return metadata structs
 *
 */


// Marshal META-INF/container.xml and return it as byte array
func _GetMarshalledEpubContainerXML(zipPath string) ([]byte, error) {

    // Epub spec guarantees this file
    containerPath := "META-INF/container.xml"

    zipFile, err := zip.OpenReader(zipPath)
    if err != nil {
        slog.Error("Failed to open archive", "path", zipPath)
        return nil, errors.New("File could not be opened")
    }
    defer zipFile.Close()

    var b []byte

    for _, file := range zipFile.File {
        if strings.EqualFold(file.Name, containerPath) {
            v, err := file.Open()
            if err != nil {
                slog.Error("Failed to open archived file", "name", file.Name)
                return []byte(""), errors.New("META-INF/container.xml corrupted or insufficient permissions")
            }
            defer v.Close()

            b, err = io.ReadAll(v)
            if err != nil {
                slog.Error("Failed to read content of archived file", "name", file.Name)
                return []byte(""), errors.New("META-INF/container.xml content could not be readed")
            }
            break
        }
    }
    if b != nil {
        return b, nil
    }
    slog.Error("Zip file is not an Epub container", "path", zipPath)
    return []byte(""), errors.New("META-INF/container.xml not present in Zip container")
}


// Unmarshal META-INF/container.xml and return it as struct
func _GetUnmarshalledEpubContainerXML(marshalledEpubContainer []byte) (Container, error) {

    var containerStruct Container

    err := xml.Unmarshal(marshalledEpubContainer, &containerStruct)
    if err != nil {
        slog.Error("Error unmarhsalling .epub::container.xml", "error", err)
        return Container{}, errors.New("Failed to unmarshall xml")
    }

    return containerStruct, nil

}


// Marshal OPF and return it as byte array
func _GetMarshalledEpubOpfXML(zipPath string, opfPath string) ([]byte, error) {

    zipFile, err := zip.OpenReader(zipPath)
    if err != nil {
        slog.Error("Failed to open archive", "path", zipPath)
        return []byte(""), errors.New("File could not be opened")
    }
    defer zipFile.Close()

    var b []byte

    for _, file := range zipFile.File {
        if strings.EqualFold(file.Name, opfPath) {
            v, err := file.Open()
            if err != nil {
                slog.Error("Failed to open archived file", "name", file.Name)
                return []byte(""), errors.New("OPF corrupted or insufficient permissions")
            }
            defer v.Close()

            b, err = io.ReadAll(v)
            if err != nil {
                slog.Error("Failed to read content of archived file", "name", file.Name)
                return []byte(""), errors.New("OPF content could not be readed")
            }
        }
    }
    if b != nil {
        return b, nil
    }
    slog.Error("OPF not found in Zip container", "zip_path", zipPath, "opf_path", opfPath)
    return []byte(""), errors.New("OPF not present in Epub container")
}


// Unmarshal OPF and return it as byte array
func _GetUnmarshalledEpubOpfXML(marshalledEpubOpf []byte) (OpfPackage, error) {

    var opfPackageStruct OpfPackage

    err := xml.Unmarshal(marshalledEpubOpf, &opfPackageStruct)
    if err != nil {
        slog.Error("Error unmarhsalling OPF file", "error", err)
        return OpfPackage{}, errors.New("Failed to unmarshall xml")
    }

    return opfPackageStruct, nil

}

/*
 * EPUB METADATA GETTERS
 *
 * Public functions to ask for book's metadata
 *
 */


// Receives an epub file path.
// Returns a struct with the EPUB container metadata.
func EpubContainerAsStruct(zipPath string) Container {

    b, err := _GetMarshalledEpubContainerXML(zipPath)
    if err != nil {
        panic(errors.New("Failed to extract Epub metadata\nIs this file a valid Epub?"))
    }

    c, err := _GetUnmarshalledEpubContainerXML(b)
    if err != nil {
        panic(errors.New("Failed to create struct with Epub metadata\nIs this file properly tagged?"))
    }

    return c

}


// Receives an epub file path.
// Returns a struct with the OPF container metadata.
func EpubOpfAsStruct(zipPath string) OpfPackage {

    container := EpubContainerAsStruct(zipPath)

    b, err := _GetMarshalledEpubOpfXML(zipPath, container.Rootfiles.Rootfile.FullPath)
    if err != nil {
        panic(errors.New("Failed to extract OPF metadata\nCould not find OPF file"))
    }

    o, err := _GetUnmarshalledEpubOpfXML(b)
    if err != nil {
        panic(errors.New("Failed to create struct with OPF metadata\nIs this file properly tagged?"))
    }

    return o

}
