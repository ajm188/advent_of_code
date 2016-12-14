all: run

run: build cleanup.sh
	cat input.txt | ./cleanup.sh | ./Main

build: *.hs
	ghc -o Main $^

clean:
	rm -f *.hi
	rm -f *.o
	rm -f Main
