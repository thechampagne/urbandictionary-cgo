CGO := go build
STATIC := -buildmode=c-archive
SHARED := -buildmode=c-shared
LIBS := build/urbandictionary.a build/urbandictionary.so

.PHONY: all
all: $(LIBS)

build/urbandictionary.a: urbandictionary.go
	$(CGO) $(STATIC) -o build/urbandictionary.a $<

build/urbandictionary.so: urbandictionary.go
	$(CGO) $(SHARED) -o build/urbandictionary.so $<

.PHONY: clean
clean:
	find build -type f \( -name '*.h' -o -name '*.so' -o -name '*.a' \) -delete
