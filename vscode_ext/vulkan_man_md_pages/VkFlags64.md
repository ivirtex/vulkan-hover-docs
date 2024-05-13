# VkFlags64(3) Manual Page

## Name

VkFlags64 - Vulkan 64-bit bitmasks



## <a href="#_c_specification" class="anchor"></a>C Specification

A collection of 64-bit flags is represented by a bitmask using the type
[VkFlags64](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFlags64.html):

``` c
// Provided by VK_VERSION_1_3, VK_KHR_synchronization2
typedef uint64_t VkFlags64;
```

## <a href="#_description" class="anchor"></a>Description

When the 31 bits available in [VkFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFlags.html) are insufficient,
the [VkFlags64](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFlags64.html) type can be passed to commands and
structures to represent up to 64 options. [VkFlags64](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFlags64.html) is
not used directly in the API. Instead, a `Vk*Flags2` type which is an
alias of [VkFlags64](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFlags64.html), and whose name matches the
corresponding `Vk*FlagBits2` that are valid for that type, is used.

Any `Vk*Flags2` member or parameter used in the API as an input **must**
be a valid combination of bit flags. A valid combination is either zero
or the bitwise OR of valid bit flags.

An individual bit flag is valid for a `Vk*Flags2` type if it would be a
<a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#fundamentals-validusage-enums"
target="_blank" rel="noopener">valid enumerant</a> when used with the
equivalent `Vk*FlagBits2` type, where the bits type is obtained by
taking the flag type and replacing the trailing `Flags2` with
`FlagBits2`. For example, a flag value of type
[VkAccessFlags2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccessFlags2KHR.html) **must** contain only bit
flags defined by [VkAccessFlagBits2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccessFlagBits2KHR.html).

Any `Vk*Flags2` member or parameter returned from a query command or
otherwise output from Vulkan to the application **may** contain bit
flags undefined in its corresponding `Vk*FlagBits2` type. An application
**cannot** rely on the state of these unspecified bits.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr class="odd">
<td class="icon"><em></em></td>
<td class="content">Note
<p>Both the <code>Vk*FlagBits2</code> type, and the individual bits
defined for that type, are defined as <code>uint64_t</code> integers in
the C API. This is in contrast to the 32-bit types, where the
<code>Vk*FlagBits</code> type is defined as a C <code>enum</code> and
the individual bits as enumerants belonging to that <code>enum</code>.
As a result, there is less compile time type checking possible for the
64-bit types. This is unavoidable since there is no sufficiently
portable way to define a 64-bit <code>enum</code> type in C99.</p></td>
</tr>
</tbody>
</table>

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_synchronization2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_synchronization2.html),
[VK_VERSION_1_3](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_3.html), [VkFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFlags.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkFlags64"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
