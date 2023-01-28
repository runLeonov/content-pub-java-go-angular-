import {Component} from '@angular/core';
import {Title} from "../title";
import {HttpClient} from "@angular/common/http";
import {environment} from "../environment";

@Component({
  selector: 'app-titles',
  templateUrl: './titles.component.html',
  styleUrls: ['./titles.component.css']
})
export class TitlesComponent {
  titles: Title[] = []

  constructor(private httpclient: HttpClient) {
    this.loadTitlesList();
  }


  ngOnInit(): void {

  }

  private loadTitlesList() {
    this.httpclient.get<Title[]>(`${environment.serverUrl}/titles/`).subscribe(titles => this.titles = titles)
  }
}
