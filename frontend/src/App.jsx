import { Container, CssBaseline } from "@mui/material";
import { GlobalStyles } from "@mui/material";
import { Form } from "./components/Form";
import { Timeline } from "./components/Timeline";
import { UserSelectBox } from "./components/UserSelectBox";
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
        <ToggleTheme />
        <GlobalStyles
          styles={{
            body: {
              margin: 0,
            },
          }}
        />
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
      </ColorModeProvider>
    </>
  );
}

export default App;