# VkBaseInStructure(3) Manual Page

## Name

VkBaseInStructure - Base structure for a read-only pointer chain



## [](#_c_specification)C Specification

The `VkBaseInStructure` structure is defined as:

```c++
// Provided by VK_VERSION_1_0
typedef struct VkBaseInStructure {
    VkStructureType                    sType;
    const struct VkBaseInStructure*    pNext;
} VkBaseInStructure;
```

## [](#_members)Members

- `sType` is the structure type of the structure being iterated through.
- `pNext` is `NULL` or a pointer to the next structure in a structure chain.

## [](#_description)Description

`VkBaseInStructure` can be used to facilitate iterating through a read-only structure pointer chain.

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkBaseInStructure](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBaseInStructure.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkBaseInStructure)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0