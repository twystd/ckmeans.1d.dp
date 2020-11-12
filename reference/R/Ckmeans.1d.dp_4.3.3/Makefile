CC = g++
CXXFLAGS = -pedantic -Wall -std=c++11
SRC = src
BUILD = build
BIN = bin
OBJ = $(BUILD)/ckmeans.o \
      $(BUILD)/Ckmeans.1d.dp.o $(BUILD)/dynamic_prog.o \
      $(BUILD)/fill_log_linear.o $(BUILD)/fill_quadratic.o $(BUILD)/fill_SMAWK.o \
      $(BUILD)/select_levels.o $(BUILD)/weighted_select_levels.o \
      $(BUILD)/EWL2_dynamic_prog.o \
      $(BUILD)/EWL2_fill_log_linear.o $(BUILD)/EWL2_fill_quadratic.o $(BUILD)/EWL2_fill_SMAWK.o

.PHONY: clean
.PHONY: build

$(BUILD)/ckmeans.o: ckmeans.cpp 
	mkdir -p $(BUILD)
	$(CC) $(CXXFLAGS) -c ckmeans.cpp -o $@

$(BUILD)/Ckmeans.1d.dp.o: $(SRC)/Ckmeans.1d.dp.cpp
	mkdir -p $(BUILD)
	$(CC) $(CXXFLAGS) -c src/Ckmeans.1d.dp.cpp -o $@

$(BUILD)/dynamic_prog.o: $(SRC)/dynamic_prog.cpp
	$(CC) $(CXXFLAGS) -c $^ -o $@

$(BUILD)/fill_log_linear.o: $(SRC)/fill_log_linear.cpp
	mkdir -p $(BUILD)
	$(CC) $(CXXFLAGS) -c $^ -o $@

$(BUILD)/fill_quadratic.o: $(SRC)/fill_quadratic.cpp
	mkdir -p $(BUILD)
	$(CC) $(CXXFLAGS) -c $^ -o $@

$(BUILD)/select_levels.o: $(SRC)/select_levels.cpp
	mkdir -p $(BUILD)
	$(CC) $(CXXFLAGS) -c $^ -o $@

$(BUILD)/weighted_select_levels.o: $(SRC)/weighted_select_levels.cpp
	mkdir -p $(BUILD)
	$(CC) $(CXXFLAGS) -c $^ -o $@

$(BUILD)/fill_SMAWK.o: $(SRC)/fill_SMAWK.cpp
	mkdir -p $(BUILD)
	$(CC) $(CXXFLAGS) -c $^ -o $@

$(BUILD)/EWL2_dynamic_prog.o: $(SRC)/EWL2_dynamic_prog.cpp
	mkdir -p $(BUILD)
	$(CC) $(CXXFLAGS) -c $^ -o $@

$(BUILD)/EWL2_fill_log_linear.o: $(SRC)/EWL2_fill_log_linear.cpp
	mkdir -p $(BUILD)
	$(CC) $(CXXFLAGS) -c $^ -o $@

$(BUILD)/EWL2_fill_quadratic.o: $(SRC)/EWL2_fill_quadratic.cpp
	mkdir -p $(BUILD)
	$(CC) $(CXXFLAGS) -c $^ -o $@

$(BUILD)/EWL2_fill_SMAWK.o: $(SRC)/EWL2_fill_SMAWK.cpp
	mkdir -p $(BUILD)
	$(CC) $(CXXFLAGS) -c $^ -o $@

ckmeans: $(OBJ)
	mkdir -p $(BIN)
	$(CC) $(CXXFLAGS) -g -o $(BIN)/$@ $^ 

build: ckmeans

test: ckmeans
	$(BIN)/ckmeans

run: ckmeans
	$(BIN)/ckmeans

clean:
	rm -f $(BUILD)/*
	rm -f $(BIN)/*

all: clean build test
