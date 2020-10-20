#include "Ckmeans.1d.dp.h"
#include "EWL2.h"

// #include <algorithm>
// #include <cmath>
#include <iostream>
#include <string>
// #include <vector>
// #include <cstring>

template <class ForwardIterator>

size_t numberOfUnique(ForwardIterator first, ForwardIterator last) {
       size_t nUnique;
 
       if (first == last) {
          nUnique = 0;
       } else {
          nUnique = 1;
          for (ForwardIterator itr=first+1; itr!=last; ++itr) {
              if (*itr != *(itr -1)) {
                 nUnique ++;
              }
          }
   }

   return nUnique;
}

static const double * px;

bool compi(size_t i, size_t j) {
   return px[i] < px[j];
}

void kmeans_1d_dp(const double *x, const size_t N, const double *y, size_t Kmin, size_t Kmax,
                  int    *cluster, double *centers, double *withinss, double *size, double *BIC,
                  const std::string &estimate_k, const std::string &method, const enum DISSIMILARITY criterion)
{
  // Input:
  //  x -- an array of double precision numbers, not necessarily sorted
  //  Kmin -- the minimum number of clusters expected
  //  Kmax -- the maximum number of clusters expected
  // NOTE: All vectors in this program is considered starting at position 0.

  std::vector<size_t> order(N);

  //Number generation using lambda function, not supported by all g++:
  //std::size_t n(0);
  //std::generate(order.begin(), order.end(), [&]{ return n++; });

  for(size_t i=0; i<order.size(); ++i) {
    order[i] = i;
  }

  bool is_sorted(true);
  for(size_t i=0; i<N-1; ++i) {
    if(x[i] > x[i+1]) {
      is_sorted = false;
      break;
    }
  }

  std::vector<double> x_sorted(x, x+N);
  std::vector<double> y_sorted;
  bool is_equally_weighted = true;

  if (!is_sorted) {
     // Sort the index of x in increasing order of x

     // Option 1.
     // Sorting using lambda function, not supported by all g++ versions:
     // std::sort(order.begin(), order.end(),
     //           [&](size_t i1, size_t i2) { return x[i1] < x[i2]; } );
 
     // Option 2. The following is not supported by C++98:
     // struct CompareIndex {
     //   const double * m_x;
     //   CompareIndex(const double * x) : m_x(x) {}
     //   bool operator() (size_t i, size_t j) { return (m_x[i] < m_x[j]);}
     //} compi(x);
     // 
     // std::sort(order.begin(), order.end(), compi);

     // Option 3:
     px = x;
     std::sort(order.begin(), order.end(), compi);
 
     for (size_t i=0ul; i<order.size(); ++i) {
         x_sorted[i] = x[order[i]];
     }
  }

  // check to see if unequal weight is provided
  if (y != NULL) {
     is_equally_weighted = true;
     for (size_t i=1; i<N; ++i) {
         if (y[i] != y[i-1]) {
            is_equally_weighted = false;
            break;
         }
     }
  }

  if (!is_equally_weighted) {
     y_sorted.resize(N);
     for (size_t i=0; i<order.size(); ++i) {
         y_sorted[i] = y[order[i]];
     }
  }

  const size_t nUnique = numberOfUnique(x_sorted.begin(), x_sorted.end());

  Kmax = nUnique < Kmax ? nUnique : Kmax;

  if (nUnique > 1) { // The case when not all elements are equal.
     std::vector< std::vector< ldouble > > S( Kmax, std::vector<ldouble>(N) );
     std::vector< std::vector< size_t > > J( Kmax, std::vector<size_t>(N) );
     size_t Kopt;
 
     // Fill in dynamic programming matrix
     if (is_equally_weighted) {
        if (criterion == L2) {
           EWL2::fill_dp_matrix(x_sorted, y_sorted, S, J, method);
        } else {
           fill_dp_matrix(x_sorted, y_sorted, S, J, method, criterion);
        }

       // Choose an optimal number of levels between Kmin and Kmax
       if (estimate_k == "BIC") {
          Kopt = select_levels(x_sorted, J, Kmin, Kmax, BIC);
       } else {
          Kopt = select_levels_3_4_12(x_sorted, J, Kmin, Kmax, BIC);
       }

     } else {
       fill_dp_matrix(x_sorted, y_sorted, S, J, method, criterion);

       switch(criterion) {
       case L2Y:
            if (estimate_k=="BIC") {
               Kopt = select_levels(y_sorted, J, Kmin, Kmax, BIC);
            } else {
               Kopt = select_levels_3_4_12(y_sorted, J, Kmin, Kmax, BIC);
            }
            break;
 
       default:
            if (estimate_k=="BIC") {
               // Choose an optimal number of levels between Kmin and Kmax
               Kopt = select_levels_weighted(x_sorted, y_sorted, J, Kmin, Kmax, BIC);
            } else {
               Kopt = select_levels_weighted_3_4_12(x_sorted, y_sorted, J, Kmin, Kmax, BIC);
            }
       }
     }

     if (Kopt < Kmax) { // Reform the dynamic programming matrix S and J
       J.erase(J.begin() + Kopt, J.end());
     }

     std::vector<int> cluster_sorted(N);

     // Backtrack to find the clusters beginning and ending indices
     if(is_equally_weighted && criterion == L1) {
         backtrack_L1(x_sorted, J, &cluster_sorted[0], centers, withinss, size);
     } else if (is_equally_weighted && criterion == L2) {
         backtrack(x_sorted, J, &cluster_sorted[0], centers, withinss, size);
     } else if(criterion == L2Y) {
       backtrack_L2Y(x_sorted, y_sorted, J, &cluster_sorted[0],
                     centers, withinss, size);
 
     } else {
       backtrack_weighted(x_sorted, y_sorted, J, &cluster_sorted[0], centers, withinss, size);
     }

#ifdef DEBUG
    std::cout << "backtrack done." << std::endl;
#endif

     for (size_t i = 0; i < N; ++i) {
         // Obtain clustering on data in the original order
         cluster[order[i]] = cluster_sorted[i];
     }

  } else {  // A single cluster that contains all elements
    for (size_t i=0; i<N; ++i) {
        cluster[i] = 0;
    }
 
    centers[0] = x[0];
    withinss[0] = 0.0;
    size[0] = N * (is_equally_weighted ? 1 : y[0]);
  }
}  //end of kmeans_1d_dp()
