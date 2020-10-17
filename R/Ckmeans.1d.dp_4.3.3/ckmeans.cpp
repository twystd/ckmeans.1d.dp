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

// FORWARD DECLARATIONS

template <int N, int K> void compare(cluster<N,K>& p, cluster<N,K>& q);

void testWeightedInput(const std::string&);
void testGivenK(const std::string&);
void testNlteK(const std::string&);
void testKeq2(const std::string&);
void testKeq1(const std::string&);


// MAIN
int main(int argc,char *argv[]) {
    std::cout << "ckmeans v4.3.3" << std::endl;

    std::string methods[] = { "linear", "loglinear", "quadratic" };

    for (const std::string &method: methods) {
    //     std::cout << std::endl << method << std::endl;
    // 
    //     testWeightedInput("linear");
    //     testGivenK(method);
    //     testNlteK(method);
    //     testKeq2(method);
    //     testKeq1(method);
    }
    
    testWeightedInput("linear");

    return 0;
}

// test_that("Weighted input", {
void testWeightedInput(const std::string& method) {
     std::cout << "   test with weighted input" << std::endl;

     { double        data[]    = {-1, 2, 4, 5, 6};
       double        weights[] = { 4, 3, 1, 1, 1};
       cluster<5,3> p = {{},{},{},{}};
       cluster<5,3> q = {{1,2,3,3,3},{-1, 2, 5},{0,0,2},{4,3,3}};
       double BIC;

       kmeans_1d_dp(data, 5, weights, 3, 3,
                    p.clusters.data(), p.centers.data(), p.withins.data(), p.size.data(), &BIC,
                    "BIC", method, L2);
 
       // rebase cluster indices to match 'R'
       for (size_t i=0; i<5; ++i) {
           p.clusters[i]++;
       }

       compare(p,q);
     }

     // NOTE: kmeans_1d_dp with range of K somehow seems to corrupt struct cluster arrays (no idea why - code below works fine though)
     { double data[]    = {-0.9, 1, 1.1, 1.9, 2, 2.1};
       double weights[] = { 3,   1, 2,   2,   1, 1  };
       int    clusters[6];
       double centers[6];
       double withins[6];
       double sizes[6];
       double BIC;

       kmeans_1d_dp(data, 6, weights, 1, 6,
                    clusters, centers, withins, sizes, &BIC,
                    "BIC", method, L2);
 
       // rebase cluster indices to match 'R'
       for (size_t i=0; i<6; ++i) {
           clusters[i]++;
       }

       cluster<6,3> p = {{clusters[0],clusters[1],clusters[2],clusters[3],clusters[4],clusters[5]},{centers[0],centers[1],centers[2]},{withins[0],withins[1],withins[2]},{sizes[0],sizes[1],sizes[2]}};
       cluster<6,3> q = {{1,2,2,3,3,3},{-0.9, (1+2.2)/3, (1.9*2+2+2.1)/4},{0,0.00666667,0.0275},{3,3,4}};

       compare(p,q);
     }
}

// test_that("Given the number of clusters", {
void testGivenK(const std::string& method) {
     std::cout << "   test with given number of clusters" << std::endl;

     double        data[] = {-1, 2, -1, 2, 4, 5, 6, -1, 2, -1};
     cluster<10,3> p = {{},{},{},{}};
     cluster<10,3> q = {{1,2,1,2,3,3,3,1,2,1},{-1, 2, 5},{0,0,2},{4,3,3}};
     double BIC;

     kmeans_1d_dp(data, 10, NULL, 3, 3,
                  p.clusters.data(), p.centers.data(), p.withins.data(), p.size.data(), &BIC,
                  "BIC", method, L2);
 
     // rebase cluster indices to match 'R'
     for (size_t i=0; i<10; ++i) {
         p.clusters[i]++;
     }

     compare(p,q);

     // Ref. https://stackoverflow.com/questions/8637460/k-means-return-value-in-r

     // totss.truth <- sum(scale(x, scale=FALSE)^2)
     // expect_equal(result$totss, totss.truth)
     // expect_equal(result$tot.withinss, 2)
     // expect_equal(result$betweenss, totss.truth - sum(withinss.truth))
}

// test_Ckmeans.1d.dp::test_that("n<=k"...
void testNlteK(const std::string& method) {
     std::cout << "   test with N<=K" << std::endl;

     double        data[] = {3, 2, -5.4, 0.1};
     cluster<4,4> p = {{},{},{},{}};
     cluster<4,4> q = {{4, 3, 1, 2},{-5.4, 0.1, 2, 3},{0, 0, 0, 0},{1, 1, 1, 1}};
     double BIC;

     kmeans_1d_dp(data, 4, NULL, 4, 4,
                  p.clusters.data(), p.centers.data(), p.withins.data(), p.size.data(), &BIC,
                  "BIC", method, L2);
 
     // rebase cluster indices to match 'R'
     for (size_t i=0; i<10; ++i) {
         p.clusters[i]++;
     }

     compare(p,q);
}

// test_Ckmeans.1d.dp::test_that("k==2"...
void testKeq2(const std::string& method) {
     std::cout << "   test with K=2" << std::endl;

     double        data[] = {1,2,3,4,5,6,7,8,9,10};
     cluster<10,2> p = {{},{},{},{}};
     cluster<10,2> q = {{1, 1, 1, 1, 1, 2, 2, 2, 2, 2},{3,8},{10,10},{5,5}};
     double BIC;

     kmeans_1d_dp(data, 10, NULL, 2, 2,
                  p.clusters.data(), p.centers.data(), p.withins.data(), p.size.data(), &BIC,
                  "BIC", method, L2);
 
     // rebase cluster indices to match 'R'
     for (size_t i=0; i<10; ++i) {
         p.clusters[i]++;
     }

     compare(p,q);
}


// test_Ckmeans.1d.dp::test_that("k==1"...
void testKeq1(const std::string& method) {
     { std::cout << "   test with single unique value" << std::endl;
     
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

     { static std::random_device rd;
       static std::mt19937 e2(rd());
       static std::uniform_real_distribution<> dist(-100.0, +100.0);
  
       std::cout << "   test with K=1" << std::endl;
  
       double         data[100];
       cluster<100,1> p = {{},{},{},{}};
       cluster<100,1> q = {{}, {}, {}, {100}};
       double         BIC;
  
       for (int i=0; i<100; i++) {
           data[i] = dist(rd);
       }
  
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
}

// test_that("n==10, k==3", {
// test_that("n==14, k==8", {
// test_that("Estimating k example set 1", {
// test_that("Estimating k example set 2", {
// test_that("Estimating k example set 3 cosine", {
// test_that("Estimating k example set 4 gamma", {

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
        for (int i=0; i<K; i++) {
            double d = abs(p.centers[i] - q.centers[i]);

            if (d > 0.00001) {
               std::cout << " returned invalid centers" << std::endl;
               std::cout << "   expected: [ ";
               std::copy(std::begin(q.centers), std::end(q.centers), std::ostream_iterator<double>(std::cout, " "));
               std::cout << "]" << std::endl;
               std::cout << "   got:      [ ";
               std::copy(std::begin(p.centers), std::end(p.centers), std::ostream_iterator<double>(std::cout, " "));
               std::cout << "]" << std::endl;
               break;
            }
        }
     }

     if (p.withins != q.withins) {
        for (int i=0; i<K; i++) {
            double d = abs(p.withins[i] - q.withins[i]);

            if (d > 0.00001) {
               std::cout << " returned invalid withins" << std::endl;
               std::cout << "   expected: [ ";
               std::copy(std::begin(q.withins), std::end(q.withins), std::ostream_iterator<double>(std::cout, " "));
               std::cout << "]" << std::endl;
               std::cout << "   got:      [ ";
               std::copy(std::begin(p.withins), std::end(p.withins), std::ostream_iterator<double>(std::cout, " "));
               std::cout << "]" << std::endl;
               break;
            }
        }
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

