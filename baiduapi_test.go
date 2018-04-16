// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package wxspider

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"
	// "golang.org/x/net/html/charset"
)

func Test_XTags(t *testing.T) {
	tags := AiTags{}
	tt := `{"log_id": 7971186989667946038}`
	rbyte := []byte(tt)

	json.Unmarshal(rbyte, &tags)
	t.Fatal(tags)
}

func Test_XCategories(t *testing.T) {
	tags := AiCategories{}
	tt := `{"log_id": 8163873610937095637, "item": {"lv2_tag_list": [{"score": 0.877436, "tag": "足球"}, {"score": 0.793682, "tag": "国际足球"}, {"score": 0.775911, "tag": "英超"}], "lv1_tag_list": [{"score": 0.824329, "tag": "体育"}]}}`
	rbyte := []byte(tt)

	json.Unmarshal(rbyte, &tags)
	t.Fatal(tags)
}
func Test_GetTags(t *testing.T) {

	tt := "你若记得，他便不悔"
	cc := `2001年4月1日，飞行员王伟因捍卫祖国领空而牺牲。17年，他的儿子已经长大，妻子不再年轻；17年，中国的海空实力今非昔比。对于英雄的思念，从未停止——17年里的每一天，亲人都在等他团聚、战友都在等他仗剑、中华儿女都在等他共同见证祖国的强大。《我和王伟的爱情故事》（节选）作者/王伟的妻子：阮国琴也许，爱上他就在这一瞬间。1986年，我们高中毕业前夕，空军飞行学院在湖州招收飞行学员。这时我给他写了一封信，信中有这么一句话：“是一名真正的男子汉，你就去蓝天翱翔吧！”王伟说：“我的人生第一是飞行，第二是我们之间的爱情，它们对我来讲，就像飞机的两翼缺一不可。我一定要飞出来，否则决不回来见你！”△中学时期的王伟“感情上，爱你胜过爱飞行，而理智上要我爱飞行比爱你更甚。”在此后的近三年的时间里，我们没有再见面。毕业后的王伟并没有立即回到分别了两年半的我的身边，而是直接去了海军航空兵的训练基地报到。他写信告诉我：“此生唯有飞行和你左右我。你能给我最大的幸福和快乐，飞行也能。我爱飞行事业，同样爱你。感情上，爱你胜过爱飞行，而理智上要我爱飞行比爱你更甚——这就是我目前（也许是一生）的主要矛盾，以及它们的辩证关系。”△王伟和妻子收到这样的来信，我觉得我爱的王伟成熟了。我在给王伟的信中写到：“在我的心中，一个热爱飞行事业、刚毅、坚强、奋斗的王伟，他吃得起苦、受得起累，他自信地向前走，这种对目标的追求和未来美好的渴望，使我热爱他。如此一个王伟，使我不在乎他对我关心有多少，不在乎他跟我在一起的时间有多长，不在乎别人一对对卿卿我我的甜蜜感受，不在乎他不能在我身边的孤独和寂寞，在乎的是他对未来的追求和生命的热爱。”当年江南雨巷中徘徊的少年终于成长为搏击蓝天的雄鹰了。向往蓝天的王伟是一个浪漫的人，我们的家中墙上挂着的是他亲手画的油画。一次，王伟画了一幅中国海军航空兵驾驶着未来的新型战机从战舰上起飞的油画挂在了家中的墙上，并对前来家中做客的团政委讲：“画上那名飞行员就是我自己！”△王伟的油画遗作。我知道，这是他的追求与梦想！看着英姿勃勃的丈夫，我在想：部队真是锻炼人，当年在江南雨巷中徘徊的那个少年终于成长为搏击蓝天的雄鹰了。我觉得好幸福，在空旷的机场跑道旁，我和王伟手挽着手，漫步在一望无际的天地，享受落日的霞光……△王伟油画遗作。最后的告别，刻骨铭心。2001年4月17日，大海肃穆。人民海军以最隆重的方式——海祭，为我的丈夫王伟送行……战舰在微微的海风中起航，缓缓驶向那朵洁白的伞花飘落的海域……△阮国琴把花环抛进大海，祭奠自己的丈夫。我靠在甲板右舷，已经没有了泪水，面朝大海，仰望着蓝天上飘动的白云，在心中对丈夫默默地诉说着自己的思念和悲伤：阿伟啊，记得你曾对我说：你是天上的风筝，飞得再高再远，线还在我手中攥着，只要我不松手，不管飞到哪里，你都会回到我的身边，可我没松手、真的没有松手啊，你怎么就飞走了呢？△导弹护卫舰官兵脱帽肃立，为英雄王伟送行。王伟离开我十七年来，我有太多的话想对他诉说，我把这些话写成了歌词《等你回家》：记得那年春暖花烂漫你匆匆离家没顾上给我说再见知道你驾驶战鹰又去那片天那矫健的英姿映满我眼帘等你回家，为你沏好家乡的新茶一遍又一遍等你回家，插满鲜花的房间芳华溢满等你回家，等着给你一个长长的吻等你回家，我们手挽着手去看晚霞漫天记得那年春暖花烂漫你匆匆离家没顾上给我说再见知道你肩负职责生死一瞬间那坚毅的背影深深刻在我心田等你回家，为你采撷满山的红豆一年又一年等你回家，儿子已长成英俊的男子汉等你回家，等到我两鬓霜花飘满等你回家，望眼欲穿哪怕海枯石又烂英雄，危难中孕育，奉献中诞生。他们奉献着、燃烧着，我们感动着，铭记着。01张超，海军航母舰载战斗机一级飞行员。2016年4月27日，他在驾驶歼-15飞机进行陆基模拟着舰训练时，突遇电传故障。危机关头，他果断处置，尽最大努力保住战机，也正因此，张超错过了最佳跳伞时机，最终不幸牺牲，年仅29岁。他是为中国航母舰载机事业殉职的第一位飞行员。认识六年，结婚四年多，女儿两岁三个月。我们在一起的日子数得太清楚，后悔没有早点来陪你，直到现在，我还是不想相信。这两天，我看着你瘦了，黑了，我不敢太伤心，我舍不得你走，也知道你舍不得走。看见你还是那么帅，你那么年轻，我真的好怕自己老到你不认识我，我该怎么赴约！我仍然觉得，你只是太累了，睡着了。——张超的妻子 张亚02危难面前，维和战士申亮亮作出了另一种选择。西非马里，当地时间2016年5月31日，一辆载有500多公斤炸药的汽车，驶入申亮亮所在维和部队的营区。申亮亮发出预警信号，引导后方部队隐蔽。在汽车侧翻时，本可以选择撤离现场，但年仅29岁的他用生命中的最后37秒，换回了部队其他人员的平安。好弟弟，忠孝不能两全，你为国家尽忠了，父母跟前的这些孝，姐姐和哥哥替你完成。你在天堂那边，自己照顾好自己就行了，也要好好的。一家人永远都会想你的，放心吧。——申亮亮的姐姐 申海霞03余旭，中国首位歼-10女飞行员，也是中国首批能够驾驶三代战机执行任务的女飞行员之一。2016年11月12日，她所在的八一飞行表演队在河北省唐山市玉田县进行飞行训练中发生一等事故，余旭跳伞失败，壮烈牺牲。没有你的日子，一天一天过得特别长。我总想在梦里抱抱你，和你说说话，问你为什么这么狠心离开我们，可是我咋个都梦不到你。旭儿，你从小就匪，这次真的跑得太远了呀。那些报道上都说你是英雄，我们是伟大的父母。可旭儿，我从没想过这辈子会和伟大两个字连在一起。你晓不晓得每次看你飞行，我和你爸爸的心都是悬着的，什么激动啊、骄傲啊，都是你平安落地后才会有的感觉。那年劝说你停飞，你哭着说不愿意，我就晓得，我这个女儿，是要嫁给飞机了。——余旭的母亲 胡中秋042007年11月30日，一名军人感动了一座城，感动了全中国。年仅28岁的火箭军某部中尉参谋孟祥斌，为救轻生女青年，从浙江金华城南桥头纵身跃入婺江壮烈牺牲，被中央军委追授“舍己救人模范军官”荣誉称号，当选感动中国2007年度人物。十年前我带孩子去部队探亲，你用一次平凡的跳跃为自己的人生与军旅生涯画上一个圆！那刻我在悲伤中寻找活下去的理由，我在绝望中追问生命的意义，我在反思中思考，你平凡的一跃后，作为妻子的我如何将平凡继续？也就在那天整理你遗物时看到本子上这样一句话———认认真真做人，踏踏实实做事！——孟祥斌的妻子 叶庆华资料来源/老海军 军报记者等图/视觉中国 军报记者所谓英雄，不关乎时代，不关乎环境。你若感恩，他便值得，△中央广播电视总台7位播音员、主持人诵读秋瑾、夏明翰、赵一曼、黄继光、雷锋、孔繁森、焦裕禄等先烈的家书或诗歌，追思英雄，铭记他们的牺牲和奉献。点击「写留言」表达追思`

	// t.Fatal(strings.Count(tt, "") - 1)
	a := &Article{
		Title: tt,
		Cont:  cc,
	}
	t.Fatal(a.AiGetTags())
}
func Test_GetCategories(t *testing.T) {

	// tt := "欧洲冠军联赛"
	// cc := `欧洲冠军联赛是欧洲足球协会联盟主办的年度足球比赛，代表欧洲俱乐部足球最高荣誉和水平，被认为是全世界最高素质、最具影响力以及最高水平的俱乐部赛事，亦是世界上奖金最高的足球赛事和体育赛事之一。`

	tt := "你若记得，他便不悔"
	cc := `2001年4月1日，飞行员王伟因捍卫祖国领空而牺牲。17年，他的儿子已经长大，妻子不再年轻；17年，中国的海空实力今非昔比。对于英雄的思念，从未停止——17年里的每一天，亲人都在等他团聚、战友都在等他仗剑、中华儿女都在等他共同见证祖国的强大。《我和王伟的爱情故事》（节选）作者/王伟的妻子：阮国琴也许，爱上他就在这一瞬间。1986年，我们高中毕业前夕，空军飞行学院在湖州招收飞行学员。这时我给他写了一封信，信中有这么一句话：“是一名真正的男子汉，你就去蓝天翱翔吧！”王伟说：“我的人生第一是飞行，第二是我们之间的爱情，它们对我来讲，就像飞机的两翼缺一不可。我一定要飞出来，否则决不回来见你！”△中学时期的王伟“感情上，爱你胜过爱飞行，而理智上要我爱飞行比爱你更甚。”在此后的近三年的时间里，我们没有再见面。毕业后的王伟并没有立即回到分别了两年半的我的身边，而是直接去了海军航空兵的训练基地报到。他写信告诉我：“此生唯有飞行和你左右我。你能给我最大的幸福和快乐，飞行也能。我爱飞行事业，同样爱你。感情上，爱你胜过爱飞行，而理智上要我爱飞行比爱你更甚——这就是我目前（也许是一生）的主要矛盾，以及它们的辩证关系。”△王伟和妻子收到这样的来信，我觉得我爱的王伟成熟了。我在给王伟的信中写到：“在我的心中，一个热爱飞行事业、刚毅、坚强、奋斗的王伟，他吃得起苦、受得起累，他自信地向前走，这种对目标的追求和未来美好的渴望，使我热爱他。如此一个王伟，使我不在乎他对我关心有多少，不在乎他跟我在一起的时间有多长，不在乎别人一对对卿卿我我的甜蜜感受，不在乎他不能在我身边的孤独和寂寞，在乎的是他对未来的追求和生命的热爱。”当年江南雨巷中徘徊的少年终于成长为搏击蓝天的雄鹰了。向往蓝天的王伟是一个浪漫的人，我们的家中墙上挂着的是他亲手画的油画。一次，王伟画了一幅中国海军航空兵驾驶着未来的新型战机从战舰上起飞的油画挂在了家中的墙上，并对前来家中做客的团政委讲：“画上那名飞行员就是我自己！”△王伟的油画遗作。我知道，这是他的追求与梦想！看着英姿勃勃的丈夫，我在想：部队真是锻炼人，当年在江南雨巷中徘徊的那个少年终于成长为搏击蓝天的雄鹰了。我觉得好幸福，在空旷的机场跑道旁，我和王伟手挽着手，漫步在一望无际的天地，享受落日的霞光……△王伟油画遗作。最后的告别，刻骨铭心。2001年4月17日，大海肃穆。人民海军以最隆重的方式——海祭，为我的丈夫王伟送行……战舰在微微的海风中起航，缓缓驶向那朵洁白的伞花飘落的海域……△阮国琴把花环抛进大海，祭奠自己的丈夫。我靠在甲板右舷，已经没有了泪水，面朝大海，仰望着蓝天上飘动的白云，在心中对丈夫默默地诉说着自己的思念和悲伤：阿伟啊，记得你曾对我说：你是天上的风筝，飞得再高再远，线还在我手中攥着，只要我不松手，不管飞到哪里，你都会回到我的身边，可我没松手、真的没有松手啊，你怎么就飞走了呢？△导弹护卫舰官兵脱帽肃立，为英雄王伟送行。王伟离开我十七年来，我有太多的话想对他诉说，我把这些话写成了歌词《等你回家》：记得那年春暖花烂漫你匆匆离家没顾上给我说再见知道你驾驶战鹰又去那片天那矫健的英姿映满我眼帘等你回家，为你沏好家乡的新茶一遍又一遍等你回家，插满鲜花的房间芳华溢满等你回家，等着给你一个长长的吻等你回家，我们手挽着手去看晚霞漫天记得那年春暖花烂漫你匆匆离家没顾上给我说再见知道你肩负职责生死一瞬间那坚毅的背影深深刻在我心田等你回家，为你采撷满山的红豆一年又一年等你回家，儿子已长成英俊的男子汉等你回家，等到我两鬓霜花飘满等你回家，望眼欲穿哪怕海枯石又烂英雄，危难中孕育，奉献中诞生。他们奉献着、燃烧着，我们感动着，铭记着。01张超，海军航母舰载战斗机一级飞行员。2016年4月27日，他在驾驶歼-15飞机进行陆基模拟着舰训练时，突遇电传故障。危机关头，他果断处置，尽最大努力保住战机，也正因此，张超错过了最佳跳伞时机，最终不幸牺牲，年仅29岁。他是为中国航母舰载机事业殉职的第一位飞行员。认识六年，结婚四年多，女儿两岁三个月。我们在一起的日子数得太清楚，后悔没有早点来陪你，直到现在，我还是不想相信。这两天，我看着你瘦了，黑了，我不敢太伤心，我舍不得你走，也知道你舍不得走。看见你还是那么帅，你那么年轻，我真的好怕自己老到你不认识我，我该怎么赴约！我仍然觉得，你只是太累了，睡着了。——张超的妻子 张亚02危难面前，维和战士申亮亮作出了另一种选择。西非马里，当地时间2016年5月31日，一辆载有500多公斤炸药的汽车，驶入申亮亮所在维和部队的营区。申亮亮发出预警信号，引导后方部队隐蔽。在汽车侧翻时，本可以选择撤离现场，但年仅29岁的他用生命中的最后37秒，换回了部队其他人员的平安。好弟弟，忠孝不能两全，你为国家尽忠了，父母跟前的这些孝，姐姐和哥哥替你完成。你在天堂那边，自己照顾好自己就行了，也要好好的。一家人永远都会想你的，放心吧。——申亮亮的姐姐 申海霞03余旭，中国首位歼-10女飞行员，也是中国首批能够驾驶三代战机执行任务的女飞行员之一。2016年11月12日，她所在的八一飞行表演队在河北省唐山市玉田县进行飞行训练中发生一等事故，余旭跳伞失败，壮烈牺牲。没有你的日子，一天一天过得特别长。我总想在梦里抱抱你，和你说说话，问你为什么这么狠心离开我们，可是我咋个都梦不到你。旭儿，你从小就匪，这次真的跑得太远了呀。那些报道上都说你是英雄，我们是伟大的父母。可旭儿，我从没想过这辈子会和伟大两个字连在一起。你晓不晓得每次看你飞行，我和你爸爸的心都是悬着的，什么激动啊、骄傲啊，都是你平安落地后才会有的感觉。那年劝说你停飞，你哭着说不愿意，我就晓得，我这个女儿，是要嫁给飞机了。——余旭的母亲 胡中秋042007年11月30日，一名军人感动了一座城，感动了全中国。年仅28岁的火箭军某部中尉参谋孟祥斌，为救轻生女青年，从浙江金华城南桥头纵身跃入婺江壮烈牺牲，被中央军委追授“舍己救人模范军官”荣誉称号，当选感动中国2007年度人物。十年前我带孩子去部队探亲，你用一次平凡的跳跃为自己的人生与军旅生涯画上一个圆！那刻我在悲伤中寻找活下去的理由，我在绝望中追问生命的意义，我在反思中思考，你平凡的一跃后，作为妻子的我如何将平凡继续？也就在那天整理你遗物时看到本子上这样一句话———认认真真做人，踏踏实实做事！——孟祥斌的妻子 叶庆华资料来源/老海军 军报记者等图/视觉中国 军报记者所谓英雄，不关乎时代，不关乎环境。你若感恩，他便值得，△中央广播电视总台7位播音员、主持人诵读秋瑾、夏明翰、赵一曼、黄继光、雷锋、孔繁森、焦裕禄等先烈的家书或诗歌，追思英雄，铭记他们的牺牲和奉献。点击「写留言」表达追思`

	a := &Article{
		Title: tt,
		Cont:  cc,
	}

	t.Fatal(a.AiGetCategories())
}

func Test_Cp(t *testing.T) {
	song := make(map[string]interface{})

	tt := "go语言标准json库Marshal不支持gbk格式string?"
	cc := `整个go的语言基础就是utf8，包括源码，运行时都是utf8，除此之外，还有很多编码，你觉得gbk常用是因为你是中国人，日本和韩国的还不一定呢。json那么实现，就简化整个json的解码实现，至于其它编码，再做个转码的就好了。`

	// gbkTitle, _ := ioutil.ReadAll(transform.NewReader(bytes.NewReader([]byte(tt)), simplifiedchinese.GBK.NewEncoder()))

	// gbkContent, _ := ioutil.ReadAll(transform.NewReader(bytes.NewReader([]byte(cc)), simplifiedchinese.GBK.NewEncoder()))

	song["title"] = tt
	song["content"] = cc
	// src := "编码转换内容内容"
	// enc := mahonia.NewEncoder("GBK")
	// // output := enc.ConvertString(tt)

	// song["title"] = enc.ConvertString(tt)
	// song["content"] = enc.ConvertString(cc)

	bytesData, err := json.Marshal(song)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	gbkBytesData, err := UTF8ToGBK(bytesData)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	reader := bytes.NewReader(gbkBytesData)
	url := `https://aip.baidubce.com/rpc/2.0/nlp/v1/keyword?access_token=24.01a7fba39af897d7e5c3141b28962bd4.2592000.1526178157.282335-11067381`
	request, err := http.NewRequest("POST", url, reader)
	if err != nil {
		t.Fatal(err.Error())
		return
	}
	// request.Header.Set("Content-Type", "application/json")
	client := http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		t.Fatal(err.Error())
		return
	}

	// reader2, err := charset.NewReader(resp.Body, strings.ToLower(resp.Header.Get("Content-Type")))

	// respBytes, err := ioutil.ReadAll(reader2)

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err.Error())
		return
	}
	//byte数组直接转成string，优化内存
	str := string(respBytes)

	// t.Fatal(str)
	udata := ConvertStrEncode(str, "gbk", "utf-8")
	t.Fatal(udata)
}

func UrlEncode(instr string) string {
	return url.QueryEscape(instr)
}
func Test_Cp2(t *testing.T) {

	tt := "iphone手机出现“白苹果”原因及解决办法，用苹果手机的可以看下"
	cc := `如果下面的方法还是没有解决你的问题建议来我们门店看下成都市锦江区红星路三段99号银石广场24层01室。在通电的情况下掉进清水，这种情况一不需要拆机处理。尽快断电。用力甩干，但别把机器甩掉，主意要把屏幕内的水甩出来。如果屏幕残留有水滴，干后会有痕迹。^H3 放在台灯，射灯等轻微热源下让水分慢慢散去。`

	gbkTitle := UrlEncode(ConvertStrEncode(tt, "utf-8", "gbk"))

	gbkContent := UrlEncode(ConvertStrEncode(cc, "utf-8", "gbk"))

	post_arg := map[string]interface{}{
		"title":   gbkTitle,
		"content": gbkContent,
	}
	post_json, err := json.Marshal(post_arg)
	if nil != err {
		panic(err)
	}
	real_uri := `https://aip.baidubce.com/rpc/2.0/nlp/v1/keyword?access_token=24.01a7fba39af897d7e5c3141b28962bd4.2592000.1526178157.282335-11067381`

	req, err := http.NewRequest("POST", real_uri, bytes.NewReader(post_json))
	client := http.Client{}

	resp, _ := client.Do(req)

	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if nil != err {
		panic(err)
	}

	udata := ConvertStrEncode(string(data), "gbk", "utf-8")

	if "null" == udata {
		panic(errors.New("貌似返回空"))
	}

	map_result := make(map[string]interface{})
	json.Unmarshal([]byte(udata), &map_result)

	panic(map_result)
	error_msg, ok := map_result["error_msg"]
	if ok {
		panic(error_msg)
	}

	// map_result["tags"].([]interface{})

	fmt.Println(map_result["tags"].([]interface{}))
}
