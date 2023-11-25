let UUID = "";
let IsPlayer1Turn= true;

function extractGameStateData(pos) {
    UUID = pos["uuid"]
    console.log("UUID: ", UUID)
    IsPlayer1Turn = pos["isPlayer1Turn"]
    document.getElementById("IsPlayer1Turn").innerHTML = IsPlayer1Turn ? "White's turn!" : "Black's turn!"
    delete pos["uuid"]
    delete pos["isPlayer1Turn"]
    return pos
}
function onDrop(oldPos, newPos) {
    let boardPosition = board.position()
    boardPosition["uuid"] = UUID
    fetch("/move", {
        method: "POST",
        body: JSON.stringify({
            start: oldPos,
            end: newPos,
            board: boardPosition,
        }),
        headers: {
            "Content-Type": "application/json"
        }
    }).then((response) => {
        response.json().then((data) => {
            let pos = extractGameStateData(data.BoardPosition)
            p = {}
            for (let k in pos) {
                if (pos[k] !== "empty") {
                    p[k] = pos[k]
                }
            }
            board.position(p, false)
        }
        )
        });
}

function flipBoard() {
    console.log("Flipping board...")
    let boardPosition = board.position()
    boardPosition["uuid"] = UUID
    fetch("/flip", {
        method: "POST",
        body: JSON.stringify({
            board: boardPosition,
        }),
        headers: {
            "Content-Type": "application/json"
        }
    }).then((response) => {
        response.json().then((data) => {
                let pos = extractGameStateData(data.BoardPosition)
                p = {}
                for (let k in pos) {
                    if (pos[k] !== "empty") {
                        p[k] = pos[k]
                    }
                }
                board.position(p, false)
            }
        )
    });

}

function initBoard() {
    console.log("Initializing board...")
    fetch("/init", {
        method: "GET",
        headers: {
            "Content-Type": "application/json"
        }
    }).then((response) => {
        response.json().then((data) => {
                let pos = extractGameStateData(data.BoardPosition)
                console.log(pos)
                p = {}
                for (let k in pos) {
                    if (pos[k] !== "empty") {
                        p[k] = pos[k]
                    }
                }
                board.position(p, false)
            }
        )
    });

}


let board = Chessboard('board1', {
    draggable: true,
    dropOffBoard: 'snapback',
    onDrop: onDrop,
})

$('#startBtn').on('click', initBoard)
$('#flipBoardBtn').on('click', flipBoard)