/*
bham or the "blocky hypertext abstraction markup"
is an attempt to take what is good about languages
like haml, jade, slim, etc. and port it to Go, but
not blindly. It will take into account the capabilities
of Go's template libraries to parse directly into the
internal template structures that the stdlib template
libraries use to provide both speed and interoperability
with standard Go templates.

Plain Text

Most web templates are mostly html content, or just plain text
content, as opposed to executable code. If bham doesn't think
a line can be executed, it will just output the raw version to
html. If bham thinks your line could be executable, but is going
to error it should return you an error during the parse step.

  %first
    %second
      %frist
        <a href="#">Showdown</a>

would be turned into the html content:

  <first>
    <second>
      <frist>
        <a href="#">Showdown</a>
      </frist>
    </second>
  </first>

You should notice that placing a % before a word will turn it into
a tag, while at the same time the a tag was sent though as well. For
those of us wondering about automatic escaping, putting plain html
into a bham template will not escape any characters.

Todo: Escaping

To have a line display as written, whether that line would be executed
or is html code just add a \ as the first non-whitespace character,
and that line will be html escaped before it is inserted into the
template.

Attributes

The normal attribute syntax is the html style attributes, in the form
of

  %script(type="text/javascript" src="/js/{{.script}}.js")

TODO: To set attributes on a tag by a map, put the map in single { }, multiple
maps map be put in the { } and the attribute will be put together. If
attributes are present in multiple maps, but they aren't of the key
class or id, the first map with that key will be the value appearing.

 %script(type="text/javascript){ .script }

Classes and ID's

Period and pound signs are used in the same manner as haml. Use them
to quickly add classes or an id to tags. Multiple class names are each
added to the class attribute, while multiple id parts will be joined
with underscores in the compiled form.

  %span#tesla(id="tower")
    %img.watt.ohm(src="img1.png")

would be compiled into

  <span id="tesla_tower">
    <img class="watt ohm" src="img1.png"></img>
  </span>

If you do not specify the tag, it will be assumed to be a div. The
following two lines are identical after compilation.

  %div.classy#restaurant
  .classy#restaurant

TODO: Empty tags. Tags that have no content, and are on the autoclose
list of tags (see bham.AutocloseTags) will either have no closing tag,
or will end with />

TODO: Whitespace non-removal. By default, bham likes to remove
whitespace from your templates. It still needs some tweaking. But
maybe there needs to be some operation to keep whitespace?

TODO Reference Structs

There should be some interface definition to allow developers to
use the an instance of a struct that implements an interface to
use that instance to template out some attributes, content, etc.

Doctypes

Doctypes are placed near the beginning of files and start with three
exclamation marks. There is a list of default Doctypes in the
Doctypes variable. You can add or modify your own doctypes. Here's
an example:

  !!! Strict

would be compiled into

  <!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Strict//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-strict.dtd">

Todo Comments

Adding a forward slash should turn the more indented lines into an
html comment. Adding [] should make it a conditional comment for IE.
Adding a -# should make it a silent comment that doesn't show up in
the compiled form.

Filters

Sometimes you need to embed a bit of javascript or css into your
html template, in this case you can use a filter to automatically
turn it into a different format. For example

  :javascript
    $(".name").html("Hello World");

would compile into

  <script type="text/javascipt">
    $(".name").html("Hello World");
  </script>

Currently you cannot embed values into lines that are embedded into
filters, this will be fixed.

Template code

Pretty much the same as the standard go templates, but I'm not good
at complex  number parsing.
*/
package bham
