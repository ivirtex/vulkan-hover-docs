# VK\_SHADER\_INDEX\_UNUSED\_AMDX(3) Manual Page

## Name

VK\_SHADER\_INDEX\_UNUSED\_AMDX - Sentinel for an unused shader index



## [](#_c_specification)C Specification

`VK_SHADER_INDEX_UNUSED_AMDX` is a special shader index used to indicate that the created node does not override the index. In this case, the shader index is determined through other means. It is defined as:

```c++
#define VK_SHADER_INDEX_UNUSED_AMDX       (~0U)
```

## [](#_see_also)See Also

[VK\_AMDX\_shader\_enqueue](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_AMDX_shader_enqueue.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_SHADER_INDEX_UNUSED_AMDX).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0