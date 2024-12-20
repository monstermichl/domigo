import argparse
import os
import re
from typing import List, Tuple

from pick import pick
from notes_classes import AccessType, NotesComponent, NotesMethod, NotesParameter, NotesProperty, NotesType, Type, first_letter_lower, first_letter_upper, get_class


def get_urls() -> str:
    parser = argparse.ArgumentParser()

    parser.add_argument('--url', '-u', required=True, help='URL to LotusScript class page', action='append')
    args = parser.parse_args()

    return args.url


def get_go_type(type: Type) -> Tuple[str, str]:
    cast_type = type.type
    return f'[]{cast_type}' if type.is_array else cast_type, cast_type


def get_go_params(params: List[NotesParameter], is_call) -> List[str]:
    def function(p: NotesParameter):
        name = p.name

        if is_call:
            if p.type.type == NotesType.OBJECT and not p.type.is_array:
                name = f'{name}.com().Dispatch()'  # Make sure to pass the dispatch pointer, not the object.
            s = name
        else:
            s = f'{name} {get_go_type(p.type)[0]}'
        return s

    return list(map(function, params))


def properties_stubs(class_name: str, props: List[NotesProperty]) -> Tuple[List[str], List[str]]:
    class_name = first_letter_upper(class_name)
    stubs = []
    test_stubs = []

    properties_comment = '/* --------------------------------- Properties --------------------------------- */'
    stubs.append(properties_comment)
    test_stubs.append(properties_comment)

    for p in props:
        name = p.name
        value_type = p.value_type
        notes_type = value_type.type
        is_array = value_type.is_array
        specify_string = ' /* TODO: Specify. */' if notes_type == NotesType.UNKNOWN else ''
        go_method_return_type, go_cast_type = get_go_type(value_type)
        url = p.url

        if p.access_type == AccessType.UNKNOWN:
            info = f'/* TODO: Access type for {name} could not be evaluated, check yourself if getter/setter is needed. */'
            p.access_type = AccessType.READ_WRITE

            stubs.append(info)
            test_stubs.append(info)

        if p.access_type in [AccessType.READ, AccessType.READ_WRITE]:
            stubs.extend([
                f'/* {url} */',
                f'func ({receiver} {class_name}) {name}() ({go_method_return_type}{specify_string}, error) {{',
            ])

            if notes_type == NotesType.OBJECT:
                if is_array:
                    stubs.extend([
                        f'return com.GetObjectArrayProperty({receiver}.com(), New{go_cast_type}, "{name}")',
                    ])
                else:
                    stubs.extend([
                        f'    dispatchPtr, err := {receiver}.com().GetObjectProperty("{name}")',
                        f'    return New{go_cast_type}(dispatchPtr), err',
                    ])
            else:
                stubs.extend([
                    f'    val{"s" if is_array else ""}, err := {receiver}.com().Get{"Array" if is_array else ""}Property("{name}")',
                    f'    return helpers.Cast{"Slice" if is_array else "Value"}[{go_cast_type}](val{"s" if is_array else ""}), err',
                ])

            stubs.extend([
                '}',
                '',
            ])

            test_stubs.extend([
                f'/* {url} */',
                f'func Test{name}(t *testing.T) {{',
                f'    _, err := {variable}.{name}()',
                '    require.Nil(t, err)',
                '}',
                '',
            ])

        if p.access_type in [AccessType.WRITE, AccessType.READ_WRITE]:
            stubs.extend([
                f'/* {url} */',
                f'func ({receiver} {class_name}) Set{name}(v {go_method_return_type}) error {{',
                f'    return {receiver}.com().PutProperty("{name}", v)',
                '}',
            ])

            test_stubs.extend([
                f'/* {url} */',
                f'func TestSet{name}(t *testing.T) {{',
                f'    s, err := {variable}.{name}()',
                '    require.Nil(t, err)',
                '',
                f'    err = {variable}.Set{name}(s)',
                '    require.Nil(t, err)',
                '}',
            ])

    return stubs, test_stubs


def method_stubs(class_name: str, methods: List[NotesMethod]) -> Tuple[List[str], List[str]]:
    class_name = first_letter_upper(class_name)
    stubs = []
    test_stubs = []

    methods_comment = '/* --------------------------------- Methods ------------------------------------ */'
    stubs.append(methods_comment)
    test_stubs.append(methods_comment)

    for m in methods:
        name = m.name
        return_type = m.return_type
        is_array = return_type.is_array
        notes_type = return_type.type
        void = notes_type == NotesType.VOID
        go_return_type, go_cast_type = get_go_type(return_type)
        go_method_return_type = f"({go_return_type}, error)" if not void else "error"
        url = m.url
        params = m.params
        optional_params = [p for p in params if p.optional]
        has_optional_params = len(optional_params)

        stubs.append(f'/* {url} */')

        # Prepare optional params structs and functions.
        if has_optional_params:
            first_letter_lower_name = f'{first_letter_lower(class_name)}{first_letter_upper(name)}'
            params_type_name = f'{first_letter_lower_name}Params'
            param_type_name = f'{first_letter_lower_name}Param'

            def p_string(p: NotesParameter, pointer: bool):
                return f'{p.name} {"*" if pointer else ""}{get_go_type(p.type)[0]}'

            # Prepare helper struct.
            stubs.append(f'type {params_type_name} struct {{')
            stubs.extend([f'    {p_string(p, True)}' for p in optional_params])
            stubs.extend([
                '}',
                '',
                f'type {param_type_name} func(*{params_type_name})',
                '',
            ])

            for p in optional_params:
                stubs.extend([
                    f'func With{class_name}{first_letter_upper(name)}{first_letter_upper(p.name)}({p_string(p, False)}) {param_type_name} {{',
                    f'    return func(c *{params_type_name}) {{',
                    f'        c.{p.name} = &{p.name}',
                    '    }',
                    '}',
                    '',
                ])

        method_params = ", ".join(get_go_params([p for p in params if not p.optional], False))

        if has_optional_params:
            if method_params:
                method_params = f'{method_params}, '
            method_params = f'{method_params}params ...{param_type_name}'
        stubs.append(f'func ({receiver} {class_name}) {name}({method_params}) {go_method_return_type} {{')

        # Prepare call params.
        STRUCT_VARIABLE_NAME = 'paramsStruct'
        call_params = ", ".join(get_go_params([p for p in params if not p.optional], True))

        if has_optional_params:
            PARAM_VARIABLE_NAME = 'paramsOrdered'

            stubs.extend([
                f'    {STRUCT_VARIABLE_NAME} := &{params_type_name}{{}}',
                f'    paramsOrdered := []interface{{}}{{{call_params}}}',
                '',
                '    for _, p := range params {',
                f'        p({STRUCT_VARIABLE_NAME})',
                '    }',
                '',
            ])

            def parameter_check(ps: List[NotesParameter], indent=1):
                lines = []

                if len(ps) > 0:
                    p = ps[0]
                    p_name = p.name
                    p_is_object_array = p.type.is_array and p.type.type == NotesType.OBJECT
                    parameter_name = f'{STRUCT_VARIABLE_NAME}.{p_name}'

                    if p.com_handling:
                        lines.append(f'/* TODO: {p_name} probably needs to be handled differently with COM. */')

                    lines.extend([
                        f'if {parameter_name} != nil {{',
                        f'    paramsOrdered = append(paramsOrdered, {"DispatchSlice(" if p_is_object_array else ""}*{parameter_name}{")" if p_is_object_array else ""})',
                        *parameter_check(ps[1:], indent + 1),
                        '}',
                    ])
                return list(map(lambda line: f'{(" " * 4) * indent}{line}', lines))

            stubs.extend(parameter_check(optional_params))
            call_params = f'{PARAM_VARIABLE_NAME}...'
        call_params = f', {call_params}' if call_params else ''

        if notes_type == NotesType.OBJECT:
            if is_array:
                stubs.extend([
                    f'return com.CallObjectArrayMethod({receiver}.com(), New{go_cast_type}, "{name}"{call_params})',
                ])
            else:
                stubs.extend([
                    f'    dispatchPtr, err := {receiver}.com().CallObjectMethod("{name}"{call_params})',
                    f'    return New{go_cast_type}(dispatchPtr), err',
                ])
        elif notes_type == NotesType.VOID:
            stubs.extend([
                f'    _, err := {receiver}.com().CallMethod("{name}"{call_params})',
                '    return err',
            ])
        else:
            stubs.extend([
                f'    val{"s" if is_array else ""}, err := {receiver}.com().Call{"Array" if is_array else ""}Method("{name}"{call_params})',
                f'    return helpers.Cast{"Slice" if is_array else "Value"}[{go_cast_type}](val{"s" if is_array else ""}), err',
            ])

        stubs.extend([
            '}',
            '',
        ])

        test_stubs.extend([
            f'/* {url} */',
            f'func Test{name}(t *testing.T) {{',
        ])
        pass_values_string = '/* TODO: Pass test values. */' if len(params) > 0 else ''

        if notes_type == NotesType.VOID:
            test_stubs.extend([
                f'    err := {variable}.{name}({pass_values_string})',
                '    require.Nil(t, err)',
            ])
        else:
            test_stubs.extend([
                f'    _, err := {variable}.{name}({pass_values_string})',
                '    require.Nil(t, err)',
            ])

        test_stubs.extend([
            '}',
            '',
        ])

    return stubs, test_stubs


def is_optional(param: NotesParameter) -> bool:
    lines = [
        f'Is the parameter {param.name} optional? (y/n)',
        f'See {param.url}',
        '',
        'Description:',
        param.description,
        '',
    ]
    i = input('\n'.join(lines))

    return i[0].lower() == 'y' if len(i) > 0 else False


def get_type(prompt: str, component: NotesComponent, type: Type) -> Type:
    options = [
        NotesType.UNKNOWN,
        NotesType.BOOLEAN,
        NotesType.BYTE,
        NotesType.INTEGER,
        NotesType.LONG,
        NotesType.SINGLE,
        NotesType.DOUBLE,
        NotesType.CURRENCY,
        NotesType.STRING,
        NotesType.VARIANT,
        NotesType.OBJECT,
        NotesType.VOID,
    ]
    _, index = pick(list(map(lambda t: t.value, options)), '\n'.join([
        prompt,
        f'See {component.url}',
        '',
        'Description:',
        component.description,
        '',
    ]))
    type.type = options[index]

    if not type.type_name and type.type == NotesType.OBJECT:
        type.type_name = input('Please specify the type name: ')
    return type


def get_param_type(param: NotesParameter) -> Type:
    return get_type(f'Please select a type for parameter {param.name}', param, param.type)


def get_return_type(method: NotesMethod) -> Type:
    return get_type(f'Please select a return type for method {method.name}', method, method.return_type)


if __name__ == '__main__':
    DATABASE_FILE = 'GoInterface.nsf'

    urls = get_urls()

    for url in urls:
        url_comment = f'/* {url} */'
        c = get_class(url, callouts=NotesMethod.Callouts(
            is_optional=is_optional,
            #get_param_type=get_param_type,
            #get_return_type=get_return_type,
        ))
        class_name = c.name
        class_name_lower = class_name.lower()
        variable = re.sub('notes', '', class_name_lower)  # Remove the "Notes" prefix.
        receiver = variable[0]  # Just use the first letter.
        path = os.path.join('generated', class_name_lower)
        stubs = []
        test_stubs = []

        stubs.extend([
            url_comment,
            'package domigo',
            '',
            'import (',
            '    "github.com/monstermichl/domigo/internal/helpers"',
            '',
            '    ole "github.com/go-ole/go-ole"',
            ')',
            '',
            f'type {class_name} struct {{',
            '    NotesStruct',
            '}',
            '',
            f'func New{class_name}(dispatchPtr *ole.IDispatch) {class_name} {{',
            f'    return {class_name}{{NewNotesStruct(dispatchPtr)}}',
            '}',
            '',
        ])

        test_stubs.extend([
            url_comment,
            f'package {class_name_lower}_test',
            '',
            'import (',
            '    "github.com/monstermichl/domigo"',
            '    testhelpers "github.com/monstermichl/domigo/test/helpers"',
            '    "testing"',
            '',
            '    "github.com/stretchr/testify/require"',
            ')',
            '',
            f'var {variable} domigo.{class_name}',
            '',
            '/* https://pkg.go.dev/testing#hdr-Main */',
            'func TestMain(m *testing.M) {',
            '    testhelpers.Initialize(func(session domigo.NotesSession, db domigo.NotesDatabase) (string, error) {',
            '        var err error',
            f'        {variable}, err = /* TODO: Get variable. */',
            f'        defer {variable}.Release()',
            '',
            '        if err != nil {',
            f'            return "{class_name} could not be created", err',
            '        }',
            '        m.Run()',
            '        return "", nil',
            '    })',
            '}',
            '',
        ])

        # Add properties stubs.
        properties_stubs_temp, properties_test_stubs_temp = properties_stubs(class_name, c.properties)

        stubs.extend(properties_stubs_temp)
        test_stubs.extend(properties_test_stubs_temp)

        # Add method stubs.
        method_stubs_temp, method_test_stubs_temp = method_stubs(class_name, c.methods)
        
        stubs.extend(method_stubs_temp)
        test_stubs.extend(method_test_stubs_temp)

        # Write files.
        with open(f'{path}.go', 'w') as f:
            f.write('\n'.join(stubs))

        with open(f'{path}_test.go', 'w') as f:
            f.write('\n'.join(test_stubs))
