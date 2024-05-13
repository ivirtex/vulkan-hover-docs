# VkPhysicalDeviceDisplacementMicromapPropertiesNV(3) Manual Page

## Name

VkPhysicalDeviceDisplacementMicromapPropertiesNV - Structure describing
the displacement micromap properties of a physical device



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPhysicalDeviceDisplacementMicromapPropertiesNV` structure is
defined as:

``` c
// Provided by VK_NV_displacement_micromap
typedef struct VkPhysicalDeviceDisplacementMicromapPropertiesNV {
    VkStructureType    sType;
    void*              pNext;
    uint32_t           maxDisplacementMicromapSubdivisionLevel;
} VkPhysicalDeviceDisplacementMicromapPropertiesNV;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `maxDisplacementMicromapSubdivisionLevel` is the maximum allowed
  `subdivisionLevel` for displacement micromaps.

## <a href="#_description" class="anchor"></a>Description

If the `VkPhysicalDeviceDisplacementMicromapPropertiesNV` structure is
included in the `pNext` chain of the
[VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties2.html)
structure passed to
[vkGetPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceProperties2.html),
it is filled in with each corresponding implementation-dependent
property.

Valid Usage (Implicit)

- <a
  href="#VUID-VkPhysicalDeviceDisplacementMicromapPropertiesNV-sType-sType"
  id="VUID-VkPhysicalDeviceDisplacementMicromapPropertiesNV-sType-sType"></a>
  VUID-VkPhysicalDeviceDisplacementMicromapPropertiesNV-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_DISPLACEMENT_MICROMAP_PROPERTIES_NV`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NV_displacement_micromap](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_displacement_micromap.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceDisplacementMicromapPropertiesNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
