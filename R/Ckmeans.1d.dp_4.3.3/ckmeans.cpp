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

// MAIN
int main(int argc,char *argv[]) {
    std::cout << "ckmeans v4.3.3" << std::endl;

    testForSingleUniqueValue("linear");
    testForSingleUniqueValue("loglinear");
    testForSingleUniqueValue("quadratic");

    return 0;
}

void testForSingleUniqueValue(const std::string& method) {
    std::cout << "testForSingleUniqueValue:" << method << std::endl;
     
    double  data[] = {-2.5,-2.5,-2.5,-2.5};
    pq p = {{0,0,0,0}, {0},    {0}, {0}};
    pq q = {{1,1,1,1}, {-2.5}, {0}, {4}};
    double BIC;
    std::string estimate = "";
   

    kmeans_1d_dp(data, 4, NULL, 1, 1,
                 p.clusters.data(), p.centers.data(), p.withins.data(), p.size.data(), &BIC,
                 estimate, method, L2);

    // rebase cluster indices to match 'R'
    for (size_t i=0; i<4; ++i) {
        p.clusters[i]++;
    }

    compare(p,q);
}

// test_that("k==1", {
// 
//     x <- rep(1, 100)
//     result <- Ckmeans.1d.dp(x, 1, method=method)
//     expect_equal(result$size, 100)
//   }
// 
// })

