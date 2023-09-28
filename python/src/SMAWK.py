from array import array

class SMAWK:

    def __init__(self, criterion):
        self.criterion = criterion

    def dissimilarity(self, j, i, sum_x, sum_x_sq, sum_w, sum_w_sq):
        return self.criterion.dissimilarity(j, i, sum_x, sum_x_sq, sum_w, sum_w_sq)

    def fill_row_q_SMAWK(self, imin, imax, q, S, J, sum_x, sum_x_sq, sum_w, sum_w_sq):
        js = array('i', [0]) * (imax - q + 1)
        absv = q

        for i in range(0, len(js)):
            js[i] = absv
            absv = absv + 1

        # TODO 
        # self.smawk(imin, imax, 1, q, js, S, J, sum_x, sum_x_sq, sum_w, sum_w_sq)

    def smawk(self, imin, imax, istep, q, js, S, J, sum_x, sum_x_sq, sum_w, sum_w_sq):
        if imax - imin <= 0 * istep:
            self.find_min_from_candidates(imin, imax, istep, q, js, S, J, sum_x, sum_x_sq, sum_w, sum_w_sq)
        else:
            js_odd = array('i', [0]) * len(js)

        self.reduce_in_place(imin, imax, istep, q, js, js_odd, S, J, sum_x, sum_x_sq, sum_w, sum_w_sq)

        istepx2 = istep << 1
        imin_odd = imin + istep
        imax_odd = imin_odd + (imax - imin_odd) / istepx2 * istepx2

        self.smawk(imin_odd, imax_odd, istepx2, q, js_odd, S, J, sum_x, sum_x_sq, sum_w, sum_w_sq)
        self.fill_even_positions(imin, imax, istep, q, js, S, J, sum_x, sum_x_sq, sum_w, sum_w_sq)

    def fill_even_positions(self, imin, imax, istep, q, js, S, J, sum_x, sum_x_sq, sum_w, sum_w_sq):
        n = len(js)
        istepx2 = istep << 1
        jl = js[0]
        r = 0

        for i in range(imin, imax + 1, istepx2):
            while js[r] < jl:
                r = r + 1

            S[q][i] = S[q - 1][js[r] - 1] + self.dissimilarity(js[r], i, sum_x, sum_x_sq, sum_w, sum_w_sq)
            J[q][i] = js[r]

            # Look for minimum S upto jmax within js
            jh = 0

            if i + istep <= imax:
                jh = J[q][i + istep]
            else:
                jh = js[n - 1]

            jmax = jh
            if i < jmax:
                jmax = i

            sjimin = self.dissimilarity(jmax, i, sum_x, sum_x_sq, sum_w, sum_w_sq)

            r = r + 1
            while r < n and js[r] <= jmax:
                jabs = js[r]

                if jabs > i:
                    break

                if jabs >= J[q - 1][i]:
                    s = self.dissimilarity(jabs, i, sum_x, sum_x_sq, sum_w, sum_w_sq)
                    Sj = S[q - 1][jabs - 1] + s

                    if Sj <= S[q][i]:
                        S[q][i] = Sj
                        J[q][i] = js[r]
                    elif S[q - 1][jabs - 1] + sjimin > S[q][i]:
                        break

                    r = r + 1

        r = r - 1
        jl = jh

    def reduce_in_place(self, imin, imax, istep, q, js, js_red, S, J, sum_x, sum_x_sq, sum_w, sum_w_sq):
        N = (imax - imin) // istep + 1

        for i in range(0, len(js)):
            js_red[i] = js[i]

        if N >= len(js):
            return

        # Two positions to move candidate j's back and forth
        left = -1  # points to last favorable position / column
        right = 0  # points to current position / column

        m = len(js_red)

        while m > N:  # js_reduced has more than N positions / columns
            p = left + 1

            i = imin + p * istep
            j = js_red[right]
            Sl = S[q - 1][j - 1] + self.dissimilarity(j, i, sum_x, sum_x_sq, sum_w, sum_w_sq)

            jplus1 = js_red[right + 1]
            Slplus1 = S[q - 1][jplus1 - 1] + self.dissimilarity(jplus1, i, sum_x, sum_x_sq, sum_w, sum_w_sq)

            if Sl < Slplus1 and p < N - 1:
                left = left + 1
                js_red[left] = j
                right = right + 1
            elif Sl < Slplus1 and p == N - 1:
                right = right + 1
                js_red[right] = j
                m = m - 1
            else:
                if p > 0:
                    js_red[right] = js_red[left]
                    left = left - 1
                else:
                    right = right + 1
            m = m - 1

        for r in range(left + 1, m):
            js_red[r] = js_red[right]
            right = right + 1

        # ??? leftover weirdness from 'R' probably
        tmp = array('i', [0]) * m
        for i in range(0, m):
            tmp[i] = js_red[i]

        js_red = tmp

    def find_min_from_candidates(self, imin, imax, istep, q, js, S, J, sum_x, sum_x_sq, sum_w, sum_w_sq):
        rmin_prev = 0

        for i in range(imin, imax + 1, istep):
            rmin = rmin_prev

        S[q][i] = S[q - 1][js[rmin] - 1] + self.dissimilarity(js[rmin], i, sum_x, sum_x_sq, sum_w, sum_w_sq)
        J[q][i] = js[rmin]

        for r in range(rmin + 1, len(js)):
            j_abs = js[r]

            if j_abs < J[q - 1][i]:
                r = r + 1
                continue

            if j_abs > i:
                break

            Sj = (S[q - 1][j_abs - 1] + self.dissimilarity(j_abs, i, sum_x, sum_x_sq, sum_w, sum_w_sq))
            if Sj <= S[q][i]:
                S[q][i] = Sj
                J[q][i] = js[r]
                rmin_prev = r
