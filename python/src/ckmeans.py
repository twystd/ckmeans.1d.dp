'''
Port of Ckmeans.1d.dp_4.3.3 from R to Python.
'''

from array import array


def ckmeans_1d_dp(data, weights):
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

    print(x[0:16])
    # calculate range of K
    # // TODO: should this include weights??
    kmin = 1
    kmax = 1

    p = x[0]
    for q in x[1:]:
        if q != p:
            kmax += 1
            q = p

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
    return 0, [], [], []
