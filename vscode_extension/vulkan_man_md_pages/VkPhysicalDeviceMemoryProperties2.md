# VkPhysicalDeviceMemoryProperties2(3) Manual Page

## Name

VkPhysicalDeviceMemoryProperties2 - Structure specifying physical device
memory properties



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPhysicalDeviceMemoryProperties2` structure is defined as:

``` c
// Provided by VK_VERSION_1_1
typedef struct VkPhysicalDeviceMemoryProperties2 {
    VkStructureType                     sType;
    void*                               pNext;
    VkPhysicalDeviceMemoryProperties    memoryProperties;
} VkPhysicalDeviceMemoryProperties2;
```

or the equivalent

``` c
// Provided by VK_KHR_get_physical_device_properties2
typedef VkPhysicalDeviceMemoryProperties2 VkPhysicalDeviceMemoryProperties2KHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `memoryProperties` is a
  [VkPhysicalDeviceMemoryProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceMemoryProperties.html)
  structure which is populated with the same values as in
  [vkGetPhysicalDeviceMemoryProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceMemoryProperties.html).

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-VkPhysicalDeviceMemoryProperties2-sType-sType"
  id="VUID-VkPhysicalDeviceMemoryProperties2-sType-sType"></a>
  VUID-VkPhysicalDeviceMemoryProperties2-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MEMORY_PROPERTIES_2`

- <a href="#VUID-VkPhysicalDeviceMemoryProperties2-pNext-pNext"
  id="VUID-VkPhysicalDeviceMemoryProperties2-pNext-pNext"></a>
  VUID-VkPhysicalDeviceMemoryProperties2-pNext-pNext  
  `pNext` **must** be `NULL` or a pointer to a valid instance of
  [VkPhysicalDeviceMemoryBudgetPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceMemoryBudgetPropertiesEXT.html)

- <a href="#VUID-VkPhysicalDeviceMemoryProperties2-sType-unique"
  id="VUID-VkPhysicalDeviceMemoryProperties2-sType-unique"></a>
  VUID-VkPhysicalDeviceMemoryProperties2-sType-unique  
  The `sType` value of each struct in the `pNext` chain **must** be
  unique

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_1](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_1.html),
[VkPhysicalDeviceMemoryProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceMemoryProperties.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkGetPhysicalDeviceMemoryProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceMemoryProperties2.html),
[vkGetPhysicalDeviceMemoryProperties2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceMemoryProperties2KHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceMemoryProperties2"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
