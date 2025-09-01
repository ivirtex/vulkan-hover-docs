# VkStreamDescriptorSurfaceCreateInfoGGP(3) Manual Page

## Name

VkStreamDescriptorSurfaceCreateInfoGGP - Structure specifying parameters of a newly created Google Games Platform stream surface object



## [](#_c_specification)C Specification

The `VkStreamDescriptorSurfaceCreateInfoGGP` structure is defined as:

```c++
// Provided by VK_GGP_stream_descriptor_surface
typedef struct VkStreamDescriptorSurfaceCreateInfoGGP {
    VkStructureType                            sType;
    const void*                                pNext;
    VkStreamDescriptorSurfaceCreateFlagsGGP    flags;
    GgpStreamDescriptor                        streamDescriptor;
} VkStreamDescriptorSurfaceCreateInfoGGP;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `flags` is reserved for future use.
- `streamDescriptor` is a `GgpStreamDescriptor` referring to the GGP stream descriptor to associate with the surface.

## [](#_description)Description

Valid Usage

- [](#VUID-VkStreamDescriptorSurfaceCreateInfoGGP-streamDescriptor-02681)VUID-VkStreamDescriptorSurfaceCreateInfoGGP-streamDescriptor-02681  
  `streamDescriptor` **must** be a valid `GgpStreamDescriptor`

Valid Usage (Implicit)

- [](#VUID-VkStreamDescriptorSurfaceCreateInfoGGP-sType-sType)VUID-VkStreamDescriptorSurfaceCreateInfoGGP-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_STREAM_DESCRIPTOR_SURFACE_CREATE_INFO_GGP`
- [](#VUID-VkStreamDescriptorSurfaceCreateInfoGGP-pNext-pNext)VUID-VkStreamDescriptorSurfaceCreateInfoGGP-pNext-pNext  
  `pNext` **must** be `NULL`
- [](#VUID-VkStreamDescriptorSurfaceCreateInfoGGP-flags-zerobitmask)VUID-VkStreamDescriptorSurfaceCreateInfoGGP-flags-zerobitmask  
  `flags` **must** be `0`

## [](#_see_also)See Also

[VK\_GGP\_stream\_descriptor\_surface](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_GGP_stream_descriptor_surface.html), [VkStreamDescriptorSurfaceCreateFlagsGGP](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStreamDescriptorSurfaceCreateFlagsGGP.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkCreateStreamDescriptorSurfaceGGP](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateStreamDescriptorSurfaceGGP.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkStreamDescriptorSurfaceCreateInfoGGP).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0