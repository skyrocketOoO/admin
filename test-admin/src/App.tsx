import {
  Admin,
  Resource,
  ListGuesser,
  EditGuesser,
  ShowGuesser,
} from "react-admin";
import { dataProvider } from "./dataProvider";
import { UserList, UserShow, UserEdit, UserCreate } from "./users";
import { PostList, PostEdit, PostCreate } from "./posts";

import PostIcon from "@mui/icons-material/Book";
import UserIcon from "@mui/icons-material/Group";
import { Dashboard } from './Dashboard';
import { authProvider } from './authProvider';

export const App = () => (
    <Admin authProvider={authProvider} dataProvider={dataProvider} dashboard={Dashboard}>
        <Resource
            name="users"
            list={UserList}
            show={UserShow}
            edit={UserEdit}
            create={UserCreate}
            recordRepresentation="email"
            icon={UserIcon}
        />
    </Admin>
);