package entity

import (
	"time"
)

var FileFixtures = map[string]File{
	"exampleFileName.jpg": {
		ID:              1000000,
		Photo:           &PhotoFixture19800101_000002_D640C559,
		PhotoID:         1000000,
		PhotoUUID:       "pt9jtdre2lvl0yh7",
		FileUUID:        "ft8es39w45bnlqdw",
		FileName:        "exampleFileName.jpg",
		OriginalName:    "exampleFileNameOriginal.jpg",
		FileHash:        "2cad9168fa6acc5c5c2965ddf6ec465ca42fd818",
		FileModified:    time.Date(2020, 3, 6, 2, 6, 51, 0, time.UTC),
		FileSize:        4278906,
		FileType:        "jpg",
		FileMime:        "image/jpg",
		FilePrimary:     true,
		FileSidecar:     false,
		FileVideo:       false,
		FileMissing:     false,
		FileDuplicate:   false,
		FilePortrait:    false,
		FileWidth:       3648,
		FileHeight:      2736,
		FileOrientation: 1,
		FileAspectRatio: 1.33333,
		FileMainColor:   "green",
		FileColors:      "929299991",
		FileLuminance:   "8836BD496",
		FileDiff:        968,
		FileChroma:      25,
		FileNotes:       "",
		FileError:       "",
		Share:           []FileShare{},
		Sync:            []FileSync{},
		Links:           []Link{},
		CreatedAt:       time.Date(2020, 3, 6, 2, 6, 51, 0, time.UTC),
		CreatedIn:       2,
		UpdatedAt:       time.Date(2020, 3, 28, 14, 6, 0, 0, time.UTC),
		UpdatedIn:       0,
		DeletedAt:       nil,
	},
	"exampleDNGFile.dng": {
		ID:              1000001,
		Photo:           &PhotoFixturePhoto01,
		PhotoID:         1000001,
		PhotoUUID:       "pt9jtdre2lvl0yh8",
		FileUUID:        "ft9es39w45bnlqdw",
		FileName:        "exampleDNGFile.dng",
		OriginalName:    "exampleDNGFile.dng",
		FileHash:        "3cad9168fa6acc5c5c2965ddf6ec465ca42fd818",
		FileModified:    time.Date(2019, 3, 6, 2, 6, 51, 0, time.UTC),
		FileSize:        661858,
		FileType:        "dng",
		FileMime:        "image/dng",
		FilePrimary:     true,
		FileSidecar:     false,
		FileVideo:       false,
		FileMissing:     false,
		FileDuplicate:   false,
		FilePortrait:    false,
		FileWidth:       1200,
		FileHeight:      1600,
		FileOrientation: 6,
		FileAspectRatio: 0.75,
		FileMainColor:   "gold",
		FileColors:      "5552E2222",
		FileLuminance:   "444428399",
		FileDiff:        747,
		FileChroma:      12,
		FileNotes:       "File notes for dng",
		FileError:       "",
		Share:           []FileShare{},
		Sync:            []FileSync{},
		Links:           []Link{},
		CreatedAt:       time.Date(2019, 3, 6, 2, 6, 51, 0, time.UTC),
		CreatedIn:       2,
		UpdatedAt:       time.Date(2020, 3, 28, 14, 6, 0, 0, time.UTC),
		UpdatedIn:       0,
		DeletedAt:       nil,
	},
	"exampleXmpFile.xmp": {
		ID:              1000002,
		Photo:           &PhotoFixturePhoto01,
		PhotoID:         1000001,
		PhotoUUID:       "pt9jtdre2lvl0yh8",
		FileUUID:        "ft1es39w45bnlqdw",
		FileName:        "exampleXmpFile.xmp",
		OriginalName:    "exampleXmpFile.xmp",
		FileHash:        "ocad9168fa6acc5c5c2965ddf6ec465ca42fd818",
		FileModified:    time.Date(2019, 3, 6, 2, 6, 51, 0, time.UTC),
		FileSize:        858,
		FileType:        "xmp",
		FileMime:        "text/xmp",
		FilePrimary:     false,
		FileSidecar:     true,
		FileVideo:       false,
		FileMissing:     false,
		FileDuplicate:   false,
		FilePortrait:    false,
		FileWidth:       0,
		FileHeight:      0,
		FileOrientation: 0,
		FileAspectRatio: 0,
		FileMainColor:   "",
		FileColors:      "",
		FileLuminance:   "",
		FileDiff:        0,
		FileChroma:      0,
		FileNotes:       "File notes for xmp",
		FileError:       "",
		Share:           []FileShare{},
		Sync:            []FileSync{},
		Links:           []Link{},
		CreatedAt:       time.Date(2019, 3, 6, 2, 6, 51, 0, time.UTC),
		CreatedIn:       2,
		UpdatedAt:       time.Date(2020, 3, 28, 14, 6, 0, 0, time.UTC),
		UpdatedIn:       0,
		DeletedAt:       nil,
	},
	"bridge.jpg": {
		ID:              1000003,
		Photo:           &PhotoFixturePhoto04,
		PhotoID:         1000004,
		PhotoUUID:       "pt9jtdre2lvl0y11",
		FileUUID:        "ft2es39w45bnlqdw",
		FileName:        "bridge.jpg",
		OriginalName:    "bridgeOriginal.jpg",
		FileHash:        "pcad9168fa6acc5c5c2965ddf6ec465ca42fd818",
		FileModified:    time.Date(2017, 2, 6, 2, 6, 51, 0, time.UTC),
		FileSize:        961858,
		FileType:        "jpg",
		FileMime:        "image/jpg",
		FilePrimary:     true,
		FileSidecar:     false,
		FileVideo:       false,
		FileMissing:     false,
		FileDuplicate:   false,
		FilePortrait:    false,
		FileWidth:       1200,
		FileHeight:      1600,
		FileOrientation: 6,
		FileAspectRatio: 0.75,
		FileMainColor:   "magenta",
		FileColors:      "225221C1E",
		FileLuminance:   "DC42844C8",
		FileDiff:        986,
		FileChroma:      32,
		FileNotes:       "",
		FileError:       "",
		Share:           []FileShare{},
		Sync:            []FileSync{},
		Links:           []Link{},
		CreatedAt:       time.Date(2019, 1, 1, 2, 6, 51, 0, time.UTC),
		CreatedIn:       2,
		UpdatedAt:       time.Date(2020, 3, 28, 14, 6, 0, 0, time.UTC),
		UpdatedIn:       0,
		DeletedAt:       nil,
	},
	"reunion.jpg": {
		ID:              1000004,
		Photo:           &PhotoFixturePhoto05,
		PhotoID:         1000005,
		PhotoUUID:       "pt9jtdre2lvl0y12",
		FileUUID:        "ft3es39w45bnlqdw",
		FileName:        "reunion.jpg",
		OriginalName:    "reunionOriginal.jpg",
		FileHash:        "acad9168fa6acc5c5c2965ddf6ec465ca42fd818",
		FileModified:    time.Date(2017, 1, 6, 2, 6, 51, 0, time.UTC),
		FileSize:        81858,
		FileType:        "jpg",
		FileMime:        "image/jpg",
		FilePrimary:     true,
		FileSidecar:     false,
		FileVideo:       false,
		FileMissing:     false,
		FileDuplicate:   false,
		FilePortrait:    false,
		FileWidth:       1200,
		FileHeight:      1600,
		FileOrientation: 6,
		FileAspectRatio: 0.75,
		FileMainColor:   "blue",
		FileColors:      "266111000",
		FileLuminance:   "DC42844C8",
		FileDiff:        800,
		FileChroma:      26,
		FileNotes:       "",
		FileError:       "",
		Share:           []FileShare{},
		Sync:            []FileSync{},
		Links:           []Link{},
		CreatedAt:       time.Date(2018, 1, 1, 2, 6, 51, 0, time.UTC),
		CreatedIn:       2,
		UpdatedAt:       time.Date(2029, 3, 28, 14, 6, 0, 0, time.UTC),
		UpdatedIn:       0,
		DeletedAt:       nil,
	},
}

// CreateFileFixtures inserts known entities into the database for testing.
func CreateFileFixtures() {
	for _, entity := range FileFixtures {
		Db().Create(&entity)
	}
}
