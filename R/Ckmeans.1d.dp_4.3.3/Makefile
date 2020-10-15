CC = g++
CFLAGS = -pedantic -Wall

TARGET = ckmeans
OUTDIR = ./bin
OBJ = ./bin

objects = Ckmeans.1d.dp.o ckmeans.o

all: bin/ckmeans

Ckmeans.1d.dp.o: Ckmeans.1d.dp.cpp Ckmeans.1d.dp.h
	$(CC) $(CFLAGS) -c Ckmeans.1d.dp.cpp -o bin/Ckmeans.1d.dp.o

ckmeans.o: ckmeans.cpp
	$(CC) $(CFLAGS) -c ckmeans.cpp -o bin/ckmeans.o

ckmeans: ckmeans.o Ckmeans.1d.dp.o
	$(CC) ./bin/ckmeans.o ./bin/Ckmeans.1d.dp.o -o bin/ckmeans

run: ckmeans
	./bin/ckmeans

clean:
	rm -f bin/*.o bin/ckmeans
