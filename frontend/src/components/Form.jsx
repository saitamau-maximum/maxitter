import { Box, Button, FormLabel, TextField } from "@mui/material";
import { useState } from "react";

export const Form = ({ onSubmitted ,selectedUser}) => {
  const [body, setBody] = useState("");
  const [isSending, setIsSending] = useState(false);

  const sendPost = async (e) => {
    e.preventDefault();
    setIsSending(true);
    const res = await fetch(`/api/posts?user_id=${selectedUser}`, {
      method: "POST",
      body: JSON.stringify({ body }),
      headers: {
        "Content-Type": "application/json",
      },
    });
    setIsSending(false);
    setBody("");
    if (res.ok) {
      onSubmitted(await res.json());
    }
  };

  return (
    <>
      <form onSubmit={sendPost}>
        <FormLabel htmlFor="body">Maxitterに投稿する</FormLabel>
        <TextField
          variant="outlined"
          margin="normal"
          required
          fullWidth
          multiline
          rows={4}
          id="body"
          label="いまどうしてる？"
          name="body"
          autoFocus
          value={body}
          onChange={(e) => setBody(e.target.value)}
        />
        <Box textAlign="center" m={3}>
          <Button
            type="submit"
            variant="contained"
            color="primary"
            disabled={!body || isSending}
          >
            投稿する
          </Button>
        </Box>
      </form>
    </>
  );
};
