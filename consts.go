package main

type TabDataType int

const (
	TabDataType_Placeholder TabDataType = 0
	TabDataType_Data        TabDataType = 1
	TabDataType_Buttons     TabDataType = 2
)

type SourceType string
const(
	SourceType_SQ SourceType = "SQ&formatType=SQ&resourceType=E"
	SourceType_HQ SourceType = "HQ&formatType=HQ&resourceType=2"
)

var (
	SourceType2FileExt = map[SourceType]string{
		SourceType_SQ: ".flac",
		SourceType_HQ: ".mp3",
	}
)