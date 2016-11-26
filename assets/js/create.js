$(function(){
    $("#DescriptionTextarea").bind("input change", function(){
    console.log("sdfjgjfbhjkdnf")
         $.post("/markdown-html", {md: $("#DescriptionTextarea").val()}, function(response){
            $("#markdown").html(response.html)
         })
    })
})