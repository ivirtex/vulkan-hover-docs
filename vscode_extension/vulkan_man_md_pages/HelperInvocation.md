# HelperInvocation(3) Manual Page

## Name

HelperInvocation - Indication of whether a fragment shader is a helper
invocation



## <a href="#_description" class="anchor"></a>Description

`HelperInvocation`  
Decorating a variable with the `HelperInvocation` built-in decoration
will make that variable contain whether the current invocation is a
helper invocation. This variable is non-zero if the current fragment
being shaded is a helper invocation and zero otherwise. A helper
invocation is an invocation of the shader that is produced to satisfy
internal requirements such as the generation of derivatives.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr>
<td class="icon"><em></em></td>
<td class="content">Note
<p>It is very likely that a helper invocation will have a value of
<code>SampleMask</code> fragment shader input value that is
zero.</p></td>
</tr>
</tbody>
</table>

Valid Usage

- <a href="#VUID-HelperInvocation-HelperInvocation-04239"
  id="VUID-HelperInvocation-HelperInvocation-04239"></a>
  VUID-HelperInvocation-HelperInvocation-04239  
  The `HelperInvocation` decoration **must** be used only within the
  `Fragment` `Execution` `Model`

- <a href="#VUID-HelperInvocation-HelperInvocation-04240"
  id="VUID-HelperInvocation-HelperInvocation-04240"></a>
  VUID-HelperInvocation-HelperInvocation-04240  
  The variable decorated with `HelperInvocation` **must** be declared
  using the `Input` `Storage` `Class`

- <a href="#VUID-HelperInvocation-HelperInvocation-04241"
  id="VUID-HelperInvocation-HelperInvocation-04241"></a>
  VUID-HelperInvocation-HelperInvocation-04241  
  The variable decorated with `HelperInvocation` **must** be declared as
  a boolean value

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#HelperInvocation"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
