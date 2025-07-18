import React, { useState } from 'react'

const App = () => {
  const [roomName, setRoomName] = useState('')
  const [bookedRooms, setBookedRooms] = useState([])

  const bookRoom = () => {
    if (roomName.trim()) {
      setBookedRooms([...bookedRooms, roomName])
      setRoomName('')
    }
  }

  return (
    <div className="app">
      <h1>🏨 Hotel Room Booking Anand</h1>
      <input
        type="text"
        placeholder="Enter room name or number"
        value={roomName}
        onChange={(e) => setRoomName(e.target.value)}
      />
      <button onClick={bookRoom}>Book Room</button>
      <ul>
        {bookedRooms.map((room, idx) => (
          <li key={idx}>✅ {room} booked</li>
        ))}
      </ul>
    </div>
  )
}

export default App
