class L2:

    def __init__(self):
        pass

    # Returns ssq(....)
    def dissimilarity(self, j, i, sum_x, sum_x_sq, sum_w, sum_w_sq):
        sji = 0.0

        if sum_w[j] >= sum_w[i]:
            sji = 0.0
        elif j > 0:
            muji = (sum_x[i] - sum_x[j - 1]) / (sum_w[i] - sum_w[j - 1])
            sji = sum_x_sq[i] - sum_x_sq[j - 1] - (sum_w[i] - sum_w[j - 1]) * muji * muji
        else:
            sji = sum_x_sq[i] - sum_x[i] * sum_x[i] / sum_w[i]

        if sji < 0.0:
            sji = 0.0

        return sji
