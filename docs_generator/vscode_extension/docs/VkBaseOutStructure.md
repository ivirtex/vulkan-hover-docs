# VkBaseOutStructure(3) Manual Page

## Name

VkBaseOutStructure - Base structure for a read-only pointer chain



## [](#_c_specification)C Specification

The `VkBaseOutStructure` structure is defined as:

```c++
// Provided by VK_VERSION_1_0
typedef struct VkBaseOutStructure {
    VkStructureType               sType;
    struct VkBaseOutStructure*    pNext;
} VkBaseOutStructure;
```

## [](#_members)Members

- `sType` is the structure type of the structure being iterated through.
- `pNext` is `NULL` or a pointer to the next structure in a structure chain.

## [](#_description)Description

`VkBaseOutStructure` can be used to facilitate iterating through a structure pointer chain that returns data back to the application.

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkBaseOutStructure](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBaseOutStructure.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkGetPipelinePropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPipelinePropertiesEXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkBaseOutStructure)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0