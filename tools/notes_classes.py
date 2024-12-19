from __future__ import annotations
from dataclasses import dataclass
from enum import StrEnum
from hashlib import sha256
import json
import os
import re
from typing import Callable, Dict, List, Tuple
from urllib.parse import urljoin

from bs4 import BeautifulSoup, Tag
import requests
from mdutils import MdUtils


class NotesType(StrEnum):
    UNKNOWN = 'Unknown'
    BOOLEAN = 'Boolean'
    BYTE = 'Byte'
    INTEGER = 'Integer'
    LONG = 'Long'
    SINGLE = 'Single'
    DOUBLE = 'Double'
    CURRENCY = 'Currency'
    STRING = 'String'
    VARIANT = 'Variant'
    OBJECT = 'Object'
    VOID = 'Void'


class Type:
    def __init__(self, type=NotesType.UNKNOWN, is_array=False, type_name=''):
        self.type = type
        self.is_array = is_array
        self.type_name = first_letter_upper(type_name) if type_name and type == NotesType.OBJECT else type_name

    def __str__(self):
        type = self.type_name if self.type_name else self.type.value

        if self.is_array:
            type = f'{type}[]'
        return type


class AccessType(StrEnum):
    UNKNOWN = 'unknown'
    READ = 'read'
    WRITE = 'write'
    READ_WRITE = 'read-write'


class SectionEntry:
    def __init__(self, title: str, url=''):
        self.title = title
        self.url = url


class Section:
    def __init__(self, title: str, entries: List[SectionEntry]):
        self.title = title
        self.entries = entries


class NotesComponent:
    def __init__(self, name: str, description='', url=''):
        self.name = name
        self.description = re.sub(r'\s+', ' ', description.replace('\n', ' '))
        self.url = url


class NotesProperty(NotesComponent):
    def __init__(self, name: str, value_type: Type, access_type: AccessType, description='', url=''):
        super().__init__(name, description, url)

        self.value_type = value_type
        self.access_type = access_type
    

class NotesParameter(NotesComponent):
    def __init__(self, name: str, type: Type, optional=False, com_handling=False, description='', url=''):
        super().__init__(re.sub(r'[^a-zA-Z_]+', '', name), description, url)

        self.optional = optional
        self.type = type
        self.com_handling = com_handling


class NotesMethod(NotesComponent):
    @dataclass(kw_only=True)
    class Callouts:
        is_optional: Callable[[NotesParameter], bool] = None
        get_param_type: Callable[[NotesParameter], Type] = None
        get_return_type: Callable[[NotesMethod], Type] = None


    def __init__(self, clazz: str, name: str, return_type: Type, params: List[NotesParameter], description='', url=''):
        super().__init__(name, description, url)

        self.clazz = clazz
        self.return_type = return_type
        self.params = params

    def as_markdown_text(self):
        return self._markdown_file().get_md_text()
    
    def as_markdown(self, path=''):
        return self._markdown_file(path).create_md_file()

    def _markdown_file(self, path='') -> MdUtils:
        file = MdUtils(os.path.join(path, f'{self.name}.md'))
        file.new_header(title=f'{f"{self.clazz}." if self.clazz else ""}{self.name}', level=1)

        if self.description:
            file.new_line(self.description)

        if len(self.params):
            table_values = ['Type', 'Name', 'Description']
            columns = len(table_values)
            [table_values.extend([
                str(param.type),
                param.name,
                param.description
            ]) for param in self.params]

            file.new_header(title='Parameters', level=2)
            file.new_table(text=table_values, columns=columns, rows=int(len(table_values) / columns))

        if self.return_type.type != NotesType.VOID:
            file.new_header(title='Return type', level=2)
            file.new_line(str(self.return_type))
        
        return file


class NotesClass(NotesComponent):
    def __init__(self, name: str, properties: List[NotesProperty], methods: List[NotesMethod], description='', url=''):
        super().__init__(name, description, url)

        self.properties = properties
        self.methods = methods

    def as_markdown_text(self, method_details=False):
        return self._markdown_file(method_details=method_details).get_md_text()
    
    def as_markdown(self, path='', method_details=False):
        return self._markdown_file(path, method_details=method_details).create_md_file()

    def _markdown_file(self, path='', method_details=False) -> MdUtils:
        file = MdUtils(os.path.join(path, f'{self.name}.md'))
        file.new_header(title=self.name, level=1)

        if self.description:
            file.new_line(self.description)

        if len(self.properties):
            table_values = ['Type', 'Name', 'Access', 'Description']
            columns = len(table_values)

            [table_values.extend([
                str(property.value_type),
                property.name,
                property.access_type.value,
                property.description
            ]) for property in self.properties]

            file.new_header(title='Properties', level=2)
            file.new_table(text=table_values, columns=columns, rows=int(len(table_values) / columns))

        if len(self.methods):
            file.new_header(title='Methods', level=2)

            for method in self.methods:
                file.new_header(title=f'{method.return_type} {method.name}', level=3)

                if method.description:
                    file.new_line(method.description)

                table_values = ['Type', 'Name', 'Description']
                columns = len(table_values)
                [table_values.extend([
                    param.type,
                    param.name,
                    param.description
                ]) for param in method.params]

                file.new_table(text=table_values, columns=columns, rows=int(len(table_values) / columns))

        return file
    

TYPE_MAPPINGS = (
    ['bool', NotesType.BOOLEAN],
    ['true', NotesType.BOOLEAN],
    ['false', NotesType.BOOLEAN],
    ['byte', NotesType.BYTE],
    ['int', NotesType.INTEGER],
    ['long', NotesType.LONG],
    ['single', NotesType.SINGLE],
    ['double', NotesType.DOUBLE],
    ['currency', NotesType.CURRENCY],
    ['string', NotesType.STRING],
    [r'notes\w+', NotesType.OBJECT],
    ['variant', NotesType.VARIANT],
    ['void', NotesType.VOID],
)


def first_letter(s: str, upper: bool):
    return f'{s[0].upper() if upper else s[0].lower()}{s[1:]}' if s else s


def first_letter_lower(s: str):
    return first_letter(s, False)


def first_letter_upper(s: str):
    return first_letter(s, True)


def map_type(type_string: str, mappings: Tuple[Tuple[str, any]]) -> any:
    mapped_type = None

    if type_string and mappings:
        for key, t in mappings:
            if re.search(key, type_string, re.IGNORECASE):
                mapped_type = t
                break

    return mapped_type
    

def evaluate_tag_text(tag: Tag):
    # HINT: Had to use encode_contents here because for some reason the titles on
    # the page contain newline which are not parsed properly by BeautifulSoup4,
    # using the text property. 'Data type' for example would end up as 'Data'.
    return tag.encode_contents().replace(b'\n', b' ').decode() if tag else ''


def evaluate_type_string_from_description(description: str):
    type_string = ''
    description = description.strip()

    if description:
        type_string_regex = r'^((\w|-)+)(\s*\.\s*|$)'  # Type must be a word (dashes are allowed).
        match = re.match(type_string_regex, description, flags=re.IGNORECASE)

        if match:
            type_string = match.group(1)
            description = re.sub(type_string_regex, '', description, flags=re.IGNORECASE)

    return type_string, description


def evaluate_type_from_description(description: str, mappings: Tuple[Tuple[str, any]], default_type=None) -> Tuple[any, str, str]:
    description = description.strip() if description else ''
    type = default_type
    type_string = ''

    if description:
        type_string, description_new = evaluate_type_string_from_description(description)
        type_temp = map_type(type_string, mappings)

        if type_temp:
            type = type_temp
            description = description_new

    return type, description, type_string


def evaluate_access_type(description: str) -> Tuple[AccessType, str]:
    return evaluate_type_from_description(description, (
        ['read-write', AccessType.READ_WRITE],
        ['read-only', AccessType.READ],
        ['write-only', AccessType.WRITE],
    ), AccessType.UNKNOWN)[0:2]


def evaluate_value_type(description: str) -> Tuple[Type, str]:
    notes_type, description, type_string = evaluate_type_from_description(description, TYPE_MAPPINGS, NotesType.UNKNOWN)
    value_type = Type(notes_type, re.search('array', description, re.IGNORECASE), type_string)

    if description:
        if value_type.type == NotesType.UNKNOWN:
            notes_type_temp = map_type(description, TYPE_MAPPINGS)

            if notes_type_temp:
                value_type.type = notes_type_temp

    return value_type, description


def read_json(path: str):
    object = None

    if os.path.exists(path):
        with open(path, 'r', encoding='utf-8') as f:
            object = json.loads(f.read())
    return object


def write_json(path: str, object: any):
    with open(path, 'w', encoding='utf-8') as f:
        f.write(json.dumps(object, indent=2))


def parse_html(html: str) -> BeautifulSoup:
    return BeautifulSoup(html, features='lxml')


def get_site(url: str) -> BeautifulSoup:
    KEY_PATH = 'path'
    KEY_ENCODING = 'encoding'
    SITES_DIR = 'sites'
    SITES_FILE = 'sites.json'

    object = read_json(SITES_FILE)
    site_dict = object if object else {}
    url_hash = sha256(url.encode())
    path = os.path.join(SITES_DIR, f'{url_hash.hexdigest()}.html')
    content = ''

    # If site not cached yet, cache it.
    if url not in site_dict or not os.path.exists(site_dict[url][KEY_PATH]):
        response = requests.get(url)

        if response.status_code == 200:
            content = response.text.encode().decode()
            encoding = response.encoding
            site_dict[url] = {
                KEY_PATH: path,
                KEY_ENCODING: encoding
            }

            with open(path, 'w', encoding=encoding) as f:
                f.write(content)
            write_json(SITES_FILE, site_dict)

    if url in site_dict:
        entry = site_dict[url]
        path = entry[KEY_PATH]

        if os.path.exists(path):
            # Read as UTF-8, no matter what the original encoding was (why?).
            with open(path, 'r', encoding='utf-8') as f:
                content = f.read()

                # Get rid of all newlines between words as they cause troubles with BeautifulSoup.
                content = re.sub(r'(?<=\w)\n(?=\w)', ' ', content)

    return parse_html(content)


def get_section(soup: BeautifulSoup, title_pattern: str, base_url='') -> Section:
    section = None
    html_section_titles: List[Tag] = list(soup.find_all(attrs={'class': 'topic/title'}))  # https://stackoverflow.com/a/5041056
    html_section_titles = list(filter(lambda t: re.match(title_pattern, evaluate_tag_text(t).strip()), html_section_titles))

    if len(html_section_titles) == 1:
        html_section_title = html_section_titles[0]
        html_section = html_section_title.parent
        entries: List[SectionEntry] = []

        for t in html_section.find_all(lambda t: t != html_section_title, recursive=False):  # Find all section elements excluding the title.
            entry = SectionEntry(t.text.replace('\n', ' ').strip())

            if t.a:
                entry.url = urljoin(base_url, t.a.get('href'))
            entries.append(entry)
        section = Section(title_pattern, entries)
    return section


def get_name_and_description(soup: BeautifulSoup) -> Tuple[str, str]:
    name = re.match('\w+', soup.h1.text.strip()).group(0).strip()
    description = evaluate_tag_text(soup.find('p', {'class': 'topic/shortdesc'})).strip()

    return name, description


def get_table_from_header_pattern(soup: BeautifulSoup, pattern: str) -> Tuple[str, Dict[str, List[str]]]:
    table: Tag = None
    table_header = soup.find(lambda t: t.name == 'th' and re.match(pattern, t.text))
    header_text: str = None
    table_dict = {}

    if table_header:
        table = table_header.find_parent('table')
        header_text = table_header.text

    if table:
        table_body = table.find('tbody')

        if table_body:
            table_headers = [evaluate_tag_text(th) for th in table.find_all('th')]
            table_rows = table_body.find_all('tr')

            for tr in table_rows:
                row_cells = tr.find_all('td')

                for i, cell in enumerate(row_cells):
                    header = table_headers[i]
                    
                    if header not in table_dict:
                        table_dict[header] = []
                    table_dict[header].append(cell.text)
            
    return header_text, table_dict


def get_section_from_table_header_pattern(soup: BeautifulSoup, pattern: str) -> Section:
    header_text, table_dict = get_table_from_header_pattern(soup, pattern)
    section: Section = None

    if header_text:
        section = Section(header_text, [])
        entries = section.entries
        names = table_dict[header_text]
        descriptions = table_dict['Description']

        for name, description in zip(names, descriptions):
            entries.append(SectionEntry(name))
            entries.append(SectionEntry(description))

    return section


def get_property(url: str) -> NotesProperty:
    soup = get_site(url)
    name, description = get_name_and_description(soup)
    data_type_section = get_section(soup, 'Data type')
    value_type = Type()
    access_type, description = evaluate_access_type(description)

    if data_type_section:
        entries = data_type_section.entries
        description = ' '.join(map(lambda e: e.title, entries)) if len(entries) > 0 else ''
        value_type, description = evaluate_value_type(description)
    return NotesProperty(name, value_type, access_type, description, url)


def get_properties(soup: BeautifulSoup, base_url: str):
    section = get_section(soup, 'Properties', base_url)
    properties = []
    
    if section:
        for entry in section.entries:
            url = entry.url

            if url:
                properties.append(get_property(url))

    return properties


def get_method(url: str, callouts: NotesMethod.Callouts=None):
    soup = get_site(url)
    name, description = get_name_and_description(soup)
    defined_in_section = get_section(soup, 'Defined in')
    clazz = defined_in_section.entries[0].title if defined_in_section and len(defined_in_section.entries) else ''
    params_section = get_section(soup, r'Parameter(s)?')
    return_value_section = get_section(soup, 'Return value')
    params = []
    return_type = Type()
    callouts = callouts if callouts else NotesMethod.Callouts()

    # If params section could not be found, params are probably listed in a table.
    if not params_section:
        params_section = get_section_from_table_header_pattern(soup, r'Parameter(s)?')

    if params_section:
        entries = []

        # Put parameter entry and description entries into one list entry.
        for entry in params_section.entries:
            title = entry.title.strip()

            # If entry has just one word, assume it's the parameter name.
            if title and len(title.split(' ')) == 1:
                entries.append([entry, []])
            elif len(entries):
                entries[-1][1].append(entry.title)

        for entry in entries:
            title = entry[0].title

            if title and title.lower() != 'none':
                param_name = re.sub('[^a-zA-Z_]', '', title)
                param_description = ' '.join(entry[1]).strip()
                param_type, param_description = evaluate_value_type(param_description)
                param = NotesParameter(
                    param_name,
                    param_type,
                    re.search('optional', param_description, re.IGNORECASE) or re.search('optional', title, re.IGNORECASE),
                    re.search('COM', param_description),
                    param_description,
                    url
                )

                # If evaluation is not sure if the parameter is optional, use callout if provided.
                if not param.optional and re.search('(default|omit)', param_description, re.IGNORECASE) and callouts.is_optional:
                    param.optional = callouts.is_optional(param)

                # If evaluation is not sure about the parameter type, use callout if provided.
                if param.type.type == NotesType.UNKNOWN and callouts.get_param_type:
                    param.type = callouts.get_param_type(param)

                params.append(param)

    # If return value section could not be found, return value is probably listed in a table.
    if not return_value_section:
        return_value_section = get_section_from_table_header_pattern(soup, r'Return value(s)?')

    if return_value_section:
        for entry in return_value_section.entries:
            title = entry.title

            if not title or title.lower() == 'none':
                return_type_temp = Type(NotesType.VOID)
            else:
                return_type_temp, _ = evaluate_value_type(title)

            if return_type_temp.type != NotesType.UNKNOWN:
                return_type = return_type_temp
                break
    else:
        return_type = Type(NotesType.VOID)

    method = NotesMethod(clazz, name, return_type, params, description, url)

    if method.return_type.type == NotesType.UNKNOWN and callouts.get_return_type:
        method.return_type = callouts.get_return_type(method)
    return method


def get_methods(soup: BeautifulSoup, base_url: str, callouts: NotesMethod.Callouts=None):
    section = get_section(soup, 'Methods', base_url)
    methods = []

    if section:
        for entry in section.entries:
            url = entry.url

            if url:
                methods.append(get_method(url, callouts))

    return methods


def get_class(url: str, callouts: NotesMethod.Callouts=None) -> NotesClass:
    soup = get_site(url)
    name, description = get_name_and_description(soup)
    properties = get_properties(soup, url)
    methods = get_methods(soup, url, callouts)

    return NotesClass(name, properties, methods, description, url)
