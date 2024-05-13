# VK_DEFINE_NON_DISPATCHABLE_HANDLE(3) Manual Page

## Name

VK_DEFINE_NON_DISPATCHABLE_HANDLE - Declare a non-dispatchable object
handle



## <a href="#_c_specification" class="anchor"></a>C Specification

`VK_DEFINE_NON_DISPATCHABLE_HANDLE` defines a <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#fundamentals-objectmodel-overview"
target="_blank" rel="noopener">non-dispatchable handle</a> type.

``` c
// Provided by VK_VERSION_1_0

#ifndef VK_DEFINE_NON_DISPATCHABLE_HANDLE
    #if (VK_USE_64_BIT_PTR_DEFINES==1)
        #define VK_DEFINE_NON_DISPATCHABLE_HANDLE(object) typedef struct object##_T *object;
    #else
        #define VK_DEFINE_NON_DISPATCHABLE_HANDLE(object) typedef uint64_t object;
    #endif
#endif
```

## <a href="#_description" class="anchor"></a>Description

- `object` is the name of the resulting C type.

Most Vulkan handle types, such as [VkBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBuffer.html), are
non-dispatchable.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr class="odd">
<td class="icon"><em></em></td>
<td class="content">Note
<p>The <code>vulkan_core.h</code> header allows the <a
href="VK_DEFINE_NON_DISPATCHABLE_HANDLE.html">VK_DEFINE_NON_DISPATCHABLE_HANDLE</a>
and <a href="VK_NULL_HANDLE.html">VK_NULL_HANDLE</a> definitions to be
overridden by the application. If <a
href="VK_DEFINE_NON_DISPATCHABLE_HANDLE.html">VK_DEFINE_NON_DISPATCHABLE_HANDLE</a>
is already defined when <code>vulkan_core.h</code> is compiled, the
default definitions for <a
href="VK_DEFINE_NON_DISPATCHABLE_HANDLE.html">VK_DEFINE_NON_DISPATCHABLE_HANDLE</a>
and <a href="VK_NULL_HANDLE.html">VK_NULL_HANDLE</a> are skipped. This
allows the application to define a binary-compatible custom handle which
<strong>may</strong> provide more type-safety or other features needed
by the application. Applications <strong>must</strong> not define
handles in a way that is not binary compatible - where binary
compatibility is platform dependent.</p></td>
</tr>
</tbody>
</table>

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html), [VkBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBuffer.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_DEFINE_NON_DISPATCHABLE_HANDLE"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
