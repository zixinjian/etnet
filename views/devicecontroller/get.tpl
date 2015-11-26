<!DOCTYPE html>
<html>
<head>
    <meta charset="utf-8" />
    <title>Webo</title>
    <meta name="description" content="管理系统" />
    <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1" />
    <link rel="stylesheet" href="../../lib/simple-line-icons/css/simple-line-icons.css" type="text/css" />
    <link rel="stylesheet" href="../../lib/font-awesome/css/font-awesome.min.css" type="text/css" />
    <link rel="stylesheet" href="../../lib/app/css/app.min.css" type="text/css" />
    <link rel="stylesheet" href="../../lib/bootstrap-table/bootstrap-table.css">

    <!--<link rel="stylesheet" href="../../lib/datatables/css/dataTables.bootstrap.css">-->
</head>
<body>
<div class="app">
    <div class="bg-light lter b-b wrapper-md">
        <div class="row">
            <div class="col-sm-3 col-xs-6">设备: {{.DeviceName}}[ID:{{.DeviceId}}]</div>
            <div class="pull-right m-r">当前状态:<i class="fa fa-circle-o text-info"></i>待机
                <a class="btn btn-success btn-sm m-l-lg">启动</a>
                <a class="btn btn-danger btn-sm">停止</a>
                <a class="btn btn-info btn-sm" id="btnShowPosition" data-toggle="collapse" data-target="#position"><i class="icon-pointer"></i>查看位置</a>
            </div>
        </div>
    </div>
    <div class="panel collapse m-r m-l" id="position" style="padding: 0">
        <div id="positionContent" style="height: 400px;"></div>
    </div>
    <div class="m">
        <div class="row">
            <div class="col-md-4">
                <div class="panel no-border">
                    <div class="panel-body text-center">
                        <div class="text-center" id="vChart" style="height: 300px;"></div>
                    </div>
                </div>
            </div>
            <div class="col-md-4">
                <div class="panel no-border">
                    <div class="panel-body">
                        <div id="aChart" style="height: 300px"></div>
                    </div>
                </div>
            </div>
            <div class="col-md-4">
                <div class="panel no-border">
                    <div class="panel-body">
                        <div id="pChart" style="height: 300px"></div>
                    </div>
                </div>
            </div>
        </div>
        <div class="wrapper-md bg-white-only b-b">
            <div class="row text-center">
                <div class="col-sm-3 col-xs-6">
                    <div>参数1<i class="fa fa-fw fa-caret-up text-success text-sm"></i></div>
                    <div class="h2 m-b-sm webo-c-param" name="param1">219k</div>
                </div>
                <div class="col-sm-3 col-xs-6">
                    <div>参数1<i class="fa fa-fw fa-caret-up text-success text-sm"></i></div>
                    <div class="h2 m-b-sm webo-c-param" name="param2">219k</div>
                </div>
                <div class="col-sm-3 col-xs-6">
                    <div>参数1<i class="fa fa-fw fa-caret-up text-success text-sm"></i></div>
                    <div class="h2 m-b-sm webo-c-param" name="param3">219k</div>
                </div>
                <div class="col-sm-3 col-xs-6">
                    <div>参数1<i class="fa fa-fw fa-caret-up text-success text-sm"></i></div>
                    <div class="h2 m-b-sm webo-c-param" name="param4">219k</div>
                </div>
            </div>
        </div>
        <div class="panel panel-default m-t">
            <div class="panel-heading bg-white">
                参数历史
                <a class="pull-right" id="showHistory" data-toggle="collapse" data-target="#history"><i class="icon-arrow-down"></i>点击展开</a>
            </div>
            <div class="collapse" id="history">
                <table class="table table-striped b-t b-b" id="historyTable">
                    <thead>
                    <tr>
                        <th  style="width:20%">时间</th>
                        <th  style="width:25%">电流</th>
                        <th  style="width:25%">电压</th>
                        <th  style="width:15%">功率</th>
                        <th  style="width:15%">参数1</th>
                        <th  style="width:15%">参数2</th>
                        <th  style="width:15%">参数3</th>
                        <th  style="width:15%">参数4</th>
                    </tr>
                    </thead>
                    <tbody>
                    <tr>
                        <td  style="width:20%">Rendering engine</td>
                        <td  style="width:25%">Browser</td>
                        <td  style="width:25%">Platform(s)</td>
                        <td  style="width:15%">Engine version</td>
                        <td  style="width:15%">CSS grade</td>
                        <td  style="width:15%">CSS grade</td>
                        <td  style="width:15%">CSS grade</td>
                        <td  style="width:15%">CSS grade</td>
                    </tr><tr>
                        <td  style="width:20%">Rendering engine</td>
                        <td  style="width:25%">Browser</td>
                        <td  style="width:25%">Platform(s)</td>
                        <td  style="width:15%">Engine version</td>
                        <td  style="width:15%">CSS grade</td>
                        <td  style="width:15%">CSS grade</td>
                        <td  style="width:15%">CSS grade</td>
                        <td  style="width:15%">CSS grade</td>
                    </tr><tr>
                        <td  style="width:20%">Rendering engine</td>
                        <td  style="width:25%">Browser</td>
                        <td  style="width:25%">Platform(s)</td>
                        <td  style="width:15%">Engine version</td>
                        <td  style="width:15%">CSS grade</td>
                        <td  style="width:15%">CSS grade</td>
                        <td  style="width:15%">CSS grade</td>
                        <td  style="width:15%">CSS grade</td>
                    </tr><tr>
                        <td  style="width:20%">Rendering engine</td>
                        <td  style="width:25%">Browser</td>
                        <td  style="width:25%">Platform(s)</td>
                        <td  style="width:15%">Engine version</td>
                        <td  style="width:15%">CSS grade</td>
                        <td  style="width:15%">CSS grade</td>
                        <td  style="width:15%">CSS grade</td>
                        <td  style="width:15%">CSS grade</td>
                    </tr>
                    </tbody>
                </table>
            </div>
        </div>
    </div>
</div>
<script src="../../lib/app/js/app.min.js"></script>
<script src="../../lib/echart/echarts-all.js"></script>
<script src="../../lib/bootstrap-table/bootstrap-table.js"></script>
<script src="../../lib/bootstrap-table/locale/bootstrap-table-zh-CN.js"></script>
<script src="http://api.map.baidu.com/api?v=1.5&ak=55Rsk2ZW0d6xqrr8XfYT8QHB"></script>
<script>

var map, vChart, aChart, pChart, vOption, aOption, pOption

function showVChart(){
    vOption = {
        toolbox: {
            show : true,
            feature : {
                saveAsImage : {show: true}
            }
        },
        series : [
            {
                name:'电压',
                type:'gauge',
                splitNumber:10,
                max:440,
                detail : {formatter:'{value}V'},
                data:[{value: 220, name: '电压'}],
                axisLine: {            // 坐标轴线
                    lineStyle: {       // 属性lineStyle控制线条样式
                        color: [[0.4, '#228b22'],[0.85, '#48b'],[1, '#ff4500']],
                        width: 8
                    }
                }
            }
        ]
    };
    vChart = echarts.init(document.getElementById("vChart"))
    vChart.setOption(vOption)
    vChart.setTheme("macarons")
}
function showAChart(){
    aOption = {
        toolbox: {
            show : true,
            feature : {
                saveAsImage : {show: true}
            }
        },
        series : [
            {
                name:'电流',
                type:'gauge',
                splitNumber:10,
                max:100,
                detail : {formatter:'{value}A'},
                data:[{value: 50, name: '电流'}],
                axisLine: {            // 坐标轴线
                    lineStyle: {       // 属性lineStyle控制线条样式
                        color: [[0.4, '#228b22'],[0.85, '#48b'],[1, '#ff4500']],
                        width: 8
                    }
                }
            }
        ]
    };
    aChart = echarts.init(document.getElementById("aChart"))
    aChart.setOption(aOption)
    aChart.setTheme("macarons")
}
function showPChart(){
    pOption = {
        toolbox: {
            show : true,
            feature : {
                saveAsImage : {show: true}
            }
        },
        series : [
            {
                name:'功率',
                type:'gauge',
                splitNumber:10,
                max:1000,
                detail : {formatter:'{value}KW'},
                data:[{value: 800, name: '电流'}],
                axisLine: {            // 坐标轴线
                    lineStyle: {       // 属性lineStyle控制线条样式
                        color: [[0.4, '#228b22'],[0.85, '#48b'],[1, '#ff4500']],
                        width: 8
                    }
                }
            }
        ]
    };
    pChart = echarts.init(document.getElementById("pChart"))
    pChart.setOption(pOption)
    pChart.setTheme("macarons")
}
function setChart(chart, option, data){
    option.series[0].data[0].value = data;
    chart.setOption(option, true);
}
function showPostion(x, y){
    map = new BMap.Map("positionContent");          // 创建地图实例
    map.addControl(new BMap.NavigationControl());
    map.addControl(new BMap.ScaleControl());
    map.addControl(new BMap.OverviewMapControl());
    map.addControl(new BMap.MapTypeControl());
    var point = new BMap.Point(x, y);  // 创建点坐标
    map.centerAndZoom(point, 15);
    var marker = new BMap.Marker(point);        // 创建标注
    map.addOverlay(marker);                     // 将标注添加到地图中
}
function setLocation(x, y){
    map.panTo(new BMap.Point(x, y), 15)
}
function refreshData(){
    $.post("/device/params",
            {
                sn:"1234"
            },
            function(data,status){
                console.log("adat", data,$("positionContent").is(":visible"))
                setChart(aChart, aOption, data["a"])
                setChart(vChart, vOption, data["v"])
                setChart(pChart, pOption, data["p"])
                $(".webo-c-param").each(function(){
                    $el = $(this)
                    $el.text(data[$el.attr("name")])
                })
                if(!$("positionContent").is(":hidden")){
                    setLocation(data["locX"], data["locY"])
                }
            });
}
$(function(){
    showVChart()
    showAChart()
    showPChart()
    showPostion(116.404, 39.915)
    $('#historyTable').bootstrapTable();
    setInterval(function(){
        refreshData()
    },5000);
})
</script>
</body>
</html>