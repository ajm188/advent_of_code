all: run

run: build
	cat input.txt | ./Main

build: *.hs
	ghc -o Main $^

clean:
	rm -f *.hi
	rm -f *.o
	rm -f Main
