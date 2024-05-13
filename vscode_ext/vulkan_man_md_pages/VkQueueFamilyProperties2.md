# VkQueueFamilyProperties2(3) Manual Page

## Name

VkQueueFamilyProperties2 - Structure providing information about a queue
family



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkQueueFamilyProperties2` structure is defined as:

``` c
// Provided by VK_VERSION_1_1
typedef struct VkQueueFamilyProperties2 {
    VkStructureType            sType;
    void*                      pNext;
    VkQueueFamilyProperties    queueFamilyProperties;
} VkQueueFamilyProperties2;
```

or the equivalent

``` c
// Provided by VK_KHR_get_physical_device_properties2
typedef VkQueueFamilyProperties2 VkQueueFamilyProperties2KHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `queueFamilyProperties` is a
  [VkQueueFamilyProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueueFamilyProperties.html) structure
  which is populated with the same values as in
  [vkGetPhysicalDeviceQueueFamilyProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceQueueFamilyProperties.html).

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-VkQueueFamilyProperties2-sType-sType"
  id="VUID-VkQueueFamilyProperties2-sType-sType"></a>
  VUID-VkQueueFamilyProperties2-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_QUEUE_FAMILY_PROPERTIES_2`

- <a href="#VUID-VkQueueFamilyProperties2-pNext-pNext"
  id="VUID-VkQueueFamilyProperties2-pNext-pNext"></a>
  VUID-VkQueueFamilyProperties2-pNext-pNext  
  Each `pNext` member of any structure (including this one) in the
  `pNext` chain **must** be either `NULL` or a pointer to a valid
  instance of
  [VkQueueFamilyCheckpointProperties2NV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueueFamilyCheckpointProperties2NV.html),
  [VkQueueFamilyCheckpointPropertiesNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueueFamilyCheckpointPropertiesNV.html),
  [VkQueueFamilyGlobalPriorityPropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueueFamilyGlobalPriorityPropertiesKHR.html),
  [VkQueueFamilyQueryResultStatusPropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueueFamilyQueryResultStatusPropertiesKHR.html),
  or
  [VkQueueFamilyVideoPropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueueFamilyVideoPropertiesKHR.html)

- <a href="#VUID-VkQueueFamilyProperties2-sType-unique"
  id="VUID-VkQueueFamilyProperties2-sType-unique"></a>
  VUID-VkQueueFamilyProperties2-sType-unique  
  The `sType` value of each struct in the `pNext` chain **must** be
  unique

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_1](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_1.html),
[VkQueueFamilyProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueueFamilyProperties.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkGetPhysicalDeviceQueueFamilyProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceQueueFamilyProperties2.html),
[vkGetPhysicalDeviceQueueFamilyProperties2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceQueueFamilyProperties2KHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkQueueFamilyProperties2"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
