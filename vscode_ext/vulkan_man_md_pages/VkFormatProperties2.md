# VkFormatProperties2(3) Manual Page

## Name

VkFormatProperties2 - Structure specifying image format properties



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkFormatProperties2` structure is defined as:

``` c
// Provided by VK_VERSION_1_1
typedef struct VkFormatProperties2 {
    VkStructureType       sType;
    void*                 pNext;
    VkFormatProperties    formatProperties;
} VkFormatProperties2;
```

or the equivalent

``` c
// Provided by VK_KHR_get_physical_device_properties2
typedef VkFormatProperties2 VkFormatProperties2KHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `formatProperties` is a [VkFormatProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormatProperties.html)
  structure describing features supported by the requested format.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-VkFormatProperties2-sType-sType"
  id="VUID-VkFormatProperties2-sType-sType"></a>
  VUID-VkFormatProperties2-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_FORMAT_PROPERTIES_2`

- <a href="#VUID-VkFormatProperties2-pNext-pNext"
  id="VUID-VkFormatProperties2-pNext-pNext"></a>
  VUID-VkFormatProperties2-pNext-pNext  
  Each `pNext` member of any structure (including this one) in the
  `pNext` chain **must** be either `NULL` or a pointer to a valid
  instance of
  [VkDrmFormatModifierPropertiesList2EXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDrmFormatModifierPropertiesList2EXT.html),
  [VkDrmFormatModifierPropertiesListEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDrmFormatModifierPropertiesListEXT.html),
  [VkFormatProperties3](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormatProperties3.html), or
  [VkSubpassResolvePerformanceQueryEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubpassResolvePerformanceQueryEXT.html)

- <a href="#VUID-VkFormatProperties2-sType-unique"
  id="VUID-VkFormatProperties2-sType-unique"></a>
  VUID-VkFormatProperties2-sType-unique  
  The `sType` value of each struct in the `pNext` chain **must** be
  unique

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_1](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_1.html),
[VkFormatProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormatProperties.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkGetPhysicalDeviceFormatProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceFormatProperties2.html),
[vkGetPhysicalDeviceFormatProperties2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceFormatProperties2KHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkFormatProperties2"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
