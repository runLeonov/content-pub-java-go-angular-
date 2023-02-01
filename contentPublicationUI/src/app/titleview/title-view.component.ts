import {Component} from '@angular/core';
import {Like, Title} from "../title";
import {HttpClient} from "@angular/common/http";
import {environment} from "../environment";
import {ActivatedRoute, Router} from "@angular/router";
import {User} from "../user";

@Component({
  selector: 'app-titleview',
  templateUrl: './title-view.component.html',
  styleUrls: ['./title-view.component.css']
})
export class TitleViewComponent {
  title: Title = Title.getEmptyTitle();
  user: User;
  isLiked: boolean = false;
  clickCount: number = 0;
  countSec: number = 0;
  isWorking: boolean = false;


  constructor(private httpclient: HttpClient, private router: Router, private route: ActivatedRoute) {
    let id: number = Number(this.route.snapshot.paramMap.get('id'));
    this.user = JSON.parse(String(localStorage.getItem("user")).toString());
    if (id) {
      this.loadTitleByID().subscribe(title => {
        this.title = title;
        this.isLiked = checkIfLiked(this.user, this.title.content.likes);
      });
    } else {
      this.loadTitleRandom().subscribe(title => {
        this.title = title;
        this.isLiked = checkIfLiked(this.user, this.title.content.likes);
      });
    }

    function checkIfLiked(user: User, likes: Like[]): boolean {
      let liked = false;
      likes.forEach(x => {
          if (x.userId === user.id) {
            liked = true;
            return;
          }
        }
      );
      return liked;
    }
  }


  ngOnInit(): void {

  }

  private loadTitleByID() {
    let id = this.route.snapshot.paramMap.get('id');
    return this.httpclient.get<Title>(`${environment.serverUrl}/titles/${id}`);
  }

  private loadTitleRandom() {
    return this.httpclient.get<Title>(`${environment.serverUrl}/titles/random`);
  }

  like(doLike: boolean) {
    this.clickCount++;
    if (this.isWorking) {
      return;
    }

    this.isWorking = true;
    let interval = setInterval(() => {
      if (this.countSec === 5 && this.clickCount % 2 === 0) {
        this.clear(interval);
        return;
      }
      if (this.countSec === 5 && this.clickCount % 2 !== 0) {
        if (doLike) {
          this.httpclient.post<Title>(`${environment.serverUrl}/titles/like`, {
            titleContentId: this.title.id,
            userId: this.user.id
          }).subscribe();
        } else {
          this.httpclient.post<Title>(`${environment.serverUrl}/titles/unlike`, {
            titleContentId: this.title.id,
            userId: this.user.id
          }).subscribe();
        }
        this.clear(interval);
        return;
      }
      this.countSec++;
    }, 100);
  }

  clear(interval: NodeJS.Timer) {
    this.countSec = 0;
    this.isWorking = false;
    this.clickCount = 0;
    clearInterval(interval);
  }


}
