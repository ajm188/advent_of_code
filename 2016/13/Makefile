all: run

run: build
	./Main

build: *.hs
	ghc -o Main $^

clean:
	rm -f *.hi
	rm -f *.o
	rm -f Main
