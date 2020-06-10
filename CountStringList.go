package process

func CountStringList(Stringlist []string) map[string]int {
	demo := map[string]int{}
	// 创建空map
	for i := 0; i < (len(Stringlist)); i++ {
		_, ok := demo[Stringlist[i]] //判断a是否存在
		if !ok {
			demo[Stringlist[i]] = 0
		}
	}
	for i := 0; i < (len(Stringlist)); i++ {
		_, ok := demo[Stringlist[i]] //判断a是否存在
		if ok {
			demo[Stringlist[i]] += 1
		}
	}
	return demo
}
