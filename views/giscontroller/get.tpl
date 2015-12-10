<!DOCTYPE html>
<html>
<head>
    <meta name="viewport" content="initial-scale=1.0, user-scalable=no" />
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
    <title>Hello, World</title>
    <style type="text/css">
        html{height:100%}
        body{height:100%;margin:0px;padding:0px}
        #container{height:100%}
    </style>
    <script type="text/javascript" src="http://api.map.baidu.com/api?v=1.5&ak=55Rsk2ZW0d6xqrr8XfYT8QHB">
        //v1.5版本的引用方式：src="http://api.map.baidu.com/api?v=1.5&ak=您的密钥"
        //v1.4版本及以前版本的引用方式：src="http://api.map.baidu.com/api?v=1.4&key=您的密钥&callback=initialize"
    </script>
</head>

<body>
<div id="container"></div>
<script src="../../lib/app/js/app.min.js"></script>
<script type="text/javascript">
    var map
    function showPostion(x, y){
        map = new BMap.Map("container");          // 创建地图实例
        map.addControl(new BMap.NavigationControl());
        map.addControl(new BMap.ScaleControl());
        map.addControl(new BMap.OverviewMapControl());
        map.addControl(new BMap.MapTypeControl());
        var point = new BMap.Point(x, y);  // 创建点坐标
        map.centerAndZoom(point, 10);
        var marker = new BMap.Marker(point);        // 创建标注
        map.addOverlay(marker);                     // 将标注添加到地图中
    }
    function setLocation(x, y){
        map.panTo(new BMap.Point(x, y), 15)
    }
    function refreshData(){
        $.post("/device/params",
                {
                    sn:"{{.DeviceId}}"
                },
                function(data,status){
                    x= data("x")
                    y= data("y")
                    function setLocation(x, y){
                        map.panTo(new BMap.Point(x, y), 15)
                    }
                })
    }
    $(function(){
        showPostion(116.4821,35.7107)
    })
</script>
</body>
</html>