# VkSpecializationMapEntry(3) Manual Page

## Name

VkSpecializationMapEntry - Structure specifying a specialization map
entry



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkSpecializationMapEntry` structure is defined as:

``` c
// Provided by VK_VERSION_1_0
typedef struct VkSpecializationMapEntry {
    uint32_t    constantID;
    uint32_t    offset;
    size_t      size;
} VkSpecializationMapEntry;
```

## <a href="#_members" class="anchor"></a>Members

- `constantID` is the ID of the specialization constant in SPIR-V.

- `offset` is the byte offset of the specialization constant value
  within the supplied data buffer.

- `size` is the byte size of the specialization constant value within
  the supplied data buffer.

## <a href="#_description" class="anchor"></a>Description

If a `constantID` value is not a specialization constant ID used in the
shader, that map entry does not affect the behavior of the pipeline.

Valid Usage

- <a href="#VUID-VkSpecializationMapEntry-constantID-00776"
  id="VUID-VkSpecializationMapEntry-constantID-00776"></a>
  VUID-VkSpecializationMapEntry-constantID-00776  
  For a `constantID` specialization constant declared in a shader,
  `size` **must** match the byte size of the `constantID`. If the
  specialization constant is of type `boolean`, `size` **must** be the
  byte size of [VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html)

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkSpecializationInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSpecializationInfo.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkSpecializationMapEntry"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
