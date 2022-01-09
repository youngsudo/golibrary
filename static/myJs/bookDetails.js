$(document).ready(() => {
    $("#changeDetails").click(() => {
        console.log(
        $("#book_id").val(),
        $("#book_title").val(),
        $("#book_author").val(),
        $("#book_state").val(),
        $("#book_details").val()
        )

        $.ajax({
            type: "post",
            url: "/admin/books/bookDetails",
            data: {
                "id" : $("#book_id").val(),
                "title" : $("#book_title").val(),
                "author" : $("#book_author").val(),
                "state" : $("#book_state").val(),
                "content" : $("#book_details").val()
            },
            success: (data) => {
                if (data.result == 0) {
                    $("#result").val("修改成功")
                }
            }
        })
    })
})