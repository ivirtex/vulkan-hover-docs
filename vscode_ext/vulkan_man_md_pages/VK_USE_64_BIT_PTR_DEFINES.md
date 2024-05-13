# VK_USE_64_BIT_PTR_DEFINES(3) Manual Page

## Name

VK_USE_64_BIT_PTR_DEFINES - Defines whether non-dispatchable handles are
a 64-bit pointer type or a 64-bit unsigned integer type



## <a href="#_c_specification" class="anchor"></a>C Specification

`VK_USE_64_BIT_PTR_DEFINES` defines whether the default non-dispatchable
handles are declared using either a 64-bit pointer type or a 64-bit
unsigned integer type.

`VK_USE_64_BIT_PTR_DEFINES` is set to '1' to use a 64-bit pointer type
or any other value to use a 64-bit unsigned integer type.

``` c
// Provided by VK_VERSION_1_0

#ifndef VK_USE_64_BIT_PTR_DEFINES
    #if defined(__LP64__) || defined(_WIN64) || (defined(__x86_64__) && !defined(__ILP32__) ) || defined(_M_X64) || defined(__ia64) || defined (_M_IA64) || defined(__aarch64__) || defined(__powerpc64__) || (defined(__riscv) && __riscv_xlen == 64)
        #define VK_USE_64_BIT_PTR_DEFINES 1
    #else
        #define VK_USE_64_BIT_PTR_DEFINES 0
    #endif
#endif
```

## <a href="#_description" class="anchor"></a>Description

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr class="odd">
<td class="icon"><em></em></td>
<td class="content">Note
<p>The <code>vulkan_core.h</code> header allows the
<code>VK_USE_64_BIT_PTR_DEFINES</code> definition to be overridden by
the application. This allows the application to select either a 64-bit
pointer type or a 64-bit unsigned integer type for non-dispatchable
handles in the case where the predefined preprocessor check does not
identify the desired configuration.</p></td>
</tr>
</tbody>
</table>

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr class="odd">
<td class="icon"><em></em></td>
<td class="content">Note
<p>This macro was introduced starting with the Vulkan 1.2.174 headers,
and its availability can be checked at compile time by requiring
<code>VK_HEADER_VERSION</code><code> &gt;= 174</code>.</p>
<p>It is not available if you are using older headers, such as may be
shipped with an older Vulkan SDK. Developers requiring this
functionality may wish to include a copy of the current Vulkan headers
with their project in this case.</p></td>
</tr>
</tbody>
</table>

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_USE_64_BIT_PTR_DEFINES"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
