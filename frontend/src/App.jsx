import { useEffect, useState } from 'react'
import axios from 'axios'
import './App.css'

function App() {
  const [newNumber, setNewNumber] = useState('')
  const [numbers, setNumbers] = useState([])
  const [error, setError] = useState('')

  useEffect(() => {
    getNumbers()
  }, [])   

  async function getNumbers() {
    try {
      const res = await axios.get('/api/numbers')
      setNumbers(res.data.list)
      setError('')
    } catch (err) {
      alert(err.response?.data?.error || 'Failed to fetch numbers')
    }
  }

  async function addNumber() {
    try {
      await axios.post('/api/numbers', { number: parseInt(newNumber) })
      getNumbers()
      setNewNumber('')
      setError('')
    } catch (err) {
      alert(err.response?.data?.error || 'Failed to add number')
    }
  }

  return (
    <>
      <div className="input-container">
        <input
          type="number"
          value={newNumber}
          onChange={(e) => {
            const val = e.target.value
            if (!isNaN(val)) setNewNumber(val)
          }}
        />
        <button onClick={addNumber}>Add Number</button>
      </div>
      {numbers && (
        <p className="number-list">
          {numbers.join(", ")}
        </p>
      )}
    </>
  )
}

export default App
