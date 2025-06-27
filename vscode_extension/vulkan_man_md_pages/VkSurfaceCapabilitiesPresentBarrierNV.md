# VkSurfaceCapabilitiesPresentBarrierNV(3) Manual Page

## Name

VkSurfaceCapabilitiesPresentBarrierNV - Structure describing present
barrier capabilities of a surface



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkSurfaceCapabilitiesPresentBarrierNV` structure is defined as:

``` c
// Provided by VK_NV_present_barrier
typedef struct VkSurfaceCapabilitiesPresentBarrierNV {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           presentBarrierSupported;
} VkSurfaceCapabilitiesPresentBarrierNV;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `presentBarrierSupported` is a boolean describing whether the surface
  is able to make use of the present barrier feature.

## <a href="#_description" class="anchor"></a>Description

This structure **can** be included in the `pNext` chain of
[VkSurfaceCapabilities2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceCapabilities2KHR.html) to determine
support for present barrier access. If `presentBarrierSupported` is
`VK_FALSE`, it indicates that the present barrier feature is not
obtainable for this surface.

Valid Usage (Implicit)

- <a href="#VUID-VkSurfaceCapabilitiesPresentBarrierNV-sType-sType"
  id="VUID-VkSurfaceCapabilitiesPresentBarrierNV-sType-sType"></a>
  VUID-VkSurfaceCapabilitiesPresentBarrierNV-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_SURFACE_CAPABILITIES_PRESENT_BARRIER_NV`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NV_present_barrier](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_present_barrier.html),
[VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkSurfaceCapabilitiesPresentBarrierNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
