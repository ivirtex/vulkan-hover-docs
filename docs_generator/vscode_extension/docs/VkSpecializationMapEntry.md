# VkSpecializationMapEntry(3) Manual Page

## Name

VkSpecializationMapEntry - Structure specifying a specialization map entry



## [](#_c_specification)C Specification

The `VkSpecializationMapEntry` structure is defined as:

```c++
// Provided by VK_VERSION_1_0
typedef struct VkSpecializationMapEntry {
    uint32_t    constantID;
    uint32_t    offset;
    size_t      size;
} VkSpecializationMapEntry;
```

## [](#_members)Members

- `constantID` is the ID of the specialization constant in SPIR-V.
- `offset` is the byte offset of the specialization constant value within the supplied data buffer.
- `size` is the byte size of the specialization constant value within the supplied data buffer.

## [](#_description)Description

If a `constantID` value is not a specialization constant ID used in the shader, that map entry does not affect the behavior of the pipeline.

Valid Usage

- [](#VUID-VkSpecializationMapEntry-constantID-00776)VUID-VkSpecializationMapEntry-constantID-00776  
  For a `constantID` specialization constant declared in a shader, `size` **must** match the byte size of the `constantID`. If the specialization constant is of type `boolean`, `size` **must** be the byte size of [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html)

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkSpecializationInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSpecializationInfo.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkSpecializationMapEntry)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0