//选择省市区 
mui.init();
mui(function () {
    cityPicker = new mui.PopPicker({
        layer: 3
    });
    cityPicker.setData(cityData3);
});
var address = '';
//选择区域
function chooseArea() {
    address = '';
    cityPicker.show(function (items) {
        // items为选中的三级区域信息
        for (var i = 0; i < items.length; i++) {
            address += items[i].text;
        }
        $('#address').html('')
        $('#address').html(address)
    });
}
//关闭选择器
function closeChoose() {
    cityPicker.hide();
}
$('#address').on('click', function () {
    chooseArea()
})
//  出生日期
var dtpicker = new mui.DtPicker({
    type: "date",//设置日历初始视图模式 
    beginDate: new Date(1900, 01, 01),//设置开始日期 
    endDate: new Date(),//设置结束日期 
});
$('#birTime').on('tap', function (e) {
    dtpicker.show(function (selectItems) {
        $(e.target).html(selectItems.text);
    });
});