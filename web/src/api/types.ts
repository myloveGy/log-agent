export interface PageQuery {
  page: number;
  page_size: number;
}

export interface Pagination<T> extends PageQuery {
  items: T[];
  total: number;
}

export interface DatabaseQuery extends PageQuery {
  query: Record<string, any>;
  collection: string;
  order: 'asc' | 'desc';
  start_time?: string;
  end_time?: string;
}

export interface User {
	password: string;
	last_login_time: string;
	updated_at: string;
	last_login_ip: string;
	created_at: string;
	id: string;
	username: string;
	status: string;
}









