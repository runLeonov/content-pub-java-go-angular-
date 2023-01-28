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
import {TitleviewComponent} from './titleview/titleview.component';
import {TitleViewRandomComponent} from './title-view-random/title-view-random.component';

const appRoutes: Routes = [
  {path: 'titles', component: TitlesComponent},
  {path: 'profile', component: ProfileComponent},
  {path: 'title/create', component: PublishComponent},
  {path: 'random', component: TitleViewRandomComponent},
  {path: 'titles/:id', component: TitleviewComponent},
]

@NgModule({
  declarations: [
    AppComponent,
    HeadercComponent,
    TitlesComponent,
    ProfileComponent,
    PublishComponent,
    TitleviewComponent,
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
}
