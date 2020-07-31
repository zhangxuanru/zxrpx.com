
function getUrlParam(name) {
    var reg = new RegExp("(^|&)" + name + "=([^&]*)(&|$)"); //构造一个含有目标参数的正则表达式对象
    var r = window.location.search.substr(1).match(reg); //匹配目标参数
    if (r != null) return unescape(r[2]);
    return null; //返回参数值
}


order = getUrlParam("order");
$(".box_order_div a >i").remove();
orderText="热门";
if(order == "latest"){
    orderText = "最新";
    $(".box_order_div a:eq(0)").html("最新<i class='check'>✓</i>");
}else{
    $(".box_order_div a:eq(1)").html("热门<i class='check'>✓</i>");
}
$(".box_order").text(orderText);


pType = getUrlParam("type");
$(".box_type_div a >i").remove();
typeText="图片";
if(pType == "photo"){
    typeText="照片";
    $(".box_type_div a:eq(1)").html("照片<i class='check'>✓</i>");
}
if(pType == "vector"){
    typeText="矢量图";
    $(".box_type_div a:eq(2)").html("矢量图<i class='check'>✓</i>");
}
if(pType == "illustration"){
    typeText="插画";
    $(".box_type_div a:eq(3)").html("插画<i class='check'>✓</i>");
}
$(".box_type").text(typeText);



ori = getUrlParam("orientation");
$(".box_orientation_div a >i").remove();
oriText="图像方向";
if(ori == "horizontal"){
    oriText="水平";
    $(".box_orientation_div a:eq(1)").html("水平<i class='check'>✓</i>");
}
if(ori == "vertical"){
    oriText="垂直";
    $(".box_orientation_div a:eq(2)").html("垂直<i class='check'>✓</i>");
}
$(".box_orientation").text(oriText);


cat = getUrlParam("cat");
catText= "类别";
if(cat == "transportation"){
    catText = "交通运输";
}
if(cat == "industry"){
    catText = "产业/技术";
}
if(cat == "people"){
    catText = "人物";
}
if(cat == "animals"){
    catText = "动物";
}
if(cat == "health"){
    catText = "医学/健康";
}
if(cat == "business"){
    catText = "商业/金融";
}
if(cat == "places"){
    catText = "地点/地标";
}
if(cat == "religion"){
    catText = "宗教";
}
if(cat == "buildings"){
    catText = "建筑";
}
if(cat == "education"){
    catText = "教育";
}
if(cat == "travel"){
    catText = "旅游度假";
}
if(cat == "science"){
    catText = "科学/技术";
}
if(cat == "fashion"){
    catText = "美妆/时尚";
}
if(cat == "backgrounds"){
    catText = "背景/花纹";
}
if(cat == "nature"){
    catText = "自然风景";
}
if(cat == "feelings"){
    catText = "表情";
}
if(cat == "computer"){
    catText = "计算机/沟通";
}
if(cat == "sports"){
    catText = "运动";
}
if(cat == "music"){
    catText = "音乐";
}
if(cat == "food"){
    catText = "食物/饮料";
}
$(".box_category").text(catText);