CC = g++
CXXFLAGS = -pedantic -Wall -std=c++11 -g

src = $(wildcard *.cpp)
obj = $(src:.cpp=.o)

.PHONY: clean

# ckmeans.o: ckmeans.cpp
# 	$(CC) $(CFLAGS) -c ckmeans.cpp -o $(OBJ)/$@

all: ckmeans

ckmeans: $(obj)
	$(CC) $(CXXFLAGS) -o $@ $^ 

run: ckmeans
	./ckmeans

clean:
	rm *.o
	rm ckmeans
