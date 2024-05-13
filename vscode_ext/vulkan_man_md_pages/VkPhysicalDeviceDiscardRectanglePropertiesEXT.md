# VkPhysicalDeviceDiscardRectanglePropertiesEXT(3) Manual Page

## Name

VkPhysicalDeviceDiscardRectanglePropertiesEXT - Structure describing
discard rectangle limits that can be supported by an implementation



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPhysicalDeviceDiscardRectanglePropertiesEXT` structure is defined
as:

``` c
// Provided by VK_EXT_discard_rectangles
typedef struct VkPhysicalDeviceDiscardRectanglePropertiesEXT {
    VkStructureType    sType;
    void*              pNext;
    uint32_t           maxDiscardRectangles;
} VkPhysicalDeviceDiscardRectanglePropertiesEXT;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- <span id="limits-maxDiscardRectangles"></span> `maxDiscardRectangles`
  is the maximum number of active discard rectangles that **can** be
  specified.

## <a href="#_description" class="anchor"></a>Description

If the `VkPhysicalDeviceDiscardRectanglePropertiesEXT` structure is
included in the `pNext` chain of the
[VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties2.html)
structure passed to
[vkGetPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceProperties2.html),
it is filled in with each corresponding implementation-dependent
property.

Valid Usage (Implicit)

- <a
  href="#VUID-VkPhysicalDeviceDiscardRectanglePropertiesEXT-sType-sType"
  id="VUID-VkPhysicalDeviceDiscardRectanglePropertiesEXT-sType-sType"></a>
  VUID-VkPhysicalDeviceDiscardRectanglePropertiesEXT-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_DISCARD_RECTANGLE_PROPERTIES_EXT`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_discard_rectangles](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_discard_rectangles.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceDiscardRectanglePropertiesEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
