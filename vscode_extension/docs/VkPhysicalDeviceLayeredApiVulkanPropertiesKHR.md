# VkPhysicalDeviceLayeredApiVulkanPropertiesKHR(3) Manual Page

## Name

VkPhysicalDeviceLayeredApiVulkanPropertiesKHR - Structure describing
physical device properties of a layered Vulkan implementation underneath
the Vulkan physical device



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPhysicalDeviceLayeredApiVulkanPropertiesKHR` structure is defined
as:

``` c
// Provided by VK_KHR_maintenance7
typedef struct VkPhysicalDeviceLayeredApiVulkanPropertiesKHR {
    VkStructureType                sType;
    void*                          pNext;
    VkPhysicalDeviceProperties2    properties;
} VkPhysicalDeviceLayeredApiVulkanPropertiesKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `properties` is a
  [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties2.html) in
  which properties of the underlying layered Vulkan implementation are
  returned.

## <a href="#_description" class="anchor"></a>Description

The implementation **must** zero-fill the contents of
`properties.properties.limits` and
`properties.properties.sparseProperties`.

Valid Usage

- <a
  href="#VUID-VkPhysicalDeviceLayeredApiVulkanPropertiesKHR-pNext-10011"
  id="VUID-VkPhysicalDeviceLayeredApiVulkanPropertiesKHR-pNext-10011"></a>
  VUID-VkPhysicalDeviceLayeredApiVulkanPropertiesKHR-pNext-10011  
  Only
  [VkPhysicalDeviceDriverProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceDriverProperties.html)
  and [VkPhysicalDeviceIDProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceIDProperties.html)
  are allowed in the `pNext` chain of `properties`

Valid Usage (Implicit)

- <a
  href="#VUID-VkPhysicalDeviceLayeredApiVulkanPropertiesKHR-sType-sType"
  id="VUID-VkPhysicalDeviceLayeredApiVulkanPropertiesKHR-sType-sType"></a>
  VUID-VkPhysicalDeviceLayeredApiVulkanPropertiesKHR-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_LAYERED_API_VULKAN_PROPERTIES_KHR`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_maintenance7](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_maintenance7.html),
[VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties2.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceLayeredApiVulkanPropertiesKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
