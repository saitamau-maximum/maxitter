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
  const [selectedUser, setSelectedUser] = useState('');//現在選択されているユーザーを保持する変数
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
  //セレクトボックスの人が変更されたときに呼ばれるハンドラ関数
  const handleChange = (event) => {
    setSelectedUser(event.target.value);
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
          <UserSelectBox 
            user = {selectedUser}
            handleChange={handleChange}
          />
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