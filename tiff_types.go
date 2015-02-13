// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tiff

type CompressType uint16

// Compression types (defined in various places in the spec and supplements).
const (
	CompressType_None       CompressType = 1     //
	CompressType_CCITT      CompressType = 2     //
	CompressType_G3         CompressType = 3     // Group 3 Fax.
	CompressType_G4         CompressType = 4     // Group 4 Fax.
	CompressType_LZW        CompressType = 5     //
	CompressType_JPEGOld    CompressType = 6     // Superseded by cJPEG.
	CompressType_JPEG       CompressType = 7     //
	CompressType_Deflate    CompressType = 8     // zlib compression.
	CompressType_PackBits   CompressType = 32773 //
	CompressType_DeflateOld CompressType = 32946 // Superseded by cDeflate.
)

type DataType uint16

// Data types (p. 14-16 of the spec).
const (
	DataType_Nil       DataType = iota //  0, invalid
	DataType_Byte                      //  1
	DataType_ASCII                     //  2
	DataType_Short                     //  3
	DataType_Long                      //  4
	DataType_Rational                  //  5
	DataType_SByte                     //  6
	DataType_Undefined                 //  7
	DataType_SShort                    //  8
	DataType_SLong                     //  9
	DataType_SRational                 // 10
	DataType_Float                     // 11
	DataType_Double                    // 12
	DataType_IFD                       // 13
	DataType_Unicode                   // 14
	DataType_Complex                   // 15
	DataType_Long8                     // 16
	DataType_SLong8                    // 17
	DataType_IFD8                      // 18
)

func (d DataType) Valid() bool {
	if d <= DataType_Nil || d > DataType_IFD8 {
		return false
	}
	return true
}

type TagType uint16

// Tags (see p. 28-41 of the spec).
const (
	_                                  TagType = 0     // Type, Num
	TagType_NewSubfileType             TagType = 254   // LONG, 1
	TagType_SubfileType                TagType = 255   // SHORT, 1
	TagType_ImageWidth                 TagType = 256   // SHORT/LONG/LONG8, 1
	TagType_ImageLength                TagType = 257   // SHORT/LONG/LONG8, 1
	TagType_BitsPerSample              TagType = 258   // SHORT, SamplesPerPixel
	TagType_Compression                TagType = 259   // SHORT, 1
	TagType_PhotometricInterpretation  TagType = 262   // SHORT, 1
	TagType_Threshholding              TagType = 263   // SHORT, 1
	TagType_CellWidth                  TagType = 264   // SHORT, 1
	TagType_CellLenght                 TagType = 265   // SHORT, 1
	TagType_FillOrder                  TagType = 266   // SHORT, 1
	TagType_DocumentName               TagType = 269   // ASCII
	TagType_ImageDescription           TagType = 270   // ASCII
	TagType_Make                       TagType = 271   // ASCII
	TagType_Model                      TagType = 272   // ASCII
	TagType_StripOffsets               TagType = 273   // SHORT/LONG/LONG8, StripsPerImage
	TagType_Orientation                TagType = 274   // SHORT, 1
	TagType_SamplesPerPixel            TagType = 277   // SHORT, 1
	TagType_RowsPerStrip               TagType = 278   // SHORT/LONG/LONG8, 1
	TagType_StripByteCounts            TagType = 279   // SHORT/LONG/LONG8, StripsPerImage
	TagType_MinSampleValue             TagType = 280   // SHORT, SamplesPerPixel
	TagType_MaxSampleValue             TagType = 281   // SHORT, SamplesPerPixel
	TagType_XResolution                TagType = 282   // RATIONAL, 1
	TagType_YResolution                TagType = 283   // RATIONAL, 1
	TagType_PlanarConfiguration        TagType = 284   // SHORT
	TagType_PageName                   TagType = 285   // ASCII
	TagType_XPosition                  TagType = 286   // RATIONAL
	TagType_YPosition                  TagType = 287   // RATIONAL
	TagType_FreeOffsets                TagType = 288   // LONG/LONG8
	TagType_FreeByteCounts             TagType = 289   // LONG/LONG8
	TagType_GrayResponseUnit           TagType = 290   // SHORT, 1
	TagType_GrayResponseCurve          TagType = 291   // SHORT, 2**BitPerSample
	TagType_T4Options                  TagType = 292   // LONG, 1
	TagType_T6Options                  TagType = 293   // LONG, 1
	TagType_ResolutionUnit             TagType = 296   // SHORT, 1
	TagType_PageNumber                 TagType = 297   // SHORT, 2
	TagType_TransferFunction           TagType = 301   // SHORT, {1 or SamplesPerPixel}*2**BitPerSample
	TagType_Software                   TagType = 305   // ASCII
	TagType_DateTime                   TagType = 306   // ASCII, 20
	TagType_Artist                     TagType = 315   // ASCII
	TagType_HostComputer               TagType = 316   // ASCII
	TagType_Predictor                  TagType = 317   // SHORT, 1
	TagType_WhitePoint                 TagType = 318   // RATIONAL, 2
	TagType_PrimaryChromaticities      TagType = 319   // RATIONAL, 6
	TagType_ColorMap                   TagType = 320   // SHORT, 3*(2**BitPerSample)
	TagType_HalftoneHints              TagType = 321   // SHORT, 2
	TagType_TileWidth                  TagType = 322   // SHORT/LONG, 1
	TagType_TileLength                 TagType = 323   // SHORT/LONG, 1
	TagType_TileOffsets                TagType = 324   // LONG/LONG8, TilesPerImage
	TagType_TileByteCounts             TagType = 325   // SHORT/LONG, TilesPerImage
	TagType_InkSet                     TagType = 332   // SHORT, 1
	TagType_InkNames                   TagType = 333   // ASCII
	TagType_NumberOfInks               TagType = 334   // SHORT
	TagType_DotRange                   TagType = 336   // BYTE/SHORT, 2 or 2*NumberOfInks
	TagType_TargetPrinter              TagType = 337   // ASCII
	TagType_ExtraSamples               TagType = 338   // BYTE
	TagType_SampleFormat               TagType = 339   // SHORT, SamplesPerPixel
	TagType_SMinSampleValue            TagType = 340   // *, SamplesPerPixel
	TagType_SMaxSampleValue            TagType = 341   // *, SamplesPerPixel
	TagType_TransferRange              TagType = 342   // SHORT, 6
	TagType_JPEGProc                   TagType = 512   // SHORT, 1
	TagType_JPEGInterchangeFormat      TagType = 513   // LONG, 1
	TagType_JPEGInterchangeFormatLngth TagType = 514   // LONG, 1
	TagType_JPEGRestartInterval        TagType = 515   // SHORT, 1
	TagType_JPEGLosslessPredictors     TagType = 517   // SHORT, SamplesPerPixel
	TagType_JPEGPointTransforms        TagType = 518   // SHORT, SamplesPerPixel
	TagType_JPEGQTables                TagType = 519   // LONG, SamplesPerPixel
	TagType_JPEGDCTables               TagType = 520   // LONG, SamplesPerPixel
	TagType_JPEGACTables               TagType = 521   // LONG, SamplesPerPixel
	TagType_YCbCrCoefficients          TagType = 529   // RATIONAL, 3
	TagType_YCbCrSubSampling           TagType = 530   // SHORT, 2
	TagType_YCbCrPositioning           TagType = 531   // SHORT, 1
	TagType_ReferenceBlackWhite        TagType = 532   // LONG, 2*SamplesPerPixel
	TagType_Copyright                  TagType = 33432 // ASCII
)

const (
	TagType_GeoKeyDirectoryTag     TagType = 34735 // SHORT, >= 4
	TagType_GeoDoubleParamsTag     TagType = 34736 // DOUBLE
	TagType_GeoAsciiParamsTag      TagType = 34737 // ASCII
	TagType_ModelTiepointTag       TagType = 33922 // DOUBLE
	TagType_ModelPixelScaleTag     TagType = 33550 // DOUBLE
	TagType_ModelTransformationTag TagType = 34264 // DOUBLE, 16
	TagType_IntergraphMatrixTag    TagType = 33920 // DOUBLE, 17
)

// Photometric interpretation values (see p. 37 of the spec).
const (
	TagValue_Photometric_WhiteIsZero = 0
	TagValue_Photometric_BlackIsZero = 1
	TagValue_Photometric_RGB         = 2
	TagValue_Photometric_Paletted    = 3
	TagValue_Photometric_TransMask   = 4 // transparency mask
	TagValue_Photometric_CMYK        = 5
	TagValue_Photometric_YCbCr       = 6
	TagValue_Photometric_CIELab      = 8
)

// Values for the tPredictor tag (page 64-65 of the spec).
const (
	TagValue_Predictor_None       = 1
	TagValue_Predictor_Horizontal = 2
)

// Values for the tResolutionUnit tag (page 18).
const (
	TagValue_ResolutionUnit_None    = 1
	TagValue_ResolutionUnit_PerInch = 2 // Dots per inch.
	TagValue_ResolutionUnit_PerCM   = 3 // Dots per centimeter.
)