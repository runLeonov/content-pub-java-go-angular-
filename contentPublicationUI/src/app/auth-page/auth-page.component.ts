import {Component, ElementRef, ViewChild} from '@angular/core';
import {environment} from "../environment";
import {HttpClient} from "@angular/common/http";
import {Router} from "@angular/router";

@Component({
  selector: 'app-auth-page',
  templateUrl: './auth-page.component.html',
  styleUrls: ['./auth-page.component.css']
})
export class AuthPageComponent {
  email: string | undefined
  password: string | undefined
  name: string | undefined
  img: any = ''

  incorrectPass: boolean = false
  userExist: boolean = false
  emptyEmail: boolean = false
  emptyName: boolean = false
  emptyPass: boolean = false

  msg = "";
  public edited = false;
  url: any;

  constructor(private httpclient: HttpClient, private router: Router) {
  }

  //@ts-ignore
  @ViewChild('pass') pass: ElementRef;
  //@ts-ignore
  @ViewChild('passRep') passRep: ElementRef;

  ngAfterViewInit() {
  }

  validation() {
    if (!this.email) {
      this.emptyEmail = true;
      return;
    }

    if (!this.name) {
      this.emptyName = true;
      return;
    }

    let pass = this.pass.nativeElement.value;
    let passRep = this.passRep.nativeElement.value;

    if (pass.length < 6) {
      this.emptyPass = true;
      return;
    }

    if (pass !== passRep) {
      this.incorrectPass = true;
      return;
    }

    this.incorrectPass = false;
    this.emptyName = false;
    this.emptyEmail = false;
    this.emptyPass = false;
    this.userExist = false;

    this.httpclient.post<boolean>(`${environment.serverUrl}/auth/check-exist`, {
      email: this.email
    }).subscribe(exist => {
      this.userExist = exist;
      if (!this.userExist) {
        this.singUp()
      }
    });
  }

  selectImg(event: any) {
    let mimeType = event.target.files[0].type;

    if (mimeType.match(/image\/*/) == null) {
      this.msg = "Only images are supported";
      return;
    }

    let reader = new FileReader();
    reader.readAsDataURL(event.target.files[0]);

    reader.onload = (_event) => {
      this.msg = "";
      this.url = reader.result;
      this.img = reader.result;
    }

    this.edited = true
  }

  singUp() {
    this.httpclient.post<bigint>(`${environment.serverUrl}/auth/register`, {
      email: this.email,
      password: this.password,
      name: this.name,
      img64: this.img
    }, {withCredentials: true}).subscribe(() => this.singIn());
  }

  singIn() {
    this.httpclient.post<any>(`${environment.serverUrl}/auth/login`, {
      email: this.email,
      password: this.password,
    }, {withCredentials: true}).subscribe(() => {
      this.getUser();
    });
  }

  getUser() {
    this.httpclient.get<any>(`${environment.serverUrl}/auth/user`, {withCredentials: true}).subscribe(user => {
      localStorage.setItem("user", JSON.stringify(user));
    });
  }
}
