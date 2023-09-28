'''
Port of Ckmeans.1d.dp_4.3.3 from R to Python.
'''

from array import array

from SMAWK import SMAWK
from L2 import L2
from dp import fill_dp_matrix


def ckmeans_1d_dp(data, weights, compare=None, kmin=None, kmax=None):
    # validate inputs
    if data == None or len(data) == 0:
        return []

    if weights and len(weights) != len(data):
        raise ValueError('number of weights must match the data')

    # sort and order data
    x = array('d', [0.0]) * len(data)
    w = array('d', [0.0]) * len(data)
    order = array('i', range(0, len(data)))

    order = sorted(order, key=lambda x: data[x])

    if weights == None:
        for i in range(len(data)):
            x[i] = data[order[i]]
            w[i] = 1.0
    else:
        for i in range(len(data)):
            x[i] = data[order[i]]
            w[i] = weights[order[i]]

    # calculate range of K
    # // TODO: should this include weights??
    kmin = kmin if kmin and kmin > 1 else 1
    kmax = kmax if kmax and kmax > 1 else 1

    p = x[0]
    unique = 1
    for q in x[1:]:
        if q != p:
            unique = unique + 1
            q = p

    if kmax < unique:
        kmax = unique

    k, clusters, centers, variance = ckmeans(x, w, kmin, kmax)
    # index := make([]int, len(x))
    # for i := range clusters {
    #     index[order[i]] = clusters[i]
    # }

    # clustered := make([]Cluster, k)

    # for i := 0; i < k; i++ {
    #     clustered[i].Center = centers[i]
    #     clustered[i].Variance = variance[i]
    # }

    # for i, ix := range index {
    #     clustered[ix].Values = append(clustered[ix].Values, data[i])
    # }

    # return clustered
    raise NotImplementedError('NOT IMPLEMENTED YET')


def ckmeans(x, w, kmin, kmax):
    smawk = SMAWK(L2())

    N = len(x)
    S = []
    J = []

    for i in range(0, kmax):
        S.append(array('d', [0.0]) * N)
        J.append(array('i', [0]) * N)

    fill_dp_matrix(x, w, S, J, smawk)

    #     bic := make([]float64, kmax)
    #     kopt := select_levels_weighted(x, w, J, kmin, kmax, bic)
    #
    #     if kopt < kmax {
    #         J = J[0:kopt]
    #     }
    #
    #     clusters := backtrackWeightedX(x, w, J)
    #
    #     // ... calulate mean and variance
    #
    #     centers := make([]float64, kopt)
    #     variance := make([]float64, kopt)
    #     count := make([]int, kopt)
    #     withinss := make([]float64, kopt)
    #     sum := make([]float64, kopt)
    #     sumw := make([]float64, kopt)
    #
    #     for i := range x {
    #         ix := clusters[i]
    #         sum[ix] += x[i] * w[i]
    #         sumw[ix] += w[i]
    #     }
    #
    #     for i := 0; i < kopt; i++ {
    #         centers[i] = sum[i] / sumw[i]
    #     }
    #
    #     for i := range x {
    #         ix := clusters[i]
    #         withinss[ix] += w[i] * (x[i]*x[i] - 2*x[i]*centers[ix] + centers[ix]*centers[ix])
    #         count[ix] += 1
    #     }
    #
    #     for i := 0; i < kopt; i++ {
    #         if count[i] > 1 {
    #             variance[i] = withinss[i] / float64(count[i]-1)
    #         } else {
    #             variance[i] = 0
    #         }
    #     }
    #
    #     return kopt, clusters, centers, variance
    # }

    return 0, [], [], []
