export class User {
  public id: number;
  public email: string;
  public name: string;
  public role: string;


  constructor(id: number, email: string, name: string, role: string) {
    this.id = id;
    this.email = email;
    this.name = name;
    this.role = role;
  }
}
