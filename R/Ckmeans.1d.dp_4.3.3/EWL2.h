#include <cstddef> // For size_t
#include <vector>
#include <string>

#include "EWL2_within_cluster.h"

namespace EWL2 {

void fill_dp_matrix(
    const std::vector<double> & x,
    const std::vector<double> & w,
    std::vector< std::vector< ldouble > > & S,
    std::vector< std::vector< size_t > > & J,
    const std::string & method);

void fill_row_q_SMAWK(
    int imin, int imax, int q,
    std::vector< std::vector<ldouble> > & S,
    std::vector< std::vector<size_t> > & J,
    const std::vector<ldouble> & sum_x,
    const std::vector<ldouble> & sum_x_sq);

void fill_row_q(
    int imin, int imax, int q,
    std::vector< std::vector<ldouble> > & S,
    std::vector< std::vector<size_t> > & J,
    const std::vector<ldouble> & sum_x,
    const std::vector<ldouble> & sum_x_sq);

void fill_row_q_log_linear(
    int imin, int imax, int q, int jmin, int jmax,
    std::vector< std::vector<ldouble> > & S,
    std::vector< std::vector<size_t> > & J,
    const std::vector<ldouble> & sum_x,
    const std::vector<ldouble> & sum_x_sq);

}
