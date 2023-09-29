from array import array


def fill_dp_matrix(x, w, S, J, smawk):
    K = len(S)
    N = len(S[0])

    sum_x = array('d', [0.0]) * N
    sum_x_sq = array('d', [0.0]) * N
    sum_w = array('d', [0.0]) * len(w)
    sum_w_sq = array('d', [0.0]) * len(w)

    # jseq = []int{}
    shift = x[N // 2]  # median (used to shift the values of x to improve numerical stability)

    sum_x[0] = w[0] * (x[0] - shift)
    sum_x_sq[0] = w[0] * (x[0] - shift) * (x[0] - shift)
    sum_w[0] = w[0]
    sum_w_sq[0] = w[0] * w[0]

    S[0][0] = 0
    J[0][0] = 0

    for i in range(1, N):
        sum_x[i] = sum_x[i - 1] + w[i] * (x[i] - shift)
        sum_x_sq[i] = sum_x_sq[i - 1] + w[i] * (x[i] - shift) * (x[i] - shift)
        sum_w[i] = sum_w[i - 1] + w[i]
        sum_w_sq[i] = sum_w_sq[i - 1] + w[i] * w[i]

        # NTS: using same dissimilarity as SMAWK - original R algorithm potentially (but not really) allowed
        #      for alternative criterion here i.e. not convinced embedding criterion in SMAWK is all that
        #      necessary (and/or correct)
        S[0][i] = smawk.dissimilarity(0, i, sum_x, sum_x_sq, sum_w, sum_w_sq)
        J[0][i] = 0

    for q in range(1, K):
        imin = 0
        if q < K - 1:
            imin = 1
            if q > imin:
                imin = q
        else:
            imin = N - 1

        smawk.fill_row_q_SMAWK(imin, N - 1, q, S, J, sum_x, sum_x_sq, sum_w, sum_w_sq)
