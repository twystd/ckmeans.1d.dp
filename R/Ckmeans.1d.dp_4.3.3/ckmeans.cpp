#include <iostream>
#include <array>
#include <vector>
#include <iomanip>
#include <map>
#include <random>
#include <cmath>

#include "Ckmeans.1d.dp.h"

template <int N, int K> struct cluster {
   std::array<int,N>    clusters;
   std::array<double,N> centers;
   std::array<double,N> withins;
   std::array<double,K> size;
   std::array<double,K> BIC; // TODO check - fixes the segfault but should probably be kMax
};

// FORWARD DECLARATIONS

template <int N, int K> void rebase(cluster<N,K>&);
template <int N, int K> void compare(cluster<N,K>&,cluster<N,K>&);

void testWeightedInput(const std::string&);
void testGivenK(const std::string&);
void testNlteK(const std::string&);
void testKeq2(const std::string&);
void testKeq1(const std::string&);
void testN10K3(const std::string&);
void testN14K8(const std::string&);
void testEstimateKExampleSet1(const std::string& method);
void testEstimateKExampleSet2(const std::string& method);
void testEstimateKExampleSet3(const std::string& method);
void testEstimateKExampleSet4(const std::string& method);

void testTaps();

// MAIN
int main(int argc,char *argv[]) {
    std::cout << "ckmeans v4.3.3" << std::endl;

//    std::string methods[] = { "linear", "loglinear", "quadratic" };
//
//    for (const std::string &method: methods) {
//         std::cout << std::endl << method << std::endl;
//     
//         testWeightedInput(method);
//         testGivenK(method);
//         testNlteK(method);
//         testKeq2(method);                    // Go'd
//         testKeq1(method);                    // Go'd
//         testN10K3(method);                   // Go'd
//         testN14K8(method);                   // Go'd
//         testEstimateKExampleSet1(method);
//         testEstimateKExampleSet2(method);
//         testEstimateKExampleSet3(method);
//         testEstimateKExampleSet4(method);
//
//         testTaps();
//    }
    
    testN14K8("linear");

    return 0;
}


void testTaps() {
     double taps[] = { 4.570271991, 5.063594027, 5.603539973, 6.102690998, 6.642708943, 7.141796968, 7.710649857, 8.192470916,
                       4.506176116, 5.045971061, 5.591722996, 6.114172975, 6.619153989, 7.13578898,  7.693071891, 8.203885893,
                       4.52956007,  5.057670039, 5.591721996, 6.13742393,  6.630941966, 7.1766839,   7.69897488,  8.227207848,
                       4.52956007,  5.069284016, 5.603428973, 6.102591998, 6.613455,    7.147644957, 7.69912088,  8.215609871,
                       4.517865093, 5.022782107, 5.580101018, 6.096715009, 6.654118921, 7.1763719,   7.681405914, 8.215537871,
                                    5.133092891, 5.545395086, 6.067721066, 6.578564068, 7.130096991, 7.652464971, 8.13427303,
                       4.494581138, 5.040234073, 5.562732052, 6.079333043, 6.624973977, 7.141650968, 7.664070948, 8.198270905,
                       4.52940807,  5.040295073, 5.556940064, 6.131584941, 6.654145921, 7.193876866, 7.722112835, 8.244539814,
                       4.523631082, 5.046071061, 5.586102007, 6.09099502,  6.596029034, 7.130224991, 7.652501971, 8.180805939,
                       4.517979093, 5.046165061, 5.551068075, 6.073547054, 6.607636011, 7.165018923, 7.687334903, 8.238953825,
                       4.517911093, 5.069403016, 5.586174007, 6.108568986, 6.578649068, 7.147523957, 7.681606914, 8.26211078
                     };

     cluster<87,87> p = {{},{},{},{}};
     std::map<int, std::vector<double>> beats;

     kmeans_1d_dp(taps, 87, NULL, 1, 87,
                  p.clusters.data(), p.centers.data(), p.withins.data(), p.size.data(), p.BIC.data(),
                 "BIC", "linear", L2);
 
     rebase(p);

     for (int i=0; i<87; i++) {
         int beat = p.clusters[i];

         beats[beat].push_back(taps[i]) ;
     }

     std::cout << "TAPS:" << std::endl << "   ";
     std::copy(std::begin(p.clusters), std::end(p.clusters), std::ostream_iterator<int>(std::cout, " "));
     std::cout << std::endl;

     for (std::map<int,std::vector<double>>::iterator it=beats.begin(); it != beats.end(); ++it) {
         std::cout << "   " << it->first << " => ";
         for (std::vector<double>::iterator j=it->second.begin(); j != it->second.end(); ++j) {
             std::cout << std::setprecision(6) << std::setfill(' ') << std::setw(9) << std::left << *j;
         }
         // std::copy(std::begin(it->second), std::end(it->second), std::ostream_iterator<double>(std::cout, " "));
         std::cout << std::endl;
     }
}

// test_that("Weighted input", {
void testWeightedInput(const std::string& method) {
     std::cout << "   test with weighted input" << std::endl;

     { double        data[]    = {-1, 2, 4, 5, 6};
       double        weights[] = { 4, 3, 1, 1, 1};
       cluster<5,3> p = {{},{},{},{}};
       cluster<5,3> q = {{1,2,3,3,3},{-1, 2, 5},{0,0,2},{4,3,3}};

       kmeans_1d_dp(data, 5, weights, 3, 3,
                    p.clusters.data(), p.centers.data(), p.withins.data(), p.size.data(), p.BIC.data(),
                    "BIC", method, L2);
 
       rebase(p);
       compare(p,q);
     }

     { double data[]    = {-0.9, 1, 1.1, 1.9, 2, 2.1};
       double weights[] = { 3,   1, 2,   2,   1, 1  };
       cluster<6,3> p = {{},{},{},{}};
       cluster<6,3> q = {{1,2,2,3,3,3},{-0.9, (1+2.2)/3, (1.9*2+2+2.1)/4},{0,0.00666667,0.0275},{3,3,4}};

       kmeans_1d_dp(data, 6, weights, 1, 6,
                    p.clusters.data(), p.centers.data(), p.withins.data(), p.size.data(), p.BIC.data(),
                    "BIC", method, L2);
 
       rebase(p);
       compare(p,q);
     }
}

// test_that("Given the number of clusters", {
void testGivenK(const std::string& method) {
     std::cout << "   test with given number of clusters" << std::endl;

     double        data[] = {-1, 2, -1, 2, 4, 5, 6, -1, 2, -1};
     cluster<10,3> p = {{},{},{},{}};
     cluster<10,3> q = {{1,2,1,2,3,3,3,1,2,1},{-1, 2, 5},{0,0,2},{4,3,3}};

     kmeans_1d_dp(data, 10, NULL, 3, 3,
                  p.clusters.data(), p.centers.data(), p.withins.data(), p.size.data(), p.BIC.data(),
                  "BIC", method, L2);
 
     rebase(p);
     compare(p,q);

     // Ref. https://stackoverflow.com/questions/8637460/k-means-return-value-in-r
     //
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

     kmeans_1d_dp(data, 4, NULL, 4, 4,
                  p.clusters.data(), p.centers.data(), p.withins.data(), p.size.data(), p.BIC.data(),
                  "BIC", method, L2);
 
     rebase(p);
     compare(p,q);
}

// test_Ckmeans.1d.dp::test_that("k==2"...
void testKeq2(const std::string& method) {
     std::cout << "   test with K=2" << std::endl;

     double        data[] = {1,2,3,4,5,6,7,8,9,10};
     cluster<10,2> p = {{},{},{},{}};
     cluster<10,2> q = {{1, 1, 1, 1, 1, 2, 2, 2, 2, 2},{3,8},{10,10},{5,5}};

     kmeans_1d_dp(data, 10, NULL, 2, 2,
                  p.clusters.data(), p.centers.data(), p.withins.data(), p.size.data(), p.BIC.data(),
                  "BIC", method, L2);
 
     rebase(p);
     compare(p,q);
}


// test_Ckmeans.1d.dp::test_that("k==1"...
void testKeq1(const std::string& method) {
     { std::cout << "   test with single unique value" << std::endl;
     
       double       data[] = {-2.5,-2.5,-2.5,-2.5};
       cluster<4,1> p = {{},{},{},{}};
       cluster<4,1> q = {{1,1,1,1}, {-2.5}, {0}, {4}};

       kmeans_1d_dp(data, 4, NULL, 1, 1,
                    p.clusters.data(), p.centers.data(), p.withins.data(), p.size.data(), p.BIC.data(),
                    "BIC" ,method, L2);

       rebase(p);
       compare(p,q);
     }

     { static std::random_device rd;
       static std::mt19937 e2(rd());
       static std::uniform_real_distribution<> dist(-100.0, +100.0);
  
       std::cout << "   test with K=1" << std::endl;
  
       double         data[100];
       cluster<100,1> p = {{},{},{},{}};
       cluster<100,1> q = {{}, {}, {}, {100}};
  
       for (int i=0; i<100; i++) {
           data[i] = dist(rd);
       }
  
       kmeans_1d_dp(data, 100, NULL, 1, 1,
                    p.clusters.data(), p.centers.data(), p.withins.data(), p.size.data(), p.BIC.data(),
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

// test_that("n==10, k==3"...
void testN10K3(const std::string& method) {
     std::cout << "   test with n=10, k=3" << std::endl;

     double        data[]    = {3, 3, 3, 3, 1, 1, 1, 2, 2, 2};
     cluster<10,3> p = {{},{},{},{}};
     cluster<10,3> q = {{3, 3, 3, 3, 1, 1, 1, 2, 2, 2},{1, 2, 3},{0, 0, 0},{3, 3, 4}};

     kmeans_1d_dp(data, 10, NULL, 3, 3,
                  p.clusters.data(), p.centers.data(), p.withins.data(), p.size.data(), p.BIC.data(),
                  "BIC", method, L2);
 
     rebase(p);
     compare(p,q);
}

// test_that("n==14, k==8"...
void testN14K8(const std::string& method) {
     std::cout << "   test with n=14, k=8" << std::endl;

     double        data[] = {-3, 2.2, -6, 7, 9, 11, -6.3, 75, 82.6, 32.3, -9.5, 62.5, 7, 95.2};
     cluster<14,8> p = {{},{},{},{}};
     cluster<14,8> q = {{2, 2, 1, 3, 3, 3, 1, 6, 7, 4, 1, 5, 3, 8},
                        {-7.266666667, -0.4, 8.5, 32.3, 62.5, 75.0, 82.6, 95.2},
                        {7.526666667, 13.52, 11.0, 0.0, 0.0, 0.0, 0.0, 0.0},
                        {3, 2, 4, 1, 1, 1, 1, 1}};

     kmeans_1d_dp(data, 14, NULL, 8, 8,
                  p.clusters.data(), p.centers.data(), p.withins.data(), p.size.data(), p.BIC.data(),
                  "BIC", method, L2);
 
     rebase(p);
     compare(p,q);
}

// test_that("Estimating k example set 1"...
void testEstimateKExampleSet1(const std::string& method) {
     std::cout << "   test estimate K, example set 1" << std::endl;

     { double       data[] = {0.9, 1, 1.1, 1.9, 2, 2.1};
       cluster<6,2> p = {{},{},{},{}};
       cluster<6,2> q = {{1,1,1,2,2,2},{1,2},{0.02,0.02},{3,3}};

       kmeans_1d_dp(data, 6, NULL, 1, 6,
                    p.clusters.data(), p.centers.data(), p.withins.data(), p.size.data(), p.BIC.data(),
                    "BIC", method, L2);
 
       rebase(p);
       compare(p,q);
     }

     { double       data[] = {2.1, 2, 1.9, 1.1, 1, 0.9};
       cluster<6,2> p = {{},{},{},{}};
       cluster<6,2> q = {{2,2,2,1,1,1},{1,2},{0.02,0.02},{3,3}};

       kmeans_1d_dp(data, 6, NULL, 1, 6,
                    p.clusters.data(), p.centers.data(), p.withins.data(), p.size.data(), p.BIC.data(),
                    "BIC", method, L2);
 
       rebase(p);
       compare(p,q);
     }

     { double       data[] = {2.1, 2, 1.9, 1.1, 1, 0.9};
       cluster<6,2> p = {{},{},{},{}};
       cluster<6,2> q = {{2,2,2,1,1,1},{1,2},{0.02,0.02},{3,3}};

       kmeans_1d_dp(data, 6, NULL, 1, 10,
                    p.clusters.data(), p.centers.data(), p.withins.data(), p.size.data(), p.BIC.data(),
                    "BIC", method, L2);

       rebase(p);
       compare(p,q);
     }
}

// test_that("Estimating k example set 2"...
void testEstimateKExampleSet2(const std::string& method) {
     std::cout << "   test estimate K, example set 2" << std::endl;

     double        data[] = {3.5, 3.6, 3.7, 3.1, 1.1, 0.9, 0.8, 2.2, 1.9, 2.1};
     cluster<10,3> p = {{},{},{},{}};
     cluster<10,3> q = {{3, 3, 3, 3, 1, 1, 1, 2, 2, 2},{0.933333333333, 2.066666666667, 3.475},{0.0466666666667, 0.0466666666667, 0.2075},{3, 3, 4}};

     kmeans_1d_dp(data, 10, NULL, 2, 5,
                  p.clusters.data(), p.centers.data(), p.withins.data(), p.size.data(), p.BIC.data(),
                  "BIC", method, L2);
 
     rebase(p);
     compare(p,q);
}

// test_that("Estimating k example set 3 cosine"...
void testEstimateKExampleSet3(const std::string& method) {
     std::cout << "   test estimate K, example set 3 (cosine)" << std::endl;

     // x <- cos((-10:10))
     double data[] = { -0.8390715,-0.9111303,-0.1455000,0.7539023,0.9601703,0.2836622,
                       -0.6536436,-0.9899925,-0.4161468,0.5403023,1.0000000,0.5403023,
                       -0.4161468,-0.9899925,-0.6536436,0.2836622,0.9601703,0.7539023,
                       -0.1455000,-0.9111303,-0.8390715
                     };

     cluster<21,2> p = {{},{},{},{}};
     cluster<21,2> q = {{1,1,1,2,2,2,1,1,1,2,2,2,1,1,1,2,2,2,1,1,1}, {-0.6592474631, 0.6751193405},{1.0564793100, 0.6232976959},{12,9}};

     kmeans_1d_dp(data, 21, NULL, 1, 21,
                  p.clusters.data(), p.centers.data(), p.withins.data(), p.size.data(), p.BIC.data(),
                  "BIC", method, L2);
 
     rebase(p);
     compare(p,q);
     
}

// test_that("Estimating k example set 4 gamma", {
void testEstimateKExampleSet4(const std::string& method) {
     std::cout << "   test estimate K, example set 4 (gamma)" << std::endl;

     // x <- dgamma(seq(1,10, by=0.5), shape=2, rate=1)
     double data[] = { 0.3678794412,0.3346952402,0.2706705665,0.2052124966,0.1493612051,
                       0.1056908420,0.0732625556,0.0499904844,0.0336897350,0.0224772429,
                       0.0148725131,0.0097723548,0.0063831738,0.0041481328,0.0026837010,
                       0.0017294811,0.0011106882,0.0007110924,0.0004539993
                     };

     cluster<19,3> p = {{},{},{},{}};
     cluster<19,3> q = {{3,3,3,2,2,2,1,1,1,1,1,1,1,1,1,1,1,1,1}, {0.01702193495, 0.15342151455, 0.32441508262},{0.006126754998,0.004977009034,0.004883305120},{13,3,3}};

     kmeans_1d_dp(data, 19, NULL, 1, 19,
                  p.clusters.data(), p.centers.data(), p.withins.data(), p.size.data(), p.BIC.data(),
                  "BIC", method, L2);
 
     rebase(p);
     compare(p,q);
}

template <int N, int K> void rebase(cluster<N,K>& p) {
     for (size_t i=0; i<N; ++i) {
         p.clusters[i] += 1;
     }
}

template <int N, int K> void compare(cluster<N,K>& p, cluster<N,K>& q) {
     if (p.clusters != q.clusters) {
        std::cout << "     returned invalid clusters" << std::endl;
        std::cout << "        expected: [ ";
        std::copy(std::begin(q.clusters), std::end(q.clusters), std::ostream_iterator<int>(std::cout, " "));
        std::cout << "]" << std::endl;
        std::cout << "        got:      [ ";
        std::copy(std::begin(p.clusters), std::end(p.clusters), std::ostream_iterator<int>(std::cout, " "));
        std::cout << "]" << std::endl;
     }

     if (p.centers != q.centers) {
        for (int i=0; i<K; i++) {
            double d = abs(p.centers[i] - q.centers[i]);

            if (d > 0.00001) {
               std::cout << "     returned invalid centers" << std::endl;
               std::cout << "        expected: [ ";
               std::copy(std::begin(q.centers), std::end(q.centers), std::ostream_iterator<double>(std::cout, " "));
               std::cout << "]" << std::endl;
               std::cout << "        got:      [ ";
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
               std::cout << "     returned invalid withins" << std::endl;
               std::cout << "       expected: [ ";
               std::copy(std::begin(q.withins), std::end(q.withins), std::ostream_iterator<double>(std::cout, " "));
               std::cout << "]" << std::endl;
               std::cout << "       got:      [ ";
               std::copy(std::begin(p.withins), std::end(p.withins), std::ostream_iterator<double>(std::cout, " "));
               std::cout << "]" << std::endl;
               break;
            }
        }
     }

     if (p.size != q.size) {
        std::cout << "     returned invalid size" << std::endl;
        std::cout << "       expected: [ ";
        std::copy(std::begin(q.size), std::end(q.size), std::ostream_iterator<double>(std::cout, " "));
        std::cout << "]" << std::endl;
        std::cout << "       got:      [ ";
        std::copy(std::begin(p.size), std::end(p.size), std::ostream_iterator<double>(std::cout, " "));
        std::cout << "]" << std::endl;
     }
}

