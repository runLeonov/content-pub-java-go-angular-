import { Component } from '@angular/core';
import {Title} from "../title";
import {HttpClient} from "@angular/common/http";
import {environment} from "../environment";
import {ActivatedRoute, Router} from "@angular/router";

@Component({
  selector: 'app-titleview',
  templateUrl: './titleview.component.html',
  styleUrls: ['./titleview.component.css']
})
export class TitleviewComponent {
  title: Title = Title.getEmptyTitle()

  constructor(private httpclient: HttpClient, private router: Router, private route: ActivatedRoute) {
    let id = this.route.snapshot.paramMap.get('id');
    if (id) {
      this.loadTitleByID();
    } else {
      this.loadTitleRandom();
    }
  }


  ngOnInit(): void {

  }

  private loadTitleByID() {
    let id = this.route.snapshot.paramMap.get('id');
    this.httpclient.get<Title>(`${environment.serverUrl}/titles/${id}`).subscribe(title => this.title = title);
  }

  private loadTitleRandom() {
    this.httpclient.get<Title>(`${environment.serverUrl}/titles/random`).subscribe(title => this.title = title);
  }
}
