# VkPhysicalDeviceOpacityMicromapPropertiesEXT(3) Manual Page

## Name

VkPhysicalDeviceOpacityMicromapPropertiesEXT - Structure describing the
opacity micromap properties of a physical device



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPhysicalDeviceOpacityMicromapPropertiesEXT` structure is defined
as:

``` c
// Provided by VK_EXT_opacity_micromap
typedef struct VkPhysicalDeviceOpacityMicromapPropertiesEXT {
    VkStructureType    sType;
    void*              pNext;
    uint32_t           maxOpacity2StateSubdivisionLevel;
    uint32_t           maxOpacity4StateSubdivisionLevel;
} VkPhysicalDeviceOpacityMicromapPropertiesEXT;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `maxOpacity2StateSubdivisionLevel` is the maximum allowed
  `subdivisionLevel` when `format` is
  `VK_OPACITY_MICROMAP_FORMAT_2_STATE_EXT`

- `maxOpacity4StateSubdivisionLevel` is the maximum allowed
  `subdivisionLevel` when `format` is
  `VK_OPACITY_MICROMAP_FORMAT_4_STATE_EXT`

## <a href="#_description" class="anchor"></a>Description

If the `VkPhysicalDeviceOpacityMicromapPropertiesEXT` structure is
included in the `pNext` chain of the
[VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties2.html)
structure passed to
[vkGetPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceProperties2.html),
it is filled in with each corresponding implementation-dependent
property.

Valid Usage (Implicit)

- <a href="#VUID-VkPhysicalDeviceOpacityMicromapPropertiesEXT-sType-sType"
  id="VUID-VkPhysicalDeviceOpacityMicromapPropertiesEXT-sType-sType"></a>
  VUID-VkPhysicalDeviceOpacityMicromapPropertiesEXT-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_OPACITY_MICROMAP_PROPERTIES_EXT`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_opacity_micromap](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_opacity_micromap.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceOpacityMicromapPropertiesEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
