
// chat message object
var chat =
    {
        Id: 0,
        Nickname: '',
        Body: '',
        Type: '',
        Point_X: 0,
        Point_Y: 0,
        Point_Top: 0,
        Point_Left: 0,
        Image: null,
    };

var listers = [];

const url = "ws://127.0.0.1:8080/ws";

var socket = new WebSocket("ws://127.0.0.1:8080/ws");

if (window.MozWebSocket) {
    socket = new MozWebSocket(url);
}
socket.addEventListener("open", () => {
    console.log("socket is open");
});
$("#file").on('change', function (e) {
    var file = e.originalEvent.target.files[0];
    var reader = new FileReader();
    reader.onload = function (evt) {
        chat.Image = evt.target.result;
    };
    reader.readAsDataURL(file);
});
$("#chat-form").on('submit', function (e) {

    e.preventDefault();

    chat.Id = 1;
    chat.Nickname = localStorage.getItem("nickname");
    chat.Body = $("#msg").val();
    chat.Type = "chat";
    let message = JSON.stringify(chat);
    try {
        socket.send(message);
    } catch (e) {
        console.log(e);
    }
    $("#file").val('');
    chat.Image = "";
    $(".message_input").val('');
});

socket.onclose = () => {
    console.log("websocket is disconnected now .");
};


// When true, moving the mouse draws on the canvas
let isDrawing = false;
let x = 0;
let y = 0;

const myPics = document.getElementById('canvas');
const context = myPics.getContext('2d');

// The x and y offset of the canvas from the edge of the page
const rect = myPics.getBoundingClientRect();

// Add the event listeners for mousedown, mousemove, and mouseup
myPics.addEventListener('mousedown', e => {
    x = e.clientX - rect.left;
    y = e.clientY - rect.top;
    isDrawing = true;
});

let msg = chat;
myPics.addEventListener('mousemove', e => {
    if (isDrawing === true) {
        drawLine(context, x, y, e.clientX - rect.left, e.clientY - rect.top);
        msg.Point_X = x;
        msg.Point_Y = y;
        msg.Point_Left = e.clientX - rect.left;
        msg.Point_Top = e.clientY - rect.top;
        x = e.clientX - rect.left;
        y = e.clientY - rect.top;
        msg.Type = "paint";
        let jsonMsg = JSON.stringify(msg);
        socket.send(jsonMsg);
    }
});

window.addEventListener('mouseup', e => {
    if (isDrawing === true) {
        drawLine(context, x, y, e.clientX - rect.left, e.clientY - rect.top);
        x = 0;
        y = 0;
        isDrawing = false;
    }
});

function drawLine(context, x1, y1, x2, y2) {
    context.beginPath();
    context.strokeStyle = '#ffffff';
    context.lineWidth = 1;
    context.moveTo(x1, y1);
    context.lineTo(x2, y2);
    context.stroke();
    context.closePath();

}

var i = 20;

socket.addEventListener("message", (message) => {

    let response = JSON.parse(message.data);

    if (response.type === "chat") {

        let chat_box = $("#chat-box");
        let className = "me";
        let bodyClass = "body_me";
        let user = '<span class="writer">' + '<i class="fa fa-user"></i> ' + response.nickname + '</span>';
        let fileClass = "file_me";
        if (response.nickname !== localStorage.getItem("nickname")) {
            className = "other";
            bodyClass = "body_other";
            fileClass = "";
        }
        var div = `<p class="${className}">`;
        div = div + user + `<p class="${bodyClass}">${response.body}</p>`;
        div = div + `<p><a target="_blank" href="${response.Image}"><img width="45%" class="${fileClass}"  src="${response.Image}"></a></p>`;
        chat_box.append(div);

        chat_box.stop().animate({scrollTop: chat_box[0].scrollHeight}, 1000);

    } else if (response.type === "paint") {
        // console.log(response.point_x, response.point_y, response.point_left, response.point_top);

        drawLine(context, response.point_x, response.point_y, response.point_left, response.point_top);
    }
});
