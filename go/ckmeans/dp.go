package ckmeans

func backtrack(x []float64, J [][]int, k int) []int {
	count := make([]int, k)
	N := len(J[0])
	cluster_right := N - 1
	var cluster_left int

	for q := k - 1; q >= 0; q-- {
		cluster_left = J[q][cluster_right]
		count[q] = cluster_right - cluster_left + 1
		if q > 0 {
			cluster_right = cluster_left - 1
		}
	}

	return count
}

//void backtrack(const std::vector<double> & x,
//               const std::vector< std::vector< size_t > > & J,
//               std::vector<size_t> & count, const int K)
//{ std::cout << "backtrack/4" << std::endl;
//  // const int K = (int) J.size();
//  const size_t N = J[0].size();
//  size_t cluster_right = N-1;
//  size_t cluster_left;
//
//  // Backtrack the clusters from the dynamic programming matrix
//  for(int q = K-1; q >= 0; --q) {
//    cluster_left = J[q][cluster_right];
//    count[q] = cluster_right - cluster_left + 1;
//    if(q > 0) {
//      cluster_right = cluster_left - 1;
//    }
//  }
//}
