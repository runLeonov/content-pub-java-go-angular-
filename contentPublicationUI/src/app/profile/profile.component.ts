import {Component, OnInit} from '@angular/core';
import {User} from "../user";
import {AppComponent} from "../app.component";
import {HttpClient} from "@angular/common/http";
import {Router} from "@angular/router";
import {environment} from "../environment";
import {Comment, Title} from "../title";

@Component({
  selector: 'app-profile',
  templateUrl: './profile.component.html',
  styleUrls: ['./profile.component.css', '../title-view/title-view.component.css', '../titles/titles.component.css']
})
export class ProfileComponent  implements OnInit {
  user: User | undefined;
  likedTitles: Title[] | undefined
  commentedTitles: Title[] | undefined

  constructor(private httpclient: HttpClient, private router: Router) {
    this.httpclient.get<User>(`${environment.serverUrl}/account/`).subscribe(user => {
      this.user = user
    });

  }

  getLimitLikes() {
    this.httpclient.get<Title[]>(`${environment.serverUrl}/account/likes-limit`).subscribe(titles => {
      this.likedTitles = titles;
    });
  }
  getLimitComments() {
    this.httpclient.get<Title[]>(`${environment.serverUrl}/account/commented-limit`).subscribe(titles => {
      this.commentedTitles = titles;
    });
  }

  getReleaseYear(u: User): string {
    return new Date(u.creationDate).toLocaleDateString();
  }

  parseDate(date: Date): Date {
    return new Date(date);
  }

  ngOnInit(): void {
    this.getLimitLikes();
    this.getLimitComments();
  }
}
