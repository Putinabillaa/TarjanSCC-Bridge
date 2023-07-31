import './App.css';
import React, { useState, useRef } from 'react';
import { toast, ToastContainer } from 'react-toastify';
import 'react-toastify/dist/ReactToastify.css';
import loading from './loading.png';

function App() {
  const [fileInput, setFileInput] = useState(null);
  const [textInput, setTextInput] = useState("");
  const [selectedOption, setSelectedOption] = useState(null);
  const [result, setResult] = useState(null);
  const [isLoading, setIsLoading] = useState(false);
  const fileInputRef = useRef(null);

  const cacheBustingQueryParam = Date.now();

  const handleTextChange = (event) => {
    setTextInput(event.target.value);
  };

  const handleFileChange = (event) => {
    const file = event.target.files[0];
    if (!file) {return;}
    if (file.type !== 'text/plain') {showErrorToast("Error file format, file should be a txt");}
    setFileInput(file);
  };

  const handleOptionChange = (option) => {
    setSelectedOption(option);
    setFileInput(null);
    setTextInput("");
  };

  const handleBrowseClick = () => {
    fileInputRef.current.click();
  };

  const handleFind = async () => {
    try {
      setIsLoading(true);
      const input = selectedOption === 'text' ? textInput : await readFileAsText(fileInput)
      
      const response = await fetch('http://localhost:8080/sccandbridge', {
          method: 'POST',
          headers: {
              'Content-Type': 'application/json',
          },
          body: JSON.stringify({
              input: input,
          }),
      });

      if (!response.ok) {
          const errorData = await response.json();
          throw new Error(errorData.error);
      }

      const data = await response.json();
      setResult(data)
      console.log('Graph Visualization Result:', data);
      setIsLoading(false);
    } catch (error) {
        showErrorToast(error.message);
        console.error('Error fetching data from the server:', error);
        setIsLoading(false);
    }
  };

  const readFileAsText = (file) => {
    return new Promise((resolve, reject) => {
      const reader = new FileReader();
      reader.onload = (event) => {
        resolve(event.target.result);
      };
      reader.onerror = (error) => {
        reject(error);
      };
      reader.readAsText(file);
    });
  };

  const showErrorToast = (message) => {
    toast.error(message, {
      position: toast.POSITION.TOP_RIGHT,
      autoClose: 5000,
      hideProgressBar: false,
      closeOnClick: true,
      pauseOnHover: true,
      draggable: true,
      progress: undefined,
    });
  };

  function areInputsEmpty() {
    return (selectedOption === 'text' && textInput.trim() === '') || (selectedOption === 'file' && !fileInput) || selectedOption === null;
  }

  return (
    <div className="App">
      <ToastContainer style={{textAlign:'left'}}/>
      <h1>SCC & Bridge Finder</h1>
      <div className="input-div">
        <h2>input:</h2>
        <div className="radio-buttons">
          <label htmlFor="file">
            <input
              type="radio"
              id="file"
              value="file"
              checked={selectedOption === "file"}
              onChange={() => handleOptionChange("file")}
            />
            File
          </label>
          <label htmlFor="text">
            <input
              type="radio"
              id="text"
              value="text"
              checked={selectedOption === "text"}
              onChange={() => handleOptionChange("text")}
            />
            Text
          </label>
        </div>
        {selectedOption === "file" && (
          <>
            <input
              type="file"
              onChange={handleFileChange}
              style={{ display: "none" }}
              ref={fileInputRef}
            />
            <button onClick={handleBrowseClick}>
            &#128196; Choose .txt File
            </button>
            {fileInput && <p>Selected File: {fileInput.name}</p>}
          </>
        )}
        {selectedOption === "text" && (
          <>
            <textarea
              rows={5}
              value={textInput}
              onChange={handleTextChange}
            />
          </>
        )}
      </div>
      <div className="buttons-div">
        <button onClick={handleFind} disabled={isLoading || areInputsEmpty()}> &#128270; Find</button>
        {isLoading && (
          <img className="loading-spinner" src={loading} alt="loading..."/>
        )}
      </div>
      <div className="output-div">
        <h2>output:</h2>
        {result && result.SCCs && result.Bridges && result.ImageSrc && (
          <>
            <p> SCCs: {result.SCCs.map((scc) => JSON.stringify(scc)).join(", ")}</p>
            <p> Bridges: {result.Bridges.map((bridge) => JSON.stringify(bridge)).join(", ")}</p>
            <p> {result.ExecTime} </p>
            <div className='result-images'>
              <p> SCCs Visualization: </p>
              <img key={0} src={`${result.ImageSrc[0]}?cb=${cacheBustingQueryParam}`} alt={`Graph`} />
              <p> SCCs Visualization: </p>
              <img key={1} src={`${result.ImageSrc[1]}?cb=${cacheBustingQueryParam}`} alt={`SCC`} />
              <p> Bridges Visualization (bridge shown as blue edges): </p>
              <img key={2} src={`${result.ImageSrc[2]}?cb=${cacheBustingQueryParam}`} alt={`Bridge`} />
            </div>
          </>
        )}
      </div>
    </div>
  );
}

export default App;
