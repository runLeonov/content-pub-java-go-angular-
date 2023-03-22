import {Component} from '@angular/core';
import {User} from "../user";
import {AppComponent} from "../app.component";
import {HttpClient} from "@angular/common/http";
import {Router} from "@angular/router";
import {environment} from "../environment";

@Component({
  selector: 'app-profile',
  templateUrl: './profile.component.html',
  styleUrls: ['./profile.component.css', '../title-view/title-view.component.css']
})
export class ProfileComponent {
  user: User | undefined;

  constructor(private httpclient: HttpClient, private router: Router) {
    this.httpclient.get<User>(`${environment.serverUrl}/account/`).subscribe(user => {
      this.user = user
    });
  }
}
