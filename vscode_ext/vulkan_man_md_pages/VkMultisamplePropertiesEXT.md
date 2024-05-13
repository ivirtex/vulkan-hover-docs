# VkMultisamplePropertiesEXT(3) Manual Page

## Name

VkMultisamplePropertiesEXT - Structure returning information about
sample count specific additional multisampling capabilities



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkMultisamplePropertiesEXT` structure is defined as

``` c
// Provided by VK_EXT_sample_locations
typedef struct VkMultisamplePropertiesEXT {
    VkStructureType    sType;
    void*              pNext;
    VkExtent2D         maxSampleLocationGridSize;
} VkMultisamplePropertiesEXT;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `maxSampleLocationGridSize` is the maximum size of the pixel grid in
  which sample locations **can** vary.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-VkMultisamplePropertiesEXT-sType-sType"
  id="VUID-VkMultisamplePropertiesEXT-sType-sType"></a>
  VUID-VkMultisamplePropertiesEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_MULTISAMPLE_PROPERTIES_EXT`

- <a href="#VUID-VkMultisamplePropertiesEXT-pNext-pNext"
  id="VUID-VkMultisamplePropertiesEXT-pNext-pNext"></a>
  VUID-VkMultisamplePropertiesEXT-pNext-pNext  
  `pNext` **must** be `NULL`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_sample_locations](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_sample_locations.html),
[VkExtent2D](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExtent2D.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkGetPhysicalDeviceMultisamplePropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceMultisamplePropertiesEXT.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkMultisamplePropertiesEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
