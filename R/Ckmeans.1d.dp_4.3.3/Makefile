CC = g++
CFLAGS = -pedantic -Wall

TARGET = ckmeans
OUTDIR = ./bin
OBJ = ./bin

objects = EWL2_dynamic_prog.o Ckmeans.1d.dp.o ckmeans.o EWL2_fill_SMAWK.o

all: bin/ckmeans

EWL2_fill_SMAWK.o: EWL2_fill_SMAWK.cpp EWL2.h
	$(CC) $(CFLAGS) -c EWL2_fill_SMAWK.cpp -o bin/EWL2_fill_SMAWK.o

EWL2_dynamic_prog.o: EWL2_dynamic_prog.cpp EWL2.h
	$(CC) $(CFLAGS) -c EWL2_dynamic_prog.cpp -o bin/EWL2_dynamic_prog.o

Ckmeans.1d.dp.o: Ckmeans.1d.dp.cpp Ckmeans.1d.dp.h
	$(CC) $(CFLAGS) -c Ckmeans.1d.dp.cpp -o bin/Ckmeans.1d.dp.o

ckmeans.o: ckmeans.cpp
	$(CC) $(CFLAGS) -c ckmeans.cpp -o bin/ckmeans.o

ckmeans: ckmeans.o Ckmeans.1d.dp.o EWL2_dynamic_prog.o EWL2_fill_SMAWK.o
	$(CC) bin/ckmeans.o bin/Ckmeans.1d.dp.o bin/EWL2_dynamic_prog.o bin/EWL2_fill_SMAWK.o -o bin/ckmeans

run: ckmeans
	./bin/ckmeans

clean:
	rm -f bin/*.o bin/ckmeans
