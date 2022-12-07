package main

/*
#include <stdint.h>
#include <stdlib.h>

typedef struct {
   char*  definition;
   char*  permalink;
   int    thumbs_up;
   char** sound_urls;
   size_t sound_urls_len;
   char*  author;
   char*  word;
   int    defid;
   char*  written_on;
   char*  example;
   int    thumbs_down;
} urban_dictionary_response;

typedef struct {
   urban_dictionary_response* response;
   char* error_msg;
} urban_dictionary_response_t;
*/
import "C"
import (
	"unsafe"
	ub "github.com/thechampagne/urbandictionary-go/urbandictionary"
)

//export urban_dictionary_definition_by_id
func urban_dictionary_definition_by_id(id C.int64_t) C.urban_dictionary_response_t {
	var self C.urban_dictionary_response_t
	res, err := ub.DefinitionById(int64(id))
	if err != nil {
		self.error_msg = C.CString(err.Error())
		return self
	}
	response := (*C.urban_dictionary_response) (C.malloc(C.size_t(unsafe.Sizeof(C.urban_dictionary_response{}))))
	response.definition =  C.CString(res.Definition)
	response.permalink =  C.CString(res.Permalink)
	response.thumbs_up =  C.int(res.ThumbsUp)
	sound_urls := C.malloc(C.size_t(len(res.SoundUrls)) * C.size_t(unsafe.Sizeof(uintptr(0))))
	sound_urls_slice := (*[1<<30 - 1]*C.char)(sound_urls)

	for i, v := range res.SoundUrls {
		sound_urls_slice[i] = C.CString(v)
	}
	response.sound_urls = (**C.char) (sound_urls)
	response.sound_urls_len =  C.size_t(len(res.SoundUrls))
	response.author =  C.CString(res.Author)
	response.word =  C.CString(res.Word)
	response.defid =  C.int(res.Defid)
	response.written_on =  C.CString(res.WrittenOn)
	response.example =  C.CString(res.Example)
	response.thumbs_down =  C.int(res.ThumbsDown)
	self.response = response
	return self
}

//export urban_dictionary_random
func urban_dictionary_random(array_len *C.size_t) C.urban_dictionary_response_t {
	var self C.urban_dictionary_response_t
	res, err := ub.Random()
	if err != nil {
		self.error_msg = C.CString(err.Error())
		return self
	}
	array := C.malloc(C.size_t(len(res)) * C.size_t(unsafe.Sizeof(uintptr(0))))

	slice := (*[1<<30 - 1]*C.urban_dictionary_response)(array)

	for i, v := range res {
		response := (*C.urban_dictionary_response) (C.malloc(C.size_t(unsafe.Sizeof(C.urban_dictionary_response{}))))
		response.definition =  C.CString(v.Definition)
		response.permalink =  C.CString(v.Permalink)
		response.thumbs_up =  C.int(v.ThumbsUp)

		sound_urls := C.malloc(C.size_t(len(v.SoundUrls)) * C.size_t(unsafe.Sizeof(uintptr(0))))
		sound_urls_slice := (*[1<<30 - 1]*C.char)(sound_urls)

		for idx, val := range v.SoundUrls {
			sound_urls_slice[idx] = C.CString(val)
		}
		response.sound_urls = (**C.char) (sound_urls)
		response.sound_urls_len =  C.size_t(len(v.SoundUrls))
		response.author =  C.CString(v.Author)
		response.word =  C.CString(v.Word)
		response.defid =  C.int(v.Defid)
		response.written_on =  C.CString(v.WrittenOn)
		response.example =  C.CString(v.Example)
		response.thumbs_down =  C.int(v.ThumbsDown)
		slice[i] = response
	}
	self.response = (*C.urban_dictionary_response) (array)
	*array_len = C.size_t(len(res))
	return self
}

//export urban_dictionary_tool_tip
func urban_dictionary_tool_tip(term *C.char, is_err *C.int) *C.char {
	res, err := ub.ToolTip(C.GoString(term))
	if err != nil {
		return  C.CString(err.Error())
	}
	return C.CString(res)
}

//export urban_dictionary_data
func urban_dictionary_data(input *C.char, page C.int) C.urban_dictionary_response_t {
	var self C.urban_dictionary_response_t
	urban := ub.New(C.GoString(input), int32(page))
	res, err := urban.Data()
	if err != nil {
		self.error_msg = C.CString(err.Error())
		return self
	}
	array := C.malloc(C.size_t(len(res)) * C.size_t(unsafe.Sizeof(uintptr(0))))

	slice := (*[1<<30 - 1]*C.urban_dictionary_response)(array)

	for i, v := range res {
		response := (*C.urban_dictionary_response) (C.malloc(C.size_t(unsafe.Sizeof(C.urban_dictionary_response{}))))
		response.definition =  C.CString(v.Definition)
		response.permalink =  C.CString(v.Permalink)
		response.thumbs_up =  C.int(v.ThumbsUp)

		sound_urls := C.malloc(C.size_t(len(v.SoundUrls)) * C.size_t(unsafe.Sizeof(uintptr(0))))
		sound_urls_slice := (*[1<<30 - 1]*C.char)(sound_urls)

		for idx, val := range v.SoundUrls {
			sound_urls_slice[idx] = C.CString(val)
		}
		response.sound_urls = (**C.char) (sound_urls)
		response.sound_urls_len =  C.size_t(len(v.SoundUrls))
		response.author =  C.CString(v.Author)
		response.word =  C.CString(v.Word)
		response.defid =  C.int(v.Defid)
		response.written_on =  C.CString(v.WrittenOn)
		response.example =  C.CString(v.Example)
		response.thumbs_down =  C.int(v.ThumbsDown)
		slice[i] = response
	}
	self.response = (*C.urban_dictionary_response) (array)
	return self
}

//export urban_dictionary_response_clean
func urban_dictionary_response_clean(self *C.urban_dictionary_response) {
	if self != nil {
		if self.definition != nil { C.free(unsafe.Pointer(self.definition)) }
		if self.permalink != nil { C.free(unsafe.Pointer(self.permalink)) }
		if self.sound_urls != nil {
			slice := (*[1<<30 - 1]*C.char)(unsafe.Pointer(self.sound_urls))[:self.sound_urls_len:self.sound_urls_len]
			for i := 0 ; i < len(slice); i++ {
				if slice[i] != nil {
					C.free(unsafe.Pointer(slice[i]))
				}
			}
			C.free(unsafe.Pointer(self.sound_urls))
		}
		if self.author != nil { C.free(unsafe.Pointer(self.author)) }
		if self.word != nil { C.free(unsafe.Pointer(self.word)) }
		if self.written_on != nil { C.free(unsafe.Pointer(self.written_on)) }
		if self.example != nil { C.free(unsafe.Pointer(self.example)) }
	}
}

func main() {}
