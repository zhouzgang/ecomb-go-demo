package cfg

const DefaultLocale = "en"

// 自定义类型
type KV map[string]string

type I18Config map[string]KV

type I18Configs map[string]I18Config

var I I18Configs

func (i I18Configs) Get(locale, sort, key string) string {
	if v, ok := i[locale]; ok {
		if v2, ok := v[sort]; ok {
			if v3, ok := v2[key]; ok {
				return v3
			}
		}
	}
	return ""
}

/**
	golang 执行顺序：import --> const --> var --> init() --> main()
	![go执行顺序](doc/images/go执行顺序.png)
 */
func init() {
	/**
	make 的作用是初始化内置的数据结构，也就是我们在前面提到的切片、哈希表和 Channel；
	new 的作用是根据传入的类型分配一片内存空间并返回指向这片内存空间的指针；
	 */
	I = make(I18Configs)
}

func LoadI18n(dir string, locales []string) {

}