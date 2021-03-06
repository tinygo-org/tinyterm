package vga

import "tinygo.org/x/tinyfont"

var Regular8pt8bBitmaps = []byte{
	0x7E, 0x81, 0xA5, 0x81, 0x81, 0xBD, 0x99, 0x81, 0x81, 0x7E, 0x7E, 0xFF,
	0xDB, 0xFF, 0xFF, 0xC3, 0xE7, 0xFF, 0xFF, 0x7E, 0x6D, 0xFF, 0xFF, 0xFF,
	0xEF, 0x8E, 0x08, 0x10, 0x71, 0xF7, 0xF7, 0xC7, 0x04, 0x00, 0x18, 0x3C,
	0x3C, 0xE7, 0xE7, 0xE7, 0x18, 0x18, 0x3C, 0x18, 0x3C, 0x7E, 0xFF, 0xFF,
	0x7E, 0x18, 0x18, 0x3C, 0x6F, 0xF6, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
	0xE7, 0xC3, 0xC3, 0xE7, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0x7B, 0x38,
	0x61, 0xCD, 0xE0, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xC3, 0x99, 0xBD, 0xBD,
	0x99, 0xC3, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0x1E, 0x1C, 0x69, 0x97, 0x99,
	0xB3, 0x66, 0xCC, 0xF0, 0x7B, 0x3C, 0xF3, 0xCD, 0xE3, 0x3F, 0x30, 0xC0,
	0x3F, 0x33, 0x3F, 0x30, 0x30, 0x30, 0x30, 0x70, 0xF0, 0xE0, 0x7F, 0x63,
	0x7F, 0x63, 0x63, 0x63, 0x63, 0x67, 0xE7, 0xE6, 0xC0, 0x18, 0x18, 0xDB,
	0x3C, 0xE7, 0x3C, 0xDB, 0x18, 0x18, 0x81, 0x83, 0x87, 0x8F, 0x9F, 0xFE,
	0x78, 0xE1, 0x82, 0x00, 0x02, 0x0C, 0x38, 0xF3, 0xFF, 0xCF, 0x8F, 0x0E,
	0x0C, 0x08, 0x31, 0xEF, 0xCC, 0x30, 0xCF, 0xDE, 0x30, 0xCF, 0x3C, 0xF3,
	0xCF, 0x3C, 0xC0, 0xCF, 0x30, 0x7F, 0xDB, 0xDB, 0xDB, 0x7B, 0x1B, 0x1B,
	0x1B, 0x1B, 0x1B, 0x7D, 0x8D, 0x81, 0xC6, 0xD8, 0xF1, 0xB6, 0x38, 0x1B,
	0x1B, 0xE0, 0xFF, 0xFF, 0xFF, 0xF0, 0x31, 0xEF, 0xCC, 0x30, 0xCF, 0xDE,
	0x33, 0xF0, 0x31, 0xEF, 0xCC, 0x30, 0xC3, 0x0C, 0x30, 0xC0, 0x30, 0xC3,
	0x0C, 0x30, 0xC3, 0x3F, 0x78, 0xC0, 0x18, 0x1B, 0xF8, 0x61, 0x80, 0x30,
	0xC3, 0xFB, 0x03, 0x00, 0xC1, 0x83, 0x07, 0xF0, 0x28, 0xDB, 0xFB, 0x62,
	0x80, 0x10, 0x70, 0xE3, 0xE7, 0xDF, 0xFF, 0x80, 0xFF, 0xFD, 0xF3, 0xE3,
	0x87, 0x04, 0x00, 0x6F, 0xFF, 0x66, 0x60, 0x66, 0xCF, 0x3C, 0xD2, 0x6C,
	0xDB, 0xFB, 0x66, 0xCD, 0xBF, 0xB6, 0x6C, 0x18, 0x31, 0xF6, 0x3C, 0x38,
	0x1F, 0x03, 0x07, 0x0F, 0x1B, 0xE1, 0x83, 0x00, 0xC3, 0x8C, 0x30, 0xC3,
	0x0C, 0x31, 0xC3, 0x38, 0xD9, 0xB1, 0xC7, 0x7B, 0xB3, 0x66, 0xCC, 0xEC,
	0x6D, 0xE0, 0x36, 0xCC, 0xCC, 0xCC, 0x63, 0xC6, 0x33, 0x33, 0x33, 0x6C,
	0x66, 0x3C, 0xFF, 0x3C, 0x66, 0x30, 0xCF, 0xCC, 0x30, 0x6D, 0xE0, 0xFE,
	0xF0, 0x02, 0x0C, 0x30, 0xC3, 0x0C, 0x30, 0x40, 0x38, 0xDB, 0x1E, 0x3D,
	0x7A, 0xF1, 0xE3, 0x6C, 0x70, 0x31, 0xCF, 0x0C, 0x30, 0xC3, 0x0C, 0x33,
	0xF0, 0x7D, 0x8C, 0x18, 0x61, 0x86, 0x18, 0x60, 0xC7, 0xFC, 0x7D, 0x8C,
	0x18, 0x33, 0xC0, 0xC1, 0x83, 0xC6, 0xF8, 0x0C, 0x38, 0xF3, 0x6C, 0xDF,
	0xC3, 0x06, 0x0C, 0x3C, 0xFF, 0x83, 0x06, 0x0F, 0xC0, 0xC1, 0x83, 0xC6,
	0xF8, 0x38, 0xC3, 0x06, 0x0F, 0xD8, 0xF1, 0xE3, 0xC6, 0xF8, 0xFF, 0x8C,
	0x18, 0x30, 0xC3, 0x0C, 0x18, 0x30, 0x60, 0x7D, 0x8F, 0x1E, 0x37, 0xD8,
	0xF1, 0xE3, 0xC6, 0xF8, 0x7D, 0x8F, 0x1E, 0x37, 0xE0, 0xC1, 0x83, 0x0C,
	0xF0, 0xF0, 0x3C, 0x6C, 0x00, 0xDE, 0x0C, 0x63, 0x18, 0xC1, 0x83, 0x06,
	0x0C, 0xFC, 0x00, 0x3F, 0xC1, 0x83, 0x06, 0x0C, 0x63, 0x18, 0xC0, 0x7D,
	0x8F, 0x18, 0x61, 0x83, 0x06, 0x00, 0x18, 0x30, 0x7D, 0x8F, 0x1E, 0xFD,
	0xFB, 0xF7, 0x60, 0x7C, 0x10, 0x71, 0xB6, 0x3C, 0x7F, 0xF1, 0xE3, 0xC7,
	0x8C, 0xFC, 0xCD, 0x9B, 0x37, 0xCC, 0xD9, 0xB3, 0x67, 0xF8, 0x3C, 0xCF,
	0x0E, 0x0C, 0x18, 0x30, 0x61, 0x66, 0x78, 0xF8, 0xD9, 0x9B, 0x36, 0x6C,
	0xD9, 0xB3, 0x6D, 0xF0, 0xFE, 0xCD, 0x8B, 0x47, 0x8D, 0x18, 0x31, 0x67,
	0xFC, 0xFE, 0xCD, 0x8B, 0x47, 0x8D, 0x18, 0x30, 0x61, 0xE0, 0x3C, 0xCF,
	0x0E, 0x0C, 0x1B, 0xF1, 0xE3, 0x66, 0x74, 0xC7, 0x8F, 0x1E, 0x3F, 0xF8,
	0xF1, 0xE3, 0xC7, 0x8C, 0xF6, 0x66, 0x66, 0x66, 0x6F, 0x1E, 0x18, 0x30,
	0x60, 0xC1, 0xB3, 0x66, 0xCC, 0xF0, 0xE6, 0xCD, 0x9B, 0x67, 0x8F, 0x1B,
	0x33, 0x67, 0xCC, 0xF0, 0xC1, 0x83, 0x06, 0x0C, 0x18, 0x31, 0x67, 0xFC,
	0xC7, 0xDF, 0xFF, 0xFD, 0x78, 0xF1, 0xE3, 0xC7, 0x8C, 0xC7, 0xCF, 0xDF,
	0xFD, 0xF9, 0xF1, 0xE3, 0xC7, 0x8C, 0x7D, 0x8F, 0x1E, 0x3C, 0x78, 0xF1,
	0xE3, 0xC6, 0xF8, 0xFC, 0xCD, 0x9B, 0x37, 0xCC, 0x18, 0x30, 0x61, 0xE0,
	0x7D, 0x8F, 0x1E, 0x3C, 0x78, 0xF1, 0xEB, 0xDE, 0xF8, 0x30, 0x70, 0xFC,
	0xCD, 0x9B, 0x37, 0xCD, 0x99, 0xB3, 0x67, 0xCC, 0x7D, 0x8F, 0x1B, 0x03,
	0x81, 0x81, 0xE3, 0xC6, 0xF8, 0xFF, 0xFB, 0x4C, 0x30, 0xC3, 0x0C, 0x31,
	0xE0, 0xC7, 0x8F, 0x1E, 0x3C, 0x78, 0xF1, 0xE3, 0xC6, 0xF8, 0xC7, 0x8F,
	0x1E, 0x3C, 0x78, 0xF1, 0xB6, 0x38, 0x20, 0xC7, 0x8F, 0x1E, 0x3D, 0x7A,
	0xF5, 0xFF, 0xEE, 0xD8, 0xC7, 0x8D, 0xB3, 0xE3, 0x87, 0x1F, 0x36, 0xC7,
	0x8C, 0xCF, 0x3C, 0xF3, 0x78, 0xC3, 0x0C, 0x31, 0xE0, 0xFF, 0x8E, 0x18,
	0x61, 0x86, 0x18, 0x61, 0xC7, 0xFC, 0xFC, 0xCC, 0xCC, 0xCC, 0xCF, 0x81,
	0x83, 0x83, 0x83, 0x83, 0x83, 0x83, 0x02, 0xF3, 0x33, 0x33, 0x33, 0x3F,
	0x10, 0x71, 0xB6, 0x30, 0xFF, 0xD9, 0x80, 0x78, 0x19, 0xF6, 0x6C, 0xD9,
	0x9D, 0x80, 0xE0, 0xC1, 0x83, 0xC6, 0xCC, 0xD9, 0xB3, 0x66, 0xF8, 0x7D,
	0x8F, 0x06, 0x0C, 0x18, 0xDF, 0x00, 0x1C, 0x18, 0x31, 0xE6, 0xD9, 0xB3,
	0x66, 0xCC, 0xEC, 0x7D, 0x8F, 0xFE, 0x0C, 0x18, 0xDF, 0x00, 0x39, 0xB6,
	0x58, 0xF1, 0x86, 0x18, 0x63, 0xC0, 0x77, 0x9B, 0x36, 0x6C, 0xD9, 0x9F,
	0x06, 0xCC, 0xF0, 0xE0, 0xC1, 0x83, 0x67, 0x6C, 0xD9, 0xB3, 0x67, 0xCC,
	0x66, 0x0E, 0x66, 0x66, 0x6F, 0x0C, 0x30, 0x07, 0x0C, 0x30, 0xC3, 0x0C,
	0x3C, 0xF3, 0x78, 0xE0, 0xC1, 0x83, 0x36, 0xCF, 0x1E, 0x36, 0x67, 0xCC,
	0xE6, 0x66, 0x66, 0x66, 0x6F, 0xED, 0xFF, 0x5E, 0xBD, 0x7A, 0xF1, 0x80,
	0xDC, 0xCD, 0x9B, 0x36, 0x6C, 0xD9, 0x80, 0x7D, 0x8F, 0x1E, 0x3C, 0x78,
	0xDF, 0x00, 0xDC, 0xCD, 0x9B, 0x36, 0x6C, 0xDF, 0x30, 0x61, 0xE0, 0x77,
	0x9B, 0x36, 0x6C, 0xD9, 0x9F, 0x06, 0x0C, 0x3C, 0xDC, 0xED, 0x9B, 0x06,
	0x0C, 0x3C, 0x00, 0x7D, 0x8D, 0x81, 0xC0, 0xD8, 0xDF, 0x00, 0x10, 0x60,
	0xC7, 0xE3, 0x06, 0x0C, 0x18, 0x36, 0x38, 0xCD, 0x9B, 0x36, 0x6C, 0xD9,
	0x9D, 0x80, 0xCF, 0x3C, 0xF3, 0xCD, 0xE3, 0x00, 0xC7, 0x8F, 0x5E, 0xBD,
	0x7F, 0xDB, 0x00, 0xC6, 0xD8, 0xE1, 0xC3, 0x8D, 0xB1, 0x80, 0xC7, 0x8F,
	0x1E, 0x3C, 0x78, 0xDF, 0x83, 0x0D, 0xF0, 0xFF, 0x98, 0x61, 0x86, 0x18,
	0xFF, 0x80, 0x1C, 0xC3, 0x0C, 0xE0, 0xC3, 0x0C, 0x30, 0x70, 0xFF, 0x3F,
	0xF0, 0xE0, 0xC3, 0x0C, 0x1C, 0xC3, 0x0C, 0x33, 0x80, 0x77, 0xB8, 0x10,
	0x71, 0xB6, 0x3C, 0x78, 0xFF, 0x80, 0xFC, 0x63, 0x18, 0xC6, 0x31, 0x8C,
	0x7E, 0xFC, 0x63, 0x18, 0xC6, 0x31, 0x8C, 0x7E, 0xFC, 0x63, 0x18, 0xC6,
	0x31, 0x8C, 0x7E, 0xFC, 0x63, 0x18, 0xC6, 0x31, 0x8C, 0x7E, 0xFC, 0x63,
	0x18, 0xC6, 0x31, 0x8C, 0x7E, 0xFC, 0x63, 0x18, 0xC6, 0x31, 0x8C, 0x7E,
	0xFC, 0x63, 0x18, 0xC6, 0x31, 0x8C, 0x7E, 0xFC, 0x63, 0x18, 0xC6, 0x31,
	0x8C, 0x7E, 0xFC, 0x63, 0x18, 0xC6, 0x31, 0x8C, 0x7E, 0xFC, 0x63, 0x18,
	0xC6, 0x31, 0x8C, 0x7E, 0xFC, 0x63, 0x18, 0xC6, 0x31, 0x8C, 0x7E, 0xFC,
	0x63, 0x18, 0xC6, 0x31, 0x8C, 0x7E, 0xFC, 0x63, 0x18, 0xC6, 0x31, 0x8C,
	0x7E, 0xFC, 0x63, 0x18, 0xC6, 0x31, 0x8C, 0x7E, 0xFC, 0x63, 0x18, 0xC6,
	0x31, 0x8C, 0x7E, 0xFC, 0x63, 0x18, 0xC6, 0x31, 0x8C, 0x7E, 0xFC, 0x63,
	0x18, 0xC6, 0x31, 0x8C, 0x7E, 0xFC, 0x63, 0x18, 0xC6, 0x31, 0x8C, 0x7E,
	0xFC, 0x63, 0x18, 0xC6, 0x31, 0x8C, 0x7E, 0xFC, 0x63, 0x18, 0xC6, 0x31,
	0x8C, 0x7E, 0xFC, 0x63, 0x18, 0xC6, 0x31, 0x8C, 0x7E, 0xFC, 0x63, 0x18,
	0xC6, 0x31, 0x8C, 0x7E, 0xFC, 0x63, 0x18, 0xC6, 0x31, 0x8C, 0x7E, 0xFC,
	0x63, 0x18, 0xC6, 0x31, 0x8C, 0x7E, 0xFC, 0x63, 0x18, 0xC6, 0x31, 0x8C,
	0x7E, 0xFC, 0x63, 0x18, 0xC6, 0x31, 0x8C, 0x7E, 0xFC, 0x63, 0x18, 0xC6,
	0x31, 0x8C, 0x7E, 0xFC, 0x63, 0x18, 0xC6, 0x31, 0x8C, 0x7E, 0xFC, 0x63,
	0x18, 0xC6, 0x31, 0x8C, 0x7E, 0xFC, 0x63, 0x18, 0xC6, 0x31, 0x8C, 0x7E,
	0xFC, 0x63, 0x18, 0xC6, 0x31, 0x8C, 0x7E, 0xFC, 0x63, 0x18, 0xC6, 0x31,
	0x8C, 0x7E, 0x66, 0x06, 0x66, 0xFF, 0xF6, 0x30, 0xC7, 0xB3, 0xC3, 0x0C,
	0x33, 0x78, 0xC3, 0x00, 0x38, 0xD9, 0x93, 0x0F, 0x0C, 0x18, 0x30, 0x61,
	0xCF, 0xF0, 0xFC, 0x63, 0x18, 0xC6, 0x31, 0x8C, 0x7E, 0xCF, 0x37, 0x8C,
	0xFC, 0xCF, 0xCC, 0x30, 0xC0, 0xFC, 0x63, 0x18, 0xC6, 0x31, 0x8C, 0x7E,
	0x7D, 0x8D, 0x81, 0xC6, 0xD8, 0xF1, 0xB6, 0x38, 0x1B, 0x1B, 0xE0, 0xFC,
	0x63, 0x18, 0xC6, 0x31, 0x8C, 0x7E, 0xFC, 0x63, 0x18, 0xC6, 0x31, 0x8C,
	0x7E, 0x7B, 0x6D, 0x9F, 0x03, 0xF0, 0x36, 0xDB, 0x63, 0x63, 0x60, 0xFE,
	0x0C, 0x18, 0x30, 0x60, 0xFC, 0x63, 0x18, 0xC6, 0x31, 0x8C, 0x7E, 0xFC,
	0x63, 0x18, 0xC6, 0x31, 0x8C, 0x7E, 0xFC, 0x63, 0x18, 0xC6, 0x31, 0x8C,
	0x7E, 0x76, 0xF6, 0xE0, 0x18, 0x18, 0x7E, 0x18, 0x18, 0x00, 0x00, 0xFF,
	0x76, 0xCC, 0xCC, 0xFC, 0xFC, 0x63, 0x18, 0xC6, 0x31, 0x8C, 0x7E, 0xFC,
	0x63, 0x18, 0xC6, 0x31, 0x8C, 0x7E, 0x66, 0xCD, 0x9B, 0x36, 0x6F, 0x98,
	0x30, 0xC0, 0x7F, 0xDB, 0xDB, 0xDB, 0x7B, 0x1B, 0x1B, 0x1B, 0x1B, 0x1B,
	0xC0, 0xFC, 0x63, 0x18, 0xC6, 0x31, 0x8C, 0x7E, 0xFC, 0x63, 0x18, 0xC6,
	0x31, 0x8C, 0x7E, 0x76, 0xF6, 0xE0, 0x7C, 0xD8, 0xD8, 0xDB, 0x6D, 0x80,
	0xC1, 0x83, 0x0E, 0x3C, 0xC3, 0x0C, 0x33, 0xCF, 0x3C, 0xF8, 0x30, 0x60,
	0xC1, 0x83, 0x0E, 0x3C, 0xC3, 0x0C, 0x30, 0xDD, 0x0C, 0x30, 0xC3, 0xE0,
	0xFC, 0x63, 0x18, 0xC6, 0x31, 0x8C, 0x7E, 0x30, 0x60, 0x01, 0x83, 0x0C,
	0x30, 0x63, 0xC6, 0xF8, 0xFC, 0x63, 0x18, 0xC6, 0x31, 0x8C, 0x7E, 0xFC,
	0x63, 0x18, 0xC6, 0x31, 0x8C, 0x7E, 0xFC, 0x63, 0x18, 0xC6, 0x31, 0x8C,
	0x7E, 0xFC, 0x63, 0x18, 0xC6, 0x31, 0x8C, 0x7E, 0xC6, 0x00, 0x41, 0xC6,
	0xD8, 0xF1, 0xFF, 0xC7, 0x8F, 0x18, 0x38, 0xD8, 0xE0, 0x03, 0x8D, 0xB1,
	0xE3, 0xFF, 0x8F, 0x1E, 0x30, 0x3E, 0xDB, 0x36, 0x6F, 0xF9, 0xB3, 0x66,
	0xCD, 0x9C, 0x3C, 0xCF, 0x0E, 0x0C, 0x18, 0x30, 0xB3, 0x3C, 0x18, 0x1B,
	0xE0, 0xFC, 0x63, 0x18, 0xC6, 0x31, 0x8C, 0x7E, 0x18, 0x61, 0x80, 0x0F,
	0xEC, 0xD8, 0x3E, 0x60, 0xC1, 0x9F, 0xF0, 0xFC, 0x63, 0x18, 0xC6, 0x31,
	0x8C, 0x7E, 0xFC, 0x63, 0x18, 0xC6, 0x31, 0x8C, 0x7E, 0xFC, 0x63, 0x18,
	0xC6, 0x31, 0x8C, 0x7E, 0xFC, 0x63, 0x18, 0xC6, 0x31, 0x8C, 0x7E, 0xFC,
	0x63, 0x18, 0xC6, 0x31, 0x8C, 0x7E, 0xFC, 0x63, 0x18, 0xC6, 0x31, 0x8C,
	0x7E, 0xFC, 0x63, 0x18, 0xC6, 0x31, 0x8C, 0x7E, 0x77, 0xB8, 0x06, 0x3E,
	0x7E, 0xFF, 0xEF, 0xCF, 0x8F, 0x1E, 0x30, 0xFC, 0x63, 0x18, 0xC6, 0x31,
	0x8C, 0x7E, 0xFC, 0x63, 0x18, 0xC6, 0x31, 0x8C, 0x7E, 0xFC, 0x63, 0x18,
	0xC6, 0x31, 0x8C, 0x7E, 0xFC, 0x63, 0x18, 0xC6, 0x31, 0x8C, 0x7E, 0xC6,
	0x01, 0xF6, 0x3C, 0x78, 0xF1, 0xE3, 0xC7, 0x8D, 0xF0, 0xFC, 0x63, 0x18,
	0xC6, 0x31, 0x8C, 0x7E, 0xFC, 0x63, 0x18, 0xC6, 0x31, 0x8C, 0x7E, 0xFC,
	0x63, 0x18, 0xC6, 0x31, 0x8C, 0x7E, 0xFC, 0x63, 0x18, 0xC6, 0x31, 0x8C,
	0x7E, 0xFC, 0x63, 0x18, 0xC6, 0x31, 0x8C, 0x7E, 0xC6, 0x03, 0x1E, 0x3C,
	0x78, 0xF1, 0xE3, 0xC7, 0x8D, 0xF0, 0xFC, 0x63, 0x18, 0xC6, 0x31, 0x8C,
	0x7E, 0xFC, 0x63, 0x18, 0xC6, 0x31, 0x8C, 0x7E, 0x79, 0x9B, 0x36, 0x6D,
	0x99, 0xB1, 0xE3, 0xC7, 0x98, 0x60, 0x60, 0x60, 0x07, 0x81, 0x9F, 0x66,
	0xCD, 0x99, 0xD8, 0x18, 0x61, 0x80, 0x07, 0x81, 0x9F, 0x66, 0xCD, 0x99,
	0xD8, 0x10, 0x71, 0xB0, 0x07, 0x81, 0x9F, 0x66, 0xCD, 0x99, 0xD8, 0xFC,
	0x63, 0x18, 0xC6, 0x31, 0x8C, 0x7E, 0xCC, 0x00, 0x03, 0xC0, 0xCF, 0xB3,
	0x66, 0xCC, 0xEC, 0x38, 0xD8, 0xE0, 0x07, 0x81, 0x9F, 0x66, 0xCD, 0x99,
	0xD8, 0xCC, 0xEC, 0xDB, 0xFD, 0x9B, 0x1B, 0x80, 0x7B, 0x3C, 0x30, 0xCD,
	0xE1, 0x83, 0x78, 0x60, 0x60, 0x60, 0x07, 0xD8, 0xFF, 0xE0, 0xC1, 0x8D,
	0xF0, 0x0C, 0x30, 0xC0, 0x07, 0xD8, 0xFF, 0xE0, 0xC1, 0x8D, 0xF0, 0x10,
	0x71, 0xB0, 0x07, 0xD8, 0xFF, 0xE0, 0xC1, 0x8D, 0xF0, 0xC6, 0x00, 0x03,
	0xEC, 0x7F, 0xF0, 0x60, 0xC6, 0xF8, 0xC3, 0x0C, 0x07, 0x18, 0xC6, 0x31,
	0x9E, 0x36, 0xC0, 0xE6, 0x66, 0x66, 0xF0, 0x31, 0xEC, 0xC0, 0x70, 0xC3,
	0x0C, 0x30, 0xC7, 0x80, 0xCC, 0x00, 0x1C, 0x30, 0xC3, 0x0C, 0x31, 0xE0,
	0xFC, 0x63, 0x18, 0xC6, 0x31, 0x8C, 0x7E, 0x77, 0xB8, 0x06, 0xE6, 0x6C,
	0xD9, 0xB3, 0x66, 0xCC, 0x60, 0x60, 0x60, 0x07, 0xD8, 0xF1, 0xE3, 0xC7,
	0x8D, 0xF0, 0x18, 0x61, 0x80, 0x07, 0xD8, 0xF1, 0xE3, 0xC7, 0x8D, 0xF0,
	0x10, 0x71, 0xB0, 0x07, 0xD8, 0xF1, 0xE3, 0xC7, 0x8D, 0xF0, 0xFC, 0x63,
	0x18, 0xC6, 0x31, 0x8C, 0x7E, 0xC6, 0x00, 0x03, 0xEC, 0x78, 0xF1, 0xE3,
	0xC6, 0xF8, 0x30, 0xC0, 0x3F, 0x00, 0xC3, 0x00, 0xFC, 0x63, 0x18, 0xC6,
	0x31, 0x8C, 0x7E, 0x60, 0x60, 0x60, 0x0C, 0xD9, 0xB3, 0x66, 0xCD, 0x99,
	0xD8, 0x18, 0x61, 0x80, 0x0C, 0xD9, 0xB3, 0x66, 0xCD, 0x99, 0xD8, 0x30,
	0xF3, 0x30, 0x0C, 0xD9, 0xB3, 0x66, 0xCD, 0x99, 0xD8, 0xCC, 0x00, 0x06,
	0x6C, 0xD9, 0xB3, 0x66, 0xCC, 0xEC, 0xFC, 0x63, 0x18, 0xC6, 0x31, 0x8C,
	0x7E, 0xFC, 0x63, 0x18, 0xC6, 0x31, 0x8C, 0x7E, 0xC6, 0x00, 0x06, 0x3C,
	0x78, 0xF1, 0xE3, 0xC6, 0xFC, 0x18, 0x67, 0x80}

var Regular8pt8bGlyphs = []tinyfont.Glyph{
	{0, 8, 10, 8, 0, -9},     // 0x01
	{10, 8, 10, 8, 0, -9},    // 0x02
	{20, 7, 8, 8, 0, -7},     // 0x03
	{27, 7, 7, 8, 0, -7},     // 0x04
	{34, 8, 9, 8, 0, -8},     // 0x05
	{43, 8, 9, 8, 0, -8},     // 0x06
	{52, 4, 4, 8, 2, -5},     // 0x07
	{54, 8, 16, 8, 0, -11},   // 0x08
	{70, 6, 6, 8, 1, -6},     // 0x09
	{75, 8, 16, 8, 0, -11},   // 0x0A
	{91, 7, 10, 8, 0, -9},    // 0x0B
	{100, 6, 10, 8, 1, -9},   // 0x0C
	{108, 8, 10, 8, 0, -9},   // 0x0D
	{118, 8, 11, 8, 0, -9},   // 0x0E
	{129, 8, 9, 8, 0, -8},    // 0x0F
	{138, 7, 11, 8, 0, -10},  // 0x10
	{148, 7, 11, 8, 0, -10},  // 0x11
	{158, 6, 9, 8, 1, -9},    // 0x12
	{165, 6, 10, 8, 1, -9},   // 0x13
	{173, 8, 10, 8, 0, -9},   // 0x14
	{183, 7, 12, 8, 0, -10},  // 0x15
	{194, 7, 4, 8, 0, -3},    // 0x16
	{198, 6, 10, 8, 1, -9},   // 0x17
	{206, 6, 10, 8, 1, -9},   // 0x18
	{214, 6, 10, 8, 1, -9},   // 0x19
	{222, 7, 5, 8, 0, -6},    // 0x1A
	{227, 7, 5, 8, 0, -6},    // 0x1B
	{232, 7, 4, 8, 0, -5},    // 0x1C
	{236, 7, 5, 8, 0, -6},    // 0x1D
	{241, 7, 7, 8, 0, -7},    // 0x1E
	{248, 7, 7, 8, 0, -7},    // 0x1F
	{255, 0, 0, 8, 0, 1},     // 0x20 ' '
	{255, 4, 10, 8, 2, -9},   // 0x21 '!'
	{260, 6, 4, 8, 1, -10},   // 0x22 '"'
	{263, 7, 9, 8, 0, -8},    // 0x23 '#'
	{271, 7, 14, 8, 0, -11},  // 0x24 '$'
	{284, 7, 8, 8, 0, -7},    // 0x25 '%'
	{291, 7, 10, 8, 0, -9},   // 0x26 '&'
	{300, 3, 4, 8, 1, -10},   // 0x27 '''
	{302, 4, 10, 8, 2, -9},   // 0x28 '('
	{307, 4, 10, 8, 2, -9},   // 0x29 ')'
	{312, 8, 5, 8, 0, -6},    // 0x2A '*'
	{317, 6, 5, 8, 1, -6},    // 0x2B '+'
	{321, 3, 4, 8, 2, -2},    // 0x2C ','
	{323, 7, 1, 8, 0, -4},    // 0x2D '-'
	{324, 2, 2, 8, 3, -1},    // 0x2E '.'
	{325, 7, 8, 8, 0, -7},    // 0x2F '/'
	{332, 7, 10, 8, 0, -9},   // 0x30 '0'
	{341, 6, 10, 8, 1, -9},   // 0x31 '1'
	{349, 7, 10, 8, 0, -9},   // 0x32 '2'
	{358, 7, 10, 8, 0, -9},   // 0x33 '3'
	{367, 7, 10, 8, 0, -9},   // 0x34 '4'
	{376, 7, 10, 8, 0, -9},   // 0x35 '5'
	{385, 7, 10, 8, 0, -9},   // 0x36 '6'
	{394, 7, 10, 8, 0, -9},   // 0x37 '7'
	{403, 7, 10, 8, 0, -9},   // 0x38 '8'
	{412, 7, 10, 8, 0, -9},   // 0x39 '9'
	{421, 2, 7, 8, 3, -7},    // 0x3A ':'
	{423, 3, 8, 8, 2, -7},    // 0x3B ';'
	{426, 6, 9, 8, 1, -8},    // 0x3C '<'
	{433, 6, 4, 8, 1, -6},    // 0x3D '='
	{436, 6, 9, 8, 1, -8},    // 0x3E '>'
	{443, 7, 10, 8, 0, -9},   // 0x3F '?'
	{452, 7, 9, 8, 0, -8},    // 0x40 '@'
	{460, 7, 10, 8, 0, -9},   // 0x41 'A'
	{469, 7, 10, 8, 0, -9},   // 0x42 'B'
	{478, 7, 10, 8, 0, -9},   // 0x43 'C'
	{487, 7, 10, 8, 0, -9},   // 0x44 'D'
	{496, 7, 10, 8, 0, -9},   // 0x45 'E'
	{505, 7, 10, 8, 0, -9},   // 0x46 'F'
	{514, 7, 10, 8, 0, -9},   // 0x47 'G'
	{523, 7, 10, 8, 0, -9},   // 0x48 'H'
	{532, 4, 10, 8, 2, -9},   // 0x49 'I'
	{537, 7, 10, 8, 0, -9},   // 0x4A 'J'
	{546, 7, 10, 8, 0, -9},   // 0x4B 'K'
	{555, 7, 10, 8, 0, -9},   // 0x4C 'L'
	{564, 7, 10, 8, 0, -9},   // 0x4D 'M'
	{573, 7, 10, 8, 0, -9},   // 0x4E 'N'
	{582, 7, 10, 8, 0, -9},   // 0x4F 'O'
	{591, 7, 10, 8, 0, -9},   // 0x50 'P'
	{600, 7, 12, 8, 0, -9},   // 0x51 'Q'
	{611, 7, 10, 8, 0, -9},   // 0x52 'R'
	{620, 7, 10, 8, 0, -9},   // 0x53 'S'
	{629, 6, 10, 8, 1, -9},   // 0x54 'T'
	{637, 7, 10, 8, 0, -9},   // 0x55 'U'
	{646, 7, 10, 8, 0, -9},   // 0x56 'V'
	{655, 7, 10, 8, 0, -9},   // 0x57 'W'
	{664, 7, 10, 8, 0, -9},   // 0x58 'X'
	{673, 6, 10, 8, 1, -9},   // 0x59 'Y'
	{681, 7, 10, 8, 0, -9},   // 0x5A 'Z'
	{690, 4, 10, 8, 2, -9},   // 0x5B '['
	{695, 7, 9, 8, 0, -8},    // 0x5C '\'
	{703, 4, 10, 8, 2, -9},   // 0x5D ']'
	{708, 7, 4, 8, 0, -11},   // 0x5E '^'
	{712, 8, 1, 8, 0, 2},     // 0x5F '_'
	{713, 3, 3, 8, 2, -11},   // 0x60 '`'
	{715, 7, 7, 8, 0, -6},    // 0x61 'a'
	{722, 7, 10, 8, 0, -9},   // 0x62 'b'
	{731, 7, 7, 8, 0, -6},    // 0x63 'c'
	{738, 7, 10, 8, 0, -9},   // 0x64 'd'
	{747, 7, 7, 8, 0, -6},    // 0x65 'e'
	{754, 6, 10, 8, 0, -9},   // 0x66 'f'
	{762, 7, 10, 8, 0, -6},   // 0x67 'g'
	{771, 7, 10, 8, 0, -9},   // 0x68 'h'
	{780, 4, 10, 8, 2, -9},   // 0x69 'i'
	{785, 6, 13, 8, 1, -9},   // 0x6A 'j'
	{795, 7, 10, 8, 0, -9},   // 0x6B 'k'
	{804, 4, 10, 8, 2, -9},   // 0x6C 'l'
	{809, 7, 7, 8, 0, -6},    // 0x6D 'm'
	{816, 7, 7, 8, 0, -6},    // 0x6E 'n'
	{823, 7, 7, 8, 0, -6},    // 0x6F 'o'
	{830, 7, 10, 8, 0, -6},   // 0x70 'p'
	{839, 7, 10, 8, 0, -6},   // 0x71 'q'
	{848, 7, 7, 8, 0, -6},    // 0x72 'r'
	{855, 7, 7, 8, 0, -6},    // 0x73 's'
	{862, 7, 10, 8, 0, -9},   // 0x74 't'
	{871, 7, 7, 8, 0, -6},    // 0x75 'u'
	{878, 6, 7, 8, 1, -6},    // 0x76 'v'
	{884, 7, 7, 8, 0, -6},    // 0x77 'w'
	{891, 7, 7, 8, 0, -6},    // 0x78 'x'
	{898, 7, 10, 8, 0, -6},   // 0x79 'y'
	{907, 7, 7, 8, 0, -6},    // 0x7A 'z'
	{914, 6, 10, 8, 1, -9},   // 0x7B '{'
	{922, 2, 10, 8, 3, -9},   // 0x7C '|'
	{925, 6, 10, 8, 1, -9},   // 0x7D '}'
	{933, 7, 2, 8, 0, -9},    // 0x7E '~'
	{935, 7, 7, 8, 0, -7},    // 0x7F
	{942, 5, 11, 8, 1, -10},  // 0x80
	{949, 5, 11, 8, 1, -10},  // 0x81
	{956, 5, 11, 8, 1, -10},  // 0x82
	{963, 5, 11, 8, 1, -10},  // 0x83
	{970, 5, 11, 8, 1, -10},  // 0x84
	{977, 5, 11, 8, 1, -10},  // 0x85
	{984, 5, 11, 8, 1, -10},  // 0x86
	{991, 5, 11, 8, 1, -10},  // 0x87
	{998, 5, 11, 8, 1, -10},  // 0x88
	{1005, 5, 11, 8, 1, -10}, // 0x89
	{1012, 5, 11, 8, 1, -10}, // 0x8A
	{1019, 5, 11, 8, 1, -10}, // 0x8B
	{1026, 5, 11, 8, 1, -10}, // 0x8C
	{1033, 5, 11, 8, 1, -10}, // 0x8D
	{1040, 5, 11, 8, 1, -10}, // 0x8E
	{1047, 5, 11, 8, 1, -10}, // 0x8F
	{1054, 5, 11, 8, 1, -10}, // 0x90
	{1061, 5, 11, 8, 1, -10}, // 0x91
	{1068, 5, 11, 8, 1, -10}, // 0x92
	{1075, 5, 11, 8, 1, -10}, // 0x93
	{1082, 5, 11, 8, 1, -10}, // 0x94
	{1089, 5, 11, 8, 1, -10}, // 0x95
	{1096, 5, 11, 8, 1, -10}, // 0x96
	{1103, 5, 11, 8, 1, -10}, // 0x97
	{1110, 5, 11, 8, 1, -10}, // 0x98
	{1117, 5, 11, 8, 1, -10}, // 0x99
	{1124, 5, 11, 8, 1, -10}, // 0x9A
	{1131, 5, 11, 8, 1, -10}, // 0x9B
	{1138, 5, 11, 8, 1, -10}, // 0x9C
	{1145, 5, 11, 8, 1, -10}, // 0x9D
	{1152, 5, 11, 8, 1, -10}, // 0x9E
	{1159, 5, 11, 8, 1, -10}, // 0x9F
	{1166, 0, 0, 8, 0, 1},    // 0xA0
	{1166, 4, 10, 8, 2, -9},  // 0xA1
	{1171, 6, 11, 8, 1, -10}, // 0xA2
	{1180, 7, 11, 8, 0, -10}, // 0xA3
	{1190, 5, 11, 8, 1, -10}, // 0xA4
	{1197, 6, 10, 8, 1, -9},  // 0xA5
	{1205, 5, 11, 8, 1, -10}, // 0xA6
	{1212, 7, 12, 8, 0, -10}, // 0xA7
	{1223, 5, 11, 8, 1, -10}, // 0xA8
	{1230, 5, 11, 8, 1, -10}, // 0xA9
	{1237, 6, 6, 8, 1, -10},  // 0xAA
	{1242, 7, 5, 8, 0, -6},   // 0xAB
	{1247, 7, 5, 8, 0, -5},   // 0xAC
	{1252, 5, 11, 8, 1, -10}, // 0xAD
	{1259, 5, 11, 8, 1, -10}, // 0xAE
	{1266, 5, 11, 8, 1, -10}, // 0xAF
	{1273, 5, 4, 8, 1, -10},  // 0xB0
	{1276, 8, 8, 8, 0, -7},   // 0xB1
	{1284, 5, 6, 8, 0, -10},  // 0xB2
	{1288, 5, 11, 8, 1, -10}, // 0xB3
	{1295, 5, 11, 8, 1, -10}, // 0xB4
	{1302, 7, 9, 8, 0, -7},   // 0xB5
	{1310, 8, 10, 8, 0, -9},  // 0xB6
	{1320, 2, 1, 8, 3, -3},   // 0xB7
	{1321, 5, 11, 8, 1, -10}, // 0xB8
	{1328, 5, 11, 8, 1, -10}, // 0xB9
	{1335, 5, 6, 8, 1, -10},  // 0xBA
	{1339, 7, 5, 8, 0, -6},   // 0xBB
	{1344, 7, 13, 8, 0, -10}, // 0xBC
	{1356, 7, 13, 8, 0, -10}, // 0xBD
	{1368, 5, 11, 8, 1, -10}, // 0xBE
	{1375, 7, 10, 8, 0, -9},  // 0xBF
	{1384, 5, 11, 8, 1, -10}, // 0xC0
	{1391, 5, 11, 8, 1, -10}, // 0xC1
	{1398, 5, 11, 8, 1, -10}, // 0xC2
	{1405, 5, 11, 8, 1, -10}, // 0xC3
	{1412, 7, 11, 8, 0, -10}, // 0xC4
	{1422, 7, 12, 8, 0, -11}, // 0xC5
	{1433, 7, 10, 8, 0, -9},  // 0xC6
	{1442, 7, 12, 8, 0, -9},  // 0xC7
	{1453, 5, 11, 8, 1, -10}, // 0xC8
	{1460, 7, 12, 8, 0, -11}, // 0xC9
	{1471, 5, 11, 8, 1, -10}, // 0xCA
	{1478, 5, 11, 8, 1, -10}, // 0xCB
	{1485, 5, 11, 8, 1, -10}, // 0xCC
	{1492, 5, 11, 8, 1, -10}, // 0xCD
	{1499, 5, 11, 8, 1, -10}, // 0xCE
	{1506, 5, 11, 8, 1, -10}, // 0xCF
	{1513, 5, 11, 8, 1, -10}, // 0xD0
	{1520, 7, 12, 8, 0, -11}, // 0xD1
	{1531, 5, 11, 8, 1, -10}, // 0xD2
	{1538, 5, 11, 8, 1, -10}, // 0xD3
	{1545, 5, 11, 8, 1, -10}, // 0xD4
	{1552, 5, 11, 8, 1, -10}, // 0xD5
	{1559, 7, 11, 8, 0, -10}, // 0xD6
	{1569, 5, 11, 8, 1, -10}, // 0xD7
	{1576, 5, 11, 8, 1, -10}, // 0xD8
	{1583, 5, 11, 8, 1, -10}, // 0xD9
	{1590, 5, 11, 8, 1, -10}, // 0xDA
	{1597, 5, 11, 8, 1, -10}, // 0xDB
	{1604, 7, 11, 8, 0, -10}, // 0xDC
	{1614, 5, 11, 8, 1, -10}, // 0xDD
	{1621, 5, 11, 8, 1, -10}, // 0xDE
	{1628, 7, 10, 8, 0, -9},  // 0xDF
	{1637, 7, 11, 8, 0, -10}, // 0xE0
	{1647, 7, 11, 8, 0, -10}, // 0xE1
	{1657, 7, 11, 8, 0, -10}, // 0xE2
	{1667, 5, 11, 8, 1, -10}, // 0xE3
	{1674, 7, 10, 8, 0, -9},  // 0xE4
	{1683, 7, 11, 8, 0, -10}, // 0xE5
	{1693, 7, 7, 8, 0, -6},   // 0xE6
	{1700, 6, 9, 8, 1, -7},   // 0xE7
	{1707, 7, 11, 8, 0, -10}, // 0xE8
	{1717, 7, 11, 8, 0, -10}, // 0xE9
	{1727, 7, 11, 8, 0, -10}, // 0xEA
	{1737, 7, 10, 8, 0, -9},  // 0xEB
	{1746, 5, 11, 8, 1, -10}, // 0xEC
	{1753, 4, 11, 8, 2, -10}, // 0xED
	{1759, 6, 11, 8, 1, -10}, // 0xEE
	{1768, 6, 10, 8, 1, -9},  // 0xEF
	{1776, 5, 11, 8, 1, -10}, // 0xF0
	{1783, 7, 10, 8, 0, -9},  // 0xF1
	{1792, 7, 11, 8, 0, -10}, // 0xF2
	{1802, 7, 11, 8, 0, -10}, // 0xF3
	{1812, 7, 11, 8, 0, -10}, // 0xF4
	{1822, 5, 11, 8, 1, -10}, // 0xF5
	{1829, 7, 10, 8, 0, -9},  // 0xF6
	{1838, 6, 7, 8, 1, -7},   // 0xF7
	{1844, 5, 11, 8, 1, -10}, // 0xF8
	{1851, 7, 11, 8, 0, -10}, // 0xF9
	{1861, 7, 11, 8, 0, -10}, // 0xFA
	{1871, 7, 11, 8, 0, -10}, // 0xFB
	{1881, 7, 10, 8, 0, -9},  // 0xFC
	{1890, 5, 11, 8, 1, -10}, // 0xFD
	{1897, 5, 11, 8, 1, -10}, // 0xFE
	{1904, 7, 13, 8, 0, -9}}  // 0xFF

var Regular8pt8b = tinyfont.Font{
	Regular8pt8bBitmaps,
	Regular8pt8bGlyphs,
	0x01, 0xFF, 16}

// Approx. 3708 bytes
