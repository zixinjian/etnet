<!DOCTYPE html>
<html>
<head>
    <meta charset="utf-8" />
    <title>Webo</title>
    <meta name="description" content="管理系统" />
    <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1" />
    <link rel="stylesheet" href="../lib/simple-line-icons/css/simple-line-icons.css" type="text/css" />
    <link rel="stylesheet" href="../lib/font-awesome/css/font-awesome.min.css" type="text/css" />
    <link rel="stylesheet" href="../lib/app/css/app.min.css" type="text/css" />
</head>
<body>
<div class="app app-header-fixed">
    <header id="header" class="app-header navbar" role="menu">
        <!-- navbar header -->
        <div class="navbar-header bg-dark">
            <button class="pull-right visible-xs dk" ui-toggle-class="show" target=".navbar-collapse">
                <i class="glyphicon glyphicon-cog"></i>
            </button>
            <button class="pull-right visible-xs" ui-toggle-class="off-screen" target=".app-aside" ui-scroll="app">
                <i class="glyphicon glyphicon-align-justify"></i>
            </button>
            <!-- brand -->
            <a class="navbar-brand text-lt" ui-toggle-class="app-aside-folded" target=".app">
                <i class="glyphicon glyphicon-th-large"></i>
                <span class="hidden-folded m-l-xs">HuaLi</span>
            </a>
            <!-- / brand -->
        </div>
        <div class="collapse pos-rlt navbar-collapse box-shadow bg-white-only">
            <div class="nav navbar-nav hidden-xs">
                <a href="#" class="btn no-shadow navbar-btn" ui-toggle-class="app-aside-folded" target=".app">
                    <i class="fa fa-dedent fa-fw text"></i>
                    <i class="fa fa-indent fa-fw text-active"></i>
                </a>
            </div>
            <ul class="nav navbar-nav navbar-right">
                <li class="hidden-xs">
                    <a ui-fullscreen>
                        <i class="fa fa-expand fa-fw text"></i>
                        <i class="fa fa-compress fa-fw text-active"></i>
                    </a>
                </li>
                <li>
                    <a href="#" data-toggle="dropdown" class="dropdown-toggle clear" data-toggle="dropdown">
                        <i class="icon-user icon text-success-lter m-r-xs"></i>
                        <span class="hidden-sm hidden-md">{{.UserName}}</span>
                    </a>
                </li>
                <li><a href="/logout"><i class="icon-arrow-right icon text-success-lter"></i></a></li>
            </ul>
        </div>
    </header>
    <!-- / header -->
    <aside id="aside" class="app-aside hidden-xs bg-dark">
        <div class="aside-wrap">
            <div class="navi-wrap">
                <!-- nav -->
                <nav ui-nav class="navi clearfix">
                    <ul class="nav">
                        <li class="hidden-folded padder m-t m-b-sm text-muted text-xs">
                            <span>全部设备</span>
                        </li>
                        <li class="active">
                            <a href="/gis" target="main" class="auto">
                                <i class="icon-pointer text-primary-lter"></i>
                                <span>位置信息</span>
                            </a>
                        </li>
                        <li class="hidden-folded padder m-t m-b-sm text-muted text-xs">
                            <span>设备列表</span>
                        </li>
                        <li>
                            <a href="/device?sn=10001" target="main" class="auto">
                                <i class="fa fa-truck text-success-lter"></i>
                                <span>应急电源车A</span>
                            </a>
                        </li>
                        <li>
                            <a href class="auto">
                                <span class="pull-right text-muted">
                                    <i class="fa fa-fw fa-angle-right text"></i>
                                    <i class="fa fa-fw fa-angle-down text-active"></i>
                                </span>
                                <i class="icon-grid text-success-lter"></i>
                                <span class="font-bold">发电机组(东区)</span>
                            </a>
                            <ul class="nav nav-sub dk">
                                <li>
                                    <a href="/ui/purchase/buyertimely" target="main">
                                        <span>发电机组A</span>
                                    </a>
                                </li>
                                <li>
                                    <a href="/ui/purchase/producttimely" target="main">
                                        <span>500W</span>
                                    </a>
                                </li>
                                <li>
                                    <a href="/ui/purchase/suppliertimely" target="main">
                                        <span>5500W</span>
                                    </a>
                                </li>
                            </ul>
                        </li>
                        <li>
                            <a href="/ui/expense/list" target="main" class="auto">
                                <i class="glyphicon-send text-success-lter"></i>
                                <span>雾霾炮</span>
                            </a>
                        </li>
                        <li>
                            <a href="/ui/expense/accountcurrentlist" target="main" class="auto">
                                <i class="glyphicon-send text-success-lter"></i>
                                <span>雾霾炮1</span>
                            </a>
                        </li>
                    </ul>
                </nav>
                <!-- nav -->
            </div>
        </div>
    </aside>
    <!-- / aside -->
    <div id="content" class="app-content" role="main">
        <iframe class="bg-light " name="main" src="/device?sn=10001" layout-auto-height="-50" style="width:100%;border:none"></iframe>
    </div>
</div>
<script src="../lib/app/js/app.min.js"></script>
<script src="../lib/screenfull/screenfull.min.js"></script>
<script src="../lib/webo/js/ui.js"></script>
<script>
    function showTopModal(options){
        $("#content").hide()
        $("#topModalContent").attr({src:options.url})
        $("#topModal").show()
        $("#topModalBtnOk").off("click").on("click", function(evt){
            var topModalWindows = $("#topModalContent")[0].contentWindow;
            if(topModalWindows && topModalWindows.onTopModalOk){
                topModalWindows.onTopModalOk(options)
            }
        })
    }
    function hideTopModal(){
        if($("#content").is(":hidden")){
            $("#topModalBtnOk").off("click")
            $('#topModal').hide();
            $("#topModalContent").replaceWith('<iframe id="topModalContent" class="bg-light " name="topContent" layout-auto-height="-100" style="width:100%;border:none"></iframe>');
            layoutAutoHeight()
            $("#content").show()
        }
    }

    $(function(){
        $('#topModal').hide()
        $("#topModalBtnCancel").click(hideTopModal)
        $('a[target=main]').click(function(e){
            hideTopModal()
        })
        initFullScreen();
    });
</script>
</body>
</html>
