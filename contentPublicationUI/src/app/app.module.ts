import {NgModule} from '@angular/core';
import {BrowserModule} from '@angular/platform-browser';

import {AppComponent} from './app.component';
import {HeadercComponent} from './headerc/headerc.component';
import {HttpClientModule} from "@angular/common/http";
import {RouterModule, Routes} from "@angular/router";
import {TitlesComponent} from './titles/titles.component';
import {ProfileComponent} from './profile/profile.component';
import {PublishComponent} from './publish/publish.component';
import {FormsModule} from "@angular/forms";
import {TitleViewComponent} from './titleview/title-view.component';
import {TitleViewRandomComponent} from './title-view-random/title-view-random.component';
import {User} from "./user";

const appRoutes: Routes = [
  {path: 'titles', component: TitlesComponent},
  {path: 'profile', component: ProfileComponent},
  {path: 'title/create', component: PublishComponent},
  {path: 'random', component: TitleViewRandomComponent},
  {path: 'titles/:id', component: TitleViewComponent},
]

@NgModule({
  declarations: [
    AppComponent,
    HeadercComponent,
    TitlesComponent,
    ProfileComponent,
    PublishComponent,
    TitleViewComponent,
    TitleViewRandomComponent
  ],
  imports: [
    BrowserModule,
    HttpClientModule,
    RouterModule.forRoot(appRoutes),
    FormsModule,
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule {


  constructor() {
    let user: User = {
      email: "leonovcasha@gmail.com",
      name: "Sasha",
      role: "USER",
      id: 1,
      likes: [
        {titleContentId: 1, userId: 1},
        {titleContentId: 2, userId: 1},
      ]
    };

    localStorage.setItem("user", JSON.stringify(user));
  }
}
