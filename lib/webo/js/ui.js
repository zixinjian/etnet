/**
 * Created by rick on 15/7/19.
 */
// 重定义部分控件属性
if ($.fn.bootstrapTable){
    $.fn.bootstrapTable.columnDefaults.sortable = true
}
//if($.validator){
//    $.validator.setDefaults(
//    )
//}
/* ==========================================================
 * sco.message.js
 * http://github.com/terebentina/sco.js
 * ==========================================================
 * Copyright 2013 Dan Caragea.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 * ========================================================== */

/*jshint laxcomma:true, sub:true, browser:true, jquery:true, eqeqeq: false */

;(function($, undefined) {
    "use strict";

    var pluginName = 'scojs_message';

    $[pluginName] = function(message, type) {
        clearTimeout($[pluginName].timeout);
        var $selector = $('#' + $[pluginName].options.id);
        if (!$selector.length) {
            $selector = $('<div/>', {id: $[pluginName].options.id}).appendTo($[pluginName].options.appendTo);
        }
        if ($[pluginName].options.animate) {
            $selector.addClass('page_mess_animate');
        } else {
            $selector.removeClass('page_mess_animate');
        }
        $selector.html(message);
        if (type === undefined || type == $[pluginName].TYPE_ERROR) {
            $selector.removeClass($[pluginName].options.okClass).addClass($[pluginName].options.errClass);
        } else if (type == $[pluginName].TYPE_OK) {
            $selector.removeClass($[pluginName].options.errClass).addClass($[pluginName].options.okClass);
        }
        $selector.slideDown('fast', function() {
            $[pluginName].timeout = setTimeout(function() { $selector.slideUp('fast'); }, $[pluginName].options.delay);
        });
    };
    $.extend($[pluginName], {
        options: {
            id: 'page_message'
            ,okClass: 'alert alert-success page_mess_ok'
            ,errClass: 'alert alert-danger page_mess_error'
            ,animate: true
            ,delay: 4000
            ,appendTo: 'body'	// where should the modal be appended to (default to document.body). Added for unit tests, not really needed in real life.
        },

        TYPE_ERROR: 1,
        TYPE_OK: 2
    });
})(jQuery);


function hideAlert(){
    $(".alert").hide()
}
function showSuccess(tip){
    if($.scojs_message){
        $.scojs_message(tip, $.scojs_message.TYPE_OK);
    }else{
        showAlert("success", tip)
    }
}
function showError(tip){
    if($.scojs_message){
        $.scojs_message(tip, $.scojs_message.TYPE_ERROR);
    }else{
        showAlert("danger", tip)
    }
}
function showAlert(type, tip){
    $(".alert").addClass("alert-"+type)
    $(".alert").text(tip)
    $(".alert").show()
}
function showInputError(select, tip){
    group=wbGetParentFromGroup(select).addClass("has-error")
    group.find(".help-block").text(tip)
}
function clearInputError(select){
    group = wbGetParentFromGroup(select).removeClass("has-error")
    group.find(".help-block").text("")
}
function initFullScreen(){
    $el = $('[ui-fullscreen]')
    if (screenfull.enabled && !navigator.userAgent.match(/Trident.*rv:11\./)) {
        $el.removeClass('hide');
    }
    $el.on('click', function(){
        screenfull.toggle();
    });
    $(document).on(screenfull.raw.fullscreenchange, function () {
        if(screenfull.isFullscreen){
            $el.addClass('active');
        }else{
            $el.removeClass('active');
        }
    });
}

function layoutAutoHeight(){
    $.each($("[layout-auto-height]"), function(){
        var outHeight = $(this).attr("layout-auto-height")
        //console.log("outHeight", outHeight, $(window).height())
        $(this).height($(window).height() + parseInt(outHeight))
});
}

function wbGetParentFromGroupLabel(selecter){
    return wbGetParent(selecter, ".form-group").text()
}
function wbGetParentFromGroup(selecter){
    return wbGetParent(selecter, ".form-group")
}

function wbGetParent(selecter, parentSelecter){
    $self = $(selecter)
    $parent = $self.closest(parentSelecter)
    return $parent
}

function getTableHeight() {
    return $(window).height();
}
function getTopModelHeight(){
    return $(top).height()-100
}
$(function(){
    layoutAutoHeight()
    $(window).resize(function () {
        layoutAutoHeight()
    });
});