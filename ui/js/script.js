function onDrop(oldPos, newPos) {
    console.log('Position changed:')
    console.log('Old position: ' + JSON.stringify(oldPos))
    console.log('~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~')
    // use go API to check validity of the move
    // if it's valid, accept the move, otherwise, snapback
    console.log(board.position())
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
    position: 'start',
    onDrop: onDrop,
})

$('#startBtn').on('click', board.start)
$('#clearBtn').on('click', board.clear)