// 跳转
var classNa;
$('.cateUl li').on('click', function () {
    classNa = $(this).attr('class');
    switch (classNa) {
        case 'comments':
            window.location.href = './comments.html'
            break;
        case 'myPost':
            window.location.href = './myPost.html'
            break;
        case 'mylike':
            window.location.href = './mylike.html'
            break;
        case 'look':
            window.location.href = './look.html'
            break;
        case 'fans':
            window.location.href = './fans.html'
            break;
    }
})
$('.firstP').on('click', function () {
    window.location.href = './information.html'
})
// 首页HOME
$(".home").on('click', function () {
    $(this).css('font-weight', '700');
    window.location.href = './index.html'
})
//个人中心
$(".me").on('click', function () {
    $(this).css('font-weight', '700');
    window.location.href = './memberIndex.html'
})
