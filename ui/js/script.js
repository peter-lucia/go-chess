function onChange(oldPos, newPos) {
    console.log('Position changed:')
    console.log('Old position: ' + Chessboard.objToFen(oldPos))
    console.log('New position: ' + Chessboard.objToFen(newPos))
    console.log('~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~')
    // use go API to check validity of the move
    // if it's valid, accept the move, otherwise, snapback
    fetch("/move", {
        method: "POST",
        body: JSON.stringify({
            piece: "P1K1",
            start: "a1",
            end: "a2",
        }),
        headers: {
            "Content-Type": "application/json"
        }
    })
        .then((response) => {
            console.log("Response")
            console.log(response.json())
        });
}

let board = Chessboard('board1', {
    draggable: true,
    dropOffBoard: 'snapback',
    position: 'start',
    onChange: onChange
})

$('#startBtn').on('click', board.start)
$('#clearBtn').on('click', board.clear)