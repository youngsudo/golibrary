$(document).ready(() => {
    // console.log($("#showtable"))
    var table = $("#showtable")[0];
    //获取所有行
    var rows = table.rows;
    console.log(rows)
    for (var i = 0; i < rows.length; i++) {
        var row = rows[i];//获取每一行
        var list = row.cells; //获取每一列
        for (var j = 0; j < list.length - 1; j++) {  //除去删除按钮
            // var td = list[j].innerHTML;//获取具体单元格    string
            // console.log(td)
            var td_input = list[j].children[0]  //获取具体单元格 HTML


            // 修改数据
            td_input.onchange = function () {
                //    console.log(this)    // td_input
                //    console.log(this.value)  // input里面的更改后的值
                //    console.log(this.parentNode.parentNode) // tr
                //    console.log(this.parentNode.parentNode.attributes.ind.value) // tr 的ind属性的值,即用户目前id,用户id可能也会被改变
                //    console.log(this.attributes.ind.value)  // 该input的类型ind
                let value = this.value;
                let id = this.parentNode.parentNode.attributes.ind.value;
                let flag = this.attributes.flag.value;

                $.ajax({
                    type: "post",
                    url: "/admin/users/change",
                    data: {
                        "id": id,
                        "flag": flag,
                        "value": value
                    },
                    success: (data) => {
                        // 同时因为页面不会刷新,如果更改了id,还应该将tr中的ind更改
                        if (flag == "id") {
                            this.parentNode.parentNode.attributes.ind.value = value
                        }
                        console.log(data)
                    }
                })
            }
        }
    }

})

// 删除数据
function func_users_delete(i) {
    // console.log(i)      // button
    // console.log(i.parentNode.parentNode)
    // console.log(i.parentNode.parentNode.attributes.ind.value)
    let id = i.parentNode.parentNode.attributes.ind.value;
    $.ajax({
        type: "post",
        url: "/admin/users/delete",
        data: {
            "id": id
        },
        success: (data) => {
            console.log(data.result)
            // 删除页面上的数据
            i.parentNode.parentNode.parentNode.removeChild(i.parentNode.parentNode)
        }
    })
}
function func_record_delete(i) {
    // console.log(i)      // button
    // console.log(i.parentNode.parentNode)
    // console.log(i.parentNode.parentNode.attributes.ind.value)
    let id = i.parentNode.parentNode.attributes.ind.value;
    $.ajax({
        type: "post",
        url: "/admin/record/delete",
        data: {
            "id": id
        },
        success: (data) => {
            console.log(data.result)
            // 删除页面上的数据
            i.parentNode.parentNode.parentNode.removeChild(i.parentNode.parentNode)
        }
    })
}

// 添加用户数据 form表单

// 上下滑动添加输入框
$("#slideToggle_but").click(() => {
    $("#slideToggle_div").slideToggle("slow");
})
