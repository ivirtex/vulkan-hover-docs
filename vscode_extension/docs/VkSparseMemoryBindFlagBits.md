# VkSparseMemoryBindFlagBits(3) Manual Page

## Name

VkSparseMemoryBindFlagBits - Bitmask specifying usage of a sparse memory binding operation



## [](#_c_specification)C Specification

Bits which **can** be set in [VkSparseMemoryBind](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSparseMemoryBind.html)::`flags`, specifying usage of a sparse memory binding operation, are:

```c++
// Provided by VK_VERSION_1_0
typedef enum VkSparseMemoryBindFlagBits {
    VK_SPARSE_MEMORY_BIND_METADATA_BIT = 0x00000001,
} VkSparseMemoryBindFlagBits;
```

## [](#_description)Description

- `VK_SPARSE_MEMORY_BIND_METADATA_BIT` specifies that the memory being bound is only for the metadata aspect.

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkSparseMemoryBindFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSparseMemoryBindFlags.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkSparseMemoryBindFlagBits).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0