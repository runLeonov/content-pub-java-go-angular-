import {Component, OnInit} from '@angular/core';
import {environment} from "../../environment";
import {HttpClient} from "@angular/common/http";
import {Router} from "@angular/router";

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {

  email: string = ''
  password: string = ''
  emptyEmail: boolean = false
  emptyPass: boolean = false

  constructor(private httpclient: HttpClient, private router: Router) {

  }

  ngOnInit(): void {

  }

  submit(): void {
    this.httpclient.post<any>(`${environment.serverUrl}/auth/login`, {
      email: this.email,
      password: this.password,
    }).subscribe(() => {
      this.getUser();
    });
  }

  getUser() {
    if (!this.email) {
      this.emptyEmail = true;
      return;
    }
    if (!this.password) {
      this.emptyPass = true;
      return;
    }

    this.httpclient.get<any>(`${environment.serverUrl}/auth/user`).subscribe(user => {
      localStorage.setItem("user", JSON.stringify(user));
      this.router.navigate(['/titles']);
      document.location.reload();
    });
  }
}
