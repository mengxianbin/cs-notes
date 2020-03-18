* 编码平面
    * U+0000 ~ U+10FFFF
    * 65536 * 17
    * Plane 0 ～ Plane 17
    * BMP
        * Plane 0
        * Basic Multilingual Plane
        * 基本多语言平面
        * U+0000 ~ U+FFFF
    * SP
        * Supplementary Planes
        * 增补平面

* 代理区
    * Surrogate Area
    * D8~DF
        * D800–DBFF 属于高代理区（High Surrogate Area）
        * DC00–DFFF 属于低代理区（Low Surrogate Area）
        * 4 * 256 = 1024
        * 1024 * 2 = 2048
        * 1024 * 1024 = 2 ^ 20 = 16 * 2 ^ 16
    * 代理对 
        * Surrogate Pair

* UTF-16
    * 代理区

* UTF-32
    * 直接编解码

* UTF-8
    * 高位无重叠
        * 便于搜索、匹配
    * 多字节高位固定模版
        * 首字节
            * 0xxxxxxx
            * 110xxxxx
            * 1110xxxx
            * 11110xxx
        * 跟随字节
            * 10xxxxxx
    * 低位由 UTF-16 二进制依次填入
