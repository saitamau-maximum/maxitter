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
  const [users, setUsers] = useState([]);
  const [isLoading, setIsLoading] = useState(false);
  const [selectedUser, setSelectedUser] = useState('');//現在選択されているユーザーidを保持する変数
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
  
  const fetchUsers = async () =>{
    const res = await fetch("/api/users");
    if(res.ok){
      const data = await res.json();
      setUsers(data);
    } else{
      console.error(data);
    }
  }
  //セレクトボックスの人が変更されたときに呼ばれるハンドラ関数
  const handleChange = (event) => {
    setSelectedUser(event.target.value);
  };

  useEffect(() => {
    fetchPosts();
    fetchUsers();
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
            users = {users}
            selectedUser = {selectedUser}
            handleChange={handleChange}
          />
          <Form 
            onSubmitted={onSubmitted}
            selectedUser={selectedUser}
          />
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