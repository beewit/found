package main

import "regexp"

func main() {
	reStr := `<html lang="zh-cn" class="skin-dark" style="padding-right: 0px; margin-top: 0px;"><head>
<noscript>&lt;meta http-equiv="refresh" content="0; url=//qzs.qq.com/qzone/v6/troubleshooter/noscript.html" /&gt;</noscript>
<meta charset="UTF-8">
<meta http-equiv="content-type" content="text/html; charset=UTF-8">
<title>→ωó暧р女孩♂ [http://294477044.qzone.qq.com]</title>
<meta name="keywords" content="QQ空间,黄钻,免费装扮,开心农场,QQ农场,QQ牧场">



	<meta name="description" content="对一个女孩的思念">

<link rel="apple-touch-icon" href="//qzonestyle.gtimg.cn/qzone/v8/index/touch-icon-ipad-retina.png">
<link rel="apple-touch-icon" sizes="76x76" href="//qzonestyle.gtimg.cn/qzone/v8/index/touch-icon-ipad.png">
<link rel="apple-touch-icon" sizes="120x120" href="//qzonestyle.gtimg.cn/qzone/v8/index/touch-icon-iphone-retina.png">
<link rel="apple-touch-icon" sizes="152x152" href="//qzonestyle.gtimg.cn/qzone/v8/index/touch-icon-ipad-retina.png">
<link rel="icon" sizes="any" mask="" href="//qzonestyle.gtimg.cn/qzone/v8/img/Qzone.svg">
<meta name="theme-color" content="#FFC028">

<script type="text/javascript">
	window.diyitems_1_url = ''

		window.g_cdn_proto = 'https:'



    	var g_domain = "qq.com";

    document.domain=g_domain;

    window.g_point0=Date.now();
    var _s_=new Date(),
	g_T={},
	siDomain="qzonestyle.gtimg.cn",
	imgcacheDomain="qzs."+g_domain,

	g_iUin=294477044,
	g_iLoginUin=294477044;



g_T.fwp=[_s_];


//引入配置文件

window.g_dms = {
            'u.qzone.qq.com' : 1,
            'base.s2.qzone.qq.com' : 1,
            'base.s8.qzone.qq.com' : 1,
            'base.qzone.qq.com' : 1,
            'r.qzone.qq.com' : 1,
            'rsh.qzone.qq.com' : 1,
            'ic2.cnc.qzone.qq.com' : 1,
            'ic2.edu.qzone.qq.com':1,
            'ic2.qzone.qq.com':1,
            'ic2.s1.qzone.qq.com':1,
            'ic2.s11.qzone.qq.com' : 1,
            'ic2.s12.qzone.qq.com':1,
            'ic2.s2.qzone.qq.com' : 1,
            'ic2.s21.qzone.qq.com' : 1,
            'ic2.s5.qzone.qq.com':1,
            'ic2.s51.qzone.qq.com':1,
            'ic2.s6.qzone.qq.com':1,
            'ic2.s7.qzone.qq.com':1,
            'ic2.s8.qzone.qq.com' : 1,
            'ic2.qzone.qq.com' : 1,
            'xalist.photo.qzone.qq.com' :1,
            'hzalist.photo.qzone.qq.com' :1,
            'alist.photo.qzone.qq.com' :1,
            'shalist.photo.qzone.qq.com' :1,
            'plist.photo.qzone.qq.com':1,
            'xaplist.photo.qzone.qq.com':1,
            'hzplist.photo.qzone.qq.com':1,
            'gzplist.photo.qzone.qq.com':1,
            'shplist.photo.qzone.qq.com':1,
            'photo.qzone.qq.com':1,
            'xa.photo.qzone.qq.com':1,
            'hz.photo.qzone.qq.com':1,
            'gz.photo.qzone.qq.com':1,
            'shanghai.photo.qzone.qq.com':1,
            'app.photo.qzone.qq.com':1,
            'b.qzone.qq.com':1,
            'b1.cnc.qzone.qq.com':1,
            'b1.edu.qzone.qq.com':1,
            'b1.qzone.qq.com':1,
            'b11.cnc.qzone.qq.com':1,
            'b11.edu.qzone.qq.com':1,
            'b11.qzone.qq.com':1,
            'b1.ctc.qzone.qq.com':1,
            'b11.ctc.qzone.qq.com':1,
            'b.ctc.qzone.qq.com':1,
            'b.edu.qzone.qq.com':1,
            'fav.qzone.qq.com' : 1,
            'sns.qzone.qq.com' : 1,
            'm.qzone.qq.com' : 1,
            'snsapp.qzone.qq.com' : 1,
            'taotao.qzone.qq.com' : 1,
            'w.qzone.qq.com' : 1,
            'g.qzone.qq.com' : 1
        };


window.g_sdms = {
    'taotao.qq.com' : 1,
    'taotao.qzone.qq.com' : 1,
    'b.qzone.qq.com' : 1,
    'pageapp.qzone.qq.com' : 1,
    'b11.qzone.qq.com' : 1,
    'b1.qzone.qq.com' : 1,
    'br.qzone.qq.com' : 1,
    'm.qzone.qq.com' : 1,
    'base.qzone.qq.com' : 1,
    'w.qzone.qq.com' : 1,
    'g.qzone.qq.com' : 1,
    'r.qzone.qq.com' : 1,
    'ic2.qzone.qq.com' : 1,
    'boss.qzone.qq.com' : 1,
    'mall.qzone.qq.com' : 1,
    'statistic.qzone.qq.com' : 1,
    'fav.qzone.qq.com' : 1,
    'snsapp.qzone.qq.com' : 1,
    'vip.qzone.qq.com' : 1,
    'route.store.qq.com' : 1,
            'drift.qzone.qq.com' : 1,
            'p1.qzone.qq.com' : 1,
            'p2.qzone.qq.com' : 1,
            'xalist.photo.qzone.qq.com' :1,
            'hzalist.photo.qzone.qq.com' :1,
            'alist.photo.qzone.qq.com' :1,
            'shalist.photo.qzone.qq.com' :1,
            'plist.photo.qzone.qq.com':1,
            'xaplist.photo.qzone.qq.com':1,
            'hzplist.photo.qzone.qq.com':1,
            'gzplist.photo.qzone.qq.com':1,
            'shplist.photo.qzone.qq.com':1,
            'photo.qzone.qq.com':1,
            'xa.photo.qzone.qq.com':1,
            'hz.photo.qzone.qq.com':1,
            'gz.photo.qzone.qq.com':1,
            'shanghai.photo.qzone.qq.com':1,
            'tj.photo.qzone.qq.com':1,
            'app.photo.qzone.qq.com':1,
            'memo.qq.com' : 1,
            'analy.qq.com' : 1,
            'analy.qzone.qq.com' : 1,
            'page.qq.com' : 1,
            'rsh.qzone.qq.com' : 1,
            'search.qzone.qq.com' : 1,
            'flower.qzone.qq.com' : 1,
            'sz.ic2.qzone.qq.com' : 1,//通过qqtips到PC空间时，ic2域名统一走深圳
            'up.photo.qq.com' : 1
};

window.g_cgidomain = location.host.replace('.qzone.qq.com','') || 'h5';



    window.g_proto = 'https:';


!function(e){var n,i,t,o={maxBatchReportCount:6,interval:3e3,reportInterface:"//h5.qzone.qq.com/wspeed.qq.com/w.cgi?"},r={},a=navigator.userAgent;navigator.appVersion;if(r.adjustBehaviors=function(){},window.ActiveXObject||window.msIsStaticHTML){if(r.ie=6,(window.XMLHttpRequest||a.indexOf("MSIE 7.0")>-1)&&(r.ie=7),(window.XDomainRequest||a.indexOf("Trident/4.0")>-1)&&(r.ie=8),a.indexOf("Trident/5.0")>-1&&(r.ie=9),a.indexOf("Trident/6.0")>-1&&(r.ie=10),a.indexOf("Trident/7.0")>-1&&(r.ie=11),r.isBeta=navigator.appMinorVersion&&navigator.appMinorVersion.toLowerCase().indexOf("beta")>-1,r.ie<7)try{document.execCommand("BackgroundImageCache",!1,!0)}catch(s){}t=function(e){return function(n,i){var t;return"string"==typeof n?e(n,i):(t=Array.prototype.slice.call(arguments,2),e(function(){n.apply(null,t)},i))}}}else document.getBoxObjectFor||"undefined"!=typeof window.mozInnerScreenX?(n=/(?:Firefox|GranParadiso|Iceweasel|Minefield).(\d+\.\d+)/i,r.firefox=parseFloat((n.exec(a)||n.exec("Firefox/3.3"))[1],10)):navigator.taintEnabled?window.opera?r.opera=parseFloat(window.opera.version(),10):r.ie=6:(i=/AppleWebKit.(\d+\.\d+)/i.exec(a),r.webkit=i?parseFloat(i[1],10):document.evaluate?document.querySelector?525:420:419,(i=/Chrome.(\d+\.\d+)/i.exec(a))||window.chrome?r.chrome=i?parseFloat(i[1],10):"2.0":((i=/Version.(\d+\.\d+)/i.exec(a))||window.safariHandler)&&(r.safari=i?parseFloat(i[1],10):"3.3"),r.air=a.indexOf("AdobeAIR")>-1?1:0,r.isiPod=a.indexOf("iPod")>-1,r.isiPad=a.indexOf("iPad")>-1,r.isiPhone=a.indexOf("iPhone")>-1);(r.macs=a.indexOf("Mac OS X")>-1)||(r.windows=(i=/Windows.+?(\d+\.\d+)/i.exec(a),i&&parseFloat(i[1],10)),r.linux=a.indexOf("Linux")>-1,r.android=a.indexOf("Android")>-1),r.iOS=a.indexOf("iPhone OS")>-1,!r.iOS&&(i=/OS (\d+(?:_\d+)*) like Mac OS X/i.exec(a),r.iOS=!(!i||!i[1]));var d=function(){var e=[],n=!1,i=new Image;return i.onload=i.onerror=function(){n=!1,e.length&&(n=!0,i.src=e.shift())},function(t){e.push(t),n||(n=!0,i.src=e.shift())}}(),c=function(e){return null===e?"null":void 0===e?"undefined":Object.prototype.toString.call(e).slice(8,-1).toLowerCase()},u=function(e){e.indexOf("://")<1&&(e=location.protocol+"//"+location.host+(0==e.indexOf("/")?"":location.pathname.substr(0,location.pathname.lastIndexOf("/")+1))+e);var n=e.split("://");if("array"==c(n)&&n.length>1&&/^[a-zA-Z]+$/.test(n[0])){this.protocol=n[0].toLowerCase();var i=n[1].split("/");if("array"==c(i)){this.host=i[0],this.pathname="/"+i.slice(1).join("/").replace(/(\?|\#).+/i,""),this.href=e;var t=n[1].lastIndexOf("?"),o=n[1].lastIndexOf("#");return this.search=t>=0?n[1].substring(t):"",this.hash=o>=0?n[1].substring(o):"",this.search.length>0&&this.hash.length>0&&(o<t?this.search="":this.search=n[1].substring(t,o)),this}return null}return null},f=[],l=function(){var e="unknown";return e="",e=r.windows?"window_"+r.windows:r.macs?"macos_"+r.macs:r.isiPad?"ipad":"unknown_os",e+=r.chrome?"_chrome_"+r.chrome:r.ie?"_ie_"+r.ie:r.firefox?"_ff_"+r.firefox:"_unknown_br"},h=function(e){e.frequency||(0===e.resultcode?e.frequency=20:e.frequency=1);var n={appid:e.appid||"1000372",releaseversion:l(),stime:e.stime||11e3,commandid:e.commandid||"",apn:"wifi",resultcode:e.resultcode||0,touin:window.g_iLoginUin||"0",frequency:e.frequency||1,r:Math.random()};n.frequency>1&&Math.random()*n.frequency>1||(n.commandid=m(n.commandid),f.push(n),x())},p=function(e){return/^(?:ht|f)tp(?:s)?\:\/\/(?:[\w\-\.]+)\.\w+/i.test(e)},m=function(e){if(!e)return"";if(!p(e))return e;var n=u(e),i=(n.host||"unknown").replace(".qzone.qq.com","/p").replace(".qq.com","/q").replace("qzonestyle.gtimg.cn","cdn").replace("qzs.qq.com","cdn"),t=n.pathname.split("/");return i+"/"+t[t.length-1]},w=function(){var e,n=[];if(!f.length)return"";var i=0;for(n.push("key=appid,releaseversion,stime,commandid,resultcode,touin,frequency");(e=f.shift())&&(n.push(i+1+"_1="+e.appid),n.push(i+1+"_2="+e.releaseversion),n.push(i+1+"_3="+e.stime),n.push(i+1+"_4="+e.commandid),n.push(i+1+"_5="+e.resultcode),n.push(i+1+"_6="+e.touin),n.push(i+1+"_7="+e.frequency),i++,!(i>=h.maxBatchReportCount)););return n.join("&")},x=function(){if(1!=h._isSending){h._isSending=!0;var e=o.reportInterface;setTimeout(function(){var n=w();n&&n.length&&d(e+n),h._isSending=!1,f.length&&x()},o.interval)}};h.getVersion=l,window.haboStat=h}(window);

(function() {
        var reportNum = 0;
        var beforeLoad = 1;
        var needReport = window.g_iLoginUin%100 == 43;

		var getReportResult = function(){
			var result;
            var pathName = location.pathname;

            if(!pathName){
            	pathName = '/'
			}

            var frags = pathName.split('/');
            var appid = window.g_app_identifier.toLowerCase();
            if(window.g_isDirectApp){
                if(appid == 'mood' || appid=='311' || appid=='photo' || appid=='4' || appid=='share' || appid=='202' || appid=='blog' || appid=='2'){//说说和相册的appid有点特别，要修正一下
                    if(frags[1]==appid){
                        frags[2] && (frags[2] ='num')
					}else if(frags[2]==appid){
                        frags[3] && (frags[3] = 'num')
                        frags[4] && (frags[4] = 'num')

                    }else{
						frags[4] && (frags[4] = 'num')
                        frags[5] && (frags[5] = 'num')
					}
                }
            }
            if(frags && frags.length>0){
                pathName = frags.join('/');
            }
			result = (location.host + pathName).replace(/\d{5,30}/,'num').replace(/\/+$/ig,'');
			return result;
		};

        var onErrorOcurred = function(event) {
            if(needReport && reportNum == 0) {
                reportNum = 1;

                var result = getReportResult();

                // 堆栈信息
                var errorObj = event.error || {};
                var stack = errorObj.stack || '';

                var xhr = new XMLHttpRequest();
                xhr.withCredentials = true;
				var url = '//h5.qzone.qq.com/log/post/error/' + result+'?from=pcqzone&qua='+window.haboStat.getVersion();
                xhr.open('post', url, true);
                xhr.setRequestHeader('Content-Type', 'application/x-www-form-urlencoded');
                var errMsg = event.message + ";url:" + location.href + "\n stack: " + stack + ";";
                errMsg += "scriptURI:" + event.filename + ";";
                errMsg += "lineNumber:" + event.lineno + ";";
                errMsg += "columnNumber:" + event.colno + ";";
                errMsg += "beforeLoad:" + beforeLoad + ";";
                xhr.send(errMsg);
            }
        };

        var onWindowLoaded = function(evt){

            beforeLoad = 0;
            // reportNum为0（即成功）才上报，失败上报会在上报错误时在Node端过滤后上报
            // 先灰度下
             needReport && !reportNum && window.haboStat && window.haboStat({
				appid	: 1000416,
				commandid : getReportResult(),
                stime : 1000,
                resultcode : 0, //这里只上报成功
                touin : window.g_iLoginUin || 0,
                frequency :  1
            });

        }

        if(window.addEventListener){
             window.addEventListener('error', function(e){
                onErrorOcurred(e);
                });
             window.listenError = true;
            window.addEventListener('load', function(e){
                onWindowLoaded(e);
            });
        }else if(window.attachEvent){
            window.attachEvent('onerror', function(e){
                onErrorOcurred(e);
            })
            window.listenError = true;
            window.attachEvent('onload', function(e){
                onWindowLoaded(e);
            });
        }


}());




document.namespaces&&document.namespaces.add&&(document.namespaces.add('qz', location.protocol + '//qzone.qq.com/'),document.namespaces.add('x', location.protocol+'//qzone.qq.com/'));

	var QZFL={},
		QZFF_M_img_ribr=[];
	QZFL.event={
		getEvent : function(evt){
			var evt=window.event||evt,c,cnt;
			if(!evt&&window.Event){
				c=arguments.callee;
				cnt=0;
				while(c){
					if((evt=c.arguments[0])&&typeof(evt.srcElement)!="undefined"){
						break;
					}else if(cnt>9){
						break;
					}
					c=c.caller;
					++cnt;
				}
			}
			return evt;
		},
		getTarget : function(evt){
			var e=QZFL.event.getEvent(evt);
			if(e){
				return e.srcElement||e.target;
			}
			return null;
		}
	};
	QZFL.object = {
		getType: function(o){return o === null ? 'null' : (o === undefined ? 'undefined' : Object.prototype.toString.call(o).slice(8,-1).toLowerCase());}
	};
	QZFL.media={
		reduceImgByRule:function(ew,eh,opts,cb){
			QZFF_M_img_ribr.push(QZFL.event.getTarget());
		},
		adjustImageSize:function(w,h,s,cb,ecb){
			var op = {trueSrc:s,callback:function(o, type, ew, eh, p){(QZFL.object.getType(cb) == "function") && cb(o, ew, eh, null, p.ow, p.oh,p);},errCallback:ecb};
			QZFL.media.reduceImage(0, w, h, op);
		},
		reduceImage:function(type, ew, eh, opts){
			var rd = function(o, t, ew, eh, p, cb){
				var rl, k;
				if(p.rate==1){
					p.direction[0] = ( ew>eh ? 'height' : 'width');
					p.direction[1] = ( ew>eh ? 'width' : 'height');
				}
				rl = ( p.direction[t] == "width" ? ew : eh );
				t ? ( ( ( rl>p.shortSize ) ? ( rl = p.shortSize ) : 1 ) && ( p.k = p.shortSize/rl ) ) : ( ( ( rl>p.longSize ) ? ( rl = p.longSize ) : 1 ) && ( p.k = p.longSize/rl ) );
				o.setAttribute(p.direction[t],rl);
				(QZFL.object.getType(cb) == "function") && cb(o, t, ew, eh, p);
			};
			opts = opts || {};
			opts.img = (opts.img && (typeof(opts.img.nodeName) != 'undefined' || typeof(opts.img.nodeType) != 'undefined')  ? opts.img : QZFL.event.getTarget());
			opts.img.onload=null;
			opts.trueSrc && (opts.img.src = opts.trueSrc);
			if(opts.img){
				if( ! ( opts.direction && opts.rate && opts.longSize && opts.shortSize ) ){
					r = QZFL.media.getImageInfo(function(o,p){
						if(!o||!p){
							return;
						}
						rd(opts.img, type, ew, eh, p, opts.callback)
					},opts);
				}else{
					rd(opts.img, type, ew, eh, opts, opts.callback)
				}
			}
		},
		getImageInfo:function(callback, opts){
			var gif = function(img,cb,opts){
				if(img){
					var _w = opts.ow || img.width, _h = opts.oh || img.height, r,ls,ss,d;
					if(_w && _h){
						if(_w >=_h){
							ls = _w;
							ss = _h;
							d = ["width","height"];
						}else{
							ls = _h;
							ss = _w;
							d = ["height","width"];
						}
						r = {
							direction:d,
							rate : ls/ ss,
							longSize :ls,
							shortSize :ss
						};
						r.ow = _w;
						r.oh = _h;
					}
					(QZFL.object.getType(cb) == "function") && cb(img,r,opts);
				}
			};
			opts = opts || {};
			if(QZFL.object.getType(opts.trueSrc)== "string"){
				var _i = new Image();
				_i.onload = (function(ele, cb, p){
					return function(){
						gif( ele, cb, p );
						ele = ele.onerror =  ele.onload = null;
					};
				})(_i, callback, opts);

				_i.onerror =  (function(ele, cb, p){
					return function(){
						if (typeof(p.errCallback) == 'function') {
							p.errCallback();
						}
						ele = ele.onerror =  ele.onload = null;
					};
				})(_i, callback, opts);

				_i.src = opts.trueSrc;
			}else if(opts.img && opts.img.nodeType == 1){
				gif(opts.img, callback, opts );
			}
		}
	};


	function restXHTML(s){
		return s.replace(/&amp;|&lt;|&gt;|&apos;|&#0?39;|&quot;/g,function(a){
			switch (a){
				case "&amp;":return "&";
				case "&lt;":return"<";
				case "&gt;":return ">";
				case "&apos;":
				case "&#39;":
				case "&#039;":return "\x27";
				case "&quot;":return "\x22";
			}
		});
	}







window.g_isGrayRelease = true;




</script>


	<link href="//qzonestyle.gtimg.cn/aoi/skin/31.css?max_age=19830212&amp;d=20160530165317" rel="stylesheet">
<link href="//qzonestyle.gtimg.cn/aoi/icenter-170517152946.css?max_age=31536000" rel="stylesheet">
<link href="//qzonestyle.gtimg.cn/aoi/icenter-poster-comment-170109104014.css?max_age=31536000" rel="stylesheet">








		<style id="mainJSBg" type="text/css">

			.background-container{
				background-repeat:no-repeat;

				background-position:center top;

				background-attachment:scroll;

				background-image:url(//i.gtimg.cn/qzone/space_item/orig/13/123053_top.jpg);
			}



				.bg-body{
					background-image:url(//i.gtimg.cn/qzone/space_item/orig/13/123053_bg.jpg); background-repeat:repeat;}




				.layout-nav{background-image: url(//i.gtimg.cn/qzone/space_item/orig/13/123053_menu_bg.png);_background-image: url(//i.gtimg.cn/qzone/space_item/orig/13/123053_menu_bg_ie6.png);border: none;}


				.layout-background{background-image:url(//i.gtimg.cn/qzone/space_item/orig/13/123053_bg1.png); background-repeat:repeat;}




		</style>




<style id="mainJSTitleBar" type="text/css">
	.layout-head-inner {
		height:250px;



		}
</style>





<style type="text/css">

</style>
<style type="text/css" id="dynamicStyle">
.ownermode{display:;}
.clientmode{display:none;}
.editmode{display:none;}
.customoff{display:;}
.alphamode{display:none;}
</style>






<link rel="alternate" type="application/rss+xml" title="RSS news feed" href="http://feeds.qzone.qq.com/cgi-bin/cgi_rss_out?uin=294477044">



	<link rel="Shortcut Icon" href="//qzonestyle.gtimg.cn/aoi/img/logo/favicon.ico?max_age=31536000" type="image/x-icon">



	<!--[if IE]><base href="//qzs.qq.com/"></base><![endif]-->
	<!--[if !(IE)]><!--><link charset="utf-8" rel="stylesheet" href="//qzonestyle.gtimg.cn/aoi/ui-dialog.css"><link charset="utf-8" rel="stylesheet" href="https://qzonestyle.gtimg.cn/aoi/icenter-act.css"><link charset="utf-8" rel="stylesheet" href="https://qzonestyle.gtimg.cn/open_proj/vip/c_vip_hintbar_v2.css?max_age=19888888&amp;v=20150522"><link charset="utf-8" rel="stylesheet" href="https://qzonestyle.gtimg.cn/qzone_v6/qz_vip_tip_v2.css?max_age=31536000"><link charset="utf-8" rel="stylesheet" href="https://qzonestyle.gtimg.cn/aoi/dialog-vip-tip.css?max_age=31536000"><link charset="utf-8" rel="stylesheet" href="https://qzonestyle.gtimg.cn/aoi/hot-tips2.css"><link charset="utf-8" rel="stylesheet" href="https://qzonestyle.gtimg.cn/aoi/photo-scan-layer.css?d=140918103710"><link charset="utf-8" rel="stylesheet" href="https://qzonestyle.gtimg.cn/aoi/ui-dialog.css"><link charset="utf-8" rel="stylesheet" href="https://qzonestyle.gtimg.cn/aoi/dialog-qzone-app-download.css"><base href="//qzs.qq.com/"><!--<![endif]-->




<!--[if IE]>
<script type="text/javascript">
    function toAbsURL(s) {
     var l = location, h, p, f, i;
     if (/^\w+:/.test(s)) {
       return s;
     }
     h = l.protocol + '//' + l.host + (l.port!=''?(':' + l.port):'');
     if (s.indexOf('/') == 0) {
       return h + s;
     }
     p = l.pathname.replace(/\/[^\/]*$/, '');
     f = s.match(/\.\.\//g);
     if (f) {
       s = s.substring(f.length * 3);
       for (i = f.length; i--;) {
         p = p.substring(0, p.lastIndexOf('/'));
       }
     }
     return h + p + '/' + s;
   }


    var base = document.getElementsByTagName('base')[0];
    base.href = toAbsURL(base.href);
</script>
<![endif]-->
<link rev="stylesheet" rel="stylesheet" type="text/css" media="screen" href="https://qzonestyle.gtimg.cn/aoi/qz-tips-box.css"><script id="_gdt_loader_0.8997824479281149" data-name="1648225091" type="text/javascript" charset="utf-8" src="https://qzonestyle.gtimg.cn/qzone/biz/ac/comm/qbscomm.20150907.js"></script><script id="_gdt_loader_0.9368626653947674" data-name="607126241" type="text/javascript" charset="utf-8" src="https://qzonestyle.gtimg.cn/qzone/biz/ac/comm/gdtlib.20160810.js"></script><script id="_gdt_loader_0.29862186343507147" data-name="1513054522" type="text/javascript" charset="utf-8" src="https://qzonestyle.gtimg.cn/qzone/biz/ac/comm/ver.20170622.js"></script><link rev="stylesheet" rel="stylesheet" type="text/css" media="screen" href="https://qzonestyle.gtimg.cn/aoi/qz_smart_search.css"><link rev="stylesheet" rel="stylesheet" type="text/css" media="screen" href="https://qzonestyle.gtimg.cn/aoi/icenter-delay-170302163216.css?max_age=31536000"><script id="_gdt_loader_0.9170674689566871" data-name="2003264745" type="text/javascript" charset="UTF-8" src="//qzonestyle.gtimg.cn/qzone/biz/gdt/display/click.js"></script><style type="text/css" id="gdt-corner-css">*{margin:0;padding:0}.j-corner{display:inline-block;width:44px;overflow:hidden;transition:all 0.8s;cursor:pointer;position:absolute;right:0;bottom:0;z-index:3;}.j-corner:hover{width:92px;}.j-corner .tr{width:100%;height:18px;display:inline-block;background-image:url(//qzonestyle.gtimg.cn/gdt_ui_proj/dist/adpoly/logo-adsign/images/logo-bg-defalt.png);transform-origin:0 18px;transition:all 0.6s ease;position:absolute;top:0;left:0;}.j-corner:hover .tr{width:92px;height:18px;display:inline-block;background-image:url(//qzonestyle.gtimg.cn/gdt_ui_proj/dist/adpoly/logo-adsign/images/logo-bg-open.png);}.j-corner .info{position:relative;z-index:2;font-size:12px;height:18px;padding-left:7px;line-height:18px;white-space:nowrap;color:#fff;}.j-corner .text1, .text2{display:inline-block;position:absolute;top:0;left:7px;height:18px;transition:all 0.6s ease;}.j-corner .text1{width:44px;background:url(//qzonestyle.gtimg.cn/gdt_ui_proj/dist/adpoly/logo-adsign/images/logo-defalt.png) no-repeat 0 center;opacity:1;filter:alpha(opacity=100);}.j-corner:hover .text1{opacity:0;filter:alpha(opacity=0);}.j-corner .text2{width:92px;background:url(//qzonestyle.gtimg.cn/gdt_ui_proj/dist/adpoly/logo-adsign/images/logo-open.png) no-repeat 0 center;opacity:0;filter:alpha(opacity=0);}.j-corner:hover .text2{opacity:1;filter:alpha(opacity=100);}</style><link rev="stylesheet" rel="stylesheet" type="text/css" media="screen" href="https://qzonestyle.gtimg.cn/aoi/fifa-flags.css?max_age=31536000"><script type="text/javascript" charset="UTF-8" src="https://qzonestyle.gtimg.cn/qzone/biz/gdt/display/modules/mod/stat_ff8856c.js?max_age=31536000"></script><link rev="stylesheet" rel="stylesheet" type="text/css" media="screen" href="https://qzonestyle.gtimg.cn/open_proj/gg-qz-icenter.css"><script type="text/javascript" charset="UTF-8" src="https://qzonestyle.gtimg.cn/qzone/biz/gdt/display/modules/mod/asynload_a051ecb.js?max_age=31536000"></script><script type="text/javascript" charset="UTF-8" src="https://qzonestyle.gtimg.cn/qzone/biz/gdt/display/modules/mod/sns_f946688.js?max_age=31536000"></script><script type="text/javascript" charset="UTF-8" src="https://qzonestyle.gtimg.cn/qzone/biz/gdt/display/modules/mod/extrender_8b8eaf8.js?max_age=31536000"></script><link rev="stylesheet" rel="stylesheet" type="text/css" media="screen" href="https://qzonestyle.gtimg.cn/qzone_v6/gb_hintbar.css"><link id="qznccss" rev="stylesheet" rel="stylesheet" type="text/css" media="screen" href="https://qzonestyle.gtimg.cn/aoi/gb-card.css"><style id="viewerStyle">
.js_fade_like{
 line-height:0px;
 display:none;
}

.figure-comment{
 display : none;
}
.figure-comment.js-can-comment{
 display : block;
}

.js_show_comment #mod_comment_module{}
.js_show_comment #mod_retweet{
 display : none;
}
.js_show_comment #mod_retweet .textinput{
 visibility:hidden;
}

.js_show_retweet #mod_comment_module{
 display : none;
}
.js_show_retweet #mod_retweet{}
.mod_comment_report{
 visibility:hidden;
}

.js_standalone_mode .photo_layer_close{
 visibility:hidden;
}
.js_standalone_mode .photo_layer {
 padding : 0;
}
.js_standalone_mode .js_reproduced{
 visibility :hidden;
}
.js-scrolling-inner{
 background-color:#c9c9c9;
}

.js-scrollbox{
 overflow:hidden;
 height:100%;
 position:relative;
}
.js-scrollcont{
 height:100%;
 overflow-x:hidden;
 overflow-y:scroll;
 padding-right:60px;
 position:relative;
 width:100%;
}
.js-scrollbar{
 height:100%;
 position:absolute;
 right:0;
 top:0;
 width:15px;
 z-index:9;
}
.js-scrolling-inner{
 border-radius:4px 4px 4px 4px;
 display:block;
 height:100%;
 width:6px;
 float:right;
}
.js-scrolling-inner-hover{
 width:10px;
 border-radius:6px 6px 6px 6px;
}
.js-scrollbar-hover{
 width:15px;
}

a:hover .js-scrolling-inner,a:active .js-scrolling-inner{
 background-color:#a8a8a8;
}

.rotate{
 transition:transform 0.2s linear;
 -webkit-transition:-webkit-transform 0.2s linear;
 -moz-transition:-moz-transform 0.2s linear;
 -o-transition:-o-transform 0.2s linear;
 -ms-transition:-ms-transform 0.2s linear;
}

.js-fullscreen-wrapper{
 position:absolute;
 top:0;
 left:0;
 right:0;
 bottom:0;
 z-index:9999;
 background-color:#000;
 -webkit-backface-visibility:hidden;
}

.js-fullscreen-layer-transition{
 transition:opacity 1.2s ease-out;
 -webkit-transition:opacity 1.2s ease-out;
 -moz-transition:opacity 1.2s ease-out;
 -o-transition:opacity 1.2s ease-out;
 -ms-transition:opacity 1.2s ease-out;
 opacity: 1;
}
.js-fullscreen-layer-transition-none{
 transition:none;
 -webkit-transition:none;
 -moz-transition:none;
 -o-transition:none;
 -ms-transition:none;
 opacity: 0;
}

.js-fullscreen-cont{
 position:absolute;
 top:0;
 left:0;
 right:0;
 bottom:0;
}

.js-fullscreen-cont-transition{
 transition:1.2s ease-out;
 -webkit-transition:1.2s ease-out;
 -moz-transition:1.2s ease-out;
 -o-transition:1.2s ease-out;
 -ms-transition:1.2s ease-out;
}

.js-fullscreen-img{
 position:absolute;
 display:none;
 top:0;
 left:0;
 right:0;
 bottom:0;
 background-repeat:no-repeat;
 background-size:cover;
 background-position:top left;
 -webkit-transform-origin:0px 0px;
}

.js-rightmenu-box{
 width:110px;
 position:absolute;
 z-index:9999;
}
.js-rightmenu-box ul{
 background-color:#FFF
}
.js-rightmenu-box ul li{
 margin:0;
 height:auto;
 line-height:normal;
}
.js-rightmenu-box ul li a{
 display:block;
 width:100%;
 padding:8px 0;
 text-align:center;
 color:#333;
}
.js-rightmenu-box ul li a:hover{
 background-color:#5CAAE6;
 color:#FFF;
 text-decoration:none;
}

.js-hidden {
    visibility:hidden
}

.handle-tab .btn-txt{
 display:none;
}
.handle-tab .btn-txt-num{
 display:none;
}
.handle-tab.j-show-txt .btn-txt{
 display:inline-block;
}
.handle-tab.j-show-txt-num .btn-txt-num{
 display:inline-block;
}

.js-thumb-item {
 position: relative;
 display: inline-block;
 float: left;
}

</style><link id="friendSelector_qz_at" rev="stylesheet" rel="stylesheet" type="text/css" media="screen" href="https://qzonestyle.gtimg.cn/qzone_v6/qz_at.css"><style type="text/css">.mod_commnets_poster .qzappwrap{display:none}.mod_retweet .qzappwrap{display:none}</style><link rev="stylesheet" rel="stylesheet" type="text/css" media="screen" href="https://qzonestyle.gtimg.cn/aoi/feed_likelist.css"></head>



    <body class="os-mac bg-body mode-theme date-20170921" style="">

<div style="position:absolute;top:-200px;">
    <a href="http://rc.qzone.qq.com/im?visitor=blind" tabindex="1" class="blind" accesskey="o">用读屏软件的朋友从这里进入QQ空间读屏版</a>
    <a href="http://qzs.qq.com/qzone/v6/accessibility/help.html" tabindex="2" accesskey="0">QQ空间无障碍使用帮助，请点击这里</a>
    <a href="http://support.qq.com/write.shtml?fid=811" tabindex="3" target="_blank">使用QQ空间遇到问题，点击这里反馈</a>
    <a href="//user.qzone.qq.com/294477044#pageContent" tabindex="4" accesskey="8">跳到内容区</a>
</div>
<div class="layout-invisible">
    <span id="trs_tip" style="background-color:#FFF2D1;padding:10px;">如果您看到这个提示，说明QQ空间无法正常打开，请尝试<a href="//user.qzone.qq.com/troubleshooter#style" title="空间小助手"><strong>使用空间小助手修复</strong></a></span>
</div>

















<div class="top-fix-bar">
	<div class="top-fix-inner" style="">
        <div class="top-fix-container" id="QZ_Toolbar_Container" style="margin-right: 0px; display: block;">
            <div class="top-fix-wrap">

                <a class="logo" href="//qzone.qq.com/" title="QQ空间">QQ空间</a>

                <ul class="top-nav">
                    <li class="nav-list" id="tb_ic_li">
                        <div class="nav-list-inner">
                            <a id="aIcenter" class="nav-hover" href="//user.qzone.qq.com/294477044/infocenter?via=toolbar" accesskey="1"><i class="ui-icon icon-icenter"></i><span>个人中心</span></a>
                        </div>
                    </li>
                    <li class="nav-list" id="tb_index_li">
                        <div class="nav-list-inner">
                            <a id="aMainPage" href="//user.qzone.qq.com/294477044/main" class="homepage-link a-link" accesskey="z"><i class="ui-icon icon-homepage"></i><span>我的主页</span><i class="drop-down-arrow"></i></a>
                        </div>

                        <div class="nav-drop-down" id="tb_menu_panel" style="display: none;">
                            <div class="side-area">
                                <div class="homepage-link">
                                    <a class="link-icon" href="//user.qzone.qq.com/294477044/main"><i class="icon-homepage"></i></a>
                                    <a class="link-text" href="//user.qzone.qq.com/294477044/main">主页</a>
                                </div>
                            </div>
                            <div class="main-area">
                                <div class="main-application">
                                    <ul class="main-application-list clearfix" id="tb_menu_list">
                                        <li class="menu_item_2">
                                            <a href="javascript:;"><i class="icon-blog"></i></a>
                                            <a class="app-name">日志</a>
                                        </li>
                                        <li class="menu_item_4">
                                            <a href="javascript:;"><i class="icon-album"></i></a>
                                            <a class="app-name">相册</a>
                                        </li>
                                        <li class="menu_item_311">
                                            <a href="javascript:;"><i class="icon-say"></i></a>
                                            <a class="app-name">说说</a>
                                        </li>
                                        <li class="menu_item_202">
                                            <a href="javascript:;"><i class="icon-share"></i></a>
                                            <a class="app-name">分享</a>
                                        </li>
                                        <li class="menu_item_847">
                                            <a href="javascript:;"><i class="icon-video"></i></a>
                                            <a class="app-name">视频</a>
                                        </li>
                                        <li class="menu_item_305">
                                            <a href="javascript:;"><i class="icon-music"></i></a>
                                            <a class="app-name">音乐</a>
                                        </li>
                                        <li class="menu_item_334">
                                            <a href="javascript:;"><i class="icon-message-book"></i></a>
                                            <a class="app-name">留言板</a>
                                        </li>
                                        <li class="menu_item_1">
                                            <a href="javascript:;"><i class="icon-personal"></i></a>
                                            <a class="app-name">个人档</a>
                                        </li>
                                         <li class="menu_item_907">
                                             <a href="javascript:;"><i class="icon-collect"></i></a>
                                             <a class="app-name">收藏</a>
                                         </li>
                                    </ul>
                                </div>
                            </div>
                        </div>
                    </li>

					<li class="nav-list" id="tb_friend_li">
						<div class="nav-list-inner">
							<a id="aMyFriends" href="//user.qzone.qq.com/294477044/myhome/friends/index" class="friends-link a-link"><i class="ui-icon icon-friend"></i><span>好友</span><i class="drop-down-arrow"></i></a>
						</div>

						<div class="nav-drop-down friends-drop-down" id="friends-drop-down" style="display: none;">
							<div class="side-area">
								<div class="friends-link">
									<a href="javascript:" class="link-icon"><i class="icon-friends"></i></a>
									<a href="javascript:" class="link-text">寻找好友</a>
								</div>

								<div class="friends-relation">
									<a href="javascript:" class="link-icon"><i class="icon-relation"></i></a>
									<a href="javascript:" class="link-text">亲密度</a>
								</div>


							</div>
							<div class="main-area">
								<div class="mod-friends-topbar">
									<div class="inner">
										<div class="hd">
											<h3>最近联络的好友</h3>
											<div class="friends-search-topbar ">
												<input type="text" placeholder="搜索好友" id="friend_search_input"><button><i class="ui-icon sp-top-search"></i></button>
											</div>

											<div class="friends-results-topbar" style="display:none">
												<ul class="results-list-topbar qz-scrollbar" id="search_friend_result"></ul>
												<p class="null" style="display:none">没有找到相关用户</p>
											</div>
										</div>
										<div class="bd">
											<ul class="friends-list-topbar clearfix" id="friends-list-topbar"><li><a href="http://user.qzone.qq.com/442769390" target="_blank" class="pic"><img src="http://qlogo3.store.qq.com/qzone/442769390/442769390/100?1354419484" alt="朱红兵的头像" title="你最近看过的好友"><span class="name ellipsis">朱红兵</span></a><a class="number" href="http://user.qzone.qq.com/442769390/friendship?fuin=442769390" target="_blank"><i class="ui-icon icon-relation"></i>11</a></li><li><a href="http://user.qzone.qq.com/87789607" target="_blank" class="pic"><img src="http://qlogo4.store.qq.com/qzone/87789607/87789607/100?1505584997" alt="黄健的头像" title="你最近看过的好友"><span class="name ellipsis">黄健</span></a><a class="number" href="http://user.qzone.qq.com/87789607/friendship?fuin=87789607" target="_blank"><i class="ui-icon icon-relation"></i>45</a></li><li><a href="http://user.qzone.qq.com/87001095" target="_blank" class="pic"><img src="http://qlogo4.store.qq.com/qzone/87001095/87001095/100?1357195352" alt="曹丑的头像" title="你最近看过的好友"><span class="name ellipsis">曹丑</span></a><a class="number" href="http://user.qzone.qq.com/87001095/friendship?fuin=87001095" target="_blank"><i class="ui-icon icon-relation"></i>51</a></li><li><a href="http://user.qzone.qq.com/364517951" target="_blank" class="pic"><img src="http://qlogo4.store.qq.com/qzone/364517951/364517951/100?1260246358" alt="静的头像" title="你最近看过的好友"><span class="name ellipsis">静</span></a><a class="number" href="http://user.qzone.qq.com/364517951/friendship?fuin=364517951" target="_blank"><i class="ui-icon icon-relation"></i>45</a></li><li><a href="http://user.qzone.qq.com/1814985685" target="_blank" class="pic"><img src="http://qlogo2.store.qq.com/qzone/1814985685/1814985685/100?1504886899" alt="德品美工的头像" title="你最近看过的好友"><span class="name ellipsis">德品美工</span></a><a class="number" href="http://user.qzone.qq.com/1814985685/friendship?fuin=1814985685" target="_blank"><i class="ui-icon icon-relation"></i>54</a></li><li><a href="http://user.qzone.qq.com/2287304030" target="_blank" class="pic"><img src="http://qlogo3.store.qq.com/qzone/2287304030/2287304030/100?1420985512" alt="陈琳的头像" title="你最近看过的好友"><span class="name ellipsis">陈琳</span></a><a class="number" href="http://user.qzone.qq.com/2287304030/friendship?fuin=2287304030" target="_blank"><i class="ui-icon icon-relation"></i>43</a></li><li><a href="http://user.qzone.qq.com/800013811" target="_blank" class="pic"><img src="http://qlogo4.store.qq.com/qzone/800013811/800013811/100" alt="腾讯开放平台服务团队的头像" title="你新添加的好友"><span class="name ellipsis">腾讯开放平台服务团队</span></a><a class="number" href="http://user.qzone.qq.com/800013811/friendship?fuin=800013811" target="_blank"><i class="ui-icon icon-relation"></i>11</a></li><li><a href="http://user.qzone.qq.com/5055632" target="_blank" class="pic"><img src="http://qlogo1.store.qq.com/qzone/5055632/5055632/100?1434524664" alt="前段的头像" title="你新添加的好友"><span class="name ellipsis">前段</span></a><a class="number" href="http://user.qzone.qq.com/5055632/friendship?fuin=5055632" target="_blank"><i class="ui-icon icon-relation"></i>47</a></li></ul>
										</div>
										<div class="ft"></div>
									</div>
								</div>
							</div>
						</div>
					</li>


                    <li class="nav-list" id="tb_app_li">
                        <div class="nav-list-inner">
                            <a id="aAppstore" href="//user.qzone.qq.com/294477044/appstore" class="application-link a-link"><i class="ui-icon icon-application"></i><span>游戏</span><i class="drop-down-arrow"></i></a>
                        </div>

                        <div class="nav-drop-down application-drop-down" id="tb_app_wall" style="display: none;">
                            <div class="side-area">
                                <div class="application-shop">
                                    <a class="link-icon" href="//user.qzone.qq.com/294477044/apppoints?via=QZ.APPWALL.APPCENTER"><i class="icon-shop"></i></a>
                                    <a class="link-text" href="//user.qzone.qq.com/294477044/apppoints?via=QZ.APPWALL.APPCENTER">积分商城</a>
                                </div>
                            </div>
                            <div class="main-area">
                                <div class="main-application">
                                    <ul id="tb_recent_app_list" class="main-application-list clearfix"></ul>
                                </div>
                                <div class="publicity-msg" style="display: none;">
                                    <a id="tb_qboss_text_ad" href="javascript:;" target="_blank"></a>
                                </div>
                            </div>
                        </div>
                    </li>

                    <li class="nav-list" id="tb_dress_li">
                        <div class="nav-list-inner">
                            <a id="aMall" class="dress-up-link a-link" href="//rc.qzone.qq.com/onekeydressup"><i class="ui-icon icon-dress-up"></i><span>装扮</span><i class="drop-down-arrow"></i></a>
                        </div>

                        <div class="nav-drop-down dress-drop-down" id="tb_dress_panel" style="display: none;"></div>
                    </li>

                </ul>


                <div class="user-info">
                    <a class="user-home" href="//user.qzone.qq.com/294477044/main">
                        <img class="user-avatar" src="//qlogo1.store.qq.com/qzone/294477044/294477044/30?1382948711" alt="您的头像">
                        <span class="user-name textoverflow">承諾，/aiq一世的誓言</span>
                    </a>
					<a id="tb_logout" class="logout-new" href="javascript:;" title="退出"><i class="ui-icon icon-logout"></i></a>
                    <div id="tb_setting_li" class="user-setting">
                        <i id="tb_setting_i" class="drop-down-arrow"></i>
                        <div class="user-drop-down" id="tb_setting_panel" style="display:none;">
                            <b class="white-line"></b>
                            <div class="drop-down-setting">
                                <a type="info" href="javascript:;">修改资料</a>
                                <a type="dress" href="javascript:;">主页排版</a>
                                <a type="set" href="javascript:;">空间设置</a>
                                <a type="auth" href="javascript:;">权限设置</a>
                            </div>

			<a target="_blank" class="online-service" href="http://service.qq.com/special_auto/qzone.html" id="aHelpCenter">帮助中心</a>

                            <a href="http://support.qq.com/discuss/46_1.shtml" target="_blank" title="意见反馈">意见反馈</a>
                            <a href="http://guanjia.qq.com/user/?ADTAG=media.outenter.qzone" target="_blank" title="安全防护" id="securityDefence" onclick="TCISD &amp;&amp; TCISD.hotClick('toolbar.houseKeeper', 'user.qzone.qq.com')">安全防护</a>
                        </div>
                    </div>
                    <a href="http://pay.qq.com/ipay/index.shtml?c=xxjzgw,xxjzghh&amp;ch=self&amp;aid=zone.tbarshuang" id="tb_vip_li" target="_blank" class="vip-setting"><i class="ui-icon icon-vip"></i></a>
				</div>


                <!-- 顶部搜索框 默认放出 2014-04-22 markqin -->
                <div class="top-search">
                    <div class="search-box">
                        <input class="search-input" type="text" placeholder="用户/游戏/动态"><!--搜索框默认态-->
                        <a href="javascript:;" class="search-button"><i class="ui-icon icon-search"></i></a>
                    </div>
                    <div class="search-drop-down" id="search_smart_panel_navigation_bar" style="display: none;">
                        <div id="search_smart_search_loading" class="waiting-box" style="display:none">
                            <img src="//qzonestyle.gtimg.cn/aoi/img/loading2.gif" alt="等待图标">
                        </div>
                        <div id="search_smart_search_list">
                            <div class="result-box r-modle j-default-result-list" style="display:none">
                                <a href="https://user.qzone.qq.com/294477044/infocenter?qz_referrer=special&amp;clickReport=searchSpecial" class="result-list">
                                    <div class="avatar"><i class="ui-icon icon-ts-special"></i></div>
                                    <div class="result-details">
                                        <p class="info">特别关心</p>
                                    </div>
                                </a>
                                <a href="https://user.qzone.qq.com/294477044/myhome/friends/ofpmd?clickReport=searchMaybeKnow" class="result-list">
                                    <div class="avatar"><i class="ui-icon icon-ts-qq"></i></div>
                                    <div class="result-details">
                                        <p class="info">可能认识的人</p>
                                    </div>
                                </a>
                                <a href="https://user.qzone.qq.com/294477044/myhome/217?clickReport=searchAuthenQzone" class="result-list">
                                    <div class="avatar"><i class="ui-icon icon-ts-approve"></i></div>
                                    <div class="result-details">
                                        <p class="info">认证空间</p>
                                    </div>
                                </a>
                                <a href="https://rc.qzone.qq.com/appstore?via=qzonesearchbar&amp;clickReport=searchApp" class="result-list">
                                    <div class="avatar"><i class="ui-icon icon-ts-app"></i></div>
                                    <div class="result-details">
                                        <p class="info">游戏应用</p>
                                    </div>
                                </a>
                            </div>
                            <div class="smart-result-list" style="display:none"></div>
                        </div>
                        <a id="search_smart_search_more" class="search-for-friends" style="display:none" href="javascript:;">
                            <span>查看更多</span><i class="more-icon"></i>
                        </a>
                    </div>
                </div>


                <div class="music-container" id="tb_music_li" title="点击设置背景音乐" style="cursor: pointer;">
                    <div class="music-panel">
                        <a class="music-play" href="javascript:;"><i class="ui-icon ico-music-play" title="点击设置背景音乐"></i></a>
                        <b class="music-dynamic"><i class="ui-icon ico-music-dynamic" title="点击设置背景音乐" style="cursor: pointer;"></i></b>
                    </div>

                    <div id="tb_music_panel" class="music-drop-down" style="display: none;"><i class="hline"></i>
                        <div class="drop-down-container">
                            <div class="background-music">
                                <i class="icon-qq-music"></i>
                                <b>我的背景音乐</b>
                                <div class="setting-box">
                                    <a href="javascript:;" class="music-setting">设置</a>
                                </div>
                            </div>
                            <div class="music-list">
                                <div class="music-scroll-container">
                                    <div style="top:0" class="music-choose">
                                    </div>
                                </div>
                                <b class="scroll-bar-container">
                                    <b style="top:0" class="scroll-bar"></b>
                                </b>
                            </div>
                        </div>
                    </div>
                </div>

                <div class="vip-drop-down" id="tb_vip_panel" style="display: none;"><b class="white-line" data-type="normal"></b><div class="drop-down-setting"><a id="aPayBySelf" href="http://pay.qq.com/ipay/index.shtml?c=xxjzgw,xxjzghh&amp;ch=self&amp;aid=zone.tbarshuang" target="_blank">开通黄钻</a><a id="aPayByYear" href="http://pay.qq.com/ipay/index.shtml?c=xxjzgw,xxjzghh&amp;ch=self&amp;aid=zone.tbarshuangnf&amp;paytime=year" target="_blank">开年费黄钻</a><a id="aPayByFriend" href="javascript:;">向好友索取</a><a id="aPayToFriend" target="_blank" href="http://pay.qq.com/ipay/index.shtml?c=xxjzgw,xxjzghh&amp;ch=send&amp;aid=zone.tbar.send&amp;ADTAG=pay.service.qzone.send&amp;u=0">赠送给好友</a><a id="aVipHome" target="_blank" href="http://vip.qzone.com/?login=qq">黄钻官网</a></div></div>
            </div>
        <div class="top-fix-extend">    							<a id="speed_signal_btn" class="icon-speed icon-speed-3" target="_blank" href="http://user.qzone.qq.com/troubleshooter"></a>							</div></div>
    </div>
</div>

<div class="background-container" id="layBackground">



    <div class="layout-head ">
        <div class="layout-head-inner" id="headContainer" style="background: no-repeat; height: 250px;">

            <div class="head-info"><div class="dragdrop_item_title" style="display: none;"><a class="edit_title_link right" href="javascript:;">编辑</a></div>
                <h1 class="head-title" id="top_head_title"><span class="title-text ui-mr5">→ωó暧р女孩♂</span>

<a href="javascript:;" id="qz-head-level" class="qz-level-flag" onclick="window.QZONE &amp;&amp; QZONE.FP.toApp('/qzscore');window.TCISD &amp;&amp; window.TCISD.hotClick('point_ic_image', 'flower.qzone.qq.com');return false;">
	<span class="qz_qzone_lv qz_qzone_lv_3 qz_qzone_no_2" title="当前空间等级：41级；积分：11569分"><span class="no"><b class="d4"></b><b class="d1"></b></span></span>
</a>



                        <span class="qz-app-flag wing-fly"><i class="ui-icon icon-qzphone"></i></span>

                </h1>
                <div class="head-description">
                    <span class="description-text" id="QZ_Space_Desc">对一个女孩的思念</span>
                </div>
            </div>
            <div class="head-detail">
                <div class="head-detail-name">
                    <span class="user-name textoverflow">承諾，/aiq一世的誓言</span>




                            <span class="user-vip-info">



                            </span>

                </div>












    <div class="head-detail-info clearfix"><!-- 未开通黄钻的用户成长信息不可见 -->
        <div class="detail-info-level">
            <a id="diamon" data-year="" data-super="0" class="f_user_vip qz_ichotclick " target="_blank" href="http://pay.qq.com/ipay/index.shtml?n=12&amp;c=xxjzgw,xxjzghh&amp;aid=info.nickname&amp;ch=qdqb,kj" style="cursor:pointer;" data-isleave="true">
            <i class="qz_vip_icon_l qz_vip_icon_l_gray_1"></i></a>
        </div>
        <div class="detail-info-con">
            <div class="qz-progress-bar  qz-progress-bar-gray" style="cursor: pointer;">
                <div class="progress-bar-info">
                    <span class="txt-value">成长值<b class="count">1</b></span><span class="txt-speed">成长速度<b class="count">-10点/天</b></span>
                </div>
                <div class="progress-bar-panel">


                    <div class="progress-bar-count" style="width:0%"><span class="count">0%</span></div>
                </div>
            </div>
        </div>



<div class="user-vip-info">



                <a class="qz-btn-vip qz-btn-vip-fee" title="续费黄钻" href="http://pay.qq.com/ipay/index.shtml?n=3&amp;c=xxjzgw,xxjzghh&amp;aid=info.nickname&amp;ch=qdqb,kj" target="_blank"></a>









            <a href="http://sweet.snsapp.qq.com/vip/index.html?ref=qzone" target="_blank" class="qz-lover-flag" title="情侣黄钻" onclick="window.TCISD&amp;&amp;window.TCISD.hotClick('isd.vip.lover.name','mall.qzone.qq.com');"><i class="ui-icon icon-lover-gray"></i></a>


</div>
    </div>

            </div>

            <!--挂件-->
            <div class="layout-shop-item" id="QZ_Shop_Items_Container" style="z-index: 101;">



















































			<div class="shop-item" id="menuContainer" style="width: 528px; height: 70px; left: 140px; top: 232px; " reged="1"><div class="">
			<!--[if IE]>
			<object id="customMenu" classid="clsid:d27cdb6e-ae6d-11cf-96b8-444553540000" width="528" height="70">
				<param name="movie" value="//qzonestyle.gtimg.cn/qzone/space_item/orig/14/123054.swf">
				<param name="allowScriptAccess" value="always">
				<param name="allownetworking" value="all">
				<param name="allowFullScreen" value="true">
				<param name="wmode" value="transparent">
				<param name="menu" value="false">
				<param name="scale" value="noScale">
				<param name="salign" value="1">
				<param name="flashvars" value="xml_path=//qzs.qq.com/qzone/v5/custom_menu/custom_menu.xml&amp;json_str=[{&quot;name&quot;:&quot;主页&quot;,&quot;href&quot;:&quot;N1&quot;},{&quot;name&quot;:&quot;日志&quot;,&quot;href&quot;:&quot;2&quot;},{&quot;name&quot;:&quot;相册&quot;,&quot;href&quot;:&quot;4&quot;},{&quot;name&quot;:&quot;留言板&quot;,&quot;href&quot;:&quot;334&quot;},{&quot;name&quot;:&quot;说说&quot;,&quot;href&quot;:&quot;311&quot;},{&quot;name&quot;:&quot;个人档&quot;,&quot;href&quot;:&quot;1&quot;},{&quot;name&quot;:&quot;音乐&quot;,&quot;href&quot;:&quot;305&quot;},{&quot;name&quot;:&quot;更多&quot;,&quot;href&quot;:&quot;more&quot;}]&amp;tips_str=">
				<![endif]-->
				<!--[if !IE]>-->
					<object id="customMenu" name="customMenu" data="//qzonestyle.gtimg.cn/qzone/space_item/orig/14/123054.swf" width="528" height="70" type="application/x-shockwave-flash">
						<param name="movie" value="//qzonestyle.gtimg.cn/qzone/space_item/orig/14/123054.swf">
						<param name="allowScriptAccess" value="always">
						<param name="allownetworking" value="all">
						<param name="allowFullScreen" value="true">
						<param name="wmode" value="transparent">
						<param name="menu" value="false">
						<param name="scale" value="noScale">
						<param name="salign" value="1">
						<param name="flashvars" value="xml_path=//qzs.qq.com/qzone/v5/custom_menu/custom_menu.xml&amp;json_str=[{&quot;name&quot;:&quot;主页&quot;,&quot;href&quot;:&quot;N1&quot;},{&quot;name&quot;:&quot;日志&quot;,&quot;href&quot;:&quot;2&quot;},{&quot;name&quot;:&quot;相册&quot;,&quot;href&quot;:&quot;4&quot;},{&quot;name&quot;:&quot;留言板&quot;,&quot;href&quot;:&quot;334&quot;},{&quot;name&quot;:&quot;说说&quot;,&quot;href&quot;:&quot;311&quot;},{&quot;name&quot;:&quot;个人档&quot;,&quot;href&quot;:&quot;1&quot;},{&quot;name&quot;:&quot;音乐&quot;,&quot;href&quot;:&quot;305&quot;},{&quot;name&quot;:&quot;更多&quot;,&quot;href&quot;:&quot;more&quot;}]&amp;tips_str=">
						<!--<![endif]-->
					<!--[if !IE]>-->
				</object>
			<!--<![endif]-->
			<!--[if IE]></object><![endif]-->
			</div></div>




            <div class="shop-item" id="visitorsDiv" data-left="750" data-top="186" style="top: 186px; left: 828px; width: 218px; height: 56px; cursor: pointer;"><div class="visit-module">            <div class="other-info">                <div class="today-wrapper">                    <p class="visit-today">今日访客<span class="count count-margin ">16</span></p>                    <p class="visit-refuse">被挡访客<span class="count ">402</span></p>                </div>                <div class="count-wrapper">                    <p class="visit-count">访客总量<span class="count ">11120</span></p>                </div>                <i class="sn-count" style="display: none;"></i>            </div>            <i class="icon-visit icon-star-2" title="二星"></i>        </div></div></div>


    <div class="actions profile-hd-actions" id="like_button_pannel">









        <span class="btn-head btn-head-like">
            <a href="javascript:;" class="lnk-left" id="view_like_list" title="4人赞过" style="text-decoration: none;">
                <s class="ui-icon sp-btn-like"></s>
                <span class="txt good-num">4</span>
                <span class="txt message-num none"></span>
            </a>
        </span>
    </div>

        <div id="flowerContainerDiv" class="layout-shop-item"></div></div>
    </div>

    <div class="layout-nav">
        <div class="layout-nav-inner">
            <div class="head-avatar">
                <a id="QM_OwnerInfo_ModifyIcon" title="" href="javascript:;">
                    <div class="head-dress" style="pointer-events:none"></div>
                    <img src="//qlogo1.store.qq.com/qzone/294477044/294477044/100?1382948711" class="user-avatar" id="QM_OwnerInfo_Icon">
                </a>
            </div>
        </div>
    </div>
    <div class="layout-background">

    <div class="layout-body">

        <div id="pageContent" class="layout-page clearfix">

            <div id="main_feed_container" class="col-main">
                <div class="col-main-feed">
                    <div id="QM_Mood_Poster_Container" data-version="20131111">







<div class="qz-poster  bg b-test" id="QM_Mood_Poster_Inner"><div class="qz-poster-inner qz-poster-2017-09-21"><div class="qz-poster-top none"></div><div class="qz-poster-bd"><div class="qz-poster-editor-cont" id="qz_poster_v4_editor_container_1">            <div class="qz-inputer bor2" data-version="20130625">                <div idprefix="$1"><div x:id="content_content" spellcheck="false" x:idprefix="content" class="textinput textarea c_tx2" contenteditable="true" accesskey="q" style="display:none;" id="$1_content_content" idprefix="$1_content"></div><div x:id="substitutor_content" x:idprefix="substitutor" class="textinput textarea c_tx3" contenteditable="true" tabindex="0" accesskey="q" id="$1_substitutor_content" idprefix="$1_substitutor">说点儿什么吧</div><div x:id="tips" class="mod_at_wrap" style="display:none;" id="$1_tips"><div class="mod_at bg bor2"><div class="mod_at_list"><div x:id="tips_content" class="mod_at_tips c_tx3 bg3" id="$1_tips_content">正在加载中...</div></div></div></div><div x:id="options" id="$1_options"></div></div>            <a href="javascript:void(0);" class="close-fest" title="关闭">关闭</a></div>            <!-- attach -->            <div aria-hidden="true" class="qz-poster-attach-side">                <div class="attach">                    <div class="item bor3 bg4 item-pic  evt_hover" data-hottag="attachbar.pic" style="">   <a href="javascript:void(0);" onclick="return false;" class="pic"><i class="icon icon-pic"></i><span class="txt">照片</span></a>                    </div>                    <div class="item bor3 bg4 item-other" style="">                        <a href="javascript:void(0);" onclick="return false;" class="other"><i class="icon icon-other"></i><span class="txt">其他</span></a>                    </div>  <div class="item bor3 bg4 item-funny" style="">   <a href="javascript:void(0)" onclick="return false;" class="funny evt_hover">    <i class="icon icon-funny"></i>    <span class="txt">趣味表情</span>   </a>   <i class="icon-new-fun"></i>  </div>                </div>                <!-- attach -->  <div class="img-upload-drop cell-4 the_worlds_end" style="">                    <div class="pop-drop bg4 bor2">                        <div class="arr"><b class="in bor2"></b><b class="out bor_bg"></b></div><!-- 箭头 -->                        <ul class="list">   <li class="bor3 evt_click evt_hover qz_poster_btn_local_pic" data-hottag="ATTACHBAR.PIC_FROM_LOCAL">    <a href="javascript:void(0);" class="c_tx3" onclick="return false;"> <i class="icon icon-from-cmp"></i><span class="txt">本地</span></a>   </li>   <li class="bor3 evt_click evt_hover  qz_poster_btn_album_select" data-hottag="ATTACHBAR.PIC_FROM_ALBUM"><!-- hover 加bg2 -->    <a href="javascript:void(0);" class="111 c_tx3" onclick="return false;"> <i class="icon icon-video"></i><span class="txt">相册</span> </a>   </li>   <li class="bor3 evt_click evt_hover qz_poster_btn_screen_capture" data-hottag="ATTACHBAR.SCREENCAPTURE"><!-- hover 加bg2 -->     <a href="javascript:void(0);" class="c_tx3" onclick="return false;"> <i class="icon icon-shot"></i><span class="txt">截屏</span> </a>   </li>   <li class="bor3 evt_click evt_hover qz_poster_btn_meitu" data-hottag="ATTACHBAR.MEITU">                       <a href="javascript:void(0);" class="c_tx3" onclick="return false;"> <i class="icon icon-meitu"></i><span class="txt">拼图</span> </a>   </li>                        </ul>                    </div><!-- drop -->                </div> <div class="funny-img-drop cell-4 none" style="">                    <div class="pop-drop bg4 bor2">                        <div class="arr bg4"><b class="in bor2"></b><b class="out bor_bg"></b></div><!-- 箭头 -->                        <ul class="list">   <li class="evt_click evt_hover bor3" data-hottag="ATTACHBAR.UPUP">    <a href="javascript:void(0)" class="c_tx3" onclick="return false"><i class="icon icon-qutu"></i><span class="txt">趣图</span> </a>   </li>   <li class="evt_click evt_hover bor3" data-hottag="ATTACHBAR.TEMPLATE">    <a href="javascript:void(0)" class="c_tx3" onclick="return false"><i class="icon icon-template"></i><span class="txt">气泡</span></a>   </li>   <li class="evt_click evt_hover bor3" data-hottag="ATTACHBAR.QUTU">    <a href="javascript:void(0)" class="c_tx3" onclick="return false"><i class="icon icon-qutu-emo"></i><span class="txt">表情</span></a>   </li>   <li class="evt_click evt_hover bor3" data-hottag="ATTACHBAR.FLARE_FONT" style="position:relative;">    <a href="javascript:void(0)" class="c_tx3" onclick="return false"><i class="icon icon-flash-img"></i><span class="txt">闪图</span></a><i class="icon-new-fun" style="top:4px;"></i>   </li>                        </ul>                    </div><!-- drop --> </div> <div class="other-attach-drop cell-2 none" style="">                    <div class="pop-drop bg4 bor2">                        <div class="arr bg4"><b class="in bor2"></b><b class="out bor_bg"></b></div><!-- 箭头 -->                        <ul class="list">   <li class="evt_click evt_hover bor3" data-hottag="ATTACHBAR.MUSIC">    <a href="javascript:void(0)" class="c_tx3" onclick="return false"><i class="icon icon-music"></i><span class="txt">音乐</span> </a>   </li>   <li class="evt_click evt_hover bor3" data-hottag="ATTACHBAR.VIDEO">    <a href="javascript:void(0)" class="c_tx3" onclick="return false"><i class="icon icon-video"></i><span class="txt">视频</span></a>   </li>   <li class="evt_click evt_hover bor3 none" data-hottag="ATTACHBAR.FILE">    <a href="javascript:void(0)" class="c_tx3" onclick="return false"><i class="icon icon-attach"></i><span class="txt">附件</span></a>   </li>                        </ul>                    </div><!-- drop -->                </div>                <!-- other-drop -->            </div>            <!-- attach-side -->   </div><div class="qz-poster-msg none"></div></div><div class="qz-poster-mid none"></div><div class="qz-poster-ft"><div class="qz-poster-attach "><a href="javascript:void(0);" class="emot "><i class="icon icon-emot"></i><i class="icon-red-dot"></i><span class="txt">表情</span></a><a href="javascript:void(0)" class="at "><i class="icon icon-at"></i><span class="txt">好友</span></a><a href="javascript:void(0)" class="topic "><i class="icon icon-topic"></i><span class="txt">话题</span></a></div><div class="qz-poster-sync"><div class="sync-nuts"><a class="sync-qq evt_click" role="checkbox" aria-checked="false" aria-disabled="false" data-hottag="moodposter.syncqq" href="javascript:void(0);" onclick="return false;" title="同步至QQ签名" style=""><i class="icon icon-qq"></i></a><a class="retweet-it evt_click" role="checkbox" aria-checked="false" aria-disabled="false" data-hottag="moodposter.retweetit" onclick="return false;" href="javascript:void(0);" title="同时转发" style="display:none"><i class="icon icon-retweet"></i></a><a class="sync-weibo evt_click" role="checkbox" aria-checked="false" aria-disabled="false" data-hottag="moodposter.syncweibo" href="javascript:void(0)" onclick="return false;" title="同步至腾讯微博" style="display:none"><i class="icon icon-weibo"></i></a><span class="dividor" style="display:none">|</span><span class="ui-checkbox ui-checkbox-tint js-private-comment" title="" style="display:none"><input type="checkbox" data-hottag="moodposter.private" class="evt_click" id="private_checkbox_1"><label for="private_checkbox_1">私密评论<i class="ui-icon icon-diamond"></i></label></span><a class="sync-timing evt_click" role="checkbox" aria-checked="false" aria-disabled="false" data-hottag="moodposter.timer" href="javascript:void(0);" onclick="return false;" title="发表为定时说说" style="display:none"><i class="icon icon-timing"></i><span class="txt">定时说说</span></a><div class="set-audience" data-hottag="moodposter.secret" style=""><span class="dividor c_tx3">|</span><span class="private-label c_tx3">可见范围：</span><a href="javascript:void(0);" class="audience c_tx3" role="button" aria-haspop="true"><span class="text-place">所有人可见</span><b class="ui_trigbox ui_trigbox_b"><b class="ui_trig c_tx3"></b><b class="ui_trig ui_trig_up bor_bg"></b></b></a></div></div><div class="set-audience-drop none"></div><div class="sync-timing-drop none" style="right:106px;"><div class="pop-drop bg bor2"><div class="arr"><b class="in bor2"></b><b class="out bor_bg"></b></div><ul class="list"><li class="bg5_hover"><a href="javascript:void(0)" onclick="return false;" class="evt_click c_tx2" data-hottag="moodposter.selectPublishTime">设置发布时间</a></li><li class="bg5_hover"><a href="javascript:void(0)" onclick="return false;" class="evt_click c_tx2" data-hottag="moodposter.mytimershuoshuo">我的定时列表</a></li><li class="bg5_hover cancel-timer none"><a href="javascript:void(0)" onclick="return false;" class="evt_click c_tx2" data-hottag="moodposter.canceltimer">取消本条定时说说</a></li></ul></div></div><div class="sync-ex-drop"></div></div><div class="mod-say-authority none"></div><div class="op"><a href="javascript:void(0)" onclick="return false;" class="btn-dismiss-post bgr2 c_tx2 evt_click none" data-hottag="MOODPOSTER.CANCEL"><span class="txt">取消</span></a><a href="javascript:void(0)" onclick="return false;" class="btn-post gb_bt  evt_click" data-hottag="MOODPOSTER.POST"><i class="icon icon-loading"></i><span class="txt">发表</span></a></div><div class="qz-poster-wall-con"></div><div class="tmp-sync-list"></div></div><div class="qz-poster-bottom"><div class="qz-poster-bottom-ext-top none"></div><div class="attachbar-holder the_worlds_end"></div><div class="qz-poster-status-holder the_worlds_end"></div><div class="qz-poster-bottom-ext none"></div></div></div></div>
</div>









<div style="display:none" data-isfest="0"></div>



                    <div id="feed_friend" class=""><div class="feed-v9"><div class="update-more check-more" id="feed_friend_updatetips"><p class="b-line"><i class="radio-hot"></i><span id="feed_friend_update_nick"><a href="javascript:;">？              ..等动态有更新</a></span><i class="fui-icon icon-arrow-down"></i></p></div></div>
                        <div class="fn-feed-control-v2 clearfix">
                            <div class="control-inner">
                                <div class="feed-control-tab">
                                    <a id="feed_tab_hover" class="item-on item-on-slt" href="javascript:;">
                                        <span id="feed_hover_text">全部动态</span>
                                        <i class="ui-icon icon-feed-down"></i>
                                    </a>
                                    <div class="qz-bubble tab-bubble" style="display:none;" id="friend_feed_control">
                                        <div class="bubble-i">
                                            <div class="op-list">
                                                <a id="feed_tab_all" class="item item-slt" href="javascript:;">全部动态</a>
                                                <a id="feed_tab_photo" class="item" href="javascript:;">相册</a>
                                                <a id="feed_tab_blog" class="item" href="javascript:;">日志</a>
                                                <a id="feed_tab_renzheng" class="item none" href="javascript:;">认证空间</a>
                                            </div>
                                        </div>
                                    </div>
                                </div>
                                <div class="feed-control-tab-op">
                                    <a id="feed_friend_refresh" title="刷新动态" href="javascript:;" class="item-left"><i class="ui-icon  icon-refresh-v9"></i></a>
                                    <span class="line"></span>
                                    <a id="feed_friend_set" title="动态设置" href="javascript:;" class="item-right"><i class="ui-icon  icon-set-v9"></i></a>
                                </div>
                            </div>
                        </div>
                        <div class="fn-feed-container">
                            <div class="feed  feed-v9">
                                <div class="feed_inner">
                                    <ul id="feed_friend_list" style="display: block;"><li class="feed-fn-loading" style="display: none;"><div class="update-more loading-more" style="text-align:center;height:2000px;border:0px;cursor:default;"><p><img style="margin-top:50px;" src="http://qzonestyle.gtimg.cn/aoi/img/icenter/loading.gif" class="img-loading"></p></div></li><li class="f-single f-s-s" id="fct_1010140943_311_0_1506002803_0_1" read="1"><div class="f-single-head f-aside"><div class="f-adorn-top"></div>    <div class="user-pto"><a href="http://user.qzone.qq.com/1010140943" target="_blank" class="user-avatar q_namecard f-s-a" link="nameCard_1010140943" data-clicklog="avatar"><img src="https://qlogo4.store.qq.com/qzone/1010140943/1010140943/50"></a></div><div class="user-op"><a href="javascript:;" class="arrow-down" data-cmd="qz_opr_more" data-moreoperate="1"><i class="fui-icon icon-arrow-down"></i></a>        </div><div class="user-info"><div class="f-nick"><a target="_blank" href="http://user.qzone.qq.com/1010140943" data-clicklog="nick" class="f-name q_namecard " link="nameCard_1010140943">那些年。不曾忘記</a>  </div><div class="info-detail"><span class=" ui-mr8 state"> 22:06</span><a href="javascript:;" data-cmd="qz_sign" class="f-sign-show state" title="我也要设置"></a></div></div></div><div class="f-single-content f-wrap"><div class="f-item f-s-i" id="feed_1010140943_311_0_1506002803_0_1" data-feedsflag="" data-iswupfeed="1" data-key="0f87353c73c7c35906590400" data-specialtype="" data-extend-info="0_0_0_0_0_0_0|185495a00102c001|0048000000000000" data-functype="func_friend_feed" data-hasfollowed="1"><div class="f-info">一天抢几十个红包，最大就是这个了<img src="http://qzonestyle.gtimg.cn/qzone/em/e100007.gif" title=""><img src="http://qzonestyle.gtimg.cn/qzone/em/e100007.gif" title=""><img src="http://qzonestyle.gtimg.cn/qzone/em/e100007.gif" title=""><img src="http://qzonestyle.gtimg.cn/qzone/em/e100007.gif" title=""></div><div class="qz_summary wupfeed" id="hex_1010140943_311_0_1506002803_0_1"><i class="none" name="feed_data" data-bitmap="185495a00102c001" data-yybitmap="0048000000000000" data-vipstarbitmap="0000000000000000" data-fkey="0f87353c73c7c35906590400" data-tid="0f87353c73c7c35906590400" data-uin="1010140943" data-origfkey="" data-origtid="0f87353c73c7c35906590400" data-origuin="1010140943" data-subid="" data-totweet="" data-issignin="" data-source="" data-retweetcount="0" data-stat="OrykmfNMKB5LnQTMaOWypSr/4MtCRJ2S0OWYX2xWaFkzajicOJy49fX8H4IQQpJVS!9645U7El!kOqOeLVIt2Zs4LqxuruChkZu1jT6UTMyqt8keKhBM5VsYDNkcd3GlPNC7J3Y25fc7IvI3aIufU3R!xXo9wdavh2ZbASS7f!OUNoF8lLpjgyHgVSovENw73urF3UhC0LYhsLME3F8v6I4R8U1WkdKc_" data-topicid="1010140943_0f87353c73c7c35906590400__1" data-feedstype="100" data-abstime="1506002803" data-iswupfeed="1" data-platformid="54" data-accessright="1"></i><div class="f-ct "><div class="f-ct-txtimg fui-txtimg  "><div class="txt-box ">                                    </div><div class="img-box "><a class="img-item  " data-cmd="qz_popup" href="https://user.qzone.qq.com/1010140943/311/" data-topicid="1010140943_0f87353c73c7c35906590400__1" data-pickey="0f87353c73c7c35906590400,1010140943,V14UZqco3V76O2,NDR0D4c1PHTHw1n1*ooGPgEAAAAAAAA!" data-clicklog="pic" data-originurl="||" hotclickpath="0_appid_311_v8.pic_count_1.pic_0" hotdomain="icv6act.qzone.qq.com" data-version="2" data-param="0f87353c73c7c35906590400|1010140943|0" data-src="/qzone/photo/zone/icenter_popup.html" data-width="810" data-height="1440" data-type="popup" data-title="" data-config=""><img src="http://a3.qpic.cn/psb?/V14UZqco3V76O2/EGIHCqfQNLZ0GsXKZ8zrRtXRHJC2kBDvtY7p2eefXbo!/c/dD4BAAAAAAAA&amp;ek=1&amp;kp=1&amp;pt=0&amp;bo=KgOgBQAAAAAAAKo!&amp;t=5&amp;vuin=294477044&amp;tm=1506006000&amp;sce=60-2-2&amp;rf=0-0" style="height:420px;"></a></div></div></div></div></div></div><div class="f-single-foot"><div class="f-op-detail f-detail "><p class="op-list"><a class="item qz_retweet_btn " href="javascript:;" data-cmd="qz_popup" data-version="4" data-isnewtype="1" data-type="ForwardingBox" data-src="/qzone/app/controls/forwardingBox/forwardingBoxFacade.js" data-clicklog="retweet"><i class="fui-icon icon-op-share"></i></a><span class="item-line"></span>&nbsp;<a href="javascript:;" data-version="6.3" data-cmd="qz_reply" data-link="1" data-clicklog="comment" class=" qz_btn_reply item "><i class="fui-icon icon-op-comment"></i></a>&nbsp;<span class="item-line"></span><a class="item qz_like_btn_v3 " data-islike="0" data-likecnt="0" data-showcount="0" data-unikey="http://user.qzone.qq.com/1010140943/mood/0f87353c73c7c35906590400" data-curkey="http://user.qzone.qq.com/1010140943/mood/0f87353c73c7c35906590400" data-clicklog="like" href="javascript:;"><i class="fui-icon icon-op-praise"></i></a></p><a href="javascript:;" class="state qz_feed_plugin" data-role="Visitor" data-config="311|0f87353c73c7c35906590400|1010140943" data-clicklog="visitor">浏览8次</a></div>        <div class="mod-comments" style="padding:0"><div class="mod-commnets-poster feedClickCmd comment_default_inputentry" data-cmd="qz_reply" data-version="6" data-action="http://taotao.qq.com/cgi-bin/emotion_cgi_re_feeds" data-param="t1_source=&amp;t1_uin=1010140943&amp;t1_tid=0f87353c73c7c35906590400&amp;signin=0&amp;sceneid=100" data-charset="utf-8" data-maxlength="" data-tuin="1010140943" data-config="1|1|1|1,b52,with_fwd,同时转发;0|1,taotaoact.qzone.qq.com,@InputReply|1,taotaoact.qzone.qq.com,@ClickReply|1,taotaoact.qzone.qq.com,commentPresentClick" data-tid=""><div class="comments-poster-bd comments-poster-default"><div class="comments-box" data-clicklog="comment"><div class="textinput textinput-default bor2" contenteditable="true" alt="replybtn" placeholder="评论"><a class="c_tx3" href="javascript:void(0);" alt="replybtn">评论</a></div><div class="mod-insert-img"><a href="javascript:;" data-cmd="qz_quick_upload_img" class="btn-insert-img bg"><i class="icon-camera-16"></i></a></div></div></div></div></div>    </div></li><li class="f-single f-s-s" id="fct_595533711_311_0_1505999535_0_1" read="1"><div class="f-single-head f-aside"><div class="f-adorn-top"></div>    <div class="user-pto"><a href="http://user.qzone.qq.com/595533711" target="_blank" class="user-avatar q_namecard f-s-a" link="nameCard_595533711" data-clicklog="avatar"><img src="https://qlogo4.store.qq.com/qzone/595533711/595533711/50?1420426936"></a></div><div class="user-op"><a href="javascript:;" class="arrow-down" data-cmd="qz_opr_more" data-moreoperate="1"><i class="fui-icon icon-arrow-down"></i></a>        </div><div class="user-info"><div class="f-nick"><a target="_blank" href="http://user.qzone.qq.com/595533711" data-clicklog="nick" class="f-name q_namecard " link="nameCard_595533711">熊吉海</a>  </div><div class="info-detail"><span class=" ui-mr8 state"> 21:12</span><a href="javascript:;" data-cmd="qz_sign" class="f-sign-show state" title="我也要设置"></a></div></div></div><div class="f-single-content f-wrap"><div class="f-item f-s-i" id="feed_595533711_311_0_1505999535_0_1" data-feedsflag="" data-iswupfeed="1" data-key="8f1f7f23afbac35963ce0500" data-specialtype="" data-extend-info="0_0_0_0_0_0_0|984095a40100e009|00001000549cac05" data-functype="func_friend_feed" data-hasfollowed="1"><div class="qz_summary wupfeed" id="hex_595533711_311_0_1505999535_0_1"><i class="none" name="feed_data" data-bitmap="984095a40100e009" data-yybitmap="00001000549cac05" data-vipstarbitmap="0000000000000000" data-fkey="8f1f7f23afbac35963ce0500" data-tid="8f1f7f23afbac35963ce0500" data-uin="595533711" data-origfkey="" data-origtid="8f1f7f23afbac35963ce0500" data-origuin="595533711" data-subid="" data-totweet="" data-issignin="" data-source="" data-retweetcount="0" data-stat="75p2HcTdgjzFZp1CtGmTvo6WgzPcJtM8WThewh!mhq!TDDX8irKsFSAych1rdY/eq3nsjaXBv6YPmmf1TUic8MIRIlRA1dNkwr!TVDZz4YEW2xz/hdsLWr!WHVu9Z3FCf!mdv4HfHsHQ2K1Ng6QNBUEaH5UddivrrcB0FdGZrrbj0QD1jJ!wWMqRXYJxIjy5_" data-topicid="595533711_8f1f7f23afbac35963ce0500__1" data-feedstype="100" data-abstime="1505999535" data-iswupfeed="1" data-platformid="52" data-accessright="1"></i><div class="f-ct "><div class="f-ct-txtimg fui-txtimg  "><div class="txt-box ">                                    </div><div class="img-box "><a class="img-item  " data-cmd="qz_popup" href="https://user.qzone.qq.com/595533711/311/" data-topicid="595533711_8f1f7f23afbac35963ce0500__1" data-pickey="8f1f7f23afbac35963ce0500,http://b242.photo.store.qq.com/psb?/eaaeb4e8-e297-4b57-a2ee-66f0a4ab95ce/bL2RSIS0dnMogvh3.SmCPH5cC5h6qQc675GGtv3zHB0!/b/dPIAAAAAAAAA&amp;bo=KgMABSoDAAURECc!" data-clicklog="pic" data-originurl="||" hotclickpath="0_appid_311_v8.pic_count_1.pic_0" hotdomain="icv6act.qzone.qq.com" data-version="2" data-param="8f1f7f23afbac35963ce0500|595533711|0" data-src="/qzone/photo/zone/icenter_popup.html" data-width="810" data-height="1280" data-type="popup" data-title="" data-config=""><img src="http://a3.qpic.cn/psb?/eaaeb4e8-e297-4b57-a2ee-66f0a4ab95ce/bL2RSIS0dnMogvh3.SmCPH5cC5h6qQc675GGtv3zHB0!/c/dPIAAAAAAAAA&amp;ek=1&amp;kp=1&amp;pt=0&amp;bo=KgMABSoDAAURECc!&amp;t=5&amp;vuin=294477044&amp;tm=1506006000&amp;sce=60-2-2&amp;rf=0-0" style="height:420px;"></a></div></div></div></div></div></div><div class="f-single-foot"><div class="f-op-detail f-detail content-line"><p class="op-list"><a class="item qz_retweet_btn " href="javascript:;" data-cmd="qz_popup" data-version="4" data-isnewtype="1" data-type="ForwardingBox" data-src="/qzone/app/controls/forwardingBox/forwardingBoxFacade.js" data-clicklog="retweet"><i class="fui-icon icon-op-share"></i></a><span class="item-line"></span>&nbsp;<a href="javascript:;" data-version="6.3" data-cmd="qz_reply" data-link="1" data-clicklog="comment" class=" qz_btn_reply item "><i class="fui-icon icon-op-comment"></i></a>&nbsp;<span class="item-line"></span><a class="item qz_like_btn_v3 " data-islike="0" data-likecnt="4" data-showcount="3" data-unikey="http://user.qzone.qq.com/595533711/mood/8f1f7f23afbac35963ce0500" data-curkey="http://user.qzone.qq.com/595533711/mood/8f1f7f23afbac35963ce0500" data-clicklog="like" href="javascript:;"><i class="fui-icon icon-op-praise"></i></a></p><a href="javascript:;" class="state qz_feed_plugin" data-role="Visitor" data-config="311|8f1f7f23afbac35963ce0500|595533711" data-clicklog="visitor">浏览32次</a></div>        <div class="f-ang-t"></div><div class="f-like-list f-like _likeInfo" likeinfo="4"><div class="icon-btn"><a href="javascript:;" data-islike="0" data-likecnt="4" data-showcount="4" data-unikey="http://user.qzone.qq.com/595533711/mood/8f1f7f23afbac35963ce0500" data-curkey="http://user.qzone.qq.com/595533711/mood/8f1f7f23afbac35963ce0500" data-clicklog="like" class="praise qz_like_prase"><i class="fui-icon icon-list-praise"></i></a><div class="bubble" style="display:none;"><div class="bd">+1</div><b class="arrow arrow-down"></b></div></div><div class="user-list"><a href="http://user.qzone.qq.com/374189364" class="item q_namecard" target="_blank" link="nameCard_374189364 des_374189364">熊吉英</a>、<a href="http://user.qzone.qq.com/1486657508" class="item q_namecard" target="_blank" link="nameCard_1486657508 des_1486657508">♚ 心软是病°</a>、<a href="http://user.qzone.qq.com/1506519712" class="item q_namecard" target="_blank" link="nameCard_1506519712 des_1506519712">薰衣草的绚烂</a>等<span class="f-like-cnt">4</span>人觉得很赞</div></div><div class="mod-comments" style="padding:0"><div class="comments-list "><ul><li class="comments-item bor3" data-type="commentroot" data-tid="1" data-uin="2267337450" data-nick="毵醴" data-who="1"><div class="comments-item-bd"><div class="single-reply"><div class="ui-avatar"><a href="http://user.qzone.qq.com/2267337450" target="_blank"><img class="q_namecard" link="nameCard_2267337450 des_2267337450" alt="毵醴" src="http://qlogo3.store.qq.com/qzone/2267337450/2267337450/30"></a></div><div class="comments-content"><a class="nickname name c_tx q_namecard" link="nameCard_2267337450" target="_blank" href="http://user.qzone.qq.com/2267337450">毵醴</a>&nbsp; : 我不会逃避也不会迟疑，愿去尝试生命不同的绮丽。——《疯狂动物城》<div class="comments-op"><span class=" ui-mr10 state"> 21:59</span><a class="act-reply none" href="javascript:;" data-cmd="qz_reply" data-version="6.4" data-action="http://taotao.qq.com/cgi-bin/emotion_cgi_re_feeds" data-param="t1_source=&amp;t1_uin=595533711&amp;t1_tid=8f1f7f23afbac35963ce0500&amp;t2_uin=2267337450&amp;t2_tid=1&amp;subdotype=55702&amp;signin=0&amp;sceneid=100" data-charset="utf-8" data-tuin="" data-config="1|1|1|0|1,taotaoact.qzone.qq.com,@InputReply|1,taotaoact.qzone.qq.com,@ClickReply|1,taotaoact.qzone.qq.com,commentPresentClick"><b class="hide-clip">回复</b></a></div></div></div></div></li></ul></div><div class="mod-commnets-poster feedClickCmd comment_default_inputentry" data-cmd="qz_reply" data-version="6" data-action="http://taotao.qq.com/cgi-bin/emotion_cgi_re_feeds" data-param="t1_source=&amp;t1_uin=595533711&amp;t1_tid=8f1f7f23afbac35963ce0500&amp;signin=0&amp;sceneid=100" data-charset="utf-8" data-maxlength="" data-tuin="595533711" data-config="1|1|1|1,b52,with_fwd,同时转发;0|1,taotaoact.qzone.qq.com,@InputReply|1,taotaoact.qzone.qq.com,@ClickReply|1,taotaoact.qzone.qq.com,commentPresentClick" data-tid=""><div class="comments-poster-bd comments-poster-default"><div class="comments-box" data-clicklog="comment"><div class="textinput textinput-default bor2" contenteditable="true" alt="replybtn" placeholder="评论"><a class="c_tx3" href="javascript:void(0);" alt="replybtn">评论</a></div><div class="mod-insert-img"><a href="javascript:;" data-cmd="qz_quick_upload_img" class="btn-insert-img bg"><i class="icon-camera-16"></i></a></div></div></div></div></div>    </div></li><li class="f-single f-s-s f-single-biz" id="fct_0_6600_4_1506006854_0_1" read="1"><div class="f-single-top"><span>广告</span><a href="javascript:;" class="arrow-down adFeedsItem" data-cmd="qz_opr_more"><i class="fui-icon icon-arrow-down"></i></a><div class="qz-bubble" style="display: none;">            <div class="bubble-i op-list">                <a href="javascript:;" class="item hot-close"><i class="fui-icon icon-hide"></i>关闭</a>            </div>            <b class="bubble-trig ui-trigbox ui-trigbox-t">                <b class="ui-trig"></b>                <b class="ui-trig ui-trig-up"></b>            </b>        </div></div><div class="f-wrap"><div class="f-item"><i class="none" name="feed_data" data-bitmap="0000000000000000" data-yybitmap="0000000000000000" data-vipstarbitmap="0000000000000000" data-fkey="advertisement_outlink_0" data-tid="advertisement_app" data-uin="0" data-origfkey="" data-origtid="advertisement_app" data-origuin="0" data-subid="" data-totweet="" data-issignin="" data-source="" data-retweetcount="0" data-stat="/LyQj/2KdLSUKuiIupFewcgHyWqxbhLwuAfKZ1xwXDuUKuiIupFewcgHyWqxbhLw3hLsSXvdVJaA6ABc!VtPEdBCy2QqfRDddsJiX!tuzLUS6os4DUT8Z9TO/7oaFbVUaFnEnC9TBDot8Wr2GtXACCdhOillG!QJ_" data-feedstype="100" data-abstime="1506006854" data-iswupfeed="1" data-platformid="0" data-advfeed-click-url="http://c.gdt.qq.com/gdt_click.fcg?viewid=7YuqCNwgU37S4MjHF3SRlR4R3s!VFU4zd1wbXVTbChQVvroheCLbESo!P_WpIVldUIX5CubfQCKde76!GzmZfG4ieTyW5d60soOXbCW6ryUj5hqxTZYf9cekegG2hbP8otCa1WMrbdf1aP3b8L7eAbjiZMaPiLmqLj7zSkpa8uUUBuSsh7ypf7tJ9upXsN2r4kzzkcgIbw90XCYlP4U7SQNLl77nQTxqDi5A95pCBb9fwhpG8yB5sQ&amp;jtype=0&amp;wapver=0&amp;i=1&amp;os=7&amp;ad_container=0&amp;container_index=1" data-advfeed-pv-url="http://v.gdt.qq.com/gdt_stats.fcg?viewid=7YuqCNwgU37S4MjHF3SRlR4R3s!VFU4zd1wbXVTbChQVvroheCLbESo!P_WpIVldUIX5CubfQCKde76!GzmZfG4ieTyW5d60soOXbCW6ryUj5hqxTZYf9cekegG2hbP8otCa1WMrbdf1aP3b8L7eAbjiZMaPiLmqLj7zSkpa8uUUBuSsh7ypf7tJ9upXsN2r4kzzkcgIbw90XCYlP4U7SQNLl77nQTxqDi5A95pCBb9fwhpG8yB5sQ&amp;i=1&amp;os=7&amp;ad_container=0&amp;container_index=1"></i><div class="qz_summary wupfeed" id="hex_0_6600_4_1506006854_0_1"><div class="f-ct"><div class="f-vqz-ad">                        <div class="f-single-head"><div class="f-adorn-top"></div><div class="user-pto">                                <a class="user-avatar" target="_blank" href="javascript:;">                                    <img src="http://pgdt.gtimg.cn/gdt/0/DAAJT90ABQABQAACBZZLDKCwkpbB4t.jpg/0?ck=368977316e7a6b67408fba0363fa3ca0">                                </a>                            </div><div class="user-info"><div class="f-nick"><a class="f-name" href="javascript:;" target="_blank">中国光大银行</a></div><div class="info-detail">                                <span class="state ui-mr8"> 23:14</span>                            </div></div></div><div class="f-single-content f-ct-txtimg"><div class="f-info">                            终身免年费信用卡，额度可达5万，免费派发，不要白不要                        </div>                        <div class="f-ct">                            <div class="fui-txtimg">                                <div class="img-box">                                    <a href="http://c.gdt.qq.com/gdt_click.fcg?viewid=7YuqCNwgU37S4MjHF3SRlR4R3s!VFU4zd1wbXVTbChQVvroheCLbESo!P_WpIVldUIX5CubfQCKde76!GzmZfG4ieTyW5d60soOXbCW6ryUj5hqxTZYf9cekegG2hbP8otCa1WMrbdf1aP3b8L7eAbjiZMaPiLmqLj7zSkpa8uUUBuSsh7ypf7tJ9upXsN2r4kzzkcgIbw90XCYlP4U7SQNLl77nQTxqDi5A95pCBb9fwhpG8yB5sQ&amp;jtype=0&amp;wapver=0&amp;i=1&amp;os=7&amp;ad_container=0&amp;container_index=1&amp;acttype=1014" target="_blank" data-clicklog="pic_jump"><img src="http://pgdt.gtimg.cn/gdt/0/DAAJT90AIAAB4AAcBZmqg-CjGXp8gC.jpg/0?ck=6f84f450a022a110ca46a47a42ce742c" style="width:560px"></a>                                </div>                                <div class="detail-box clearfix">                                    <p class="info tips">                                            </p>                                    <div class="fixed-right">                                        <a href="http://c.gdt.qq.com/gdt_click.fcg?viewid=7YuqCNwgU37S4MjHF3SRlR4R3s!VFU4zd1wbXVTbChQVvroheCLbESo!P_WpIVldUIX5CubfQCKde76!GzmZfG4ieTyW5d60soOXbCW6ryUj5hqxTZYf9cekegG2hbP8otCa1WMrbdf1aP3b8L7eAbjiZMaPiLmqLj7zSkpa8uUUBuSsh7ypf7tJ9upXsN2r4kzzkcgIbw90XCYlP4U7SQNLl77nQTxqDi5A95pCBb9fwhpG8yB5sQ&amp;jtype=0&amp;wapver=0&amp;i=1&amp;os=7&amp;ad_container=0&amp;container_index=1&amp;acttype=1025" class="fixed-btn" target="_blank" data-clicklog="follow">去看看</a>                                    </div>                                </div>                            </div>                        </div></div></div></div></div></div></div></li><li class="f-single f-s-s" id="fct_610018923_4_0_1505998636_0_1" read="1"><div class="f-single-head f-aside"><div class="f-adorn-top"></div>    <div class="user-pto"><a href="http://user.qzone.qq.com/610018923" target="_blank" class="user-avatar q_namecard f-s-a" link="nameCard_610018923" data-clicklog="avatar"><img src="https://qlogo4.store.qq.com/qzone/610018923/610018923/50?1400138040"></a></div><div class="user-op"><a href="javascript:;" class="arrow-down" data-cmd="qz_opr_more" data-moreoperate="1"><i class="fui-icon icon-arrow-down"></i></a>        </div><div class="user-info"><div class="f-nick"><a target="_blank" href="http://user.qzone.qq.com/610018923" data-clicklog="nick" class="f-name q_namecard " link="nameCard_610018923">梦里有鬼桃</a>  </div><div class="info-detail"><span class=" ui-mr8 state"> 20:57</span><a href="javascript:;" data-cmd="qz_sign" class="f-sign-show state" title="我也要设置"></a></div></div></div><div class="f-single-content f-wrap"><div class="f-item f-s-i" id="feed_610018923_4_0_1505998636_0_1" data-feedsflag="" data-iswupfeed="1" data-key="1505998675248V14BzO970zmaxV" data-specialtype="" data-extend-info="0_0_0_0_0_0_0|981095a00102e201|004a000000000005" data-functype="func_friend_feed" data-hasfollowed="1"><div class="qz_summary wupfeed" id="hex_610018923_4_0_1505998636_0_1"><i class="none" name="feed_data" data-bitmap="981095a00102e201" data-yybitmap="004a000000000005" data-vipstarbitmap="0000000040000000" data-fkey="1505998675248V14BzO970zmaxV" data-tid="V14BzO970zmaxV" data-uin="610018923" data-origfkey="" data-origtid="1505998675248" data-origuin="610018923" data-subid="" data-totweet="" data-issignin="" data-source="" data-retweetcount="0" data-stat="jZB1hcagyqlw1UvIBnGqbbEN2D6b8ev64qjEOdkp9w/1BEHlZewQyr7RSKfgvW6hOtfrQX/3TI/n7VwR4K!EvSbeaMdO0uxAvUOZVjgRYHSDO00G7W83B0PKWS!IwJW8Q8A5PU7H7uVI8gKFwW!IYLlyAPxT6i751UlrzvyxIgRJ1hODPd6YawMb2XhuWr68Q8pZL4jAlbzUpEa/nKiX86GFEJ0dQzejPQyz!Z7zRwg_" data-topicid="V14BzO970zmaxV_NDR0ayZcJCy3w1ltACw6PAEAAAAAAAA!_1505998675248_4" data-feedstype="100" data-abstime="1505998636" data-iswupfeed="1" data-platformid="52" data-batchid="1505998675248" data-accessright="1"></i><div class="f-ct"><div class="f-info"><div class="txt-box txt-big-size">25</div></div><div class="fui-txtimg fui-imgbox-row-wrap"><div class="img-box img-box-row row-three"><a class="img-item  " data-cmd="qz_popup" href="https://h5.qzone.qq.com/init=photo.v7/common/viewer2/index&amp;ownerUin=610018923&amp;appid=4&amp;topicId=V14BzO970zmaxV_NDR0ayZcJCy3w1ltACw6PAEAAAAAAAA!_1505998675248_4&amp;pre=http%3A%2F%2Fa1.qpic.cn%2Fpsb%3F%2FV14BzO970zmaxV%2FxauNt67x8Or2TutRr3begB8Nk7qU2RNXMpTH*XTdgYw!%2Fm%2FdDwBAAAAAAAA%26ek%3D1%26kp%3D1%26pt%3D0%26bo%3DwAMABcADAAURGS4!%26t%3D5%26vuin%3D294477044%26tm%3D1506006000%26sce%3D60-3-3%26rf%3D0-0&amp;useqzfl=1&amp;useinterface=1&amp;noCloseBtn=0&amp;inqq=1" data-topicid="V14BzO970zmaxV" data-pickey="NDR0ayZcJCy3w1ltACw6PAEAAAAAAAA!" data-clicklog="pic" data-originurl="http://r.photo.store.qq.com/psb?/V14BzO970zmaxV/xauNt67x8Or2TutRr3begB8Nk7qU2RNXMpTH*XTdgYw!/r/dDwBAAAAAAAA&amp;ek=1&amp;kp=1&amp;pt=0&amp;bo=wAMABcADAAURGS4!&amp;t=5&amp;vuin=294477044&amp;tm=1506006000&amp;sce=60-0-0&amp;rf=0-0|960|1280" hotclickpath="0_appid_4_v8.pic_count_.pic_0" hotdomain="icv6act.qzone.qq.com" data-version="2" data-param="610018923|V14BzO970zmaxV|NDR0ayZcJCy3w1ltACw6PAEAAAAAAAA!|http://a1.qpic.cn/psb?/V14BzO970zmaxV/xauNt67x8Or2TutRr3begB8Nk7qU2RNXMpTH*XTdgYw!/c/dDwBAAAAAAAA&amp;ek=1&amp;kp=1&amp;pt=0&amp;bo=wAMABcADAAURGS4!&amp;t=5&amp;vuin=294477044&amp;tm=1506006000&amp;sce=60-2-2&amp;rf=0-0" data-src="/qzone/photo/zone/icenter_popup.html" data-width="0" data-height="0" data-type="popup" data-title="" data-config=""><img src="http://a1.qpic.cn/psb?/V14BzO970zmaxV/xauNt67x8Or2TutRr3begB8Nk7qU2RNXMpTH*XTdgYw!/m/dDwBAAAAAAAA&amp;ek=1&amp;kp=1&amp;pt=0&amp;bo=wAMABcADAAURGS4!&amp;t=5&amp;vuin=294477044&amp;tm=1506006000&amp;sce=60-4-3&amp;rf=0-0" style="margin-left:0px;margin-top:-31px;height:246px;width:185px;"></a><a class="img-item  " data-cmd="qz_popup" href="https://h5.qzone.qq.com/init=photo.v7/common/viewer2/index&amp;ownerUin=610018923&amp;appid=4&amp;topicId=V14BzO970zmaxV_NDR0ayZcJCy3w1ltACw6PAEAAAAAAAA!_1505998675248_4&amp;pre=http%3A%2F%2Fa1.qpic.cn%2Fpsb%3F%2FV14BzO970zmaxV%2FxauNt67x8Or2TutRr3begB8Nk7qU2RNXMpTH*XTdgYw!%2Fm%2FdDwBAAAAAAAA%26ek%3D1%26kp%3D1%26pt%3D0%26bo%3DwAMABcADAAURGS4!%26t%3D5%26vuin%3D294477044%26tm%3D1506006000%26sce%3D60-3-3%26rf%3D0-0&amp;useqzfl=1&amp;useinterface=1&amp;noCloseBtn=0&amp;inqq=1" data-topicid="V14BzO970zmaxV" data-pickey="NDR0ayZcJCu3w1lv20gzjgEAAAAAAAA!" data-clicklog="pic" data-originurl="http://r.photo.store.qq.com/psb?/V14BzO970zmaxV/aCneRTo60Vj1O6hUzoETB7jOo4h3b.OW9Nh.H0c.zaA!/r/dI4BAAAAAAAA&amp;ek=1&amp;kp=1&amp;pt=0&amp;bo=wAMABcADAAURGS4!&amp;t=5&amp;vuin=294477044&amp;tm=1506006000&amp;sce=60-0-0&amp;rf=0-0|960|1280" hotclickpath="0_appid_4_v8.pic_count_.pic_1" hotdomain="icv6act.qzone.qq.com" data-version="2" data-param="610018923|V14BzO970zmaxV|NDR0ayZcJCu3w1lv20gzjgEAAAAAAAA!|http://a3.qpic.cn/psb?/V14BzO970zmaxV/aCneRTo60Vj1O6hUzoETB7jOo4h3b.OW9Nh.H0c.zaA!/c/dI4BAAAAAAAA&amp;ek=1&amp;kp=1&amp;pt=0&amp;bo=wAMABcADAAURGS4!&amp;t=5&amp;vuin=294477044&amp;tm=1506006000&amp;sce=60-2-2&amp;rf=0-0" data-src="/qzone/photo/zone/icenter_popup.html" data-width="0" data-height="0" data-type="popup" data-title="" data-config=""><img src="http://a3.qpic.cn/psb?/V14BzO970zmaxV/aCneRTo60Vj1O6hUzoETB7jOo4h3b.OW9Nh.H0c.zaA!/m/dI4BAAAAAAAA&amp;ek=1&amp;kp=1&amp;pt=0&amp;bo=wAMABcADAAURGS4!&amp;t=5&amp;vuin=294477044&amp;tm=1506006000&amp;sce=60-4-3&amp;rf=0-0" style="margin-left:0px;margin-top:0px;height:246px;width:185px;"></a><a class="img-item  " data-cmd="qz_popup" href="https://h5.qzone.qq.com/init=photo.v7/common/viewer2/index&amp;ownerUin=610018923&amp;appid=4&amp;topicId=V14BzO970zmaxV_NDR0ayZcJCy3w1ltACw6PAEAAAAAAAA!_1505998675248_4&amp;pre=http%3A%2F%2Fa1.qpic.cn%2Fpsb%3F%2FV14BzO970zmaxV%2FxauNt67x8Or2TutRr3begB8Nk7qU2RNXMpTH*XTdgYw!%2Fm%2FdDwBAAAAAAAA%26ek%3D1%26kp%3D1%26pt%3D0%26bo%3DwAMABcADAAURGS4!%26t%3D5%26vuin%3D294477044%26tm%3D1506006000%26sce%3D60-3-3%26rf%3D0-0&amp;useqzfl=1&amp;useinterface=1&amp;noCloseBtn=0&amp;inqq=1" data-topicid="V14BzO970zmaxV" data-pickey="NDR0ayZcJCq3w1my*nYU5wAAAAAAAAA!" data-clicklog="pic" data-originurl="http://r.photo.store.qq.com/psb?/V14BzO970zmaxV/aW0B.NJv*zm3kWJM74U4a66Wg*KvZ.aFRNL0OkMjC9Q!/r/dOcAAAAAAAAA&amp;ek=1&amp;kp=1&amp;pt=0&amp;bo=wAMABcADAAURGS4!&amp;t=5&amp;vuin=294477044&amp;tm=1506006000&amp;sce=60-0-0&amp;rf=0-0|960|1280" hotclickpath="0_appid_4_v8.pic_count_.pic_2" hotdomain="icv6act.qzone.qq.com" data-version="2" data-param="610018923|V14BzO970zmaxV|NDR0ayZcJCq3w1my*nYU5wAAAAAAAAA!|http://a4.qpic.cn/psb?/V14BzO970zmaxV/aW0B.NJv*zm3kWJM74U4a66Wg*KvZ.aFRNL0OkMjC9Q!/c/dOcAAAAAAAAA&amp;ek=1&amp;kp=1&amp;pt=0&amp;bo=wAMABcADAAURGS4!&amp;t=5&amp;vuin=294477044&amp;tm=1506006000&amp;sce=60-2-2&amp;rf=0-0" data-src="/qzone/photo/zone/icenter_popup.html" data-width="0" data-height="0" data-type="popup" data-title="" data-config=""><img src="http://a4.qpic.cn/psb?/V14BzO970zmaxV/aW0B.NJv*zm3kWJM74U4a66Wg*KvZ.aFRNL0OkMjC9Q!/m/dOcAAAAAAAAA&amp;ek=1&amp;kp=1&amp;pt=0&amp;bo=wAMABcADAAURGS4!&amp;t=5&amp;vuin=294477044&amp;tm=1506006000&amp;sce=60-4-3&amp;rf=0-0" style="margin-left:0px;margin-top:-16px;height:246px;width:185px;"></a><a class="img-item  " data-cmd="qz_popup" href="https://h5.qzone.qq.com/init=photo.v7/common/viewer2/index&amp;ownerUin=610018923&amp;appid=4&amp;topicId=V14BzO970zmaxV_NDR0ayZcJCy3w1ltACw6PAEAAAAAAAA!_1505998675248_4&amp;pre=http%3A%2F%2Fa1.qpic.cn%2Fpsb%3F%2FV14BzO970zmaxV%2FxauNt67x8Or2TutRr3begB8Nk7qU2RNXMpTH*XTdgYw!%2Fm%2FdDwBAAAAAAAA%26ek%3D1%26kp%3D1%26pt%3D0%26bo%3DwAMABcADAAURGS4!%26t%3D5%26vuin%3D294477044%26tm%3D1506006000%26sce%3D60-3-3%26rf%3D0-0&amp;useqzfl=1&amp;useinterface=1&amp;noCloseBtn=0&amp;inqq=1" data-topicid="V14BzO970zmaxV" data-pickey="NDR0ayZcJCm3w1nlWsEOPQEAAAAAAAA!" data-clicklog="pic" data-originurl="http://r.photo.store.qq.com/psb?/V14BzO970zmaxV/w4sfsEZUmue81cAft5DqAVIu3yNHPCmCRKbOrngRZuI!/r/dD0BAAAAAAAA&amp;ek=1&amp;kp=1&amp;pt=0&amp;bo=wAMABcADAAURGS4!&amp;t=5&amp;vuin=294477044&amp;tm=1506006000&amp;sce=60-0-0&amp;rf=0-0|960|1280" hotclickpath="0_appid_4_v8.pic_count_.pic_3" hotdomain="icv6act.qzone.qq.com" data-version="2" data-param="610018923|V14BzO970zmaxV|NDR0ayZcJCm3w1nlWsEOPQEAAAAAAAA!|http://a2.qpic.cn/psb?/V14BzO970zmaxV/w4sfsEZUmue81cAft5DqAVIu3yNHPCmCRKbOrngRZuI!/c/dD0BAAAAAAAA&amp;ek=1&amp;kp=1&amp;pt=0&amp;bo=wAMABcADAAURGS4!&amp;t=5&amp;vuin=294477044&amp;tm=1506006000&amp;sce=60-2-2&amp;rf=0-0" data-src="/qzone/photo/zone/icenter_popup.html" data-width="0" data-height="0" data-type="popup" data-title="" data-config=""><img src="http://a2.qpic.cn/psb?/V14BzO970zmaxV/w4sfsEZUmue81cAft5DqAVIu3yNHPCmCRKbOrngRZuI!/m/dD0BAAAAAAAA&amp;ek=1&amp;kp=1&amp;pt=0&amp;bo=wAMABcADAAURGS4!&amp;t=5&amp;vuin=294477044&amp;tm=1506006000&amp;sce=60-4-3&amp;rf=0-0" style="margin-left:0px;margin-top:0px;height:246px;width:185px;"></a><a class="img-item  " data-cmd="qz_popup" href="https://h5.qzone.qq.com/init=photo.v7/common/viewer2/index&amp;ownerUin=610018923&amp;appid=4&amp;topicId=V14BzO970zmaxV_NDR0ayZcJCy3w1ltACw6PAEAAAAAAAA!_1505998675248_4&amp;pre=http%3A%2F%2Fa1.qpic.cn%2Fpsb%3F%2FV14BzO970zmaxV%2FxauNt67x8Or2TutRr3begB8Nk7qU2RNXMpTH*XTdgYw!%2Fm%2FdDwBAAAAAAAA%26ek%3D1%26kp%3D1%26pt%3D0%26bo%3DwAMABcADAAURGS4!%26t%3D5%26vuin%3D294477044%26tm%3D1506006000%26sce%3D60-3-3%26rf%3D0-0&amp;useqzfl=1&amp;useinterface=1&amp;noCloseBtn=0&amp;inqq=1" data-topicid="V14BzO970zmaxV" data-pickey="NDR0ayZcJCi3w1lPRsoZhQEAAAAAAAA!" data-clicklog="pic" data-originurl="http://r.photo.store.qq.com/psb?/V14BzO970zmaxV/ZyoWrOTsWw3PKrjbGrJNpBbMErSrTHI4FY0N1Knqk6Y!/r/dIUBAAAAAAAA&amp;ek=1&amp;kp=1&amp;pt=0&amp;bo=0AIABdACAAURGS4!&amp;t=5&amp;vuin=294477044&amp;tm=1506006000&amp;sce=60-0-0&amp;rf=0-0|720|1280" hotclickpath="0_appid_4_v8.pic_count_.pic_4" hotdomain="icv6act.qzone.qq.com" data-version="2" data-param="610018923|V14BzO970zmaxV|NDR0ayZcJCi3w1lPRsoZhQEAAAAAAAA!|http://a2.qpic.cn/psb?/V14BzO970zmaxV/ZyoWrOTsWw3PKrjbGrJNpBbMErSrTHI4FY0N1Knqk6Y!/c/dIUBAAAAAAAA&amp;ek=1&amp;kp=1&amp;pt=0&amp;bo=0AIABdACAAURGS4!&amp;t=5&amp;vuin=294477044&amp;tm=1506006000&amp;sce=60-2-2&amp;rf=0-0" data-src="/qzone/photo/zone/icenter_popup.html" data-width="0" data-height="0" data-type="popup" data-title="" data-config=""><img src="http://a2.qpic.cn/psb?/V14BzO970zmaxV/ZyoWrOTsWw3PKrjbGrJNpBbMErSrTHI4FY0N1Knqk6Y!/m/dIUBAAAAAAAA&amp;ek=1&amp;kp=1&amp;pt=0&amp;bo=0AIABdACAAURGS4!&amp;t=5&amp;vuin=294477044&amp;tm=1506006000&amp;sce=60-4-3&amp;rf=0-0" style="margin-left:0px;margin-top:-71px;height:328px;width:185px;"></a><a class="img-item  " data-cmd="qz_popup" href="https://h5.qzone.qq.com/init=photo.v7/common/viewer2/index&amp;ownerUin=610018923&amp;appid=4&amp;topicId=V14BzO970zmaxV_NDR0ayZcJCy3w1ltACw6PAEAAAAAAAA!_1505998675248_4&amp;pre=http%3A%2F%2Fa1.qpic.cn%2Fpsb%3F%2FV14BzO970zmaxV%2FxauNt67x8Or2TutRr3begB8Nk7qU2RNXMpTH*XTdgYw!%2Fm%2FdDwBAAAAAAAA%26ek%3D1%26kp%3D1%26pt%3D0%26bo%3DwAMABcADAAURGS4!%26t%3D5%26vuin%3D294477044%26tm%3D1506006000%26sce%3D60-3-3%26rf%3D0-0&amp;useqzfl=1&amp;useinterface=1&amp;noCloseBtn=0&amp;inqq=1" data-topicid="V14BzO970zmaxV" data-pickey="NDR0ayZcJCe3w1kh8aExhQEAAAAAAAA!" data-clicklog="pic" data-originurl="http://r.photo.store.qq.com/psb?/V14BzO970zmaxV/eQp**X5chgOjU1GRtN2haL**vtIkJjolBSR2KFqSxrs!/r/dIUBAAAAAAAA&amp;ek=1&amp;kp=1&amp;pt=0&amp;bo=0AIABdACAAURGS4!&amp;t=5&amp;vuin=294477044&amp;tm=1506006000&amp;sce=60-0-0&amp;rf=0-0|720|1280" hotclickpath="0_appid_4_v8.pic_count_.pic_5" hotdomain="icv6act.qzone.qq.com" data-version="2" data-param="610018923|V14BzO970zmaxV|NDR0ayZcJCe3w1kh8aExhQEAAAAAAAA!|http://a2.qpic.cn/psb?/V14BzO970zmaxV/eQp**X5chgOjU1GRtN2haL**vtIkJjolBSR2KFqSxrs!/c/dIUBAAAAAAAA&amp;ek=1&amp;kp=1&amp;pt=0&amp;bo=0AIABdACAAURGS4!&amp;t=5&amp;vuin=294477044&amp;tm=1506006000&amp;sce=60-2-2&amp;rf=0-0" data-src="/qzone/photo/zone/icenter_popup.html" data-width="0" data-height="0" data-type="popup" data-title="" data-config=""><img src="http://a2.qpic.cn/psb?/V14BzO970zmaxV/eQp**X5chgOjU1GRtN2haL**vtIkJjolBSR2KFqSxrs!/m/dIUBAAAAAAAA&amp;ek=1&amp;kp=1&amp;pt=0&amp;bo=0AIABdACAAURGS4!&amp;t=5&amp;vuin=294477044&amp;tm=1506006000&amp;sce=60-4-3&amp;rf=0-0" style="margin-left:0px;margin-top:-71px;height:328px;width:185px;"></a><a class="img-item  " data-cmd="qz_popup" href="https://h5.qzone.qq.com/init=photo.v7/common/viewer2/index&amp;ownerUin=610018923&amp;appid=4&amp;topicId=V14BzO970zmaxV_NDR0ayZcJCy3w1ltACw6PAEAAAAAAAA!_1505998675248_4&amp;pre=http%3A%2F%2Fa1.qpic.cn%2Fpsb%3F%2FV14BzO970zmaxV%2FxauNt67x8Or2TutRr3begB8Nk7qU2RNXMpTH*XTdgYw!%2Fm%2FdDwBAAAAAAAA%26ek%3D1%26kp%3D1%26pt%3D0%26bo%3DwAMABcADAAURGS4!%26t%3D5%26vuin%3D294477044%26tm%3D1506006000%26sce%3D60-3-3%26rf%3D0-0&amp;useqzfl=1&amp;useinterface=1&amp;noCloseBtn=0&amp;inqq=1" data-topicid="V14BzO970zmaxV" data-pickey="NDR0ayZcJCa3w1k66cQhPAEAAAAAAAA!" data-clicklog="pic" data-originurl="http://r.photo.store.qq.com/psb?/V14BzO970zmaxV/4JbNpYTL7Z8fW1T11toxpGh5XuFuK6qpU4VhgVxtBVs!/r/dDwBAAAAAAAA&amp;ek=1&amp;kp=1&amp;pt=0&amp;bo=0AIABdACAAURGS4!&amp;t=5&amp;vuin=294477044&amp;tm=1506006000&amp;sce=60-0-0&amp;rf=0-0|720|1280" hotclickpath="0_appid_4_v8.pic_count_.pic_6" hotdomain="icv6act.qzone.qq.com" data-version="2" data-param="610018923|V14BzO970zmaxV|NDR0ayZcJCa3w1k66cQhPAEAAAAAAAA!|http://a1.qpic.cn/psb?/V14BzO970zmaxV/4JbNpYTL7Z8fW1T11toxpGh5XuFuK6qpU4VhgVxtBVs!/c/dDwBAAAAAAAA&amp;ek=1&amp;kp=1&amp;pt=0&amp;bo=0AIABdACAAURGS4!&amp;t=5&amp;vuin=294477044&amp;tm=1506006000&amp;sce=60-2-2&amp;rf=0-0" data-src="/qzone/photo/zone/icenter_popup.html" data-width="0" data-height="0" data-type="popup" data-title="" data-config=""><img src="http://a1.qpic.cn/psb?/V14BzO970zmaxV/4JbNpYTL7Z8fW1T11toxpGh5XuFuK6qpU4VhgVxtBVs!/m/dDwBAAAAAAAA&amp;ek=1&amp;kp=1&amp;pt=0&amp;bo=0AIABdACAAURGS4!&amp;t=5&amp;vuin=294477044&amp;tm=1506006000&amp;sce=60-4-3&amp;rf=0-0" style="margin-left:0px;margin-top:-71px;height:328px;width:185px;"></a><a class="img-item  " data-cmd="qz_popup" href="https://h5.qzone.qq.com/init=photo.v7/common/viewer2/index&amp;ownerUin=610018923&amp;appid=4&amp;topicId=V14BzO970zmaxV_NDR0ayZcJCy3w1ltACw6PAEAAAAAAAA!_1505998675248_4&amp;pre=http%3A%2F%2Fa1.qpic.cn%2Fpsb%3F%2FV14BzO970zmaxV%2FxauNt67x8Or2TutRr3begB8Nk7qU2RNXMpTH*XTdgYw!%2Fm%2FdDwBAAAAAAAA%26ek%3D1%26kp%3D1%26pt%3D0%26bo%3DwAMABcADAAURGS4!%26t%3D5%26vuin%3D294477044%26tm%3D1506006000%26sce%3D60-3-3%26rf%3D0-0&amp;useqzfl=1&amp;useinterface=1&amp;noCloseBtn=0&amp;inqq=1" data-topicid="V14BzO970zmaxV" data-pickey="NDR0ayZcJCW3w1kxvScvhQEAAAAAAAA!" data-clicklog="pic" data-originurl="http://r.photo.store.qq.com/psb?/V14BzO970zmaxV/LgdIfBqoXwjpJuedpx7E*Q1AmDB74I6q7GFdtg1QpXc!/r/dIUBAAAAAAAA&amp;ek=1&amp;kp=1&amp;pt=0&amp;bo=0AIABdACAAURGS4!&amp;t=5&amp;vuin=294477044&amp;tm=1506006000&amp;sce=60-0-0&amp;rf=0-0|720|1280" hotclickpath="0_appid_4_v8.pic_count_.pic_7" hotdomain="icv6act.qzone.qq.com" data-version="2" data-param="610018923|V14BzO970zmaxV|NDR0ayZcJCW3w1kxvScvhQEAAAAAAAA!|http://a2.qpic.cn/psb?/V14BzO970zmaxV/LgdIfBqoXwjpJuedpx7E*Q1AmDB74I6q7GFdtg1QpXc!/c/dIUBAAAAAAAA&amp;ek=1&amp;kp=1&amp;pt=0&amp;bo=0AIABdACAAURGS4!&amp;t=5&amp;vuin=294477044&amp;tm=1506006000&amp;sce=60-2-2&amp;rf=0-0" data-src="/qzone/photo/zone/icenter_popup.html" data-width="0" data-height="0" data-type="popup" data-title="" data-config=""><img src="http://a2.qpic.cn/psb?/V14BzO970zmaxV/LgdIfBqoXwjpJuedpx7E*Q1AmDB74I6q7GFdtg1QpXc!/m/dIUBAAAAAAAA&amp;ek=1&amp;kp=1&amp;pt=0&amp;bo=0AIABdACAAURGS4!&amp;t=5&amp;vuin=294477044&amp;tm=1506006000&amp;sce=60-4-3&amp;rf=0-0" style="margin-left:0px;margin-top:-71px;height:328px;width:185px;"></a><a class="img-item  " data-cmd="qz_popup" href="https://h5.qzone.qq.com/init=photo.v7/common/viewer2/index&amp;ownerUin=610018923&amp;appid=4&amp;topicId=V14BzO970zmaxV_NDR0ayZcJCy3w1ltACw6PAEAAAAAAAA!_1505998675248_4&amp;pre=http%3A%2F%2Fa1.qpic.cn%2Fpsb%3F%2FV14BzO970zmaxV%2FxauNt67x8Or2TutRr3begB8Nk7qU2RNXMpTH*XTdgYw!%2Fm%2FdDwBAAAAAAAA%26ek%3D1%26kp%3D1%26pt%3D0%26bo%3DwAMABcADAAURGS4!%26t%3D5%26vuin%3D294477044%26tm%3D1506006000%26sce%3D60-3-3%26rf%3D0-0&amp;useqzfl=1&amp;useinterface=1&amp;noCloseBtn=0&amp;inqq=1" data-topicid="V14BzO970zmaxV" data-pickey="NDR0ayZcJCS3w1l2EbAhPQEAAAAAAAA!" data-clicklog="pic" data-originurl="http://r.photo.store.qq.com/psb?/V14BzO970zmaxV/MzHAANuwEVZy57hmHPL9syVPlAIZp6W0XjfNnf.1RfU!/r/dD0BAAAAAAAA&amp;ek=1&amp;kp=1&amp;pt=0&amp;bo=0AIABdACAAURGS4!&amp;t=5&amp;vuin=294477044&amp;tm=1506006000&amp;sce=60-0-0&amp;rf=0-0|720|1280" hotclickpath="0_appid_4_v8.pic_count_.pic_8" hotdomain="icv6act.qzone.qq.com" data-version="2" data-param="610018923|V14BzO970zmaxV|NDR0ayZcJCS3w1l2EbAhPQEAAAAAAAA!|http://a2.qpic.cn/psb?/V14BzO970zmaxV/MzHAANuwEVZy57hmHPL9syVPlAIZp6W0XjfNnf.1RfU!/c/dD0BAAAAAAAA&amp;ek=1&amp;kp=1&amp;pt=0&amp;bo=0AIABdACAAURGS4!&amp;t=5&amp;vuin=294477044&amp;tm=1506006000&amp;sce=60-2-2&amp;rf=0-0" data-src="/qzone/photo/zone/icenter_popup.html" data-width="0" data-height="0" data-type="popup" data-title="" data-config=""><img src="http://a2.qpic.cn/psb?/V14BzO970zmaxV/MzHAANuwEVZy57hmHPL9syVPlAIZp6W0XjfNnf.1RfU!/m/dD0BAAAAAAAA&amp;ek=1&amp;kp=1&amp;pt=0&amp;bo=0AIABdACAAURGS4!&amp;t=5&amp;vuin=294477044&amp;tm=1506006000&amp;sce=60-4-3&amp;rf=0-0" style="margin-left:0px;margin-top:-71px;height:328px;width:185px;"></a><div class="img-num" data-url="https://user.qzone.qq.com/610018923/photo/V14BzO970zmaxV">                <span class="bg"></span>                <span class="num"><i class="fui-icon icon-album-num"></i>12</span>            </div></div></div><div class="f-reprint">                        <p class="item state">        <i class="fui-icon icon-print-album"></i>        <a href="http://user.qzone.qq.com/610018923/photo/V14BzO970zmaxV" target="_blank" style="cursor:pointer;">                上传20张照片/视频到《2016年》        </a>        </p>            </div>    </div></div></div></div><div class="f-single-foot"><div class="f-op-detail f-detail content-line"><p class="op-list"><a class="item qz_retweet_btn " href="javascript:;" data-cmd="qz_popup" data-version="4" data-isnewtype="1" data-type="ForwardingBox" data-src="/qzone/app/controls/forwardingBox/forwardingBoxFacade.js" data-clicklog="retweet"><i class="fui-icon icon-op-share"></i></a><span class="item-line"></span>&nbsp;<a href="javascript:;" data-version="6.3" data-cmd="qz_reply" data-link="1" data-clicklog="comment" class=" qz_btn_reply item "><i class="fui-icon icon-op-comment"></i></a>&nbsp;<span class="item-line"></span><a class="item qz_like_btn_v3 " data-islike="0" data-likecnt="11" data-showcount="11" data-unikey="http://user.qzone.qq.com/610018923/batchphoto/V14BzO970zmaxV/1505998675248" data-curkey="http://user.qzone.qq.com/610018923/batchphoto/V14BzO970zmaxV/1505998675248" data-clicklog="like" href="javascript:;"><i class="fui-icon icon-op-praise"></i></a></p><a href="javascript:;" class="state qz_feed_plugin" data-role="Visitor" data-config="4|33;V14BzO970zmaxV%7C1505998675248|610018923" data-clicklog="visitor">浏览55次</a></div>        <div class="f-ang-t"></div><div class="f-like-list f-like _likeInfo" likeinfo="11"><div class="icon-btn"><a href="javascript:;" data-islike="0" data-likecnt="11" data-showcount="11" data-unikey="http://user.qzone.qq.com/610018923/batchphoto/V14BzO970zmaxV/1505998675248" data-curkey="http://user.qzone.qq.com/610018923/batchphoto/V14BzO970zmaxV/1505998675248" data-clicklog="like" class="praise qz_like_prase"><i class="fui-icon icon-list-praise"></i></a><div class="bubble" style="display:none;"><div class="bd">+1</div><b class="arrow arrow-down"></b></div></div><div class="user-list"><a href="http://user.qzone.qq.com/654858397" class="item q_namecard" target="_blank" link="nameCard_654858397 des_654858397">谭超</a>、<a href="http://user.qzone.qq.com/769914037" class="item q_namecard" target="_blank" link="nameCard_769914037 des_769914037">白大梅</a>、<a href="http://user.qzone.qq.com/383462820" class="item q_namecard" target="_blank" link="nameCard_383462820 des_383462820">～</a>、<a href="http://user.qzone.qq.com/460345424" class="item q_namecard" target="_blank" link="nameCard_460345424 des_460345424">箹錠  川々少</a>、<a href="http://user.qzone.qq.com/502206796" class="item q_namecard" target="_blank" link="nameCard_502206796 des_502206796">夜月/db漂泊</a>、<a href="http://user.qzone.qq.com/625488350" class="item q_namecard" target="_blank" link="nameCard_625488350 des_625488350">忧伤、旋律</a>、<a href="http://user.qzone.qq.com/751279539" class="item q_namecard" target="_blank" link="nameCard_751279539 des_751279539">迷澜之媛</a>、<a href="http://user.qzone.qq.com/863436255" class="item q_namecard" target="_blank" link="nameCard_863436255 des_863436255">爱你￥成依赖</a>、<a href="http://user.qzone.qq.com/1027295304" class="item q_namecard" target="_blank" link="nameCard_1027295304 des_1027295304">彬彬有礼</a>、<a href="http://user.qzone.qq.com/1184748225" class="item q_namecard" target="_blank" link="nameCard_1184748225 des_1184748225">莫失莫忘。</a>、<a href="http://user.qzone.qq.com/2507820334" class="item q_namecard" target="_blank" link="nameCard_2507820334 des_2507820334">思＆</a>共<span class="f-like-cnt">11</span>人觉得很赞</div></div><div class="mod-comments" style="padding:0"><div class="comments-list "><ul><li class="comments-item bor3" data-type="commentroot" data-tid="34451106" data-uin="2507820334" data-nick="思＆" data-who="1"><div class="comments-item-bd"><div class="single-reply"><div class="ui-avatar"><a href="http://user.qzone.qq.com/2507820334" target="_blank"><img class="q_namecard" link="nameCard_2507820334 des_2507820334" alt="思＆" src="http://qlogo3.store.qq.com/qzone/2507820334/2507820334/30"></a></div><div class="comments-content"><a class="nickname name c_tx q_namecard" link="nameCard_2507820334" target="_blank" href="http://user.qzone.qq.com/2507820334">思＆</a>&nbsp; : 好多美女呀<img src="http://qzonestyle.gtimg.cn/qzone/em/e113.png" title=""><img src="http://qzonestyle.gtimg.cn/qzone/em/e113.png" title=""><img src="http://qzonestyle.gtimg.cn/qzone/em/e113.png" title=""><div class="comments-op"><span class=" ui-mr10 state"> 21:32</span><a class="act-reply none" href="javascript:;" data-cmd="qz_reply" data-version="6.4" data-action="http://photo.qq.com/cgi-bin/common/cgi_add_icreply" data-param="oweruin=610018923&amp;albumid=V14BzO970zmaxV&amp;lloc=NDR0ayZcJCy3w1ltACw6PAEAAAAAAAA!&amp;sloc=NDR0ayZcJCy3w1ltACw6PAEAAAAAAAA!&amp;refer=qzone&amp;reqfrom=801&amp;zz=0&amp;cmtid=34451106&amp;batchId=1505998675248&amp;cmtType=4&amp;picname=2017-09-21" data-charset="GB" data-tuin="" data-config="{config:'1|1|0|0|0|0|0',targetUserInfo:{id:2507820334,nick:'鎬濓紗',who:1,auto:1}}"><b class="hide-clip">回复</b></a></div></div></div><div class="comments-list mod-comments-sub"><ul><li class="comments-item bor3" data-type="replyroot" data-tid="1506003303" data-uin="610018923" data-nick="梦里有鬼桃" data-who="1"><div class="comments-item-bd"><div class="single-reply"><div class="ui-avatar"><a href="http://user.qzone.qq.com/610018923" target="_blank"><img class="q_namecard" link="nameCard_610018923 des_610018923" alt="梦里有鬼桃" src="http://qlogo4.store.qq.com/qzone/610018923/610018923/30?1400138040"></a></div><div class="comments-content"><a class="nickname name c_tx q_namecard" link="nameCard_610018923" target="_blank" href="http://user.qzone.qq.com/610018923">梦里有鬼桃</a>&nbsp;回复<a class="nickname name c_tx q_namecard" link="nameCard_2507820334" target="_blank" href="http://user.qzone.qq.com/2507820334">思＆</a>&nbsp; : 去年的照片，我手机没错吧不够了。我上传到qq上面<div class="comments-op"><span class=" ui-mr10 state"> 22:15</span><a class="act-reply none" href="javascript:;" data-cmd="qz_reply" data-version="6.4" data-action="http://photo.qq.com/cgi-bin/common/cgi_add_icreply" data-param="oweruin=610018923&amp;albumid=V14BzO970zmaxV&amp;lloc=NDR0ayZcJCy3w1ltACw6PAEAAAAAAAA!&amp;sloc=NDR0ayZcJCy3w1ltACw6PAEAAAAAAAA!&amp;refer=qzone&amp;reqfrom=801&amp;zz=0&amp;cmtid=34451106&amp;batchId=1505998675248&amp;cmtType=4&amp;picname=2017-09-21" data-charset="GB" data-tuin="" data-config="{config:'1|1|0|0|0|0|0',targetUserInfo:{id:610018923,nick:'',who:1,auto:1}}"><b class="hide-clip">回复</b></a></div></div></div></div></li><li class="comments-item bor3" data-type="replyroot" data-tid="1506003310" data-uin="610018923" data-nick="梦里有鬼桃" data-who="1"><div class="comments-item-bd"><div class="single-reply"><div class="ui-avatar"><a href="http://user.qzone.qq.com/610018923" target="_blank"><img class="q_namecard" link="nameCard_610018923 des_610018923" alt="梦里有鬼桃" src="http://qlogo4.store.qq.com/qzone/610018923/610018923/30?1400138040"></a></div><div class="comments-content"><a class="nickname name c_tx q_namecard" link="nameCard_610018923" target="_blank" href="http://user.qzone.qq.com/610018923">梦里有鬼桃</a>&nbsp;回复<a class="nickname name c_tx q_namecard" link="nameCard_2507820334" target="_blank" href="http://user.qzone.qq.com/2507820334">思＆</a>&nbsp; : 你看哈你<div class="comments-op"><span class=" ui-mr10 state"> 22:15</span><a class="act-reply none" href="javascript:;" data-cmd="qz_reply" data-version="6.4" data-action="http://photo.qq.com/cgi-bin/common/cgi_add_icreply" data-param="oweruin=610018923&amp;albumid=V14BzO970zmaxV&amp;lloc=NDR0ayZcJCy3w1ltACw6PAEAAAAAAAA!&amp;sloc=NDR0ayZcJCy3w1ltACw6PAEAAAAAAAA!&amp;refer=qzone&amp;reqfrom=801&amp;zz=0&amp;cmtid=34451106&amp;batchId=1505998675248&amp;cmtType=4&amp;picname=2017-09-21" data-charset="GB" data-tuin="" data-config="{config:'1|1|0|0|0|0|0',targetUserInfo:{id:610018923,nick:'',who:1,auto:1}}"><b class="hide-clip">回复</b></a></div></div></div></div></li><li class="comments-item bor3" data-type="replyroot" data-tid="1506003353" data-uin="2507820334" data-nick="思＆" data-who="1"><div class="comments-item-bd"><div class="single-reply"><div class="ui-avatar"><a href="http://user.qzone.qq.com/2507820334" target="_blank"><img class="q_namecard" link="nameCard_2507820334 des_2507820334" alt="思＆" src="http://qlogo3.store.qq.com/qzone/2507820334/2507820334/30"></a></div><div class="comments-content"><a class="nickname name c_tx q_namecard" link="nameCard_2507820334" target="_blank" href="http://user.qzone.qq.com/2507820334">思＆</a>&nbsp;回复<a class="nickname name c_tx q_namecard" link="nameCard_610018923" target="_blank" href="http://user.qzone.qq.com/610018923">梦里有鬼桃</a>&nbsp; :  就是耶，看到了，你有反应没有得<div class="comments-op"><span class=" ui-mr10 state"> 22:15</span><a class="act-reply none" href="javascript:;" data-cmd="qz_reply" data-version="6.4" data-action="http://photo.qq.com/cgi-bin/common/cgi_add_icreply" data-param="oweruin=610018923&amp;albumid=V14BzO970zmaxV&amp;lloc=NDR0ayZcJCy3w1ltACw6PAEAAAAAAAA!&amp;sloc=NDR0ayZcJCy3w1ltACw6PAEAAAAAAAA!&amp;refer=qzone&amp;reqfrom=801&amp;zz=0&amp;cmtid=34451106&amp;batchId=1505998675248&amp;cmtType=4&amp;picname=2017-09-21" data-charset="GB" data-tuin="" data-config="{config:'1|1|0|0|0|0|0',targetUserInfo:{id:2507820334,nick:'',who:1,auto:1}}"><b class="hide-clip">回复</b></a></div></div></div></div></li><li class="comments-item bor3" data-type="replyroot" data-tid="1506003461" data-uin="610018923" data-nick="梦里有鬼桃" data-who="1"><div class="comments-item-bd"><div class="single-reply"><div class="ui-avatar"><a href="http://user.qzone.qq.com/610018923" target="_blank"><img class="q_namecard" link="nameCard_610018923 des_610018923" alt="梦里有鬼桃" src="http://qlogo4.store.qq.com/qzone/610018923/610018923/30?1400138040"></a></div><div class="comments-content"><a class="nickname name c_tx q_namecard" link="nameCard_610018923" target="_blank" href="http://user.qzone.qq.com/610018923">梦里有鬼桃</a>&nbsp;回复<a class="nickname name c_tx q_namecard" link="nameCard_2507820334" target="_blank" href="http://user.qzone.qq.com/2507820334">思＆</a>&nbsp; : 肚子偶尔痛，没有规律，不晓得那天会生<img src="http://qzonestyle.gtimg.cn/qzone/em/e120.png" title=""><img src="http://qzonestyle.gtimg.cn/qzone/em/e120.png" title=""><div class="comments-op"><span class=" ui-mr10 state"> 22:17</span><a class="act-reply none" href="javascript:;" data-cmd="qz_reply" data-version="6.4" data-action="http://photo.qq.com/cgi-bin/common/cgi_add_icreply" data-param="oweruin=610018923&amp;albumid=V14BzO970zmaxV&amp;lloc=NDR0ayZcJCy3w1ltACw6PAEAAAAAAAA!&amp;sloc=NDR0ayZcJCy3w1ltACw6PAEAAAAAAAA!&amp;refer=qzone&amp;reqfrom=801&amp;zz=0&amp;cmtid=34451106&amp;batchId=1505998675248&amp;cmtType=4&amp;picname=2017-09-21" data-charset="GB" data-tuin="" data-config="{config:'1|1|0|0|0|0|0',targetUserInfo:{id:610018923,nick:'',who:1,auto:1}}"><b class="hide-clip">回复</b></a></div></div></div></div></li><li class="comments-item bor3" data-type="replyroot" data-tid="1506003488" data-uin="2507820334" data-nick="思＆" data-who="1"><div class="comments-item-bd"><div class="single-reply"><div class="ui-avatar"><a href="http://user.qzone.qq.com/2507820334" target="_blank"><img class="q_namecard" link="nameCard_2507820334 des_2507820334" alt="思＆" src="http://qlogo3.store.qq.com/qzone/2507820334/2507820334/30"></a></div><div class="comments-content"><a class="nickname name c_tx q_namecard" link="nameCard_2507820334" target="_blank" href="http://user.qzone.qq.com/2507820334">思＆</a>&nbsp;回复<a class="nickname name c_tx q_namecard" link="nameCard_610018923" target="_blank" href="http://user.qzone.qq.com/610018923">梦里有鬼桃</a>&nbsp; :  好吧，自己注意点<div class="comments-op"><span class=" ui-mr10 state"> 22:18</span><a class="act-reply none" href="javascript:;" data-cmd="qz_reply" data-version="6.4" data-action="http://photo.qq.com/cgi-bin/common/cgi_add_icreply" data-param="oweruin=610018923&amp;albumid=V14BzO970zmaxV&amp;lloc=NDR0ayZcJCy3w1ltACw6PAEAAAAAAAA!&amp;sloc=NDR0ayZcJCy3w1ltACw6PAEAAAAAAAA!&amp;refer=qzone&amp;reqfrom=801&amp;zz=0&amp;cmtid=34451106&amp;batchId=1505998675248&amp;cmtType=4&amp;picname=2017-09-21" data-charset="GB" data-tuin="" data-config="{config:'1|1|0|0|0|0|0',targetUserInfo:{id:2507820334,nick:'',who:1,auto:1}}"><b class="hide-clip">回复</b></a></div></div></div></div></li></ul></div></div></li></ul></div><div class="mod-commnets-poster feedClickCmd comment_default_inputentry" data-cmd="qz_reply" data-version="6" data-action="http://photo.qq.com/cgi-bin/common/cgi_add_piccomment" data-param="uin=610018923&amp;albumid=V14BzO970zmaxV&amp;forumindex=0&amp;lloc=NDR0ayZcJCy3w1ltACw6PAEAAAAAAAA!&amp;sloc=NDR0ayZcJCy3w1ltACw6PAEAAAAAAAA!&amp;privacy=1&amp;refer=qzone&amp;reqfrom=801&amp;batchId=1505998675248&amp;cmtType=4&amp;zz=0" data-charset="GB" data-maxlength="" data-tuin="610018923" data-config="1|1|0|0|0|0|0" data-tid=""><div class="comments-poster-bd comments-poster-default"><div class="comments-box" data-clicklog="comment"><div class="textinput textinput-default bor2" contenteditable="true" alt="replybtn" placeholder="评论"><a class="c_tx3" href="javascript:void(0);" alt="replybtn">评论</a></div><div class="mod-insert-img"><a href="javascript:;" data-cmd="qz_quick_upload_img" class="btn-insert-img bg"><i class="icon-camera-16"></i></a></div></div></div></div></div>    </div></li><li class="feed_holder_top" style="height:0"></li><li class="feed_page_container"><ul data-page="0"><li class="f-single f-s-s" id="fct_1172992530_311_0_1505997686_0_1" read="1"><div class="f-single-head f-aside"><div class="f-adorn-top"></div>    <div class="user-pto"><a href="http://user.qzone.qq.com/1172992530" target="_blank" class="user-avatar q_namecard f-s-a" link="nameCard_1172992530" data-clicklog="avatar"><img src="https://qlogo3.store.qq.com/qzone/1172992530/1172992530/50?1505278502"></a></div><div class="user-op"><a href="javascript:;" class="arrow-down" data-cmd="qz_opr_more" data-moreoperate="1"><i class="fui-icon icon-arrow-down"></i></a>        </div><div class="user-info"><div class="f-nick"><a target="_blank" href="http://user.qzone.qq.com/1172992530" data-clicklog="nick" class="f-name q_namecard " link="nameCard_1172992530">T181-朱雪梅</a>  <a class="user-medal" hotclickpath="isd.qzonemall.year.feeds" hotdomain="mall.qzone.qq.com" href="http://pay.qq.com/ipay/index.shtml?n=3&amp;c=xxjzghh,xxjzgw&amp;aid=feed_guajian&amp;ch=qdqb,kj" target="_blank" title="点击查看黄钻特权详情"><span class="qz-f-vip-l-y qz-f-vip-l-gray-y-gray-6"></span></a></div><div class="info-detail"><span class=" ui-mr8 state"> 20:41</span><a href="javascript:;" data-cmd="qz_sign" class="f-sign-show state" title="我也要设置"></a></div></div></div><div class="f-single-content f-wrap"><div class="f-item f-s-i" id="feed_1172992530_311_0_1505997686_0_1" data-feedsflag="" data-iswupfeed="1" data-key="1272ea4576b3c35965850e00" data-specialtype="" data-extend-info="0_0_0_0_0_0_0|18d58de0000cc201|1018000056ae3805" data-functype="func_friend_feed" data-hasfollowed="1"><div class="f-info">关掉了手机管他谁是谁！</div><div class="qz_summary wupfeed" id="hex_1172992530_311_0_1505997686_0_1"><i class="none" name="feed_data" data-bitmap="18d58de0000cc201" data-yybitmap="1018000056ae3805" data-vipstarbitmap="000000004a001000" data-fkey="1272ea4576b3c35965850e00" data-tid="1272ea4576b3c35965850e00" data-uin="1172992530" data-origfkey="" data-origtid="1272ea4576b3c35965850e00" data-origuin="1172992530" data-subid="" data-totweet="" data-issignin="" data-source="" data-retweetcount="0" data-stat="OrykmfNMKB7vg4zLvYDkd3Te3gBZJ0AnrRA2MqEMMyEqotL!59BpxPIRqoYkz9dU3gvh1NzY8amYHxmPwPf2SWB42W7v5HFTkF9st28!47FJlxpzOyNqiyvDVw2pTdYyPNC7J3Y25ffADjhmVzKfoXR!xXo9wdavvmmHBl4/hAkqxZdXpaXN0rcIfbQQeG8x3urF3UhC0LYwTGGthWENnBpIkX30!RXT_" data-topicid="1172992530_1272ea4576b3c35965850e00__1" data-feedstype="100" data-abstime="1505997686" data-iswupfeed="1" data-platformid="54" data-accessright="1"></i><div class="f-ct "><div class="f-ct-txtimg fui-txtimg  "><div class="txt-box ">                                    </div><div class="img-box "><a class="img-item  " data-cmd="qz_popup" href="https://user.qzone.qq.com/1172992530/311/" data-topicid="1172992530_1272ea4576b3c35965850e00__1" data-pickey="1272ea4576b3c35965850e00,1172992530,V10dFNCz3c0tTJ,NDR0EnLqRXezw1nOd9kG8gAAAAAAAAA!" data-clicklog="pic" data-originurl="||" hotclickpath="0_appid_311_v8.pic_count_1.pic_0" hotdomain="icv6act.qzone.qq.com" data-version="2" data-param="1272ea4576b3c35965850e00|1172992530|0" data-src="/qzone/photo/zone/icenter_popup.html" data-width="640" data-height="1136" data-type="popup" data-title="" data-config=""><img src="http://a3.qpic.cn/psb?/V10dFNCz3c0tTJ/eBBzE2zx6Db*UuXc*9aj3u5csDhybKXIYbxhCTplZbM!/c/dPIAAAAAAAAA&amp;ek=1&amp;kp=1&amp;pt=0&amp;bo=gAJwBAAAAAAAANA!&amp;t=5&amp;vuin=294477044&amp;tm=1506006000&amp;sce=60-2-2&amp;rf=0-0" style="height:420px;"></a></div></div></div></div></div></div><div class="f-single-foot"><div class="f-op-detail f-detail content-line"><p class="op-list"><a class="item qz_retweet_btn " href="javascript:;" data-cmd="qz_popup" data-version="4" data-isnewtype="1" data-type="ForwardingBox" data-src="/qzone/app/controls/forwardingBox/forwardingBoxFacade.js" data-clicklog="retweet"><i class="fui-icon icon-op-share"></i></a><span class="item-line"></span>&nbsp;<a href="javascript:;" data-version="6.3" data-cmd="qz_reply" data-link="1" data-clicklog="comment" class=" qz_btn_reply item "><i class="fui-icon icon-op-comment"></i></a>&nbsp;<span class="item-line"></span><a class="item qz_like_btn_v3 " data-islike="0" data-likecnt="2" data-showcount="2" data-unikey="http://user.qzone.qq.com/1172992530/mood/1272ea4576b3c35965850e00" data-curkey="http://user.qzone.qq.com/1172992530/mood/1272ea4576b3c35965850e00" data-clicklog="like" href="javascript:;"><i class="fui-icon icon-op-praise"></i></a></p><a href="javascript:;" class="state qz_feed_plugin" data-role="Visitor" data-config="311|1272ea4576b3c35965850e00|1172992530" data-clicklog="visitor">浏览24次</a></div>        <div class="f-ang-t"></div><div class="f-like-list f-like _likeInfo" likeinfo="2"><div class="icon-btn"><a href="javascript:;" data-islike="0" data-likecnt="2" data-showcount="2" data-unikey="http://user.qzone.qq.com/1172992530/mood/1272ea4576b3c35965850e00" data-curkey="http://user.qzone.qq.com/1172992530/mood/1272ea4576b3c35965850e00" data-clicklog="like" class="praise qz_like_prase"><i class="fui-icon icon-list-praise"></i></a><div class="bubble" style="display:none;"><div class="bd">+1</div><b class="arrow arrow-down"></b></div></div><div class="user-list"><a href="http://user.qzone.qq.com/424500179" class="item q_namecard" target="_blank" link="nameCard_424500179 des_424500179">是梦，终会醒</a>、<a href="http://user.qzone.qq.com/2567822529" class="item q_namecard" target="_blank" link="nameCard_2567822529 des_2567822529">1y19</a>共<span class="f-like-cnt">2</span>人觉得很赞</div></div><div class="mod-comments" style="padding:0"><div class="mod-commnets-poster feedClickCmd comment_default_inputentry" data-cmd="qz_reply" data-version="6" data-action="http://taotao.qq.com/cgi-bin/emotion_cgi_re_feeds" data-param="t1_source=&amp;t1_uin=1172992530&amp;t1_tid=1272ea4576b3c35965850e00&amp;signin=0&amp;sceneid=100" data-charset="utf-8" data-maxlength="" data-tuin="1172992530" data-config="1|1|1|1,b52,with_fwd,同时转发;0|1,taotaoact.qzone.qq.com,@InputReply|1,taotaoact.qzone.qq.com,@ClickReply|1,taotaoact.qzone.qq.com,commentPresentClick" data-tid=""><div class="comments-poster-bd comments-poster-default"><div class="comments-box" data-clicklog="comment"><div class="textinput textinput-default bor2" contenteditable="true" alt="replybtn" placeholder="评论"><a class="c_tx3" href="javascript:void(0);" alt="replybtn">评论</a></div><div class="mod-insert-img"><a href="javascript:;" data-cmd="qz_quick_upload_img" class="btn-insert-img bg"><i class="icon-camera-16"></i></a></div></div></div></div></div>    </div></li><li class="f-single f-s-s" id="fct_610018923_4_11_1505996388_0_1"><div class="f-single-head f-aside"><div class="f-adorn-top"></div>    <div class="user-pto"><a href="http://user.qzone.qq.com/610018923" target="_blank" class="user-avatar q_namecard f-s-a" link="nameCard_610018923" data-clicklog="avatar"><img src="https://qlogo4.store.qq.com/qzone/610018923/610018923/50?1400138040"></a></div><div class="user-op"><a href="javascript:;" class="arrow-down" data-cmd="qz_opr_more" data-moreoperate="1"><i class="fui-icon icon-arrow-down"></i></a>        </div><div class="user-info"><div class="f-nick"><a target="_blank" href="http://user.qzone.qq.com/610018923" data-clicklog="nick" class="f-name q_namecard " link="nameCard_610018923">梦里有鬼桃</a>  </div><div class="info-detail"><span class=" ui-mr8 state"> 20:19</span><a href="javascript:;" data-cmd="qz_sign" class="f-sign-show state" title="我也要设置"></a></div></div></div><div class="f-single-content f-wrap"><div class="f-item f-s-i" id="feed_610018923_4_11_1505996388_0_1" data-feedsflag="" data-iswupfeed="1" data-key="NDR0ayZcJGSuw1lvmFAihQEAAAAAAAA" data-specialtype="" data-extend-info="0_0_0_0_0_0_0|981095a00102e201|004a000000000005" data-functype="func_friend_feed" data-hasfollowed="1"><div class="qz_summary wupfeed" id="hex_610018923_4_11_1505996388_0_1"><i class="none" name="feed_data" data-bitmap="981095a00102e201" data-yybitmap="004a000000000005" data-vipstarbitmap="0000000040000000" data-fkey="NDR0ayZcJGSuw1lvmFAihQEAAAAAAAA" data-tid="V14BzO971Yb2cu" data-uin="610018923" data-origfkey="" data-origtid="NDR0ayZcJGSuw1lvmFAihQEAAAAAAAA!" data-origuin="610018923" data-subid="NDR0ayZcJGSuw1lvmFAihQEAAAAAAAA!" data-totweet="" data-issignin="" data-source="" data-retweetcount="0" data-stat="4VxYpxNSXHPecIt409yxmr8z6D5wd/6qe4rQAv6LRvgbPKlkzt4JnlyCrl5urHJfqCxcWoXHGhnvZjuf7SW/h2bTWguEiv0W6bpm8OzJTWwp9n2i8PARCGlXMA5RPPdQ0lKgpIlhEnYNytP9jzC5jtYdrY7jUyi9lcDB5!vi8wyNoakRtfEN/dqnWd5LwlgvIhz4enB/nNwbPKlkzt4JnlyCrl5urHJfqCxcWoXHGhnvZjuf7SW/hzMwJG84NQMCIPborQjWSHYKsRSV1oWB6A_" data-topicid="V14BzO971Yb2cu_NDR0ayZcJGSuw1lvmFAihQEAAAAAAAA!_0_0" data-feedstype="100" data-abstime="1505996388" data-iswupfeed="1" data-platformid="52" data-accessright="1"></i><div class="f-ct"><div class="f-info"></div><div class="fui-txtimg "><div class="img-box f-video-wrap play" url3="http://vqzone.gtimg.com/1097_27c1ab4a42e44c06a90e94c11b1adbbb.f20.mp4?ptype=http&amp;vkey=2DA002E9F78A628227547C3FA6FEA868B0D81AD87D20E97CAC42EC0F7B5590EB509640CCD76ADB30795AFA638A5FB6BB2E1F8FFB2DA4CEE3&amp;sdtfrom=v1000&amp;owner=610018923" style="width: 560px; height: 560px; min-width: 560px; min-height: 560px;"><a class="img-item  " data-cmd="qz_popup" href="https://h5.qzone.qq.com/init=photo.v7/common/viewer2/index&amp;ownerUin=610018923&amp;appid=4&amp;topicId=V14BzO971Yb2cu_NDR0ayZcJGSuw1lvmFAihQEAAAAAAAA!_0_0&amp;pre=http%3A%2F%2Fb389.photo.store.qq.com%2Fpsb%3F%2FV14BzO971Yb2cu%2FYTQJ70RU5PuarrxjdCqkQUYVFt9sUHD8FuDTmVbrXKQ!%2Fc%2FdIUBAAAAAAAA%26bo%3DQASAB0AEgAcRFyA!&amp;useqzfl=1&amp;useinterface=1&amp;noCloseBtn=0&amp;inqq=1" data-v_itemid="NDR0ayZcJGSuw1lvmFAihQEAAAAAAAA!" data-v_type="1" data-v_picinfo_url="https://qqadapt.qpic.cn/wecam_pic/0/1097_27c1ab4a42e44c06a90e94c11b1adbbb_1/200" data-v_picinfo_width="200" data-v_picinfo_height="353" data-v_vidiourl="http://vqzone.gtimg.com/1097_27c1ab4a42e44c06a90e94c11b1adbbb.f20.mp4?ptype=http&amp;vkey=2DA002E9F78A628227547C3FA6FEA868B0D81AD87D20E97CAC42EC0F7B5590EB509640CCD76ADB30795AFA638A5FB6BB2E1F8FFB2DA4CEE3&amp;sdtfrom=v1000&amp;owner=610018923" data-v_vidioswfurl="http://vqzone.gtimg.com/1097_27c1ab4a42e44c06a90e94c11b1adbbb.f20.mp4?ptype=http&amp;vkey=2DA002E9F78A628227547C3FA6FEA868B0D81AD87D20E97CAC42EC0F7B5590EB509640CCD76ADB30795AFA638A5FB6BB2E1F8FFB2DA4CEE3&amp;sdtfrom=v1000&amp;owner=610018923" data-v_h265="" data-v_source_website="" data-v_writefrom="" data-videotype="mood" data-clicklog="video" hotclickpath="" hotdomain="" data-version="3" data-param="3&amp;videosrc=http%3A%2F%2Fvqzone.gtimg.com%2F1097_27c1ab4a42e44c06a90e94c11b1adbbb.f20.mp4%3Fptype%3Dhttp%26vkey%3D2DA002E9F78A628227547C3FA6FEA868B0D81AD87D20E97CAC42EC0F7B5590EB509640CCD76ADB30795AFA638A5FB6BB2E1F8FFB2DA4CEE3%26sdtfrom%3Dv1000%26owner%3D610018923&amp;type=1&amp;org_vidio_url=http%3A%2F%2Fvqzone.gtimg.com%2F1097_27c1ab4a42e44c06a90e94c11b1adbbb.f20.mp4%3Fptype%3Dhttp%26vkey%3D2DA002E9F78A628227547C3FA6FEA868B0D81AD87D20E97CAC42EC0F7B5590EB509640CCD76ADB30795AFA638A5FB6BB2E1F8FFB2DA4CEE3%26sdtfrom%3Dv1000%26owner%3D610018923" data-src="/qzone/app/mood/richinfo_view.html" data-width="512" data-height="512" data-type="popup" data-title="" data-config="" data-vfeed-id="vfeed_610018923_4_V14BzO971Yb2cu_NDR0ayZcJGSuw1lvmFAihQEAAAAAAAA!" data-newplayer="1" data-newplayer-id="vfeed_610018923_4_V14BzO971Yb2cu_NDR0ayZcJGSuw1lvmFAihQEAAAAAAAA!" style="background-color: rgb(0, 0, 0); width: 560px; height: 560px;"><div class="video-img j-videofeed-imgctn" style="width: 560px; height: 560px;"><img src="http://qqadapt.qpic.cn/wecam_pic/0/1097_27c1ab4a42e44c06a90e94c11b1adbbb_1/200" data-oriwidth="200" data-oriheight="353" style="width: 317.28px; height: 560px; margin: 0px 121.36px;"><i class="ui-icon icon-media-play j-videofeed-icon-play"></i><span class="video-loading j-videofeed-icon-loading" style="display:none;"><i class="inner"></i></span></div><div class="j-videofeed-flashctn" data-need_hide_when_inited="1" style="position: absolute; left: 0px; top: 0px; width: 560px; height: 560px; overflow: hidden;"><object width="100%" height="100%" id="vfeed_610018923_4_V14BzO971Yb2cu_NDR0ayZcJGSuw1lvmFAihQEAAAAAAAA.__flashobj" name="vfeed_610018923_4_V14BzO971Yb2cu_NDR0ayZcJGSuw1lvmFAihQEAAAAAAAA.__flashobj" data="//qzs.qq.com/qzone/client/photo/swf/MPlayer/MicroVideoPlayerEx.swf" type="application/x-shockwave-flash"><param name="allowFullScreen" value="true"><param name="allowNetworking" value="all"><param name="allowScriptAccess" value="always"><param name="wmode" value="opaque"><param name="flashvars" value="log=0&amp;vurl=http%3A%2F%2Fvqzone.gtimg.com%2F1097_27c1ab4a42e44c06a90e94c11b1adbbb.f20.mp4%3Fptype%3Dhttp%26vkey%3D2DA002E9F78A628227547C3FA6FEA868B0D81AD87D20E97CAC42EC0F7B5590EB509640CCD76ADB30795AFA638A5FB6BB2E1F8FFB2DA4CEE3%26sdtfrom%3Dv1000%26owner%3D610018923&amp;vurlBak=&amp;noloading=1&amp;auto=1&amp;mute=1&amp;hasPopup=1&amp;hasViewer=1&amp;onFlashInited=window.__qzvCbsMap%5B%22vfeed_610018923_4_V14BzO971Yb2cu_NDR0ayZcJGSuw1lvmFAihQEAAAAAAAA.__flashobj%22%5D.onFlashInited&amp;onChangePlayer=window.__qzvCbsMap%5B%22vfeed_610018923_4_V14BzO971Yb2cu_NDR0ayZcJGSuw1lvmFAihQEAAAAAAAA.__flashobj%22%5D.onChangePlayer&amp;onReportLog=window.__qzvCbsMap%5B%22vfeed_610018923_4_V14BzO971Yb2cu_NDR0ayZcJGSuw1lvmFAihQEAAAAAAAA.__flashobj%22%5D.onReportLog&amp;onError=window.__qzvCbsMap%5B%22vfeed_610018923_4_V14BzO971Yb2cu_NDR0ayZcJGSuw1lvmFAihQEAAAAAAAA.__flashobj%22%5D.onError&amp;onMetaData=window.__qzvCbsMap%5B%22vfeed_610018923_4_V14BzO971Yb2cu_NDR0ayZcJGSuw1lvmFAihQEAAAAAAAA.__flashobj%22%5D.onMetaData&amp;onChangeState=window.__qzvCbsMap%5B%22vfeed_610018923_4_V14BzO971Yb2cu_NDR0ayZcJGSuw1lvmFAihQEAAAAAAAA.__flashobj%22%5D.onChangeState&amp;onPlayStart=window.__qzvCbsMap%5B%22vfeed_610018923_4_V14BzO971Yb2cu_NDR0ayZcJGSuw1lvmFAihQEAAAAAAAA.__flashobj%22%5D.onPlayStart&amp;onPlayStop=window.__qzvCbsMap%5B%22vfeed_610018923_4_V14BzO971Yb2cu_NDR0ayZcJGSuw1lvmFAihQEAAAAAAAA.__flashobj%22%5D.onPlayStop&amp;onSlideStart=window.__qzvCbsMap%5B%22vfeed_610018923_4_V14BzO971Yb2cu_NDR0ayZcJGSuw1lvmFAihQEAAAAAAAA.__flashobj%22%5D.onSlideStart&amp;onSlideStop=window.__qzvCbsMap%5B%22vfeed_610018923_4_V14BzO971Yb2cu_NDR0ayZcJGSuw1lvmFAihQEAAAAAAAA.__flashobj%22%5D.onSlideStop&amp;onSkinAction=window.__qzvCbsMap%5B%22vfeed_610018923_4_V14BzO971Yb2cu_NDR0ayZcJGSuw1lvmFAihQEAAAAAAAA.__flashobj%22%5D.onSkinAction&amp;onChangeMute=window.__qzvCbsMap%5B%22vfeed_610018923_4_V14BzO971Yb2cu_NDR0ayZcJGSuw1lvmFAihQEAAAAAAAA.__flashobj%22%5D.onChangeMute&amp;onChangeFull=window.__qzvCbsMap%5B%22vfeed_610018923_4_V14BzO971Yb2cu_NDR0ayZcJGSuw1lvmFAihQEAAAAAAAA.__flashobj%22%5D.onChangeFull&amp;onOpenPopup=window.__qzvCbsMap%5B%22vfeed_610018923_4_V14BzO971Yb2cu_NDR0ayZcJGSuw1lvmFAihQEAAAAAAAA.__flashobj%22%5D.onOpenPopup&amp;onOpenViewer=window.__qzvCbsMap%5B%22vfeed_610018923_4_V14BzO971Yb2cu_NDR0ayZcJGSuw1lvmFAihQEAAAAAAAA.__flashobj%22%5D.onOpenViewer&amp;onClickMask=window.__qzvCbsMap%5B%22vfeed_610018923_4_V14BzO971Yb2cu_NDR0ayZcJGSuw1lvmFAihQEAAAAAAAA.__flashobj%22%5D.onClickMask"></object></div></a></div></div><div class="f-reprint">                        <p class="item state">        <i class="fui-icon icon-print-album"></i>        <a href="http://user.qzone.qq.com/610018923/photo/V14BzO971Yb2cu" target="_blank" style="cursor:pointer;">                                上传1个视频到《手机相册》        </a>        </p>            </div>    </div></div></div></div><div class="f-single-foot"><div class="f-op-detail f-detail content-line"><p class="op-list"><a class="item qz_retweet_btn " href="javascript:;" data-cmd="qz_popup" data-version="4" data-isnewtype="1" data-type="ForwardingBox" data-src="/qzone/app/controls/forwardingBox/forwardingBoxFacade.js" data-clicklog="retweet"><i class="fui-icon icon-op-share"></i></a><span class="item-line"></span>&nbsp;<a href="javascript:;" data-version="6.3" data-cmd="qz_reply" data-link="1" data-clicklog="comment" class=" qz_btn_reply item "><i class="fui-icon icon-op-comment"></i></a>&nbsp;<span class="item-line"></span><a class="item qz_like_btn_v3 " data-islike="0" data-likecnt="0" data-showcount="0" data-unikey="http://user.qzone.qq.com/610018923/photo/V14BzO971Yb2cu/NDR0ayZcJGSuw1lvmFAihQEAAAAAAAA!" data-curkey="http://user.qzone.qq.com/610018923/photo/V14BzO971Yb2cu/NDR0ayZcJGSuw1lvmFAihQEAAAAAAAA!" data-clicklog="like" href="javascript:;"><i class="fui-icon icon-op-praise"></i></a></p><a href="javascript:;" class="state qz_feed_plugin" data-role="Visitor" data-config="4|9;V14BzO971Yb2cu%7CNDR0ayZcJGSuw1lvmFAihQEAAAAAAAA!|610018923" data-clicklog="visitor">浏览23次</a></div>        <div class="f-ang-t"></div><div class="mod-comments" style="padding:0"><div class="comments-list "><ul><li class="comments-item bor3" data-type="commentroot" data-tid="23143585" data-uin="402636981" data-nick="门吉" data-who="1"><div class="comments-item-bd"><div class="single-reply"><div class="ui-avatar"><a href="http://user.qzone.qq.com/402636981" target="_blank"><img class="q_namecard" link="nameCard_402636981 des_402636981" alt="门吉" src="http://qlogo2.store.qq.com/qzone/402636981/402636981/30"></a></div><div class="comments-content"><a class="nickname name c_tx q_namecard" link="nameCard_402636981" target="_blank" href="http://user.qzone.qq.com/402636981">门吉</a>&nbsp; : 这是啥子东东？<div class="comments-op"><span class=" ui-mr10 state"> 20:44</span><a class="act-reply none" href="javascript:;" data-cmd="qz_reply" data-version="6.4" data-action="http://photo.qq.com/cgi-bin/common/cgi_add_icreply" data-param="oweruin=610018923&amp;albumid=V14BzO971Yb2cu&amp;lloc=NDR0ayZcJGSuw1lvmFAihQEAAAAAAAA!&amp;sloc=NDR0ayZcJGSuw1lvmFAihQEAAAAAAAA!&amp;refer=qzone&amp;reqfrom=601&amp;zz=0&amp;cmtid=23143585&amp;picname=" data-charset="GB" data-tuin="" data-config="{config:'1|1|0|0|0|0|0',targetUserInfo:{id:402636981,nick:'闂ㄥ悏',who:1,auto:1}}"><b class="hide-clip">回复</b></a></div></div></div><div class="comments-list mod-comments-sub"><ul><li class="comments-item bor3" data-type="replyroot" data-tid="1506003320" data-uin="610018923" data-nick="梦里有鬼桃" data-who="1"><div class="comments-item-bd"><div class="single-reply"><div class="ui-avatar"><a href="http://user.qzone.qq.com/610018923" target="_blank"><img class="q_namecard" link="nameCard_610018923 des_610018923" alt="梦里有鬼桃" src="http://qlogo4.store.qq.com/qzone/610018923/610018923/30?1400138040"></a></div><div class="comments-content"><a class="nickname name c_tx q_namecard" link="nameCard_610018923" target="_blank" href="http://user.qzone.qq.com/610018923">梦里有鬼桃</a>&nbsp;回复<a class="nickname name c_tx q_namecard" link="nameCard_402636981" target="_blank" href="http://user.qzone.qq.com/402636981">门吉</a>&nbsp; : 吹糖人<div class="comments-op"><span class=" ui-mr10 state"> 22:15</span><a class="act-reply none" href="javascript:;" data-cmd="qz_reply" data-version="6.4" data-action="http://photo.qq.com/cgi-bin/common/cgi_add_icreply" data-param="oweruin=610018923&amp;albumid=V14BzO971Yb2cu&amp;lloc=NDR0ayZcJGSuw1lvmFAihQEAAAAAAAA!&amp;sloc=NDR0ayZcJGSuw1lvmFAihQEAAAAAAAA!&amp;refer=qzone&amp;reqfrom=601&amp;zz=0&amp;cmtid=23143585&amp;picname=" data-charset="GB" data-tuin="" data-config="{config:'1|1|0|0|0|0|0',targetUserInfo:{id:610018923,nick:'',who:1,auto:1}}"><b class="hide-clip">回复</b></a></div></div></div></div></li></ul></div></div></li></ul></div><div class="mod-commnets-poster feedClickCmd comment_default_inputentry" data-cmd="qz_reply" data-version="6" data-action="http://photo.qq.com/cgi-bin/common/cgi_add_piccomment" data-param="uin=610018923&amp;albumid=V14BzO971Yb2cu&amp;forumindex=0&amp;lloc=NDR0ayZcJGSuw1lvmFAihQEAAAAAAAA!&amp;sloc=NDR0ayZcJGSuw1lvmFAihQEAAAAAAAA!&amp;privacy=1&amp;refer=qzone&amp;reqfrom=601&amp;cmtType=0&amp;zz=0&amp;picname=" data-charset="GB" data-maxlength="" data-tuin="610018923" data-config="1|1|0|0|0|0|0" data-tid=""><div class="comments-poster-bd comments-poster-default"><div class="comments-box" data-clicklog="comment"><div class="textinput textinput-default bor2" contenteditable="true" alt="replybtn" placeholder="评论"><a class="c_tx3" href="javascript:void(0);" alt="replybtn">评论</a></div><div class="mod-insert-img"><a href="javascript:;" data-cmd="qz_quick_upload_img" class="btn-insert-img bg"><i class="icon-camera-16"></i></a></div></div></div></div></div>    </div></li><li class="f-single f-s-s" id="fct_994199947_311_0_1505995597_0_1"><div class="f-single-head f-aside"><div class="f-adorn-top"></div>    <div class="user-pto"><a href="http://user.qzone.qq.com/994199947" target="_blank" class="user-avatar q_namecard f-s-a" link="nameCard_994199947" data-clicklog="avatar"><img src="https://qlogo4.store.qq.com/qzone/994199947/994199947/50?1356758777"></a></div><div class="user-op"><a href="javascript:;" class="arrow-down" data-cmd="qz_opr_more" data-moreoperate="1"><i class="fui-icon icon-arrow-down"></i></a>        </div><div class="user-info"><div class="f-nick"><a target="_blank" href="http://user.qzone.qq.com/994199947" data-clicklog="nick" class="f-name q_namecard " link="nameCard_994199947">易小军</a>  </div><div class="info-detail"><span class=" ui-mr8 state"> 20:06</span><a href="javascript:;" data-cmd="qz_sign" class="f-sign-show state" title="我也要设置"></a></div></div></div><div class="f-single-content f-wrap"><div class="f-item f-s-i" id="feed_994199947_311_0_1505995597_0_1" data-feedsflag="" data-iswupfeed="1" data-key="8b49423b4cabc359f53f0f00" data-specialtype="" data-extend-info="0_0_0_0_0_0_0|18d195a00100e109|0000000000000001" data-functype="func_friend_feed" data-hasfollowed="1"><div class="f-info">2017·9.21乐道教育私人健身房，肌肉男的一天<img src="http://qzonestyle.gtimg.cn/qzone/em/e113.png" title=""></div><div class="qz_summary wupfeed" id="hex_994199947_311_0_1505995597_0_1"><i class="none" name="feed_data" data-bitmap="18d195a00100e109" data-yybitmap="0000000000000001" data-vipstarbitmap="0000000000002000" data-fkey="8b49423b4cabc359f53f0f00" data-tid="8b49423b4cabc359f53f0f00" data-uin="994199947" data-origfkey="" data-origtid="8b49423b4cabc359f53f0f00" data-origuin="994199947" data-subid="" data-totweet="" data-issignin="" data-source="" data-retweetcount="0" data-stat="75p2HcTdgjya1goCq1GSC3I8CRqpElNGvsFU8ZZlCVzi1t4Pzu8snyAych1rdY/ef5djyQOjDhQwpDhjzBFTJiIU3T83MGnOwr!TVDZz4YEW2xz/hdsLWoK8KHB1CrL!Nv1yjzGmQKEse6hFY6t2GKmrPqqAN3dFH9XL!CTyRHgTKmtoRHU78Xuqwh!9fo8i_" data-topicid="994199947_8b49423b4cabc359f53f0f00__1" data-feedstype="100" data-abstime="1505995597" data-iswupfeed="1" data-platformid="52" data-accessright="1"></i><div class="f-ct "><div class="f-ct-txtimg fui-txtimg f-ct-video  "><div class="txt-box ">                                    </div><div class="img-box f-video-wrap play" url3="http://vwecam.gtimg.com/1006_081dff72e4004afc99d55064aac8ccbb.f0.mp4?ptype=http&amp;vkey=C6F0EA8CF63D1C68176C3F7314EF152ADD3E27A48A776895DB8B107B34E3588DAD8BDF26EDFC7E3C9357EA4DB0CA70782268E6D614D75DCE&amp;sdtfrom=v1000&amp;owner=994199947" style="width: 560px; height: 560px; min-width: 560px; min-height: 560px;"><a class="img-item  " data-cmd="qz_popup" href="https://user.qzone.qq.com/994199947/311/1006_081dff72e4004afc99d55064aac8ccbb" data-v_itemid="1006_081dff72e4004afc99d55064aac8ccbb" data-v_type="1" data-v_picinfo_url="https://qqadapt.qpic.cn/wecam_pic/0/1006_081dff72e4004afc99d55064aac8ccbb_1/200" data-v_picinfo_width="200" data-v_picinfo_height="356" data-v_vidiourl="" data-v_vidioswfurl="http://vwecam.gtimg.com/1006_081dff72e4004afc99d55064aac8ccbb.f0.mp4?ptype=http&amp;vkey=C6F0EA8CF63D1C68176C3F7314EF152ADD3E27A48A776895DB8B107B34E3588DAD8BDF26EDFC7E3C9357EA4DB0CA70782268E6D614D75DCE&amp;sdtfrom=v1000&amp;owner=994199947" data-v_h265="" data-v_source_website="" data-v_writefrom="" data-videotype="mood" data-clicklog="video" hotclickpath="" hotdomain="" data-version="3" data-param="3&amp;videosrc=http%3A%2F%2Fvwecam.gtimg.com%2F1006_081dff72e4004afc99d55064aac8ccbb.f0.mp4%3Fptype%3Dhttp%26vkey%3DC6F0EA8CF63D1C68176C3F7314EF152ADD3E27A48A776895DB8B107B34E3588DAD8BDF26EDFC7E3C9357EA4DB0CA70782268E6D614D75DCE%26sdtfrom%3Dv1000%26owner%3D994199947&amp;type=1&amp;org_vidio_url=" data-src="/qzone/app/mood/richinfo_view.html" data-width="512" data-height="512" data-type="popup" data-title="" data-config="" data-vfeed-id="vfeed_994199947_311_8b49423b4cabc359f53f0f00_" data-newplayer="1" data-newplayer-id="vfeed_994199947_311_8b49423b4cabc359f53f0f00_" style="background-color: rgb(0, 0, 0); width: 560px; height: 560px;"><div class="video-img j-videofeed-imgctn" style="width: 560px; height: 560px;"><img src="http://b242.photo.store.qq.com/psb?/V11Q0mso0RE9ia/y8Hn5PLyZV0jJVfyO9q4uoITS44WUCw1s89yNqWfaI4!/c/dPIAAAAAAAAA&amp;bo=HALAAwAAAAARF*0!" data-oriwidth="512" data-oriheight="910" style="width: 315.077px; height: 560px; margin: 0px 122.462px;"><i class="ui-icon icon-media-play j-videofeed-icon-play"></i><span class="video-loading j-videofeed-icon-loading" style="display:none;"><i class="inner"></i></span></div><div class="j-videofeed-flashctn" data-need_hide_when_inited="1" style="position: absolute; left: 0px; top: 0px; width: 560px; height: 560px; overflow: hidden;"><object width="100%" height="100%" id="vfeed_994199947_311_8b49423b4cabc359f53f0f00___flashobj" name="vfeed_994199947_311_8b49423b4cabc359f53f0f00___flashobj" data="//qzs.qq.com/qzone/client/photo/swf/MPlayer/MicroVideoPlayerEx.swf" type="application/x-shockwave-flash"><param name="allowFullScreen" value="true"><param name="allowNetworking" value="all"><param name="allowScriptAccess" value="always"><param name="wmode" value="opaque"><param name="flashvars" value="log=0&amp;vurl=http%3A%2F%2Fvwecam.gtimg.com%2F1006_081dff72e4004afc99d55064aac8ccbb.f0.mp4%3Fptype%3Dhttp%26vkey%3DC6F0EA8CF63D1C68176C3F7314EF152ADD3E27A48A776895DB8B107B34E3588DAD8BDF26EDFC7E3C9357EA4DB0CA70782268E6D614D75DCE%26sdtfrom%3Dv1000%26owner%3D994199947&amp;vurlBak=&amp;noloading=1&amp;auto=1&amp;mute=1&amp;hasPopup=1&amp;hasViewer=1&amp;onFlashInited=window.__qzvCbsMap%5B%22vfeed_994199947_311_8b49423b4cabc359f53f0f00___flashobj%22%5D.onFlashInited&amp;onChangePlayer=window.__qzvCbsMap%5B%22vfeed_994199947_311_8b49423b4cabc359f53f0f00___flashobj%22%5D.onChangePlayer&amp;onReportLog=window.__qzvCbsMap%5B%22vfeed_994199947_311_8b49423b4cabc359f53f0f00___flashobj%22%5D.onReportLog&amp;onError=window.__qzvCbsMap%5B%22vfeed_994199947_311_8b49423b4cabc359f53f0f00___flashobj%22%5D.onError&amp;onMetaData=window.__qzvCbsMap%5B%22vfeed_994199947_311_8b49423b4cabc359f53f0f00___flashobj%22%5D.onMetaData&amp;onChangeState=window.__qzvCbsMap%5B%22vfeed_994199947_311_8b49423b4cabc359f53f0f00___flashobj%22%5D.onChangeState&amp;onPlayStart=window.__qzvCbsMap%5B%22vfeed_994199947_311_8b49423b4cabc359f53f0f00___flashobj%22%5D.onPlayStart&amp;onPlayStop=window.__qzvCbsMap%5B%22vfeed_994199947_311_8b49423b4cabc359f53f0f00___flashobj%22%5D.onPlayStop&amp;onSlideStart=window.__qzvCbsMap%5B%22vfeed_994199947_311_8b49423b4cabc359f53f0f00___flashobj%22%5D.onSlideStart&amp;onSlideStop=window.__qzvCbsMap%5B%22vfeed_994199947_311_8b49423b4cabc359f53f0f00___flashobj%22%5D.onSlideStop&amp;onSkinAction=window.__qzvCbsMap%5B%22vfeed_994199947_311_8b49423b4cabc359f53f0f00___flashobj%22%5D.onSkinAction&amp;onChangeMute=window.__qzvCbsMap%5B%22vfeed_994199947_311_8b49423b4cabc359f53f0f00___flashobj%22%5D.onChangeMute&amp;onChangeFull=window.__qzvCbsMap%5B%22vfeed_994199947_311_8b49423b4cabc359f53f0f00___flashobj%22%5D.onChangeFull&amp;onOpenPopup=window.__qzvCbsMap%5B%22vfeed_994199947_311_8b49423b4cabc359f53f0f00___flashobj%22%5D.onOpenPopup&amp;onOpenViewer=window.__qzvCbsMap%5B%22vfeed_994199947_311_8b49423b4cabc359f53f0f00___flashobj%22%5D.onOpenViewer&amp;onClickMask=window.__qzvCbsMap%5B%22vfeed_994199947_311_8b49423b4cabc359f53f0f00___flashobj%22%5D.onClickMask"></object></div></a></div></div><div class="f-reprint">        <p class="item">        <i class="fui-icon icon-print-phone"></i><span class="ui-mr8 state">来自&nbsp;<a href="http://z.qzone.com?from=iphonegrzxpl" target="_blank" class=" phone-style state">iPhone 7</a>&nbsp;</span>        </p>                            </div>    </div></div></div></div><div class="f-single-foot"><div class="f-op-detail f-detail content-line"><p class="op-list"><a class="item qz_retweet_btn " href="javascript:;" data-cmd="qz_popup" data-version="4" data-isnewtype="1" data-type="ForwardingBox" data-src="/qzone/app/controls/forwardingBox/forwardingBoxFacade.js" data-clicklog="retweet"><i class="fui-icon icon-op-share"></i></a><span class="item-line"></span>&nbsp;<a href="javascript:;" data-version="6.3" data-cmd="qz_reply" data-link="1" data-clicklog="comment" class=" qz_btn_reply item "><i class="fui-icon icon-op-comment"></i></a>&nbsp;<span class="item-line"></span><a class="item qz_like_btn_v3 " data-islike="0" data-likecnt="32" data-showcount="19" data-unikey="http://user.qzone.qq.com/994199947/mood/8b49423b4cabc359f53f0f00" data-curkey="http://user.qzone.qq.com/994199947/mood/8b49423b4cabc359f53f0f00" data-clicklog="like" href="javascript:;"><i class="fui-icon icon-op-praise"></i></a></p><a href="javascript:;" class="state qz_feed_plugin" data-role="Visitor" data-config="311|8b49423b4cabc359f53f0f00|994199947" data-clicklog="visitor">浏览208次</a></div>        <div class="f-ang-t"></div><div class="f-like-list f-like _likeInfo" likeinfo="32"><div class="icon-btn"><a href="javascript:;" data-islike="0" data-likecnt="32" data-showcount="32" data-unikey="http://user.qzone.qq.com/994199947/mood/8b49423b4cabc359f53f0f00" data-curkey="http://user.qzone.qq.com/994199947/mood/8b49423b4cabc359f53f0f00" data-clicklog="like" class="praise qz_like_prase"><i class="fui-icon icon-list-praise"></i></a><div class="bubble" style="display:none;"><div class="bd">+1</div><b class="arrow arrow-down"></b></div></div><div class="user-list"><a href="http://user.qzone.qq.com/994199947" class="item q_namecard" target="_blank" link="nameCard_994199947 des_994199947">易小军</a>、<a href="http://user.qzone.qq.com/210191109" class="item q_namecard" target="_blank" link="nameCard_210191109 des_210191109">か^＆海℡</a>、<a href="http://user.qzone.qq.com/395461248" class="item q_namecard" target="_blank" link="nameCard_395461248 des_395461248">yul</a>、<a href="http://user.qzone.qq.com/464075338" class="item q_namecard" target="_blank" link="nameCard_464075338 des_464075338">☆ξηrich★</a>、<a href="http://user.qzone.qq.com/569417689" class="item q_namecard" target="_blank" link="nameCard_569417689 des_569417689">丢</a>、<a href="http://user.qzone.qq.com/690631518" class="item q_namecard" target="_blank" link="nameCard_690631518 des_690631518">糯米糕</a>、<a href="http://user.qzone.qq.com/704320324" class="item q_namecard" target="_blank" link="nameCard_704320324 des_704320324">something for nothing</a>、<a href="http://user.qzone.qq.com/773482175" class="item q_namecard" target="_blank" link="nameCard_773482175 des_773482175">胡理山</a>、<a href="http://user.qzone.qq.com/804030147" class="item q_namecard" target="_blank" link="nameCard_804030147 des_804030147">遇见！</a>、<a href="http://user.qzone.qq.com/863144225" class="item q_namecard" target="_blank" link="nameCard_863144225 des_863144225">放空</a>、<a href="http://user.qzone.qq.com/879092096" class="item q_namecard" target="_blank" link="nameCard_879092096 des_879092096">西西里艳阳下、</a>、<a href="http://user.qzone.qq.com/961346218" class="item q_namecard" target="_blank" link="nameCard_961346218 des_961346218">一条</a>、<a href="http://user.qzone.qq.com/996322423" class="item q_namecard" target="_blank" link="nameCard_996322423 des_996322423">狗样</a>、<a href="http://user.qzone.qq.com/1007938230" class="item q_namecard" target="_blank" link="nameCard_1007938230 des_1007938230">南舟</a>、<a href="http://user.qzone.qq.com/1041876105" class="item q_namecard" target="_blank" link="nameCard_1041876105 des_1041876105">.</a>、<a href="http://user.qzone.qq.com/1158955536" class="item q_namecard" target="_blank" link="nameCard_1158955536 des_1158955536">[QQ红包]恭喜发财！</a>、<a href="http://user.qzone.qq.com/1278548659" class="item q_namecard" target="_blank" link="nameCard_1278548659 des_1278548659">会走路的三百块</a>、<a href="http://user.qzone.qq.com/1300683599" class="item q_namecard" target="_blank" link="nameCard_1300683599 des_1300683599">          籽沐.گ</a>、<a href="http://user.qzone.qq.com/1324039562" class="item q_namecard" target="_blank" link="nameCard_1324039562 des_1324039562">自由摄影</a>等<span class="f-like-cnt">32</span>人觉得很赞</div></div><div class="mod-comments" style="padding:0"><div class="comments-list "><ul><li class="comments-item bor3" data-type="commentroot" data-tid="1" data-uin="1158955536" data-nick="[QQ红包]恭喜发财！" data-who="1"><div class="comments-item-bd"><div class="single-reply"><div class="ui-avatar"><a href="http://user.qzone.qq.com/1158955536" target="_blank"><img class="q_namecard" link="nameCard_1158955536 des_1158955536" alt="[QQ红包]恭喜发财！" src="http://qlogo1.store.qq.com/qzone/1158955536/1158955536/30"></a></div><div class="comments-content"><a class="nickname name c_tx q_namecard" link="nameCard_1158955536" target="_blank" href="http://user.qzone.qq.com/1158955536">[QQ红包]恭喜发财！</a>&nbsp; : 哇哦<div class="comments-op"><span class=" ui-mr10 state"> 20:08</span><a class="act-reply none" href="javascript:;" data-cmd="qz_reply" data-version="6.4" data-action="http://taotao.qq.com/cgi-bin/emotion_cgi_re_feeds" data-param="t1_source=&amp;t1_uin=994199947&amp;t1_tid=8b49423b4cabc359f53f0f00&amp;t2_uin=1158955536&amp;t2_tid=1&amp;subdotype=55702&amp;signin=0&amp;sceneid=100" data-charset="utf-8" data-tuin="" data-config="1|1|1|0|1,taotaoact.qzone.qq.com,@InputReply|1,taotaoact.qzone.qq.com,@ClickReply|1,taotaoact.qzone.qq.com,commentPresentClick"><b class="hide-clip">回复</b></a></div></div></div><div class="comments-list mod-comments-sub"><ul><li class="comments-item bor3" data-type="replyroot" data-tid="1" data-uin="994199947" data-nick="易小军" data-who="1"><div class="comments-item-bd"><div class="single-reply"><div class="ui-avatar"><a href="http://user.qzone.qq.com/994199947" target="_blank"><img class="q_namecard" link="nameCard_994199947 des_994199947" alt="易小军" src="http://qlogo4.store.qq.com/qzone/994199947/994199947/30?1356758777"></a></div><div class="comments-content"><a class="nickname name c_tx q_namecard" link="nameCard_994199947" target="_blank" href="http://user.qzone.qq.com/994199947">易小军</a>&nbsp;回复<a class="nickname name c_tx q_namecard" link="nameCard_1158955536" target="_blank" href="http://user.qzone.qq.com/1158955536">[QQ红包]恭喜发财！</a>&nbsp; :  哇哇哦哦哦<div class="comments-op"><span class=" ui-mr10 state"> 20:15</span><a class="act-reply none" href="javascript:;" data-cmd="qz_reply" data-version="6.4" data-action="http://taotao.qq.com/cgi-bin/emotion_cgi_re_feeds" data-param="t1_source=&amp;t1_uin=994199947&amp;t1_tid=8b49423b4cabc359f53f0f00&amp;t2_uin=1158955536&amp;t2_tid=1&amp;subdotype=55702&amp;signin=0&amp;sceneid=100" data-charset="utf-8" data-tuin="" data-config="1|1|1|0|1,taotaoact.qzone.qq.com,@InputReply|1,taotaoact.qzone.qq.com,@ClickReply|1,taotaoact.qzone.qq.com,commentPresentClick"><b class="hide-clip">回复</b></a></div></div></div></div></li><li class="comments-item bor3" data-type="replyroot" data-tid="2" data-uin="1158955536" data-nick="[QQ红包]恭喜发财！" data-who="1"><div class="comments-item-bd"><div class="single-reply"><div class="ui-avatar"><a href="http://user.qzone.qq.com/1158955536" target="_blank"><img class="q_namecard" link="nameCard_1158955536 des_1158955536" alt="[QQ红包]恭喜发财！" src="http://qlogo1.store.qq.com/qzone/1158955536/1158955536/30"></a></div><div class="comments-content"><a class="nickname name c_tx q_namecard" link="nameCard_1158955536" target="_blank" href="http://user.qzone.qq.com/1158955536">[QQ红包]恭喜发财！</a>&nbsp;回复<a class="nickname name c_tx q_namecard" link="nameCard_994199947" target="_blank" href="http://user.qzone.qq.com/994199947">易小军</a>&nbsp; :  帅爆了！<div class="comments-op"><span class=" ui-mr10 state"> 20:28</span><a class="act-reply none" href="javascript:;" data-cmd="qz_reply" data-version="6.4" data-action="http://taotao.qq.com/cgi-bin/emotion_cgi_re_feeds" data-param="t1_source=&amp;t1_uin=994199947&amp;t1_tid=8b49423b4cabc359f53f0f00&amp;t2_uin=1158955536&amp;t2_tid=1&amp;subdotype=55702&amp;signin=0&amp;sceneid=100" data-charset="utf-8" data-tuin="" data-config="1|1|1|0|1,taotaoact.qzone.qq.com,@InputReply|1,taotaoact.qzone.qq.com,@ClickReply|1,taotaoact.qzone.qq.com,commentPresentClick"><b class="hide-clip">回复</b></a></div></div></div></div></li><li class="comments-item bor3" data-type="replyroot" data-tid="3" data-uin="994199947" data-nick="易小军" data-who="1"><div class="comments-item-bd"><div class="single-reply"><div class="ui-avatar"><a href="http://user.qzone.qq.com/994199947" target="_blank"><img class="q_namecard" link="nameCard_994199947 des_994199947" alt="易小军" src="http://qlogo4.store.qq.com/qzone/994199947/994199947/30?1356758777"></a></div><div class="comments-content"><a class="nickname name c_tx q_namecard" link="nameCard_994199947" target="_blank" href="http://user.qzone.qq.com/994199947">易小军</a>&nbsp;回复<a class="nickname name c_tx q_namecard" link="nameCard_1158955536" target="_blank" href="http://user.qzone.qq.com/1158955536">[QQ红包]恭喜发财！</a>&nbsp; :  爆瘦<div class="comments-op"><span class=" ui-mr10 state"> 20:28</span><a class="act-reply none" href="javascript:;" data-cmd="qz_reply" data-version="6.4" data-action="http://taotao.qq.com/cgi-bin/emotion_cgi_re_feeds" data-param="t1_source=&amp;t1_uin=994199947&amp;t1_tid=8b49423b4cabc359f53f0f00&amp;t2_uin=1158955536&amp;t2_tid=1&amp;subdotype=55702&amp;signin=0&amp;sceneid=100" data-charset="utf-8" data-tuin="" data-config="1|1|1|0|1,taotaoact.qzone.qq.com,@InputReply|1,taotaoact.qzone.qq.com,@ClickReply|1,taotaoact.qzone.qq.com,commentPresentClick"><b class="hide-clip">回复</b></a></div></div></div></div></li><li class="comments-item bor3" data-type="replyroot" data-tid="4" data-uin="1158955536" data-nick="[QQ红包]恭喜发财！" data-who="1"><div class="comments-item-bd"><div class="single-reply"><div class="ui-avatar"><a href="http://user.qzone.qq.com/1158955536" target="_blank"><img class="q_namecard" link="nameCard_1158955536 des_1158955536" alt="[QQ红包]恭喜发财！" src="http://qlogo1.store.qq.com/qzone/1158955536/1158955536/30"></a></div><div class="comments-content"><a class="nickname name c_tx q_namecard" link="nameCard_1158955536" target="_blank" href="http://user.qzone.qq.com/1158955536">[QQ红包]恭喜发财！</a>&nbsp;回复<a class="nickname name c_tx q_namecard" link="nameCard_994199947" target="_blank" href="http://user.qzone.qq.com/994199947">易小军</a>&nbsp; :  好帅<div class="comments-op"><span class=" ui-mr10 state"> 20:49</span><a class="act-reply none" href="javascript:;" data-cmd="qz_reply" data-version="6.4" data-action="http://taotao.qq.com/cgi-bin/emotion_cgi_re_feeds" data-param="t1_source=&amp;t1_uin=994199947&amp;t1_tid=8b49423b4cabc359f53f0f00&amp;t2_uin=1158955536&amp;t2_tid=1&amp;subdotype=55702&amp;signin=0&amp;sceneid=100" data-charset="utf-8" data-tuin="" data-config="1|1|1|0|1,taotaoact.qzone.qq.com,@InputReply|1,taotaoact.qzone.qq.com,@ClickReply|1,taotaoact.qzone.qq.com,commentPresentClick"><b class="hide-clip">回复</b></a></div></div></div></div></li></ul></div></div></li><li class="comments-item bor3" data-type="commentroot" data-tid="2" data-uin="398269967" data-nick="FUCK ℡" data-who="1"><div class="comments-item-bd"><div class="single-reply"><div class="ui-avatar"><a href="http://user.qzone.qq.com/398269967" target="_blank"><img class="q_namecard" link="nameCard_398269967 des_398269967" alt="FUCK ℡" src="http://qlogo4.store.qq.com/qzone/398269967/398269967/30"></a></div><div class="comments-content"><a class="nickname name c_tx q_namecard" link="nameCard_398269967" target="_blank" href="http://user.qzone.qq.com/398269967">FUCK ℡</a>&nbsp; : 哇<div class="comments-op"><span class=" ui-mr10 state"> 20:09</span><a class="act-reply none" href="javascript:;" data-cmd="qz_reply" data-version="6.4" data-action="http://taotao.qq.com/cgi-bin/emotion_cgi_re_feeds" data-param="t1_source=&amp;t1_uin=994199947&amp;t1_tid=8b49423b4cabc359f53f0f00&amp;t2_uin=398269967&amp;t2_tid=2&amp;subdotype=55702&amp;signin=0&amp;sceneid=100" data-charset="utf-8" data-tuin="" data-config="1|1|1|0|1,taotaoact.qzone.qq.com,@InputReply|1,taotaoact.qzone.qq.com,@ClickReply|1,taotaoact.qzone.qq.com,commentPresentClick"><b class="hide-clip">回复</b></a></div></div></div><div class="comments-list mod-comments-sub"><ul><li class="comments-item bor3" data-type="replyroot" data-tid="1" data-uin="994199947" data-nick="易小军" data-who="1"><div class="comments-item-bd"><div class="single-reply"><div class="ui-avatar"><a href="http://user.qzone.qq.com/994199947" target="_blank"><img class="q_namecard" link="nameCard_994199947 des_994199947" alt="易小军" src="http://qlogo4.store.qq.com/qzone/994199947/994199947/30?1356758777"></a></div><div class="comments-content"><a class="nickname name c_tx q_namecard" link="nameCard_994199947" target="_blank" href="http://user.qzone.qq.com/994199947">易小军</a>&nbsp;回复<a class="nickname name c_tx q_namecard" link="nameCard_398269967" target="_blank" href="http://user.qzone.qq.com/398269967">FUCK ℡</a>&nbsp; :  哇哇哇<img src="http://qzonestyle.gtimg.cn/qzone/em/e120.png" title=""><div class="comments-op"><span class=" ui-mr10 state"> 20:15</span><a class="act-reply none" href="javascript:;" data-cmd="qz_reply" data-version="6.4" data-action="http://taotao.qq.com/cgi-bin/emotion_cgi_re_feeds" data-param="t1_source=&amp;t1_uin=994199947&amp;t1_tid=8b49423b4cabc359f53f0f00&amp;t2_uin=398269967&amp;t2_tid=2&amp;subdotype=55702&amp;signin=0&amp;sceneid=100" data-charset="utf-8" data-tuin="" data-config="1|1|1|0|1,taotaoact.qzone.qq.com,@InputReply|1,taotaoact.qzone.qq.com,@ClickReply|1,taotaoact.qzone.qq.com,commentPresentClick"><b class="hide-clip">回复</b></a></div></div></div></div></li><li class="comments-item bor3" data-type="replyroot" data-tid="2" data-uin="398269967" data-nick="FUCK ℡" data-who="1"><div class="comments-item-bd"><div class="single-reply"><div class="ui-avatar"><a href="http://user.qzone.qq.com/398269967" target="_blank"><img class="q_namecard" link="nameCard_398269967 des_398269967" alt="FUCK ℡" src="http://qlogo4.store.qq.com/qzone/398269967/398269967/30"></a></div><div class="comments-content"><a class="nickname name c_tx q_namecard" link="nameCard_398269967" target="_blank" href="http://user.qzone.qq.com/398269967">FUCK ℡</a>&nbsp;回复<a class="nickname name c_tx q_namecard" link="nameCard_994199947" target="_blank" href="http://user.qzone.qq.com/994199947">易小军</a>&nbsp; : 好爽哦，我在了时候都没有<div class="comments-op"><span class=" ui-mr10 state"> 20:15</span><a class="act-reply none" href="javascript:;" data-cmd="qz_reply" data-version="6.4" data-action="http://taotao.qq.com/cgi-bin/emotion_cgi_re_feeds" data-param="t1_source=&amp;t1_uin=994199947&amp;t1_tid=8b49423b4cabc359f53f0f00&amp;t2_uin=398269967&amp;t2_tid=2&amp;subdotype=55702&amp;signin=0&amp;sceneid=100" data-charset="utf-8" data-tuin="" data-config="1|1|1|0|1,taotaoact.qzone.qq.com,@InputReply|1,taotaoact.qzone.qq.com,@ClickReply|1,taotaoact.qzone.qq.com,commentPresentClick"><b class="hide-clip">回复</b></a></div></div></div></div></li></ul></div></div></li><li class="comments-item bor3" data-type="commentroot" data-tid="3" data-uin="569417689" data-nick="丢" data-who="1"><div class="comments-item-bd"><div class="single-reply"><div class="ui-avatar"><a href="http://user.qzone.qq.com/569417689" target="_blank"><img class="q_namecard" link="nameCard_569417689 des_569417689" alt="丢" src="http://qlogo2.store.qq.com/qzone/569417689/569417689/30"></a></div><div class="comments-content"><a class="nickname name c_tx q_namecard" link="nameCard_569417689" target="_blank" href="http://user.qzone.qq.com/569417689">丢</a>&nbsp; : 乐道有健身房了<div class="comments-op"><span class=" ui-mr10 state"> 20:10</span><a class="act-reply none" href="javascript:;" data-cmd="qz_reply" data-version="6.4" data-action="http://taotao.qq.com/cgi-bin/emotion_cgi_re_feeds" data-param="t1_source=&amp;t1_uin=994199947&amp;t1_tid=8b49423b4cabc359f53f0f00&amp;t2_uin=569417689&amp;t2_tid=3&amp;subdotype=55702&amp;signin=0&amp;sceneid=100" data-charset="utf-8" data-tuin="" data-config="1|1|1|0|1,taotaoact.qzone.qq.com,@InputReply|1,taotaoact.qzone.qq.com,@ClickReply|1,taotaoact.qzone.qq.com,commentPresentClick"><b class="hide-clip">回复</b></a></div></div></div><div class="comments-list mod-comments-sub"><ul><li class="comments-item bor3" data-type="replyroot" data-tid="1" data-uin="994199947" data-nick="易小军" data-who="1"><div class="comments-item-bd"><div class="single-reply"><div class="ui-avatar"><a href="http://user.qzone.qq.com/994199947" target="_blank"><img class="q_namecard" link="nameCard_994199947 des_994199947" alt="易小军" src="http://qlogo4.store.qq.com/qzone/994199947/994199947/30?1356758777"></a></div><div class="comments-content"><a class="nickname name c_tx q_namecard" link="nameCard_994199947" target="_blank" href="http://user.qzone.qq.com/994199947">易小军</a>&nbsp;回复<a class="nickname name c_tx q_namecard" link="nameCard_569417689" target="_blank" href="http://user.qzone.qq.com/569417689">丢</a>&nbsp; :  嗯嗯，过来可以耍啦<div class="comments-op"><span class=" ui-mr10 state"> 20:14</span><a class="act-reply none" href="javascript:;" data-cmd="qz_reply" data-version="6.4" data-action="http://taotao.qq.com/cgi-bin/emotion_cgi_re_feeds" data-param="t1_source=&amp;t1_uin=994199947&amp;t1_tid=8b49423b4cabc359f53f0f00&amp;t2_uin=569417689&amp;t2_tid=3&amp;subdotype=55702&amp;signin=0&amp;sceneid=100" data-charset="utf-8" data-tuin="" data-config="1|1|1|0|1,taotaoact.qzone.qq.com,@InputReply|1,taotaoact.qzone.qq.com,@ClickReply|1,taotaoact.qzone.qq.com,commentPresentClick"><b class="hide-clip">回复</b></a></div></div></div></div></li><li class="comments-item bor3" data-type="replyroot" data-tid="2" data-uin="569417689" data-nick="丢" data-who="1"><div class="comments-item-bd"><div class="single-reply"><div class="ui-avatar"><a href="http://user.qzone.qq.com/569417689" target="_blank"><img class="q_namecard" link="nameCard_569417689 des_569417689" alt="丢" src="http://qlogo2.store.qq.com/qzone/569417689/569417689/30"></a></div><div class="comments-content"><a class="nickname name c_tx q_namecard" link="nameCard_569417689" target="_blank" href="http://user.qzone.qq.com/569417689">丢</a>&nbsp;回复<a class="nickname name c_tx q_namecard" link="nameCard_994199947" target="_blank" href="http://user.qzone.qq.com/994199947">易小军</a>&nbsp; : 哈哈哈，二天要过来的<div class="comments-op"><span class=" ui-mr10 state"> 20:24</span><a class="act-reply none" href="javascript:;" data-cmd="qz_reply" data-version="6.4" data-action="http://taotao.qq.com/cgi-bin/emotion_cgi_re_feeds" data-param="t1_source=&amp;t1_uin=994199947&amp;t1_tid=8b49423b4cabc359f53f0f00&amp;t2_uin=569417689&amp;t2_tid=3&amp;subdotype=55702&amp;signin=0&amp;sceneid=100" data-charset="utf-8" data-tuin="" data-config="1|1|1|0|1,taotaoact.qzone.qq.com,@InputReply|1,taotaoact.qzone.qq.com,@ClickReply|1,taotaoact.qzone.qq.com,commentPresentClick"><b class="hide-clip">回复</b></a></div></div></div></div></li><li class="comments-item bor3" data-type="replyroot" data-tid="3" data-uin="994199947" data-nick="易小军" data-who="1"><div class="comments-item-bd"><div class="single-reply"><div class="ui-avatar"><a href="http://user.qzone.qq.com/994199947" target="_blank"><img class="q_namecard" link="nameCard_994199947 des_994199947" alt="易小军" src="http://qlogo4.store.qq.com/qzone/994199947/994199947/30?1356758777"></a></div><div class="comments-content"><a class="nickname name c_tx q_namecard" link="nameCard_994199947" target="_blank" href="http://user.qzone.qq.com/994199947">易小军</a>&nbsp;回复<a class="nickname name c_tx q_namecard" link="nameCard_569417689" target="_blank" href="http://user.qzone.qq.com/569417689">丢</a>&nbsp; :  <img src="http://qzonestyle.gtimg.cn/qzone/em/e401210.gif" title=""><div class="comments-op"><span class=" ui-mr10 state"> 20:24</span><a class="act-reply none" href="javascript:;" data-cmd="qz_reply" data-version="6.4" data-action="http://taotao.qq.com/cgi-bin/emotion_cgi_re_feeds" data-param="t1_source=&amp;t1_uin=994199947&amp;t1_tid=8b49423b4cabc359f53f0f00&amp;t2_uin=569417689&amp;t2_tid=3&amp;subdotype=55702&amp;signin=0&amp;sceneid=100" data-charset="utf-8" data-tuin="" data-config="1|1|1|0|1,taotaoact.qzone.qq.com,@InputReply|1,taotaoact.qzone.qq.com,@ClickReply|1,taotaoact.qzone.qq.com,commentPresentClick"><b class="hide-clip">回复</b></a></div></div></div></div></li></ul></div></div></li><li class="comments-item bor3" data-type="commentroot" data-tid="4" data-uin="2501062583" data-nick="如果" data-who="1"><div class="comments-item-bd"><div class="single-reply"><div class="ui-avatar"><a href="http://user.qzone.qq.com/2501062583" target="_blank"><img class="q_namecard" link="nameCard_2501062583 des_2501062583" alt="如果" src="http://qlogo4.store.qq.com/qzone/2501062583/2501062583/30"></a></div><div class="comments-content"><a class="nickname name c_tx q_namecard" link="nameCard_2501062583" target="_blank" href="http://user.qzone.qq.com/2501062583">如果</a>&nbsp; : nb<div class="comments-op"><span class=" ui-mr10 state"> 20:14</span><a class="act-reply none" href="javascript:;" data-cmd="qz_reply" data-version="6.4" data-action="http://taotao.qq.com/cgi-bin/emotion_cgi_re_feeds" data-param="t1_source=&amp;t1_uin=994199947&amp;t1_tid=8b49423b4cabc359f53f0f00&amp;t2_uin=2501062583&amp;t2_tid=4&amp;subdotype=55702&amp;signin=0&amp;sceneid=100" data-charset="utf-8" data-tuin="" data-config="1|1|1|0|1,taotaoact.qzone.qq.com,@InputReply|1,taotaoact.qzone.qq.com,@ClickReply|1,taotaoact.qzone.qq.com,commentPresentClick"><b class="hide-clip">回复</b></a></div></div></div><div class="comments-list mod-comments-sub"><ul><li class="comments-item bor3" data-type="replyroot" data-tid="1" data-uin="994199947" data-nick="易小军" data-who="1"><div class="comments-item-bd"><div class="single-reply"><div class="ui-avatar"><a href="http://user.qzone.qq.com/994199947" target="_blank"><img class="q_namecard" link="nameCard_994199947 des_994199947" alt="易小军" src="http://qlogo4.store.qq.com/qzone/994199947/994199947/30?1356758777"></a></div><div class="comments-content"><a class="nickname name c_tx q_namecard" link="nameCard_994199947" target="_blank" href="http://user.qzone.qq.com/994199947">易小军</a>&nbsp;回复<a class="nickname name c_tx q_namecard" link="nameCard_2501062583" target="_blank" href="http://user.qzone.qq.com/2501062583">如果</a>&nbsp; :  必须滴，杠杠滴<div class="comments-op"><span class=" ui-mr10 state"> 20:16</span><a class="act-reply none" href="javascript:;" data-cmd="qz_reply" data-version="6.4" data-action="http://taotao.qq.com/cgi-bin/emotion_cgi_re_feeds" data-param="t1_source=&amp;t1_uin=994199947&amp;t1_tid=8b49423b4cabc359f53f0f00&amp;t2_uin=2501062583&amp;t2_tid=4&amp;subdotype=55702&amp;signin=0&amp;sceneid=100" data-charset="utf-8" data-tuin="" data-config="1|1|1|0|1,taotaoact.qzone.qq.com,@InputReply|1,taotaoact.qzone.qq.com,@ClickReply|1,taotaoact.qzone.qq.com,commentPresentClick"><b class="hide-clip">回复</b></a></div></div></div></div></li><li class="comments-item bor3" data-type="replyroot" data-tid="2" data-uin="2501062583" data-nick="如果" data-who="1"><div class="comments-item-bd"><div class="single-reply"><div class="ui-avatar"><a href="http://user.qzone.qq.com/2501062583" target="_blank"><img class="q_namecard" link="nameCard_2501062583 des_2501062583" alt="如果" src="http://qlogo4.store.qq.com/qzone/2501062583/2501062583/30"></a></div><div class="comments-content"><a class="nickname name c_tx q_namecard" link="nameCard_2501062583" target="_blank" href="http://user.qzone.qq.com/2501062583">如果</a>&nbsp;回复<a class="nickname name c_tx q_namecard" link="nameCard_994199947" target="_blank" href="http://user.qzone.qq.com/994199947">易小军</a>&nbsp; : 私人健身房都搞起了<div class="comments-op"><span class=" ui-mr10 state"> 20:18</span><a class="act-reply none" href="javascript:;" data-cmd="qz_reply" data-version="6.4" data-action="http://taotao.qq.com/cgi-bin/emotion_cgi_re_feeds" data-param="t1_source=&amp;t1_uin=994199947&amp;t1_tid=8b49423b4cabc359f53f0f00&amp;t2_uin=2501062583&amp;t2_tid=4&amp;subdotype=55702&amp;signin=0&amp;sceneid=100" data-charset="utf-8" data-tuin="" data-config="1|1|1|0|1,taotaoact.qzone.qq.com,@InputReply|1,taotaoact.qzone.qq.com,@ClickReply|1,taotaoact.qzone.qq.com,commentPresentClick"><b class="hide-clip">回复</b></a></div></div></div></div></li><li class="comments-item bor3" data-type="replyroot" data-tid="3" data-uin="994199947" data-nick="易小军" data-who="1"><div class="comments-item-bd"><div class="single-reply"><div class="ui-avatar"><a href="http://user.qzone.qq.com/994199947" target="_blank"><img class="q_namecard" link="nameCard_994199947 des_994199947" alt="易小军" src="http://qlogo4.store.qq.com/qzone/994199947/994199947/30?1356758777"></a></div><div class="comments-content"><a class="nickname name c_tx q_namecard" link="nameCard_994199947" target="_blank" href="http://user.qzone.qq.com/994199947">易小军</a>&nbsp;回复<a class="nickname name c_tx q_namecard" link="nameCard_2501062583" target="_blank" href="http://user.qzone.qq.com/2501062583">如果</a>&nbsp; :  那必须滴嘛<img src="http://qzonestyle.gtimg.cn/qzone/em/e104.png" title=""><div class="comments-op"><span class=" ui-mr10 state"> 20:19</span><a class="act-reply none" href="javascript:;" data-cmd="qz_reply" data-version="6.4" data-action="http://taotao.qq.com/cgi-bin/emotion_cgi_re_feeds" data-param="t1_source=&amp;t1_uin=994199947&amp;t1_tid=8b49423b4cabc359f53f0f00&amp;t2_uin=2501062583&amp;t2_tid=4&amp;subdotype=55702&amp;signin=0&amp;sceneid=100" data-charset="utf-8" data-tuin="" data-config="1|1|1|0|1,taotaoact.qzone.qq.com,@InputReply|1,taotaoact.qzone.qq.com,@ClickReply|1,taotaoact.qzone.qq.com,commentPresentClick"><b class="hide-clip">回复</b></a></div></div></div></div></li></ul></div></div></li><li class="comments-item bor3" data-type="commentroot" data-tid="5" data-uin="1501206418" data-nick="“          ”" data-who="1"><div class="comments-item-bd"><div class="single-reply"><div class="ui-avatar"><a href="http://user.qzone.qq.com/1501206418" target="_blank"><img class="q_namecard" link="nameCard_1501206418 des_1501206418" alt="“          ”" src="http://qlogo3.store.qq.com/qzone/1501206418/1501206418/30"></a></div><div class="comments-content"><a class="nickname name c_tx q_namecard" link="nameCard_1501206418" target="_blank" href="http://user.qzone.qq.com/1501206418">“          ”</a>&nbsp; : 厉害<div class="comments-op"><span class=" ui-mr10 state"> 20:15</span><a class="act-reply none" href="javascript:;" data-cmd="qz_reply" data-version="6.4" data-action="http://taotao.qq.com/cgi-bin/emotion_cgi_re_feeds" data-param="t1_source=&amp;t1_uin=994199947&amp;t1_tid=8b49423b4cabc359f53f0f00&amp;t2_uin=1501206418&amp;t2_tid=5&amp;subdotype=55702&amp;signin=0&amp;sceneid=100" data-charset="utf-8" data-tuin="" data-config="1|1|1|0|1,taotaoact.qzone.qq.com,@InputReply|1,taotaoact.qzone.qq.com,@ClickReply|1,taotaoact.qzone.qq.com,commentPresentClick"><b class="hide-clip">回复</b></a></div></div></div><div class="comments-list mod-comments-sub"><ul><li class="comments-item bor3" data-type="replyroot" data-tid="1" data-uin="994199947" data-nick="易小军" data-who="1"><div class="comments-item-bd"><div class="single-reply"><div class="ui-avatar"><a href="http://user.qzone.qq.com/994199947" target="_blank"><img class="q_namecard" link="nameCard_994199947 des_994199947" alt="易小军" src="http://qlogo4.store.qq.com/qzone/994199947/994199947/30?1356758777"></a></div><div class="comments-content"><a class="nickname name c_tx q_namecard" link="nameCard_994199947" target="_blank" href="http://user.qzone.qq.com/994199947">易小军</a>&nbsp;回复<a class="nickname name c_tx q_namecard" link="nameCard_1501206418" target="_blank" href="http://user.qzone.qq.com/1501206418">“          ”</a>&nbsp; :  必须滴<div class="comments-op"><span class=" ui-mr10 state"> 20:15</span><a class="act-reply none" href="javascript:;" data-cmd="qz_reply" data-version="6.4" data-action="http://taotao.qq.com/cgi-bin/emotion_cgi_re_feeds" data-param="t1_source=&amp;t1_uin=994199947&amp;t1_tid=8b49423b4cabc359f53f0f00&amp;t2_uin=1501206418&amp;t2_tid=5&amp;subdotype=55702&amp;signin=0&amp;sceneid=100" data-charset="utf-8" data-tuin="" data-config="1|1|1|0|1,taotaoact.qzone.qq.com,@InputReply|1,taotaoact.qzone.qq.com,@ClickReply|1,taotaoact.qzone.qq.com,commentPresentClick"><b class="hide-clip">回复</b></a></div></div></div></div></li><li class="comments-item bor3" data-type="replyroot" data-tid="2" data-uin="1501206418" data-nick="“          ”" data-who="1"><div class="comments-item-bd"><div class="single-reply"><div class="ui-avatar"><a href="http://user.qzone.qq.com/1501206418" target="_blank"><img class="q_namecard" link="nameCard_1501206418 des_1501206418" alt="“          ”" src="http://qlogo3.store.qq.com/qzone/1501206418/1501206418/30"></a></div><div class="comments-content"><a class="nickname name c_tx q_namecard" link="nameCard_1501206418" target="_blank" href="http://user.qzone.qq.com/1501206418">“          ”</a>&nbsp;回复<a class="nickname name c_tx q_namecard" link="nameCard_994199947" target="_blank" href="http://user.qzone.qq.com/994199947">易小军</a>&nbsp; :  小易加油<img src="http://qzonestyle.gtimg.cn/qzone/em/e401148.gif" title=""><div class="comments-op"><span class=" ui-mr10 state"> 20:16</span><a class="act-reply none" href="javascript:;" data-cmd="qz_reply" data-version="6.4" data-action="http://taotao.qq.com/cgi-bin/emotion_cgi_re_feeds" data-param="t1_source=&amp;t1_uin=994199947&amp;t1_tid=8b49423b4cabc359f53f0f00&amp;t2_uin=1501206418&amp;t2_tid=5&amp;subdotype=55702&amp;signin=0&amp;sceneid=100" data-charset="utf-8" data-tuin="" data-config="1|1|1|0|1,taotaoact.qzone.qq.com,@InputReply|1,taotaoact.qzone.qq.com,@ClickReply|1,taotaoact.qzone.qq.com,commentPresentClick"><b class="hide-clip">回复</b></a></div></div></div></div></li><li class="comments-item bor3" data-type="replyroot" data-tid="3" data-uin="994199947" data-nick="易小军" data-who="1"><div class="comments-item-bd"><div class="single-reply"><div class="ui-avatar"><a href="http://user.qzone.qq.com/994199947" target="_blank"><img class="q_namecard" link="nameCard_994199947 des_994199947" alt="易小军" src="http://qlogo4.store.qq.com/qzone/994199947/994199947/30?1356758777"></a></div><div class="comments-content"><a class="nickname name c_tx q_namecard" link="nameCard_994199947" target="_blank" href="http://user.qzone.qq.com/994199947">易小军</a>&nbsp;回复<a class="nickname name c_tx q_namecard" link="nameCard_1501206418" target="_blank" href="http://user.qzone.qq.com/1501206418">“          ”</a>&nbsp; :  那是那是，得加油<img src="http://qzonestyle.gtimg.cn/qzone/em/e113.png" title=""><div class="comments-op"><span class=" ui-mr10 state"> 20:16</span><a class="act-reply none" href="javascript:;" data-cmd="qz_reply" data-version="6.4" data-action="http://taotao.qq.com/cgi-bin/emotion_cgi_re_feeds" data-param="t1_source=&amp;t1_uin=994199947&amp;t1_tid=8b49423b4cabc359f53f0f00&amp;t2_uin=1501206418&amp;t2_tid=5&amp;subdotype=55702&amp;signin=0&amp;sceneid=100" data-charset="utf-8" data-tuin="" data-config="1|1|1|0|1,taotaoact.qzone.qq.com,@InputReply|1,taotaoact.qzone.qq.com,@ClickReply|1,taotaoact.qzone.qq.com,commentPresentClick"><b class="hide-clip">回复</b></a></div></div></div></div></li></ul></div></div></li><li class="comments-item bor3" data-type="commentroot" data-tid="6" data-uin="1309408821" data-nick="不变的小春" data-who="1"><div class="comments-item-bd"><div class="single-reply"><div class="ui-avatar"><a href="http://user.qzone.qq.com/1309408821" target="_blank"><img class="q_namecard" link="nameCard_1309408821 des_1309408821" alt="不变的小春" src="http://qlogo2.store.qq.com/qzone/1309408821/1309408821/30"></a></div><div class="comments-content"><a class="nickname name c_tx q_namecard" link="nameCard_1309408821" target="_blank" href="http://user.qzone.qq.com/1309408821">不变的小春</a>&nbsp; : 哇哦。小易好帅了哦<img src="http://qzonestyle.gtimg.cn/qzone/em/e121009.gif" title=""><img src="http://qzonestyle.gtimg.cn/qzone/em/e121009.gif" title=""><div class="comments-op"><span class=" ui-mr10 state"> 20:21</span><a class="act-reply none" href="javascript:;" data-cmd="qz_reply" data-version="6.4" data-action="http://taotao.qq.com/cgi-bin/emotion_cgi_re_feeds" data-param="t1_source=&amp;t1_uin=994199947&amp;t1_tid=8b49423b4cabc359f53f0f00&amp;t2_uin=1309408821&amp;t2_tid=6&amp;subdotype=55702&amp;signin=0&amp;sceneid=100" data-charset="utf-8" data-tuin="" data-config="1|1|1|0|1,taotaoact.qzone.qq.com,@InputReply|1,taotaoact.qzone.qq.com,@ClickReply|1,taotaoact.qzone.qq.com,commentPresentClick"><b class="hide-clip">回复</b></a></div></div></div><div class="comments-list mod-comments-sub"><ul><li class="comments-item bor3" data-type="replyroot" data-tid="1" data-uin="994199947" data-nick="易小军" data-who="1"><div class="comments-item-bd"><div class="single-reply"><div class="ui-avatar"><a href="http://user.qzone.qq.com/994199947" target="_blank"><img class="q_namecard" link="nameCard_994199947 des_994199947" alt="易小军" src="http://qlogo4.store.qq.com/qzone/994199947/994199947/30?1356758777"></a></div><div class="comments-content"><a class="nickname name c_tx q_namecard" link="nameCard_994199947" target="_blank" href="http://user.qzone.qq.com/994199947">易小军</a>&nbsp;回复<a class="nickname name c_tx q_namecard" link="nameCard_1309408821" target="_blank" href="http://user.qzone.qq.com/1309408821">不变的小春</a>&nbsp; :  谢谢<img src="http://qzonestyle.gtimg.cn/qzone/em/e104.png" title=""><div class="comments-op"><span class=" ui-mr10 state"> 20:23</span><a class="act-reply none" href="javascript:;" data-cmd="qz_reply" data-version="6.4" data-action="http://taotao.qq.com/cgi-bin/emotion_cgi_re_feeds" data-param="t1_source=&amp;t1_uin=994199947&amp;t1_tid=8b49423b4cabc359f53f0f00&amp;t2_uin=1309408821&amp;t2_tid=6&amp;subdotype=55702&amp;signin=0&amp;sceneid=100" data-charset="utf-8" data-tuin="" data-config="1|1|1|0|1,taotaoact.qzone.qq.com,@InputReply|1,taotaoact.qzone.qq.com,@ClickReply|1,taotaoact.qzone.qq.com,commentPresentClick"><b class="hide-clip">回复</b></a></div></div></div></div></li><li class="comments-item bor3" data-type="replyroot" data-tid="2" data-uin="1309408821" data-nick="不变的小春" data-who="1"><div class="comments-item-bd"><div class="single-reply"><div class="ui-avatar"><a href="http://user.qzone.qq.com/1309408821" target="_blank"><img class="q_namecard" link="nameCard_1309408821 des_1309408821" alt="不变的小春" src="http://qlogo2.store.qq.com/qzone/1309408821/1309408821/30"></a></div><div class="comments-content"><a class="nickname name c_tx q_namecard" link="nameCard_1309408821" target="_blank" href="http://user.qzone.qq.com/1309408821">不变的小春</a>&nbsp;回复<a class="nickname name c_tx q_namecard" link="nameCard_994199947" target="_blank" href="http://user.qzone.qq.com/994199947">易小军</a>&nbsp; :  <img src="http://qzonestyle.gtimg.cn/qzone/em/e400644.gif" title=""><img src="http://qzonestyle.gtimg.cn/qzone/em/e400644.gif" title=""><div class="comments-op"><span class=" ui-mr10 state"> 20:24</span><a class="act-reply none" href="javascript:;" data-cmd="qz_reply" data-version="6.4" data-action="http://taotao.qq.com/cgi-bin/emotion_cgi_re_feeds" data-param="t1_source=&amp;t1_uin=994199947&amp;t1_tid=8b49423b4cabc359f53f0f00&amp;t2_uin=1309408821&amp;t2_tid=6&amp;subdotype=55702&amp;signin=0&amp;sceneid=100" data-charset="utf-8" data-tuin="" data-config="1|1|1|0|1,taotaoact.qzone.qq.com,@InputReply|1,taotaoact.qzone.qq.com,@ClickReply|1,taotaoact.qzone.qq.com,commentPresentClick"><b class="hide-clip">回复</b></a></div></div></div></div></li><li class="comments-item bor3" data-type="replyroot" data-tid="3" data-uin="994199947" data-nick="易小军" data-who="1"><div class="comments-item-bd"><div class="single-reply"><div class="ui-avatar"><a href="http://user.qzone.qq.com/994199947" target="_blank"><img class="q_namecard" link="nameCard_994199947 des_994199947" alt="易小军" src="http://qlogo4.store.qq.com/qzone/994199947/994199947/30?1356758777"></a></div><div class="comments-content"><a class="nickname name c_tx q_namecard" link="nameCard_994199947" target="_blank" href="http://user.qzone.qq.com/994199947">易小军</a>&nbsp;回复<a class="nickname name c_tx q_namecard" link="nameCard_1309408821" target="_blank" href="http://user.qzone.qq.com/1309408821">不变的小春</a>&nbsp; :  <img src="http://qzonestyle.gtimg.cn/qzone/em/e400644.gif" title=""><div class="comments-op"><span class=" ui-mr10 state"> 20:24</span><a class="act-reply none" href="javascript:;" data-cmd="qz_reply" data-version="6.4" data-action="http://taotao.qq.com/cgi-bin/emotion_cgi_re_feeds" data-param="t1_source=&amp;t1_uin=994199947&amp;t1_tid=8b49423b4cabc359f53f0f00&amp;t2_uin=1309408821&amp;t2_tid=6&amp;subdotype=55702&amp;signin=0&amp;sceneid=100" data-charset="utf-8" data-tuin="" data-config="1|1|1|0|1,taotaoact.qzone.qq.com,@InputReply|1,taotaoact.qzone.qq.com,@ClickReply|1,taotaoact.qzone.qq.com,commentPresentClick"><b class="hide-clip">回复</b></a></div></div></div></div></li></ul></div></div></li><li class="comments-item bor3" data-type="commentroot" data-tid="7" data-uin="2110859882" data-nick="OoO" data-who="1"><div class="comments-item-bd"><div class="single-reply"><div class="ui-avatar"><a href="http://user.qzone.qq.com/2110859882" target="_blank"><img class="q_namecard" link="nameCard_2110859882 des_2110859882" alt="OoO" src="http://qlogo3.store.qq.com/qzone/2110859882/2110859882/30"></a></div><div class="comments-content"><a class="nickname name c_tx q_namecard" link="nameCard_2110859882" target="_blank" href="http://user.qzone.qq.com/2110859882">OoO</a>&nbsp; : 哇，越来越牛逼了<div class="comments-op"><span class=" ui-mr10 state"> 20:31</span><a class="act-reply none" href="javascript:;" data-cmd="qz_reply" data-version="6.4" data-action="http://taotao.qq.com/cgi-bin/emotion_cgi_re_feeds" data-param="t1_source=&amp;t1_uin=994199947&amp;t1_tid=8b49423b4cabc359f53f0f00&amp;t2_uin=2110859882&amp;t2_tid=7&amp;subdotype=55702&amp;signin=0&amp;sceneid=100" data-charset="utf-8" data-tuin="" data-config="1|1|1|0|1,taotaoact.qzone.qq.com,@InputReply|1,taotaoact.qzone.qq.com,@ClickReply|1,taotaoact.qzone.qq.com,commentPresentClick"><b class="hide-clip">回复</b></a></div></div></div></div></li><li class="comments-item bor3" data-type="commentroot" data-tid="8" data-uin="1570211662" data-nick="◇" data-who="1"><div class="comments-item-bd"><div class="single-reply"><div class="ui-avatar"><a href="http://user.qzone.qq.com/1570211662" target="_blank"><img class="q_namecard" link="nameCard_1570211662 des_1570211662" alt="◇" src="http://qlogo3.store.qq.com/qzone/1570211662/1570211662/30"></a></div><div class="comments-content"><a class="nickname name c_tx q_namecard" link="nameCard_1570211662" target="_blank" href="http://user.qzone.qq.com/1570211662">◇</a>&nbsp; : 不得不说一句，小易好帅了<img src="http://qzonestyle.gtimg.cn/qzone/em/e113.png" title=""><img src="http://qzonestyle.gtimg.cn/qzone/em/e113.png" title=""><img src="http://qzonestyle.gtimg.cn/qzone/em/e113.png" title=""><div class="comments-op"><span class=" ui-mr10 state"> 20:34</span><a class="act-reply none" href="javascript:;" data-cmd="qz_reply" data-version="6.4" data-action="http://taotao.qq.com/cgi-bin/emotion_cgi_re_feeds" data-param="t1_source=&amp;t1_uin=994199947&amp;t1_tid=8b49423b4cabc359f53f0f00&amp;t2_uin=1570211662&amp;t2_tid=8&amp;subdotype=55702&amp;signin=0&amp;sceneid=100" data-charset="utf-8" data-tuin="" data-config="1|1|1|0|1,taotaoact.qzone.qq.com,@InputReply|1,taotaoact.qzone.qq.com,@ClickReply|1,taotaoact.qzone.qq.com,commentPresentClick"><b class="hide-clip">回复</b></a></div></div></div></div></li><li class="comments-item bor3" data-type="commentroot" data-tid="9" data-uin="932012043" data-nick="♡[em]e257868[/em]   QiuQiu ╮" data-who="1"><div class="comments-item-bd"><div class="single-reply"><div class="ui-avatar"><a href="http://user.qzone.qq.com/932012043" target="_blank"><img class="q_namecard" link="nameCard_932012043 des_932012043" alt="♡[em]e257868[/em]   QiuQiu ╮" src="http://qlogo4.store.qq.com/qzone/932012043/932012043/30"></a></div><div class="comments-content"><a class="nickname name c_tx q_namecard" link="nameCard_932012043" target="_blank" href="http://user.qzone.qq.com/932012043">♡<img src="http://qzonestyle.gtimg.cn/qzone/em/e257868.gif" title="">   QiuQiu ╮</a>&nbsp; : 哇，易师傅瘦了！！！<div class="comments-op"><span class=" ui-mr10 state"> 20:43</span><a class="act-reply none" href="javascript:;" data-cmd="qz_reply" data-version="6.4" data-action="http://taotao.qq.com/cgi-bin/emotion_cgi_re_feeds" data-param="t1_source=&amp;t1_uin=994199947&amp;t1_tid=8b49423b4cabc359f53f0f00&amp;t2_uin=932012043&amp;t2_tid=9&amp;subdotype=55702&amp;signin=0&amp;sceneid=100" data-charset="utf-8" data-tuin="" data-config="1|1|1|0|1,taotaoact.qzone.qq.com,@InputReply|1,taotaoact.qzone.qq.com,@ClickReply|1,taotaoact.qzone.qq.com,commentPresentClick"><b class="hide-clip">回复</b></a></div></div></div></div></li><li class="comments-item bor3" data-type="commentroot" data-tid="10" data-uin="704320324" data-nick="something for nothing" data-who="1"><div class="comments-item-bd"><div class="single-reply"><div class="ui-avatar"><a href="http://user.qzone.qq.com/704320324" target="_blank"><img class="q_namecard" link="nameCard_704320324 des_704320324" alt="something for nothing" src="http://qlogo1.store.qq.com/qzone/704320324/704320324/30"></a></div><div class="comments-content"><a class="nickname name c_tx q_namecard" link="nameCard_704320324" target="_blank" href="http://user.qzone.qq.com/704320324">something for nothing</a>&nbsp; : 666<div class="comments-op"><span class=" ui-mr10 state"> 20:49</span><a class="act-reply none" href="javascript:;" data-cmd="qz_reply" data-version="6.4" data-action="http://taotao.qq.com/cgi-bin/emotion_cgi_re_feeds" data-param="t1_source=&amp;t1_uin=994199947&amp;t1_tid=8b49423b4cabc359f53f0f00&amp;t2_uin=704320324&amp;t2_tid=10&amp;subdotype=55702&amp;signin=0&amp;sceneid=100" data-charset="utf-8" data-tuin="" data-config="1|1|1|0|1,taotaoact.qzone.qq.com,@InputReply|1,taotaoact.qzone.qq.com,@ClickReply|1,taotaoact.qzone.qq.com,commentPresentClick"><b class="hide-clip">回复</b></a></div></div></div></div></li><li class="comments-item bor3" data-type="commentroot" data-tid="11" data-uin="1476463990" data-nick="城" data-who="1"><div class="comments-item-bd"><div class="single-reply"><div class="ui-avatar"><a href="http://user.qzone.qq.com/1476463990" target="_blank"><img class="q_namecard" link="nameCard_1476463990 des_1476463990" alt="城" src="http://qlogo3.store.qq.com/qzone/1476463990/1476463990/30"></a></div><div class="comments-content"><a class="nickname name c_tx q_namecard" link="nameCard_1476463990" target="_blank" href="http://user.qzone.qq.com/1476463990">城</a>&nbsp; : 瘦了<div class="comments-op"><span class=" ui-mr10 state"> 20:56</span><a class="act-reply none" href="javascript:;" data-cmd="qz_reply" data-version="6.4" data-action="http://taotao.qq.com/cgi-bin/emotion_cgi_re_feeds" data-param="t1_source=&amp;t1_uin=994199947&amp;t1_tid=8b49423b4cabc359f53f0f00&amp;t2_uin=1476463990&amp;t2_tid=11&amp;subdotype=55702&amp;signin=0&amp;sceneid=100" data-charset="utf-8" data-tuin="" data-config="1|1|1|0|1,taotaoact.qzone.qq.com,@InputReply|1,taotaoact.qzone.qq.com,@ClickReply|1,taotaoact.qzone.qq.com,commentPresentClick"><b class="hide-clip">回复</b></a></div></div></div></div></li><li class="comments-item bor3" data-type="commentroot" data-tid="12" data-uin="1007938230" data-nick="南舟" data-who="1"><div class="comments-item-bd"><div class="single-reply"><div class="ui-avatar"><a href="http://user.qzone.qq.com/1007938230" target="_blank"><img class="q_namecard" link="nameCard_1007938230 des_1007938230" alt="南舟" src="http://qlogo3.store.qq.com/qzone/1007938230/1007938230/30"></a></div><div class="comments-content"><a class="nickname name c_tx q_namecard" link="nameCard_1007938230" target="_blank" href="http://user.qzone.qq.com/1007938230">南舟</a>&nbsp; : 要瘦了要瘦了<div class="comments-op"><span class=" ui-mr10 state"> 21:07</span><a class="act-reply none" href="javascript:;" data-cmd="qz_reply" data-version="6.4" data-action="http://taotao.qq.com/cgi-bin/emotion_cgi_re_feeds" data-param="t1_source=&amp;t1_uin=994199947&amp;t1_tid=8b49423b4cabc359f53f0f00&amp;t2_uin=1007938230&amp;t2_tid=12&amp;subdotype=55702&amp;signin=0&amp;sceneid=100" data-charset="utf-8" data-tuin="" data-config="1|1|1|0|1,taotaoact.qzone.qq.com,@InputReply|1,taotaoact.qzone.qq.com,@ClickReply|1,taotaoact.qzone.qq.com,commentPresentClick"><b class="hide-clip">回复</b></a></div></div></div></div></li><li class="comments-item bor3" data-type="commentroot" data-tid="13" data-uin="863144225" data-nick="放空" data-who="1"><div class="comments-item-bd"><div class="single-reply"><div class="ui-avatar"><a href="http://user.qzone.qq.com/863144225" target="_blank"><img class="q_namecard" link="nameCard_863144225 des_863144225" alt="放空" src="http://qlogo2.store.qq.com/qzone/863144225/863144225/30"></a></div><div class="comments-content"><a class="nickname name c_tx q_namecard" link="nameCard_863144225" target="_blank" href="http://user.qzone.qq.com/863144225">放空</a>&nbsp; : 瘦了耶<div class="comments-op"><span class=" ui-mr10 state"> 22:26</span><a class="act-reply none" href="javascript:;" data-cmd="qz_reply" data-version="6.4" data-action="http://taotao.qq.com/cgi-bin/emotion_cgi_re_feeds" data-param="t1_source=&amp;t1_uin=994199947&amp;t1_tid=8b49423b4cabc359f53f0f00&amp;t2_uin=863144225&amp;t2_tid=13&amp;subdotype=55702&amp;signin=0&amp;sceneid=100" data-charset="utf-8" data-tuin="" data-config="1|1|1|0|1,taotaoact.qzone.qq.com,@InputReply|1,taotaoact.qzone.qq.com,@ClickReply|1,taotaoact.qzone.qq.com,commentPresentClick"><b class="hide-clip">回复</b></a></div></div></div></div></li><li class="comments-item bor3" data-type="commentroot" data-tid="14" data-uin="1403663865" data-nick="重庆老渣男" data-who="1"><div class="comments-item-bd"><div class="single-reply"><div class="ui-avatar"><a href="http://user.qzone.qq.com/1403663865" target="_blank"><img class="q_namecard" link="nameCard_1403663865 des_1403663865" alt="重庆老渣男" src="http://qlogo2.store.qq.com/qzone/1403663865/1403663865/30"></a></div><div class="comments-content"><a class="nickname name c_tx q_namecard" link="nameCard_1403663865" target="_blank" href="http://user.qzone.qq.com/1403663865">重庆老渣男</a>&nbsp; : 男神<img src="http://qzonestyle.gtimg.cn/qzone/em/e102.png" title=""><div class="comments-op"><span class=" ui-mr10 state"> 22:27</span><a class="act-reply none" href="javascript:;" data-cmd="qz_reply" data-version="6.4" data-action="http://taotao.qq.com/cgi-bin/emotion_cgi_re_feeds" data-param="t1_source=&amp;t1_uin=994199947&amp;t1_tid=8b49423b4cabc359f53f0f00&amp;t2_uin=1403663865&amp;t2_tid=14&amp;subdotype=55702&amp;signin=0&amp;sceneid=100" data-charset="utf-8" data-tuin="" data-config="1|1|1|0|1,taotaoact.qzone.qq.com,@InputReply|1,taotaoact.qzone.qq.com,@ClickReply|1,taotaoact.qzone.qq.com,commentPresentClick"><b class="hide-clip">回复</b></a></div></div></div></div></li></ul></div><div class="mod-commnets-poster feedClickCmd comment_default_inputentry" data-cmd="qz_reply" data-version="6" data-action="http://taotao.qq.com/cgi-bin/emotion_cgi_re_feeds" data-param="t1_source=&amp;t1_uin=994199947&amp;t1_tid=8b49423b4cabc359f53f0f00&amp;signin=0&amp;sceneid=100" data-charset="utf-8" data-maxlength="" data-tuin="994199947" data-config="1|1|1|1,b52,with_fwd,同时转发;0|1,taotaoact.qzone.qq.com,@InputReply|1,taotaoact.qzone.qq.com,@ClickReply|1,taotaoact.qzone.qq.com,commentPresentClick" data-tid=""><div class="comments-poster-bd comments-poster-default"><div class="comments-box" data-clicklog="comment"><div class="textinput textinput-default bor2" contenteditable="true" alt="replybtn" placeholder="评论"><a class="c_tx3" href="javascript:void(0);" alt="replybtn">评论</a></div><div class="mod-insert-img"><a href="javascript:;" data-cmd="qz_quick_upload_img" class="btn-insert-img bg"><i class="icon-camera-16"></i></a></div></div></div></div></div>    </div></li><li class="f-single f-s-s f-single-biz" id="fct_20050606_5000_1_1506003254_0_1"><div class="f-single-top"><span>广告</span><a href="javascript:;" class="arrow-down adFeedsItem" data-cmd="qz_opr_more"><i class="fui-icon icon-arrow-down"></i></a><div class="qz-bubble" style="display: none;">            <div class="bubble-i op-list">                <a href="javascript:;" class="item hot-close"><i class="fui-icon icon-hide"></i>关闭</a>            </div>            <b class="bubble-trig ui-trigbox ui-trigbox-t">                <b class="ui-trig"></b>                <b class="ui-trig ui-trig-up"></b>            </b>        </div></div><div class="f-single-content f-wrap"><div class="f-info">      <a class="report-qboss" href="http://rc.qzone.qq.com/1105370739?via=QZ.SEVENTH.FEED&amp;app_custom=essence " target="_blank" data-traceinfo="2386_190423_578_1_0_0">美人姿态，3D视角，领悟不可描述的风采~点击进入&gt;&gt;&gt;</a>    </div>    <div class="f-ct">            <div class="fui-txtimg">                <div class="img-box"><a class="report-qboss" href="http://rc.qzone.qq.com/1105370739?via=QZ.SEVENTH.FEED&amp;app_custom=essence " target="_blank" data-traceinfo="2386_190423_578_1_0_0"><img src="http://qzonestyle.gtimg.cn/qzone/space_item/boss_pic/2386_2017_9/1505899715801_780154.jpg" style="width:560px"></a>                </div>                <div class="detail-box clearfix">                    <p class="state tip">                    <b style="color: #5d7895;">1901680</b>个网友在玩</p>                    <div class="fixed-right">                        <a class="fixed-btn" href="http://rc.qzone.qq.com/1105370739?via=QZ.SEVENTH.FEED&amp;app_custom=essence " target="_blank" data-traceinfo="2386_190423_578_1_0_0">开始游戏                        </a>                    </div>                </div>            </div>        </div></div></li></ul></li><li class="feed_holder_bottom" style="height:0"></li></ul>
                                    <div style="visibility: hidden;" id="feed_friend_tips" class="f-more-label"></div>
                                    <img style="width:1px;height:1px;" src="about:blank" onerror="window.FIRST_PAGE_FEEDS_TIME=new Date();">
                                    <script type="text/javascript">
                                        window.FIRST_PAGE_FEEDS_TIME2=new Date();
                                        setTimeout(function(){window.FIRST_PAGE_FEEDS_TIME3=new Date();}, 0);
                                    </script>
                                </div>
                            </div>
                        </div>
                    </div>

                    <div id="feed_me" class="none">
                        <div class="fn-feed-control-v2 clearfix">
                            <div class="control-inner">
                                <div class="feed-control-tab" id="me_feed_control">
                                    <a title="查看我的空间动态" class="item-on item-on-slt" id="feed_tab_my" href="javascript:;">我的空间动态<span id="tab_menu_my_cnt" class="sn-count none">0</span><i class="dot dot-tl"></i><i class="dot dot-tr"></i><i class="dot dot-bl"></i><i class="dot dot-br"></i></a>
                                </div>
                                <div class="feed-control-tab-op">
                                    <label class="feed-choose-item" for="feed_me_filter_friend"><i class="ui-icon qz-checkbox-on"></i><input id="feed_me_filter_friend" class="none" type="checkbox" checked="checked">显示非好友评论</label>
                                    <a id="feed_me_refresh" title="刷新动态" href="javascript:;" class="item-left"><i class="ui-icon  icon-refresh-v9"></i></a>
                                </div>
                            </div>
                        </div>
                        <div class="fn-feed-container">
                            <div class="feed  feed-v9">
                                <div class="feed_inner"><div style="display:none;" id="feed_me_update" class="tipsbox bg2"></div><ul id="feed_me_list"></ul><div id="feed_me_tips" style="visibility: visible;" class="f-more-label"><div class="update-more check-more"><p class="b-line"><span class="state">数据加载中</span><img class="hot-loading" src="//qzonestyle.gtimg.cn/aoi/img/icenter/loading.gif"></p></div></div></div>
                            </div>
                        </div>
                    </div>


                    <div id="feed_mv" class="none">
                        <div class="fn-feed-control-v2 clearfix">
                            <div class="control-inner">
                                <div class="feed-control-tab" id="mv_feed_control">
                                    <a title="查看直播广场" class="item-on item-on-slt" id="feed_tab_mv" href="javascript:;">直播广场<i class="dot dot-tl"></i><i class="dot dot-tr"></i><i class="dot dot-bl"></i><i class="dot dot-br"></i></a>
                                </div>
                                <div class="feed-control-tab-op">
                                    <a id="feed_mv_refresh" title="刷新动态" href="javascript:;" class="item-left"><i class="ui-icon  icon-refresh-v9"></i></a>
                                </div>
                            </div>
                        </div>
                        <div class="fn-feed-content">
                            <div class="fn-feed-slide">
                                <div class="slide-wrapper">
                                    <div class="slide-content">
                                        <ul style="position: relative;width: 766px;height: 252px;">
                                            <!-- slide content -->
                                        </ul>
                                    </div>
                                   <div class="slide-btn" style="cursor:pointer">
                                        <a class="prev" href="javascript:void(0);">
                                            <i class="ui-icon icon-prev"></i>
                                        </a>
                                        <a class="next" href="javascript:void(0);">
                                            <i class="ui-icon icon-next"></i>
                                        </a>
                                        <div class="hd">
                                            <ul>
                                                <!-- slide btn -->
                                            </ul>
                                        </div>
                                    </div>
                                </div>
                            </div>
                            <div class="fn-feed-title clearfix">
                                <div class="feed-control-tab">
                                    <span><!-- 直播广场 --></span></div>
                                <div class="feed-control-tab-op">
                                    <label class="feed-live-star">
                                        <i class="ui-icon icon-star"></i>
                                        <span class="j-balance"></span>
                                        <a href="//h5.qzone.qq.com/vip/paycoin" target="_blank"> 充值</a>
                                    </label>
                                </div>
                            </div>
                            <div class="fn-feed-list">
                                <ul class="item-list clearfix" id="feed_mv_list"></ul>
                            </div>
                        </div>
                    </div>


                    <div id="feed_sdc" class="none">
                        <div class="fn-feed-control-v2 clearfix">
                            <div class="control-inner">
                                <div class="feed-control-tab">
                                    <a id="yearAll" href="javascript:;" class="item-on item-on-slt" data-year="2016">全部</a>
                                    <span class="line"></span>
                                    <a id="year2016" href="javascript:;" class="item-on">2016</a>
                                    <span class="line"></span>
                                    <a id="year2015" href="javascript:;" class="item-on">2015</a>
                                    <span class="line"></span>
                                    <a id="year2014" href="javascript:;" class="item-on">2014</a>
                                </div>
                                <div class="feed-control-tab-op">
                                    <a href="javascript:;" title="刷新动态" class="item-left"><i class="ui-icon  icon-refresh-v9"></i></a>
                                    <span class="line"></span>
                                    <a href="javascript:;" title="动态设置" class="item-right"><i class="ui-icon  icon-set-v9"></i></a>
                                </div>
                            </div>
                        </div>
                        <div class="fn-feed-container beforeyear-pop">
                            <div class="feed  feed-v9">
                                <div class="feed_inner">
                                    <ul id="feed_sdc_list_2016" class="feed_sdc_list"></ul>
                                    <ul id="feed_sdc_list_2015" class="feed_sdc_list"></ul>
                                    <ul id="feed_sdc_list_2014" class="feed_sdc_list"></ul>
                                    <div id="feed_sdc_tips" style="visibility:hidden;" class="f-more-label"></div>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>

                <div class="col-main-sidebar">

                    <div id="QM_Container_31">
                        <div id="QM_100004_Title" class="hd none"></div>
                        <div id="QM_100004_Body" class="bd">
                            <div class="fn-checkin-btn">
    <div class="checkin-btn m9d21">
        <div class="ck-time">
            <span class="ck-week">周四</span>
            <span class="ck-day">09.21</span>
        </div>
        <a href="javascript:void(0);" id="checkin_button" onclick="QM&amp;&amp;QM.checkin&amp;&amp;QM.checkin.clickCheckinButton();return false;" class="ck-btn">签到<span class="ck-clip"></span></a>
        <div class="ck-count">
            <span class="ck-count-bg icenter-font">o</span>
            <span class="ck-count-num">113</span>
            <span class="ck-count-word">Days</span>
        </div>
        <span class="ck-flag c_tx4"></span>
        <div class="ck-new-flag"></div>
    </div>
</div>


                        </div>
                        <div id="QM_100004_Bottom" class="ft none"></div>
                    </div>


                        <div id="QM_Container_100005" class="icenter-right-mod icenter-right-ad">
                            <div id="QM_100005_Title" class="hd none">
                                <h3>精彩生活</h3>
                            </div>
                            <div id="QM_100005_Body" class="bd" style="height:120px;">


<div class="gdtads_box">
<div class="gdtads_img" id="gdtLifeCnt">
<a href="http://c.gdt.qq.com/gdt_click.fcg?viewid=Hf0Sijq0KeqxRX8ziIARF2xQuatLW5CLBg1SxIVnfLI0yrH3cSgojotL1NJvRy9fnXu!vhs5mXxuInk8luXetPt3Ld5qJRGB2tpmmqFlSoQBSyQr8m3KdWwEhFUJVFbwZs1RXkHPMWWORHkxHBPArw&amp;jtype=0&amp;i=1&amp;os=0" target="_blank" id="gdtLifelink" class="gdtads_img_link" gdtoriurl="http://c.gdt.qq.com/gdt_click.fcg?viewid=Hf0Sijq0KeqxRX8ziIARF2xQuatLW5CLBg1SxIVnfLI0yrH3cSgojotL1NJvRy9fnXu!vhs5mXxuInk8luXetPt3Ld5qJRGB2tpmmqFlSoQBSyQr8m3KdWwEhFUJVFbwZs1RXkHPMWWORHkxHBPArw&amp;jtype=0&amp;i=1&amp;os=0"><img src="http://pgdt.gtimg.cn/gdt/0/DAAQNXOADcAB4AASBZGnWDBHbdIiDh.jpg/0?ck=a5c005bfeda650e593cdb817e7fb7acb" id="gdtLifeImg" alt=""></a>
<div class="change_box _js_next" style="" id="gdtLifeNext">
<a href="javascript:;" onclick="return false;"><i class="icon_change"></i></a>
</div>
<div class="gdtads_loading" id="gdtLifeLoading" style="display: none">Loading...</div>
<img src="http://v.gdt.qq.com/gdt_stats.fcg?viewid=Hf0Sijq0KeqxRX8ziIARF2xQuatLW5CLBg1SxIVnfLI0yrH3cSgojotL1NJvRy9fnXu!vhs5mXxuInk8luXetPt3Ld5qJRGB2tpmmqFlSoQBSyQr8m3KdWwEhFUJVFbwZs1RXkHPMWWORHkxHBPArw&amp;i=1&amp;os=0" style="width:1px;height: 1px; visibility: hidden;">
<input type="hidden" id="gdtLifealist" value="">
<a href="http://e.qq.com/?from=qzone_pc" class="j-corner" target="_blank" style="bottom: 0px;"><div class="tr"></div>
<div class="info">
<p class="text1"></p>
<p class="text2"></p>
</div></a></div>
</div>


<input type="hidden" id="gdtLifeHtml" value="1">
<input type="hidden" id="gdtLifecn" value="">
<input type="hidden" id="gdtLifeexptime" value="">
<input type="hidden" id="gdtLifeordlist" value="">

<!--[if !IE]>|xGv00|a99ae7a130aa7e7004caa47c0fffc884<![endif]-->

                            </div>
                            <div id="QM_100005_Bottom" class="ft none">

                            </div>
                        </div>


                    <div id="QM_Container_100002" class="icenter-right-mod">
                        <div id="QM_100002_Title" class="hd">
                            <h3>大家都在看</h3>
                        </div>
                        <div id="QM_100002_Body" class="bd">
                            <div class="fn_hotTopic_v5">
                                <ol>
                                    <li class="topic_item" key="https://fm.qq.com/show/rd001qLxxM3ChgAU" feedkey="" data-url="https://fm.qq.com/show/rd001qLxxM3ChgAU?referer=hot_topic" data-actmore="url|" qbosskey="2129_190287_563_1_100453_91183" exposure="1"><a class="topic_avatar" target="_blank" href="https://fm.qq.com/show/rd001qLxxM3ChgAU?referer=hot_topic"><img src="http://qzonestyle.gtimg.cn/qzone/space_item/boss_pic/2129_2017_9/1505830430029_400674.jpg" alt=""></a><div class="topic_main"><div class="topic_title"><a target="_blank" class="c_tx2 title_txt topic_text" href="https://fm.qq.com/show/rd001qLxxM3ChgAU?referer=hot_topic">范冰冰终于被求婚，其实她也只是小女人</a></div><div class="topic_op"><a href="javascript:;" class="topic_op_link"><i class="ui_icon icon_hottopic_prefer"></i><span class="act_text c_tx">赞</span></a> <span class="topic_op_num bor2"><span class="act_count">499</span><b class="ui_trigbox ui_trigbox_l"><b class="ui_trig bor2"></b><b class="ui_trig ui_trig_up bor_bg"></b></b></span></div></div></li><li class="topic_item" key="http://rc.qzone.qq.com/1105731023?via=QZ.ISWATCHING&amp;app_custom=essence" feedkey="" data-url="http://rc.qzone.qq.com/1105731023?via=QZ.ISWATCHING&amp;app_custom=essence&amp;referer=hot_topic" qbosskey="2129_190359_439_1_100458_92276" exposure="1"><a class="topic_avatar" target="_blank" href="http://rc.qzone.qq.com/1105731023?via=QZ.ISWATCHING&amp;app_custom=essence&amp;referer=hot_topic"><img src="http://qzonestyle.gtimg.cn/qzone/space_item/boss_pic/2129_2017_8/1503050472656_974149.jpg" alt=""></a><div class="topic_main"><div class="topic_title"><a target="_blank" class="c_tx2 title_txt topic_text" href="http://rc.qzone.qq.com/1105731023?via=QZ.ISWATCHING&amp;app_custom=essence&amp;referer=hot_topic">姜子牙封神为何不封自己，原因竟是这样</a></div><div class="topic_op"><a href="javascript:;" class="topic_op_link"><i class="ui_icon icon_hottopic_praise"></i><span class="act_text c_tx">赞</span></a> <span class="topic_op_num bor2"><span class="act_count">621</span><b class="ui_trigbox ui_trigbox_l"><b class="ui_trig bor2"></b><b class="ui_trig ui_trig_up bor_bg"></b></b></span></div></div></li>



<li class="topic_item" data-gdt="http://analy.qq.com/cgi-bin/apptrace?appid_via=1105876681_1A000101310DD4CF&amp;platform=1&amp;action=200" data-url="http://c.gdt.qq.com/gdt_click.fcg?viewid=Xlv_MisZDPyxRX8ziIARF83ioVQ7qfAV47JJOVn!ihtIl4UR!oUbyYtL1NJvRy9fnXu!vhs5mXxuInk8luXetPt3Ld5qJRGB2tpmmqFlSoQBSyQr8m3KddWPRl9p_ETiU4P1ExPeHhTKkUMkGXU5gA&amp;jtype=0&amp;via=1A000101310DD4CF&amp;i=1&amp;os=0">
<a class="topic_avatar" target="_blank" href="http://c.gdt.qq.com/gdt_click.fcg?viewid=Xlv_MisZDPyxRX8ziIARF83ioVQ7qfAV47JJOVn!ihtIl4UR!oUbyYtL1NJvRy9fnXu!vhs5mXxuInk8luXetPt3Ld5qJRGB2tpmmqFlSoQBSyQr8m3KddWPRl9p_ETiU4P1ExPeHhTKkUMkGXU5gA&amp;jtype=0&amp;via=1A000101310DD4CF&amp;i=1&amp;os=0" gdtoriurl="http://c.gdt.qq.com/gdt_click.fcg?viewid=Xlv_MisZDPyxRX8ziIARF83ioVQ7qfAV47JJOVn!ihtIl4UR!oUbyYtL1NJvRy9fnXu!vhs5mXxuInk8luXetPt3Ld5qJRGB2tpmmqFlSoQBSyQr8m3KddWPRl9p_ETiU4P1ExPeHhTKkUMkGXU5gA&amp;jtype=0&amp;via=1A000101310DD4CF&amp;i=1&amp;os=0">
<img src="http://pgdt.gtimg.cn/gdt/0/DAAEQ_AABQAA8AAJBYExKcDXiNULrA.jpg/0?ck=ed065b744e92b6c444c10b33b99505f4" alt="">
<span style="position: absolute; z-index: 11; bottom: 3px; left: 3px; height: 13px; width: 24px; background: url(&quot;http://qzonestyle.gtimg.cn/open_proj/proj-gdt-toufang/ggExternal/img/sign-ad-10.png&quot;);"></span></a>
<div class="topic_main">
<div class="topic_title"><a target="_blank" class="c_tx2 title_txt topic_text" href="http://c.gdt.qq.com/gdt_click.fcg?viewid=Xlv_MisZDPyxRX8ziIARF83ioVQ7qfAV47JJOVn!ihtIl4UR!oUbyYtL1NJvRy9fnXu!vhs5mXxuInk8luXetPt3Ld5qJRGB2tpmmqFlSoQBSyQr8m3KddWPRl9p_ETiU4P1ExPeHhTKkUMkGXU5gA&amp;jtype=0&amp;via=1A000101310DD4CF&amp;i=1&amp;os=0" gdtoriurl="http://c.gdt.qq.com/gdt_click.fcg?viewid=Xlv_MisZDPyxRX8ziIARF83ioVQ7qfAV47JJOVn!ihtIl4UR!oUbyYtL1NJvRy9fnXu!vhs5mXxuInk8luXetPt3Ld5qJRGB2tpmmqFlSoQBSyQr8m3KddWPRl9p_ETiU4P1ExPeHhTKkUMkGXU5gA&amp;jtype=0&amp;via=1A000101310DD4CF&amp;i=1&amp;os=0">唐僧的九个前世为什么都被沙僧吃了？</a>
</div>
<div class="topic_op">
<a href="http://c.gdt.qq.com/gdt_click.fcg?viewid=Xlv_MisZDPyxRX8ziIARF83ioVQ7qfAV47JJOVn!ihtIl4UR!oUbyYtL1NJvRy9fnXu!vhs5mXxuInk8luXetPt3Ld5qJRGB2tpmmqFlSoQBSyQr8m3KddWPRl9p_ETiU4P1ExPeHhTKkUMkGXU5gA&amp;jtype=0&amp;via=1A000101310DD4CF&amp;i=1&amp;os=0" target="_blank" class="topic_op_link" gdtoriurl="http://c.gdt.qq.com/gdt_click.fcg?viewid=Xlv_MisZDPyxRX8ziIARF83ioVQ7qfAV47JJOVn!ihtIl4UR!oUbyYtL1NJvRy9fnXu!vhs5mXxuInk8luXetPt3Ld5qJRGB2tpmmqFlSoQBSyQr8m3KddWPRl9p_ETiU4P1ExPeHhTKkUMkGXU5gA&amp;jtype=0&amp;via=1A000101310DD4CF&amp;i=1&amp;os=0"><i class="ui_icon icon_hottopic_prefer"></i><span class="act_text c_tx">去看看</span></a>
<span class="topic_op_num bor2">
<span class="act_count">50875</span>
<b class="ui_trigbox ui_trigbox_l"><b class="ui_trig bor2"></b><b class="ui_trig ui_trig_up bor_bg"></b></b>
</span>
</div>
</div>
<img src="http://v.gdt.qq.com/gdt_stats.fcg?viewid=Xlv_MisZDPyxRX8ziIARF83ioVQ7qfAV47JJOVn!ihtIl4UR!oUbyYtL1NJvRy9fnXu!vhs5mXxuInk8luXetPt3Ld5qJRGB2tpmmqFlSoQBSyQr8m3KddWPRl9p_ETiU4P1ExPeHhTKkUMkGXU5gA&amp;i=1&amp;os=0" width="1" height="1" style="visibility: hidden;">

<img src="http://analy.qq.com/cgi-bin/apptrace?appid_via=1105876681_1A000101310DD4CF&amp;platform=1&amp;action=100" width="1" height="1" style="visibility: hidden;">

</li>

<!--[if !IE]>|xGv00|3f78282c7313db1fc5aad9ed0a5be80f<![endif]-->


                                </ol>
                            </div>
                        </div>
                        <div id="QM_100002_Bottom" class="ft">

                            <a href="javascript:;" class="p_prev prev-page unclick" title="上一页"><i></i></a>
                            <a href="javascript:;" class="p_next next-page" title="下一页"><i></i></a>
                        </div>
                    </div>


                    <div id="QM_Container_3" class="icenter-right-mod" style="">
                        <div id="QM_3_Title" class="hd"><p><a class="title" href="javascript:">谁看过我</a><b class="divide-line">|</b><a href="javascript:;" id="visitYou">我看过谁</a><b class="divide-line">|</b><a href="javascript:;" id="refuseYou">被挡访客</a></p></div>
                        <div class="fn_accessLog_tips bg2 bor2" style="padding: 8px 24px 8px 8px; width: 178px; line-height: 24px; position: relative; margin-bottom: 8px;"><p>看看大家都看了什么？开通黄钻可查看2000人的访问轨迹。<br><a href="http://pay.qq.com/ipay/index.shtml?c=xxjzgw,xxjzghh&amp;aid=zone.fangke" target="_blank"><img src="http://qzonestyle.gtimg.cn/qzone_v6/img/icenter/btn_qz_open.png" style="vertical-align:top"></a></p><a class="ui_x c_tx2 bg6_hover v_close" href="javascript:;" style="right: 10px;top: 5px;position: absolute;display: block;height: 20px;width: 10px;text-decoration: none;">×</a></div><div id="QM_3_Body" class="bd"><ul id="visitMeContainer" class="three-in-line visitor-list clearfix" style="height: 300px;"><li class="user-item"><div class="qz_del_ul w90" data-uin="573197470"><a class="c_tx3 j-set-record top_del bg6_hover" title="删除" href="javascript:;">×</a><div class="qz_del_ul_option bor2 none"><ul><li class="bg"><a class="bg6_hover j-hide-record" data-click="hide_visitor_record" href="javascript:;">隐藏他的访问</a></li><li class="bg"><a class="bg6_hover j-del-record" href="javascript:;">删除本次记录</a></li></ul></div></div><span class="hide-user-tag none j-hide-from-flag"><i class="ui-icon hide-user-icon" title="该好友来访记录仅你可见"></i><b class="floor"></b></span><a href="http://user.qzone.qq.com/573197470" class="user-avatar q_namecard" data-from="visitor" data-index="1" link="nameCard_573197470" target="_blank" data-click="visitor_head"><img src="http://qlogo3.store.qq.com/qzone/573197470/573197470/100" alt="李甜的头像"></a><div class="user-name-bg"><span class="user-name textoverflow  ">李甜</span><a href="javascript:void(0);" target="_blank" title="QQ空间独立版" class="ui-icon icon-qzphone qz-app-flag"></a></div> <span class="date"> 21:21</span></li><li class="user-item"><div class="qz_del_ul w90" data-uin="654858397"><a class="c_tx3 j-set-record top_del bg6_hover" title="删除" href="javascript:;">×</a><div class="qz_del_ul_option bor2 none"><ul><li class="bg"><a class="bg6_hover j-hide-record" data-click="hide_visitor_record" href="javascript:;">隐藏他的访问</a></li><li class="bg"><a class="bg6_hover j-del-record" href="javascript:;">删除本次记录</a></li></ul></div></div><span class="hide-user-tag none j-hide-from-flag"><i class="ui-icon hide-user-icon" title="该好友来访记录仅你可见"></i><b class="floor"></b></span><a href="http://user.qzone.qq.com/654858397" class="user-avatar q_namecard" data-from="visitor" data-index="2" link="nameCard_654858397" target="_blank" data-click="visitor_head"><img src="http://qlogo2.store.qq.com/qzone/654858397/654858397/100" alt="谭超的头像"></a><div class="user-name-bg"><span class="user-name textoverflow user-name-full ">谭超</span></div> <span class="date"> 21:19</span></li><li class="user-item"><div class="qz_del_ul w90" data-uin="769980109"><a class="c_tx3 j-set-record top_del bg6_hover" title="删除" href="javascript:;">×</a><div class="qz_del_ul_option bor2 none"><ul><li class="bg"><a class="bg6_hover j-hide-record" data-click="hide_visitor_record" href="javascript:;">隐藏他的访问</a></li><li class="bg"><a class="bg6_hover j-del-record" href="javascript:;">删除本次记录</a></li></ul></div></div><span class="hide-user-tag none j-hide-from-flag"><i class="ui-icon hide-user-icon" title="该好友来访记录仅你可见"></i><b class="floor"></b></span><a href="http://user.qzone.qq.com/769980109" class="user-avatar q_namecard" data-from="visitor" data-index="3" link="nameCard_769980109" target="_blank" data-click="visitor_head"><img src="http://qlogo2.store.qq.com/qzone/769980109/769980109/100" alt="肖亚的头像"></a><div class="user-name-bg"><span class="user-name textoverflow  ">肖亚</span><a href="javascript:void(0);" target="_blank" title="QQ空间独立版" class="ui-icon icon-qzphone qz-app-flag"></a></div> <span class="date"> 20:41</span></li><li class="user-item"><div class="qz_del_ul w90" data-uin="469396333"><a class="c_tx3 j-set-record top_del bg6_hover" title="删除" href="javascript:;">×</a><div class="qz_del_ul_option bor2 none"><ul><li class="bg"><a class="bg6_hover j-hide-record" data-click="hide_visitor_record" href="javascript:;">隐藏他的访问</a></li><li class="bg"><a class="bg6_hover j-del-record" href="javascript:;">删除本次记录</a></li></ul></div></div><span class="hide-user-tag none j-hide-from-flag"><i class="ui-icon hide-user-icon" title="该好友来访记录仅你可见"></i><b class="floor"></b></span><a href="http://user.qzone.qq.com/469396333" class="user-avatar q_namecard" data-from="visitor" data-index="4" link="nameCard_469396333" target="_blank" data-click="visitor_head"><img src="http://qlogo2.store.qq.com/qzone/469396333/469396333/100" alt="曹玲玲的头像"></a><div class="user-name-bg"><span class="user-name textoverflow  ">曹玲玲</span><a href="javascript:void(0);" target="_blank" title="QQ空间独立版" class="ui-icon icon-qzphone qz-app-flag"></a></div> <span class="date"> 18:10</span></li><li class="user-item"><div class="qz_del_ul w90" data-uin="516211120"><a class="c_tx3 j-set-record top_del bg6_hover" title="删除" href="javascript:;">×</a><div class="qz_del_ul_option bor2 none"><ul><li class="bg"><a class="bg6_hover j-hide-record" data-click="hide_visitor_record" href="javascript:;">隐藏他的访问</a></li><li class="bg"><a class="bg6_hover j-del-record" href="javascript:;">删除本次记录</a></li></ul></div></div><span class="hide-user-tag none j-hide-from-flag"><i class="ui-icon hide-user-icon" title="该好友来访记录仅你可见"></i><b class="floor"></b></span><a href="http://user.qzone.qq.com/516211120" class="user-avatar q_namecard" data-from="visitor" data-index="5" link="nameCard_516211120" target="_blank" data-click="visitor_head"><img src="http://qlogo1.store.qq.com/qzone/516211120/516211120/100" alt="杨强的头像"></a><div class="user-name-bg"><span class="user-name textoverflow user-name-full ">杨强</span></div> <span class="date"> 16:03</span></li><li class="user-item"><div class="qz_del_ul w90" data-uin="382847106"><a class="c_tx3 j-set-record top_del bg6_hover" title="删除" href="javascript:;">×</a><div class="qz_del_ul_option bor2 none"><ul><li class="bg"><a class="bg6_hover j-hide-record" data-click="hide_visitor_record" href="javascript:;">隐藏他的访问</a></li><li class="bg"><a class="bg6_hover j-del-record" href="javascript:;">删除本次记录</a></li></ul></div></div><span class="hide-user-tag none j-hide-from-flag"><i class="ui-icon hide-user-icon" title="该好友来访记录仅你可见"></i><b class="floor"></b></span><a href="http://user.qzone.qq.com/382847106" class="user-avatar q_namecard" data-from="visitor" data-index="6" link="nameCard_382847106" target="_blank" data-click="visitor_head"><img src="http://qlogo3.store.qq.com/qzone/382847106/382847106/100" alt="房亮的头像"></a><div class="user-name-bg"><span class="user-name textoverflow  ">房亮</span><a href="http://pay.qq.com/ipay/index.shtml?n=3&amp;c=xxjzghh,xxjzgw&amp;aid=mingpian_icon&amp;ch=qdqb,kj" target="_blank" title="点击查看黄钻特权详情" class="qz_vip_icon_s qz_vip_icon_s_6 "></a></div> <span class="date"> 15:34</span></li><li class="user-item"><div class="qz_del_ul w90" data-uin="747178206"><a class="c_tx3 j-set-record top_del bg6_hover" title="删除" href="javascript:;">×</a><div class="qz_del_ul_option bor2 none"><ul><li class="bg"><a class="bg6_hover j-hide-record" data-click="hide_visitor_record" href="javascript:;">隐藏他的访问</a></li><li class="bg"><a class="bg6_hover j-del-record" href="javascript:;">删除本次记录</a></li></ul></div></div><span class="hide-user-tag none j-hide-from-flag"><i class="ui-icon hide-user-icon" title="该好友来访记录仅你可见"></i><b class="floor"></b></span><a href="http://user.qzone.qq.com/747178206" class="user-avatar q_namecard" data-from="visitor" data-index="7" link="nameCard_747178206" target="_blank" data-click="visitor_head"><img src="http://qlogo3.store.qq.com/qzone/747178206/747178206/100" alt="赖雪的头像"></a><div class="user-name-bg"><span class="user-name textoverflow  ">赖雪</span><a href="javascript:void(0);" target="_blank" title="QQ空间独立版" class="ui-icon icon-qzphone qz-app-flag"></a></div> <span class="date"> 14:27</span></li><li class="user-item"><div class="qz_del_ul w90" data-uin="569453713"><a class="c_tx3 j-set-record top_del bg6_hover" title="删除" href="javascript:;">×</a><div class="qz_del_ul_option bor2 none"><ul><li class="bg"><a class="bg6_hover j-hide-record" data-click="hide_visitor_record" href="javascript:;">隐藏他的访问</a></li><li class="bg"><a class="bg6_hover j-del-record" href="javascript:;">删除本次记录</a></li></ul></div></div><span class="hide-user-tag none j-hide-from-flag"><i class="ui-icon hide-user-icon" title="该好友来访记录仅你可见"></i><b class="floor"></b></span><a href="http://user.qzone.qq.com/569453713" class="user-avatar q_namecard" data-from="visitor" data-index="8" link="nameCard_569453713" target="_blank" data-click="visitor_head"><img src="http://qlogo2.store.qq.com/qzone/569453713/569453713/100" alt="疯子的头像"></a><div class="user-name-bg"><span class="user-name textoverflow  ">疯子</span><a href="javascript:void(0);" target="_blank" title="QQ空间独立版" class="ui-icon icon-qzphone qz-app-flag"></a></div> <span class="date"> 10:19</span></li><li class="user-item"><div class="qz_del_ul w90" data-uin="454782407"><a class="c_tx3 j-set-record top_del bg6_hover" title="删除" href="javascript:;">×</a><div class="qz_del_ul_option bor2 none"><ul><li class="bg"><a class="bg6_hover j-hide-record" data-click="hide_visitor_record" href="javascript:;">隐藏他的访问</a></li><li class="bg"><a class="bg6_hover j-del-record" href="javascript:;">删除本次记录</a></li></ul></div></div><span class="hide-user-tag none j-hide-from-flag"><i class="ui-icon hide-user-icon" title="该好友来访记录仅你可见"></i><b class="floor"></b></span><a href="http://user.qzone.qq.com/454782407" class="user-avatar q_namecard" data-from="visitor" data-index="9" link="nameCard_454782407" target="_blank" data-click="visitor_head"><img src="http://qlogo4.store.qq.com/qzone/454782407/454782407/100" alt="天之娇的头像"></a><div class="user-name-bg"><span class="user-name textoverflow  ">天之娇</span><a href="javascript:void(0);" target="_blank" title="QQ空间独立版" class="ui-icon icon-qzphone qz-app-flag"></a></div> <span class="date"> 10:05</span></li></ul><ul id="visitYouContainer" class="three-in-line visitor-list clearfix none"></ul><ul id="refuseYouContainer" class="three-in-line visitor-list clearfix none"></ul></div>
                        <div id="QM_3_Bottom" class="ft">
							<div id="QM_3_pager"><a href="javascript:" class="prev-page unclick"><i></i></a><a href="javascript:;" class="next-page" onclick="QM.Visitor.nextPage();return false;" title="下一页"><i></i></a><a href="javascript:;" onclick="QZONE.FrontPage.toApp('/friends/visitor/');return false;" data-click="detail" class="more" id="visitor_detail" title="更多"><i></i></a></div>
							<div id="QM_3_info"><ul class="visitor-count  clearfix"><li><p><span>今日浏览</span> </p><a href="javascript:;" title="今天浏览量：16" onclick="QM.Visitor.clickStat(2);return false;" id="QM_3_visit_today">16</a></li><li class="gap"><p><span>总浏览</span> </p><a href="javascript:;" title="总浏览量：11120" onclick="QM.Visitor.clickStat(1);return false;" id="QM_3_visit_total">11120</a><span class="tips-num none" title="新增访客">0</span></li><li class="gap " id="today_refuse_visitor"><p><span>被挡访客</span></p><a href="javascript:;" title="被挡访问量：402">402</a><span class="tips-num none" title="新增访客">0</span></li></ul><p id="audit_cnt" class="pt_5" style="display:none;"></p></div>
                        </div>
                    </div>


                    <div id="QM_Container_333" class="icenter-right-mod">
                        <div id="QM_333_Title" class="hd">
                            <h3>礼物</h3>
                        </div>
                        <div id="QM_333_Body" class="bd">
                            <div class="bd" id="qm_333_gift_list">
	<div class="gift-bd">
		<ul class="send-gift-list">

				<li ishighactive="0" isfestival="0" isvisitor="0" isbirth="1" uin="308136086" nick="李莎莎" source="0" year="1993" month="9" day="21" lunarflag="0" is_super="0" is_year="0" vip_level="0" gender="2" age="24" relate_days="0" btntext="0">
					<div class="user-info">

						<a href="http://user.qzone.qq.com/308136086" class="user-avatar" target="_blank"><img src="http://qlogo2.store.qq.com/qzone/308136086/308136086/50"></a>

				        <div class="info-details">
				            <p>
				            	<a href="http://user.qzone.qq.com/308136086" title="李莎莎" target="_blank" class="user-name textoverflow">李莎莎</a>

				            </p>
				            <p>

									<span class="user-desc c_tx3">今天生日</span>

				            </p>
				        </div>
				        <!-- 关闭按钮 -->
				        <a class="close-button" title="3天内不再显示此人" uin="308136086" data-hottag="ISD.QZONEGIFT.QZONEINFOCENTER.CENTER-closeuin" href="javascript:;">×</a>
				    </div>


						<a class="button bgr2 c_tx_2">

							<i class="icon-cake-on"></i><b class="c_tx2">生日快乐</b>

						</a>

				</li>

				<li ishighactive="0" isfestival="0" isvisitor="0" isbirth="1" uin="1426700739" nick="陈露" source="0" year="1997" month="9" day="22" lunarflag="0" is_super="0" is_year="0" vip_level="0" gender="2" age="20" birthday="1997-9-22" relate_days="1" btntext="0">
					<div class="user-info">

						<a href="http://user.qzone.qq.com/1426700739" class="user-avatar" target="_blank"><img src="http://qlogo2.store.qq.com/qzone/1426700739/1426700739/50"></a>

				        <div class="info-details">
				            <p>
				            	<a href="http://user.qzone.qq.com/1426700739" title="陈露" target="_blank" class="user-name textoverflow">陈露</a>

				            </p>
				            <p>

									<span class="user-desc c_tx3">明天生日</span>

				            </p>
				        </div>
				        <!-- 关闭按钮 -->
				        <a class="close-button" title="3天内不再显示此人" uin="1426700739" data-hottag="ISD.QZONEGIFT.QZONEINFOCENTER.CENTER-closeuin" href="javascript:;">×</a>
				    </div>


						<a class="button bgr2 c_tx_2">

							<i class="icon-cake"></i><b class="c_tx2">生日快乐</b>

						</a>

				</li>

				<li ishighactive="0" isfestival="0" isvisitor="0" isbirth="1" uin="2235008237" nick="浅熙" source="0" year="1992" month="9" day="22" lunarflag="0" is_super="0" is_year="0" vip_level="0" gender="2" age="25" birthday="1992-9-22" relate_days="1" btntext="0">
					<div class="user-info">

						<a href="http://user.qzone.qq.com/2235008237" class="user-avatar" target="_blank"><img src="http://qlogo2.store.qq.com/qzone/2235008237/2235008237/50"></a>

				        <div class="info-details">
				            <p>
				            	<a href="http://user.qzone.qq.com/2235008237" title="浅熙" target="_blank" class="user-name textoverflow">浅熙</a>

				            </p>
				            <p>

									<span class="user-desc c_tx3">明天生日</span>

				            </p>
				        </div>
				        <!-- 关闭按钮 -->
				        <a class="close-button" title="3天内不再显示此人" uin="2235008237" data-hottag="ISD.QZONEGIFT.QZONEINFOCENTER.CENTER-closeuin" href="javascript:;">×</a>
				    </div>


						<a class="button bgr2 c_tx_2">

							<i class="icon-cake"></i><b class="c_tx2">生日快乐</b>

						</a>

				</li>

		</ul>
	</div>
</div>
                        </div>
                        <div id="QM_333_Bottom" class="ft">
                            <div class="gift-ft">

		<a href="javascript:void(0)" class="prev-page" data-hottag="ISD.QZONEINFOCENTER.CENTER-pre_group">
			<i></i>
		</a>
		<a href="javascript:void(0)" class="next-page" data-hottag="ISD.QZONEINFOCENTER.CENTER-next_group">
			<i></i>
		</a>


	<a href="javascript:void(0)" class="send-all" data-birthnum="12" data-hottag="ISD.QZONEGIFT.QZONEINFOCENTER.CENTER-more_gift">给更多好友送礼</a>
</div>
                        </div>
                    </div>

                    <div id="QM_Container_100003" data-module="friendship" class="icenter-right-mod">
    <div id="QM_100003_Title" class="hd">
        <p>
             <a id="QM_100003_care_me" data-role="button" data-rel="care_me" class="title">谁在意我</a>
             <b class="divide-line">|</b>
             <a id="QM_100003_care_who" data-role="button" data-rel="care_who" class="who-i-visit" href="javascript:void(0);">我在意谁</a>
        </p>
    </div>
    <div id="QM_100003_Body" class="bd">
        <ul id="QM_100003_friendship_care_me" class="friendship-rank-list user-vertical-list">



            <li>
                <div class="user-info">
                    <a class="user-avatar" target="_blank" href="http://user.qzone.qq.com/1017335988">
                        <img src="http://qlogo1.store.qq.com/qzone/1017335988/1017335988/100" alt="">
                        <span class="rank-tag rank-tag-first"> <b class="txt">1</b> <b class="floor"></b>
                        </span>
                    </a>
                    <div class="info-details">
                        <p>
                            <a class="user-name textoverflow" target="_blank" href="http://user.qzone.qq.com/1017335988">PN·二刀肉</a>
                        </p>
                        <p>
                            <a href="http://user.qzone.qq.com/1017335988/friendship?reverse=1&amp;via=ic" class="rank-show"> <i class="ui-icon icon-friendship"></i>
                                92
                            </a>
                        </p>
                    </div>

                </div>
                <a class="button" target="_blank" href="http://user.qzone.qq.com/1017335988/friendship?reverse=1&amp;via=ic">详情</a>
            </li>


            <li>
                <div class="user-info">
                    <a class="user-avatar" target="_blank" href="http://user.qzone.qq.com/1150783525">
                        <img src="http://qlogo2.store.qq.com/qzone/1150783525/1150783525/100" alt="">
                        <span class="rank-tag rank-tag-second"> <b class="txt">2</b> <b class="floor"></b>
                        </span>
                    </a>
                    <div class="info-details">
                        <p>
                            <a class="user-name textoverflow" target="_blank" href="http://user.qzone.qq.com/1150783525">李松</a>
                        </p>
                        <p>
                            <a href="http://user.qzone.qq.com/1150783525/friendship?reverse=1&amp;via=ic" class="rank-show"> <i class="ui-icon icon-friendship"></i>
                                67
                            </a>
                        </p>
                    </div>

                </div>
                <a class="button" target="_blank" href="http://user.qzone.qq.com/1150783525/friendship?reverse=1&amp;via=ic">详情</a>
            </li>


            <li>
                <div class="user-info">
                    <a class="user-avatar" target="_blank" href="http://user.qzone.qq.com/452242253">
                        <img src="http://qlogo2.store.qq.com/qzone/452242253/452242253/100" alt="">
                        <span class="rank-tag rank-tag-third"> <b class="txt">3</b> <b class="floor"></b>
                        </span>
                    </a>
                    <div class="info-details">
                        <p>
                            <a class="user-name textoverflow" target="_blank" href="http://user.qzone.qq.com/452242253">张森</a>
                        </p>
                        <p>
                            <a href="http://user.qzone.qq.com/452242253/friendship?reverse=1&amp;via=ic" class="rank-show"> <i class="ui-icon icon-friendship"></i>
                                58
                            </a>
                        </p>
                    </div>

                </div>
                <a class="button" target="_blank" href="http://user.qzone.qq.com/452242253/friendship?reverse=1&amp;via=ic">详情</a>
            </li>

        </ul>
        <ul id="QM_100003_friendship_care_who" class="friendship-rank-list user-vertical-list none">



            <li>
                <div class="user-info">
                    <a class="user-avatar" target="_blank" href="http://user.qzone.qq.com/1017335988">
                        <img src="http://qlogo1.store.qq.com/qzone/1017335988/1017335988/100" alt="">
                        <span class="rank-tag rank-tag-first"> <b class="txt">1</b> <b class="floor"></b>
                        </span>
                    </a>
                    <div class="info-details">
                        <p>
                            <a class="user-name textoverflow" target="_blank" href="http://user.qzone.qq.com/1017335988">PN·二刀肉</a>
                        </p>
                        <p>
                            <a href="http://user.qzone.qq.com/1017335988/friendship?via=ic" class="rank-show"> <i class="ui-icon icon-friendship"></i>
                                91
                            </a>
                        </p>
                    </div>

                </div>
                <a class="button" target="_blank" href="http://user.qzone.qq.com/1017335988/friendship?via=ic">详情</a>
            </li>


            <li>
                <div class="user-info">
                    <a class="user-avatar" target="_blank" href="http://user.qzone.qq.com/1150783525">
                        <img src="http://qlogo2.store.qq.com/qzone/1150783525/1150783525/100" alt="">
                        <span class="rank-tag rank-tag-second"> <b class="txt">2</b> <b class="floor"></b>
                        </span>
                    </a>
                    <div class="info-details">
                        <p>
                            <a class="user-name textoverflow" target="_blank" href="http://user.qzone.qq.com/1150783525">李松</a>
                        </p>
                        <p>
                            <a href="http://user.qzone.qq.com/1150783525/friendship?via=ic" class="rank-show"> <i class="ui-icon icon-friendship"></i>
                                68
                            </a>
                        </p>
                    </div>

                </div>
                <a class="button" target="_blank" href="http://user.qzone.qq.com/1150783525/friendship?via=ic">详情</a>
            </li>


            <li>
                <div class="user-info">
                    <a class="user-avatar" target="_blank" href="http://user.qzone.qq.com/452242253">
                        <img src="http://qlogo2.store.qq.com/qzone/452242253/452242253/100" alt="">
                        <span class="rank-tag rank-tag-third"> <b class="txt">3</b> <b class="floor"></b>
                        </span>
                    </a>
                    <div class="info-details">
                        <p>
                            <a class="user-name textoverflow" target="_blank" href="http://user.qzone.qq.com/452242253">张森</a>
                        </p>
                        <p>
                            <a href="http://user.qzone.qq.com/452242253/friendship?via=ic" class="rank-show"> <i class="ui-icon icon-friendship"></i>
                                59
                            </a>
                        </p>
                    </div>

                </div>
                <a class="button" target="_blank" href="http://user.qzone.qq.com/452242253/friendship?via=ic">详情</a>
            </li>

        </ul>
    </div>
</div><div id="QM_Container_100003_anchor" data-friend-num="435" class="icenter-right-mod none"></div>



                    <div id="QM_Container_100014" class="icenter-right-mod before-year none"></div>


                <div style="width: 230px;" id="sidebar-ic-fixed"></div><div id="QM_Container_100012" class="collet_box fn_fnrecm icenter-right-mod icenter-right-photo" style="">
                        <div id="QM_100012_Title" class="box_hd"><h3>猜你喜欢</h3></div>
                        <div id="QM_100012_Body" class="box_bd"><div class="qz_recommend_friend_photo qz_recommend_friend_photo_v2" style="width:228px"><div class="friend_photo_border" style="width:228px;"><div class="friend_photo_container" style="display:block;width:228px;height:228px;overflow:hidden;" onmousemove="QPHOTO.mod.friendHotPhoto.onMouseMove(this,event);" onmouseout="QPHOTO.mod.friendHotPhoto.onMouseOut(this,event);"><div id="fhp_photo_loading" class="loading" style="display: none;"><a tcisdkey="loading" onclick="QPHOTO.mod.friendHotPhoto.tcisdStats(this);return QPHOTO.mod.friendHotPhoto.gotoPhoto();"><img src="http://qzonestyle.gtimg.cn/qzone_v6/img/photo/loading32.gif"></a><p class="loading-tips">图片加载中...</p></div><a id="fhp_photo_imgctn" href="javascript:void(0);" tcisdkey="smallphoto" onclick="QPHOTO.mod.friendHotPhoto.tcisdStats(this);return QPHOTO.mod.friendHotPhoto.gotoPhoto();" title="2017-09-17"><img id="js-img-show" style="width: 380px; height: 228px; margin: 0px -76px; opacity: 1;" src="http://b242.photo.store.qq.com/psb?/V12ifKF60nitlQ/3antr9MBf910WELCyvGqhgwntDxgv8aZS*Mhdr4kBQY!/c/dPIAAAAAAAAA&amp;bo=4AEgAQAAAAAAAOY!&amp;n=1&amp;rf=friendphotoincenter&amp;t=5"></a><div id="fhp_photo_imgcenter" style="visibility:hidden;position:absolute;top:0px;left:0px;border:1px;background-color:#f00;">中心点</div><div class="bottom_cover"><p id="fhp_photo_desc" class="photo_describe"><span title="朱红兵&nbsp;《手机相册》">朱红兵&nbsp;《手机相册》</span></p></div><div class="friend_pages" style="display:"><a id="fhp_pager_left" href="javascript:void(0);" tcisdkey="prev" onclick="QPHOTO.mod.friendHotPhoto.tcisdStats(this);QPHOTO.mod.friendHotPhoto.goLeft();" class="page_nav_before" style=""><i class="icon_nav_before">上一个</i></a><a id="fhp_pager_right" href="javascript:void(0);" tcisdkey="next" onclick="QPHOTO.mod.friendHotPhoto.tcisdStats(this);QPHOTO.mod.friendHotPhoto.goRight();" class="page_nav_after" style=""><i class="icon_nav_after">下一个</i></a></div></div></div><div class="interact"><span class="gb_bt2" style="border-width:0px;width:106px;height:36px;padding:0px;" id="fhp_like" __ugclikebtnid="108"><a id="fhp_like_fake" class="gb_bt2" href="javascript:void(0)" style="display: none;"><i class="icon icon_praise"><i></i></i>赞</a><span style="overflow: hidden; zoom: 1; margin: 0px; display: block;"><a class="gb_bt2" href="javascript:void(0)" style="display: none;"><i class="icon icon_praise"><i></i></i>赞</a><a class="gb_bt2" href="javascript:void(0)" style="display: inline-block;"><i class="icon icon_praise"><i></i></i>赞</a><a class="gb_bt2" href="javascript:void(0)" style="display: none;"><i class="icon icon_praise"><i></i></i>已赞</a><a class="gb_bt2" href="javascript:void(0)" style="display: none;"><i class="icon icon_praise"><i></i></i>已赞</a></span></span><span id="_4_mod_hotfriend_plus_1" style="position:absolute;left:80px;margin-top:12px;color:#ff720b;font-weight:bold;display:none;">+1</span><a class="gb_bt2 r" href="javascript:void(0)" tcisdkey="change" onclick="QPHOTO.mod.friendHotPhoto.tcisdStats(this);QPHOTO.mod.friendHotPhoto.goRight();" title="换一张">换一张</a></div><div class="mod_like" id="fhp_like_content">来成为第一个赞的人吧</div><div class="friend_comment_lists"><div class="mod_comments" id="fhp_comment_list"><div class="comments_list"><ul></ul></div></div></div><div class="go_comment" id="fhp_op_unactive" onclick="QPHOTO.mod.friendHotPhoto.focusComment();return false;" style=""><div class="bor2"><div class="text-holder">我也说一句</div></div></div><div class="friend_photo_interactive" id="fhp_op_active" style="display:none"><div class="comment_box_active"><div id="fhp_comment_box"></div></div></div></div></div>
                        <div id="QM_100012_Bottom" class="box_ft">
                        </div>
                    </div></div>
            </div>

            <div id="care_friend_container" class="col-main none">
                <div class="col-main-feed">
                    <div class="fn-feed-control-v2 clearfix">
                        <div class="control-inner">
                            <div class="feed-control-tab">
                                <a id="care_menu_all" href="javascript:;" class="item-on item-on-slt">
                                    <span id="care_menu_nick">关心的好友</span>
                                    <i class="ui-icon icon-feed-down"></i>
                                </a>
                                <span class="line"></span>

                                <div class="scrollbar" id="care_menu_cnt" style="display:none;">
                                    <div class="fn-single-select">
                                        <ul id="care_menu_list" class="avatar-list"></ul>
                                    </div>
                                </div>
    							<span title="您可能关心的好友动态" class="item" style="font-weight:bold;display:none" id="may_care_title">您可能关心的好友动态</span>
                            </div>
                            <div class="feed-control-tab-op">
                                <a title="刷新动态" id="feed_care_refresh" href="javascript:;" class="item-left"><i class="ui-icon  icon-refresh-v9"></i></a>
                            </div>
                        </div>
                    </div>
                    <div class="fn-feed-container">
                        <div class="feed  feed-v9">
                            <div class="feed_inner">
                                <div style="display:none;" id="feed_care_update" class="tipsbox bg2"></div>
                                <ul id="feed_care_list"></ul>
                                <div id="feed_care_tips" style="visibility:hidden;" class="f-more-label"></div>
                            </div>
                        </div>
                    </div>
                </div>
                <div class="col-main-sidebar">
                    <div class="icenter-right-mod mod-add-care" id="spcare_addspecial">
                        <div class="hd">
                            <a href="javascript:;" class="btn-add-care-big" id="add_care_btn"><i class="ui-icon icon-care-red-b"></i>添加特别关心</a>
                        </div>
                        <div class="bd"></div>
                    </div>

                    <div class="icenter-right-mod mod-special-counter" id="spcare_careme" style="display:none">
	                    <div class="hd">
	                        <h3>关心我的</h3>
	                    </div>
	                    <div class="bd">
	                        <div class="box_bd qz-grid">
	                            <div class="special-counter qz-left bg5">
	                                <div class="in-hd qz-grid bg6">
	                                    <s class="qz-left ui_icon sp-sc-heart"></s><s class="qz-right ui_icon sp-sc-heart"></s>
	                                </div>
	                                <div class="in-bd"><strong class="sc-counter c_tx3" id="sp_careme_count">-</strong></div>
	                            </div>
	                            <div class="qz-main c_tx2" id="sp_careme_desc" title="为保护用户隐私，此处将不展示具体好友。">
									<p><a href="javascript:" id="sp_careme_action"></a></p>
	                            </div>
	                        </div>
	                    </div>
	                </div>



					<div class="icenter-right-mod mod-my-cared-list" id="my_care_list_con" style="display:none">
						<div class="hd">
							<h3>我关心的 <span id="carelist_count_new"></span></h3>
						</div>
						<div class="bd">
							<ul id="my_care_list"></ul>
						</div>
						<div class="ft">
							<a href="javascript:" class="prev-page unclick" title="上一页"><i></i></a>
							<a href="javascript:" class="next-page unclick" title="下一页"><i></i></a>
						</div>
					</div>



                <div style="width:230px;" id="sidebar-care-fixed"></div></div>
            </div>
            <div id="app_right_container" class="col-main none"></div>

            <div id="leftMenu" class="col-menu">



<div class="mod-side-nav mod-side-nav-message" data-version="20140912">
    <div class="inner">
        <div class="bd">

            <ul id="tab_menu_list" class="sn-list">
                <li type="friend" class="current">
                    <a id="tab_menu_friend" href="javascript:;" class="qz-grid">
                        <div class="qz-left"><i class="ui-icon sp-friend-feed"></i></div>
                        <div class="qz-right"><span class="sn-count none" id="tab_menu_friend_cnt">0</span></div>
                        <div class="qz-main"><span class="sn-title">好友动态</span></div>
                    </a>
                </li>

                <li type="care">
                    <a id="tab_menu_care" href="javascript:;" class="qz-grid" accesskey="t">
                        <div class="qz-left"><i class="ui-icon sp-close-feed"></i></div>
                        <div class="qz-right"><span class="sn-count none" id="tab_menu_care_cnt">0</span></div>
                        <div class="qz-main"><span class="sn-title">特别关心</span></div>
                    </a>
                </li>


                <li type="me" class="">
                    <a id="tab_menu_me" href="javascript:;" class="qz-grid" accesskey="y">
                        <div class="qz-left"><i class="ui-icon sp-at"></i></div>
                        <div class="qz-right"><span class="sn-count none" id="tab_menu_me_cnt">0</span></div>
                        <div class="qz-main"><span class="sn-title">与我相关</span></div>
                    </a>
                </li>


                <li type="mv">
                    <a id="tab_menu_mv" href="javascript:;" class="qz-grid" accesskey="m" style="display:none;">
                        <div class="qz-left"><i class="ui-icon sp-video"></i></div>
                        <div class="qz-right"><span class="sn-radio" id="tab_menu_mv_cnt"></span></div>
                        <div class="qz-main"><span class="sn-title">直播广场</span></div>
                    </a>
                </li>


                <li type="sdc">
                    <a id="tab_menu_sdc" href="javascript:;" class="qz-grid">
                        <div class="qz-left"><i class="ui-icon sp-past-years"></i></div>
                        <div class="qz-main"><span class="sn-title">那年今日</span></div>
                    </a>
                </li>






                <li type="class" class="ke_gray">
                    <a id="tab_menu_class" href="javascript:;" class="qz-grid txclass" data-url="https://ke.qq.com/index_new.html?source=qzone&amp;from=152">
                        <div class="qz-left"><i class="ui-icon sp-classroom"></i></div>
                        <div class="qz-main"><span class="sn-title">腾讯课堂</span></div>
                    </a>
                </li>

                <li type="app">
                    <a id="tab_menu_app" href="javascript:;" class="qz-grid">
                        <div class="qz-left"><i class="ui-icon sp-app-center"></i></div>
                        <div class="qz-right"><span class="sn-count none" id="tab_menu_app_cnt">0</span></div>
                        <div class="qz-main"><span class="sn-title">游戏应用</span></div>
                    </a>
                </li>


            </ul>

            <a id="tab_switch" title="展开" class="qz-grid more-toggle more-toggle-hide" href="javascript:;" data="______">
                <i class="ui-icon sp-nav-l-down"></i>
            </a>

            <ul id="tab_hide_list" class="sn-list none">

                    <li type="fav">
                         <a id="tab_menu_fav" href="javascript:;" class="qz-grid">
                            <div class="qz-left"><i class="ui-icon sp-fav"></i></div>
                            <div class="qz-main"><span class="sn-title">我的收藏</span></div>
                        </a>
                    </li>

            </ul>
        </div>
    </div>
</div>
<div class="mod-side-nav mod-side-nav-recently-used">
    <div class="inner" id="applist_html">
        <div class="hd"></div>
        <div class="bd">
            <ul class="sn-list" id="tab_applist_show">





                    <li>
                        <a href="http://cf.qq.com/cp/a20170907open/hz.shtml" data-bosstrace="2156_190379_3_1_0_0" onclick="TCISD &amp;&amp; TCISD.hotClick('applist.leftposition.guangbo','applist.qzone.qq.com');" class="qz-grid" target="_blank">
                            <div class="qz-left"><i class="ui-icon sp-popularize sparkle"></i></div>
                            <div class="qz-main"><span class="sn-title">黄钻特权红包！</span></div>
                        </a>
                    </li>


                <li>
                    <a href="//game.qzone.qq.com/?via=QZONE.CENTER" class="qz-grid qz-gamecenter">
                        <div class="qz-left"><i class="ui-icon sp-gamecenter"></i></div>
                        <div class="qz-main"><span class="sn-title">游戏应用中心</span></div>
                    </a>
                </li>




















            <li data-recommend="0" data-apptype="0" data-newicon="0">
                <a class="qz-grid j-to-app" data-index="" data-id="333" data-via="" data-traceinfo="" href="javascript:;">
                <div class="qz-left">
                    <img alt="礼物" src="//i.gtimg.cn/open/app_icon/00/00/03/33/333_16.png">
                </div>

                <div class="qz-main">
                    <span class="sn-title">
                        礼物
                    </span>
                </div>
                </a>
                <a id="aListSendGift" href="javascript:;" class="act">赠送</a>
            </li>

















            <li data-recommend="0" data-apptype="0" data-newicon="0">
                <a class="qz-grid j-to-app" data-index="" data-id="100622432" data-via="" data-traceinfo="" href="javascript:;">
                <div class="qz-left">
                    <img alt="QQ情侣" src="//i.gtimg.cn/open/app_icon/00/62/24/32/100622432_16.png">
                </div>

                <div class="qz-main">
                    <span class="sn-title">
                        QQ情侣
                    </span>
                </div>
                </a>

            </li>

















            <li data-recommend="0" data-apptype="0" data-newicon="0">
                <a class="qz-grid j-to-app" data-index="" data-id="100718848" data-via="" data-traceinfo="" href="javascript:;">
                <div class="qz-left">
                    <img alt="龙骑士传" src="//i.gtimg.cn/open/app_icon/00/71/88/48/100718848_16.png">
                </div>

                <div class="qz-main">
                    <span class="sn-title">
                        龙骑士传
                    </span>
                </div>
                </a>

            </li>




                <li class="sn-other">
                    <a href="//rc.qzone.qq.com/appstore/allapp?via=QZONE.CENTER">我的全部游戏应用&gt;&gt;</a>
                </li>



                    <li class="sn-line">
                        <sapn>热门游戏</sapn>
                    </li>

















            <li data-recommend="0" data-apptype="0" data-newicon="0">
                <a class="qz-grid j-to-app" data-index="" data-id="1106193011" data-via="QZ.MYAPP.HOT" data-traceinfo="2544_190318_848_1_0_0" href="javascript:;">
                <div class="qz-left">
                    <img alt="龙神契约" src="//i.gtimg.cn/open/app_icon/06/19/30/11/1106193011_16.png">
                </div>



                <div class="qz-main">
                    <span class="sn-title">
                        龙神契约
                    </span>
                </div>
                </a>

            </li>

















            <li data-recommend="0" data-apptype="0" data-newicon="0">
                <a class="qz-grid j-to-app" data-index="" data-id="1105370739" data-via="QZ.MYAPP.HOT" data-traceinfo="2544_190316_848_1_0_0" href="javascript:;">
                <div class="qz-left">
                    <img alt="剑灵洪门崛起" src="//i.gtimg.cn/open/app_icon/05/37/07/39/1105370739_16.png">
                </div>



                <div class="qz-main">
                    <span class="sn-title">
                        剑灵洪门崛起
                    </span>
                </div>
                </a>

            </li>

















            <li data-recommend="1" data-apptype="0" data-newicon="0">
                <a class="qz-grid j-to-app" data-index="" data-id="1105972482" data-via="QZ.MYAPP.HOT" data-traceinfo="2544_190317_848_1_0_0" href="javascript:;">
                <div class="qz-left">
                    <img alt="战神三十六计" src="//i.gtimg.cn/open/app_icon/05/97/24/82/1105972482_16.png">
                </div>


                        <div class="qz-right"><i class="ui_icon icon_app_recommend"></i></div>


                <div class="qz-main">
                    <span class="sn-title">
                        战神三十六计
                    </span>
                </div>
                </a>

            </li>







                    <li class="sn-line">
                        <sapn>新品抢先</sapn>
                    </li>

















            <li data-recommend="0" data-apptype="0" data-newicon="1">
                <a class="qz-grid j-to-app" data-index="" data-id="1106259913" data-via="QZ.MYAPP.NEW" data-traceinfo="2545_190172_848_1_0_0" href="javascript:;">
                <div class="qz-left">
                    <img alt="九阴九阳" src="//i.gtimg.cn/open/app_icon/06/25/99/13/1106259913_16.png">
                </div>


                        <div class="qz-right"><i class="ui_icon icon_app_new_cn"></i></div>


                <div class="qz-main">
                    <span class="sn-title">
                        九阴九阳
                    </span>
                </div>
                </a>

            </li>








            </ul>


                <a id="tab_applist_switch" title="展开" class="qz-grid more-toggle more-toggle-hide" href="javascript:;">
                   <i class="ui-icon sp-nav-l-down"></i>
                </a>

                <ul class="sn-list none" id="tab_applist_hide">
                    <li class="sn-line">
                        <sapn>休闲娱乐</sapn>
                    </li>

















            <li data-recommend="0" data-apptype="0" data-newicon="0">
                <a class="qz-grid j-to-app" data-index="" data-id="1105974527" data-via="QZ.MYAPP.CASUAL" data-traceinfo="2546_190405_848_1_0_0" href="javascript:;">
                <div class="qz-left">
                    <img alt="超神狙击" src="//i.gtimg.cn/open/app_icon/05/97/45/27/1105974527_16.png">
                </div>



                <div class="qz-main">
                    <span class="sn-title">
                        超神狙击
                    </span>
                </div>
                </a>

            </li>


















            <li data-recommend="0" data-apptype="0" data-newicon="0">
                <a class="qz-grid j-to-app" data-index="" data-id="363" data-via="QZ.MYAPP.CASUAL" data-traceinfo="2546_190406_848_1_0_0" href="javascript:;">
                <div class="qz-left">
                    <img alt="欢乐斗地主（腾讯）" src="//i.gtimg.cn/open/app_icon/00/00/03/63/363_16.png">
                </div>



                <div class="qz-main">
                    <span class="sn-title">
                        欢乐斗地主（腾讯）
                    </span>
                </div>
                </a>

            </li>

















            <li data-recommend="0" data-apptype="0" data-newicon="0">
                <a class="qz-grid j-to-app" data-index="" data-id="1105587804" data-via="QZ.MYAPP.CASUAL" data-traceinfo="2546_190407_848_1_0_0" href="javascript:;">
                <div class="qz-left">
                    <img alt="全民农场" src="//i.gtimg.cn/open/app_icon/05/58/78/04/1105587804_16.png">
                </div>



                <div class="qz-main">
                    <span class="sn-title">
                        全民农场
                    </span>
                </div>
                </a>

            </li>



                </ul>


            <ul class="sn-list" id="tab_otherlist_show">

                <li>
                    <a class="qz-grid sn-act" href="javascript:;" id="aPointMall" data-id="-10004">
                       <div class="qz-left"><i class="ui-icon sp-shop"></i></div>
                       <div class="qz-main"><span class="sn-title">积分兑换实物</span></div>
                    </a>
                </li>

                <li>
                    <a class="qz-grid sn-act" href="javascript:;" id="aMyGift" data-id="-10001">
                        <div class="qz-left"><i class="ui-icon sp-gift"></i></div>
                        <div class="qz-main"><span class="sn-title">每日抽奖</span></div>
                    </a>
                </li>
            </ul>

        </div>
    </div>
</div>

                <div class="collet_box fn_paipai" id="QM_Container_11">
                    <div id="QM_Container_100006">
                        <div class="box_bd" id="QM_100006_Body"><div class="sns_icenter" id="gdtLcModCnt">
<div class="gg-qz-icenter qbs" id="gdtApprec" style=""><a href="javascript:;" class="btn-close  _js_negative"><i class="ui-icon icon-close"></i></a><div class="mod-tuwen">
<h3 class="" id="gdtApprecTitle">广告</h3>
<span id="gdtApprecCnt"><div class="main"> <a onclick="return GDT._click( '648519450970936527', '36099302', {el: this} );" onmousedown="GDT. ClickManage.mousedown( '648519450970936527', '36099302', {el: this} );" onmouseup="GDT. ClickManage.mouseup( '648519450970936527', '36099302', {el: this} );" onmouseover="GDT. ClickManage.mouseover( '648519450970936527', '36099302', {el: this} );" href="http://c.gdt.qq.com/gdt_click.fcg?viewid=tKRYqHY6CXnFYrlFMp5dm0fvRJE4BkBeXMvlI8Nk3PObY917BsysMbn73hho4mfraC1OcdoLTOr0gD7inQmWFNdWXPik9jTf!rn19ulBYHyj1WZiSjeP0eltzuZ5MHBORYxyUPeM3iI&amp;jtype=0&amp;i=1&amp;os=0" gdtoriurl="http://c.gdt.qq.com/gdt_click.fcg?viewid=tKRYqHY6CXnFYrlFMp5dm0fvRJE4BkBeXMvlI8Nk3PObY917BsysMbn73hho4mfraC1OcdoLTOr0gD7inQmWFNdWXPik9jTf!rn19ulBYHyj1WZiSjeP0eltzuZ5MHBORYxyUPeM3iI&amp;jtype=0&amp;i=1&amp;os=0" target="_blank"><img class="_js_olink" src="https://pgdt.gtimg.cn/gdt/0/DAAWQBwACgADSAAbBZu524Bee7u7vL.jpg/0?ck=1b38bcf7cfd3d3ba0bdb7737db34a9d2"></a>  <a href="http://e.qq.com/?from=qzone_pc" class="j-corner" target="_blank" style="bottom: 3px;"><div class="tr"></div> <div class="info"> <p class="text1"></p> <p class="text2"></p> </div></a>  </div> <div class="ft _js_txtrow" style=""> <div><p style="" class="_js_show"><a onclick="return GDT._click( '648519450970936527', '36099302', {el: this} );" onmousedown="GDT. ClickManage.mousedown( '648519450970936527', '36099302', {el: this} );" onmouseup="GDT. ClickManage.mouseup( '648519450970936527', '36099302', {el: this} );" onmouseover="GDT. ClickManage.mouseover( '648519450970936527', '36099302', {el: this} );" href="http://c.gdt.qq.com/gdt_click.fcg?viewid=tKRYqHY6CXnFYrlFMp5dm0fvRJE4BkBeXMvlI8Nk3PObY917BsysMbn73hho4mfraC1OcdoLTOr0gD7inQmWFNdWXPik9jTf!rn19ulBYHyj1WZiSjeP0eltzuZ5MHBORYxyUPeM3iI&amp;jtype=0&amp;i=1&amp;os=0" gdtoriurl="http://c.gdt.qq.com/gdt_click.fcg?viewid=tKRYqHY6CXnFYrlFMp5dm0fvRJE4BkBeXMvlI8Nk3PObY917BsysMbn73hho4mfraC1OcdoLTOr0gD7inQmWFNdWXPik9jTf!rn19ulBYHyj1WZiSjeP0eltzuZ5MHBORYxyUPeM3iI&amp;jtype=0&amp;i=1&amp;os=0" target="_blank" class="_js_olink">2017年T恤特价，4件99，包邮！</a></p></div> </div> <span id="gdtApprecExt" style="display: none;"></span></span>
</div></div>
<div class="gg-qz-icenter qbs" style="display: none" id="gdtSScnt">
</div>
<div class="gg-qz-icenter qbs" id="gdtpaipai" style="display: none">
</div>
<span style="width:0;height: 0;"></span>
<div class="gg-qz-icenter qbs" id="gdtdiscover" style=""><a href="javascript:;" class="btn-close  _js_negative"><i class="ui-icon icon-close"></i></a><div class="mod-tuwen"> <h3>广告</h3> <div class="main"> <a onclick="return GDT._click( '9007200359370577103', '36113647', {el: this} );" onmousedown="GDT. ClickManage.mousedown( '9007200359370577103', '36113647', {el: this} );" onmouseup="GDT. ClickManage.mouseup( '9007200359370577103', '36113647', {el: this} );" onmouseover="GDT. ClickManage.mouseover( '9007200359370577103', '36113647', {el: this} );" href="http://c.gdt.qq.com/gdt_click.fcg?viewid=jgAHFRYYGMjFYrlFMp5dm5iEfBwep!8DKWAlX9adxVFVgqxAiQEBlCl0dXxquV4dP_4VkA!8_hdUiq8x0_S1hEDMOPbwVNC8Q_go_hczMZDI4piwyvRmb!U9A3WKMK8Of7FoiOxO4kw&amp;jtype=0&amp;i=1&amp;os=0" gdtoriurl="http://c.gdt.qq.com/gdt_click.fcg?viewid=jgAHFRYYGMjFYrlFMp5dm5iEfBwep!8DKWAlX9adxVFVgqxAiQEBlCl0dXxquV4dP_4VkA!8_hdUiq8x0_S1hEDMOPbwVNC8Q_go_hczMZDI4piwyvRmb!U9A3WKMK8Of7FoiOxO4kw&amp;jtype=0&amp;i=1&amp;os=0" target="_blank"><img src="https://pgdt.gtimg.cn/gdt/0/DAAWFeKACgADSAAaBZu56-CvGDs6RY.jpg/0?ck=139f79a0288a08b25033952f1c12b13a"></a> <a href="http://e.qq.com/?from=qzone_pc" class="j-corner" target="_blank" style="bottom: 3px;"><div class="tr"></div> <div class="info"> <p class="text1"></p> <p class="text2"></p> </div></a>  </div> <div class="ft"> <div> <p><a onclick="return GDT._click( '9007200359370577103', '36113647', {el: this} );" onmousedown="GDT. ClickManage.mousedown( '9007200359370577103', '36113647', {el: this} );" onmouseup="GDT. ClickManage.mouseup( '9007200359370577103', '36113647', {el: this} );" onmouseover="GDT. ClickManage.mouseover( '9007200359370577103', '36113647', {el: this} );" href="http://c.gdt.qq.com/gdt_click.fcg?viewid=jgAHFRYYGMjFYrlFMp5dm5iEfBwep!8DKWAlX9adxVFVgqxAiQEBlCl0dXxquV4dP_4VkA!8_hdUiq8x0_S1hEDMOPbwVNC8Q_go_hczMZDI4piwyvRmb!U9A3WKMK8Of7FoiOxO4kw&amp;jtype=0&amp;i=1&amp;os=0" gdtoriurl="http://c.gdt.qq.com/gdt_click.fcg?viewid=jgAHFRYYGMjFYrlFMp5dm5iEfBwep!8DKWAlX9adxVFVgqxAiQEBlCl0dXxquV4dP_4VkA!8_hdUiq8x0_S1hEDMOPbwVNC8Q_go_hczMZDI4piwyvRmb!U9A3WKMK8Of7FoiOxO4kw&amp;jtype=0&amp;i=1&amp;os=0" target="_blank">车载音乐U盘，99元包邮免费试用！</a></p> </div> </div> </div></div>
<div class="gg-qz-icenter qbs mod-pic140" id="gdtmergepos" style="display: none">
</div>
</div></div>
                    </div>
                </div>
            </div>
        </div>

        <div id="pageApp" class="layout-app mod_wrap none">
            <div class="mod_wrap_inner">
                <div class="mod_wrap_bd">
                    <div class="mod_wrap_iframe" id="app_container" style="min-height:435px;"><iframe class="app_canvas_frame" scrolling="no" allowtransparency="yes" frameborder="no" width="100%" height="435" src="about:blank" id="ttinfo"></iframe></div>
                </div>
            </div>
        </div>
        <div id="page3rdApp" class="layout-app none" style="min-height:750px;"></div>
    </div>
    <div class="layout-copyright">
	<p><a href="http://z.qzone.com/" onclick="TCISD &amp;&amp; TCISD.hotClick('bottomqzoneapp', 'user.qzone.com')" target="_blank">空间手机版</a> | <a href="http://vip.qzone.com/" onclick="TCISD &amp;&amp; TCISD.hotClick('threefromendofpage', 'guangxi.vip.qzone.com', 'main')" target="_blank">黄钻贵族</a> | <a href="http://user.qzone.qq.com/20050606" target="_blank">官方Qzone</a> | <a href="http://connect.qq.com/" target="_blank">QQ互联</a> | <a href="http://page.opensns.qq.com/" target="_blank">认证服务</a> | <a href="http://blog.qq.com/" target="_blank">腾讯博客</a> | <a href="http://service.qq.com/special_auto/qzone.html" target="_blank">腾讯客服</a> | <a href="http://qzone.qq.com/web/tk.html" target="_blank">QQ空间服务协议</a> | <a href="http://wiki.open.qq.com/wiki/Tencent_Open_Platform_Complaint_Guidelines" target="_blank">Complaint Guidelines</a> | <a href="http://www.qq.com/culture.shtml" target="_blank">粤网文[2017]6138-1456号</a></p>
	<p>Copyright © 2005 - 2017 Tencent.<a href="http://www.tencent.com/en-us/le/copyrightstatement.shtml" target="_blank">All Rights Reserved.</a></p>
	<p>腾讯公司 <a href="http://www.tencent.com/zh-cn/le/copyrightstatement.shtml" target="_blank">版权所有</a></p>
</div>
    </div>


</div>
<div class="fix-layout">    <div class="gb-operation-area" id="_returnTop_layout_inner">        <a href="javascript:;" id="goto_top_btn" class="gb-operation-button return-top">            <i class="gb-operation-icon" title="返回顶部"></i>            <span class="gb-operation-text">顶部</span>        </a>    <a href="javascript:;" class="gb-operation-button hot-msg" id="site_hot_btn" style="display: block;"><span class="gb-operation-text" style="display: block;">热点</span></a></div></div>
<div id="QZ_Float_Items_Container" class="lay-shop-item-fixed"><div id="QZBlindMenu" style="overflow: hidden; width: 0px; height: 0px; position: absolute; left: -1000px;"><a href="javascript:;" role="navigation" accesskey="z" tabindex="1" title="主页" onclick="custom_menu_swf('-1');TCISD.pv('homeact2.qzone.qq.com', '/blindMenu_-1');return false">主页</a><a href="javascript:;" role="navigation" accesskey="r" tabindex="1" title="日志" onclick="custom_menu_swf('2');TCISD.pv('homeact2.qzone.qq.com', '/blindMenu_2');return false">日志</a><a href="javascript:;" role="navigation" accesskey="x" tabindex="1" title="相册" onclick="custom_menu_swf('4');TCISD.pv('homeact2.qzone.qq.com', '/blindMenu_4');return false">相册</a><a href="javascript:;" role="navigation" accesskey="l" tabindex="1" title="留言板" onclick="custom_menu_swf('334');TCISD.pv('homeact2.qzone.qq.com', '/blindMenu_334');return false">留言板</a><a href="javascript:;" role="navigation" accesskey="s" tabindex="1" title="说说" onclick="custom_menu_swf('311');TCISD.pv('homeact2.qzone.qq.com', '/blindMenu_311');return false">说说</a><a href="javascript:;" role="navigation" accesskey="g" tabindex="1" title="个人档" onclick="custom_menu_swf('1');TCISD.pv('homeact2.qzone.qq.com', '/blindMenu_1');return false">个人档</a><a href="javascript:;" role="navigation" accesskey="y" tabindex="1" title="音乐" onclick="custom_menu_swf('305');TCISD.pv('homeact2.qzone.qq.com', '/blindMenu_305');return false">音乐</a><a href="javascript:;" role="navigation" accesskey="s" tabindex="1" title="时光轴" onclick="custom_menu_swf('219');TCISD.pv('homeact2.qzone.qq.com', '/blindMenu_219');return false">时光轴</a></div></div>
<div class="ui-tooltip" style="display:none;" id="g_reply_tips">
	<div class="inner">
		<div class="bd">回复</div>
		<b class="arrow arrow-down"></b>
	</div>
</div>
<div class="ui-tooltip" style="display:none;" id="g_delete_tips">
    <div class="inner">
        <div class="bd">删除</div>
        <b class="arrow arrow-down"></b>
    </div>
</div>
<script type="text/javascript">
g_T.fwp[4] = new Date();var g_V={qz:'_v8_2.1.65',jq:'2.0.3',sea:'2.1.1',core_ic:'50108',fp:'60628',ic:'50215',popcss:'?d=140918103710',delay:'-170302163216',req:'61221',ci:'_20131225'};
var g_ISP={p:"255",
i:"-3",
f:"1",
c:'10.62.148.137'},
g_type="M",
g_Set="",
g_ICSet="",
g_SPrefix="",
g_SPrefix_Server="",
g_DPrefix="",
g_dns_name="",
g_UserBitmap="18d095a00102c101",
g_LoginBitmap="18d095a00102c101",
g_UserunimBitmap="0808008020300301",
g_LoginUnimBitmap="0808008020300301",
g_UserVipStarBitmap="0000000000000000",
g_LoginVipStarBitmap="0000000000000000",
g_Bitmap_3rd="0000000000000840",
g_IZone_Flag=+"1",
g_diyTitle="",
g_SceneID=76638,
g_neatstyle=0,
g_7DDressNo=0,
g_StyleID="31",
g_fullMode=6,
g_Mode="ofp",
g_version=8,
g_isFriend=0,
g_NowTime=1506002258,
g_StaticFlag="0",
g_isDirectApp=0,
ownermode=(g_iLoginUin==g_iUin),
ownerProfileSummary=["承諾，\/aiq一世的誓言",
					1,
					26,
					"中国@重庆@沙坪坝",
					"→ωó暧р女孩♂",
					"对一个女孩的思念",
					"1990-12-11",
					"0400080000000005"],
g_isOFP="1",
g_isHideTitle="0",
g_icLayout="0",
g_isFromPengYou="",
g_isBrandQzone="",
g_isFamousQzone="",
g_xwMode="",
g_needDec=1,
g_isFastPav=1,
g_scenarioidle="1536",
g_isBGScroll=0,
g_Dressup={items:[{type:347,itemno:105258,posx:842,posy:186,posz:0,height:250,width:0,flag:0},{type:25,itemno:123052,posx:0,posy:0,posz:0,height:0,width:0,flag:1},{type:1,itemno:123053,posx:0,posy:0,posz:0,height:0,width:0,flag:218},{type:13,itemno:123054,posx:140,posy:0,posz:0,height:70,width:528,flag:1},{type:19,itemno:1,posx:0,posy:0,posz:0,height:250,width:0,flag:0} ],windows:[{appid:4,mode:2,posx:2,posy:0,posz:0,height:134,width:65535,wndid:1},{appid:2,mode:2,posx:2,posy:1,posz:0,height:330,width:65535,wndid:1},{appid:333,mode:2,posx:2,posy:2,posz:0,height:151,width:65535,wndid:1},{appid:218,mode:2,posx:2,posy:3,posz:0,height:0,width:65535,wndid:1},{appid:847,mode:2,posx:2,posy:4,posz:0,height:96,width:65535,wndid:1},{appid:311,mode:2,posx:2,posy:5,posz:0,height:589,width:65535,wndid:55}]},
g_isLite=0,
g_yPos=+'',
g_portraitTime=+'1382948711',
g_loginPortraitTime = + '1382948711',
g_firstLoginTime=+'1505994652',
g_isLikable=+'1',
g_likedNumber=+'4',
g_likeable_today=+"1",
g_supportWebp=+'1',
g_OSClass="os-mac",
g_applistWallNum = parseInt("0" || "0",10),
g_isStarVip = +"0",
g_isStarVipYear = +"0",
g_iStarVipLevel = +"0",


g_iEmoji=[],
g_PreData={app:{
			     ret:0,
			     data:{},
			     flag:1
    },
	flower:{
			ret:0,
			score:+"11569",
			gardener:+"41"
	},
	hat:{
			ret:-1,
	        id:+""
	},
	hotbar:{
		ret:-2
	},
	qbossData:{



				2161:{
					posid : "2161",
					res_traceinfo : "2161_190188_445_1_0_0",
					res_data : "{\x22jumpAction\x22:\x22https:\/\/act.qzone.qq.com\/vip\/2017\/vip-bns-v2\x22,\x22islink\x22:\x220\x22,\x22flash\x22:\x22http:\/\/qzonestyle.gtimg.cn\/qzone\/space_item\/boss_pic\/2161_2017_9\/1504666647980_417893.swf\x22,\x22flash_md5\x22:\x2247387de19b2618ef71bf67dd0b100339\x22}",
					res_preprocess : "{\x22img\x22:\x22http:\/\/qzonestyle.gtimg.cn\/qzone\/space_item\/boss_pic\/2161_2017_9\/1504666647980_417893.swf\x22,\x22link\x22:\x22https:\/\/act.qzone.qq.com\/vip\/2017\/vip-bns-v2\x22,\x22islink\x22:0}"
				}


	},
	potential:{flag:+''},
	homeAddr:{"hco":"中国","hp":"四川","hc":"广安"},
	qq_friendNum:"435"
},
g_app_identifier='',
qlogoDomain="store.qq.com",
g_starstamp_pic = +'1505914848',
g_starstamp_emotion = +'1504695062',
g_ic_fpfeedsType='',
g_Data={
	feedsPart1: {result:'',message:'',main:{needFold:'',icServerTime:'1506002258',icView:'1',daylist:'',uinlist:'',hasMoreFeeds_0:true,hasMoreFeeds_1:true,hasMoreFeeds_100:true,foldAllHostBTNClass:'',foldAllHostBTNTitle:'',foldAllHostFeedDisplay:'',friend_more:'',host_more:'',aboutHostFeeds_new_cnt:'0',replyHostFeeds_new_cnt:'0',host_pav_offset:'6',get_visitorfeed:'1',myFeeds_new_cnt:'0',friendFeeds_new_cnt:'0',friendFeeds_newblog_cnt:'',friendFeeds_newphoto_cnt:'',newfeeds_uinlist:'',newfeeds_uinlist_more:'',newfeeds_special_cnt:'0',newfeeds_friend_cnt:'0',newfeeds_follow_cnt:'0',newfeeds_group_cnt:'',tips:'',js_showtips:'',lastaccesstime:'',lastAccessRelateTime:'',begintime:'0',dayspac:'0',hidedNameList:[],QzoneFeedsKey:'',hsFlag:'',aisortNextTime:'',aisortBeginTime:'',aisortEndTime:'',aisortOffset:'',blacklist:'105096414,364517951,409994995,464571314,537719886,583536810,598781523,641282731,670785221,769914037,824063635,892962047,1130548756,1146598238,1455196303,2512531654',pagenum:'1',externparam:'offset=6&total=7&basetime=1505959930&feedsource=0'}},
	feedsPart2: {host_data:[undefined],        about_data:[                                undefined        ],friend_data:[undefined],        firstpage_data:[                                {                        ver:'0',                        appid:'403',                        typeid:'2',                        key:'403_2_294477044',                        flag:'1',                        dataonly:'0',                        feedno:'0',                        title:'',                        summary:'',                        appiconid:'403',                        clscFold:'icenter_list_extend',                        abstime:'1506000077',                        feedstime:' 21:21',                        userHome:'http://user.qzone.qq.com/573197470',                        namecardLink:'nameCard_573197470',                        ouin:'',                        uin:'573197470',                        opuin:'573197470',                        foldFeed:'',                        foldFeedTitle:'',                        showEbtn:'',                        scope:'1',                        hideExtend:'',                        nickname:'李甜',                        remark:'',score:'',                        type:'',                        vip:'vip_0',                        bitmap:'18d195a00102c109',     yybitmap:'0008000000000000',                        info_user_name:'',                        logimg:'http://qlogo3.store.qq.com/qzone/573197470/573197470/50?1503739125',                        bor:'',                        lastFeedBor:'',                        list_bor2:'',                        info_user_display:'',                        upernum:'',                        oprType:'0',                        moreflag:'',                        otherflag:'0_0_0_0_0_0_0',                        sameuser:{},                        uper_isfriend:[],                        uperlist:[],mergeData:[undefined]                },                                {                        ver:'1',                        appid:'4',                        typeid:'27',                        key:'294477044_74020d9b1505911452783',                        flag:'1',                        dataonly:'0',                        feedno:'1',                        title:'',                        summary:'',                        appiconid:'4',                        clscFold:'icenter_list_extend',                        abstime:'1505999990',                        feedstime:' 21:19',                        userHome:'http://user.qzone.qq.com/654858397',                        namecardLink:'nameCard_654858397',                        ouin:'',                        uin:'654858397',                        opuin:'654858397',                        foldFeed:'',                        foldFeedTitle:'',                        showEbtn:'',                        scope:'1',                        hideExtend:'',                        nickname:'谭超',                        remark:'',score:'',                        type:'',                        vip:'vip_0',                        bitmap:'184095800000c401',     yybitmap:'0000000000000000',                        info_user_name:'',                        logimg:'http://qlogo2.store.qq.com/qzone/654858397/654858397/50',                        bor:'',                        lastFeedBor:'',                        list_bor2:'',                        info_user_display:'',                        upernum:'',                        oprType:'0',                        moreflag:'',                        otherflag:'0_0_0_0_0_0_0',                        sameuser:{},                        uper_isfriend:[],                        uperlist:[],mergeData:[undefined]                },                                {                        ver:'1',                        appid:'217',                        typeid:'3',                        key:'1101924592_294477044_294477044',                        flag:'1',                        dataonly:'0',                        feedno:'2',                        title:'',                        summary:'',                        appiconid:'217',                        clscFold:'icenter_list_extend',                        abstime:'1505997634',                        feedstime:' 20:40',                        userHome:'http://user.qzone.qq.com/1101924592',                        namecardLink:'nameCard_1101924592',                        ouin:'',                        uin:'1101924592',                        opuin:'1101924592',                        foldFeed:'',                        foldFeedTitle:'',                        showEbtn:'',                        scope:'1',                        hideExtend:'',                        nickname:'传小凤',                        remark:'',score:'',                        type:'',                        vip:'vip_0',                        bitmap:'18d095a00000c101',     yybitmap:'0000000000000001',                        info_user_name:'',                        logimg:'http://qlogo1.store.qq.com/qzone/1101924592/1101924592/50?1395557823',                        bor:'',                        lastFeedBor:'',                        list_bor2:'',                        info_user_display:'',                        upernum:'',                        oprType:'0',                        moreflag:'',                        otherflag:'0_0_0_0_0_0_0',                        sameuser:{},                        uper_isfriend:[],                        uperlist:[],mergeData:[undefined]                },                                {                        ver:'1',                        appid:'217',                        typeid:'3',                        key:'1159106747_294477044_294477044',                        flag:'1',                        dataonly:'0',                        feedno:'3',                        title:'',                        summary:'',                        appiconid:'217',                        clscFold:'icenter_list_extend',                        abstime:'1505990743',                        feedstime:' 18:45',                        userHome:'http://user.qzone.qq.com/1159106747',                        namecardLink:'nameCard_1159106747',                        ouin:'',                        uin:'1159106747',                        opuin:'1159106747',                        foldFeed:'',                        foldFeedTitle:'',                        showEbtn:'',                        scope:'1',                        hideExtend:'',                        nickname:'段可昕的妈妈',                        remark:'',score:'',                        type:'',                        vip:'vip_0',                        bitmap:'18c195a00002c401',     yybitmap:'0000100000000001',                        info_user_name:'',                        logimg:'http://qlogo4.store.qq.com/qzone/1159106747/1159106747/50?1503326090',                        bor:'',                        lastFeedBor:'',                        list_bor2:'',                        info_user_display:'',                        upernum:'',                        oprType:'0',                        moreflag:'',                        otherflag:'0_0_0_0_0_0_0',                        sameuser:{},                        uper_isfriend:[],                        uperlist:[],mergeData:[undefined]                },                                {                        ver:'1',                        appid:'217',                        typeid:'3',                        key:'454782407_294477044_294477044_',                        flag:'1',                        dataonly:'0',                        feedno:'4',                        title:'',                        summary:'',                        appiconid:'217',                        clscFold:'icenter_list_extend',                        abstime:'1505981406',                        feedstime:' 16:10',                        userHome:'http://user.qzone.qq.com/454782407',                        namecardLink:'nameCard_454782407',                        ouin:'',                        uin:'454782407',                        opuin:'454782407',                        foldFeed:'',                        foldFeedTitle:'',                        showEbtn:'',                        scope:'1',                        hideExtend:'',                        nickname:'天之娇',                        remark:'',score:'',                        type:'',                        vip:'vip_0',                        bitmap:'1890952000008001',     yybitmap:'000000005460ac04',                        info_user_name:'',                        logimg:'http://qlogo4.store.qq.com/qzone/454782407/454782407/50',                        bor:'',                        lastFeedBor:'',                        list_bor2:'',                        info_user_display:'',                        upernum:'',                        oprType:'0',                        moreflag:'',                        otherflag:'0_0_0_0_0_0_0',                        sameuser:{},                        uper_isfriend:[],                        uperlist:[],mergeData:[undefined]                },                                {                        ver:'1',                        appid:'4',                        typeid:'27',                        key:'294477044_75f2333b1505911452783',                        flag:'1',                        dataonly:'0',                        feedno:'5',                        title:'',                        summary:'',                        appiconid:'4',                        clscFold:'icenter_list_extend',                        abstime:'1505974200',                        feedstime:' 14:10',                        userHome:'http://user.qzone.qq.com/382847106',                        namecardLink:'nameCard_382847106',                        ouin:'',                        uin:'382847106',                        opuin:'382847106',                        foldFeed:'',                        foldFeedTitle:'',                        showEbtn:'',                        scope:'1',                        hideExtend:'',                        nickname:'房亮',                        remark:'',score:'',                        type:'',                        vip:'novip',                        bitmap:'18d08de0040ce101',     yybitmap:'1000100000000003',                        info_user_name:'',                        logimg:'http://qlogo3.store.qq.com/qzone/382847106/382847106/50?1422787525',                        bor:'',                        lastFeedBor:'',                        list_bor2:'',                        info_user_display:'',                        upernum:'',                        oprType:'0',                        moreflag:'',                        otherflag:'0_0_0_0_0_0_0',                        sameuser:{},                        uper_isfriend:[],                        uperlist:[],mergeData:[undefined]                },                                {                        ver:'1',                        appid:'217',                        typeid:'3',                        key:'766792812_294477044_294477044_',                        flag:'1',                        dataonly:'0',                        feedno:'6',                        title:'',                        summary:'',                        appiconid:'217',                        clscFold:'icenter_list_extend',                        abstime:'1505959930',                        feedstime:' 10:12',                        userHome:'http://user.qzone.qq.com/766792812',                        namecardLink:'nameCard_766792812',                        ouin:'',                        uin:'766792812',                        opuin:'766792812',                        foldFeed:'',                        foldFeedTitle:'',                        showEbtn:'',                        scope:'1',                        hideExtend:'',                        nickname:'黄雯',                        remark:'',score:'',                        type:'',                        vip:'vip_0',                        bitmap:'18d095a00100c001',     yybitmap:'0000000000000001',                        info_user_name:'',                        logimg:'http://qlogo1.store.qq.com/qzone/766792812/766792812/50?1459091570',                        bor:'',                        lastFeedBor:'',                        list_bor2:'',                        info_user_display:'',                        upernum:'',                        oprType:'0',                        moreflag:'',                        otherflag:'0_0_0_0_0_0_0',                        sameuser:{},                        uper_isfriend:[],                        uperlist:[],mergeData:[undefined]                },                                undefined        ]},
	feeds:{
		type:'',
		hostFeedsFilter:'',
		view:'1'
	}
},
g_moreModules=["v8/dress/otherview_v8","v8/core/state"],
G_Param={};
G_Param.qz_referrer='';
if(''=='special'){
	document.getElementById('tab_menu_care').parentNode.className = 'current';
}else{
	if(g_ic_fpfeedsType=="friend"){
		document.getElementById('tab_menu_friend').parentNode.className = 'current';
	}else{
		document.getElementById('tab_menu_me').parentNode.className = 'current';
	}
}
g_weather={};

g_weather['save']={
	'country':'中国',
	'province':'',
	'city':'重庆',
	'temp':'27',
	'weather':'多云',
	'day':'0',
	'pm2d5':''
	};

var g_blog_loc_prefix='';
var g_BLOG_LOCATION_PREFIX='b', g_BLOG_LOCATION_LOGIN_PREFIX='b';



(function(){
	var ua = {}, agent = navigator.userAgent;
	if (window.ActiveXObject || window.msIsStaticHTML){
		ua.ie = 6;
		(window.XMLHttpRequest || (agent.indexOf('MSIE 7.0') > -1)) && (ua.ie = 7);
		(window.XDomainRequest || (agent.indexOf('Trident/4.0') > -1)) && (ua.ie = 8);
		(agent.indexOf('Trident/5.0') > -1) && (ua.ie = 9);
		(agent.indexOf('Trident/6.0') > -1) && (ua.ie = 10);
		(agent.indexOf('Trident/7.0') > -1) && (ua.ie = 11);

		if(ua.ie == 6 || ua.ie == 7 || ua.ie == 8){
			g_V.jq='1.10.2';
		}
	}
})();

(navigator.appVersion.indexOf("QQBrowser/7")>-1) && (function(){
	try{
		var styleTest = '';
		var lastCpuTime,checkInterval = 1000,checkCount=15,threadCPUData=[],
			qqbrowser = window.external.qqbrowser,
			systemInfo = eval('('+ qqbrowser.perfCounter.systemInfo +')'),
			memoryu = eval('('+ qqbrowser.perfCounter.memoryUsage +')'),
			memorya =  memoryu.Process.CommittedSize,
			memoryb,
			cpuNumber = systemInfo.LogicProcessors,
			intervalUS = checkInterval*1000;

		if(!cpuNumber){
			return;
		}

		function calcCpuUsage(){
			var cpuTime = eval('('+ qqbrowser.perfCounter.cpuTime +')');
			if(lastCpuTime){
				var delSystemTime = (cpuTime.system - lastCpuTime.system);
				threadCPUData.push(100 * (cpuTime.thread - lastCpuTime.thread) / intervalUS);
			}
			lastCpuTime = cpuTime;
		}
		function avg(arr){
			var h=0;
			for(var i=0,len=arr.length;i<len;i++){
				h+=arr[i];
			}
			return h/checkCount;
		}
		var _i=0,
		timer = setInterval(function(){
			if(_i<checkCount){
				calcCpuUsage();
				_i++;
			}else{
				clearTimeout(timer);
				setTimeout(function(){
					memoryu = eval('('+ qqbrowser.perfCounter.memoryUsage +')');
					memoryb =  memoryu.Process.CommittedSize;
					if(window.TCISD && window.TCISD.createTimeStat){
						var apiStat = TCISD.createTimeStat("apiTimeStat", [7820, 109, 6]);
						apiStat.setZero(0);
						G_Param.threadCPUUsage = avg(threadCPUData);
						G_Param.threadMemoryUsage = memoryb-memorya;
						apiStat.mark(1,G_Param.threadCPUUsage*1000);
						apiStat.mark(2,G_Param.threadMemoryUsage);
						apiStat.report();
					}
				},10000);
			}
		},checkInterval);
	}catch(err){}
})();


</script>

<script type="text/javascript" src="https://qzonestyle.gtimg.cn/ac/qzone/qzfl/qzfl_v8_2.1.65.js"></script>
<script type="text/javascript" src="https://qzonestyle.gtimg.cn/ac/lib/seajs/sea-2.1.1.js"></script>
<script type="text/javascript" src="https://qzonestyle.gtimg.cn/c/=/qzone/v8/core/seajs_config.js,/qzone/v8/engine/plugin-combo.js,/qzone/v8/core/ic.js?max_age=31536001&amp;d=50108"></script>

<script>
    (function(){
        // if(window.location.protocol!='https:' || window.g_cdn_proto !='https'){
        //     return;
        // }
        var newSiDomain = 'h5.qzone.qq.com/proxy/domain/qzonestyle.gtimg.cn';
        var cookieflag = document.cookie.indexOf('blabla')>-1;
        if(cookieflag && (/(?:^user\.(s\d\.)?|\d{5,}\.|rc\.)qzone\.qq\.com/.test(location.host))){
            newSiDomain = location.host+'/q/qzs';
        }
        (function(v,b,e,i,s){
            var arJs=[];
            if(!window.QZFL || !QZFL.lang){
                arJs.push(b,s,'/ac/qzone/qzfl/qzfl',g_V.qz,'.js',e);
            }
            if(!window.seajs){
                arJs.push(b,s,'/ac/lib/seajs/sea-'+g_V.sea+'.js',e);
            }

            if(arJs.length){


                if(s.indexOf('h5.qzone.qq.com')>-1){
                    window.g_script_retry = 'h5.qzone.qq.com';
                }else{
                    window.g_script_retry = 'user.qzone.qq.com';
                }
                arJs.push(b,s,'/c/=/qzone/v8/core/seajs_config.js,/qzone/v8/engine/plugin-combo.js,/qzone/v8/core/ic.js?max_age=31536001&amp;d='+g_V.core_ic,e);
                document.write(arJs.join(''));
            }
        })(g_V,'<script type="text/javascript" charset="utf-8" src="'+(window.g_cdn_proto || window.location.protocol)+'//','"><\/script>',imgcacheDomain,newSiDomain);
    })();

</script>

<script type="text/javascript">
(function(){

    var resultcode = 0;
    var retrycode = 0;
    var commandid;
    if(window.QZFL && QZFL.lang){
        resultcode = 1;
    }

    if(window.g_script_retry=='h5.qzone.qq.com'){
        retrycode = 1;
    }else if(window.g_script_retry =='user.qzone.qq.com'){
        retrycode = 2;
    }
    if(window.g_cdn_proto=='https:' || window.location.protocol=='https:'){
        commandid = 'http://pc.qzone.qq.com/cdn_connect_https_'+window.g_isOFP;
    }else{
        commandid = 'http://pc.qzone.qq.com/cdn_connect_http_'+window.g_isOFP;
    }
    window.haboStat({
       commandid : commandid,
       stime : 1000,
       resultcode : ('100102'+ retrycode + resultcode),
       touin : window.g_iLoginUin || 0,
       frequency :  (resultcode ===0 || location.href.indexOf('concat_debug_on')>-1) ?  1 : 10 //先临时改成全量报
    });




    var resultMap = {'https':2,'http':2};

    var doReport = function(){
        var commandid2;
        if(window.g_cdn_proto=='https:' || window.location.protocol=='https:'){
            commandid2 = 'http://pc.qzone.qq.com/cdn_https_connect_failed_type_' + window.g_isOFP;
        }else{
            return;
        }
        window.haboStat({
            commandid : commandid2,
            stime : 1000,
            resultcode : '100101'+resultMap['https']+resultMap['http'],
            touin : window.g_iLoginUin || 0,
            frequency :  1
        });
    };
    var doMark  = function(type, code){
        resultMap[type] = code;
        if(resultMap['https']<2 && resultMap['http']<2){
            clearTimeout(timerId);
            doReport();
        }
    };
    if(!resultcode){ //失败的，再增加一个探测
        var timerId = setTimeout(function(){
            doReport();
        },5000);
        var imghttps = new Image();
        var imghttp = new Image();
        imghttps.onload = function(){
            this.onload = function(){}
            doMark('https', 0);
        };
        imghttps.onerror = function(){
            this.onerror = function(){}
            doMark('https', 1);
        };
        imghttp.onload = function(){
            this.onload = function(){}
            doMark('http', 0);
        };
        imghttp.onerror = function(){
            this.onerror = function(){}
            doMark('http', 1);
        };
        imghttps.src = 'https://qzonestyle.gtimg.cn/ac/b.gif?_r='+Math.random();
        imghttp.src = 'http://qzonestyle.gtimg.cn/ac/b.gif?_r='+Math.random();
    }


})();


(function(){var _i=new Image(),_r=Math.random(),_re = new RegExp("(?:^|;+|\\s+)_qz_fix=([^;]*)"),m=document.cookie.match(_re),_v=!m ?"":m[1],_ph='//pinghot.qq.com/pingd?dm=homeact.qzone.qq.com.hot&url=/troubleshooter&tt=-&hottag=guanjia.qzonestyle.',_phe='&hotx=9999&hoty=9999&rand='+_r;_i.onerror=function(){var _t=new Image();if(_v>0){document.getElementById('trs_tip').innerHTML='如果您看到这个提示，说明QQ空间仍然无法正常打开，请您返回空间小助手，进行<a href="http://user.qzone.qq.com/troubleshooter#secondfix" title="空间小助手"><strong>深度修复</strong></a>';_t.src=_ph+'badagain'+_phe;}else{_t.src=_ph+'bad'+_phe;}var _d = document.createElement("div");_d.innerHTML = '<iframe frameborder="0" style="width:1px;height:1px;" src="http://user.qzone.qq.com/troubleshooter?traytip"></iframe>';document.body.appendChild(_d);};_i.onload=function(){if(_v==2){var _t=new Image();_t.src=_ph+'fixed'+_phe;}};_i.src='//'+siDomain+'/ac/qzone_v5/client/year_vip_icon.png?i='+_r;})();
QZFL.event.onDomReady(function(){
    g_T.fwp[3] = new Date();
});



    var fsimg = document.getElementById("full_background_img");
    if(fsimg){
        window.adjust_full_screen_skin = function(){
            var bg = document.getElementById('full_background'),
                    bg_img=document.getElementById('full_background_img'),
                    cw = QZFL.dom.getClientWidth(),
                    ch = QZFL.dom.getClientHeight(),
                    iw = bg_img.width,
                    ih = bg_img.height;

                bg.style.width=cw+"px";
                bg.style.height=ch+"px";

                if (cw / ch > iw / ih) {
                    var new_h = cw * ih / iw,
                        imgTop = (ch - new_h) / 2;
                    bg_img.style.width=cw + "px";
                    bg_img.style.height=new_h + "px";
                    bg_img.style.top="";
                    bg_img.style.left="";
                }else {
                    var new_w = ch * iw / ih,
                        imgLeft = (cw - new_w) / 2;
                    bg_img.style.width=new_w + "px";
                    bg_img.style.height=ch + "px";
                    bg_img.style.left=imgLeft + "px";
                    bg_img.style.top="";
                }
        };
        if (!QZONE.shop || !QZONE.shop.initDressUp) {
            var url = fsimg.getAttribute("data-src");
            fsimg.onload = adjust_full_screen_skin;
            fsimg.src = url;
        }
    }

    var mc = document.getElementById('menuContainer');
    var oldGuideMap = {
        N2: "myhome",
        3: "305",
        5: "4",
        6: "306",
        7: "1",
        8: "3",
        10: "311",
        11: "323",
        13: "338",
        14: "myhome",
        15: "yellowgrade",
        17: "333",
        18: "340"
    };
    var isDefaultMenu = '123054' == 1,
        isCustom = '1' & 2 > 0,
        isMobile = (window.ua && (ua.isiPhone || ua.isiPad)),
        lst;
    window.getHotTagPrefix = function(){
            var env = window.g_app_identifier != "" ? 'app' : (window.g_isOFP == "0" ? 'home' : 'center');

            return env + 'click' +  (window.ownermode ? "-host" : "-guest");
        };
    window.custom_menu_swf = function(id) {
        if(QZFL.userAgent.chrome || QZFL.userAgent.safari){
            var _cbi = document.getElementById('_chrome_bug_input');
            if(!_cbi){
                var _laybg = document.getElementById('layBackground');
                QZFL.dom.createElementIn("input",_laybg,false,{id:"_chrome_bug_input",type:'text',style:"border:1px;width:1px;position:absolute;left:-10px;"});
                _cbi =  document.getElementById('_chrome_bug_input');
            }
            _cbi.focus();
        }

        if (id != 'close') {
            id = (!isCustom && !isDefaultMenu && !isMobile) ? (oldGuideMap[id] || id) : id;
            if(id == "-1" || id == "N1"){
                id = "main";
            }
            if (id == '5' || id == '7') {
                id = oldGuideMap[id];
            }
            var _p = getHotTagPrefix();
            window.TCISD && TCISD.hotClick("menu_v8."+_p+".appid_"+id, "8.qzone.qq.com");
            if(id == '219'){
                if(QZONE.music && typeof(QZONE.music.switchPage)=="function"){
                    QZONE.music.switchPage();
                }
            }else{
                QZONE.FrontPage.toApp("/" + id);
            }
        }
        return false;
    };
    if (!QZONE.shop || !QZONE.shop.initDressUp) {
        mc && QZFL.event.delegate(mc, 'li', 'click', function (e) {
            QZFL.event.preventDefault(e);
            this.className = this.className.replace(/_\-1/, '_N1');
            var appid = this.className.match(/menu_item_([A-Z0-9\-]*)/i);
            appid && appid[1] && custom_menu_swf(appid[1]);
            if (lst) {
                QZFL.css.removeClassName(lst, 'cur');
            }
            lst = this;
            QZFL.css.addClassName(lst, 'cur');
        });
        mc && mc.setAttribute('reged', '1');
    }

(function(){try{if(parent!=self && (parent.document.domain!=document.domain || (document.referrer && !/^http(s)?:\/\/[.\w-]+\.qq\.com\//i.test(document.referrer)))){throw new Error("can't be iframed");}}catch(e){window.open(location.href, "_top");}})();
    /**
     * β 增加新接口 获取通过URL传递的参数值
     * @param {String} name 参数的名称
     */
    var getParameter = function(url, name){
        var r = new RegExp("(\\?|#|&)" + name + "=([^&#]*)(&|#|$)"),
            m = url.match(r);
        return (!m ? "" : m[2]);
    };
    //上报用iframe嵌套qzone的用户
    var qq = getParameter(location.href, 'qzoneInIframe');
    if(qq) {
        TCISD.stringStat(1000100, {
                bid: 108146,
                rc: 'QQ:' + qq
        }, {
            reportRate: 1
        });
    }
    var clickReport = getParameter(location.href, 'clickReport');
    if(window.TCISD && window.TCISD.pv) {
        TCISD.pv('user.qzone.qq.com', clickReport);
    }






</script>





<script type="text/javascript">
window.g_qzonetoken = (function(){ try{return "3a5a499b288365c95ef3d236aab3859c1412021e15f19ab03e9e7ccdbb521417992e7df7578386aad6";} catch(e) {var xhr = new XMLHttpRequest();xhr.withCredentials = true;xhr.open('post', '//h5.qzone.qq.com/log/post/error/qzonetoken', true);xhr.setRequestHeader('Content-Type', 'application/x-www-form-urlencoded');xhr.send(e);}})();

/**
 * 框架机模调上报
 */
var reportMd = function (code, type, delay) {
	//使用模调测速
	var url = "//h5.qzone.qq.com/report/md?"
		+ "fromId=204971707&toId=211006088&interfaceId=111900240"
		+ (typeof code == 'undefined' ? '' : ("&code=" + code))
		+ (typeof type == 'undefined' ? '' : ("&type=" + type))
		+ (typeof delay == 'undefined' ? '' : ("&delay=" + Math.min(delay, 10 * 1000)))
		+ '&r=' + Math.random();
	var img = new Image();
	img.src = url;
};

reportMd(0, 0, 1);
</script>



<iframe id="userData_iframe___share_db" src="https://qzs.qq.com/ac/qzfl/release/resource/html/storage_helper.html" style="display: none;"></iframe><div id="h5audio_media_con" style="height: 0px; overflow: hidden;"><audio id="h5audio_media" height="0" width="0" autoplay="false"></audio></div><div style="display: none;"><div><div x:id="content_content" spellcheck="false" x:idprefix="content" class="textinput textarea c_tx2" contenteditable="true" accesskey="q" style="display:none;" id=""></div><div x:id="substitutor_content" x:idprefix="substitutor" class="textinput textarea c_tx3" contenteditable="true" tabindex="0" accesskey="q" id=""></div><div x:id="tips" class="mod_at_wrap" style="display:none;" id=""><div class="mod_at bg bor2"><div class="mod_at_list"><div x:id="tips_content" class="mod_at_tips c_tx3 bg3" id="">正在加载中...</div></div></div></div><div x:id="options" id=""></div></div><div x:id="content" contenteditable="true" tabindex="0"></div><li><img x:id="1"><span x:id="nickname">#{memo:1,nickname:0,uin:0,organizationName:0;conv:F4A.utils.getNameForShow}</span></li><div x:id="content" contenteditable="true" class="richTextBox"></div><li class="comments_item bor3" style="display: none;"><div class="comments_item_bd"><div><div class="ui_avatar"><a target="_blank" style="display: none;"><img></a></div><div class="comments_content"><span class="comments-content-publish"><a class="nickname" target="_blank">null</a><span style="display: none;"> 回复 <a class="nickname" target="_blank" style="display: none;">null</a></span><span class="private-txt"></span> : </span><span class="comments-content-detail"></span><div class="none"><div></div></div><div class="comments_op"><span class="c_tx3 ui_mr10"></span><span class="c_tx3 ui_mr10" style="display: none;">undefined</span><a href="javascript:void(0);" class="mod_comment_reply c_tx"></a><span style="display: none;"></span><a href="javascript:;" class="c_tx mod_comment_del" style="display: none;">删除</a><span style="display: none;"></span></div><div class="photo-tooltip photo-tooltip-bottom" style="display: none;"><div class="photo-tooltip-inner">点击回复</div><span class="photo-tooltip-arrow"></span></div></div></div><div class="mod_comments_sub" style="padding: 0px;"><ol style="display: none;"></ol></div></div></li></div><div style="display: none;"></div><div id="_qz_zoom_detect" style="position: absolute; right: 0px; bottom: 0px; visibility: visible;"><object data="//qzonestyle.gtimg.cn/qzone/v6/accessory/plugin/zoom.swf" width="10" height="10" id="accessory_zoom" name="zoom_detect" type="application/x-shockwave-flash"><param name="allowScriptAccess" value="always"><param name="wmode" value="transparent"><param name="scale" value="noScale"></object></div><iframe id="userData_iframe__Y_popTips" src="https://qzs.qq.com/ac/qzfl/release/resource/html/storage_helper.html" style="display: none;"></iframe><div id="fixLayout" style="width: 100%;"><div id="qz_notification" class="msg-channel-wrapper none" style="right: 1px; z-index: 6500; width: 321px; position: fixed; bottom: -200px;">
    <div class="msg-channel">
        <div class="hd">
            <p class="user-info">
                <span class="welcome">晚上好</span>
                <span class="name textoverflow">承諾，/aiq一世的誓言</span><a href="http://pay.qq.com/ipay/index.shtml?n=3&amp;c=xxjzghh,xxjzgw&amp;aid=mingpian_icon&amp;ch=qdqb,kj" target="_blank" title="点击查看黄钻特权详情" class="qz_vip_icon_s qz_vip_icon_s_gray_1 "></a>            </p>

        </div>
        <div class="bd">
            <!-- 整片可点击 -->
            <a href="http://rc.qzone.qq.com/1106217226?via=QZ.HUANGZUANTIPS&amp;app_custom=essence" target="_blank" data-hook="reportClick" class="inner clearfix" data-via="1106217226_QZ.HUANGZUANTIPS">
                <img src="http://qzonestyle.gtimg.cn/qzone/space_item/boss_pic/2331_2017_9/1505900232641_904078.gif" class="privilege-img" alt="游戏特权推广图">

                <div class="privilege-content">
                    <h3 class="title">恭喜你</h3>

                    <p class="desp">获得限量礼包一份，点击领取&gt;&gt;&gt;</p>
                    <span class="btn-get" title="点击领取">点击领取</span>

                </div>
            </a>
        </div>
    </div>
    <!-- 操作按钮 -->
    <a class="op-icon icon-setting" data-hook="showSetup" href="javascript:void(0)" title="弹窗设置"></a>
    <a class="op-icon icon-close" data-hook="closeTips" href="javascript:void(0)" title="关闭"></a>
    <!-- 分割线 -->
    <span class="op-icon-split"></span>

    <!-- 浮层 -->
    <div class="ui-dialog ui-light-dialog ui-light-dialog-s ui-popover ui-popover-top-right" id="qz_notification_config" style="display:none">
        <div class="ui-dl-inner">
            <b class="ui-dl-arrow ui-dl-arrow1"></b>
            <b class="ui-dl-arrow ui-dl-arrow2"></b>

            <div class="ui-dl-hd">
                <div class="qz-grid">
                    <div class="qz-right"><a href="javascript:void(0);" data-hook="shutDown" class="close" title="关闭"><s class="ui-icon sp-dialog-close"></s></a></div>
                </div>
            </div>
            <div class="ui-dl-bd">
                <p class="j_qz_notification_config_select" data-hook="selectThis" data-index="0"><label for=""><input type="radio" name="show_mod" class="none" value="0"><i class="ui-icon qz-radio-on"></i>每日登录时自动显示一次</label></p>

                <p class="j_qz_notification_config_select" data-hook="selectThis" data-index="1"><label for=""><input type="radio" name="show_mod" class="none" value="1"><i class="ui-icon qz-radio"></i>从不自动显示</label>
                </p>
            </div>
            <div class="ui-dl-ft">
                <div class="qz-grid">
                    <div class="qz-right actions">
                        <a class="qz-button qz-dark-button" data-hook="sendOk" href="javascript:void(0);"><span class="txt">确定</span></a>
                    </div>
                </div>
            </div>
        </div>
    </div></div></div><div id="qzNameCardTips" style="position: absolute; top: -500px; left: 755px; background-color: white; z-index: 1023; visibility: hidden;"><div id="qzNameCardLoad" class="qzone-card-loading"><div class="loading-wrap">好友名片加载失败</div></div></div><div id="qzNameCardDiv" style="position: absolute; top: -300px; left: 379px; background-color: white; z-index: 1024; visibility: hidden; opacity: 1;"><div class="qzone-cards  "><div id="_qzone_cards" class="qzone-cards-wrap"><div class="cards-abstract"><div class="card-user-pic"><a href="http://user.qzone.qq.com/424500179" target="_blank" data-click="||0|common_namecard.picture|"><img alt="是梦，终会醒" src="http://qlogo4.store.qq.com/qzone/424500179/424500179/100?1479039109" title="访问是梦，终会醒的空间" class="photo"></a></div><div class="card-main"><div class="card-user-info" id="nc_cardUserInfo"><a href="http://user.qzone.qq.com/424500179" class="nickname ui-mr5 textoverflow" id="nc_userNickname" title="访问是梦，终会醒的空间" target="_blank" data-click="||0|common_namecard.nick|"><span>是梦，终会醒</span></a><a href="http://pay.qq.com/ipay/index.shtml?n=3&amp;c=xxjzghh,xxjzgw&amp;aid=mingpian_icon&amp;ch=qdqb,kj" target="_blank" data-click="||0|common_namecard.yellow|" title="点击查看黄钻特权详情" class="qz_vip_icon_s qz_vip_icon_s_gray_1 nickname_ico ui_mr5"></a></div><div class="card-user-basicinfo"><p>男 白羊座 内江</p></div><div class="card-user-info-more"><div class="card-user-info-more">你们有<b class="ui-b">1</b>个共同好友</div></div></div></div><div class="qzone-cards-op clearfix" id="_namecard_feed"><div class="op-list"><div id="_namecard_op" class="op-right"><div class="right-inner"><a href="javascript:;" class="expand"><i class="icon-unfold"></i></a><ul><li><a href="javascript:;" data-click="sendGift|424500179-&amp;type=556&amp;birthdaytab=&amp;lunarFlag=undefined&amp;source=undefined&amp;nick=|1|common_namecard.gift|">赠送礼物</a></li><li><a href="javascript:;" data-click="visitAction|424500179|1|common_namecard.visit|">访问设置</a></li><li><a href="javascript:;" data-click="openReportbox|424500179|1|common_namecard.police|">举报</a></li></ul></div></div><div class="op-item" style="display:none"><div class="right-inner"><a href="javascript:;" data-click="openChatbox|424500179|1|common_namecard.chat|" class="expand">聊天</a></div></div><div class="op-item"><div class="right-inner"><a href="javascript:;" data-click="addFriend|424500179|1|common_namecard.addfriend|" class="expand"><span id="applyforfriendNamecard">加为好友</span></a></div></div></div></div></div></div></div><div class="photo_layer js-viewer-container" id="js-viewer-container" style="z-index: 5500; position: fixed; top: 0px; height: 987px; display: none; padding-top: 16px;" data-comment-insert-img-state="inited">
 <div id="js-viewer-layer" class="photo_bg_layer js-viewer-layer" onclick="javascript:if(slide.util.getParameter('inqq')){return;}try{slide&amp;&amp;slide.beforeClose();}catch(err){}var ctn = document.getElementById('js-viewer-container');ctn.style.display='none';slide&amp;&amp;slide.close({noTriggerBeforeClose:true});" style="height: 987px;"></div>
 <div class="screen_handle_tab" style="display:none">
  <p class="comment" title="评论"><i class="icon icon_comment_bphoto"><i></i></i></p>
  <p class="reproduced" title="转载"><i class="icon icon_reproduced_bphoto"><i></i></i></p>
  <p class="share" title="分享"><i class="icon icon_share_bphoto"><i></i></i></p>
 </div>
 <!--图片区S-->
 <div class="photo_figure" id="js-viewer-figure" style="margin-top: 0px; margin-left: 0px;">
  <div class="photo_figure_main" id="js-viewer-main" style="width: 1148px; height: 955px;">
   <div class="photo_layer_close" data-toggle="tooltip" title="关闭" onclick="javascript:try{slide&amp;&amp;slide.beforeClose();}catch(err){}var ctn = document.getElementById('js-viewer-container');ctn.style.display='none';slide&amp;&amp;slide.close({noTriggerBeforeClose:true});" ontouchend="javascript:try{slide&amp;&amp;slide.beforeClose();}catch(err){}var ctn = document.getElementById('js-viewer-container');ctn.style.display='none';slide&amp;&amp;slide.close({noTriggerBeforeClose:true});return false;">
       <a href="javascript:void(0);" onclick="return false;">关闭</a>
   </div>
   <div id="js-hdmode-close" class="photo_layer_close" title="退出全屏模式" style="display:none;"><a href="javascript:;">关闭</a></div>

   <!--图片展示区S-->
   <div class="figure_area" id="js-viewer-imgWraper" style="width: 828px; margin-top: 0px; height: 895px;">
    <input type="text" style="position:absolute;top:0px;left:0px;width:1px;height:1px;font-size:0;border:1px;outline:none;opacity:0.1;filter:alpha(opacity=10)" value="" id="js-focus-input">
    <p class="HD" style="z-index:500;display:none;"><a href="javascript:;" target="_blank">查看原图<span id="js-hd-size"></span></a></p>
    <a href="javascript:void(0);" id="js-link-hd" class="show-original" style="display: none; right: 14px; top: 14px;" canbeshow="0">查看原图</a>
        <div class="figure_img_loading" id="js-img-loading" style="z-index: 6; position: absolute; display: none;"></div>    <span class="video-loading js-video-singletip js-video-loading" style="display: none;"><i class="inner"></i></span><div class="figure_img" id="js-image-ctn" style="-moz-user-select:none;position:relative;" onselectstart="return false;" hidefocus="true">
     <img src="http://qzonestyle.gtimg.cn/ac/b.gif" id="js-img-disp" style="position: absolute; opacity: 1; left: 0px; top: 0px; width: 0px; height: 0px; z-index: 3; display: none;" hidefocus="true" has-show="">
     <img src="http://qzonestyle.gtimg.cn/ac/b.gif" id="js-img-trans" style="display: none; position: absolute; left: 0px; top: 0px; width: 0px; height: 0px;" hidefocus="true">
    <div id="js-img-border" style="position: absolute; z-index: 4; height: 290px; width: 482px; left: 174px; top: 303px;" class="figure_img_bor"><img src="http://qzonestyle.gtimg.cn/ac/b.gif" id="" style="opacity: 0; display: none; width: 0px; height: 0px; left: 0px; top: 0px;" hidefocus="true" has-show="1"></div><div class="js-video-flash-ctn" id="videoViewerCtn" style="position: absolute; top: -3px; left: -3px; width: 3px; height: 3px; z-index: 10; display: none;"><object id="videoViewer" name="videoViewer" width="100%" height="100%" data="//qzs.qq.com/qzone/client/photo/swf/MPlayer/MicroVideoPlayerEx.swf" type="application/x-shockwave-flash" data-index=""><param name="wmode" value="opaque"><param name="allowScriptAccess" value="always"><param name="allowfullscreen" value="true"><param name="allownetworking" value="all"><param name="flashvars" value="noloading=1&amp;onFlashInited=slide._plugin_video._newVideoCallbacks.onFlashInited&amp;onMetaData=slide._plugin_video._newVideoCallbacks.onMetaData&amp;onError=slide._plugin_video._newVideoCallbacks.onError&amp;onChangeState=slide._plugin_video._newVideoCallbacks.onChangeState&amp;onPlayStart=slide._plugin_video._newVideoCallbacks.onPlayStart&amp;onPlayStop=slide._plugin_video._newVideoCallbacks.onPlayStop&amp;onSlideStart=slide._plugin_video._newVideoCallbacks.onSlideStart&amp;onSlideStop=slide._plugin_video._newVideoCallbacks.onSlideStop&amp;onChangeFull=slide._plugin_video._newVideoCallbacks.onChangeFull"></object></div></div>
    <!--TODO: -->
    <div class="mod-figure-area" id="js-figure-area" style="display: none;">
     <div id="js-ctn-infoBar" class="figure-desc figure-desc-empty" style="display: none;"></div>
     <div class="figure-handle">
      <ul>

       <li style="display: none;">
        <a id="js-btn-changeMode" style="display:none" href="javascript:;" class="icon-wrap js-show-normal func-zoom" title="点击放大"><i class="icon-m icon-magnify"></i></a>
       </li>

       <li>
        <a id="js-btn-rotateRight" href="javascript:void(0)" title="旋转" class="icon-wrap func-zoom" style="display: none;"><i class="icon-m icon-rotate"></i></a>
       </li>
       <li>
        <div class="photo-view-mode">
         <a href="javascript:void(0)" class="js-normal-mode mode-normal" title="小图模式"><i class="icon-m icon-small-view">小图模式</i></a>
         <a href="javascript:void(0)" class="js-large-mode js-large-button mode-large" title="大图模式"><i class="icon-m icon-big-view">大图模式</i></a>
         <a href="javascript:void(0)" class="js-large-mode js-hd-button mode-hd" title="大图模式" style="display: none;"><i class="icon-m icon-hd-view">高清模式</i></a>
        </div>
       </li>

       <li>
        <div class="photo-hd-mode photo-hd-all" style="display:none;">
         <a href="javascript:void(0)" class="func-more js-large-mode" title="大图模式"><i class="icon-m icon-hd-m">HD</i></a>
        </div>
       </li>

       <li>
        <div class="photo-hd-mode">
         <a id="js-btn-fullscreen" style="display: none;" href="javascript:void(0)" class="func-more" title="幻灯片播放"><i class="icon-m icon-full-view"></i>播放</a>
        </div>
       </li>


      </ul>
     </div>
     <div class="figure-area-mask"></div>
    </div>
    <div id="js-map-ctn" class="figure_img_map" style="display:none;">
     <img id="js-img-map" src="http://qzonestyle.gtimg.cn/ac/b.gif" style="display:none">
     <p id="js-map-handler" class="visible_area" style="display:none;cursor:move;"></p>
    </div>
    <div id="_slideView_scale_num" class="scale_num" style="display:none"><p>150%</p></div>
    <div class="photo-save-tip" id="js-btn-saveRotate" style="display:none">
                    <div class="save-tip-cont">
                        <p> 是否保存旋转后的照片？</p>
                        <div class="save-tip-op">
                            <a href="javascript:;" class="save-tip-select js-save-rotate-ok">保存</a>
                            <a href="javascript:;">取消</a>
                        </div>
                    </div>
                </div>
    <a id="js-btn-prevPhoto" href="javascript:;" class="js-btn-changePhoto figure-area-arrow arrow-pre " style="top: 45%;">上一张</a>
    <a id="js-btn-nextPhoto" href="javascript:;" class="js-btn-changePhoto figure-area-arrow arrow-next" style="top: 45%;">下一张</a>
    <a href="javascript:void(0)" style="z-index:100;display:none;" id="js-btn-play-gif" class="play-the-video" title="播放"></a>
   </div>
   <!--图片展示区e-->
   <!--图片右侧详情区S-->
   <div class="figure-side" id="js-sidebar-ctn" style="padding-bottom: 30px; height: 925px; overflow: hidden;">
    <div class="figure-side-wrap">
        <div class="figure-side-bg" id="js-cmt-wrap" style="">

      <div class="figure-side-inner">
       <!--照片所属信息区S-->
       <div class="figure-side-hd">
                        <div class="js-userinfo-ctn"></div>

                                 <!--照片编辑区S-->
                                <div class="info-edit" id="js-desc-editor" style="display:none;">
                                    <div class="tit-edit-warp textinput js-title-editor">
                                        <div class="tit-edit">
                                            <input type="text" class="js-desc-title">
                                            <span class="watermark">照片标题</span>
                                            <span class="num-count"><span class="num-hint js-desc-title-currword">30</span>/30</span>
                                        </div>
                                    </div>
                                    <div class="desc-edit js-desc-cont">
                                        <span class="num-count"><span class="num-hint js-desc-currword">0</span>/200</span>
                                    </div>
                                    <div class="info-submit">
                                        <a href="javascript:void(0)" class="info-confirm js-desc-ok">确定</a>
                                        <a href="javascript:void(0)" class="info-cancel js-desc-cancel">取消</a>
                                    </div>
                                </div>
       </div>


       <!--照片评论区S-->
       <div class="handle-tab" style="">
                                <ul>
                                    <li id="js-viewer-like" class="praise" title="赞" style="display: none;"></li>
                                    <li>
                                        <a href="javascript:;" class="handle-item" id="js-viewer-comment" title="评论" style="display: none;"><i class="icon-m icon-comment-m">评论</i><span class="btn-txt">评论</span><span class="btn-txt-num"></span></a>
                                    </li>
                                    <li id="js-interactive-btn" js-timer="1921">
                                        <a href="javascript:;" class="handle-item"><i class="icon-m icon-reproduced-m">互动</i><span class="btn-txt"></span><span class="btn-txt-num"></span></a>
                                    </li>
                                    <li id="js-othermenu-btn" class="more" js-timer="1922">
                                        <a href="javascript:;" class="handle-item"><i class="icon-m icon-other-m">其他</i></a>
                                    </li>
                                </ul>
                            </div>
       <p class="figure_praise_num js_fade_like" id="_slideView_like"></p>
       <div class="figure-interactive">
        <p class="figure_praise_num" id="js-like-list" style="display: none;"></p>
        <div class="comment-tab" style="display:none" id="j-comment-tab">
         <a href="javascript:;" class="tab-selected" data-type="friend">好友评论</a>
         <a href="javascript:;" data-type="cmtreply">精选评论</a>
        </div>
        <div class="handle_main js_show_comment" id="js-comment-ctn" style="visibility: visible;">
         <div class="js_mod_comment_module" id="js-comment-module"></div>
         <div class="js_mod_retweet" id="js-mod-retweet">

         </div>
        </div>
       </div>
       <!--我评论区E-->
       <!--照片评论区E-->
      </div>

     </div>
     <div class="figure-comment js-can-comment" style="display: block;"></div>
    </div>
    <div id="js-qq-ad" data-inqq="1" style="left: 20px; position: absolute; height: 110px; width: 260px; overflow: hidden; bottom: 55px;"><p style="margin: 0; padding: 0; color: #262626; font-size: 12px; font-weight: bold;">精彩生活<span style="float:right;color: #999;font-weight:bold;">广告</span></p><p style="border-bottom: 1px solid #f0f0f0;height: 0;font-size: 0; margin: 8px 0 10px; padding: 0;"></p><div style="height: 75px; margin-bottom: 14px;"><a onclick="return GDT._click( '8935142765332649167', '36428148', {el: this} );" onmousedown="GDT. ClickManage.mousedown( '8935142765332649167', '36428148', {el: this} );" onmouseup="GDT. ClickManage.mouseup( '8935142765332649167', '36428148', {el: this} );" onmouseover="GDT. ClickManage.mouseover( '8935142765332649167', '36428148', {el: this} );" href="http://c.gdt.qq.com/gdt_click.fcg?viewid=4iIfdXN8o937VlMJVgGzyDvl9wwvSAa73IVuVAlgOzihSyqA2VhUcHkzQn4F6n5vL8foYGWAugvRsD57C9OaY0Lbb!NaT6OIAmk562eVqMpxHE_e69k7D1JcGU8WwvTA&amp;jtype=0&amp;i=1&amp;os=0" gdtoriurl="http://c.gdt.qq.com/gdt_click.fcg?viewid=4iIfdXN8o937VlMJVgGzyDvl9wwvSAa73IVuVAlgOzihSyqA2VhUcHkzQn4F6n5vL8foYGWAugvRsD57C9OaY0Lbb!NaT6OIAmk562eVqMpxHE_e69k7D1JcGU8WwvTA&amp;jtype=0&amp;i=1&amp;os=0" target="_blank"><img offset="2" class="js_ad_rsp js_ad_img" src="https://pgdt.gtimg.cn/gdt/0/DAATsDUABLABLAANBZw4BAB3NYi0bA.jpg/0?ck=c063365bb748d317ec3abc129845a9fa" alt="图片加载失败" style="float: left; width: 75px; height: 75px; margin: 0; padding: 0 12px 0 0; border: 0 none; cursor: pointer;"></a><div style="overflow: hidden; position: relative;height: 75px"><p class="js_ad_rsp" style="margin: 0; padding: 0; font-weight: bold; font-size: 12px; cursor: pointer; overflow:hidden; text-overflow: ellipsis; white-space: nowrap;"><a style="color: #262626;" onclick="return GDT._click( '8935142765332649167', '36428148', {el: this} );" onmousedown="GDT. ClickManage.mousedown( '8935142765332649167', '36428148', {el: this} );" onmouseup="GDT. ClickManage.mouseup( '8935142765332649167', '36428148', {el: this} );" onmouseover="GDT. ClickManage.mouseover( '8935142765332649167', '36428148', {el: this} );" href="http://c.gdt.qq.com/gdt_click.fcg?viewid=4iIfdXN8o937VlMJVgGzyDvl9wwvSAa73IVuVAlgOzihSyqA2VhUcHkzQn4F6n5vL8foYGWAugvRsD57C9OaY0Lbb!NaT6OIAmk562eVqMpxHE_e69k7D1JcGU8WwvTA&amp;jtype=0&amp;i=1&amp;os=0" gdtoriurl="http://c.gdt.qq.com/gdt_click.fcg?viewid=4iIfdXN8o937VlMJVgGzyDvl9wwvSAa73IVuVAlgOzihSyqA2VhUcHkzQn4F6n5vL8foYGWAugvRsD57C9OaY0Lbb!NaT6OIAmk562eVqMpxHE_e69k7D1JcGU8WwvTA&amp;jtype=0&amp;i=1&amp;os=0" target="_blank">大牌设计 全屋环保</a></p><p class="js_ad_rsp" style="margin: 4px 0 8px; padding: 0;overflow: hidden;line-height: 16px; height: 32px;"><a onclick="return GDT._click( '8935142765332649167', '36428148', {el: this} );" onmousedown="GDT. ClickManage.mousedown( '8935142765332649167', '36428148', {el: this} );" onmouseup="GDT. ClickManage.mouseup( '8935142765332649167', '36428148', {el: this} );" onmouseover="GDT. ClickManage.mouseover( '8935142765332649167', '36428148', {el: this} );" href="http://c.gdt.qq.com/gdt_click.fcg?viewid=4iIfdXN8o937VlMJVgGzyDvl9wwvSAa73IVuVAlgOzihSyqA2VhUcHkzQn4F6n5vL8foYGWAugvRsD57C9OaY0Lbb!NaT6OIAmk562eVqMpxHE_e69k7D1JcGU8WwvTA&amp;jtype=0&amp;i=1&amp;os=0" gdtoriurl="http://c.gdt.qq.com/gdt_click.fcg?viewid=4iIfdXN8o937VlMJVgGzyDvl9wwvSAa73IVuVAlgOzihSyqA2VhUcHkzQn4F6n5vL8foYGWAugvRsD57C9OaY0Lbb!NaT6OIAmk562eVqMpxHE_e69k7D1JcGU8WwvTA&amp;jtype=0&amp;i=1&amp;os=0" target="_blank" style="color: #7e7e7e;font-size: 12px;">全屋环保 大牌设计 髙性价比</a></p><a onclick="return GDT._click( '8935142765332649167', '36428148', {el: this} );" onmousedown="GDT. ClickManage.mousedown( '8935142765332649167', '36428148', {el: this} );" onmouseup="GDT. ClickManage.mouseup( '8935142765332649167', '36428148', {el: this} );" onmouseover="GDT. ClickManage.mouseover( '8935142765332649167', '36428148', {el: this} );" href="http://c.gdt.qq.com/gdt_click.fcg?viewid=4iIfdXN8o937VlMJVgGzyDvl9wwvSAa73IVuVAlgOzihSyqA2VhUcHkzQn4F6n5vL8foYGWAugvRsD57C9OaY0Lbb!NaT6OIAmk562eVqMpxHE_e69k7D1JcGU8WwvTA&amp;jtype=0&amp;i=1&amp;os=0" gdtoriurl="http://c.gdt.qq.com/gdt_click.fcg?viewid=4iIfdXN8o937VlMJVgGzyDvl9wwvSAa73IVuVAlgOzihSyqA2VhUcHkzQn4F6n5vL8foYGWAugvRsD57C9OaY0Lbb!NaT6OIAmk562eVqMpxHE_e69k7D1JcGU8WwvTA&amp;jtype=0&amp;i=1&amp;os=0" target="_blank" class="js_ad_rsp js_ad_btn" style="position:absolute;bottom: 0; left:0; height:20px; line-height:20px; text-decoration: none; cursor:  pointer; margin: 0; padding: 0 4px; border-radius: 3px; color: #fff; background: #5CAAE6; background-image: none; border: 1px solid #59A7E3; font-size: 12px;">去看看</a></div></div></div>
    <!--圈人推荐-->
    <div id="js-face-area" class="figure-side-circle" style="display:none;">
    </div>
    <div class="friend-list-wrap j-selector-wrap" style="width:170px; top:-999px; left:112px; display:none;">
    </div>
    <!--互动菜单-->
    <div class="mod-layer-drop" id="js-interactive-menu" style="left: 150px; position: absolute; top: 107px; z-index: 99; display: none;">
                    <ul>
                        <li style="display: none;">
                            <a href="javascript:;" id="js-viewer-reprint">转载</a>
                        </li>
                        <li style="display:none">
                            <a href="javascript:;" id="js-viewer-retweet" class="retweet js_retweet" title="转发">转发</a>
                        </li>
                        <li style="">
                            <a href="javascript:;" id="js-btn-sharePhoto">分享</a>
                        </li>
                        <li style="">
                            <a href="javascript:;" id="js-btn-copyAddress">复制地址</a>
                        </li>
                    </ul>
                    <span class="mod-arr mod-arr-t"><span></span></span>
                </div>
                <!--其他菜单-->
    <div class="mod-layer-drop" id="js-other-menu" style="left: 222.5px; position: absolute; top: 107px; z-index: 99; display: none;">
                    <ul>
                     <li id="js-btn-open-quanren" class="js-hide-when-video" style="display:none;" data-orivisible="-1"> <a class="" id="mark_btn" href="javascript:void(0)" style="display: ;" title="圈人">圈人</a></li>
                     <li id="js-btn-cover-li" style="display:none;">
                      <a href="javascript:;" id="js-btn-cover">设为封面</a>
                     </li>
                     <li id="js-btn-qzone-cover-li" style="display:none;">
                      <a href="javascript:;" id="js-btn-qzone-cover">主页展示</a>
                     </li>
                     <li id="js-btn-movePhoto-li" style="display:none;">
                      <a href="javascript:;" id="js-btn-movePhoto">移动</a>
                     </li>
                     <li id="js-btn-downloadPhoto-li">
                      <a href="javascript:;" id="js-btn-downloadPhoto" data-downloadtype="video">下载视频</a>
                      <!--<div class="more-download" style="display:none">
        <ul>
         <li><a href="javascript:void(0)" id="js-btn-downloadNormalInViewer">大图 </a></li>
         <li><a href="javascript:void(0)" id="js-btn-downloadOrigin">原图 </a></li>
         <li><a href="javascript:void(0)" id="js-btn-downloadHighClear">高清图 </a></li>
        </ul>
        <i class="more-download-arr"></i>
       </div>-->
                     </li>
                     <li id="js-btn-collect-li" style="display: none;">
                      <a href="javascript:;" id="js-btn-collect">收藏</a>
                     </li>
                     <li id="js-btn-delPhoto-li" style="display:none;">
                      <a href="javascript:;" id="js-btn-delPhoto">删除</a>
                     </li>
                     <li id="js-btn-meihua-li" class="js-hide-when-video" style="display:none;" data-orivisible="-1">
                      <span class="drop-item-seprate"></span>
                      <a href="javascript:;" id="js-btn-meihua">美化</a>
                     </li>
                     <li id="js-btn-tuya-li" class="js-hide-when-video" style="display: none;" data-orivisible="1">
                      <a href="javascript:;" id="js-btn-tuya">涂鸦</a>
                     </li>
                     <li id="js-btn-meituxiuxiu-li" class="js-hide-when-video" style="display:none;" data-orivisible="-1">
                      <a href="javascript:;" id="js-btn-meituxiuxiu" class="open-meitu">美图秀秀</a>
                     </li>
                     <li id="js-btn-follow-li" style="display:none;">
                      <a href="javascript:;" id="js-btn-follow" class="open-meitu"></a>
                     </li>
                    </ul>
                    <span class="mod-arr mod-arr-t"><span></span></span>
                </div>
    <div class="figure-side-ft" id="js-sidebar-foot">

     <a class="js-report-btn op-pic-policy" href="javascript: void(0);"><i></i>举报</a>      <a href="http://support.qq.com/write.shtml?fid=944&amp;SSTAG=slideversion%3Dviewer2" class="suggest-support" target="_blank"><i></i>建议反馈</a>    </div>
   </div>
   <!--图片右侧详情区e-->
  </div>
 </div>
 <!--图片区e-->
 <div id="js-ctn-switch" class="photo_minimap_switch" style="display: none;">
  <div id="js-switch-inner" class="inner" style="top: 0px;">
   <div id="js-thumb-unexpand" class="photo_minimap_fold"><a href="javascript:;" class="switch"><b class="ui_trig ui_trig_b"></b><span class="txt">收起</span></a></div>
   <div id="js-thumb-expand" class="photo_minimap_unfold"><a href="javascript:;" class="switch"><b class="ui_trig ui_trig_t"></b></a></div>
  </div>
 </div>
 <!--缩略图滚动区S-->
 <div id="js-thumb-ctn" onselectstart="return false;" class="photo_minimap_v2" style="width: 1148px; opacity: 1;">
  <h4 id="js-thumb-title" class="mod-title" style="display: none; padding-left: 56.5px;"></h4>
  <div id="js-thumb-subctn" class="photo_minimap_inner video-list-wrap" style="width: 775px; margin-left: 36.5px; top: -55px;">
   <p id="js-thumb-prev" class="photo_minimap_roll roll_left btn-control btn-control-prev" style="visibility:hidden;"><a href="javascript:;"><span></span></a></p>
   <p id="js-thumb-next" class="photo_minimap_roll roll_right btn-control btn-control-next" style="visibility:hidden;"><a href="javascript:;"><span></span></a></p>
   <div id="js-thumbList-stage" class="photo_mini_img video-list clearfix" style="overflow: hidden; margin: 0px 30px; padding: 0px; width: 715px;">
    <ul id="js-thumbList-ctn" style="margin: 0px auto; height: 50px; width: 385px;"></ul>
   </div>
  </div>
 </div>
 <!--缩略图区E-->
<div class="comment-insert-img the_worlds_end" style="position: absolute; left: 1399px; top: 224px; z-index: 9000;"><div class="img-upload-drop cell-2"> <div class="pop-drop bg4 bor2">  <div class="arr"> <b class="in bor2"></b> <b class="out bor_bg"></b>  </div>  <ul class="list">   <li class="bor3">    <a href="javascript:void(0);" class="c_tx3" data-role="local"> <i class="icon icon-from-cmp"></i>     <span class="txt">本地</span>    </a>   </li>   <li class="bor3">    <a href="javascript:void(0);" class="c_tx3" data-hottag="add_photo_entry.album" data-role="button"> <i class="icon icon-album"></i>     <span class="txt">相册</span>    </a>   </li>  </ul> </div></div></div></div><div id="markContainer" class="" style="position: absolute; left: 396px; top: 1501px; z-index: 5500; background-image: url(&quot;about:blank&quot;); width: 0px; height: 0px; display: none;">
 <div id="markArea" class="mod-quan" style="position: absolute; width: 0px; height: 0px; top: 303px; left: 174px;"><div id="pic_marker" class="imgarea" style="position:absolute;display:none;left:0px; top:0px;width:100px; height:100px;z-index:1000;">
 <div class="imgarea-inner photo-tagging" style="width:100%; height:100%;">
  <div id="faceArea" class="photo-tagbox" style="width:100%; height:100%;">
   <div id="resizeWraper" class="photo-tagging-wrap" style="width: 100%; height: 100%; display: none; cursor: move;">
    <a id="close_btn" href="javascript:void(0);" class="close" title="点击退出圈人"><i>退出</i></a>
    <span id="mark_lt" class="sbox lt">&nbsp;</span>
    <span id="mark_lb" class="sbox lb">&nbsp;</span>
    <span id="mark_rt" class="sbox rt">&nbsp;</span>
    <span id="mark_rb" class="sbox rb">&nbsp;</span>
   </div>
  </div>
  <div id="selectArea" class="photo-tag-option photo-tag-optioning" style="position:absolute;z-index:1000;">
   <s><s></s></s>
   <div class="typearea"><input id="selectInputer" value="这是谁？" type="text"></div>
   <div style="width:180px;height:251px;" id="selectorContainer"></div>
  </div>
 </div>
</div></div>
 <div id="control_pane" class="quan-exit" style="display:none;z-index:1000;">退出圈人</div>
</div><div class="identity-index" style="margin: 10px; z-index: 999; display: none;" id="J_VasBubble_1">
	    <b class="ui-dl-arrow ui-dl-arrow1"></b>
	    <b class="ui-dl-arrow ui-dl-arrow2"></b>
	    <div class="conten">
	        <div class="identity-bg">
	            <div class="light1"></div>
	            <div class="light2"></div>
	        </div>
	        <div class="main-content">
	            <div class="plus identity-card-wrap threed">
	                <div class="identity-card identity-card-zero">
	                    <div class="ui-icon-vip">
	                        <i class="vip-lv1-g"></i>
	                    </div>
	                </div>
	            </div>
	            <div class="identity-tips">
	                <p>暂未获得勋章</p>
	            </div>
	            <div class="mod-identity">	                <ul class="identity-list">
	                    <li class="identity-list-item">
	                        <button class="control" style="cursor:pointer" data-ruleid="1317" data-mon="3">点亮</button>
	                        <div class="identity-list-wrap">
	                            <div class="identity-card identity-card-one">
	                                <div class="ui-icon-vip"><i></i></div>
	                            </div>
	                            <div class="identity-text">
	                                <p>续费3个月</p>
	                                <p>点亮黄钻青铜勋章</p>
	                            </div>
	                        </div>
	                    </li>
	                    <li class="identity-list-item">
	                        <button class="control" style="cursor:pointer" data-ruleid="1318" data-mon="6">点亮</button>
	                        <div class="identity-list-wrap">
	                            <div class="identity-card identity-card-two">
	                                <div class="ui-icon-vip"><i></i></div>
	                            </div>
	                            <div class="identity-text">
	                                <p>续费6个月</p>
	                                <p>点亮黄钻白银勋章</p>
	                            </div>
	                        </div>
	                    </li>
	                    <li class="identity-list-item">
	                        <button class="control" style="cursor:pointer" data-ruleid="1319" data-mon="12">点亮</button>
	                        <div class="identity-list-wrap">
	                            <div class="identity-card identity-card-three">
	                                <div class="ui-icon-vip"><i></i></div>
	                            </div>
	                            <div class="identity-text">
	                                <p>续费12个月</p>
	                                <p>点亮黄钻黄金勋章</p>
	                            </div>
	                        </div>
	                    </li>
	                </ul>	            </div>
	        </div>
	    </div>
	</div><div tabindex="-1" style="position: absolute; left: 300px; top: 70px; bottom: auto; right: auto; margin: 0px; padding: 0px; outline: 0px; border: 0px none; background: transparent; z-index: 10000;" class="ui-popup ui-popup-show ui-popover ui-popover-bottom-left" role="dialog"></div><div tabindex="-1" style="position: absolute; left: 300px; top: 70px; bottom: auto; right: auto; margin: 0px; padding: 0px; outline: 0px; border: 0px none; background: transparent; z-index: 10005;" class="ui-popup ui-popup-show ui-popover ui-popover-bottom-left" role="dialog"></div><div tabindex="-1" style="position: absolute; left: 455px; top: -53px; bottom: auto; right: auto; margin: 0px; padding: 0px; outline: 0px; border: 0px none; background: transparent; z-index: 10006;" class="ui-popup ui-popup-show ui-popover ui-popover-top-left ui-popup-focus" role="dialog"></div></body></html>`
	reg := regexp.MustCompile("http(s)?://user.qzone.qq.com/\\d+")

	strs := reg.FindAllString(reStr, -1)
	for i:=0;i<len(strs) ;i++  {
		println(strs[i])
	}
}
