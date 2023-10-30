import { Container, CssBaseline } from "@mui/material";
import { GlobalStyles } from "@mui/material";
import { BrowserRouter as Router, Route, Routes } from "react-router-dom";
import { Form } from "./components/Form";
import { Timeline } from "./components/Timeline";
import { Header } from "./components/Header";
import { useEffect, useState } from "react";
import { ColorModeProvider } from "./components/theme/ColorModeProvider.jsx";
import { ToggleTheme } from "./components/theme/ToggleTheme.jsx";

function App() {
  const [posts, setPosts] = useState([]);
  const [isLoading, setIsLoading] = useState(false);
  const onSubmitted = (post) => {
    setPosts([post, ...posts]);    
  };

  const fetchPosts = async () => {
    setIsLoading(true);
    const res = await fetch("/api/posts");
    const data = await res.json();
    if (!res.ok) {
      console.error(data);
      return;
    }
    setPosts(data);
    setIsLoading(false);
  };

  useEffect(() => {
    fetchPosts();
  }, []);

  return (
    <>
      <ColorModeProvider>
        <CssBaseline />
        <Header />
        <ToggleTheme />
        <GlobalStyles
          styles={{
            body: {
              margin: 0,
            },
          }}
        />
        <Router>
          <Routes>
            <Route
              path="/"
              element={
                <Container
                  maxWidth="md"
                  sx={{
                    py: 3,
                  }}
                >
                  <Form onSubmitted={onSubmitted} />
                  <Timeline
                    posts={posts}
                    isLoading={isLoading}
                    fetchPosts={fetchPosts}
                  />
                </Container>
              }
            ></Route>
          </Routes>
        </Router>
      </ColorModeProvider>
    </>
  );
}

export default App;
