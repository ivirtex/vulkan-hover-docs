# VkStreamDescriptorSurfaceCreateInfoGGP(3) Manual Page

## Name

VkStreamDescriptorSurfaceCreateInfoGGP - Structure specifying parameters
of a newly created Google Games Platform stream surface object



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkStreamDescriptorSurfaceCreateInfoGGP` structure is defined as:

``` c
// Provided by VK_GGP_stream_descriptor_surface
typedef struct VkStreamDescriptorSurfaceCreateInfoGGP {
    VkStructureType                            sType;
    const void*                                pNext;
    VkStreamDescriptorSurfaceCreateFlagsGGP    flags;
    GgpStreamDescriptor                        streamDescriptor;
} VkStreamDescriptorSurfaceCreateInfoGGP;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `flags` is reserved for future use.

- `streamDescriptor` is a `GgpStreamDescriptor` referring to the GGP
  stream descriptor to associate with the surface.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a
  href="#VUID-VkStreamDescriptorSurfaceCreateInfoGGP-streamDescriptor-02681"
  id="VUID-VkStreamDescriptorSurfaceCreateInfoGGP-streamDescriptor-02681"></a>
  VUID-VkStreamDescriptorSurfaceCreateInfoGGP-streamDescriptor-02681  
  `streamDescriptor` **must** be a valid `GgpStreamDescriptor`

Valid Usage (Implicit)

- <a href="#VUID-VkStreamDescriptorSurfaceCreateInfoGGP-sType-sType"
  id="VUID-VkStreamDescriptorSurfaceCreateInfoGGP-sType-sType"></a>
  VUID-VkStreamDescriptorSurfaceCreateInfoGGP-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_STREAM_DESCRIPTOR_SURFACE_CREATE_INFO_GGP`

- <a href="#VUID-VkStreamDescriptorSurfaceCreateInfoGGP-pNext-pNext"
  id="VUID-VkStreamDescriptorSurfaceCreateInfoGGP-pNext-pNext"></a>
  VUID-VkStreamDescriptorSurfaceCreateInfoGGP-pNext-pNext  
  `pNext` **must** be `NULL`

- <a href="#VUID-VkStreamDescriptorSurfaceCreateInfoGGP-flags-zerobitmask"
  id="VUID-VkStreamDescriptorSurfaceCreateInfoGGP-flags-zerobitmask"></a>
  VUID-VkStreamDescriptorSurfaceCreateInfoGGP-flags-zerobitmask  
  `flags` **must** be `0`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_GGP_stream_descriptor_surface](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_GGP_stream_descriptor_surface.html),
[VkStreamDescriptorSurfaceCreateFlagsGGP](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStreamDescriptorSurfaceCreateFlagsGGP.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkCreateStreamDescriptorSurfaceGGP](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateStreamDescriptorSurfaceGGP.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkStreamDescriptorSurfaceCreateInfoGGP"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
