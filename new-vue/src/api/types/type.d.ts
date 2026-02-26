export interface ChangePasswod {
  oldpassword: string;
  newpassword: string;
}

export interface User {
  id: number;
  nickname: string;
  password?: string;
  email: string;
  headimg: string;
  created?: number;
  updated?: number;
  create_id?: number;
  realname: string;
  disable: boolean;
  position_id: number | null;
}

export interface Login {
  username: string;
  password: string;
}

export interface Version {
  id: number;
  pid: number;
  name: string;
  urlone: string;
  urltwo: string;
  createtime: number;
  createuid: number;
}

export interface StatusGroup {
  id: number;
  name: string;
  sids: number[];
}

export interface Status {
  id: number;
  name: string;
}

export interface ReqMyBugFilter {
  page: number;
  limit: number;
  level_id: number | null;
  project_id: number | null;
  title: string;
  page_type: number;
  showstatus: number | null;
}


export interface Bug {
  id: number;
  title: string;
  status_id: number
  create_id: number
  handle_uid: number | null;
  content: string;
  important_id: number | null;
  level_id: number | null;
  env_id: number | null;
  type_id: number;
  project_id: number | null;
  deadline: number;
  dustbin: boolean;
}

export interface UserGroup {
  id: number;
  name: string;
  uids: number[];
  uid: number;
}

export interface RequestPass {
  bid: number;
  remark: string;
  spusers: number[];
}

export interface ChangeStatus {
  id: number;
  status: string;
  code: number;
}

export interface RequestRole {
  id: number;
  name: string;
  perm_ids: number[];
}

export interface Role {
  id: number;
  name: string;
  info: string;
}

export interface ProjectRequsest {
  id: number;
  name: string;
  uid: number;
  uids: number[] | null;
}

export interface Position {
  id: number;
  name: string;
  level: number;
  hypo: number | null;
  role_id: number | null;
}

export interface SearchLog {
  create_time: string;
  page: number;
  limit: number;
  count: number;
  endtime: number;
  classify: string;
  ip: string;
}

export interface Auth {
  id: number;
  name: string;
  uid: number;
  created: number;
  pri: string;
  uptime: number;
  pub: string;
  typ: number;
  user: string;
  password: string;
}

export interface RequestProject {
  project_id: number;
}

export interface Email {
  host: string;
  enable: boolean;
  id: number;
  port: number;
  email: string;
  nickname: string;
  password: string;
  to: string;
}

export interface Doc {
  id: number;
  name: string;
  uid: number;
  created: number;
  giturl: string;
  uptime: number;
  dir: string;
  port: number;
  kid: number;
  authname: string;
}

export interface DefaultValue {
  created: number;
  completed: number;
  pass: number;
  receive: number;
}

export interface Perm {
  id: number;
  label: string[];
  value: string[];
  rid: number;
  info: string;
}

export interface Role {
  id: number;
  role: string;
  info: string;
}
