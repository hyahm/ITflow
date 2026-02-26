export interface Response {
  code: number;
  msg: string;
  id: number;
  update_time: number;
  create_time: number;
  user_ids: number[];
  version_ids: number[];
  data: User | KeyName[] | RespShowBug;
  count: number;
  page: number;
  is_admin: boolean;
}

export interface User {
  id: number;
  nickname: string;
  password: string;
  email: string;
  headimg: string;
  createtime: number;
  createuid: number;
  realname: string;
  showstatus: number[];
  disable: boolean;
  jid: number;
}

export interface KeyName {
  id: number;
  name: string;
}

export interface Option {
  value: number;
  label: string;
}

export interface Informations {
  user: string;
  date: number;
  info: string;
}

export interface RespShowBug {
  status: string;
  title: string;
  content: string;
  id: number;
  selectuser: string[];
  important: string;
  level: string;
  projectname: string;
  envname: string;
  version: string;
  code: number;
  msg: string;
  url: string[];
  typ: number;
  comments: Informations[];
}
