package corealgo

// --- 数论 ---

// SieveOfEratosthenes 埃拉托斯特尼筛法：返回 [2, n] 内的全部素数（升序）
// 思路：初始化 [2, n] 全部标记为素数，从 2 开始，若 i 为素数则将其所有
// 倍数（从 i*i 起）标记为合数；筛完后收集所有仍为素数的数
// 时间复杂度：O(n log log n)
// 空间复杂度：O(n)
// 适用场景：
// 1. 素数表、素数判定的预处理
// 2. 分解质因数前的素数收集
// 3. 加密算法、哈希表容量选择中的素数查找
func SieveOfEratosthenes(n int) []int {
	if n < 2 {
		return []int{}
	}
	isPrime := make([]bool, n+1)
	for i := 2; i <= n; i++ {
		isPrime[i] = true
	}
	for i := 2; i*i <= n; i++ {
		if isPrime[i] {
			for j := i * i; j <= n; j += i {
				isPrime[j] = false
			}
		}
	}
	primes := make([]int, 0, n/10+1)
	for i := 2; i <= n; i++ {
		if isPrime[i] {
			primes = append(primes, i)
		}
	}
	return primes
}

// GCD 最大公约数（欧几里得算法，辗转相除法）
// 思路：gcd(a, b) = gcd(b, a%b)，不断取余直到余数为 0，此时的除数即为
// 最大公约数。负数先取绝对值处理
// 时间复杂度：O(log min(a, b))
// 空间复杂度：O(1)
// 适用场景：
// 1. 分数约分化简
// 2. 求最小公倍数 lcm(a, b) = a / gcd(a, b) * b
// 3. 扩展欧几里得、同余方程等数论基础
func GCD(a, b int) int {
	if a < 0 {
		a = -a
	}
	if b < 0 {
		b = -b
	}
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// PowMod 快速幂取模：返回 base^exp mod mod
// 思路：将 exp 按二进制展开，base 逐次平方，仅当 exp 当前位为 1 时累乘，
// 把 O(exp) 次乘法降为 O(log exp) 次，并在每步取模防止中间结果溢出
// 时间复杂度：O(log exp)
// 空间复杂度：O(1)
// 适用场景：
// 1. 密码学中的模幂运算（RSA 加解密）
// 2. 大数幂运算取模，避免直接计算溢出
// 3. 组合数学、概率论中的模运算
// 注意：mod 为 0 时返回 0；base 为负数时先取模转正；mod 很大（>1e9 量级）
// 时乘法中间结果可能溢出 int64，超大模数请改用 math/big
func PowMod(base, exp, mod int64) int64 {
	if mod == 0 {
		return 0
	}
	result := int64(1)
	b := base % mod
	if b < 0 {
		b += mod
	}
	for exp > 0 {
		if exp&1 == 1 {
			result = result * b % mod
		}
		b = b * b % mod
		exp >>= 1
	}
	return result
}

// Combination 组合数 C(n, k)
// 思路：采用乘法除法交替的公式 C(n, k) = Π(i=1..k) (n-k+i) / i 计算，
// 并取 k = min(k, n-k) 缩小计算规模
// 防溢出：每步"先乘后除"，中间结果恰为 C(n-k+i, i)，是整数且整体不会
// 超过最终结果的数量级；只要最终结果在 int64 范围内即可安全计算
// （例如 C(60, 30) ≈ 1.18e17 可安全计算）
// 时间复杂度：O(k)
// 空间复杂度：O(1)
// 适用场景：
// 1. 抽奖、抽样的可能性计数
// 2. 概率计算（二项分布）
// 3. 排列组合类计数问题
// 注意：k < 0 或 k > n 时返回 0
func Combination(n, k int) int {
	if k < 0 || k > n {
		return 0
	}
	if k > n-k {
		k = n - k
	}
	var res int64 = 1
	for i := 1; i <= k; i++ {
		res = res * int64(n-k+i) / int64(i)
	}
	return int(res)
}
