# VkPhysicalDeviceMultiDrawPropertiesEXT(3) Manual Page

## Name

VkPhysicalDeviceMultiDrawPropertiesEXT - Structure describing multidraw
limits of an implementation



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPhysicalDeviceMultiDrawPropertiesEXT` structure is defined as:

``` c
// Provided by VK_EXT_multi_draw
typedef struct VkPhysicalDeviceMultiDrawPropertiesEXT {
    VkStructureType    sType;
    void*              pNext;
    uint32_t           maxMultiDrawCount;
} VkPhysicalDeviceMultiDrawPropertiesEXT;
```

## <a href="#_members" class="anchor"></a>Members

The members of the `VkPhysicalDeviceMultiDrawPropertiesEXT` structure
describe the following features:

## <a href="#_description" class="anchor"></a>Description

- <span id="limits-maxMultiDrawCount"></span> `maxMultiDrawCount`
  indicates the maximum number of draw calls which **can** be batched
  into a single multidraw.

If the `VkPhysicalDeviceMultiDrawPropertiesEXT` structure is included in
the `pNext` chain of the
[VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties2.html)
structure passed to
[vkGetPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceProperties2.html),
it is filled in with each corresponding implementation-dependent
property.

Valid Usage (Implicit)

- <a href="#VUID-VkPhysicalDeviceMultiDrawPropertiesEXT-sType-sType"
  id="VUID-VkPhysicalDeviceMultiDrawPropertiesEXT-sType-sType"></a>
  VUID-VkPhysicalDeviceMultiDrawPropertiesEXT-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MULTI_DRAW_PROPERTIES_EXT`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_multi_draw](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_multi_draw.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceMultiDrawPropertiesEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
