CC = g++
CFLAGS = -pedantic -Wall

all: ckmeans

Ckmeans.1d.dp.o: Ckmeans.1d.dp.cpp Ckmeans.1d.dp.h
	$(CC) $(CFLAGS) -c Ckmeans.1d.dp.cpp -o Ckmeans.1d.dp.o

ckmeans.o: ckmeans.cpp
	$(CC) $(CFLAGS) -c ckmeans.cpp -o ckmeans.o

ckmeans: ckmeans.o Ckmeans.1d.dp.o
	$(CC) ckmeans.o Ckmeans.1d.dp.o -o ckmeans

run: ckmeans
	./ckmeans

clean:
	rm -f *.o ckmeans
