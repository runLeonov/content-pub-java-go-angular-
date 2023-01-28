import {Component} from '@angular/core';
import {HttpClient} from "@angular/common/http";
import {FormsModule} from "@angular/forms"
import {Title} from "../title";
import {environment} from "../environment";

@Component({
  selector: 'app-publish',
  templateUrl: './publish.component.html',
  styleUrls: ['./publish.component.css']
})
export class PublishComponent {
  url: any[] = [];
  msg = "";
  public edited = false;
  public titleName: string = '';
  public originalAuthor: string = '';
  public description: string = '';
  public titleImg: any = '';

  constructor(private httpclient: HttpClient) {
  }

  selectTitleImg(event: any) {
    if (!event.target.files[0] || event.target.files[0].length == 0) {
      this.msg = 'You must select an image';
      return;
    }

    let mimeType = event.target.files[0].type;

    if (mimeType.match(/image\/*/) == null) {
      this.msg = "Only images are supported";
      return;
    }

    let reader = new FileReader();
    reader.readAsDataURL(event.target.files[0]);

    reader.onload = (_event) => {
      this.msg = "";
      this.url[0] = reader.result;
      this.titleImg = reader.result;
    }

    this.edited = true
  }

  addNewFile(event: any) {
    let mimeType = event.target.files[0].type;

    if (mimeType.match(/image\/*/) == null) {
      this.msg = "Only images are supported";
      return;
    }

    let reader = new FileReader();
    reader.readAsDataURL(event.target.files[0]);

    reader.onload = (_event) => {
      this.msg = "";
      this.url[1] = reader.result;
    }
  }

  addPost() {
    this.httpclient.post<Title>(`${environment.serverUrl}/titles/`, {
      titleName: this.titleName,
      originalAuthor: this.originalAuthor,
      description: this.description,
      creationDate: new Date(),
      titleImg: this.titleImg,
      content: {
        likes: 0,
        views: 0,
        images: [
        ]
      },
      categories: [
        {id:1, genre: ""},
        {id:2,genre: ""},
        {id:3, genre: ""},
        {id:4,genre: ""},
      ],
      tags: [
      ],
      serials: [
        {serialName: ""},
      ]
    }).subscribe()
  }
}
