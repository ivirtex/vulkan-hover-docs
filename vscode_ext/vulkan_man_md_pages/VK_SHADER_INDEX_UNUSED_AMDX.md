# VK_SHADER_INDEX_UNUSED_AMDX(3) Manual Page

## Name

VK_SHADER_INDEX_UNUSED_AMDX - Sentinel for an unused shader index



## <a href="#_c_specification" class="anchor"></a>C Specification

`VK_SHADER_INDEX_UNUSED_AMDX` is a special shader index used to indicate
that the created node does not override the index. In this case, the
shader index is determined through other means. It is defined as:

``` c
#define VK_SHADER_INDEX_UNUSED_AMDX       (~0U)
```

## <a href="#_see_also" class="anchor"></a>See Also

[VK_AMDX_shader_enqueue](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_AMDX_shader_enqueue.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_SHADER_INDEX_UNUSED_AMDX"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
