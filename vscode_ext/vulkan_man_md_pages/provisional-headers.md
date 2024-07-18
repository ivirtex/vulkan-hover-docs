# provisional-headers(3) Manual Page

## Name

provisional-headers - Control inclusion of provisional extensions



## <a href="#_description" class="anchor"></a>Description

*Provisional* extensions **should** not be used in production
applications. The functionality defined by such extensions **may**
change in ways that break backwards compatibility between revisions, and
before final release of a non-provisional version of that extension.

Provisional extensions are defined in a separate *provisional header*,
`vulkan_beta.h`, allowing applications to decide whether or not to
include them. The mechanism is similar to <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#boilerplate-wsi-header"
target="_blank" rel="noopener">window system-specific headers</a>:
before including `vulkan_beta.h`, applications **must** include
`vulkan_core.h`.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr>
<td class="icon"><em></em></td>
<td class="content">Note
<p>Sometimes a provisional extension will include a subset of its
interfaces in <code>vulkan_core.h</code>. This may occur if the
provisional extension is promoted from an existing vendor or EXT
extension and some of the existing interfaces are defined as aliases of
the provisional extension interfaces. All other interfaces of that
provisional extension which are not aliased will be included in
<code>vulkan_beta.h</code>.</p></td>
</tr>
</tbody>
</table>

As a convenience for applications, `vulkan.h` conditionally includes
`vulkan_beta.h`. Applications **can** control inclusion of
`vulkan_beta.h` by \#defining the macro `VK_ENABLE_BETA_EXTENSIONS`
before including `vulkan.h`.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr>
<td class="icon"><em></em></td>
<td class="content">Note
<p>Starting in version 1.2.171 of the Specification, all provisional
enumerants are protected by the macro
<code>VK_ENABLE_BETA_EXTENSIONS</code>. Applications needing to use
provisional extensions must always define this macro, even if they are
explicitly including <code>vulkan_beta.h</code>. This is a minor change
to behavior, affecting only provisional extensions.</p></td>
</tr>
</tbody>
</table>

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr>
<td class="icon"><em></em></td>
<td class="content">Note
<p>This section describes the purpose of the provisional header
independently of the specific provisional extensions which are contained
in that header at any given time. The extension appendices for
provisional extensions note their provisional status, and link back to
this section for more information. Provisional extensions are intended
to provide early access for bleeding-edge developers, with the
understanding that extension interfaces may change in response to
developer feedback. Provisional extensions are very likely to eventually
be updated and released as non-provisional extensions, but there is no
guarantee this will happen, or how long it will take if it does
happen.</p></td>
</tr>
</tbody>
</table>

## <a href="#_see_also" class="anchor"></a>See Also

[WSIheaders](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/WSIheaders.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#boilerplate-provisional-header"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
