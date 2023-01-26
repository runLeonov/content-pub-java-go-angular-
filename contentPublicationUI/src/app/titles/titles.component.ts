import {Component} from '@angular/core';
import {Title} from "../title";
import {HttpClient} from "@angular/common/http";
import {Observable} from "rxjs";

@Component({
  selector: 'app-titles',
  templateUrl: './titles.component.html',
  styleUrls: ['./titles.component.css']
})
export class TitlesComponent {
  titles: Title[] = []

  constructor(private httpclient: HttpClient) {
    this.loadTitlesList();
    console.log(this.titles);
  }


  ngOnInit(): void {

  }

  private loadTitlesList() {
    this.httpclient.get<Title[]>('content/titles').subscribe(titles => this.titles = titles)
  }
}
