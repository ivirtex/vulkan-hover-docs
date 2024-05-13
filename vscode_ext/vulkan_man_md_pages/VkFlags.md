# VkFlags(3) Manual Page

## Name

VkFlags - Vulkan bitmasks



## <a href="#_c_specification" class="anchor"></a>C Specification

A collection of flags is represented by a bitmask using the type
[VkFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFlags.html):

``` c
// Provided by VK_VERSION_1_0
typedef uint32_t VkFlags;
```

## <a href="#_description" class="anchor"></a>Description

Bitmasks are passed to many commands and structures to compactly
represent options, but [VkFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFlags.html) is not used directly in
the API. Instead, a `Vk*Flags` type which is an alias of
[VkFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFlags.html), and whose name matches the corresponding
`Vk*FlagBits` that are valid for that type, is used.

Any `Vk*Flags` member or parameter used in the API as an input **must**
be a valid combination of bit flags. A valid combination is either zero
or the bitwise OR of valid bit flags.

An individual bit flag is valid for a `Vk*Flags` type if it would be a
<a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#fundamentals-validusage-enums"
target="_blank" rel="noopener">valid enumerant</a> when used with the
equivalent `Vk*FlagBits` type, where the bits type is obtained by taking
the flag type and replacing the trailing `Flags` with `FlagBits`. For
example, a flag value of type
[VkColorComponentFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkColorComponentFlags.html) **must** contain
only bit flags defined by
[VkColorComponentFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkColorComponentFlagBits.html).

Any `Vk*Flags` member or parameter returned from a query command or
otherwise output from Vulkan to the application **may** contain bit
flags undefined in its corresponding `Vk*FlagBits` type. An application
**cannot** rely on the state of these unspecified bits.

Only the low-order 31 bits (bit positions zero through 30) are available
for use as flag bits.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr class="odd">
<td class="icon"><em></em></td>
<td class="content">Note
<p>This restriction is due to poorly defined behavior by C compilers
given a C enumerant value of <code>0x80000000</code>. In some cases
adding this enumerant value may increase the size of the underlying
<code>Vk*FlagBits</code> type, breaking the ABI.</p></td>
</tr>
</tbody>
</table>

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkColorComponentFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkColorComponentFlags.html),
[VkFlags64](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFlags64.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkFlags"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
