import {Like} from "./title";

export class User {
  public id: number;
  public email: string;
  public name: string;
  public role: string;
  public img64: string;
  public likes: Like[];

  constructor(id: number, email: string, name: string, role: string, img64: string, likes: Like[]) {
    this.id = id;
    this.email = email;
    this.name = name;
    this.role = role;
    this.img64 = img64;
    this.likes = likes;
  }
}
