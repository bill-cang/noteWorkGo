package main

var (
	//<img src="/Uploads/Special/amami_tubasa.jpg" title="天海つばさ / Tsubasa Amami" class="img-thumbnail">
	reImg   = `"(https?://[^"]+?(\.((jpg)|(jpeg)|(png)|(gif)|(bmp)|(svg)|(swf)|(ico))))"`
	rexpImg = `<img (src="[\s\S]+?(\.jpg)") (title="[\s\S]+?") (class="[\s\S]+?")>`

	/*
	<tr>
						<td><a href="/vod-read-id-308292.html">NATR-592</a></td>
						<td class="td2"><a href="/vod-read-id-308292.html" target="_blank">まるまる！佐々木あき</a></td>
						<td>2018-10-26</td>
						<td class="td4">なでしこ$$$Nadeshiko</td>
						<td>242分钟</td>
					</tr>
	*/
	//番号信息正则
	repxFh = `<tr>[\s]+?<td><a href="/vod-read-id-\d{6}.html">([\s\S]+?)</a></td>[\s]+?<td class="td2"><a href="/vod-read-id-\d{6}.html" target="_blank">([\s\S]+?)</a></td>[\s]+?<td>([\s\S]+?)</td>[\s]+?<td class="td4">[\s\S]+?</td>[\s]+?<td>(\d{1,4})分钟</td>[\s]+?</tr>`
	//老师名字正则  <h2 class="page-header">佐々木あき 参演作品、番号列表</h2>
	repxXm =`<h2 class="page-header">(.+?) 参演作品、番号列表</h2>`

	//艺人信息与总页数
	//老师名字
	repxls_name=`女优列表</a> > (.+?)</h2>`
	//老师信息
	repxls_mass=`<p>英文名字：([\s\S]+?)</p>[\s\S]*</p>`
	//老师页数
	repxls_copa=`当前:1/(\d{1,})页`

	//所有艺人
	repxYr=`href="/special-read-id-(\d{1,5}).html" target="_blank"><img`

)
