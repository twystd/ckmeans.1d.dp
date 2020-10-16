#include <iostream>
#include <array>
#include <random>

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

    // test_Ckmeans.1d.dp::test_that("k==1"...
    testForSingleUniqueValue("linear");
    testForSingleUniqueValue("loglinear");
    testForSingleUniqueValue("quadratic");
    testForK1("linear");
    testForK1("loglinear");
    testForK1("quadratic");

    return 0;
}

void testForSingleUniqueValue(const std::string& method) {
     std::cout << "- test for single unique value:" << method << std::endl;
     
     double       data[] = {-2.5,-2.5,-2.5,-2.5};
     cluster<4,1> p = {{},{},{},{}};
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
     static std::random_device rd;
     static std::mt19937 e2(rd());
     static std::uniform_real_distribution<> dist(-100.0, +100.0);

     std::cout << "- test for K=1:" << method << std::endl;

     double data[100];
     cluster<100,1> p = {{},{},{},{}};
     double BIC;

     for (int i=0; i<100; i++) {
         data[i] = dist(rd);
     }

     cluster<100,1> q = {{}, {}, {}, {100}};

     kmeans_1d_dp(data, 100, NULL, 1, 1,
                  p.clusters.data(), p.centers.data(), p.withins.data(), p.size.data(), &BIC,
                  "BIC", method, L2);
 
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


