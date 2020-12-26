$$S_1=\sum_{i=0}^n\sum_{j=0}^ia_ia_j=\sum_{j=0}^n\sum_{i=j}^na_ia_j=\sum_{i=0}^n\sum_{j=i}^na_ia_j$$

$$\sum_{i=0}^2\sum_{j=0}^ia_ia_j=a_0a_0+(a_1a_0+a_1a_1)+(a_2a_0+a_2a_1+a_2a_2)$$

$$\sum_{j=0}^2\sum_{i=j}^2a_ia_j=(a_0a_0+a_1a_0+a_2a_0)+(a_1a_1+a_2a_1)+(a_2a_2)$$

$$\sum_{i=0}^n\sum_{j=0}^ia_ia_j=\frac{1}{2}((\sum_{i=0}^n a_i)^2+(\sum_{i=0}^n a_i^2))$$

#####ex[15]

$$S(n) = 1\times2 + 2\times2^2 + 3\times2^3 +···+ n\times2^n$$

$$2S(n) ={\quad\quad\quad} 1\times2^2 + 2\times2^3 +···+(n-1)\times2^n + n\times2^{n+1} $$

$$-S(n) = 1\times2+1\times2^2+1\times2^3+···+1\times2^n-n\times2^{n+1}$$

$$S(n)=n\times2^{n+1}-(2 + 2^2 + 2^3 +···+ 2^n) \\ {\quad\quad}=n\times2^{n+1}-\frac{2 - 2^n\times2}{1-2} \\ {\quad\quad}=n\times2^{n+1}-(2^{n+1}-2) \\ {\quad\quad}=(n-1)\times2^{n+1} + 2$$

$$O(n)=\sum_{i=1}^n(9\times(\sum_{j=0}^{i-1}(i-j)10^j)+(i+1))$$



$$\prod\limits_{j=2}^n(1-\frac{1}{j^2})$$

$$lg\prod\limits_{j=2}^n(1-\frac{1}{j^2})=\sum\limits_{j=2}^nlg(1-\frac{1}{j^2})$$



$$\sum\limits_{j=1}^n(\frac{x_j^r}{\prod\limits_{1\leq k \leq n, k\neq j}(x_j-x_k)})=\begin{cases}
0,\quad  &如果 \ \ \ 0\leq r < n-1\\
1,\quad  &如果 \ \ \ r = n-1\\
 \sum\limits_{j=1}^nx_j,\quad  &如果 \ \ \ r=n\\
\end{cases}$$



$$\sum\limits_{k=1}^n\frac{\prod_{1\leq r\leq n,r\neq m}(x+k-r)}{\prod_{1\leq r\leq n,r\neq k}(k-r)}=1$$

$$x^{\underline{k}}=\prod\limits_{j=0}^{k-1}(x-j) \\ x^{\overline{k}}=\prod\limits_{j=0}^{k-1}(x+j)$$

$$P_{n(n-1)} = n^{\underline{n-1}} = \prod\limits_{j=0}^{n-2}(n-j)=n(n-1)(n-2)\dots(n-(n-2)) \\ P_{nn} = n^{\underline{n}} = \prod\limits_{j=0}^{n-1}(n-j)=n(n-1)(n-2)\dots(n-(n-2))(n-(n-1)) $$

$$n!=\lim\limits_{m\to\infty}\frac{m^nm!}{(n+1)(n+2)\dots(n+m)}$$

$$\gamma(x)=\frac{x!}{x}=\lim\limits_{m\to\infty}\frac{m^xm!}{x(x+1)(x+2)\dots(x+m)}$$

$$(-z)!\gamma(z) = \frac{\pi}{sin(\pi z)} \quad \frac{1}{\gamma(z)}=\frac{1}{2\pi i}\oint\frac{e^tdt}{t^z}$$



$$\binom{r}{k}=(-1)^k\binom{k-r-1}{k},\quad整数k$$

$$(x+y)^n=\sum\limits_k\binom{n}{k}x(x-kz)^{k-1}(y+kz)^{n-k},\quad整数\ n \geq 0, x\neq0$$

$$\binom{n}{m}=(-1)^{n-m}\binom{-(m+1)}{n-m}$$

$$\sum\limits_{k}\binom{r}{k}\binom{s-kt}{r}(-1)^k=t^r,\quad整数\ r\geq 0$$

$$\sum\limits_{k=0}^{n}\binom{-2}{k}=(-1)^n\lceil{\frac{n+1}{2}}\rceil$$

$$\sum\limits_{k=0}^{n}\binom{m}{k}(k-\frac{m}{2})=-\frac{m}{2}\binom{m-1}{n} $$

$$\binom{n}{1}+\binom{n}{4}+\binom{n}{7}+\dots=\frac{1}{3}(2^n+2cos\frac{(n-2)\pi}{3})$$

$$(1+x)(1+qx)\dots(1+q^{n-1}x)=\sum\limits_k\binom{n}{k}_qq^{k(k-1)/2}x^k$$

$$H_n=1+\frac{1}{2}+\frac{1}{3}+\dots+\frac{1}{n}=\sum\limits_{k=1}^n\frac{1}{k},\quad n\geq 0 $$

$$H_\infty^{(r)}=\zeta(r)=\sum\limits_{k\ge1}\frac{1}{k^r}$$







