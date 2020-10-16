CC = g++
CFLAGS = -pedantic -Wall

src = $(wildcard *.cpp)
obj = $(src:.cpp=.o)

# ckmeans.o: ckmeans.cpp
# 	$(CC) $(CFLAGS) -c ckmeans.cpp -o $(OBJ)/$@

all: ckmeans

ckmeans: $(obj)
	$(CC) $(CFLAGS) -o $@ $^ $(LDFLAGS)

run: ckmeans
	./ckmeans

clean:
	rm *.o
	rm ckmeans
