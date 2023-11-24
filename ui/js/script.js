function onDrop(oldPos, newPos) {
    fetch("/move", {
        method: "POST",
        body: JSON.stringify({
            start: oldPos,
            end: newPos,
            board: board.position()
        }),
        headers: {
            "Content-Type": "application/json"
        }
    }).then((response) => {
        response.json().then((data) => {
            let pos = data.BoardPosition
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
    fetch("/flip", {
        method: "POST",
        body: JSON.stringify({
            board: board.position()
        }),
        headers: {
            "Content-Type": "application/json"
        }
    }).then((response) => {
        response.json().then((data) => {
                let pos = data.BoardPosition
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

function initBoard() {
    console.log("Initializing board...")
    fetch("/init", {
        method: "GET",
        headers: {
            "Content-Type": "application/json"
        }
    }).then((response) => {
        response.json().then((data) => {
                let pos = data.BoardPosition
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