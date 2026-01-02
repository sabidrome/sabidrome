package extractor

import (
	"encoding/xml"
)

/*
Container encapsulates all the information
present in the `META-INF/container.xml` file

https://www.w3.org/TR/epub-33/#sec-container.xml-container-elem
*/
type Container struct {
	XMLName   xml.Name  `xml:"container"`
	Version   string    `xml:"version,attr"`
	Xmlns     string    `xml:"xmlns,attr"`
	Rootfiles Rootfiles `xml:"rootfiles"`
	Links     []Links   `xml:"links"`
}

/*
Rootfiles contains a list of package documents
available in the EPUB container

https://www.w3.org/TR/epub-33/#sec-container.xml-rootfiles-elem
*/
type Rootfiles struct {
	Rootfile []Rootfile `xml:"rootfile"`
}

/*
Rootfile identifies the location of one
package document in the EPUB container

https://www.w3.org/TR/epub-33/#sec-container.xml-rootfile-elem
*/
type Rootfile struct {
	FullPath  string `xml:"full-path,attr"`
	MediaType string `xml:"media-type,attr"`
}

/*
Links identifies resources necessary
for the processing of the EPUB container

https://www.w3.org/TR/epub-33/#sec-container.xml-links-elem
*/
type Links struct {
	Link []Link `xml:"link"`
}

/*
Link identifies one resource necessary
for the processing of the EPUB container

https://www.w3.org/TR/epub-33/#sec-container.xml-links-elem
*/
type Link struct {
	Href      string `xml:"href,attr"`
	MediaType string `xml:"media-type,attr,omitempty"`
	Rel       string `xml:"rel,attr"`
}

/*
Package encapsulates all the information
present in the package document XML file

Its location is specified in the struct
attribute Rootfile.FullPath

https://www.w3.org/TR/epub-33/#sec-package-elem
*/
type Package struct {
	XMLName          xml.Name `xml:"package"`
	Dir              string   `xml:"dir,attr,omitempty"`
	Id               string   `xml:"id,attr,omitempty"`
	Prefix           string   `xml:"prefix,attr,omitempty"`
	XmlLang          string   `xml:"xml:lang,attr,omitempty"`
	UniqueIdentifier string   `xml:"unique-identifier,attr"`
	Version          string   `xml:"version,attr"`
	Metadata         Metadata `xml:"metadata"`
	Manifest         Manifest `xml:"manifest"`
	Spine            Spine    `xml:"spine"`
	// Will me implemented if necessary
	/*
		Guide            Guide        `xml:"guide,omitempty"`    // Legacy
		Bindings         Bindings     `xml:"bindings,omitempty"` // Deprecated
		Collection       []Collection `xml:"collection,omitempty"`
	*/
}

/*
Metadata encapsulates meta information

https://www.w3.org/TR/epub-33/#sec-metadata-elem
*/
type Metadata struct {
	Identifier  []DublinCoreIdentifier  `xml:"dc:identifier"` // 1 or more
	Title       []DublinCoreTitle       `xml:"dc:title"`      // 1 or more
	Language    []DublinCoreLanguage    `xml:"dc:language"`   // 1 or more
	Contributor []DublinCoreContributor `xml:"dc:contributor,omitempty"`
	Coverage    []DublinCoreCoverage    `xml:"dc:coverage,omitempty"`
	Creator     []DublinCoreCreator     `xml:"dc:creator,omitempty"`
	Date        []DublinCoreDate        `xml:"dc:date,omitempty"`
	Subject     []DublinCoreSubject     `xml:"dc:subject,omitempty"`
	Description []DublinCoreDescription `xml:"dc:description,omitempty"`
	Format      []DublinCoreFormat      `xml:"dc:format,omitempty"`
	Publisher   []DublinCorePublisher   `xml:"dc:publisher,omitempty"`
	Relation    []DublinCoreRelation    `xml:"dc:relation,omitempty"`
	Rights      []DublinCoreRights      `xml:"dc:rights,omitempty"`
	Source      []DublinCoreSource      `xml:"dc:source,omitempty"`
	Type        []DublinCoreType        `xml:"dc:type,omitempty"`
	Meta        []MetadataMeta          `xml:"meta"`
	Link        []MetadataLink          `xml:"link"`
}

/*
DublinCoreIdentifier contains an identifier
such as a UUID, DOI or ISBN.

https://www.w3.org/TR/epub-33/#dfn-dc-identifier
*/
type DublinCoreIdentifier struct {
	Identifier string `xml:"id,attr,omitempty"`
}

/*
DublinCoreTitle represents an instance
of a name for the EPUB publication.

https://www.w3.org/TR/epub-33/#sec-opf-dctitle
*/
type DublinCoreTitle struct {
	Dir     string `xml:"dir,attr,omitempty"`
	Id      string `xml:"id,attr,omitempty"`
	XmlLang string `xml:"xml:lang,attr,omitempty"`
}

/*
DublinCoreLanguage specifies the language
of the content of the EPUB publication.

https://www.w3.org/TR/epub-33/#sec-opf-dclanguage
*/
type DublinCoreLanguage struct {
	Id string `xml:"id,attr,omitempty"`
}

/*
DublinCoreOptional are optional and follow
a generalized definition

https://www.w3.org/TR/epub-33/#sec-opf-dcmes-optional-def
*/
type DublinCoreOptional struct {
	Dir     string `xml:"dir,attr,omitempty"`
	Id      string `xml:"id,attr,omitempty"`
	XmlLang string `xml:"xml:lang,attr,omitempty"`
}

/*
DublinCoreContributor is used to represent the name
of a person, organization, etc. that played a
secondary role in the creation of the content.

https://www.w3.org/TR/epub-33/#sec-opf-dccontributor
*/
type DublinCoreContributor struct {
	*DublinCoreOptional
}

/*
DublinCoreCreator represents the name
of a person, organization, etc. responsible
for the creation of the content.

https://www.w3.org/TR/epub-33/#sec-opf-dccreator
*/
type DublinCoreCreator struct {
	*DublinCoreOptional
}

/*
DublinCoreDate defines the publication date
of the EPUB publication.

The publication date is not the same as the
last modified date

https://www.w3.org/TR/epub-33/#sec-opf-dcdate
*/
type DublinCoreDate struct {
	*DublinCoreOptional
}

/*
DublinCoreSubject identifies the subject
of the EPUB publication.

https://www.w3.org/TR/epub-33/#sec-opf-dcsubject
*/
type DublinCoreSubject struct {
	*DublinCoreOptional
}

/*
DublinCoreType is used to indicate that
the EPUB publication is of a specialized type.

https://www.w3.org/TR/epub-33/#sec-opf-dctype

Examples: dictionary, preview, teacher-guide

https://idpf.github.io/epub-registries/types/
*/
type DublinCoreType struct {
	*DublinCoreOptional
}

/*
DublinCoreCoverage is listed as a Dublin Core
Optional Element but not documented on EPUB 3.3
*/
type DublinCoreCoverage struct {
	*DublinCoreOptional
}

/*
DublinCoreDescription is listed as a Dublin Core
Optional Element but not documented on EPUB 3.3
*/
type DublinCoreDescription struct {
	*DublinCoreOptional
}

/*
DublinCoreFormat is listed as a Dublin Core
Optional Element but not documented on EPUB 3.3
*/
type DublinCoreFormat struct {
	*DublinCoreOptional
}

/*
DublinCorePublisher is listed as a Dublin Core
Optional Element but not documented on EPUB 3.3
*/
type DublinCorePublisher struct {
	*DublinCoreOptional
}

/*
DublinCoreRelation is listed as a Dublin Core
Optional Element but not documented on EPUB 3.3
*/
type DublinCoreRelation struct {
	*DublinCoreOptional
}

/*
DublinCoreRights is listed as a Dublin Core
Optional Element but not documented on EPUB 3.3
*/
type DublinCoreRights struct {
	*DublinCoreOptional
}

/*
DublinCoreSource is listed as a Dublin Core
Optional Element but not documented on EPUB 3.3
*/
type DublinCoreSource struct {
	*DublinCoreOptional
}

/*
MetadataMeta provides a generic means
of including package metadata.

https://www.w3.org/TR/epub-33/#sec-meta-elem
*/
type MetadataMeta struct {
	Directory string `xml:"dir,attr,omitempty"`
	Id        string `xml:"id,attr,omitempty"`
	Property  string `xml:"property,attr"`
	Refines   string `xml:"refines,attr,omitempty"`
	Schemes   string `xml:"schemes,attr,omitempty"`
	XmlLang   string `xml:"xml:lang,attr,omitempty"`
}

/*
MetadataLink associates resources with an
EPUB publication, such as metadata records.

https://www.w3.org/TR/epub-33/#dfn-link
*/
type MetadataLink struct {
	Href       string `xml:"href,attr"`
	HrefLang   string `xml:"hreflang,attr,omitempty"`
	Id         string `xml:"id,attr,omitempty"`
	MediaType  string `xml:"media-type,attr,omitempty"`
	Properties string `xml:"properties,attr,omitempty"`
	Refines    string `xml:"refines,attr,omitempty"`
	Rel        string `xml:"rel,attr"`
}

/*
Manifest provides an exhaustive list
of publication resources used
in the rendering of the content.

https://www.w3.org/TR/epub-33/#sec-manifest-elem
*/
type Manifest struct {
	Id   string `xml:"id,attr,omitempty"`
	Item []Item `xml:"item"`
}

/*
Item represents a publication resource.

https://www.w3.org/TR/epub-33/#sec-item-elem
*/
type Item struct {
	Fallback     string `xml:"fallback,attr,omitempty"`
	Href         string `xml:"href,attr"`
	Id           string `xml:"id,attr"`
	MediaOverlay string `xml:"media-overlay,attr,omitempty"`
	MediaType    string `xml:"media-type,attr"`
	Properties   string `xml:"properties,attr,omitempty"`
}

/*
Spine defines an ordered list of manifest item
references that represent the default reading order.

https://www.w3.org/TR/epub-33/#sec-spine-elem
*/
type Spine struct {
	Id                       string    `xml:"id,attr,omitempty"`
	PageProgressionDirection string    `xml:"page-progression-direction,attr,omitempty"`
	Toc                      string    `xml:"toc,attr,omitempty"` // Legacy
	ItemRef                  []ItemRef `xml:"itemref"`
}

/*
ItemRef identifies an EPUB content document
or foreign content document in the default reading order.

https://www.w3.org/TR/epub-33/#sec-itemref-elem
*/
type ItemRef struct {
	Id         string `xml:"id,attr,omitempty"`
	IdRef      string `xml:"idref,attr"`
	Linear     string `xml:"linear,attr,omitempty"`
	Properties string `xml:"properties,attr,omitempty"`
}
