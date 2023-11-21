import { Box, Button, FormLabel, TextField } from "@mui/material";
import { useState } from "react";
//方針　とりあえずApp.jsxに追加して機能を試す
export const Assign = () => {
    const [name, setName] = useState("");
    const [profileURL, setProfileURL] = useState("");
    const [bio, setBio] = useState("");
    const [isSending, setIsSending] = useState(false);//一応書いておくが、連続して登録はしない気がするので後で考える。

    const sendPost = async (e) => {
        e.preventDefault();
        setIsSending(true);
        const res = await fetch(`/api/users/new`, {
          method: "POST",
          body: JSON.stringify({ name:name, profileURL:profileURL, bio:bio }),//ひとまず省略せずに書く
          headers: {
            "Content-Type": "application/json",
          },
        });
        setIsSending(false);
        //setBody("");  登録後の挙動は後で考える。
        if (res.ok) {
          onSubmitted(await res.json());//ここを見ると、もしやonSubmittedが必要？
        }
      };

      return (
        <>
          <form onSubmit={sendPost}>
            <FormLabel htmlFor="body">ユーザー登録</FormLabel>
            
            <TextField 
              variant="outlined"
              margin="normal"
              required
              fullWidth
              multiline
              rows={1}
              id="name"
              label="ユーザー名"
              name="name"
              autoFocus
              value={name}
              onChange={(e) => setName(e.target.value)}
            />
            <TextField 
              variant="outlined"
              margin="normal"
              required
              fullWidth
              multiline
              rows={1}
              id="profileURL"
              label="プロフィール画像"
              name="profileURL"
              autoFocus
              value={profileURL}
              onChange={(e) => setProfileURL(e.target.value)}
            />
             <TextField 
              variant="outlined"
              margin="normal"
              required
              fullWidth
              multiline
              rows={4}
              id="bio"
              label="自己紹介"
              name="bio"
              autoFocus
              value={bio}
              onChange={(e) => setBio(e.target.value)}
            />
            <Box textAlign="center" m={3}>
              <Button
                type="submit"
                variant="contained"
                color="primary"
                disabled={ !name && !profileURL && !bio }//ここの条件には後ですべてのボックスが空白でないことを追加する。
              >
                登録する
              </Button>
            </Box>
          </form>
        </>
      );
}