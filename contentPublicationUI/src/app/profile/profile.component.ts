import {Component} from '@angular/core';
import {User} from "../user";
import {AppComponent} from "../app.component";
import {HttpClient} from "@angular/common/http";
import {Router} from "@angular/router";

@Component({
  selector: 'app-profile',
  templateUrl: './profile.component.html',
  styleUrls: ['./profile.component.css']
})
export class ProfileComponent {
  user: User | undefined = AppComponent.getUser();

  constructor(private httpclient: HttpClient, private router: Router) {
  }
}
