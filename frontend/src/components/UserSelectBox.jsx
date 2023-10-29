import { 
    FormControl,
    InputLabel,
    Select,
    MenuItem,
} from '@mui/material';


export  const UserSelectBox = ({users,selectedUser, handleChange}) =>{
  return(
  <>
  <FormControl fullWidth>
    <InputLabel id="user">User</InputLabel>
    <Select
      labelId="user"
      id="user"
      value={selectedUser}
      label="User"
      onChange={handleChange}
    >
    {users.map((user) => (
      <MenuItem value = {user.id}>{user.name}</MenuItem>
    ))}
    </Select>
  </FormControl>
  </>
  )
}