import { Component } from '@angular/core';
import {Title} from "../title";
import {HttpClient} from "@angular/common/http";
import {environment} from "../environment";
import {ActivatedRoute} from "@angular/router";

@Component({
  selector: 'app-titleview',
  templateUrl: './titleview.component.html',
  styleUrls: ['./titleview.component.css']
})
export class TitleviewComponent {
  title: Title = Title.getEmptyTitle()

  constructor(private httpclient: HttpClient, private route: ActivatedRoute) {
    this.loadTitleByID();
  }


  ngOnInit(): void {

  }

  private loadTitleByID() {
    let id = this.route.snapshot.paramMap.get('id');
    this.httpclient.get<Title>(`${environment.serverUrl}/titles/${id}`).subscribe(title => this.title = title)
  }
}
