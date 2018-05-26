var player = "";

$(function(){
  for(var x = 0; x < 25; x++){
    $("#tabl").append("<tr></tr>")
    var r = $("#tabl :last")
    for(var y = 0; y < 40; y++){
      r.append("<td class=''></td>")
    }
  }

  while(player != "1" && player != "2"){
    player = prompt("player 1 or 2?")
  }

  document.onkeydown = padMove

  loadBoard()
})

function padMove(e){
  e = e || window.event
  if(e.keyCode == 37){
    $.get("/api/v1/moveleft/" + player)
  }
  if(e.keyCode == 39){
    $.get("/api/v1/moveright/" + player)
  }
  // if(e.keyCode == 67){
  //   $.get("/api/v1/moveright/2")
  // }
  // if(e.keyCode == 88){
  //   $.get("/api/v1/moveleft/2")
  // }
}

function loadBoard(){

  // $.get("/api/v1/getboard", function(dat){
  //   var x = dat.split('\n')
  //   for(var i = 0; i < x.length; i++) {
  //     for(var j = 0; j < x[0].length; j++) {
  //       var el = $("#tabl tr:eq(" + String(i) + ") td:eq(" + String(j) + ")")
  //       if(x[i][j] == "0"){
  //         el.removeClass();
  //         el.addClass("m0")
  //       }
  //       else {
  //         el.removeClass();
  //         el.addClass("m1")
  //       }
  //     }
  //   }
  //   setTimeout(function(){
  //     console.log("hola")
  //     loadBoard()
  //   }, 10)
  //
  // })

  // var soc = new WebSocket("ws://10.66.156.124:8080/api/ws/getboard")
  var soc = new WebSocket("ws://localhost:8080/api/ws/getboard")

  soc.addEventListener('message', function(event) {
    var x = event.data.split(';')

    $("#tabl > tr > td").removeClass()

    for(var i = -3; i <= 3; i++){
      var el = $("#tabl tr:eq(1) td:eq("  + String(Number(x[0]) + i) + ")")
      el.removeClass()
      if (player == "1") {
        el.addClass("m2")
      }
      else {
        el.addClass("m1")
      }

    }
    for(var i = -3; i <= 3; i++){

        var el = $("#tabl tr:eq(23) td:eq("  + String(Number(x[1]) + i) + ")")
        el.removeClass()
        if (player == "2") {
          el.addClass("m2")
        }
        else {
          el.addClass("m1")
        }
    }
    var el = $("#tabl tr:eq(" + x[3] + ") td:eq(" + x[2] + ")")
    el.removeClass()
    el.addClass("m1")
  })

}
