import { useState } from "react";
import axios from "axios";

function App() {
  const [originalUrl, setOriginalUrl] = useState("");
  const [shortUrl, setShortUrl] = useState("");
  const [error, setError] = useState("");

  const baseUrl = import.meta.env.VITE_API_BASE_URL;
  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      const response = await axios.post("/api/urls", {
        url: originalUrl.trim()
      }, {
        headers: {
          'Content-Type': 'application/json'
        }
      });
      setShortUrl(response.data.short_url);
      setError("");
    } catch (err) {
      setError(err.response?.data?.message || "An error occurred");
      setShortUrl("");
    }
  };

  return (
    <div style={{ padding: "2rem", fontFamily: "Arial, sans-serif" }}>
      <h1>URL Shortener</h1>
      <form onSubmit={handleSubmit} style={{ marginBottom: "1rem" }}>
        <div style={{ marginBottom: "1rem" }}>
          <label htmlFor="originalUrl" style={{ display: "block", marginBottom: "0.5rem" }}>
            Original URL:
          </label>
          <input
            type="url"
            id="originalUrl"
            value={originalUrl}
            onChange={(e) => setOriginalUrl(e.target.value)}
            placeholder="Enter your URL"
            required
            style={{
              width: "100%",
              padding: "0.5rem",
              fontSize: "1rem",
              border: "1px solid #ccc",
              borderRadius: "4px",
            }}
          />
        </div>
        <button
          type="submit"
          style={{
            padding: "0.75rem 1.5rem",
            fontSize: "1rem",
            color: "#fff",
            backgroundColor: "#007BFF",
            border: "none",
            borderRadius: "4px",
            cursor: "pointer",
          }}
        >
          Shorten URL
        </button>
      </form>

      {shortUrl && (
        <div style={{ marginTop: "1rem", padding: "1rem", backgroundColor: "#f9f9f9", borderRadius: "4px" }}>
          <h3>Short URL:</h3>
          <a 
            href={baseUrl + shortUrl} 
            target="_blank" 
            rel="noopener noreferrer"
            style={{ wordBreak: "break-all" }}
          >
            {baseUrl + shortUrl}
          </a>
        </div>
      )}
      {error && <p style={{ color: "red" }}>{error}</p>}
    </div>
  );
}

export default App;