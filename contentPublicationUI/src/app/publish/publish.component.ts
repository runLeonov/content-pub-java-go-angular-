import {Component} from '@angular/core';
import {HttpClient} from "@angular/common/http";
import {Category, Image, Like, PossibleContent, Serial, Tag} from "../title";
import {environment} from "../environment";
import {Router} from "@angular/router";

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
  public type: string = '';


  public allPossible: PossibleContent | undefined;
  public categories: Category[] = [];
  public serials: Serial[] = [];
  public tags: Tag[] = [];
  public images: Image[] = [];

  constructor(private httpclient: HttpClient, private router: Router) {
    this.loadContents().subscribe(cot => {
      this.allPossible = cot
      this.type = this.allPossible.types[0].typeName
    });
  }

  private loadContents() {
    return this.httpclient.get<PossibleContent>(`${environment.serverUrl}/titles/content-all`);
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
    this.httpclient.post<any>(`${environment.serverUrl}/titles/`, {
      titleName: this.titleName,
      originalAuthor: this.originalAuthor,
      typeName: this.type,
      description: this.description,
      creationDate: new Date(),
      titleImg: this.titleImg,
      content: {
        likes: [],
        views: 0,
        images: this.images
      },
      categories: this.categories,
      tags: this.tags,
      serials: this.serials
    }).subscribe(x => {
      this.router.navigate(["/titles/" + x.id]);
    });
  }

  addToArr($event: any, obj: Category | Tag | Serial, arr: any[]) {
    if ($event.target.checked) {
      arr.push(obj);
      return;
    }
    removeFromArr(arr, obj);

    function removeFromArr(arr: any[], obj: Category | Tag | Serial) {
      if (arr.includes(obj)) {
        const index = arr.indexOf(obj, 0);
        if (index > -1) {
          arr.splice(index, 1);
        }
      }
    }
  }
}
