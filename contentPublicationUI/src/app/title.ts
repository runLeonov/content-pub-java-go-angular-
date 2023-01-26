export class Title {
  private _id: number;
  private _titleName: string;
  private _originalAuthor: string;
  private _creationDate: Date;
  private _description: string;
  private _titleImgBase64: string;
  private _content: TitleContent;
  private _categories: Category[];
  private _tags: Tag[];
  private _serials: Serial[];


  constructor(id: number, titleName: string, originalAuthor: string, creationDate: Date, description: string, titleImgBase64: string, content: TitleContent, categories: Category[], tags: Tag[], serials: Serial[]) {
    this._id = id;
    this._titleName = titleName;
    this._originalAuthor = originalAuthor;
    this._creationDate = creationDate;
    this._description = description;
    this._titleImgBase64 = titleImgBase64;
    this._content = content;
    this._categories = categories;
    this._tags = tags;
    this._serials = serials;
  }

  get id(): number {
    return this._id;
  }

  set id(value: number) {
    this._id = value;
  }

  get titleName(): string {
    return this._titleName;
  }

  set titleName(value: string) {
    this._titleName = value;
  }

  get originalAuthor(): string {
    return this._originalAuthor;
  }

  set originalAuthor(value: string) {
    this._originalAuthor = value;
  }

  get creationDate(): Date {
    return this._creationDate;
  }

  set creationDate(value: Date) {
    this._creationDate = value;
  }


  get content(): TitleContent {
    return this._content;
  }

  set content(value: TitleContent) {
    this._content = value;
  }

  get description(): string {
    return this._description;
  }

  set description(value: string) {
    this._description = value;
  }

  get titleImgBase64(): string {
    return this._titleImgBase64;
  }

  set titleImgBase64(value: string) {
    this._titleImgBase64 = value;
  }

  get categories(): Category[] {
    return this._categories;
  }

  set categories(value: Category[]) {
    this._categories = value;
  }

  get tags(): Tag[] {
    return this._tags;
  }

  set tags(value: Tag[]) {
    this._tags = value;
  }

  get serials(): Serial[] {
    return this._serials;
  }

  set serials(value: Serial[]) {
    this._serials = value;
  }
}

export class TitleContent {
  private _id: number;
  private _likes: number;
  private _views: number;
  private _images: Image[];


  constructor(id: number, likes: number, views: number, images: Image[]) {
    this._id = id;
    this._likes = likes;
    this._views = views;
    this._images = images;
  }


  get id(): number {
    return this._id;
  }

  set id(value: number) {
    this._id = value;
  }

  get likes(): number {
    return this._likes;
  }

  set likes(value: number) {
    this._likes = value;
  }

  get views(): number {
    return this._views;
  }

  set views(value: number) {
    this._views = value;
  }

  get images(): Image[] {
    return this._images;
  }

  set images(value: Image[]) {
    this._images = value;
  }
}

export class Image {
  private _id: number;
  private _img64: number;


  constructor(id: number, img64: number) {
    this._id = id;
    this._img64 = img64;
  }

  get id(): number {
    return this._id;
  }

  set id(value: number) {
    this._id = value;
  }

  get img64(): number {
    return this._img64;
  }

  set img64(value: number) {
    this._img64 = value;
  }
}


export class Category {
  private _id: number;
  private _genre: string;
  private _titles: Title[];


  constructor(id: number, genre: string, titles: Title[]) {
    this._id = id;
    this._genre = genre;
    this._titles = titles;
  }

  get id(): number {
    return this._id;
  }

  set id(value: number) {
    this._id = value;
  }

  get genre(): string {
    return this._genre;
  }

  set genre(value: string) {
    this._genre = value;
  }

  get titles(): Title[] {
    return this._titles;
  }

  set titles(value: Title[]) {
    this._titles = value;
  }
}

export class Tag {
  private _id: number;
  private _tagName: string;
  private _titles: Title[];


  constructor(id: number, tagName: string, titles: Title[]) {
    this._id = id;
    this._tagName = tagName;
    this._titles = titles;
  }

  get id(): number {
    return this._id;
  }

  set id(value: number) {
    this._id = value;
  }

  get tagName(): string {
    return this._tagName;
  }

  set tagName(value: string) {
    this._tagName = value;
  }

  get titles(): Title[] {
    return this._titles;
  }

  set titles(value: Title[]) {
    this._titles = value;
  }
}

export class Serial {
  private _id: number;
  private _serialName: string;
  private _titles: Title[];


  constructor(id: number, serialName: string, titles: Title[]) {
    this._id = id;
    this._serialName = serialName;
    this._titles = titles;
  }


  get serialName(): string {
    return this._serialName;
  }

  set serialName(value: string) {
    this._serialName = value;
  }

  get id(): number {
    return this._id;
  }

  set id(value: number) {
    this._id = value;
  }

  get titles(): Title[] {
    return this._titles;
  }

  set titles(value: Title[]) {
    this._titles = value;
  }
}
