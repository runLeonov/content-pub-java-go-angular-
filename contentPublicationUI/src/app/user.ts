import {Like} from "./title";

export class User {
  public id: number;
  public email: string;
  public nickName: string;
  public role: string;
  public img64: string;
  public likes: Like[];
  public creationDate: Date = new Date();
  public lastName: string;
  public firstName: string;


  constructor(id: number, email: string, name: string, role: string, img64: string, likes: Like[], creationDate: Date, lastName: string, firstName: string) {
    this.id = id;
    this.email = email;
    this.nickName = name;
    this.role = role;
    this.img64 = img64;
    this.likes = likes;
    this.creationDate = creationDate;
    this.lastName = lastName;
    this.firstName = firstName;
  }
}
