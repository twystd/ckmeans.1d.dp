#include <iostream>
#include <array>

#include "Ckmeans.1d.dp.h"

template <int N, int K> struct cluster {
   std::array<int,N>    clusters;
   std::array<double,K> centers;
   std::array<double,K> withins;
   std::array<double,K> size;
};

template <int N, int K> void compare(cluster<N,K>& p, cluster<N,K>& q) {
     if (p.clusters != q.clusters) {
        std::cout << " returned invalid clusters" << std::endl;
        std::cout << "   expected: [ ";
        std::copy(std::begin(q.clusters), std::end(q.clusters), std::ostream_iterator<int>(std::cout, " "));
        std::cout << "]" << std::endl;
        std::cout << "   got:      [ ";
        std::copy(std::begin(p.clusters), std::end(p.clusters), std::ostream_iterator<int>(std::cout, " "));
        std::cout << "]" << std::endl;
     }

     if (p.centers != q.centers) {
        std::cout << " returned invalid centers" << std::endl;
        std::cout << "   expected: [ ";
        std::copy(std::begin(q.centers), std::end(q.centers), std::ostream_iterator<double>(std::cout, " "));
        std::cout << "]" << std::endl;
        std::cout << "   got:      [ ";
        std::copy(std::begin(p.centers), std::end(p.centers), std::ostream_iterator<double>(std::cout, " "));
        std::cout << "]" << std::endl;
     }

     if (p.withins != q.withins) {
        std::cout << " returned invalid withins" << std::endl;
        std::cout << "   expected: [ ";
        std::copy(std::begin(q.withins), std::end(q.withins), std::ostream_iterator<double>(std::cout, " "));
        std::cout << "]" << std::endl;
        std::cout << "   got:      [ ";
        std::copy(std::begin(p.withins), std::end(p.withins), std::ostream_iterator<double>(std::cout, " "));
        std::cout << "]" << std::endl;
     }

     if (p.size != q.size) {
        std::cout << " returned invalid size" << std::endl;
        std::cout << "   expected: [ ";
        std::copy(std::begin(q.size), std::end(q.size), std::ostream_iterator<double>(std::cout, " "));
        std::cout << "]" << std::endl;
        std::cout << "   got:      [ ";
        std::copy(std::begin(p.size), std::end(p.size), std::ostream_iterator<double>(std::cout, " "));
        std::cout << "]" << std::endl;
     }
}

// FORWARD DECLARATIONS
void testForSingleUniqueValue(const std::string&);
void testForK1(const std::string&);

// MAIN
int main(int argc,char *argv[]) {
    std::cout << "ckmeans v4.3.3" << std::endl;

    testForSingleUniqueValue("linear");
    testForSingleUniqueValue("loglinear");
    testForSingleUniqueValue("quadratic");

    testForK1("linear");

    return 0;
}

void testForSingleUniqueValue(const std::string& method) {
    std::cout << "- test for single unique value:" << method << std::endl;
     
    double       data[] = {-2.5,-2.5,-2.5,-2.5};
    cluster<4,1> p = {{0,0,0,0}, {0},    {0}, {0}};
    cluster<4,1> q = {{1,1,1,1}, {-2.5}, {0}, {4}};
    double       BIC;

    kmeans_1d_dp(data, 4, NULL, 1, 1,
                 p.clusters.data(), p.centers.data(), p.withins.data(), p.size.data(), &BIC,
                 "BIC" ,method, L2);

    // rebase cluster indices to match 'R'
    for (size_t i=0; i<4; ++i) {
        p.clusters[i]++;
    }

    compare(p,q);
}

void testForK1(const std::string& method) {
    std::cout << "- test for K=1:" << method << std::endl;
     
//  TODO initialise with random values x <- rep(1, 100) 
    double data[] = {     1, 2, 3, 4, 5, 6, 7, 8,10,  9,11,12,13,14,15,16,17,18,19,
                      20,21,22,23,24,25,26,27,28,29,  30,31,32,33,34,35,36,37,38,39,
                      40,41,42,43,44,45,46,47,48,49,  50,51,52,53,54,55,56,57,58,59,
                      60,61,62,63,64,65,66,67,68,69,  70,71,72,73,74,75,76,77,78,79,
                      80,81,82,83,84,85,86,87,88,89,  90,91,92,93,94,95,96,97,98,99,
                      100 };

    cluster<100,1> p = {{ 0,0,0,0,0,0,0,0,0,0,  0,0,0,0,0,0,0,0,0,0,
                          0,0,0,0,0,0,0,0,0,0,  0,0,0,0,0,0,0,0,0,0,
                          0,0,0,0,0,0,0,0,0,0,  0,0,0,0,0,0,0,0,0,0,
                          0,0,0,0,0,0,0,0,0,0,  0,0,0,0,0,0,0,0,0,0,
                          0,0,0,0,0,0,0,0,0,0,  0,0,0,0,0,0,0,0,0,0 }, {0}, {0}, {0}};

    cluster<100,1> q = {{ 0,0,0,0,0,0,0,0,0,0,  0,0,0,0,0,0,0,0,0,0,
                          0,0,0,0,0,0,0,0,0,0,  0,0,0,0,0,0,0,0,0,0,
                          0,0,0,0,0,0,0,0,0,0,  0,0,0,0,0,0,0,0,0,0,
                          0,0,0,0,0,0,0,0,0,0,  0,0,0,0,0,0,0,0,0,0,
                          0,0,0,0,0,0,0,0,0,0,  0,0,0,0,0,0,0,0,0,0 }, {0}, {0}, {100}};

    double BIC;

    kmeans_1d_dp(data, 100, NULL, 1, 1,
                 p.clusters.data(), p.centers.data(), p.withins.data(), p.size.data(), &BIC,
                 "BIC", method, L2);

    // rebase cluster indices to match 'R'
    for (size_t i=0; i<4; ++i) {
        p.clusters[i]++;
    }

    compare(p,q);
//  expect_equal(result$size, 100)
}


