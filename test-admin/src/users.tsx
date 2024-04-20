import { useMediaQuery, Theme } from "@mui/material";
import { List, SimpleList, Datagrid, TextField, EmailField, UrlField, DateField, 
    Show, SimpleShowLayout, Edit, SimpleForm, Create, TextInput } from "react-admin";
import MyUrlField from './MyUrlField';
import { useRecordContext} from "react-admin";

const UserTitle = () => {
    const record = useRecordContext();
    return <span>Email {record ? `"${record.email}"` : ''}</span>;
};

export const UserList = () => {
    const isSmall = useMediaQuery<Theme>((theme) => theme.breakpoints.down("sm"));
    return (
        <List>
            <Datagrid rowClick="show">
                <TextField source="id" />
                <EmailField source="email" />
                <TextField source="real_name" />
                <TextField source="id_card_number" />
                <TextField source="nickname" />
                <TextField source="password" />
            </Datagrid>
        </List>
    );
};

export const UserShow = () => (
    <Show>
        <SimpleShowLayout>
            <TextField source="id" />
            <EmailField source="email" />
            <TextField source="real_name" />
            <TextField source="id_card_number" />
            <TextField source="nick_name" />
            <TextField source="password" />
        </SimpleShowLayout>
    </Show>
);

export const UserEdit = () => (
    <Edit title={<UserTitle />}>
        <SimpleForm>
            <TextField source="id" />
            <TextInput source="email" />
            <TextInput source="real_name" />
            <TextInput source="id_card_number" />
            <TextInput source="nickname" />
            <TextInput source="password" />
        </SimpleForm>
    </Edit>
);

export const UserCreate = () => (
    <Create>
        <SimpleForm>
            <TextField source="id" />
            <TextInput source="email" />
            <TextInput source="real_name" />
            <TextInput source="id_card_number" />
            <TextInput source="nickname" />
            <TextInput source="password" />
        </SimpleForm>
    </Create>
);