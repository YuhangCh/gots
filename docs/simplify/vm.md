# 维斯瓦林甘算法

> org.locationtech.jts.simplify.VWLineSimplifier

相比较而言，道格拉斯-普克算法有更高的知名度，但维斯瓦林甘算法（以下简称VM算法）的效率更高且容易理解。

它依据最小的变化逐步精简点，最终甚至能压缩到为原体积的5%。

为了判断哪个点在精简后使整体视觉上来看具有最小的变化，VM算法依次遍历一整条线上的连续的三个点，整条线上三个点相关联的三角形面积最小的点会被去除。每一次精简，被去除的点旁边的两个点相关联的三角形面积会被重新计算，并重复执行该步骤。

## 参考

[Line Simplification](https://bost.ocks.org/mike/simplify/)

