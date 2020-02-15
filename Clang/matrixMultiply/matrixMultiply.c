#include "matrixMultiply.h"

void main()
{
}

int memoizedMatrixChain(int *p, int n, int **m, int **s, int i, int j)
{
    if (i == j)
    {
        return 0;
    }
    if (m[i][j] != 0)
    {
        return m[i][j];
    }
    int u = memoizedMatrixChain(p, n, m, s, i, i) + memoizedMatrixChain(p, n, m, s, i + 1, j) + p[i - 1] * p[i] * p[j];
    s[i][j] = i;
    for (int k = i + 1; k < j; k++)
    {
        int t = memoizedMatrixChain(p, n, m, s, i, k) + memoizedMatrixChain(p, n, m, s, k + 1, j) + p[i - 1] * p[k] * p[j];
        if (t < u)
        {
            u = t;
            s[i][j] = k;
        }
    }
    m[i][j] = u;
    return u;
}

int recurMatrixMutiplu(int *p, int i, int j)
{
    if (i == j)
    {
        return 0;
    }
    int u = recurMatrixMutiplu(p, i, i) + recurMatrixMutiplu(p, i + 1, j) + p[i - 1] * p[i] * p[j];
    for (int z = i + 1; z < j; z++)
    {
        int tmp = recurMatrixMutiplu(p, i, z) + recurMatrixMutiplu(p, z + 1, j) + p[i - 1] * p[z] * p[j];
        if (tmp < u)
        {
            u = tmp;
        }
    }
    return u;
}

void matrixMultiply(int *p, int **m, int n, int **s)
{
    for (int i = 0; i <= n; i++)
    {
        m[i][i] = 0;
    }
    for (int r = 2; r <= n; r++)
    {
        for (int i = 1; i <= n - r + 1; i++)
        {
            int j = i + r - 1;
            m[i][j] = m[i + 1][j] + p[i - 1] * p[i] * p[j];
            s[i][j] = i;
            for (int k = i + 1; k < j; k++)
            {
                int t = m[i][k] + m[k + 1][j] + p[i - 1] * p[k] * p[j];
                if (t < m[i][j])
                {
                    m[i][j] = t;
                    s[i][j] = k;
                }
            }
        }
    }
}