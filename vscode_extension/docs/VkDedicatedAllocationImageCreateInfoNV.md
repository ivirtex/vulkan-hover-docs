# VkDedicatedAllocationImageCreateInfoNV(3) Manual Page

## Name

VkDedicatedAllocationImageCreateInfoNV - Specify that an image is bound to a dedicated memory resource



## [](#_c_specification)C Specification

If the `pNext` chain includes a `VkDedicatedAllocationImageCreateInfoNV` structure, then that structure includes an enable controlling whether the image will have a dedicated memory allocation bound to it.

The `VkDedicatedAllocationImageCreateInfoNV` structure is defined as:

```c++
// Provided by VK_NV_dedicated_allocation
typedef struct VkDedicatedAllocationImageCreateInfoNV {
    VkStructureType    sType;
    const void*        pNext;
    VkBool32           dedicatedAllocation;
} VkDedicatedAllocationImageCreateInfoNV;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `dedicatedAllocation` specifies whether the image will have a dedicated allocation bound to it.

## [](#_description)Description

Note

Using a dedicated allocation for color and depth/stencil attachments or other large images **may** improve performance on some devices.

Valid Usage

- [](#VUID-VkDedicatedAllocationImageCreateInfoNV-dedicatedAllocation-00994)VUID-VkDedicatedAllocationImageCreateInfoNV-dedicatedAllocation-00994  
  If `dedicatedAllocation` is `VK_TRUE`, [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateInfo.html)::`flags` **must** not include `VK_IMAGE_CREATE_SPARSE_BINDING_BIT`, `VK_IMAGE_CREATE_SPARSE_RESIDENCY_BIT`, or `VK_IMAGE_CREATE_SPARSE_ALIASED_BIT`

Valid Usage (Implicit)

- [](#VUID-VkDedicatedAllocationImageCreateInfoNV-sType-sType)VUID-VkDedicatedAllocationImageCreateInfoNV-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_DEDICATED_ALLOCATION_IMAGE_CREATE_INFO_NV`

## [](#_see_also)See Also

[VK\_NV\_dedicated\_allocation](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_dedicated_allocation.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkDedicatedAllocationImageCreateInfoNV)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0