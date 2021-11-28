package main

type Buffer struct {
	WordA uint32
	WordB uint32
	WordC uint32
	WordD uint32
}

func (buffer Buffer) toByteArray() []byte{
	result := []byte{
		byte(buffer.WordA),
		byte(buffer.WordA >> 8),
		byte(buffer.WordA >> 16),
		byte(buffer.WordA >> 24),
		byte(buffer.WordB),
		byte(buffer.WordB >> 8),
		byte(buffer.WordB >> 16),
		byte(buffer.WordB >> 24),
		byte(buffer.WordC),
		byte(buffer.WordC >> 8),
		byte(buffer.WordC >> 16),
		byte(buffer.WordC >> 24),
		byte(buffer.WordD),
		byte(buffer.WordD >> 8),
		byte(buffer.WordD >> 16),
		byte(buffer.WordD >> 24),
	}

	return result
}

