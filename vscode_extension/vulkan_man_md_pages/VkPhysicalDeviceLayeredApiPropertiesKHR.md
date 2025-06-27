# VkPhysicalDeviceLayeredApiPropertiesKHR(3) Manual Page

## Name

VkPhysicalDeviceLayeredApiPropertiesKHR - Structure describing a single
layered implementation underneath the Vulkan physical device



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPhysicalDeviceLayeredApiPropertiesKHR` structure is defined as:

``` c
// Provided by VK_KHR_maintenance7
typedef struct VkPhysicalDeviceLayeredApiPropertiesKHR {
    VkStructureType                  sType;
    void*                            pNext;
    uint32_t                         vendorID;
    uint32_t                         deviceID;
    VkPhysicalDeviceLayeredApiKHR    layeredAPI;
    char                             deviceName[VK_MAX_PHYSICAL_DEVICE_NAME_SIZE];
} VkPhysicalDeviceLayeredApiPropertiesKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `vendorID` is a unique identifier for the vendor of the layered
  implementation.

- `deviceID` is a unique identifier for the layered implementation among
  devices available from the vendor.

- `layeredAPI` is a
  [VkPhysicalDeviceLayeredApiKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceLayeredApiKHR.html)
  specifying the API implemented by the layered implementation.

- `deviceName` is an array of `VK_MAX_PHYSICAL_DEVICE_NAME_SIZE` `char`
  containing a null-terminated UTF-8 string which is the name of the
  device.

## <a href="#_description" class="anchor"></a>Description

If `layeredAPI` is `VK_PHYSICAL_DEVICE_LAYERED_API_VULKAN_KHR`,
additional Vulkan-specific information can be queried by including the
[VkPhysicalDeviceLayeredApiVulkanPropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceLayeredApiVulkanPropertiesKHR.html)
structure in the `pNext` chain. Otherwise if such a structure is
included in the `pNext` chain, it is ignored.

Valid Usage (Implicit)

- <a href="#VUID-VkPhysicalDeviceLayeredApiPropertiesKHR-sType-sType"
  id="VUID-VkPhysicalDeviceLayeredApiPropertiesKHR-sType-sType"></a>
  VUID-VkPhysicalDeviceLayeredApiPropertiesKHR-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_LAYERED_API_PROPERTIES_KHR`

- <a href="#VUID-VkPhysicalDeviceLayeredApiPropertiesKHR-pNext-pNext"
  id="VUID-VkPhysicalDeviceLayeredApiPropertiesKHR-pNext-pNext"></a>
  VUID-VkPhysicalDeviceLayeredApiPropertiesKHR-pNext-pNext  
  `pNext` **must** be `NULL` or a pointer to a valid instance of
  [VkPhysicalDeviceLayeredApiVulkanPropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceLayeredApiVulkanPropertiesKHR.html)

- <a href="#VUID-VkPhysicalDeviceLayeredApiPropertiesKHR-sType-unique"
  id="VUID-VkPhysicalDeviceLayeredApiPropertiesKHR-sType-unique"></a>
  VUID-VkPhysicalDeviceLayeredApiPropertiesKHR-sType-unique  
  The `sType` value of each struct in the `pNext` chain **must** be
  unique

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_maintenance7](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_maintenance7.html),
[VkPhysicalDeviceLayeredApiKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceLayeredApiKHR.html),
[VkPhysicalDeviceLayeredApiPropertiesListKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceLayeredApiPropertiesListKHR.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceLayeredApiPropertiesKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
