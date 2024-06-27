package RestoreIP_93

import "strconv"

func restoreIpAddresses(s string) []string {
	if s == "" {
		return []string{}
	}

	result, ip := []string{}, []int{}
	dfs(s, 0, ip, &result)

	return result
}

func dfs(s string, idx int, ip []int, result *[]string) {
	if idx == len(s) {
		if len(ip) == 4 {
			*result = append(*result, ip2str(ip))
		}
		return
	}
	if idx == 0 {
		num, _ := strconv.Atoi(string(s[idx]))
		ip = append(ip, num)
		dfs(s, idx+1, ip, result)
	} else {
		num, _ := strconv.Atoi(string(s[idx]))
		next := ip[len(ip)-1]*10 + num
		if next <= 255 && ip[len(ip)-1] != 0 {
			ip[len(ip)-1] = next
			dfs(s, idx+1, ip, result)
			ip[len(ip)-1] /= 10
		}
		if len(ip) < 4 {
			ip = append(ip, num)
			dfs(s, idx+1, ip, result)
			ip = ip[:len(ip)-1]
		}
	}
}

func ip2str(ip []int) string {
	result := strconv.Itoa(ip[0])
	for i := 1; i < len(ip); i++ {
		result += "." + strconv.Itoa(ip[i])
	}
	return result
}
