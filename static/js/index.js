// 跳转搜索页
$('#searchInp').on('click',function(){
    window.location.href='./search.html'
})
// 首页HOME
$(".home").on('click', function () {
    $(this).css('font-weight', '700')
})
//个人中心
$(".me").on('click', function () {
    $(this).css('font-weight', '700');
    window.location.href = './memberIndex.html'
})
// 帖子跳转
$('.content ul li').on('click', function () {
    window.location.href = './detail.html'
})