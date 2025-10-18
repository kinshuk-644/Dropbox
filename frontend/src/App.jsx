import React, { useState, useEffect } from 'react';
import axios from 'axios';

import './index.css';

const API_URL = 'http://localhost:8000';

const formatFileSize = (bytes) => {
  if (bytes === 0) return '0 Bytes';
  const k = 1024;
  const sizes = ['Bytes', 'KB', 'MB', 'GB'];
  const i = Math.floor(Math.log(bytes) / Math.log(k));
  return parseFloat((bytes / Math.pow(k, i)).toFixed(2)) + ' ' + sizes[i];
};

function App() {
  const [files, setFiles] = useState([]);
  const [selectedFile, setSelectedFile] = useState(null);
  const [message, setMessage] = useState('');
  const [isLoading, setIsLoading] = useState(false);

  const fetchFiles = async () => {
    try {
      const response = await axios.get(`${API_URL}/files`);
      setFiles(response.data || []);
    } catch (error) {
      console.error('Error fetching files:', error);
      setMessage('Could not fetch files.');
    }
  };

  useEffect(() => {
    fetchFiles();
  }, []);

  const handleFileChange = (event) => {
    setSelectedFile(event.target.files[0]);
    setMessage('');
  };

  const handleUpload = async () => {
    if (!selectedFile) {
      setMessage('Please select a file first!');
      return;
    }

    setIsLoading(true);
    setMessage('Uploading...');
    const formData = new FormData();
    formData.append('file', selectedFile);

    try {
      const response = await axios.post(`${API_URL}/upload`, formData, {
        headers: {
          'Content-Type': 'multipart/form-data',
        },
      });
      setMessage('File uploaded successfully!');
      setSelectedFile(null);
      document.getElementById('fileInput').value = null;
      fetchFiles();
    } catch (error) {
      setMessage(error.response?.data?.error || 'Error uploading file.');
      console.error('Error uploading file:', error);
    } finally {
      setIsLoading(false);
    }
  };

  const messageClass = message.includes('success')
    ? 'message success'
    : 'message error';

  return (
    <div className="container">
      <header>
        <h1>Mini Dropbox 📁</h1>
      </header>

      {/* Upload Section */}
      <div className="box upload-section">
        <h2>Upload a File</h2>
        <input
          id="fileInput"
          type="file"
          onChange={handleFileChange}
          accept=".jpg, .jpeg, .png, .txt, .json"
        />
        <button
          onClick={handleUpload}
          disabled={!selectedFile || isLoading}
        >
          {isLoading ? 'Uploading...' : 'Upload'}
        </button>
        {message && <p className={messageClass}>{message}</p>}
      </div>

      {/* File List Section */}
      <div className="box file-list">
        <h2>Your Files</h2>
        <ul>
          {files.length > 0 ? (
            files.map((file) => (
              <li key={file.id}>
                <div className="file-details">
                  <span className="file-name">{file.original_name}</span>
                  <span className="file-meta">
                    {formatFileSize(file.size)} | {file.mimetype}
                  </span>
                </div>
                <a
                  href={`${API_URL}/files/${file.id}`}
                  target="_blank"
                  rel="noopener noreferrer"
                  className="download-button"
                >
                  View / Download
                </a>
              </li>
            ))
          ) : (
            <p>You haven't uploaded any files yet.</p>
          )}
        </ul>
      </div>
    </div>
  );
}

export default App;