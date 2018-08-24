$(".back").click(function () {
    history.back();
});

$("#getSms").click(function () {
    var phone = $("#tel").val();
    console.log(phone)
});