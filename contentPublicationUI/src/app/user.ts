import {Like} from "./title";

export class User {
  public id: number;
  public email: string;
  public name: string;
  public role: string;
  public likes: Like[];

  constructor(id: number, email: string, name: string, role: string, likes: Like[]) {
    this.id = id;
    this.email = email;
    this.name = name;
    this.role = role;
    this.likes = likes;
  }
}
