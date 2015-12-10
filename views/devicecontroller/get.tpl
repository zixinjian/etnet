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
</head>
<body>
<div class="app">
    <div class="bg-light lter b-b wrapper-md">
        <div class="row">
            <div class="col-sm-3 col-xs-6">设备: {{.DeviceName}}[ID:{{.DeviceId}}]</div>
            <div class="pull-right m-r">当前状态:<span id="gstat" class="webo-c-stat" name="gstat"></span></i>
                <a class="btn btn-success btn-sm m-l-lg" id="startBtn">启动</a>
                <a class="btn btn-danger btn-sm" id="stopBtn">停止</a>
                <a class="btn btn-info btn-sm" id="btnShowPosition" data-toggle="collapse" data-target="#position"><i class="icon-pointer"></i>查看位置</a>
            </div>
        </div>
    </div>
    <div class="panel collapse m-r m-l no-padder" id="position">
        <div id="positionContent" style="height: 400px;"></div>
    </div>
    <div class="m">
        <div class="row">
            <div class="col-md-4">
                <div class="panel no-border">
                    <div class="panel-body padder">
                        <div class="no-padder" style="width: 280px;margin: auto">
                        <div id="pChart" style="height: 270px;width: 270px"></div>
                        </div>
                    </div>
                    <div class="hbox b-t b-light">
                        <span class="col padder-v text-muted b-r b-light padder webo-c-tag">
                            <span class="label text-base bg-primary pos-rlt"><i class="arrow right arrow-primary"></i>功率因子</span>
                            <span class="padder webo-c-tag-v" name="pfct"></span>
                        </span>
                        <span class="col padder-v text-muted padder webo-c-tag">
                            <span class="label text-base bg-primary pos-rlt "><i class="arrow right arrow-primary"></i>转数</span>
                            <span class="padder webo-c-tag-v" name="rpm"></span>
                        </span>
                    </div>
                </div>
            </div>
            <div class="col-md-4">
                <div class="panel no-border">
                    <div class="panel-heading text-center b-b"><h4>发电机</h4></div>
                    <div class="hbox b-light">
                        <span class="col padder-v b-r no-border-xs">
                            <div class="m-l m-b">
                                <span class="label text-base bg-info pos-rlt"><i class="arrow right arrow-info"></i>L1H</span>
                                <span class="padder webo-c-tag-v" name="v1"></span><span>V</span>
                            </div>
                            <div class="m-l m-b">
                                <span class="label text-base bg-info pos-rlt"><i class="arrow right arrow-info"></i>L2H</span>
                                <span class="padder webo-c-tag-v" name="v2"></span><span>V</span>
                            </div>
                            <div class="m-l">
                                <span class="label text-base bg-info pos-rlt"><i class="arrow right arrow-info"></i>L3H</span>
                                <span class="padder webo-c-tag-v" name="v3"></span><span>V</span>
                            </div>
                        </span>
                        <span class="col padder-v no-border-xs">
                            <div class="m-l m-b">
                                <span class="label text-base bg-info pos-rlt"><i class="arrow right arrow-info"></i>L1L3&nbsp</span>
                                <span class="padder webo-c-tag-v" name="v12"></span><span>V</span>
                            </div>
                            <div class="m-l m-b">
                                <span class="label text-base bg-info pos-rlt"><i class="arrow right arrow-info"></i>L2L3&nbsp</span>
                                <span class="padder webo-c-tag-v" name="v23"></span><span>V</span>
                            </div>
                            <div class="m-l">
                                <span class="label text-base bg-info pos-rlt"><i class="arrow right arrow-info"></i>L3L1&nbsp</span>
                                <span class="padder webo-c-tag-v" name="v31"></span><span>V</span>
                            </div>
                        </span>
                    </div>
                    <div class="hbox b-t b-light">
                        <div class="m-l m-b m-t">
                            <span class="label text-base bg-primary pos-rlt"><i class="arrow right arrow-primary"></i>L1A</span>
                            <span class="padder webo-c-tag-v" name="al1"></span><span>A</span>
                        </div>
                        <div class="m-l m-b m-t">
                            <span class="label text-base bg-primary pos-rlt"><i class="arrow right arrow-primary"></i>L2A</span>
                            <span class="padder webo-c-tag-v" name="al2"></span><span>A</span>
                        </div>
                        <div class="m-l m-b m-t">
                            <span class="label text-base bg-primary pos-rlt"><i class="arrow right arrow-primary"></i>L3A</span>
                            <span class="padder webo-c-tag-v" name="al3"></span><span>A</span>
                        </div>
                    </div>
                    <div class="hbox b-t b-light">
                        <div class="m-l m-b m-t">
                            <span class="label text-base bg-primary pos-rlt"><i class="arrow right arrow-primary"></i>频率</span>
                            <span class="padder webo-c-tag-v" name="fqcy"></span><span>HZ</span>
                        </div>
                    </div>
                </div>
            </div>
            <div class="col-md-4">
                <div class="panel no-border">
                    <div class="panel-heading text-center b-b"><h4>状态</h4></div>
                    <div class="panel-body">
                        <div class="" style="margin-bottom: 58px">
                            <div class="webo-c-progress" name="flevel" data-max="100" data-min="0">
                                <div class="">
                                    <span class="pull-right text-primary">%</span>
                                    <span class="pull-right text-primary webo-c-progress-v"></span>
                                    <span>燃油位</span>
                                </div>
                                <div class="progress progress-xs m-t-sm bg-light">
                                    <div class="progress-bar bg-primary" data-toggle="tooltip" data-original-title="60%" style="width: 60%"></div>
                                </div>
                            </div>
                            <div class="webo-c-progress" name="opress" data-max="100" data-min="0">
                                <div class="">
                                    <span class="pull-right text-primary">Bar</span>
                                    <span class="pull-right text-primary webo-c-progress-v"></span>
                                    <span>油压</span>
                                </div>
                                <div class="progress progress-xs m-t-sm bg-light">
                                    <div class="progress-bar bg-primary" data-toggle="tooltip" data-original-title="60%" style="width: 60%"></div>
                                </div>
                            </div>
                            <div class="webo-c-progress" name="etemp" data-max="120" data-min="-50">
                                <div class="">
                                    <span class="pull-right text-primary">°C</span>
                                    <span class="pull-right text-primary webo-c-progress-v"></span>
                                    <span>水温</span>
                                </div>
                                <div class="progress progress-xs m-t-sm bg-light">
                                    <div class="progress-bar bg-primary" data-toggle="tooltip" data-original-title="60%" style="width: 60%"></div>
                                </div>
                            </div>
                            <div class="webo-c-progress" name="vbty" data-max="50">
                                <div class="">
                                    <span class="pull-right text-primary">V</span>
                                    <span class="pull-right text-primary webo-c-progress-v"></span>
                                    <span>电池电压</span>
                                </div>
                                <div class="progress progress-xs m-t-sm bg-light">
                                    <div class="progress-bar bg-primary" data-toggle="tooltip" data-original-title="60%" style="width: 60%"></div>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
        <div class=" b-b">
            <div class="row text-center">
                <div class="col-sm-3 col-xs-6">
                    <div class="panel padder-v item">
                        <div class="h2 m-b-sm webo-c-o" name="gcbglight"><span class="webo-c-o-v"></span><i class="fa fa-fw fa-circle-o text-lt text-sm" data-color="text-success"></i></div>
                        <div>GCB 绿灯</div>
                    </div>
                </div>
                <div class="col-sm-3 col-xs-6">
                    <div class="panel padder-v item">
                        <div class="h2 m-b-sm webo-c-o" name="gglight"><span class="webo-c-o-v"></span><i class="fa fa-fw fa-circle-o text-lt text-sm" data-color="text-success"></i></div>
                        <div>发电机绿灯</div>
                    </div>
                </div>
                <div class="col-sm-3 col-xs-6">
                    <div class="panel padder-v item">
                        <div class="h2 m-b-sm webo-c-o" name="grlight"><span class="webo-c-o-v"></span><i class="fa fa-fw fa-circle-o text-lt text-sm" data-color="text-danger"></i></div>
                        <div>发电机红灯</div>
                    </div>
                </div>
                <div class="col-sm-3 col-xs-6">
                    <div class="panel padder-v item">
                        <div class="h2 m-b-sm webo-c-o" name="brkstat" ><span class="webo-c-o-v"></span><i class="fa fa-fw fa-circle-o text-lt text-sm" data-color="text-info"></i></div>
                        <div>断路器状态</div>
                    </div>
                </div>
            </div>
        </div>

        <div class="wrapper-md bg-white-only b-b">
            <div class="row text-center">
                <div class="col-sm-2 col-xs-4">
                    <div>运行小时<i class="fa fa-fw fa-caret-up text-success text-sm"></i></div>
                    <div class="h2 m-b-sm webo-c-text-v" name="runtime"></div>
                </div>
                <div class="col-sm-2 col-xs-4">
                    <div>维护<i class="fa fa-fw fa-caret-up text-success text-sm"></i></div>
                    <div class="h2 m-b-sm webo-c-text-v" name="mantence"></div>
                </div>
                <div class="col-sm-2 col-xs-4">
                    <div>起动次数<i class="fa fa-fw fa-caret-up text-success text-sm"></i></div>
                    <div class="h2 m-b-sm webo-c-text-v" name="startstop"></div>
                </div>
                <div class="col-sm-2 col-xs-4">
                    <div>急停次数<i class="fa fa-fw fa-caret-up text-success text-sm"></i></div>
                    <div class="h2 m-b-sm webo-c-text-v" name="estop"></div>
                </div>
                <div class="col-sm-2 col-xs-4">
                    <div>停机次数<i class="fa fa-fw fa-caret-up text-success text-sm"></i></div>
                    <div class="h2 m-b-sm webo-c-text-v" name="stop"></div>
                </div>
            </div>
        </div>
        <div class="panel panel-default m-t">
            <div class="panel-heading bg-white">
                历史
                <a class="pull-right" id="showHistory" data-toggle="collapse" data-target="#history"><i class="icon-arrow-down"></i>点击展开</a>
            </div>
            <!--<div class="collapse" id="history">-->
                <!--<table class="table table-striped b-t b-b" id="historyTable">-->
                    <!--<thead>-->
                  <!-- -->
                    <!--</tbody>-->
                <!--</table>-->
            <!--</div>-->
        </div>
    </div>
</div>
<script src="../../lib/app/js/app.min.js"></script>
<script src="../../lib/echart/echarts-all.js"></script>
<script src="../../lib/bootstrap-table/bootstrap-table.js"></script>
<script src="../../lib/bootstrap-table/locale/bootstrap-table-zh-CN.js"></script>
<script src="http://api.map.baidu.com/api?v=1.5&ak=55Rsk2ZW0d6xqrr8XfYT8QHB"></script>
<script>

var map, pChart, vOption, aOption, pOption

function showPChart(){
    pOption = {
        toolbox: {
            show : true,
            feature : {
                saveAsImage : {show: true}
            }
        },
        axisLabel:{
          show:true
        },
        series : [{
            axisLabel: {
                show: true,
                formatter: null,
                textStyle: {
                    color: 'auto'
                }
            },
            name:'功率',
            type:'gauge',
            splitNumber:10,
            max:1000,
            detail : {show:true, offsetCenter: [0, '80%'], formatter:'{value}KW'},
            data:[{value: 0, name: '功率'}],
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
    if(data){
        option.series[0].data[0].value = data;
        chart.setOption(option, true);
    }
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
var statusMap = {
    16:"起始",
    17:"未预备",
    18:"预起动",
    19:"起动中",
    20:"间歇",
    21:"起动中",
    22:"运行中",
    23:"已合闸",
    24:"停机",
    25:"停机",
    26:"候命中",
    27:"冷却中",
    28:"应急手动",
    29:"市电合闸",
    30:"市电故障",
    31:"市电故障",
    32:"岛运行",
    33:"市电回复",
    34:"断路全分",
    35:"不计时",
    36:"MCB 合闸",
    37:"恢复延时",
    38:"市并时间",
    39:"怠速运行",
    40:"最低稳时",
    41:"最高稳时",
    42:"后冷却泵",
    43:"ＧＣＢ开",
    44:"停机阀",
    45:"起动延时",
    46:"(1Ph)",
    47:"(3PD)",
    48:"(3PY)",
    49:"MRS 模式"
}
function getStatus(code){
    if (code in statusMap){
        return statusMap[code]
    }
    return "未联网"
}
function refreshData(){
    $.post("/device/params",
            {
                sn:"{{.DeviceId}}"
            },
            function(data,status){
                gstat = data["gstat"]
                $("#gstat").text(getStatus(gstat))

                setChart(pChart, pOption, data["p"])
                $(".webo-c-text-v").each(function(){
                    $el = $(this)
                    $el.text(data[$el.attr("name")])
                })
                $(".webo-c-tag-v").each(function(){
                    $el = $(this)
                    $el.text(data[$el.attr("name")])
                })
                $(".webo-c-o").each(function(){
                    $el = $(this)
                    $i = $el.find("i")
                    color = $i.data("color")

                    if (data[$el.attr("name")] == 0){
                        $i.removeClass("fa-circle").addClass("fa-circle-o")
                        $i.removeClass(color).addClass("text-lt")
                    }else{
                        $i.removeClass("fa-circle-o").addClass("fa-circle")
                        $i.removeClass("text-lt").addClass(color)
                    }
                })
                $(".webo-c-progress").each(function(){
                    $el = $(this)
                    $elv = $el.find(".webo-c-progress-v")
                    oValue = value = data[$el.attr("name")]

                    max = $el.data("max")
                    min = $el.data("min")
                    if (value > max){
                        value = max
                    }
                    if (value < min){
                        value = min
                    }
                    rat = value *100/max + "%"
                    $progressBar = $el.find(".progress-bar")
                    $progressBar.width(rat)
                    $progressBar.attr("data-original-title", rat)
                    if (oValue < -327){
                        value = "#"
                    }
                    $elv.text(value)
                })
            });
}
$(function(){
//    showVChart()
//    showAChart()
    $("#startBtn").on("click", function(){
        $.post("/device/operate", {
                    sn:"{{.DeviceId}}",
                    operate:"start"
                },
                function(data,status){

                })
    })
    $("#stopBtn").on("click", function(){
        $.post("/device/operate", {
                    sn:"{{.DeviceId}}",
                    operate:"stop"
                },
                function(data,status){

                })
    })
    $("#position").on("shown.bs.collapse", function(){
        showPostion(116.4821,35.7107)
    })
    showPChart()
    refreshData()
    $('#historyTable').bootstrapTable();
    setInterval(function(){
        refreshData()
    },5000);
})
</script>
</body>
</html>