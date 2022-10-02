# LaTeX

[https://en.wikibooks.org/wiki/LaTeX/Mathematics](https://en.wikibooks.org/wiki/LaTeX/Mathematics)

## Symbols

* $ + - \cdot \times \div = ! / ( ) [ ] < > | ' : * \leftarrow \Rightarrow $
* $\forall x \in X, \quad \exists y \leq \epsilon$

## Greek letters

* $\alpha, \Alpha, \beta, \Beta, \gamma, \Gamma, \pi, \Pi, \phi, \varphi, \mu, \Phi$

## Operators

* $\cos (2\theta) = \cos^2 \theta - \sin^2 \theta$
* $\lim\limits_{x \to \infty} \exp(-x) = 0$
* $a \bmod b$
* $x \equiv a \pmod b$
* $\cfrac {dy} {dx}$ vs $\cfrac {\mathrm d y} {\mathrm d x}$ vs $\cfrac {\operatorname d y} {\operatorname d x}$ vs $\cfrac {\operatorname d \operatorname y} {\operatorname d \operatorname x}$

## Powers and indices

* $k_{n+1} = n^2 + k_n^2 - k_{n-1}$
* $f(n) = n^5 + 4n^2 +2 |_{n=17}$
* $\cfrac{n!}{k!(n-k)} = \displaystyle\binom{n}{k}$
* $\cfrac {\frac{1}{x} + \frac{1}{y}} {y-z}$
* $^3/_7$
* $ 3 \times \frac{1}{2} = 1^1/_2$
* $x^ {\frac {1} {2}}$

## Continued fractions

* $x = a_0 +
  \cfrac {1} {a_1 +
  \cfrac {1} {a_2 +
  \cfrac {1} {a_3 +
  \cfrac {1} {a_4}}}}
  $

## Multiplication of two numbers

* $\cfrac {
  \begin{array} {r}
  (x_1x_2) \\
  \times (x'_1x'_2)
  \end{array}
  } {
  (y_1y_2y_3y_4)
  }
  $

## Roots

* $\sqrt{\cfrac{a}{b}}$
* $\sqrt[n]{1+x+x^2+x^3+\dots+x^n}$

## Sums, products and integrals

* $\sum_{i=1}^{10} t_i$    vs     $\displaystyle\sum_{i=1}^{10} t_i$    vs     $\sum\limits_{i=1}^{10} t_i$
* $\displaystyle\int_0^\infty \mathrm{e}^{-x}\, \mathrm{d}x$
* $\displaystyle\sum_{\substack{
  0<i<m \\
  0<j<n
  }}
  P(i, j)
  $
* $\int_a^b$    vs     $\int\limits_a^b$    vs    $\displaystyle\int_a^b$
* $\displaystyle \prod _{x_m} ^{x_n} = {x_m} $

## Brackets, braces and delimiters

* $( a ), [ b ], \{ c \}, | d |, \| e \|, 
  \langle f \rangle, \lfloor g \rfloor, 
  \lceil h \rceil, \ulcorner i \urcorner, 
  / j \backslash$

## Automatic sizing

* $\left(\cfrac{x^2}{y^3}\right)$
* $P\left(A=2\middle|\frac{A^2}{B}>4\right)$
* $\left\{\cfrac{x^2}{y^3}\right\}$

- $\cfrac{x^3}{3}|_0^1$     vs    $\left.\cfrac{x^3}{3}\right|_0^1$

## Manual sizing

* $( \big( \Big( \bigg( \Bigg($
* $\cfrac{\mathrm d}{\mathrm d x} \left( k g(x) \right)$    vs    $\cfrac{\mathrm d}{\mathrm d x} \Bigg( k g(x) \Bigg)$

## Matrices and array

* $\left[
  \begin{matrix}
  a & b & c \\
  d & e & f \\
  g & h & i
  \end{matrix}
  \right]$
* $\begin{matrix}
  -1 & 3 \\
  2 & -4
  \end{matrix}
  $
* $A_{m, n} = 
  \begin{pmatrix}
  a_{1, 1} & a_{1, 2} & \cdots & a_{1, n} \\
  a_{2, 1} & a_{2, 2} & \cdots & a_{2, n} \\
  \vdots  & \vdots  & \ddots & \vdots  \\
  a_{m, 1} & a_{m, 2} & \cdots & a_{m, n} 
  \end{pmatrix}$
* $\begin{array}{c|c}
  1 & 2 \\ 
  \hline
  3 & 4
  \end{array}$
* $M = \begin{bmatrix}
  \frac{5}{6} & \frac{1}{6} & 0           \\[0.5em]
  \frac{5}{6} & 0           & \frac{1}{6} \\[0.5em]
  0           & \frac{5}{6} & \frac{1}{6}
  \end{bmatrix}$

## Adding text to equations

* $ 50 apples \times 100 apples = lots of apples^2$
* $ 50 \text{apples} \times 100 \text{apples} = \text{lots of apples}^2$
* $ 50 \text{ apples} \times 100 \text{ apples} = \text{lots of apples}^2$

## Formatted text

- $ 50 \textrm{ apples} \times 100 \textbf{ apples} = \textit{lots of apples}^2$
  - `\textrm:` roman font family
  - `\textbf:` bold
  - `\textit:` italic shape
  - **...**

## Accents

* $a' a'' a''' a'''$
* $\bar{a} \quad \acute{a} \quad \check{a} \quad \grave{a}$
* $\vec {a}$    $\overrightarrow {AB}$    $\overleftarrow {AB}$

## Color

* $k = {\color{red}x} \mathbin{\color{yellow}-} 2$

## Plus and minus signs

* `\pm` $\pm$
* `\mp` $\mp$

## Controlling horizontal spacing

- $f(n) =
  \begin{cases}
  n/2       & \quad \text{if } n \text{ > 0}\\
  5         & \quad \text{if } n \text{ = 0} \\
  -(n+1)/2  & \quad \text{if } n \text{ < 0}
  \end{cases}
  $

## Dots

* $ \dot A$
* $A \cdot B$
* $A \dots B$
* $A \ldots B$
* $A \cdots B$
* $A \dotsm B$
* $A \vdots B$
* $A \ddots B$

---

## Excecises

* $E=mc^2$
* $x = \frac {-b \pm \sqrt {\Delta}} {2a} \quad\quad (\Delta = b^2 - 4ac)$
* $a^2 + b^2 = c^2$
* $d_{circle} = 2\pi r$
* $S_{circle} = \pi r^2$
* $V_{sphere} = \frac{4}{3} \pi r^3$

---
