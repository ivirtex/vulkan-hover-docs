# VK\_WHOLE\_SIZE(3) Manual Page

## Name

VK\_WHOLE\_SIZE - Sentinel value to use entire remaining array length



## [](#_c_specification)C Specification

`VK_WHOLE_SIZE` is a special value indicating that the entire remaining length of a buffer following a given `offset` should be used. It **can** be specified for [VkBufferMemoryBarrier](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferMemoryBarrier.html)::`size` and other structures.

```c++
#define VK_WHOLE_SIZE                     (~0ULL)
```

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_WHOLE_SIZE)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0